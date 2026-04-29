package fx

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// QuoteUsage represents the quote usage type.
// QuoteUsage 报价用途类型。
type QuoteUsage = string

const (
	// QuoteUsageSingleUse is a single-use quote.
	// QuoteUsageSingleUse 一次性使用
	QuoteUsageSingleUse QuoteUsage = "SINGLE_USE"
	// QuoteUsageMultiUse is a multi-use quote.
	// QuoteUsageMultiUse 多次使用
	QuoteUsageMultiUse QuoteUsage = "MULTI_USE"
)

// Quote represents foreign exchange quote information.
// Quote 表示外汇报价信息。
type Quote struct {
	// ID is the unique quote identifier.
	// ID 报价唯一标识符。
	ID string `json:"quote_id"`
	// BuyCurrency is the buy currency (3-letter ISO-4217).
	// BuyCurrency 买入货币（3位ISO-4217代码）。
	BuyCurrency sdk.Currency `json:"buy_currency"`
	// SellCurrency is the sell currency (3-letter ISO-4217).
	// SellCurrency 卖出货币（3位ISO-4217代码）。
	SellCurrency sdk.Currency `json:"sell_currency"`
	// Rate is the client rate (returned quote).
	// Rate 客户汇率（返回的报价）。
	Rate float64 `json:"client_rate"`
	// ExpiryTime is the expiry time of the quote (ISO8601, inclusive).
	// ExpiryTime 报价有效期截止时间（ISO8601格式，包含）。
	ExpiryTime string `json:"valid_to_at"`
	// Validity is the quote validity type.
	// Validity 报价有效期类型。
	Validity QuoteValidity `json:"validity"`
	// CurrencyPair is the currency pair for the quote.
	// CurrencyPair 报价对应的货币对。
	CurrencyPair string `json:"currency_pair"`
	// DealtCurrency is the currency for which the client provided an exact amount in the request.
	// DealtCurrency 客户在请求中提供精确金额的货币。
	DealtCurrency string `json:"dealt_currency"`
	// ConversionDate is the conversion settlement date.
	// ConversionDate 兑换结算日期。
	ConversionDate string `json:"conversion_date"`
	// BuyAmount is the buy amount.
	// BuyAmount 买入金额。
	BuyAmount float64 `json:"buy_amount"`
	// SellAmount is the sell amount.
	// SellAmount 卖出金额。
	SellAmount float64 `json:"sell_amount"`
	// AwxRate is the Airwallex reference rate.
	// AwxRate Airwallex参考汇率。
	AwxRate float64 `json:"awx_rate"`
	// MidRate is the mid-market rate (midpoint between buy and sell rates).
	// MidRate 中间汇率（买入价与卖出价的中点）。
	MidRate float64 `json:"mid_rate"`
	// Usage is the quote usage, SINGLE_USE or MULTI_USE.
	// Usage 报价用途，SINGLE_USE或MULTI_USE。
	Usage QuoteUsage `json:"usage"`
	// ValidFromAt is the start time of the quote validity (ISO8601, inclusive).
	// ValidFromAt 报价有效期开始时间（ISO8601格式，包含）。
	ValidFromAt string `json:"valid_from_at"`
	// ApplicationFees is the list of application fees.
	// ApplicationFees 应用费用列表。
	ApplicationFees []ApplicationFee `json:"application_fees,omitempty"`
	// RateDetails contains rate details for different user levels.
	// RateDetails 不同用户层级的汇率明细。
	RateDetails []RateDetail `json:"rate_details"`
}

// QuoteValidity represents the quote validity type.
// QuoteValidity 报价有效期类型。
type QuoteValidity = string

const (
	// QuoteValidityMin1 is 1 minute validity.
	// QuoteValidityMin1 1分钟
	QuoteValidityMin1 QuoteValidity = "MIN_1"
	// QuoteValidityMin15 is 15 minutes validity.
	// QuoteValidityMin15 15分钟
	QuoteValidityMin15 QuoteValidity = "MIN_15"
	// QuoteValidityMin30 is 30 minutes validity.
	// QuoteValidityMin30 30分钟
	QuoteValidityMin30 QuoteValidity = "MIN_30"
	// QuoteValidityHr1 is 1 hour validity.
	// QuoteValidityHr1 1小时
	QuoteValidityHr1 QuoteValidity = "HR_1"
	// QuoteValidityHr4 is 4 hours validity.
	// QuoteValidityHr4 4小时
	QuoteValidityHr4 QuoteValidity = "HR_4"
	// QuoteValidityHr8 is 8 hours validity.
	// QuoteValidityHr8 8小时
	QuoteValidityHr8 QuoteValidity = "HR_8"
	// QuoteValidityHr24 is 24 hours validity.
	// QuoteValidityHr24 24小时
	QuoteValidityHr24 QuoteValidity = "HR_24"
)

// CreateQuoteRequest represents the request to create a foreign exchange quote.
// CreateQuoteRequest 创建外汇报价请求。
type CreateQuoteRequest struct {
	// BuyCurrency is the buy currency (3-letter ISO-4217). Required.
	// BuyCurrency 买入货币（3位ISO-4217代码）。必填。
	BuyCurrency sdk.Currency `json:"buy_currency"`
	// SellCurrency is the sell currency (3-letter ISO-4217). Required.
	// SellCurrency 卖出货币（3位ISO-4217代码）。必填。
	SellCurrency sdk.Currency `json:"sell_currency"`
	// BuyAmount is the buy amount (mutually exclusive with sell_amount).
	// BuyAmount 买入金额（与sell_amount互斥）。
	BuyAmount float64 `json:"buy_amount,omitempty"`
	// SellAmount is the sell amount (mutually exclusive with buy_amount).
	// SellAmount 卖出金额（与buy_amount互斥）。
	SellAmount float64 `json:"sell_amount,omitempty"`
	// Validity is the quote validity type. Required.
	// Validity 报价有效期类型。必填。
	Validity QuoteValidity `json:"validity"`
	// ConversionDate is the conversion settlement date.
	// ConversionDate 兑换结算日期。
	ConversionDate string `json:"conversion_date,omitempty"`
	// ApplicationFeeOptions is the list of application fee options.
	// ApplicationFeeOptions 应用费用选项列表。
	ApplicationFeeOptions []ApplicationFeeOption `json:"application_fee_options,omitempty"`
}

// CreateQuote creates a foreign exchange quote.
// 官方文档: https://www.airwallex.com/docs/api/fx/quotes/create.md
// CreateQuote 创建外汇报价。
func (s *Service) CreateQuote(ctx context.Context, req *CreateQuoteRequest, opts ...sdk.RequestOption) (*Quote, error) {
	var resp Quote
	err := s.doer.Do(ctx, "POST", "/api/v1/fx/quotes/create", req, &resp, opts...)
	return &resp, err
}

// GetQuote retrieves a foreign exchange quote by ID.
// 官方文档: https://www.airwallex.com/docs/api/fx/quotes/retrieve.md
// GetQuote 根据ID获取外汇报价。
func (s *Service) GetQuote(ctx context.Context, id string, opts ...sdk.RequestOption) (*Quote, error) {
	var resp Quote
	err := s.doer.Do(ctx, "GET", "/api/v1/fx/quotes/"+id, nil, &resp, opts...)
	return &resp, err
}
