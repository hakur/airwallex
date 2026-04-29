package core

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// Balance represents account balance information.
// Balance 表示账户余额信息。
type Balance struct {
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// AvailableAmount is the available amount. Required.
	// AvailableAmount 可用金额。必填。
	AvailableAmount float64 `json:"available_amount"`
	// PendingAmount is the pending amount. Required.
	// PendingAmount 待处理金额。必填。
	PendingAmount float64 `json:"pending_amount"`
	// TotalAmount is the total amount. Required.
	// TotalAmount 总金额。必填。
	TotalAmount float64 `json:"total_amount"`
	// FrozenAmount is the frozen amount. Required.
	// FrozenAmount 冻结金额。必填。
	FrozenAmount float64 `json:"frozen_amount"`
	// AccountType is the account type. Optional.
	// AccountType 账户类型。可选。
	AccountType string `json:"account_type,omitempty"`
	// PrepaymentAmount is the prepayment amount. Optional.
	// PrepaymentAmount 预付款金额。可选。
	PrepaymentAmount float64 `json:"prepayment_amount,omitempty"`
	// ReservedAmount is the reserved amount. Optional.
	// ReservedAmount 预留金额。可选。
	ReservedAmount float64 `json:"reserved_amount,omitempty"`
}

// BalanceHistoryItem represents a balance history record item.
// BalanceHistoryItem 余额历史记录项。
type BalanceHistoryItem struct {
	// AccountType is the account type. Optional.
	// AccountType 账户类型。可选。
	AccountType string `json:"account_type,omitempty"`
	// Amount is the balance change amount. Optional.
	// Amount 余额变动金额。可选。
	Amount float64 `json:"amount,omitempty"`
	// Balance is the balance after change. Optional.
	// Balance 变动后余额。可选。
	Balance float64 `json:"balance,omitempty"`
	// Currency is the currency code. Optional.
	// Currency 货币代码。可选。
	Currency string `json:"currency,omitempty"`
	// Description is the description. Optional.
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// Fee is the fee amount. Optional.
	// Fee 手续费金额。可选。
	Fee float64 `json:"fee,omitempty"`
	// ID is the unique identifier for the balance change. Optional.
	// ID 余额变动唯一标识符。可选。
	ID string `json:"id,omitempty"`
	// PostedAt is the time of the balance change. Optional.
	// PostedAt 余额变动时间。可选。
	PostedAt string `json:"posted_at,omitempty"`
	// Source is the source transaction ID. Optional.
	// Source 源交易ID。可选。
	Source string `json:"source,omitempty"`
	// SourceType is the source transaction type. Optional.
	// SourceType 源交易类型。可选。
	SourceType string `json:"source_type,omitempty"`
	// TransactionType is the transaction type. Optional.
	// TransactionType 交易类型。可选。
	TransactionType string `json:"transaction_type,omitempty"`
}

// BalanceHistoryResponse represents the balance history response.
// BalanceHistoryResponse 余额历史响应。
type BalanceHistoryResponse struct {
	// HasMore indicates whether there are more results. Optional.
	// HasMore 是否有更多结果。可选。
	HasMore bool `json:"has_more,omitempty"`
	// Items is the list of balance history records. Optional.
	// Items 余额历史记录列表。可选。
	Items []BalanceHistoryItem `json:"items,omitempty"`
	// PageAfter is the next page cursor. Optional.
	// PageAfter 下一页游标。可选。
	PageAfter string `json:"page_after,omitempty"`
	// PageBefore is the previous page cursor. Optional.
	// PageBefore 上一页游标。可选。
	PageBefore string `json:"page_before,omitempty"`
}

// GetBalanceHistoryRequest represents the request to query balance history.
// GetBalanceHistoryRequest 查询余额历史请求。
type GetBalanceHistoryRequest struct {
	// AccountType is the account type. Optional.
	// AccountType 账户类型。可选。
	AccountType string `json:"account_type,omitempty"`
	// Currency is the currency code. Optional.
	// Currency 货币代码。可选。
	Currency sdk.Currency `json:"currency,omitempty"`
	// FromPostAt is the start time for the query (ISO8601). Optional.
	// FromPostAt 查询起始时间（ISO8601）。可选。
	FromPostAt string `json:"from_post_at,omitempty"`
	// Page is the pagination cursor. Optional.
	// Page 分页游标。可选。
	Page string `json:"page,omitempty"`
	// PageNum is the page number, starting from 0. Optional.
	// PageNum 页码，从0开始。可选。
	PageNum int32 `json:"page_num,omitempty"`
	// PageSize is the number of items per page. Optional.
	// PageSize 每页数量。可选。
	PageSize int32 `json:"page_size,omitempty"`
	// ToPostAt is the end time for the query (ISO8601). Optional.
	// ToPostAt 查询结束时间（ISO8601）。可选。
	ToPostAt string `json:"to_post_at,omitempty"`
}

// GetCurrentBalances retrieves current account balances.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/balances/current.md
// GetCurrentBalances 查询当前账户余额。
func (s *Service) GetCurrentBalances(ctx context.Context, opts ...sdk.RequestOption) ([]Balance, error) {
	var resp []Balance
	err := s.doer.Do(ctx, "GET", "/api/v1/balances/current", nil, &resp, opts...)
	return resp, err
}

// GetBalanceHistory retrieves balance history records.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/balances/history.md
// GetBalanceHistory 查询余额历史记录。
func (s *Service) GetBalanceHistory(ctx context.Context, req *GetBalanceHistoryRequest, opts ...sdk.RequestOption) (*BalanceHistoryResponse, error) {
	var resp BalanceHistoryResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/balances/history", req, &resp, opts...)
	return &resp, err
}
