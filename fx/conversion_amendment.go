package fx

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// AmendmentType represents the foreign exchange conversion amendment type.
// AmendmentType 外汇兑换修改类型。
type AmendmentType = string

const (
	// AmendmentTypeCancel fully cancels the original conversion.
	// AmendmentTypeCancel 全额取消原始兑换。
	AmendmentTypeCancel AmendmentType = "CANCEL"
)

// ChargeType represents the adjustment charge type.
// ChargeType 调整费用类型。
type ChargeType = string

const (
	// ChargeTypeFee is a fee charge.
	// ChargeTypeFee 费用。
	ChargeTypeFee ChargeType = "FEE"
	// ChargeTypeCredit is a credit charge.
	// ChargeTypeCredit 贷记。
	ChargeTypeCredit ChargeType = "CREDIT"
)

// Charge represents a financial adjustment resulting from an amendment.
// Charge 表示修改产生的财务调整。
type Charge struct {
	// Amount is the financial adjustment amount (cancellation difference).
	// Amount 财务调整金额（取消差额）。
	Amount float64 `json:"amount,omitempty"`
	// AwxRate is the Airwallex rate used for the amendment (with liquidity partner).
	// AwxRate Airwallex执行修改时使用的汇率（与流动性合作伙伴）。
	AwxRate float64 `json:"awx_rate,omitempty"`
	// ClientRate is the rate the customer pays for the amendment (inherited from existing conversion rate config).
	// ClientRate 客户为修改支付的汇率（继承自现有兑换汇率配置）。
	ClientRate float64 `json:"client_rate,omitempty"`
	// Currency is the financial adjustment currency (always the non-fixed side currency of the original conversion).
	// Currency 财务调整货币（始终为原始兑换的非固定侧货币）。
	Currency string `json:"currency,omitempty"`
	// CurrencyPair is the currency pair associated with awx_rate and client_rate.
	// CurrencyPair 与awx_rate和client_rate关联的货币对。
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Type is the charge type (FEE or CREDIT).
	// Type 费用类型（FEE或CREDIT）。
	Type ChargeType `json:"type,omitempty"`
}

// ConversionAmendment represents foreign exchange conversion amendment information.
// ConversionAmendment 表示外汇兑换修改信息。
type ConversionAmendment struct {
	// AmendmentID is the unique amendment identifier.
	// AmendmentID 修改唯一标识符。
	AmendmentID string `json:"amendment_id,omitempty"`
	// Charges is the list of financial adjustments resulting from the amendment.
	// Charges 修改产生的财务调整列表。
	Charges []Charge `json:"charges,omitempty"`
	// ConversionID is the ID of the modified conversion.
	// ConversionID 被修改的兑换ID。
	ConversionID string `json:"conversion_id,omitempty"`
	// CreatedAt is the amendment creation time.
	// CreatedAt 修改创建时间。
	CreatedAt string `json:"created_at,omitempty"`
	// Metadata is client-provided metadata.
	// Metadata 客户端提供的元数据。
	Metadata map[string]any `json:"metadata,omitempty"`
	// RequestID is the client-provided idempotency key.
	// RequestID 客户端提供的幂等性键。
	RequestID string `json:"request_id,omitempty"`
	// ShortReferenceID is the shortened transaction ID.
	// ShortReferenceID 缩短的交易ID。
	ShortReferenceID string `json:"short_reference_id,omitempty"`
	// Type is the amendment type (only CANCEL is supported).
	// Type 修改类型（仅支持CANCEL）。
	Type AmendmentType `json:"type,omitempty"`
	// UpdatedAt is the last update time of the amendment.
	// UpdatedAt 修改最后更新时间。
	UpdatedAt string `json:"updated_at,omitempty"`
}

// ConversionAmendmentQuote represents a foreign exchange conversion amendment quote.
// ConversionAmendmentQuote 表示外汇兑换修改报价信息。
type ConversionAmendmentQuote struct {
	// Charges is the list of estimated financial adjustments resulting from the amendment.
	// Charges 修改产生的预估财务调整列表。
	Charges []Charge `json:"charges,omitempty"`
	// ConversionID is the ID of the modified conversion.
	// ConversionID 被修改的兑换ID。
	ConversionID string `json:"conversion_id,omitempty"`
	// Metadata is client-provided metadata.
	// Metadata 客户端提供的元数据。
	Metadata map[string]any `json:"metadata,omitempty"`
	// RequestID is the client-provided idempotency key.
	// RequestID 客户端提供的幂等性键。
	RequestID string `json:"request_id,omitempty"`
	// ShortReferenceID is the shortened transaction ID.
	// ShortReferenceID 缩短的交易ID。
	ShortReferenceID string `json:"short_reference_id,omitempty"`
	// Type is the amendment type (only CANCEL is supported).
	// Type 修改类型（仅支持CANCEL）。
	Type AmendmentType `json:"type,omitempty"`
}

