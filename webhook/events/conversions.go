// Package events provides typed webhook event structures for the conversions domain.
// Conversions 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/conversions.md
//
// 事件映射表:
//
//	conversion.scheduled                → ConversionScheduledEvent         (Data: ConversionEventData)
//	conversion.overdue                  → ConversionOverdueEvent           (Data: ConversionEventData)
//	conversion.settled                  → ConversionSettledEvent           (Data: ConversionEventData)
//	conversion.cancelled                → ConversionCancelledEvent         (Data: ConversionEventData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

import "encoding/json"

// --- Conversion Events ---

// ConversionScheduledEvent represents the conversion.scheduled webhook event.
// ConversionScheduledEvent 表示 conversion.scheduled webhook 事件。
type ConversionScheduledEvent struct {
	Event
	Data ConversionEventData `json:"data"`
}

// ConversionOverdueEvent represents the conversion.overdue webhook event.
// ConversionOverdueEvent 表示 conversion.overdue webhook 事件。
type ConversionOverdueEvent struct {
	Event
	Data ConversionEventData `json:"data"`
}

// ConversionSettledEvent represents the conversion.settled webhook event.
// ConversionSettledEvent 表示 conversion.settled webhook 事件。
type ConversionSettledEvent struct {
	Event
	Data ConversionEventData `json:"data"`
}

// ConversionCancelledEvent represents the conversion.cancelled webhook event.
// ConversionCancelledEvent 表示 conversion.cancelled webhook 事件。
type ConversionCancelledEvent struct {
	Event
	Data ConversionEventData `json:"data"`
}

// --- Event Data Structures ---

// ConversionEventData represents the payload data for conversion webhook events.
// ConversionEventData 表示 conversion webhook 事件的载荷数据。
// 官方文档说明 data 为 Get a specific conversion 的响应体。
type ConversionEventData struct {
	RequestID             string                           `json:"requestId"`
	ConversionID          string                           `json:"conversionId"`
	ShortReferenceID      string                           `json:"shortReferenceId"`
	Status                string                           `json:"status"`
	CurrencyPair          string                           `json:"currencyPair"`
	ClientRate            float64                          `json:"clientRate"`
	AwxRate               float64                          `json:"awxRate"`
	MidRate               float64                          `json:"midRate"`
	BuyCurrency           string                           `json:"buyCurrency"`
	BuyAmount             float64                          `json:"buyAmount"`
	SellCurrency          string                           `json:"sellCurrency"`
	SellAmount            float64                          `json:"sellAmount"`
	DealtCurrency         string                           `json:"dealtCurrency"`
	ConversionDate        string                           `json:"conversionDate"`
	SettlementCutoffTime  string                           `json:"settlementCutoffTime"`
	CreatedAt             string                           `json:"createdAt"`
	UpdatedAt             string                           `json:"updatedAt"`
	ClientData            json.RawMessage                  `json:"clientData,omitempty"`
	BatchID               string                           `json:"batchId,omitempty"`
	QuoteID               string                           `json:"quoteId,omitempty"`
	RateDetails           []ConversionRateDetail           `json:"rateDetails,omitempty"`
	FundingSource         *ConversionFundingSource         `json:"fundingSource,omitempty"`
	Funding               *ConversionFunding               `json:"funding,omitempty"`
	ApplicationFeeOptions []ConversionApplicationFeeOption `json:"applicationFeeOptions,omitempty"`
	ApplicationFees       []ConversionApplicationFee       `json:"applicationFees,omitempty"`
}

// ConversionRateDetail represents a single rate detail within a conversion event.
// ConversionRateDetail 表示兑换事件中的汇率明细。
type ConversionRateDetail struct {
	Level      string  `json:"level"`
	Rate       float64 `json:"rate"`
	BuyAmount  float64 `json:"buyAmount"`
	SellAmount float64 `json:"sellAmount"`
}

// ConversionFundingSource represents the funding source information in a conversion event.
// ConversionFundingSource 表示兑换事件中的资金来源信息。
type ConversionFundingSource struct {
	ID        string `json:"id"`
	DebitType string `json:"debitType"`
}

// ConversionFunding represents the funding information in a conversion event.
// ConversionFunding 表示兑换事件中的资金信息。
type ConversionFunding struct {
	FundingSourceID string `json:"fundingSourceId"`
	DebitType       string `json:"debitType"`
	Status          string `json:"status"`
}

// ConversionApplicationFeeOption represents an application fee option in a conversion event.
// ConversionApplicationFeeOption 表示兑换事件中的应用费用选项。
type ConversionApplicationFeeOption struct {
	SourceType string  `json:"sourceType"`
	Type       string  `json:"type"`
	Percentage float64 `json:"percentage"`
}

// ConversionApplicationFee represents an application fee in a conversion event.
// ConversionApplicationFee 表示兑换事件中的应用费用。
type ConversionApplicationFee struct {
	SourceType string  `json:"sourceType"`
	Currency   string  `json:"currency"`
	Amount     float64 `json:"amount"`
}
