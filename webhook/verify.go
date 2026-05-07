package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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
