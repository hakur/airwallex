package risk

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// WatchlistAction represents the watchlist action type.
// WatchlistAction Watchlist动作类型。
type WatchlistAction = string

const (
	// WatchlistActionBlock blocks the transaction.
	// WatchlistActionBlock 阻止交易。
	WatchlistActionBlock WatchlistAction = "BLOCK"
	// WatchlistActionVerify verifies the transaction.
	// WatchlistActionVerify 验证交易。
	WatchlistActionVerify WatchlistAction = "VERIFY"
	// WatchlistActionAccept accepts the transaction.
	// WatchlistActionAccept 接受交易。
	WatchlistActionAccept WatchlistAction = "ACCEPT"
)

// WatchlistItemType represents the watchlist item type.
// WatchlistItemType Watchlist条目类型。
type WatchlistItemType = string

const (
	// WatchlistItemTypeCardBin is the card BIN type.
	// WatchlistItemTypeCardBin 银行卡BIN。
	WatchlistItemTypeCardBin WatchlistItemType = "CARD_BIN"
	// WatchlistItemTypeCardFingerprint is the card fingerprint type.
	// WatchlistItemTypeCardFingerprint 银行卡指纹。
	WatchlistItemTypeCardFingerprint WatchlistItemType = "CARD_FINGERPRINT"
	// WatchlistItemTypeCardIssuerCountryCode is the card issuer country code type.
	// WatchlistItemTypeCardIssuerCountryCode 发卡国家代码。
	WatchlistItemTypeCardIssuerCountryCode WatchlistItemType = "CARD_ISSUER_COUNTRY_CODE"
	// WatchlistItemTypeCustomerEmail is the customer email type.
	// WatchlistItemTypeCustomerEmail 客户邮箱。
	WatchlistItemTypeCustomerEmail WatchlistItemType = "CUSTOMER_EMAIL"
	// WatchlistItemTypeCustomerID is the customer ID type.
	// WatchlistItemTypeCustomerID 客户ID。
	WatchlistItemTypeCustomerID WatchlistItemType = "CUSTOMER_ID"
)

// WatchlistItemStatus represents the watchlist item status.
// WatchlistItemStatus Watchlist条目状态。
type WatchlistItemStatus = string

const (
	// WatchlistItemStatusActive is the active status.
	// WatchlistItemStatusActive 激活状态。
	WatchlistItemStatusActive WatchlistItemStatus = "ACTIVE"
	// WatchlistItemStatusInactive is the inactive status.
	// WatchlistItemStatusInactive 非激活状态。
	WatchlistItemStatusInactive WatchlistItemStatus = "INACTIVE"
	// WatchlistItemStatusDeactivated is the deactivated status.
	// WatchlistItemStatusDeactivated 已停用状态。
	WatchlistItemStatusDeactivated WatchlistItemStatus = "DEACTIVATED"
)

// WatchlistEntry represents a watchlist entry.
// WatchlistEntry 表示Watchlist条目。
type WatchlistEntry struct {
	// Action is the action taken by the risk engine (BLOCK/VERIFY/ACCEPT).
	// Action 风险引擎执行的动作（BLOCK/VERIFY/ACCEPT）。
	Action WatchlistAction `json:"action,omitempty"`
	// CreatedAt is the entry creation timestamp.
	// CreatedAt 条目创建时间戳。
	CreatedAt string `json:"created_at,omitempty"`
	// ID is the unique entry identifier.
	// ID 条目唯一标识符。
	ID string `json:"id,omitempty"`
	// Reason is the reason for adding this entry to the watchlist.
	// Reason 添加此条目到Watchlist的原因。
	Reason string `json:"reason,omitempty"`
	// RequestID is the unique request identifier.
	// RequestID 请求唯一标识符。
	RequestID string `json:"request_id,omitempty"`
	// Status is the entry status (ACTIVE/INACTIVE/DEACTIVATED).
	// Status 条目状态（ACTIVE/INACTIVE/DEACTIVATED）。
	Status WatchlistItemStatus `json:"status,omitempty"`
	// Type is the entry type.
	// Type 条目类型。
	Type WatchlistItemType `json:"type,omitempty"`
	// UpdatedAt is the entry last update timestamp.
	// UpdatedAt 条目最后更新时间戳。
	UpdatedAt string `json:"updated_at,omitempty"`
	// Value is the value associated with the entry.
	// Value 条目对应的值。
	Value string `json:"value,omitempty"`
}

