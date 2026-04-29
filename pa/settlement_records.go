package pa

import (
	"context"
	"net/url"
	"strconv"

	"github.com/hakur/airwallex/sdk"
)

// SettlementRecord represents settlement record details.
// SettlementRecord 表示结算记录详细信息。
type SettlementRecord struct {
	// ID 唯一标识符。可选。
	ID string `json:"id,omitempty"`
	// AccountID 账户唯一标识符。可选。
	AccountID string `json:"account_id,omitempty"`
	// AcquirerReferenceNumber 收单机构参考号。可选。
	AcquirerReferenceNumber string `json:"acquirer_reference_number,omitempty"`
	// CardCategory 卡类别。可选。
	CardCategory string `json:"card_category,omitempty"`
	// CardFunding 卡资金来源。可选。
	CardFunding string `json:"card_funding,omitempty"`
	// CardIssuingOrShopperCountry 发卡国家或购物者国家。可选。
	CardIssuingOrShopperCountry string `json:"card_issuing_or_shopper_country,omitempty"`
	// CardTransactionRegion 卡交易地区。可选。
	CardTransactionRegion string `json:"card_transaction_region,omitempty"`
	// ConnectedAccountID 关联账户唯一标识符。可选。
	ConnectedAccountID string `json:"connected_account_id,omitempty"`
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// CustomerEmail 客户邮箱。可选。
	CustomerEmail string `json:"customer_email,omitempty"`
	// CustomerID 客户唯一标识符。可选。
	CustomerID string `json:"customer_id,omitempty"`
	// CustomerName 客户姓名。可选。
	CustomerName string `json:"customer_name,omitempty"`
	// CustomerPhone 客户电话。可选。
	CustomerPhone string `json:"customer_phone,omitempty"`
	// Descriptor 交易描述。可选。
	Descriptor string `json:"descriptor,omitempty"`
	// DisputeReason 争议原因。可选。
	DisputeReason string `json:"dispute_reason,omitempty"`
	// DisputeReasonCode 争议原因代码。可选。
	DisputeReasonCode string `json:"dispute_reason_code,omitempty"`
	// DisputeStatus 争议状态。可选。
	DisputeStatus string `json:"dispute_status,omitempty"`
	// ExchangeRate 汇率。可选。
	ExchangeRate float64 `json:"exchange_rate,omitempty"`
	// FeeDetailsList 费用明细列表。可选。
	FeeDetailsList []SettlementRecordFeeDetail `json:"fee_details_list,omitempty"`
	// Fees 费用金额。可选。
	Fees float64 `json:"fees,omitempty"`
	// GrossAmount 总金额。可选。
	GrossAmount float64 `json:"gross_amount,omitempty"`
	// IssuingOrOriginatingBank 发卡行或发起银行。可选。
	IssuingOrOriginatingBank string `json:"issuing_or_originating_bank,omitempty"`
	// MerchantCustomerID 商户客户唯一标识符。可选。
	MerchantCustomerID string `json:"merchant_customer_id,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// NetAmount 净额。可选。
	NetAmount float64 `json:"net_amount,omitempty"`
	// OrderID 订单唯一标识符。可选。
	OrderID string `json:"order_id,omitempty"`
	// PaymentAttemptIDs 支付尝试唯一标识符列表。可选。
	PaymentAttemptIDs []string `json:"payment_attempt_ids,omitempty"`
	// PaymentCreatedTime 支付创建时间。可选。
	PaymentCreatedTime string `json:"payment_created_time,omitempty"`
	// PaymentIntentID 支付意图唯一标识符。可选。
	PaymentIntentID string `json:"payment_intent_id,omitempty"`
	// PaymentLinkReference 支付链接引用。可选。
	PaymentLinkReference string `json:"payment_link_reference,omitempty"`
	// PaymentMethod 支付方式。可选。
	PaymentMethod string `json:"payment_method,omitempty"`
	// PaymentMethodType 支付方式类型。可选。
	PaymentMethodType string `json:"payment_method_type,omitempty"`
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// SettledAt 结算时间。可选。
	SettledAt string `json:"settled_at,omitempty"`
	// SettlementBatchID 结算批次唯一标识符。可选。
	SettlementBatchID string `json:"settlement_batch_id,omitempty"`
	// SettlementCurrency 结算货币代码。可选。
	SettlementCurrency sdk.Currency `json:"settlement_currency,omitempty"`
	// ShippingAddress 配送地址。可选。
	ShippingAddress string `json:"shipping_address,omitempty"`
	// ShippingCity 配送城市。可选。
	ShippingCity string `json:"shipping_city,omitempty"`
	// ShippingCountry 配送国家。可选。
	ShippingCountry string `json:"shipping_country,omitempty"`
	// ShippingName 配送联系人姓名。可选。
	ShippingName string `json:"shipping_name,omitempty"`
	// ShippingPostalCode 配送邮政编码。可选。
	ShippingPostalCode string `json:"shipping_postal_code,omitempty"`
	// ShippingState 配送州/省。可选。
	ShippingState string `json:"shipping_state,omitempty"`
	// SourceEntity 源实体。可选。
	SourceEntity string `json:"source_entity,omitempty"`
	// SourceID 源唯一标识符。可选。
	SourceID string `json:"source_id,omitempty"`
	// SubscriptionID 订阅唯一标识符。可选。
	SubscriptionID string `json:"subscription_id,omitempty"`
	// TaxesOnFees 费用税费。可选。
	TaxesOnFees float64 `json:"taxes_on_fees,omitempty"`
	// TransactionAmount 交易金额。可选。
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
	// TransactionCurrency 交易货币代码。可选。
	TransactionCurrency sdk.Currency `json:"transaction_currency,omitempty"`
	// TransactionID 交易唯一标识符。可选。
	TransactionID string `json:"transaction_id,omitempty"`
	// TransactionType 交易类型。可选。
	TransactionType string `json:"transaction_type,omitempty"`
}

// SettlementRecordFeeDetail represents fee details.
// SettlementRecordFeeDetail 表示费用明细。
type SettlementRecordFeeDetail struct {
	// ExchangeRate 汇率。可选。
	ExchangeRate float64 `json:"exchange_rate,omitempty"`
	// FeeCurrency 费用货币代码。可选。
	FeeCurrency string `json:"fee_currency,omitempty"`
	// FeeSubtotalAmount 费用小计金额。可选。
	FeeSubtotalAmount float64 `json:"fee_subtotal_amount,omitempty"`
	// FeeType 费用类型。可选。
	FeeType string `json:"fee_type,omitempty"`
	// IncurredAt 费用发生时间。可选。
	IncurredAt string `json:"incurred_at,omitempty"`
	// PaymentMethodNumberOrID 支付方式编号或唯一标识符。可选。
	PaymentMethodNumberOrID string `json:"payment_method_number_or_id,omitempty"`
	// SettlementCurrency 结算货币代码。可选。
	SettlementCurrency string `json:"settlement_currency,omitempty"`
	// SettlementSubtotalAmount 结算小计金额。可选。
	SettlementSubtotalAmount float64 `json:"settlement_subtotal_amount,omitempty"`
	// SettlementTaxAmount 结算税费金额。可选。
	SettlementTaxAmount float64 `json:"settlement_tax_amount,omitempty"`
	// SettlementTotalAmount 结算总金额。可选。
	SettlementTotalAmount float64 `json:"settlement_total_amount,omitempty"`
	// SourceStatus 源状态。可选。
	SourceStatus string `json:"source_status,omitempty"`
}

// ListSettlementRecordsRequest is the request to query settlement records.
// ListSettlementRecordsRequest 查询结算记录请求参数。
type ListSettlementRecordsRequest struct {
	// Page 页码。可选（查询参数）。
	Page string `json:"-"`
	// PageSize 每页数量。可选（查询参数）。
	PageSize int32 `json:"-"`
	// PaymentIntentID 支付意图唯一标识符。可选（查询参数）。
	PaymentIntentID string `json:"-"`
	// SettlementID 结算唯一标识符。可选（查询参数）。
	SettlementID string `json:"-"`
}

// GetSettlementRecord retrieves a settlement record by ID.
// GetSettlementRecord 根据 ID 获取结算记录。
// 官方文档: https://www.airwallex.com/docs/api/payments/settlement_records/retrieve.md
func (s *Service) GetSettlementRecord(ctx context.Context, id string, opts ...sdk.RequestOption) (*SettlementRecord, error) {
	var resp SettlementRecord
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/settlement_records/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListSettlementRecords lists settlement records.
// ListSettlementRecords 列出结算记录。
// 官方文档: https://www.airwallex.com/docs/api/payments/settlement_records/list.md
func (s *Service) ListSettlementRecords(ctx context.Context, req *ListSettlementRecordsRequest, opts ...sdk.RequestOption) (*sdk.ListResult[SettlementRecord], error) {
	path := "/api/v1/pa/settlement_records"
	if req != nil {
		q := url.Values{}
		if req.Page != "" {
			q.Set("page", req.Page)
		}
		if req.PageSize > 0 {
			q.Set("page_size", strconv.Itoa(int(req.PageSize)))
		}
		if req.PaymentIntentID != "" {
			q.Set("payment_intent_id", req.PaymentIntentID)
		}
		if req.SettlementID != "" {
			q.Set("settlement_id", req.SettlementID)
		}
		if len(q) > 0 {
			path += "?" + q.Encode()
		}
	}
	var resp sdk.ListResult[SettlementRecord]
	err := s.doer.Do(ctx, "GET", path, nil, &resp, opts...)
	return &resp, err
}
