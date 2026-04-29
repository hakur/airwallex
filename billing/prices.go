package billing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// Price represents a price response.
// Price 价格响应。
type Price struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Active indicates whether the price is active. Required.
	// Active 是否活跃。必填。
	Active bool `json:"active"`
	// BillingType is the billing type. Optional.
	// BillingType 账单类型。可选。
	BillingType BillingType `json:"billing_type,omitempty"`
	// CreatedAt is the creation time. Required.
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency string `json:"currency"`
	// Description is the description. Optional.
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// FlatAmount is the flat amount. Optional.
	// FlatAmount 固定金额。可选。
	FlatAmount float64 `json:"flat_amount,omitempty"`
	// Metadata is the metadata key-value pairs. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// MeterID is the meter unique identifier. Optional.
	// MeterID 计量器唯一标识符。可选。
	MeterID string `json:"meter_id,omitempty"`
	// Metered indicates whether the price is metered. Required.
	// Metered 是否按量计费。必填。
	Metered bool `json:"metered"`
	// PricingModel is the pricing model. Required.
	// PricingModel 定价模型。必填。
	PricingModel PricingModel `json:"pricing_model"`
	// ProductID is the product unique identifier. Required.
	// ProductID 产品唯一标识符。必填。
	ProductID string `json:"product_id"`
	// Recurring is the recurring configuration. Optional.
	// Recurring 周期性配置。可选。
	Recurring *Recurring `json:"recurring,omitempty"`
	// Tiers is the list of pricing tiers. Optional.
	// Tiers 价格层级列表。可选。
	Tiers []PriceTier `json:"tiers,omitempty"`
	// Type is the price type. Required.
	// Type 价格类型。必填。
	Type PriceType `json:"type"`
	// UnitAmount is the unit amount. Optional.
	// UnitAmount 单价。可选。
	UnitAmount float64 `json:"unit_amount,omitempty"`
	// UpdatedAt is the update time. Required.
	UpdatedAt string `json:"updated_at"`
}

// CreatePriceRequest represents a request to create a price.
// CreatePriceRequest 创建价格请求。
type CreatePriceRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Active 是否活跃。可选。
	Active bool `json:"active,omitempty"`
	// BillingType 账单类型。可选。
	BillingType BillingType `json:"billing_type,omitempty"`
	// Currency 货币代码。必填。
	Currency string `json:"currency"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// FlatAmount 固定金额。可选。
	FlatAmount float64 `json:"flat_amount,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// MeterID 计量器唯一标识符。可选。
	MeterID string `json:"meter_id,omitempty"`
	// Metered 是否按量计费。可选。
	Metered bool `json:"metered,omitempty"`
	// PricingModel 定价模型。可选。
	PricingModel PricingModel `json:"pricing_model,omitempty"`
	// ProductID 产品唯一标识符。必填。
	ProductID string `json:"product_id"`
	// Recurring 周期性配置。可选。
	Recurring *Recurring `json:"recurring,omitempty"`
	// Tiers 价格层级列表。可选。
	Tiers []PriceTier `json:"tiers,omitempty"`
	// UnitAmount 单价。可选。
	UnitAmount float64 `json:"unit_amount,omitempty"`
}

// UpdatePriceRequest represents a request to update a price.
// UpdatePriceRequest 更新价格请求。
// 根据官方文档，仅支持更新以下字段：active、description、metadata。
type UpdatePriceRequest struct {
	// Active 是否活跃。可选。
	Active bool `json:"active,omitempty"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
}

// ListPricesRequest represents a request to list prices.
// ListPricesRequest 获取价格列表请求。
type ListPricesRequest struct {
	// ProductID 产品唯一标识符。可选。
	ProductID string `json:"product_id,omitempty"`
	// Active 是否活跃。可选。
	Active bool `json:"active,omitempty"`
	// Currency 货币代码。可选。
	Currency string `json:"currency,omitempty"`
	// RecurringPeriod 周期数。可选。
	RecurringPeriod int32 `json:"recurring_period,omitempty"`
	// RecurringPeriodUnit 周期单位。可选。
	RecurringPeriodUnit PeriodUnit `json:"recurring_period_unit,omitempty"`
	// Page 分页游标。可选。
	Page string `json:"page,omitempty"`
	// PageSize 每页数量。可选。
	PageSize int32 `json:"page_size,omitempty"`
}

// CreatePrice creates a new price.
// CreatePrice 创建价格。
// 官方文档: https://www.airwallex.com/docs/api/billing/prices/create.md
func (s *Service) CreatePrice(ctx context.Context, req *CreatePriceRequest, opts ...sdk.RequestOption) (*Price, error) {
	var resp Price
	err := s.doer.Do(ctx, "POST", "/api/v1/prices/create", req, &resp, opts...)
	return &resp, err
}

// GetPrice retrieves a price by ID.
// GetPrice 根据 ID 获取价格。
// 官方文档: https://www.airwallex.com/docs/api/billing/prices/retrieve.md
func (s *Service) GetPrice(ctx context.Context, id string, opts ...sdk.RequestOption) (*Price, error) {
	var resp Price
	err := s.doer.Do(ctx, "GET", "/api/v1/prices/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdatePrice updates an existing price.
// UpdatePrice 更新价格。
// 官方文档: https://www.airwallex.com/docs/api/billing/prices/update.md
func (s *Service) UpdatePrice(ctx context.Context, id string, req *UpdatePriceRequest, opts ...sdk.RequestOption) (*Price, error) {
	var resp Price
	err := s.doer.Do(ctx, "POST", "/api/v1/prices/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ListPrices lists prices with optional filters.
// ListPrices 列出价格。
// 官方文档: https://www.airwallex.com/docs/api/billing/prices/list.md
func (s *Service) ListPrices(ctx context.Context, req *ListPricesRequest, opts ...sdk.RequestOption) (*ListResult[Price], error) {
	var resp ListResult[Price]
	err := s.doer.Do(ctx, "GET", "/api/v1/prices", req, &resp, opts...)
	return &resp, err
}
