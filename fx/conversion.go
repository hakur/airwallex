package fx

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// ConversionStatus represents the foreign exchange conversion status.
// ConversionStatus 外汇兑换状态。
type ConversionStatus = string

const (
	// ConversionStatusScheduled indicates the conversion is scheduled.
	// ConversionStatusScheduled 已安排
	ConversionStatusScheduled ConversionStatus = "SCHEDULED"
	// ConversionStatusCancelled indicates the conversion is cancelled.
	// ConversionStatusCancelled 已取消
	ConversionStatusCancelled ConversionStatus = "CANCELLED"
	// ConversionStatusOverdue indicates the conversion is overdue.
	// ConversionStatusOverdue 已逾期
	ConversionStatusOverdue ConversionStatus = "OVERDUE"
	// ConversionStatusSettled indicates the conversion is settled.
	// ConversionStatusSettled 已结算
	ConversionStatusSettled ConversionStatus = "SETTLED"
)

// RateLevel represents the rate level.
// RateLevel 汇率层级。
type RateLevel = string

const (
	// RateLevelClient is the client rate level.
	// RateLevelClient 客户层级
	RateLevelClient RateLevel = "CLIENT"
	// RateLevelAwx is the Airwallex rate level.
	// RateLevelAwx Airwallex层级
	RateLevelAwx RateLevel = "AWX"
	// RateLevelPlatformClient is the platform client rate level.
	// RateLevelPlatformClient 平台客户层级
	RateLevelPlatformClient RateLevel = "PLATFORM_CLIENT"
)

// FeeSourceType represents the fee source type.
// FeeSourceType 费用来源类型。
type FeeSourceType = string

const (
	// FeeSourceTypeConversion is a conversion fee.
	// FeeSourceTypeConversion 兑换费用
	FeeSourceTypeConversion FeeSourceType = "CONVERSION"
	// FeeSourceTypeTransfer is a transfer fee.
	// FeeSourceTypeTransfer 转账费用
	FeeSourceTypeTransfer FeeSourceType = "TRANSFER"
)

// FeeCalculationType represents the fee calculation type.
// FeeCalculationType 费用计算类型。
type FeeCalculationType = string

const (
	// FeeCalculationTypeFixed is a fixed fee.
	// FeeCalculationTypeFixed 固定费用
	FeeCalculationTypeFixed FeeCalculationType = "FIXED"
	// FeeCalculationTypePercentage is a percentage fee.
	// FeeCalculationTypePercentage 百分比费用
	FeeCalculationTypePercentage FeeCalculationType = "PERCENTAGE"
)

// DebitType represents the debit type.
// DebitType 扣款类型。
type DebitType = string

const (
	// DebitTypeDirectDebit is a direct debit.
	// DebitTypeDirectDebit 直接扣款
	DebitTypeDirectDebit DebitType = "DIRECT_DEBIT"
	// DebitTypeFasterDirectDebit is a faster direct debit.
	// DebitTypeFasterDirectDebit 快速直接扣款
	DebitTypeFasterDirectDebit DebitType = "FASTER_DIRECT_DEBIT"
)

// ApplicationFee represents an application fee (in response).
// ApplicationFee 应用费用（响应中的application_fees）。
type ApplicationFee struct {
	// Amount is the application fee amount.
	// Amount 应用费用金额。
	Amount string `json:"amount,omitempty"`
	// Currency is the application fee currency (ISO-4217).
	// Currency 应用费用货币（ISO-4217代码）。
	Currency sdk.Currency `json:"currency,omitempty"`
	// SourceType is the fee source transaction type.
	// SourceType 费用来源交易类型。
	SourceType FeeSourceType `json:"source_type,omitempty"`
}

// ApplicationFeeOption represents an application fee option (in request).
// ApplicationFeeOption 应用费用选项（请求中的application_fee_options）。
type ApplicationFeeOption struct {
	// Amount is the fee amount (positive integer in smallest currency unit, e.g. cents for USD).
	// Amount 费用金额（以货币最小单位表示的正整数，如USD的分）。
	Amount string `json:"amount,omitempty"`
	// Currency is the fee currency (ISO-4217).
	// Currency 费用货币（ISO-4217代码）。
	Currency sdk.Currency `json:"currency,omitempty"`
	// Metadata is the metadata for the application fee (key-value pairs).
	// Metadata 应用费用的元数据（键值对格式）。
	Metadata map[string]any `json:"metadata,omitempty"`
	// Percentage is the percentage fee value (required when type is PERCENTAGE).
	// Percentage 百分比费用值（当type为PERCENTAGE时必填）。
	Percentage string `json:"percentage,omitempty"`
	// SourceType is the fee source transaction type.
	// SourceType 费用来源交易类型。
	SourceType FeeSourceType `json:"source_type,omitempty"`
	// Type is the fee calculation type (FIXED or PERCENTAGE).
	// Type 费用计算类型（FIXED或PERCENTAGE）。
	Type FeeCalculationType `json:"type,omitempty"`
}

