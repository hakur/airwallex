package billing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// UsageEvent represents a usage event.
// UsageEvent 使用事件。
type UsageEvent struct {
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// EventName 事件名称。必填。
	EventName string `json:"event_name"`
	// HappenedAt 事件发生时间。必填。
	HappenedAt string `json:"happened_at"`
	// IngestedAt 摄入时间。必填。
	IngestedAt string `json:"ingested_at"`
	// MerchantEventID 商户事件唯一标识符。必填。
	MerchantEventID string `json:"merchant_event_id"`
	// Properties 属性。必填。
	Properties map[string]any `json:"properties"`
	// Voided 是否已作废。必填。
	Voided bool `json:"voided"`
	// VoidedAt 作废时间。可选。
	VoidedAt string `json:"voided_at,omitempty"`
}

// IngestUsageEventRequest 单个摄入请求。
type IngestUsageEventRequest struct {
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// EventName 事件名称。必填。
	EventName string `json:"event_name"`
	// HappenedAt 事件发生时间。可选。
	HappenedAt string `json:"happened_at,omitempty"`
	// MerchantEventID 商户事件唯一标识符。必填。
	MerchantEventID string `json:"merchant_event_id"`
	// Properties 属性。必填。
	Properties map[string]any `json:"properties"`
}

// BatchUsageEvent 批量事件。
type BatchUsageEvent struct {
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// EventName 事件名称。必填。
	EventName string `json:"event_name"`
	// HappenedAt 事件发生时间。可选。
	HappenedAt string `json:"happened_at,omitempty"`
	// MerchantEventID 商户事件唯一标识符。必填。
	MerchantEventID string `json:"merchant_event_id"`
	// Properties 属性。必填。
	Properties map[string]any `json:"properties"`
}

// BatchIngestUsageEventsRequest 批量摄入请求。
type BatchIngestUsageEventsRequest struct {
	// Events 事件列表。必填。
	Events []BatchUsageEvent `json:"events"`
}

// VoidUsageEventRequest 作废请求。
type VoidUsageEventRequest struct {
	// MerchantEventID 商户事件唯一标识符。必填。
	MerchantEventID string `json:"merchant_event_id"`
}

// IngestUsageEvent ingests a single usage event.
// IngestUsageEvent 单个摄入 Usage Event。
// 官方文档: https://www.airwallex.com/docs/api/billing/usage_events/ingest.md
func (s *Service) IngestUsageEvent(ctx context.Context, req *IngestUsageEventRequest, opts ...sdk.RequestOption) (*UsageEvent, error) {
	var resp UsageEvent
	err := s.doer.Do(ctx, "POST", "/api/v1/usage_events/ingest", req, &resp, opts...)
	return &resp, err
}

// BatchIngestUsageEvents batch ingests usage events.
// BatchIngestUsageEvents 批量摄入 Usage Events。
// 官方文档: https://www.airwallex.com/docs/api/billing/usage_events/batch_ingest.md
func (s *Service) BatchIngestUsageEvents(ctx context.Context, req *BatchIngestUsageEventsRequest, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/usage_events/batch_ingest", req, nil, opts...)
}

// VoidUsageEvent voids a usage event.
// VoidUsageEvent 作废 Usage Event。
// 官方文档: https://www.airwallex.com/docs/api/billing/usage_events/void.md
func (s *Service) VoidUsageEvent(ctx context.Context, req *VoidUsageEventRequest, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/usage_events/void", req, nil, opts...)
}
