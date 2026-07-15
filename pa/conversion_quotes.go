package pa

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// ValidityPeriod represents the validity period of a conversion quote.
// ValidityPeriod 汇率报价的有效期。
type ValidityPeriod = string

const (
	// ValidityPeriodHr1 is a 1-hour validity period.
	// ValidityPeriodHr1 1小时有效期。
	ValidityPeriodHr1 ValidityPeriod = "HR_1"
	// ValidityPeriodHr24 is a 24-hour validity period.
	// ValidityPeriodHr24 24小时有效期。
	ValidityPeriodHr24 ValidityPeriod = "HR_24"
)

// ConversionQuoteStatus represents the status of a conversion quote.
// ConversionQuoteStatus 汇率报价的状态。
type ConversionQuoteStatus = string

const (
	// ConversionQuoteStatusCreated indicates the quote is created and valid.
	// ConversionQuoteStatusCreated 报价已创建且有效。
	ConversionQuoteStatusCreated ConversionQuoteStatus = "CREATED"
	// ConversionQuoteStatusExpired indicates the quote has expired.
	// ConversionQuoteStatusExpired 报价已过期。
	ConversionQuoteStatusExpired ConversionQuoteStatus = "EXPIRED"
)

// ConversionQuote represents a conversion quote.
// ConversionQuote 表示汇率报价信息。
type ConversionQuote struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// MerchantCurrency is the merchant currency code. Required.
	// MerchantCurrency 商户货币代码。必填。
	MerchantCurrency sdk.Currency `json:"merchant_currency"`
	// ShopperCurrency is the shopper currency code. Required.
	// ShopperCurrency 购物者货币代码。必填。
	ShopperCurrency sdk.Currency `json:"shopper_currency"`
	// ConversionRate 汇率。必填。
	ConversionRate float64 `json:"conversion_rate"`
	// ExpiresAt is the expiration time. Optional.
	// ExpiresAt 过期时间。可选。
	ExpiresAt string `json:"expires_at,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// ValidityPeriod is the validity period (HR_1 or HR_24). Optional, defaults to HR_1.
	// ValidityPeriod 有效期（HR_1 或 HR_24）。可选，默认 HR_1。
	ValidityPeriod ValidityPeriod `json:"validity_period,omitempty"`
	// Status indicates the quote status (CREATED or EXPIRED). Required.
	// Status 报价状态（CREATED 或 EXPIRED）。必填。
	Status ConversionQuoteStatus `json:"status"`
}

// CreateConversionQuoteRequest is the request to create a conversion quote.
// CreateConversionQuoteRequest 创建汇率报价请求。
type CreateConversionQuoteRequest struct {
	// RequestID is the client-generated idempotency key (max 64 chars). Required.
	// RequestID 客户端生成的幂等性键（最多64字符）。必填。
	RequestID string `json:"request_id"`
	// MerchantCurrency 商户货币代码。必填。
	MerchantCurrency sdk.Currency `json:"merchant_currency"`
	// ShopperCurrency is the shopper currency code. Required.
	// ShopperCurrency 购物者货币代码。必填。
	ShopperCurrency sdk.Currency `json:"shopper_currency"`
	// ValidityPeriod is the validity period (HR_1 or HR_24). Optional, defaults to HR_1.
	// ValidityPeriod 有效期（HR_1 或 HR_24）。可选，默认 HR_1。
	ValidityPeriod ValidityPeriod `json:"validity_period,omitempty"`
}

// CreateConversionQuote creates a conversion quote.
// CreateConversionQuote 创建汇率报价。
// 官方文档: https://www.airwallex.com/docs/api/payments/conversion_quotes/create.md
func (s *Service) CreateConversionQuote(ctx context.Context, req *CreateConversionQuoteRequest, opts ...sdk.RequestOption) (*ConversionQuote, error) {
	var resp ConversionQuote
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/conversion_quotes/create", req, &resp, opts...)
	return &resp, err
}

// GetConversionQuote retrieves a conversion quote by ID.
// GetConversionQuote 根据 ID 获取汇率报价。
// 官方文档: https://www.airwallex.com/docs/api/payments/conversion_quotes/retrieve.md
func (s *Service) GetConversionQuote(ctx context.Context, id string, opts ...sdk.RequestOption) (*ConversionQuote, error) {
	var resp ConversionQuote
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/conversion_quotes/"+id, nil, &resp, opts...)
	return &resp, err
}
