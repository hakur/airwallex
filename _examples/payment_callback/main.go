package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/pa"
	"github.com/hakur/airwallex/sdk"
	"github.com/hakur/airwallex/webhook"
)

// 支付成功回调通知完整示例
// 流程：创建 PaymentIntent → 确认 → 捕获 → 模拟 webhook 回调验证

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	client, err := airwallex.NewFromEnv("../../.env",
		sdk.WithBaseURL(sdk.SandboxURL),
		sdk.WithDebug(true),
	)
	if err != nil {
		slog.Error("创建客户端失败", "error", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	svc := client.PA()

	// ========================================
	// 阶段一：创建并完成支付（真实 Sandbox API）
	// ========================================

	// 1. 创建 PaymentIntent
	slog.Info("=== 1. 创建 PaymentIntent ===")
	pi, err := svc.CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
		RequestID: "pi-cb-" + time.Now().Format("20060102150405"),
		Amount:    100.00,
		Currency:  sdk.CurrencyUSD,
	})
	if err != nil {
		slog.Error("创建 PaymentIntent 失败", "error", err)
		os.Exit(1)
	}
	slog.Info("PaymentIntent 已创建", "id", pi.ID, "status", pi.Status)

	// 2. 确认支付（触发扣款）
	slog.Info("=== 2. 确认支付 ===")
	confirmed, err := svc.ConfirmPaymentIntent(ctx, pi.ID, &pa.ConfirmPaymentIntentRequest{
		RequestID: "pi-confirm-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		slog.Error("确认支付失败", "error", err)
		os.Exit(1)
	}
	slog.Info("支付已确认", "id", confirmed.ID, "status", confirmed.Status)

	// 3. 捕获支付（最终完成）
	time.Sleep(500 * time.Millisecond) // 等待支付状态稳定
	slog.Info("=== 3. 捕获支付 ===")
	captured, err := svc.CapturePaymentIntent(ctx, pi.ID, &pa.CapturePaymentIntentRequest{
		RequestID: "pi-capture-" + time.Now().Format("20060102150405"),
		Amount:    100.00,
	})
	if err != nil {
		slog.Error("捕获支付失败", "error", err)
		os.Exit(1)
	}
	slog.Info("支付已捕获", "id", captured.ID, "status", captured.Status)

	// 4. 取消 PaymentIntent（清理沙箱数据）
	slog.Info("=== 4. 清理：取消 PaymentIntent ===")
	_, err = svc.CancelPaymentIntent(ctx, pi.ID, &pa.CancelPaymentIntentRequest{
		RequestID: "pi-cancel-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		slog.Warn("取消 PaymentIntent 失败（不影响主流程）", "error", err)
	} else {
		slog.Info("PaymentIntent 已取消（沙箱数据已清理）")
	}

	// ========================================
	// 阶段二：演示 Webhook 回调处理（本地验证）
	// ========================================

	slog.Info("")
	slog.Info("=== 5. 演示 Webhook 回调验证 ===")

	// 模拟 Airwallex 发送的 webhook 回调负载
	webhookSecret := "whsec_your_webhook_secret_here"
	samplePayload := demoWebhookPayload(pi.ID, pi.Amount, string(pi.Currency))

	// 计算签名（与 Airwallex 官方算法一致）
	timestampMs := fmt.Sprintf("%d", time.Now().UnixMilli())
	signature := computeHMACSignature(timestampMs, samplePayload, webhookSecret)

	slog.Info("模拟 webhook 回调",
		"x-timestamp", timestampMs,
		"x-signature", signature[:16]+"...",
	)

	// 验证签名
	if err := webhook.VerifySignature(samplePayload, timestampMs, signature, webhookSecret); err != nil {
		slog.Error("签名验证失败（不应发生）", "error", err)
		os.Exit(1)
	}
	slog.Info("✅ 签名验证通过")

	// 解析事件
	evt, err := webhook.ParseEvent(samplePayload)
	if err != nil {
		slog.Error("解析事件失败", "error", err)
		os.Exit(1)
	}
	slog.Info("事件已解析", "id", evt.ID, "event", evt.Event)

	// 解析具体的 PaymentIntent 数据
	type PaymentIntentData struct {
		ID           string `json:"id"`
		Status       string `json:"status"`
		Amount       float64 `json:"amount"`
		Currency     string `json:"currency"`
		MerchantOrderID string `json:"merchant_order_id,omitempty"`
	}
	data, err := webhook.UnmarshalData[PaymentIntentData](evt)
	if err != nil {
		slog.Error("解析事件数据失败", "error", err)
		os.Exit(1)
	}
	slog.Info("✅ 业务数据已提取",
		"payment_intent_id", data.ID,
		"amount", data.Amount,
		"currency", data.Currency,
		"status", data.Status,
	)

	slog.Info("")
	slog.Info("=== 完整流程演示完毕 ===")
	slog.Info("说明：")
	slog.Info("- 阶段一用真实 Sandbox API 完成支付全流程")
	slog.Info("- 阶段二演示本地 webhook 签名验证与事件解析")
	slog.Info("- 生产环境中应注册 webhook URL 到 Airwallex 后台")
	slog.Info("- 签名验证是必须的安全步骤，验证失败说明请求可能被篡改")
	slog.Info("- 事件类型通过 evt.Event 字段区分（如 payment_intent.succeeded）")
}

// demoWebhookPayload 构造一个符合 Airwallex 规范的 webhook 回调 JSON。
func demoWebhookPayload(piID string, amount float64, currency string) []byte {
	payload := map[string]interface{}{
		"id":         "evt_demo_" + time.Now().Format("20060102150405"),
		"event":      "payment_intent.succeeded",
		"created_at": time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"data": map[string]interface{}{
			"id":       piID,
			"status":   "SUCCEEDED",
			"amount":   amount,
			"currency": currency,
		},
	}
	b, _ := json.Marshal(payload)
	return b
}

// computeHMACSignature 用 Airwallex 官方算法计算 HMAC-SHA256 签名。
// 公式: HMAC-SHA256(secret, timestamp + body)
func computeHMACSignature(timestamp string, body []byte, secret string) string {
	valueToDigest := timestamp + string(body)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(valueToDigest))
	return hex.EncodeToString(mac.Sum(nil))
}
