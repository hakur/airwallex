package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// VerifySignature verifies the HMAC-SHA256 signature of an Airwallex webhook.
// VerifySignature 验证 Airwallex webhook 的 HMAC-SHA256 签名。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events.md
//
// 参数:
//   - payload: 请求体原始字节（不要使用解析后重新序列化的 JSON）
//   - timestampHeader: x-timestamp header 值（Unix 毫秒时间戳字符串）
//   - signatureHeader: x-signature header 值（HMAC-SHA256 hex digest）
//   - secret: webhook 通知 URL 的密钥
//
// 验证步骤:
//  1. 将 x-timestamp 字符串与请求体原始字节拼接: valueToDigest = "{timestamp}{body}"
//  2. 用 secret 作为 HMAC key，SHA-256 计算 valueToDigest 的 HMAC hex digest
//  3. 比对计算结果与 x-signature 是否一致
//  4. 检查时间戳是否在 5 分钟允许范围内
func VerifySignature(payload []byte, timestampHeader, signatureHeader, secret string) error {
	return VerifySignatureWithTolerance(payload, timestampHeader, signatureHeader, secret, 300*1000)
}

// VerifySignatureWithTolerance verifies the webhook signature with a custom timestamp tolerance (milliseconds).
// A tolerance of 0 skips the timestamp check.
// VerifySignatureWithTolerance 验证签名并指定时间戳容差（毫秒）。
// tolerance 为 0 时不检查时间戳。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events.md
func VerifySignatureWithTolerance(payload []byte, timestampHeader, signatureHeader, secret string, tolerance int64) error {
	ts, err := strconv.ParseInt(timestampHeader, 10, 64)
	if err != nil {
		return fmt.Errorf("webhook: invalid timestamp header: %w", err)
	}

	if tolerance > 0 {
		now := time.Now().UnixMilli()
		if ts < now-tolerance || ts > now+tolerance {
			return fmt.Errorf("webhook: timestamp outside tolerance window")
		}
	}

	// 官方格式: 直接拼接 timestamp 字符串 + body，无分隔符
	valueToDigest := timestampHeader + string(payload)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(valueToDigest))
	expected := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(signatureHeader), []byte(expected)) {
		return fmt.Errorf("webhook: signature mismatch")
	}
	return nil
}

// Event represents the generic structure of an Airwallex webhook event.
// Event 表示 Airwallex webhook 事件通用结构。
type Event struct {
	// ID is the unique event identifier.
	// ID 事件唯一标识符。
	ID string `json:"id"`
	// Event is the event type name.
	// Event 事件类型名称。
	Event string `json:"event"`
	// Data is the event payload (JSON object, type depends on event type).
	// Data 事件载荷（JSON 对象，类型取决于 event）。
	Data json.RawMessage `json:"data"`
	// CreatedAt is the event creation time.
	// CreatedAt 事件创建时间。
	CreatedAt string `json:"created_at"`
}

// ParseEvent parses a webhook payload into an Event struct.
// ParseEvent 将 webhook payload 解析为 Event 结构体。
func ParseEvent(payload []byte) (*Event, error) {
	var evt Event
	if err := json.Unmarshal(payload, &evt); err != nil {
		return nil, fmt.Errorf("webhook: unmarshal event: %w", err)
	}
	return &evt, nil
}

// UnmarshalData unmarshals Event.Data into a value of the specified type.
// UnmarshalData 将 Event.Data 解析为指定类型的值。
func UnmarshalData[T any](evt *Event) (*T, error) {
	var v T
	if err := json.Unmarshal(evt.Data, &v); err != nil {
		return nil, fmt.Errorf("webhook: unmarshal event data: %w", err)
	}
	return &v, nil
}