// CreateConversionAmendmentRequest represents the request to create a foreign exchange conversion amendment.
// CreateConversionAmendmentRequest 创建外汇兑换修改请求。
type CreateConversionAmendmentRequest struct {
	// RequestID is the client-generated idempotency key. Required.
	// RequestID 客户端生成的幂等性键。必填。
	RequestID string `json:"request_id"`
	// ConversionID is the ID of the conversion to amend. Required.
	// ConversionID 待修改的兑换ID。必填。
	ConversionID string `json:"conversion_id"`
	// Type is the amendment type (only CANCEL is supported). Required.
	// Type 修改类型（仅支持CANCEL）。必填。
	Type AmendmentType `json:"type"`
	// ChargeCurrency is the currency for the amendment charge (must be one of the original conversion's currencies; defaults to the non-fixed side currency).
	// ChargeCurrency 修改产生的差额所收取的货币（仅限原始兑换的货币，默认为非固定侧货币）。
	ChargeCurrency string `json:"charge_currency,omitempty"`
	// Metadata is client-provided metadata (max 10 key-value pairs, keys and values must be less than 255 characters).
	// Metadata 客户端提供的元数据（最多10个键值对，键和值均需小于255字符）。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// ListConversionAmendmentsRequest represents the request to list conversion amendments.
// ListConversionAmendmentsRequest 列出外汇兑换修改请求参数。
type ListConversionAmendmentsRequest struct {
	// ConversionID is the conversion ID. Required.
	// ConversionID 兑换ID。必填。
	ConversionID string `json:"conversion_id"`
}

// CreateConversionAmendment creates a foreign exchange conversion amendment.
// 官方文档: https://www.airwallex.com/docs/api/fx/conversion_amendments/create.md
// CreateConversionAmendment 创建外汇兑换修改。
func (s *Service) CreateConversionAmendment(ctx context.Context, req *CreateConversionAmendmentRequest, opts ...sdk.RequestOption) (*ConversionAmendment, error) {
	var resp ConversionAmendment
	err := s.doer.Do(ctx, "POST", "/api/v1/fx/conversion_amendments/create", req, &resp, opts...)
	return &resp, err
}

// ListConversionAmendments lists all amendments for a specified conversion.
// 官方文档: https://www.airwallex.com/docs/api/fx/conversion_amendments/list.md
// ListConversionAmendments 列出指定兑换的所有修改。
func (s *Service) ListConversionAmendments(ctx context.Context, req *ListConversionAmendmentsRequest, opts ...sdk.RequestOption) (*sdk.ListResult[ConversionAmendment], error) {
	var resp sdk.ListResult[ConversionAmendment]
	err := s.doer.Do(ctx, "GET", "/api/v1/fx/conversion_amendments", req, &resp, opts...)
	return &resp, err
}

// CreateConversionAmendmentQuote creates a foreign exchange conversion amendment quote (estimated fees).
// 官方文档: https://www.airwallex.com/docs/api/fx/conversion_amendments/quote.md
// CreateConversionAmendmentQuote 创建外汇兑换修改报价（获取预估费用）。
func (s *Service) CreateConversionAmendmentQuote(ctx context.Context, req *CreateConversionAmendmentRequest, opts ...sdk.RequestOption) (*ConversionAmendmentQuote, error) {
	var resp ConversionAmendmentQuote
	err := s.doer.Do(ctx, "POST", "/api/v1/fx/conversion_amendments/quote", req, &resp, opts...)
	return &resp, err
}

// GetConversionAmendment retrieves a foreign exchange conversion amendment by ID.
// 官方文档: https://www.airwallex.com/docs/api/fx/conversion_amendments/retrieve.md
// GetConversionAmendment 根据ID获取外汇兑换修改。
func (s *Service) GetConversionAmendment(ctx context.Context, id string, opts ...sdk.RequestOption) (*ConversionAmendment, error) {
	var resp ConversionAmendment
	err := s.doer.Do(ctx, "GET", "/api/v1/fx/conversion_amendments/"+id, nil, &resp, opts...)
	return &resp, err
}