// RateDetail represents rate details.
// RateDetail 汇率详情。
type RateDetail struct {
	// BuyAmount is the total buy amount at this level.
	// BuyAmount 该层级买入总金额。
	BuyAmount float64 `json:"buy_amount,omitempty"`
	// Level is the rate level (CLIENT/AWX/PLATFORM_CLIENT).
	// Level 汇率层级（CLIENT/AWX/PLATFORM_CLIENT）。
	Level RateLevel `json:"level,omitempty"`
	// Rate is the rate at this level.
	// Rate 该层级的汇率。
	Rate float64 `json:"rate,omitempty"`
	// SellAmount is the total sell amount at this level.
	// SellAmount 该层级卖出总金额。
	SellAmount float64 `json:"sell_amount,omitempty"`
}

// Funding represents funding source information (response).
// Funding 资金来源信息（响应）。
type Funding struct {
	// DebitType is the debit type (DIRECT_DEBIT or FASTER_DIRECT_DEBIT).
	// DebitType 扣款类型（DIRECT_DEBIT或FASTER_DIRECT_DEBIT）。
	DebitType DebitType `json:"debit_type,omitempty"`
	// FailureReason is the funding failure reason (non-empty only in error state).
	// FailureReason 资金失败原因（仅在错误状态时非空）。
	FailureReason string `json:"failure_reason,omitempty"`
	// FundingSourceID is the funding source ID (null when using wallet).
	// FundingSourceID 资金来源ID（使用钱包时为null）。
	FundingSourceID string `json:"funding_source_id,omitempty"`
	// Status is the funding status.
	// Status 资金状态。
	Status string `json:"status,omitempty"`
}

// FundingSource represents a funding source (request and response).
// FundingSource 资金来源（请求和响应）。
type FundingSource struct {
	// DebitType is the debit type (DIRECT_DEBIT or FASTER_DIRECT_DEBIT).
	// DebitType 扣款类型（DIRECT_DEBIT或FASTER_DIRECT_DEBIT）。
	DebitType DebitType `json:"debit_type,omitempty"`
	// ID is the funding source ID (supports linked accounts only; defaults to wallet if not specified).
	// ID 资金来源ID（仅支持关联账户，未指定则默认使用钱包）。
	ID string `json:"id"`
}

// Conversion represents foreign exchange conversion information.
// Conversion 表示外汇兑换信息。
type Conversion struct {
	// ConversionID is the unique conversion identifier.
	// ConversionID 兑换唯一标识符。
	ConversionID string `json:"conversion_id"`
	// RequestID is the client-provided idempotency key.
	// RequestID 客户端提供的幂等性键。
	RequestID string `json:"request_id"`
	// BuyCurrency is the buy currency (3-letter ISO-4217).
	// BuyCurrency 买入货币（3位ISO-4217代码）。
	BuyCurrency sdk.Currency `json:"buy_currency"`
	// SellCurrency is the sell currency (3-letter ISO-4217).
	// SellCurrency 卖出货币（3位ISO-4217代码）。
	SellCurrency sdk.Currency `json:"sell_currency"`
	// BuyAmount is the buy amount.
	// BuyAmount 买入金额。
	BuyAmount float64 `json:"buy_amount"`
	// SellAmount is the sell amount.
	// SellAmount 卖出金额。
	SellAmount float64 `json:"sell_amount"`
	// Status is the conversion status (SCHEDULED/CANCELLED/OVERDUE/SETTLED).
	// Status 兑换状态（SCHEDULED/CANCELLED/OVERDUE/SETTLED）。
	Status ConversionStatus `json:"status"`
	// CurrencyPair is the currency pair.
	// CurrencyPair 货币对。
	CurrencyPair string `json:"currency_pair"`
	// DealtCurrency is the fixed-side currency (buy_currency if buy_amount was specified, otherwise sell_currency).
	// DealtCurrency 固定侧的货币（指定buy_amount则为buy_currency，否则为sell_currency）。
	DealtCurrency string `json:"dealt_currency"`
	// ConversionDate is the conversion settlement date.
	// ConversionDate 兑换结算日期。
	ConversionDate string `json:"conversion_date"`
	// ClientRate is the client execution rate (the rate the customer is charged).
	// ClientRate 客户执行汇率（客户被收取的汇率）。
	ClientRate float64 `json:"client_rate"`
	// AwxRate is the Airwallex reference rate.
	// AwxRate Airwallex参考汇率。
	AwxRate float64 `json:"awx_rate"`
	// MidRate is the mid-market rate (midpoint between buy and sell rates).
	// MidRate 中间汇率（买入价与卖出价的中点）。
	MidRate float64 `json:"mid_rate"`
	// QuoteID is the quote ID used to execute the conversion.
	// QuoteID 执行兑换所使用的报价ID。
	QuoteID string `json:"quote_id,omitempty"`
	// SettlementCutoffAt is the latest time sell_amount needs to be available in the Airwallex account balance.
	// SettlementCutoffAt sell_amount需在Airwallex账户余额中最晚可用时间。
	SettlementCutoffAt string `json:"settlement_cutoff_at"`
	// ShortReferenceID is the shortened transaction ID (for use in Web GUI and customer support).
	// ShortReferenceID 缩短的交易ID（便于在Web GUI和客服中使用）。
	ShortReferenceID string `json:"short_reference_id"`
	// CreatedAt is the conversion creation time.
	// CreatedAt 兑换创建时间。
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the last update time of the conversion (including status changes).
	// UpdatedAt 兑换最后更新时间（包含状态变更）。
	UpdatedAt string `json:"updated_at"`
	// ApplicationFees is the list of application fees.
	// ApplicationFees 应用费用列表。
	ApplicationFees []ApplicationFee `json:"application_fees,omitempty"`
	// ApplicationFeeOptions is the list of application fee options.
	// ApplicationFeeOptions 应用费用选项列表。
	ApplicationFeeOptions []ApplicationFeeOption `json:"application_fee_options,omitempty"`
	// RateDetails contains rate details for different user levels.
	// RateDetails 不同用户层级的汇率明细。
	RateDetails []RateDetail `json:"rate_details,omitempty"`
	// Funding contains funding information.
	// Funding 资金来源信息。
	Funding *Funding `json:"funding,omitempty"`
	// FundingSource is the funding source.
	// FundingSource 资金来源。
	FundingSource *FundingSource `json:"funding_source,omitempty"`
}

