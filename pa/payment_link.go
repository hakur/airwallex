package pa

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PaymentLink represents a payment link.
// PaymentLink 表示支付链接信息。
type PaymentLink struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// URL is the payment link URL. Required.
	// URL 支付链接地址。必填。
	URL string `json:"url"`
	// Amount is the amount. Required.
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// Status is the status. Required.
	// Status 状态。必填。
	Status string `json:"status"`
	// Title is the title. Optional.
	// Title 标题。可选。
	Title string `json:"title,omitempty"`
	// Reusable 是否可重复使用。可选。
	Reusable bool `json:"reusable,omitempty"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
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

// CreatePaymentLinkRequest is the request to create a payment link.
// CreatePaymentLinkRequest 创建支付链接请求。
type CreatePaymentLinkRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Title is the title. Required.
	// Title 标题。必填。
	Title string `json:"title"`
	// Amount is the amount. Required.
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// Reusable indicates whether the link is reusable. Optional.
	// Reusable 是否可重复使用。可选。
	Reusable bool `json:"reusable,omitempty"`
	// Description is the description. Optional.
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// CreatePaymentLink creates a payment link.
// CreatePaymentLink 创建支付链接。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_links/create.md
func (s *Service) CreatePaymentLink(ctx context.Context, req *CreatePaymentLinkRequest, opts ...sdk.RequestOption) (*PaymentLink, error) {
	var resp PaymentLink
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_links/create", req, &resp, opts...)
	return &resp, err
}

// GetPaymentLink retrieves a payment link by ID.
// GetPaymentLink 根据 ID 获取支付链接。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_links/retrieve.md
func (s *Service) GetPaymentLink(ctx context.Context, id string, opts ...sdk.RequestOption) (*PaymentLink, error) {
	var resp PaymentLink
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_links/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListPaymentLinks lists payment links.
// ListPaymentLinks 列出支付链接。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_links/list.md
func (s *Service) ListPaymentLinks(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[PaymentLink], error) {
	var resp sdk.ListResult[PaymentLink]
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_links", nil, &resp, opts...)
	return &resp, err
}
