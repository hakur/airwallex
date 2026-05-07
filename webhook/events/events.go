// Package events provides typed webhook event structures for all Airwallex webhook domains.
// events 包为所有 Airwallex webhook 领域提供类型化的事件结构体。
//
// Each domain has its own file (e.g., online_payments.go, issuing.go).
// Typed events embed webhook.Event and shadow the Data field with a domain-specific type,
// enabling compile-time type safety when parsing webhook payloads.
//
// Usage:
//
//	evt, err := events.ParseEventAs[events.PaymentIntentSucceededEvent](payload)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(evt.Data.Object.ID) // typed as pa.PaymentIntent
package events

import (
	"encoding/json"
	"fmt"
)

// Event represents the generic structure of an Airwallex webhook event.
// Event 表示 Airwallex webhook 事件通用结构。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events.md
//
// 字段说明:
//   - ID: 事件唯一标识符
//   - Name: 事件类型名称（如 payment_intent.succeeded）
//   - AccountID: 事件所属账户 ID（账户级别事件）
//   - OrgID: 事件所属组织 ID（组织级别事件）
//   - SourceID: 源 ID（部分领域如 issuing 使用）
//   - Data: 事件载荷原始 JSON（类型取决于事件类型）
//   - CreatedAt: 事件创建时间
//   - Version: API 版本号
type Event struct {
	// ID is the unique event identifier.
	// ID 事件唯一标识符。
	ID string `json:"id"`
	// Name is the event type name (e.g., payment_intent.succeeded).
	// Name 事件类型名称（如 payment_intent.succeeded）。
	Name string `json:"name"`
	// AccountID is the unique identifier of the account this event belongs to.
	// Only applicable to account-level events; can be empty.
	// AccountID 事件所属账户 ID。仅适用于账户级别事件；可能为空。
	AccountID string `json:"account_id,omitempty"`
	// OrgID is the unique identifier of the organization this event belongs to.
	// Only applicable to organization-level events; can be empty.
	// OrgID 事件所属组织 ID。仅适用于组织级别事件；可能为空。
	OrgID string `json:"org_id,omitempty"`
	// SourceID is the source identifier. Used by some domains like issuing.
	// SourceID 源标识符。部分领域（如 issuing）使用。
	SourceID string `json:"source_id,omitempty"`
	// Data is the event payload raw JSON (type depends on event type).
	// Data 事件载荷原始 JSON（类型取决于事件类型）。
	Data json.RawMessage `json:"data"`
	// CreatedAt is the event creation time.
	// CreatedAt 事件创建时间。
	CreatedAt string `json:"created_at"`
	// Version is the API version for the webhook subscription.
	// Version webhook 订阅的 API 版本。
	Version string `json:"version,omitempty"`
}

// ParseEventAs parses a webhook payload into a typed event structure.
// ParseEventAs 将 webhook payload 解析为类型化事件结构体。
func ParseEventAs[T any](payload []byte) (*T, error) {
	var evt T
	if err := json.Unmarshal(payload, &evt); err != nil {
		return nil, fmt.Errorf("webhook: unmarshal typed event: %w", err)
	}
	return &evt, nil
}