// CreateWatchlistEntryRequest represents a request to create a watchlist entry.
// CreateWatchlistEntryRequest 创建Watchlist条目请求。
type CreateWatchlistEntryRequest struct {
	// Action is the action taken by the risk engine (BLOCK/VERIFY/ACCEPT). Required.
	// Action 风险引擎执行的动作（BLOCK/VERIFY/ACCEPT）。必填。
	Action WatchlistAction `json:"action"`
	// Reason is the reason for adding this entry to the watchlist.
	// Reason 添加此条目到Watchlist的原因。
	Reason string `json:"reason,omitempty"`
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Status is the entry status (ACTIVE/INACTIVE, defaults to ACTIVE).
	// Status 条目状态（ACTIVE/INACTIVE，默认为ACTIVE）。
	Status WatchlistItemStatus `json:"status,omitempty"`
	// Type is the entry type. Required.
	// Type 条目类型。必填。
	Type WatchlistItemType `json:"type"`
	// Value is the value associated with the entry. Required.
	// Value 条目对应的值。必填。
	Value string `json:"value"`
}

// UpdateWatchlistEntryRequest represents a request to update a watchlist entry.
// UpdateWatchlistEntryRequest 更新Watchlist条目请求。
type UpdateWatchlistEntryRequest struct {
	// Reason is the reason for adding this entry to the watchlist.
	// Reason 添加此条目到Watchlist的原因。
	Reason string `json:"reason,omitempty"`
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Status is the entry status (ACTIVE/INACTIVE, defaults to ACTIVE).
	// Status 条目状态（ACTIVE/INACTIVE，默认为ACTIVE）。
	Status WatchlistItemStatus `json:"status,omitempty"`
	// Value is the value associated with the entry.
	// Value 条目对应的值。
	Value string `json:"value,omitempty"`
}

// CreateWatchlistEntry creates a watchlist entry.
// CreateWatchlistEntry 创建Watchlist条目。
// 官方文档: https://www.airwallex.com/docs/api/risk/watchlist/create.md
func (s *Service) CreateWatchlistEntry(ctx context.Context, req *CreateWatchlistEntryRequest, opts ...sdk.RequestOption) (*WatchlistEntry, error) {
	var resp WatchlistEntry
	err := s.doer.Do(ctx, "POST", "/api/v1/risk/watchlist/create", req, &resp, opts...)
	return &resp, err
}

// GetWatchlistEntry retrieves a watchlist entry by ID.
// GetWatchlistEntry 根据ID获取Watchlist条目。
// 官方文档: https://www.airwallex.com/docs/api/risk/watchlist/retrieve.md
func (s *Service) GetWatchlistEntry(ctx context.Context, id string, opts ...sdk.RequestOption) (*WatchlistEntry, error) {
	var resp WatchlistEntry
	err := s.doer.Do(ctx, "GET", "/api/v1/risk/watchlist/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateWatchlistEntry updates a watchlist entry.
// UpdateWatchlistEntry 更新Watchlist条目。
// 官方文档: https://www.airwallex.com/docs/api/risk/watchlist/update.md
func (s *Service) UpdateWatchlistEntry(ctx context.Context, id string, req *UpdateWatchlistEntryRequest, opts ...sdk.RequestOption) (*WatchlistEntry, error) {
	var resp WatchlistEntry
	err := s.doer.Do(ctx, "POST", "/api/v1/risk/watchlist/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ListWatchlistEntries lists all watchlist entries.
// ListWatchlistEntries 列出Watchlist条目。
// 官方文档: https://www.airwallex.com/docs/api/risk/watchlist/list.md
func (s *Service) ListWatchlistEntries(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[WatchlistEntry], error) {
	var resp sdk.ListResult[WatchlistEntry]
	err := s.doer.Do(ctx, "GET", "/api/v1/risk/watchlist", nil, &resp, opts...)
	return &resp, err
}
