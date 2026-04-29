package webhook_test

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/sdk"
	"github.com/hakur/airwallex/webhook"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testClient *airwallex.Client

func TestMain(m *testing.M) {
	_ = godotenv.Load(sdk.ResolveEnvPath())
	var err error
	testClient, err = airwallex.NewFromEnv("", sdk.WithBaseURL(sdk.SandboxURL), sdk.WithDebug(true))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create test client: %v\n", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

// generateTestSignature 生成 Airwallex 官方格式的 webhook 签名。
// 格式: digest = HMAC-SHA256(secret, timestamp_string + payload_string)，返回 hex string。
func generateTestSignature(payload []byte, secret string, tsMillis int64) string {
	valueToDigest := strconv.FormatInt(tsMillis, 10) + string(payload)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(valueToDigest))
	return hex.EncodeToString(mac.Sum(nil))
}

// ─── 签名验证测试 ───

func TestVerifySignatureValid(t *testing.T) {
	payload := []byte(`{"event":"payment_intent.succeeded","data":{"id":"pi_123"}}`)
	secret := "whsec_test_secret"
	now := time.Now().UnixMilli()
	sig := generateTestSignature(payload, secret, now)
	require.NoError(t, webhook.VerifySignature(payload, strconv.FormatInt(now, 10), sig, secret), "valid signature should pass")
}

func TestVerifySignatureInvalid(t *testing.T) {
	payload := []byte(`{"event":"payment_intent.succeeded","data":{"id":"pi_123"}}`)
	secret := "whsec_test_secret"
	now := time.Now().UnixMilli()
	require.Error(t, webhook.VerifySignature(payload, strconv.FormatInt(now, 10), "invalid_sig", secret), "invalid signature should fail")
}

func TestVerifySignatureFutureTimestamp(t *testing.T) {
	payload := []byte(`{"event":"payment_intent.succeeded","data":{"id":"pi_123"}}`)
	secret := "whsec_test_secret"
	future := time.Now().UnixMilli() + 400*1000 // 400 秒后
	sig := generateTestSignature(payload, secret, future)
	require.Error(t, webhook.VerifySignature(payload, strconv.FormatInt(future, 10), sig, secret), "future timestamp should fail")
}

func TestVerifySignaturePastTimestamp(t *testing.T) {
	payload := []byte(`{"event":"payment_intent.succeeded","data":{"id":"pi_123"}}`)
	secret := "whsec_test_secret"
	past := time.Now().UnixMilli() - 400*1000 // 400 秒前
	sig := generateTestSignature(payload, secret, past)
	require.Error(t, webhook.VerifySignature(payload, strconv.FormatInt(past, 10), sig, secret), "past timestamp should fail")
}

func TestVerifySignatureCustomTolerance(t *testing.T) {
	payload := []byte(`{"event":"payment_intent.succeeded","data":{"id":"pi_123"}}`)
	secret := "whsec_test_secret"
	past := time.Now().UnixMilli() - 400*1000 // 400 秒前，但在 600 秒容差内
	sig := generateTestSignature(payload, secret, past)
	require.NoError(t, webhook.VerifySignatureWithTolerance(payload, strconv.FormatInt(past, 10), sig, secret, 600*1000), "custom tolerance should pass")
}

func TestVerifySignatureWrongSecret(t *testing.T) {
	payload := []byte(`{"event":"payment_intent.succeeded","data":{"id":"pi_123"}}`)
	now := time.Now().UnixMilli()
	sig := generateTestSignature(payload, "wrong_secret", now)
	require.Error(t, webhook.VerifySignature(payload, strconv.FormatInt(now, 10), sig, "whsec_test_secret"), "wrong secret should fail")
}

func TestVerifySignatureInvalidTimestamp(t *testing.T) {
	payload := []byte(`{"event":"payment_intent.succeeded","data":{"id":"pi_123"}}`)
	secret := "whsec_test_secret"
	require.Error(t, webhook.VerifySignature(payload, "not_a_number", "any_signature", secret), "invalid timestamp format should fail")
}

// ─── 事件解析测试 ───

func TestParseEvent(t *testing.T) {
	payload := []byte(`{
		"id": "evt_123",
		"event": "payment_intent.succeeded",
		"data": {"id": "pi_123", "status": "SUCCEEDED"},
		"created_at": "2024-01-15T08:30:00Z"
	}`)

	evt, err := webhook.ParseEvent(payload)
	require.NoError(t, err, "parse event failed")
	assert.Equal(t, "evt_123", evt.ID)
	assert.Equal(t, "payment_intent.succeeded", evt.Event)
	assert.Equal(t, "2024-01-15T08:30:00Z", evt.CreatedAt)
}

func TestParseEventInvalidJSON(t *testing.T) {
	_, err := webhook.ParseEvent([]byte("not json"))
	require.Error(t, err, "invalid json should fail")
}

func TestUnmarshalData(t *testing.T) {
	payload := []byte(`{
		"id": "evt_123",
		"event": "payment_intent.succeeded",
		"data": {"id": "pi_123", "status": "SUCCEEDED", "amount": 100.0, "currency": "USD"},
		"created_at": "2024-01-15T08:30:00Z"
	}`)

	evt, err := webhook.ParseEvent(payload)
	require.NoError(t, err, "parse event failed")

	type PaymentIntentData struct {
		ID       string  `json:"id"`
		Status   string  `json:"status"`
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	}
	data, err := webhook.UnmarshalData[PaymentIntentData](evt)
	require.NoError(t, err, "unmarshal data failed")
	assert.Equal(t, "pi_123", data.ID)
	assert.Equal(t, "SUCCEEDED", data.Status)
	assert.Equal(t, 100.0, data.Amount)
	assert.Equal(t, "USD", data.Currency)
}

// ─── Webhook 端点生命周期测试 ───

func TestWebhookLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Webhook()

	list, err := svc.ListWebhooks(ctx, &webhook.ListWebhooksRequest{PageSize: 10})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "list webhooks failed: %v", err)
		return
	}
	t.Logf("existing webhooks: %d", len(list.Items))

	created, err := svc.CreateWebhook(ctx, &webhook.CreateWebhookRequest{
		RequestID: "wh-req-" + time.Now().Format("20060102150405"),
		URL:       "https://example.com/webhook",
		Version:   "2022-11-11",
		Events:    []string{"payment_intent.succeeded"},
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "create webhook failed: %v", err)
		return
	}
	t.Logf("created webhook: %s", created.ID)

	fetched, err := svc.GetWebhook(ctx, created.ID)
	require.NoError(t, err, "get webhook failed")
	assert.Equal(t, created.ID, fetched.ID, "webhook id mismatch")

	updated, err := svc.UpdateWebhook(ctx, created.ID, &webhook.UpdateWebhookRequest{
		URL: "https://example.com/webhook-updated",
	})
	require.NoError(t, err, "update webhook failed")
	t.Logf("updated webhook url: %s", updated.URL)

	deleted, err := svc.DeleteWebhook(ctx, created.ID)
	require.NoError(t, err, "delete webhook failed")
	assert.True(t, deleted.Deleted, "deleted flag should be true")
	assert.Equal(t, created.ID, deleted.ID, "deleted id mismatch")
}
