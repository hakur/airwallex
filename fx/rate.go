package fx

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hakur/airwallex/sdk"
)

// Rate represents foreign exchange rate information.
// Rate 表示外汇汇率信息。
type Rate struct {
	// BuyCurrency is the buy currency (3-letter ISO-4217).
	// BuyCurrency 买入货币（3位ISO-4217代码）。
	BuyCurrency sdk.Currency `json:"buy_currency"`
	// SellCurrency is the sell currency (3-letter ISO-4217).
	// SellCurrency 卖出货币（3位ISO-4217代码）。
	SellCurrency sdk.Currency `json:"sell_currency"`
	// Rate is the exchange rate.
	// Rate 汇率。
	Rate float64 `json:"rate"`
	// ConversionDate is the conversion settlement date.
	// ConversionDate 兑换结算日期。
	ConversionDate string `json:"conversion_date"`
	// CreatedAt is the request creation time.
	// CreatedAt 请求创建时间。
	CreatedAt string `json:"created_at"`
	// CurrencyPair is the currency pair for the rate.
	// CurrencyPair 汇率对应的货币对。
	CurrencyPair string `json:"currency_pair"`
	// DealtCurrency is the currency for which the client provided an exact amount in the request.
	// DealtCurrency 客户在请求中提供精确金额的货币。
	DealtCurrency string `json:"dealt_currency"`
	// RateDetails contains rate details for different user levels.
	// RateDetails 不同用户层级的汇率明细。
	RateDetails []RateDetail `json:"rate_details"`
}

// GetRatesRequest represents the request to get current rates.
// GetRatesRequest 获取当前汇率请求参数。
type GetRatesRequest struct {
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
	// ConversionDate is the conversion settlement date.
	// ConversionDate 兑换结算日期。
	ConversionDate string `json:"conversion_date,omitempty"`
}

// GetRates retrieves current exchange rates.
// 官方文档: https://www.airwallex.com/docs/api/fx/rates/retrieve.md
// GetRates 获取当前汇率。
func (s *Service) GetRates(ctx context.Context, req *GetRatesRequest, opts ...sdk.RequestOption) (*Rate, error) {
	var resp Rate
	path := "/api/v1/fx/rates/current"
	if req != nil {
		query := url.Values{}
		query.Set("buy_currency", string(req.BuyCurrency))
		query.Set("sell_currency", string(req.SellCurrency))
		if req.BuyAmount > 0 {
			query.Set("buy_amount", fmt.Sprintf("%.0f", req.BuyAmount))
		}
		if req.SellAmount > 0 {
			query.Set("sell_amount", fmt.Sprintf("%.0f", req.SellAmount))
		}
		if req.ConversionDate != "" {
			query.Set("conversion_date", req.ConversionDate)
		}
		path = path + "?" + query.Encode()
	}
	err := s.doer.Do(ctx, "GET", path, nil, &resp, opts...)
	return &resp, err
}
