package issuing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// LifecycleBillingAmount represents the currency-level aggregated amounts for a lifecycle.
// LifecycleBillingAmount 生命周期的货币级别汇总金额。
type LifecycleBillingAmount struct {
	Currency        string  `json:"currency,omitempty"`
	TotalAuthorized float64 `json:"total_authorized,omitempty"`
	TotalCredited   float64 `json:"total_credited,omitempty"`
	TotalDebited    float64 `json:"total_debited,omitempty"`
	TotalDeclined   float64 `json:"total_declined,omitempty"`
	TotalExpired    float64 `json:"total_expired,omitempty"`
	TotalPending    float64 `json:"total_pending,omitempty"`
	TotalReversed   float64 `json:"total_reversed,omitempty"`
}

// Lifecycle represents the lifecycle aggregation container for a card transaction.
// Lifecycle 表示卡交易的生命周期聚合容器。
type Lifecycle struct {
	ID             string                   `json:"id,omitempty"`
	CardID         string                   `json:"card_id,omitempty"`
	CreatedAt      string                   `json:"created_at,omitempty"`
	UpdatedAt      string                   `json:"updated_at,omitempty"`
	BillingAmounts []LifecycleBillingAmount `json:"billing_amounts,omitempty"`
}

// ListLifecyclesRequest represents the lifecycle list query parameters.
// ListLifecyclesRequest 生命周期列表查询参数。
type ListLifecyclesRequest struct {
	CardID        string `json:"card_id,omitempty"`
	FromCreatedAt string `json:"from_created_at,omitempty"`
	ToCreatedAt   string `json:"to_created_at,omitempty"`
	Page          string `json:"page,omitempty"`
	PageSize      int32  `json:"page_size,omitempty"`
}

// ListLifecyclesResponse represents the lifecycle list response (cursor pagination).
// ListLifecyclesResponse 生命周期列表响应（cursor 分页）。
type ListLifecyclesResponse struct {
	Items      []Lifecycle `json:"items"`
	PageAfter  string      `json:"page_after,omitempty"`
	PageBefore string      `json:"page_before,omitempty"`
}

// ListLifecycles retrieves the list of card transaction lifecycles.
// 官方文档: https://www.airwallex.com/docs/api/issuing/card_transaction_lifecycles/lifecycles.md
// ListLifecycles 获取卡交易生命周期列表。
func (s *Service) ListLifecycles(ctx context.Context, req *ListLifecyclesRequest, opts ...sdk.RequestOption) (*ListLifecyclesResponse, error) {
	var resp ListLifecyclesResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/lifecycles", req, &resp, opts...)
	return &resp, err
}

// GetLifecycle retrieves a single card transaction lifecycle.
// 官方文档: https://www.airwallex.com/docs/api/issuing/card_transaction_lifecycles/retrieve_lifecycles.md
// GetLifecycle 获取单个卡交易生命周期。
func (s *Service) GetLifecycle(ctx context.Context, id string, opts ...sdk.RequestOption) (*Lifecycle, error) {
	var resp Lifecycle
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/lifecycles/"+id, nil, &resp, opts...)
	return &resp, err
}

// --- Deprecated ---
// ListCardTransactionLifecycles / GetCardTransactionLifecycle 已重命名为 ListLifecycles / GetLifecycle。
// 路径从 /card_transaction_lifecycles 修正为 /lifecycles。
