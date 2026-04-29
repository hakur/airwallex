package pa

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PaymentAttempt represents a payment attempt.
// PaymentAttempt 表示支付尝试信息。
type PaymentAttempt struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// PaymentIntentID 支付意图唯一标识符。必填。
	PaymentIntentID string `json:"payment_intent_id"`
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// Status 支付尝试状态。必填。
	Status string `json:"status"`
	// PaymentMethod 支付方式信息。可选。
	PaymentMethod map[string]any `json:"payment_method,omitempty"`
	// FailureCode 失败代码。可选。
	FailureCode string `json:"failure_code,omitempty"`
	// FailureDetails 失败详情。可选。
	FailureDetails map[string]any `json:"failure_details,omitempty"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
	// MerchantOrderID 商户订单号。可选。
	MerchantOrderID string `json:"merchant_order_id,omitempty"`
	// ProviderTransactionID 支付提供商交易号。可选。
	ProviderTransactionID string `json:"provider_transaction_id,omitempty"`
	// PaymentMethodTransactionID 支付方式交易号。可选。
	PaymentMethodTransactionID string `json:"payment_method_transaction_id,omitempty"`
	// ProviderOriginalResponseCode 支付提供商原始响应码。可选。
	ProviderOriginalResponseCode string `json:"provider_original_response_code,omitempty"`
	// AuthorizationCode 授权码。可选。
	AuthorizationCode string `json:"authorization_code,omitempty"`
	// CapturedAmount 已捕获金额。可选。
	CapturedAmount float64 `json:"captured_amount,omitempty"`
	// RefundedAmount 已退款金额。可选。
	RefundedAmount float64 `json:"refunded_amount,omitempty"`
	// SettleVia 结算方式。可选。
	SettleVia string `json:"settle_via,omitempty"`
	// AuthenticationData 认证数据。可选。
	AuthenticationData map[string]any `json:"authentication_data,omitempty"`
	// PaymentMethodOptions 支付方式选项。可选。
	PaymentMethodOptions map[string]any `json:"payment_method_options,omitempty"`
}

// GetPaymentAttempt retrieves a payment attempt by ID.
// GetPaymentAttempt 根据 ID 获取支付尝试。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_attempts/retrieve.md
func (s *Service) GetPaymentAttempt(ctx context.Context, id string, opts ...sdk.RequestOption) (*PaymentAttempt, error) {
	var resp PaymentAttempt
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_attempts/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListPaymentAttempts lists payment attempts.
// ListPaymentAttempts 列出支付尝试。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_attempts/list.md
func (s *Service) ListPaymentAttempts(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[PaymentAttempt], error) {
	var resp sdk.ListResult[PaymentAttempt]
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_attempts", nil, &resp, opts...)
	return &resp, err
}
