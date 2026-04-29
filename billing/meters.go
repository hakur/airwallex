package billing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// Meter represents a meter.
// Meter 计量器。
type Meter struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Name 名称。必填。
	Name string `json:"name"`
	// EventName 事件名称。必填。
	EventName string `json:"event_name"`
	// AggregationMethod 聚合方法。必填。
	AggregationMethod MeterAggregationMethod `json:"aggregation_method"`
	// AggregationProperty 聚合属性。可选。
	AggregationProperty string `json:"aggregation_property,omitempty"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// Archived 是否已归档。必填。
	Archived bool `json:"archived"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// UpdatedAt 更新时间。必填。
	UpdatedAt string `json:"updated_at"`
}

// CreateMeterRequest 创建请求。
type CreateMeterRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Name 名称。必填。
	Name string `json:"name"`
	// EventName 事件名称。必填。
	EventName string `json:"event_name"`
	// AggregationMethod 聚合方法。必填。
	AggregationMethod MeterAggregationMethod `json:"aggregation_method"`
	// AggregationProperty 聚合属性。可选。
	AggregationProperty string `json:"aggregation_property,omitempty"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
}

// UpdateMeterRequest 更新请求。
type UpdateMeterRequest struct {
	// Name 名称。可选。
	Name string `json:"name,omitempty"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
}

// MeterSummary 计量器汇总。
type MeterSummary struct {
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// MeterID 计量器唯一标识符。必填。
	MeterID string `json:"meter_id"`
	// FromHappenedAt 事件起始时间。必填。
	FromHappenedAt string `json:"from_happened_at"`
	// ToHappenedAt 事件截止时间。必填。
	ToHappenedAt string `json:"to_happened_at"`
	// Value 汇总值。必填。
	Value float64 `json:"value"`
	// EventCount 事件数量。必填。
	EventCount int64 `json:"event_count"`
}

// ListMetersRequest 列出计量器请求。
type ListMetersRequest struct {
	// Archived 是否已归档。可选。
	Archived bool `json:"archived,omitempty"`
	// EventName 事件名称。可选。
	EventName string `json:"event_name,omitempty"`
	// AggregationMethod 聚合方法。可选。
	AggregationMethod string `json:"aggregation_method,omitempty"`
	// AggregationProperty 聚合属性。可选。
	AggregationProperty string `json:"aggregation_property,omitempty"`
	// FromCreatedAt 创建时间起始。可选。
	FromCreatedAt string `json:"from_created_at,omitempty"`
	// ToCreatedAt 创建时间截止。可选。
	ToCreatedAt string `json:"to_created_at,omitempty"`
	// Page 分页游标。可选。
	Page string `json:"page,omitempty"`
	// PageSize 每页数量。可选。
	PageSize int32 `json:"page_size,omitempty"`
}

// MeterSummariesRequest 汇总请求。
type MeterSummariesRequest struct {
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// FromHappenedAt 事件起始时间。必填。
	FromHappenedAt string `json:"from_happened_at"`
	// ToHappenedAt 事件截止时间。必填。
	ToHappenedAt string `json:"to_happened_at"`
	// GroupedBy 分组方式。可选。
	GroupedBy string `json:"grouped_by,omitempty"`
	// Page 分页游标。可选。
	Page string `json:"page,omitempty"`
	// PageSize 每页数量。可选。
	PageSize int32 `json:"page_size,omitempty"`
}

// CreateMeter creates a new meter.
// CreateMeter 创建计量器。
// 官方文档: https://www.airwallex.com/docs/api/billing/meters/create.md
func (s *Service) CreateMeter(ctx context.Context, req *CreateMeterRequest, opts ...sdk.RequestOption) (*Meter, error) {
	var resp Meter
	err := s.doer.Do(ctx, "POST", "/api/v1/meters/create", req, &resp, opts...)
	return &resp, err
}

// GetMeter retrieves a meter by ID.
// GetMeter 根据 ID 获取计量器。
// 官方文档: https://www.airwallex.com/docs/api/billing/meters/retrieve.md
func (s *Service) GetMeter(ctx context.Context, id string, opts ...sdk.RequestOption) (*Meter, error) {
	var resp Meter
	err := s.doer.Do(ctx, "GET", "/api/v1/meters/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateMeter updates an existing meter.
// UpdateMeter 更新计量器。
// 官方文档: https://www.airwallex.com/docs/api/billing/meters/update.md
func (s *Service) UpdateMeter(ctx context.Context, id string, req *UpdateMeterRequest, opts ...sdk.RequestOption) (*Meter, error) {
	var resp Meter
	err := s.doer.Do(ctx, "POST", "/api/v1/meters/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ArchiveMeter archives a meter.
// ArchiveMeter 归档计量器。
// 官方文档: https://www.airwallex.com/docs/api/billing/meters/archive.md
func (s *Service) ArchiveMeter(ctx context.Context, id string, opts ...sdk.RequestOption) (*Meter, error) {
	var resp Meter
	err := s.doer.Do(ctx, "POST", "/api/v1/meters/"+id+"/archive", nil, &resp, opts...)
	return &resp, err
}

// RestoreMeter restores an archived meter.
// RestoreMeter 恢复计量器。
// 官方文档: https://www.airwallex.com/docs/api/billing/meters/restore.md
func (s *Service) RestoreMeter(ctx context.Context, id string, opts ...sdk.RequestOption) (*Meter, error) {
	var resp Meter
	err := s.doer.Do(ctx, "POST", "/api/v1/meters/"+id+"/restore", nil, &resp, opts...)
	return &resp, err
}

// ListMeters lists meters with optional filters.
// ListMeters 列出计量器。
// 官方文档: https://www.airwallex.com/docs/api/billing/meters/list.md
func (s *Service) ListMeters(ctx context.Context, req *ListMetersRequest, opts ...sdk.RequestOption) (*ListResult[Meter], error) {
	var resp ListResult[Meter]
	err := s.doer.Do(ctx, "GET", "/api/v1/meters", req, &resp, opts...)
	return &resp, err
}

// GetMeterSummaries retrieves meter summary data.
// GetMeterSummaries 获取计量器汇总。
// 官方文档: https://www.airwallex.com/docs/api/billing/meters/retrieve.md
func (s *Service) GetMeterSummaries(ctx context.Context, id string, req *MeterSummariesRequest, opts ...sdk.RequestOption) (*ListResult[MeterSummary], error) {
	var resp ListResult[MeterSummary]
	err := s.doer.Do(ctx, "GET", "/api/v1/meters/"+id+"/summaries", req, &resp, opts...)
	return &resp, err
}
