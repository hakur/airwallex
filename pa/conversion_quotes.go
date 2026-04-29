package pa

import (
	"context"

	"github.com/hakur/airwallex/sdk"
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
	// ValidityPeriod is the validity period. Optional.
	// ValidityPeriod 有效期。可选。
	ValidityPeriod string `json:"validity_period,omitempty"`
	// Status 报价状态。必填。
	Status string `json:"status"`
}

// CreateConversionQuoteRequest is the request to create a conversion quote.
// CreateConversionQuoteRequest 创建汇率报价请求。
type CreateConversionQuoteRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// MerchantCurrency 商户货币代码。必填。
	MerchantCurrency sdk.Currency `json:"merchant_currency"`
	// ShopperCurrency is the shopper currency code. Required.
	// ShopperCurrency 购物者货币代码。必填。
	ShopperCurrency sdk.Currency `json:"shopper_currency"`
	// ValidityPeriod is the validity period. Optional.
	// ValidityPeriod 有效期。可选。
	ValidityPeriod string `json:"validity_period,omitempty"`
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
