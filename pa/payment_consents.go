package pa

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PaymentConsent represents a payment consent.
// PaymentConsent 表示支付授权信息。
type PaymentConsent struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// RequestID is the unique request identifier. Optional.
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// CustomerID is the customer unique identifier. Required.
	// CustomerID 客户唯一标识符。必填。
	CustomerID string `json:"customer_id"`
	// PaymentMethodID is the payment method unique identifier. Optional.
	// PaymentMethodID 支付方式唯一标识符。可选。
	PaymentMethodID string `json:"payment_method_id,omitempty"`
	// Status is the consent status. Required.
	// Status 授权状态。必填。
	Status string `json:"status"`
	// NextTriggerDate is the next trigger date. Optional.
	// NextTriggerDate 下次触发日期。可选。
	NextTriggerDate string `json:"next_trigger_date,omitempty"`
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

// CreatePaymentConsentRequest is the request to create a payment consent.
// CreatePaymentConsentRequest 创建支付授权请求。
type CreatePaymentConsentRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// CustomerID is the customer unique identifier. Required.
	// CustomerID 客户唯一标识符。必填。
	CustomerID string `json:"customer_id"`
	// PaymentMethodID is the payment method unique identifier. Optional.
	// PaymentMethodID 支付方式唯一标识符。可选。
	PaymentMethodID string `json:"payment_method_id,omitempty"`
	// NextTriggeredBy is the next trigger initiator. Required.
	// NextTriggeredBy 下次触发方。必填。
	NextTriggeredBy string `json:"next_triggered_by"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// CancelPaymentConsentRequest is the request to cancel a payment consent.
// CancelPaymentConsentRequest 取消支付授权请求。
type CancelPaymentConsentRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
}

// CreatePaymentConsent creates a payment consent.
// CreatePaymentConsent 创建支付授权。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_consents/create.md
func (s *Service) CreatePaymentConsent(ctx context.Context, req *CreatePaymentConsentRequest, opts ...sdk.RequestOption) (*PaymentConsent, error) {
	var resp PaymentConsent
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_consents/create", req, &resp, opts...)
	return &resp, err
}

// GetPaymentConsent retrieves a payment consent by ID.
// GetPaymentConsent 根据 ID 获取支付授权。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_consents/retrieve.md
func (s *Service) GetPaymentConsent(ctx context.Context, id string, opts ...sdk.RequestOption) (*PaymentConsent, error) {
	var resp PaymentConsent
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_consents/"+id, nil, &resp, opts...)
	return &resp, err
}

// CancelPaymentConsent cancels a payment consent.
// CancelPaymentConsent 取消支付授权。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_consents/cancel.md
func (s *Service) CancelPaymentConsent(ctx context.Context, id string, req *CancelPaymentConsentRequest, opts ...sdk.RequestOption) (*PaymentConsent, error) {
	var resp PaymentConsent
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_consents/"+id+"/cancel", req, &resp, opts...)
	return &resp, err
}

// ListPaymentConsents lists payment consents.
// ListPaymentConsents 列出支付授权。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_consents/list.md
func (s *Service) ListPaymentConsents(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[PaymentConsent], error) {
	var resp sdk.ListResult[PaymentConsent]
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_consents", nil, &resp, opts...)
	return &resp, err
}
