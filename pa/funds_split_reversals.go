package pa

import (
	"context"
	"net/url"
	"strconv"

	"github.com/hakur/airwallex/sdk"
)

// FundsSplitReversal represents a funds split reversal.
// FundsSplitReversal 表示资金拆分撤销信息。
type FundsSplitReversal struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// FundsSplitID is the funds split unique identifier. Required.
	// FundsSplitID 资金拆分唯一标识符。必填。
	FundsSplitID string `json:"funds_split_id"`
	// Amount is the reversal amount. Required.
	// Amount 撤销金额。必填。
	Amount string `json:"amount"`
	// Status is the reversal status. Required.
	// Status 撤销状态。必填。
	Status string `json:"status"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
}

// CreateFundsSplitReversalRequest is the request to create a funds split reversal.
// CreateFundsSplitReversalRequest 创建资金拆分撤销请求。
type CreateFundsSplitReversalRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// FundsSplitID is the funds split unique identifier. Required.
	// FundsSplitID 资金拆分唯一标识符。必填。
	FundsSplitID string `json:"funds_split_id"`
	// Amount is the reversal amount. Required.
	// Amount 撤销金额。必填。
	Amount string `json:"amount"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// ListFundsSplitReversalsRequest is the request to list funds split reversals.
// ListFundsSplitReversalsRequest 列出资金拆分撤销请求参数。
type ListFundsSplitReversalsRequest struct {
	// FundsSplitID is the funds split unique identifier. Optional (query param).
	// FundsSplitID 资金拆分唯一标识符。可选（查询参数）。
	FundsSplitID string `json:"-"`
	// PageNum is the page number. Optional (query param).
	// PageNum 页码。可选（查询参数）。
	PageNum int32 `json:"-"`
	// PageSize is the page size. Optional (query param).
	// PageSize 每页数量。可选（查询参数）。
	PageSize int32 `json:"-"`
}

// CreateFundsSplitReversal creates a funds split reversal.
// CreateFundsSplitReversal 创建资金拆分撤销。
// 官方文档: https://www.airwallex.com/docs/api/payments/funds_split_reversals/create.md
func (s *Service) CreateFundsSplitReversal(ctx context.Context, req *CreateFundsSplitReversalRequest, opts ...sdk.RequestOption) (*FundsSplitReversal, error) {
	var resp FundsSplitReversal
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/funds_split_reversals/create", req, &resp, opts...)
	return &resp, err
}

// GetFundsSplitReversal retrieves a funds split reversal by ID.
// GetFundsSplitReversal 根据 ID 获取资金拆分撤销。
// 官方文档: https://www.airwallex.com/docs/api/payments/funds_split_reversals/retrieve.md
func (s *Service) GetFundsSplitReversal(ctx context.Context, id string, opts ...sdk.RequestOption) (*FundsSplitReversal, error) {
	var resp FundsSplitReversal
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/funds_split_reversals/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListFundsSplitReversals lists funds split reversals.
// ListFundsSplitReversals 列出资金拆分撤销。
// 官方文档: https://www.airwallex.com/docs/api/payments/funds_split_reversals/list.md
func (s *Service) ListFundsSplitReversals(ctx context.Context, req *ListFundsSplitReversalsRequest, opts ...sdk.RequestOption) (*sdk.ListResult[FundsSplitReversal], error) {
	path := "/api/v1/pa/funds_split_reversals"
	if req != nil {
		q := url.Values{}
		if req.FundsSplitID != "" {
			q.Set("funds_split_id", req.FundsSplitID)
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
	var resp sdk.ListResult[FundsSplitReversal]
	err := s.doer.Do(ctx, "GET", path, nil, &resp, opts...)
	return &resp, err
}
