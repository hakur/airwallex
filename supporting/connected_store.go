package supporting

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// ConnectedStore represents an e-commerce platform's connected store.
// ConnectedStore 表示电商平台已连接店铺。
type ConnectedStore struct {
	// ID is the unique identifier of the connected store.
	// ID 唯一标识符。
	ID string `json:"id,omitempty"`
	// Name is the store name.
	// Name 店铺名称。
	Name string `json:"name,omitempty"`
	// Platform is the e-commerce platform name.
	// Platform 电商平台名称。
	Platform string `json:"platform,omitempty"`
	// PlatformSellerID is the seller's ID on the platform.
	// PlatformSellerID 电商平台卖家 ID。
	PlatformSellerID string `json:"platform_seller_id,omitempty"`
	// Status is the connection status of the store.
	// Status 连接状态。
	Status string `json:"status,omitempty"`
	// URL is the store URL.
	// URL 店铺 URL。
	URL string `json:"url,omitempty"`
	// DefaultCurrency is the default currency used by the store.
	// DefaultCurrency 默认货币。
	DefaultCurrency string `json:"default_currency,omitempty"`
}

// ListConnectedStoresRequest represents query parameters for listing connected stores.
// ListConnectedStoresRequest 已连接店铺列表查询参数。
type ListConnectedStoresRequest struct {
	// FromCreatedAt is the start of the creation date range.
	// FromCreatedAt 创建时间范围起点。
	FromCreatedAt string `json:"from_created_at,omitempty"`
	// ToCreatedAt is the end of the creation date range.
	// ToCreatedAt 创建时间范围终点。
	ToCreatedAt string `json:"to_created_at,omitempty"`
	// PageNum is the page number for pagination.
	// PageNum 页码。
	PageNum int32 `json:"page_num,omitempty"`
	// PageSize is the number of items per page.
	// PageSize 每页数量。
	PageSize int32 `json:"page_size,omitempty"`
}

// ListConnectedStoresResponse represents the response for listing connected stores.
// ListConnectedStoresResponse 已连接店铺列表响应。
type ListConnectedStoresResponse struct {
	// HasMore indicates whether there are more results.
	// HasMore 是否还有更多结果。
	HasMore bool `json:"has_more,omitempty"`
	// Items is the list of connected stores.
	// Items 已连接店铺列表。
	Items []ConnectedStore `json:"items,omitempty"`
}

// ListConnectedStores lists connected stores.
// ListConnectedStores 列出已连接店铺。
// 官方文档: https://www.airwallex.com/docs/api/supporting_services/connected_stores/list.md
func (s *Service) ListConnectedStores(ctx context.Context, req *ListConnectedStoresRequest, opts ...sdk.RequestOption) (*ListConnectedStoresResponse, error) {
	var resp ListConnectedStoresResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/ecosystem/connected_stores", req, &resp, opts...)
	return &resp, err
}

// GetConnectedStore retrieves details of a connected store.
// GetConnectedStore 获取已连接店铺详情。
// 官方文档: https://www.airwallex.com/docs/api/supporting_services/connected_stores/retrieve.md
func (s *Service) GetConnectedStore(ctx context.Context, id string, opts ...sdk.RequestOption) (*ConnectedStore, error) {
	var resp ConnectedStore
	err := s.doer.Do(ctx, "GET", "/api/v1/ecosystem/connected_stores/"+id, nil, &resp, opts...)
	return &resp, err
}

// --- Deprecated ---
// CreateConnectedStore / UpdateConnectedStore 已移除。官方 API 仅支持 List/Get。
