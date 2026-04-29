package billing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PaymentSource represents a payment source.
// PaymentSource 支付来源。
type PaymentSource struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// ExternalID 外部唯一标识符。必填。
	ExternalID string `json:"external_id"`
	// LinkedPaymentAccountID 关联支付账户唯一标识符。必填。
	LinkedPaymentAccountID string `json:"linked_payment_account_id"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// UpdatedAt 更新时间。必填。
	UpdatedAt string `json:"updated_at"`
}

// CreatePaymentSourceRequest 创建请求。
type CreatePaymentSourceRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// ExternalID 外部唯一标识符。必填。
	ExternalID string `json:"external_id"`
	// LinkedPaymentAccountID 关联支付账户唯一标识符。可选。
	LinkedPaymentAccountID string `json:"linked_payment_account_id,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
}

// ListPaymentSourcesRequest 列出请求。
type ListPaymentSourcesRequest struct {
	// BillingCustomerID 账单客户唯一标识符。可选。
	BillingCustomerID string `json:"billing_customer_id,omitempty"`
	// LinkedPaymentAccountID 关联支付账户唯一标识符。可选。
	LinkedPaymentAccountID string `json:"linked_payment_account_id,omitempty"`
	// ExternalID 外部唯一标识符。可选。
	ExternalID string `json:"external_id,omitempty"`
	// FromCreatedAt 创建时间起始。可选。
	FromCreatedAt string `json:"from_created_at,omitempty"`
	// ToCreatedAt 创建时间截止。可选。
	ToCreatedAt string `json:"to_created_at,omitempty"`
	// Page 分页游标。可选。
	Page string `json:"page,omitempty"`
	// PageSize 每页数量。可选。
	PageSize int32 `json:"page_size,omitempty"`
}

// CreatePaymentSource creates a new payment source.
// CreatePaymentSource 创建支付来源。
// 官方文档: https://www.airwallex.com/docs/api/billing/payment_sources/create.md
func (s *Service) CreatePaymentSource(ctx context.Context, req *CreatePaymentSourceRequest, opts ...sdk.RequestOption) (*PaymentSource, error) {
	var resp PaymentSource
	err := s.doer.Do(ctx, "POST", "/api/v1/payment_sources/create", req, &resp, opts...)
	return &resp, err
}

// GetPaymentSource retrieves a payment source by ID.
// GetPaymentSource 根据 ID 获取支付来源。
// 官方文档: https://www.airwallex.com/docs/api/billing/payment_sources/retrieve.md
func (s *Service) GetPaymentSource(ctx context.Context, id string, opts ...sdk.RequestOption) (*PaymentSource, error) {
	var resp PaymentSource
	err := s.doer.Do(ctx, "GET", "/api/v1/payment_sources/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListPaymentSources lists payment sources with optional filters.
// ListPaymentSources 列出支付来源。
// 官方文档: https://www.airwallex.com/docs/api/billing/payment_sources/list.md
func (s *Service) ListPaymentSources(ctx context.Context, req *ListPaymentSourcesRequest, opts ...sdk.RequestOption) (*ListResult[PaymentSource], error) {
	var resp ListResult[PaymentSource]
	err := s.doer.Do(ctx, "GET", "/api/v1/payment_sources", req, &resp, opts...)
	return &resp, err
}