// CreateConversionRequest represents the request to create a foreign exchange conversion.
// CreateConversionRequest 创建外汇兑换请求。
type CreateConversionRequest struct {
	// RequestID is the client-generated idempotency key (max 64 chars; requests with the same request_id will be rejected). Required.
	// RequestID 客户端生成的幂等性键（最多64字符，同一request_id的请求会被拒绝）。必填。
	RequestID string `json:"request_id"`
	// BuyCurrency is the buy currency (3-letter ISO-4217). Required.
	// BuyCurrency 买入货币（3位ISO-4217代码）。必填。
	BuyCurrency sdk.Currency `json:"buy_currency"`
	// SellCurrency is the sell currency (3-letter ISO-4217). Required.
	// SellCurrency 卖出货币（3位ISO-4217代码）。必填。
	SellCurrency sdk.Currency `json:"sell_currency"`
	// BuyAmount is the buy amount (mutually exclusive with sell_amount, string type).
	// BuyAmount 买入金额（与sell_amount互斥，字符串类型）。
	BuyAmount string `json:"buy_amount,omitempty"`
	// SellAmount is the sell amount (mutually exclusive with buy_amount, string type).
	// SellAmount 卖出金额（与buy_amount互斥，字符串类型）。
	SellAmount string `json:"sell_amount,omitempty"`
	// ConversionDate is the conversion settlement date (defaults to current date if not specified).
	// ConversionDate 兑换结算日期（未指定则默认为当前日期）。
	ConversionDate string `json:"conversion_date,omitempty"`
	// QuoteID is the valid quote ID to execute the conversion at the quoted client_rate.
	// QuoteID 希望按报价client_rate执行兑换的有效报价ID。
	QuoteID string `json:"quote_id,omitempty"`
	// ApplicationFeeOptions is the list of application fee options.
	// ApplicationFeeOptions 应用费用选项列表。
	ApplicationFeeOptions []ApplicationFeeOption `json:"application_fee_options,omitempty"`
	// FundingSource is the funding source information (defaults to wallet if not specified).
	// FundingSource 资金来源信息（未指定则默认使用钱包）。
	FundingSource *FundingSource `json:"funding_source,omitempty"`
	// Metadata is client-defined metadata (key-value pairs).
	// Metadata 客户端自定义元数据（键值对）。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// CreateConversion creates a foreign exchange conversion.
// 官方文档: https://www.airwallex.com/docs/api/fx/conversions/create.md
// CreateConversion 创建外汇兑换。
func (s *Service) CreateConversion(ctx context.Context, req *CreateConversionRequest, opts ...sdk.RequestOption) (*Conversion, error) {
	var resp Conversion
	err := s.doer.Do(ctx, "POST", "/api/v1/fx/conversions/create", req, &resp, opts...)
	return &resp, err
}

// GetConversion retrieves a foreign exchange conversion by ID.
// 官方文档: https://www.airwallex.com/docs/api/fx/conversions/retrieve.md
// GetConversion 根据ID获取外汇兑换。
func (s *Service) GetConversion(ctx context.Context, id string, opts ...sdk.RequestOption) (*Conversion, error) {
	var resp Conversion
	err := s.doer.Do(ctx, "GET", "/api/v1/fx/conversions/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListConversions lists foreign exchange conversions.
// 官方文档: https://www.airwallex.com/docs/api/fx/conversions/list.md
// ListConversions 列出外汇兑换。
func (s *Service) ListConversions(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[Conversion], error) {
	var resp sdk.ListResult[Conversion]
	err := s.doer.Do(ctx, "GET", "/api/v1/fx/conversions", nil, &resp, opts...)
	return &resp, err
}
