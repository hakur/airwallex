package spend

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// ExpenseStatus represents the expense approval status.
// ExpenseStatus 费用审批状态。
type ExpenseStatus string

const (
	ExpenseStatusDraft            ExpenseStatus = "DRAFT"
	ExpenseStatusAwaitingApproval ExpenseStatus = "AWAITING_APPROVAL"
	ExpenseStatusRejected         ExpenseStatus = "REJECTED"
	ExpenseStatusApproved         ExpenseStatus = "APPROVED"
	ExpenseStatusArchived         ExpenseStatus = "ARCHIVED"
	ExpenseStatusDeleted          ExpenseStatus = "DELETED"
)

// ExpenseSyncStatus represents the expense sync status.
// ExpenseSyncStatus 费用同步状态。
type ExpenseSyncStatus string

const (
	ExpenseSyncStatusNotSynced   ExpenseSyncStatus = "NOT_SYNCED"
	ExpenseSyncStatusReadyToSync ExpenseSyncStatus = "READY_TO_SYNC"
	ExpenseSyncStatusSynced      ExpenseSyncStatus = "SYNCED"
	ExpenseSyncStatusSyncFailed  ExpenseSyncStatus = "SYNC_FAILED"
)

// ExpenseCardTransactionStatus represents the card transaction status.
// ExpenseCardTransactionStatus 卡交易状态。
type ExpenseCardTransactionStatus string

const (
	ExpenseCardTxStatusAuthorized ExpenseCardTransactionStatus = "AUTHORIZED"
	ExpenseCardTxStatusCleared    ExpenseCardTransactionStatus = "CLEARED"
	ExpenseCardTxStatusDeclined   ExpenseCardTransactionStatus = "DECLINED"
	ExpenseCardTxStatusReversed   ExpenseCardTransactionStatus = "REVERSED"
)

// Expense represents an expense.
// Expense 表示费用。
type Expense struct {
	ID                        string                     `json:"id"`
	AccountID                 string                     `json:"account_id"`
	LegalEntityID             string                     `json:"legal_entity_id"`
	CardID                    string                     `json:"card_id"`
	Status                    ExpenseStatus              `json:"status"`
	CardTransaction           ExpenseCardTransaction     `json:"card_transaction"`
	SyncStatus                ExpenseSyncStatus          `json:"sync_status"`
	BillingCurrency           string                     `json:"billing_currency"`
	BillingAmount             string                     `json:"billing_amount"`
	Description               string                     `json:"description,omitempty"`
	Merchant                  string                     `json:"merchant,omitempty"`
	Approvers                 []string                   `json:"approvers"`
	CreatedAt                 string                     `json:"created_at"`
	UpdatedAt                 string                     `json:"updated_at"`
	SettledAt                 string                     `json:"settled_at,omitempty"`
	AccountingFieldSelections []AccountingFieldSelection `json:"accounting_field_selections"`
	Attachments               []Attachment               `json:"attachments"`
	LineItems                 []ExpenseLineItem          `json:"line_items"`
	Comments                  []Comment                  `json:"comments"`
}

// ExpenseCardTransaction represents card transaction details.
// ExpenseCardTransaction 卡交易详情。
type ExpenseCardTransaction struct {
	Status   ExpenseCardTransactionStatus `json:"status"`
	Amount   string                       `json:"amount"`
	Currency string                       `json:"currency"`
}

// ExpenseLineItem represents an expense line item.
// ExpenseLineItem 费用行项目。
type ExpenseLineItem struct {
	ID                        string                     `json:"id"`
	TransactionAmount         string                     `json:"transaction_amount,omitempty"`
	Description               string                     `json:"description,omitempty"`
	AccountingFieldSelections []AccountingFieldSelection `json:"accounting_field_selections"`
}

// SyncExpenseRequest represents a request to sync an expense.
// SyncExpenseRequest 同步费用请求。
type SyncExpenseRequest struct {
	SyncStatus ExpenseSyncStatus `json:"sync_status"`
}

// ListExpensesRequest represents query parameters for listing expenses.
// ListExpensesRequest 费用列表查询参数。
type ListExpensesRequest struct {
	Page          string   `json:"page,omitempty"`
	FromCreatedAt string   `json:"from_created_at,omitempty"`
	ToCreatedAt   string   `json:"to_created_at,omitempty"`
	Status        []string `json:"status,omitempty"`
	SyncStatus    []string `json:"sync_status,omitempty"`
	LegalEntityID string   `json:"legal_entity_id,omitempty"`
}

// ListExpensesResponse represents an expense list response with cursor pagination.
// ListExpensesResponse 费用列表响应（cursor 分页）。
type ListExpensesResponse struct {
	Items      []Expense `json:"items"`
	PageAfter  string    `json:"page_after,omitempty"`
	PageBefore string    `json:"page_before,omitempty"`
}

// ListExpenses lists all expenses.
// ListExpenses 列出费用。
// 官方文档: https://www.airwallex.com/docs/api/spend/expenses/list.md
func (s *Service) ListExpenses(ctx context.Context, req *ListExpensesRequest, opts ...sdk.RequestOption) (*ListExpensesResponse, error) {
	var resp ListExpensesResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/spend/expenses", req, &resp, opts...)
	return &resp, err
}

// GetExpense retrieves expense details.
// GetExpense 获取费用详情。
// 官方文档: https://www.airwallex.com/docs/api/spend/expenses/retrieve.md
func (s *Service) GetExpense(ctx context.Context, id string, opts ...sdk.RequestOption) (*Expense, error) {
	var resp Expense
	err := s.doer.Do(ctx, "GET", "/api/v1/spend/expenses/"+id, nil, &resp, opts...)
	return &resp, err
}

// SyncExpense updates the expense sync status.
// SyncExpense 更新费用同步状态。
// 官方文档: https://www.airwallex.com/docs/api/spend/expenses/sync.md
func (s *Service) SyncExpense(ctx context.Context, id string, req *SyncExpenseRequest, opts ...sdk.RequestOption) (*Expense, error) {
	var resp Expense
	err := s.doer.Do(ctx, "POST", "/api/v1/spend/expenses/"+id+"/sync", req, &resp, opts...)
	return &resp, err
}

// --- Deprecated ---
// CreateExpense / UpdateExpense 已移除。官方 API 仅支持 List/Get/Sync，不支持创建和更新费用。
