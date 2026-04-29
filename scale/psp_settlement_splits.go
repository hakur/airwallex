package scale

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PSPSettlementSplit represents a PSP settlement split.
// PSPSettlementSplit 表示 PSP 结算拆分。
type PSPSettlementSplit struct {
	ID              string `json:"id,omitempty"`
	Amount          string `json:"amount,omitempty"`
	SettlementType  string `json:"settlement_type,omitempty"`
	Status          string `json:"status,omitempty"`
	TargetAccountID string `json:"target_account_id,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
}

// SplitSettlementIntentRequest represents a request to split a PSP settlement intent.
// SplitSettlementIntentRequest 拆分 PSP 结算意图请求。
type SplitSettlementIntentRequest struct {
	RequestID string                `json:"request_id"`
	Splits    []SettlementSplitItem `json:"splits"`
}

// SettlementSplitItem represents a single settlement split item.
// SettlementSplitItem 单个结算拆分项。
type SettlementSplitItem struct {
	Amount          string `json:"amount"`
	Identifier      string `json:"identifier"`
	SettlementType  string `json:"settlement_type"`
	TargetAccountID string `json:"target_account_id"`
}

// SplitSettlementIntentResponse represents the response for splitting a PSP settlement intent.
// SplitSettlementIntentResponse 拆分 PSP 结算意图响应。
type SplitSettlementIntentResponse struct {
	PSPSettlementIntentID string `json:"psp_settlement_intent_id"`
	RequestID             string `json:"request_id"`
}

// ListPSPSettlementSplitsRequest represents parameters for listing PSP settlement splits.
// ListPSPSettlementSplitsRequest 结算拆分列表查询参数。
type ListPSPSettlementSplitsRequest struct {
	Page     string `json:"page,omitempty"`
	PageSize int32  `json:"page_size,omitempty"`
}

// ListPSPSettlementSplitsResponse represents a response listing PSP settlement splits.
// ListPSPSettlementSplitsResponse 结算拆分列表响应。
type ListPSPSettlementSplitsResponse struct {
	Items      []PSPSettlementSplit `json:"items"`
	PageAfter  string               `json:"page_after,omitempty"`
	PageBefore string               `json:"page_before,omitempty"`
}

// SplitPSPSettlementIntent splits a PSP settlement intent (batch creates settlement splits).
// SplitPSPSettlementIntent 拆分 PSP 结算意图（批量创建结算拆分）。
// 官方文档: https://www.airwallex.com/docs/api/scale/psp_settlement_splits/split_psp_settlement_intents.md
func (s *Service) SplitPSPSettlementIntent(ctx context.Context, intentID string, req *SplitSettlementIntentRequest, opts ...sdk.RequestOption) (*SplitSettlementIntentResponse, error) {
	var resp SplitSettlementIntentResponse
	err := s.doer.Do(ctx, "POST", "/api/v1/psp_settlement_intents/"+intentID+"/split", req, &resp, opts...)
	return &resp, err
}

// GetPSPSettlementSplit retrieves a PSP settlement split by ID.
// GetPSPSettlementSplit 获取 PSP 结算拆分详情。
// 官方文档: https://www.airwallex.com/docs/api/scale/psp_settlement_splits/retrieve.md
func (s *Service) GetPSPSettlementSplit(ctx context.Context, id string, opts ...sdk.RequestOption) (*PSPSettlementSplit, error) {
	var resp PSPSettlementSplit
	err := s.doer.Do(ctx, "GET", "/api/v1/psp_settlement_splits/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListPSPSettlementSplits lists PSP settlement splits.
// ListPSPSettlementSplits 列出 PSP 结算拆分。
// 官方文档: https://www.airwallex.com/docs/api/scale/psp_settlement_splits/list.md
func (s *Service) ListPSPSettlementSplits(ctx context.Context, req *ListPSPSettlementSplitsRequest, opts ...sdk.RequestOption) (*ListPSPSettlementSplitsResponse, error) {
	var resp ListPSPSettlementSplitsResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/psp_settlement_splits", req, &resp, opts...)
	return &resp, err
}

// CancelPSPSettlementSplit cancels a PSP settlement split.
// CancelPSPSettlementSplit 取消 PSP 结算拆分。
// 官方文档: https://www.airwallex.com/docs/api/scale/psp_settlement_splits/cancel.md
func (s *Service) CancelPSPSettlementSplit(ctx context.Context, id string, opts ...sdk.RequestOption) (*PSPSettlementSplit, error) {
	var resp PSPSettlementSplit
	err := s.doer.Do(ctx, "POST", "/api/v1/psp_settlement_splits/"+id+"/cancel", nil, &resp, opts...)
	return &resp, err
}

// ReleasePSPSettlementSplit releases a PSP settlement split.
// ReleasePSPSettlementSplit 释放 PSP 结算拆分。
// 官方文档: https://www.airwallex.com/docs/api/scale/psp_settlement_splits/release.md
func (s *Service) ReleasePSPSettlementSplit(ctx context.Context, id string, opts ...sdk.RequestOption) (*PSPSettlementSplit, error) {
	var resp PSPSettlementSplit
	err := s.doer.Do(ctx, "POST", "/api/v1/psp_settlement_splits/"+id+"/release", nil, &resp, opts...)
	return &resp, err
}

// --- Deprecated ---
// CreatePSPSettlementSplit 已移除。请使用 SplitPSPSettlementIntent（路径为 /psp_settlement_intents/{id}/split）。
