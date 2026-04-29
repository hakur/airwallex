package pa

import (
	"context"
	"net/url"
	"strconv"

	"github.com/hakur/airwallex/sdk"
)

// FundsSplit represents a funds split.
// FundsSplit 表示资金拆分信息。
type FundsSplit struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// SourceID 源唯一标识符。必填。
	SourceID string `json:"source_id"`
	// SourceType 源类型。必填。
	SourceType string `json:"source_type"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// Amount is the amount. Required.
	// Amount 金额。必填。
	Amount string `json:"amount"`
	// Destination is the destination address. Required.
	// Destination 目标地址。必填。
	Destination string `json:"destination"`
	// Status is the split status. Required.
	// Status 拆分状态。必填。
	Status string `json:"status"`
	// AutoRelease indicates whether auto-release is enabled. Required.
	// AutoRelease 是否自动释放。必填。
	AutoRelease bool `json:"auto_release"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
}

// CreateFundsSplitRequest is the request to create a funds split.
// CreateFundsSplitRequest 创建资金拆分请求。
type CreateFundsSplitRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// SourceID is the source unique identifier. Required.
	// SourceID 源唯一标识符。必填。
	SourceID string `json:"source_id"`
	// SourceType is the source type. Required.
	// SourceType 源类型。必填。
	SourceType string `json:"source_type"`
	// Destination is the destination address. Required.
	// Destination 目标地址。必填。
	Destination string `json:"destination"`
	// Amount is the amount. Required.
	// Amount 金额。必填。
	Amount string `json:"amount"`
	// AutoRelease indicates whether auto-release is enabled. Optional.
	// AutoRelease 是否自动释放。可选。
	AutoRelease *bool `json:"auto_release,omitempty"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// ListFundsSplitsRequest is the request to list funds splits.
// ListFundsSplitsRequest 列出资金拆分请求参数。
type ListFundsSplitsRequest struct {
	// SourceID is the source unique identifier. Optional (query param).
	// SourceID 源唯一标识符。可选（查询参数）。
	SourceID string `json:"-"`
	// SourceType is the source type. Optional (query param).
	// SourceType 源类型。可选（查询参数）。
	SourceType string `json:"-"`
	// PageNum is the page number. Optional (query param).
	// PageNum 页码。可选（查询参数）。
	PageNum int32 `json:"-"`
	// PageSize is the page size. Optional (query param).
	// PageSize 每页数量。可选（查询参数）。
	PageSize int32 `json:"-"`
}

// CreateFundsSplit creates a funds split.
// CreateFundsSplit 创建资金拆分。
// 官方文档: https://www.airwallex.com/docs/api/payments/funds_splits/create.md
func (s *Service) CreateFundsSplit(ctx context.Context, req *CreateFundsSplitRequest, opts ...sdk.RequestOption) (*FundsSplit, error) {
	var resp FundsSplit
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/funds_splits/create", req, &resp, opts...)
	return &resp, err
}

// GetFundsSplit retrieves a funds split by ID.
// GetFundsSplit 根据 ID 获取资金拆分。
// 官方文档: https://www.airwallex.com/docs/api/payments/funds_splits/retrieve.md
func (s *Service) GetFundsSplit(ctx context.Context, id string, opts ...sdk.RequestOption) (*FundsSplit, error) {
	var resp FundsSplit
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/funds_splits/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListFundsSplits lists funds splits.
// ListFundsSplits 列出资金拆分。
// 官方文档: https://www.airwallex.com/docs/api/payments/funds_splits/list.md
func (s *Service) ListFundsSplits(ctx context.Context, req *ListFundsSplitsRequest, opts ...sdk.RequestOption) (*sdk.ListResult[FundsSplit], error) {
	path := "/api/v1/pa/funds_splits"
	if req != nil {
		q := url.Values{}
		if req.SourceID != "" {
			q.Set("source_id", req.SourceID)
		}
		if req.SourceType != "" {
			q.Set("source_type", req.SourceType)
		}
		if req.PageNum > 0 {
			q.Set("page_num", strconv.Itoa(int(req.PageNum)))
		}
		if req.PageSize > 0 {
			q.Set("page_size", strconv.Itoa(int(req.PageSize)))
		}
		if len(q) > 0 {
			path += "?" + q.Encode()
		}
	}
	var resp sdk.ListResult[FundsSplit]
	err := s.doer.Do(ctx, "GET", path, nil, &resp, opts...)
	return &resp, err
}
