package pa

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PaymentDispute represents a payment dispute.
// PaymentDispute 表示支付争议信息。
type PaymentDispute struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// PaymentIntentID 支付意图唯一标识符。可选。
	PaymentIntentID string `json:"payment_intent_id,omitempty"`
	// PaymentAttemptID 支付尝试唯一标识符。可选。
	PaymentAttemptID string `json:"payment_attempt_id,omitempty"`
	// Amount 争议金额。必填。
	Amount float64 `json:"amount"`
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// Reason 争议原因。可选。
	Reason map[string]any `json:"reason,omitempty"`
	// Status 争议状态。必填。
	Status string `json:"status"`
	// Stage 争议阶段。可选。
	Stage string `json:"stage,omitempty"`
	// Mode 争议模式。可选。
	Mode string `json:"mode,omitempty"`
	// TransactionType 交易类型。可选。
	TransactionType string `json:"transaction_type,omitempty"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
	// DueAt 截止日期。可选。
	DueAt string `json:"due_at,omitempty"`
	// CustomerName 客户姓名。可选。
	CustomerName string `json:"customer_name,omitempty"`
	// MerchantOrderID 商户订单号。可选。
	MerchantOrderID string `json:"merchant_order_id,omitempty"`
	// AcquirerReferenceNumber 收单机构参考号。可选。
	AcquirerReferenceNumber string `json:"acquirer_reference_number,omitempty"`
	// CardBrand 卡品牌。可选。
	CardBrand string `json:"card_brand,omitempty"`
	// PaymentMethodType 支付方式类型。可选。
	PaymentMethodType string `json:"payment_method_type,omitempty"`
	// IssuerComment 发卡行评论。可选。
	IssuerComment string `json:"issuer_comment,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// Evidence 争议证据。可选。
	Evidence map[string]any `json:"evidence,omitempty"`
	// AcceptDetails 接受详情。可选。
	AcceptDetails []map[string]any `json:"accept_details,omitempty"`
	// ChallengeDetails 挑战详情。可选。
	ChallengeDetails []map[string]any `json:"challenge_details,omitempty"`
	// Refunds 退款列表。可选。
	Refunds []map[string]any `json:"refunds,omitempty"`
}

// GetPaymentDispute retrieves a payment dispute by ID.
// GetPaymentDispute 根据 ID 获取支付争议。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_disputes/retrieve.md
func (s *Service) GetPaymentDispute(ctx context.Context, id string, opts ...sdk.RequestOption) (*PaymentDispute, error) {
	var resp PaymentDispute
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_disputes/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListPaymentDisputes lists payment disputes.
// ListPaymentDisputes 列出支付争议。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_disputes/list.md
func (s *Service) ListPaymentDisputes(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[PaymentDispute], error) {
	var resp sdk.ListResult[PaymentDispute]
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_disputes", nil, &resp, opts...)
	return &resp, err
}
