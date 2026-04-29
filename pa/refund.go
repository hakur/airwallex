package pa

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// RefundStatus represents the refund status.
// RefundStatus 表示退款状态。
type RefundStatus = string

const (
	// RefundStatusReceived indicates the refund has been received.
	// RefundStatusReceived 退款已创建。
	RefundStatusReceived RefundStatus = "RECEIVED"
	// RefundStatusAccepted indicates the refund has been accepted by the payment provider.
	// RefundStatusAccepted 退款已被支付提供商接受。
	RefundStatusAccepted RefundStatus = "ACCEPTED"
	// RefundStatusSettled indicates the refund has been settled.
	// RefundStatusSettled 退款已结算。
	RefundStatusSettled RefundStatus = "SETTLED"
	// RefundStatusFailed indicates the refund has failed.
	// RefundStatusFailed 退款失败。
	RefundStatusFailed RefundStatus = "FAILED"
)

// Refund represents a refund.
// Refund 表示退款信息。
type Refund struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// RequestID is the unique request identifier. Optional.
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// PaymentIntentID is the payment intent unique identifier. Required.
	// PaymentIntentID 支付意图唯一标识符。必填。
	PaymentIntentID string `json:"payment_intent_id"`
	// PaymentAttemptID is the payment attempt unique identifier. Optional.
	// PaymentAttemptID 支付尝试唯一标识符。可选。
	PaymentAttemptID string `json:"payment_attempt_id,omitempty"`
	// Amount is the refund amount. Required.
	// Amount 退款金额。必填。
	Amount float64 `json:"amount"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// Status is the refund status. Required.
	// Status 退款状态。必填。
	Status string `json:"status"`
	// Reason is the refund reason. Optional.
	// Reason 退款原因。可选。
	Reason string `json:"reason,omitempty"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt is the update time. Optional.
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
}

// CreateRefundRequest is the request to create a refund.
// CreateRefundRequest 创建退款请求。
type CreateRefundRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// PaymentIntentID is the payment intent unique identifier. Required.
	// PaymentIntentID 支付意图唯一标识符。必填。
	PaymentIntentID string `json:"payment_intent_id"`
	// Amount is the refund amount. Optional.
	// Amount 退款金额。可选。
	Amount float64 `json:"amount,omitempty"`
}

// CreateRefund creates a refund.
// CreateRefund 创建退款。
// 官方文档: https://www.airwallex.com/docs/api/payments/refunds/create.md
func (s *Service) CreateRefund(ctx context.Context, req *CreateRefundRequest, opts ...sdk.RequestOption) (*Refund, error) {
	var resp Refund
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/refunds/create", req, &resp, opts...)
	return &resp, err
}

// GetRefund retrieves a refund by ID.
// GetRefund 根据 ID 获取退款。
// 官方文档: https://www.airwallex.com/docs/api/payments/refunds/retrieve.md
func (s *Service) GetRefund(ctx context.Context, id string, opts ...sdk.RequestOption) (*Refund, error) {
	var resp Refund
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/refunds/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListRefunds lists refunds.
// ListRefunds 列出退款。
// 官方文档: https://www.airwallex.com/docs/api/payments/refunds/list.md
func (s *Service) ListRefunds(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[Refund], error) {
	var resp sdk.ListResult[Refund]
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/refunds", nil, &resp, opts...)
	return &resp, err
}
