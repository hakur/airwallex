package scale

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// ProgramLimitBalance represents a program limit balance.
// ProgramLimitBalance 程序限额余额。
type ProgramLimitBalance struct {
	AvailableAmount float64 `json:"available_amount,omitempty"`
	Currency        string  `json:"currency,omitempty"`
	ReservedAmount  float64 `json:"reserved_amount,omitempty"`
	TotalAmount     float64 `json:"total_amount,omitempty"`
}

// ProgramSupportedCurrency represents a program supported currency configuration.
// ProgramSupportedCurrency 程序支持的货币配置。
type ProgramSupportedCurrency struct {
	Currency            string  `json:"currency,omitempty"`
	LowBalanceThreshold float64 `json:"low_balance_threshold,omitempty"`
}

// Program represents a platform liquidity program.
// Program 表示平台流动性计划。
type Program struct {
	ID                  string                     `json:"id,omitempty"`
	Name                string                     `json:"name,omitempty"`
	Status              string                     `json:"status,omitempty"`
	CreatedAt           string                     `json:"created_at,omitempty"`
	UpdatedAt           string                     `json:"updated_at,omitempty"`
	LimitBalances       []ProgramLimitBalance      `json:"limit_balances,omitempty"`
	SupportedCurrencies []ProgramSupportedCurrency `json:"supported_currencies,omitempty"`
}

// ProgramSpendingAccount represents a program spending account.
// ProgramSpendingAccount 程序支出账户。
type ProgramSpendingAccount struct {
	ID                 string `json:"id,omitempty"`
	ConnectedAccountID string `json:"connected_account_id,omitempty"`
	FundingSourceID    string `json:"funding_source_id,omitempty"`
	Status             string `json:"status,omitempty"`
	CreatedAt          string `json:"created_at,omitempty"`
	UpdatedAt          string `json:"updated_at,omitempty"`
}

// ProgramTransaction represents a program transaction record.
// ProgramTransaction 程序交易记录。
type ProgramTransaction struct {
	ID                 string  `json:"id,omitempty"`
	Amount             float64 `json:"amount,omitempty"`
	AvailableLimit     float64 `json:"available_limit,omitempty"`
	ConnectedAccountID string  `json:"connected_account_id,omitempty"`
	Currency           string  `json:"currency,omitempty"`
	PostedAt           string  `json:"posted_at,omitempty"`
	SourceID           string  `json:"source_id,omitempty"`
	TransactionType    string  `json:"transaction_type,omitempty"`
}

// DepositFundsRequest represents a request to deposit funds.
// DepositFundsRequest 存入资金请求。
type DepositFundsRequest struct {
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	RequestID string  `json:"request_id"`
}

// WithdrawFundsRequest represents a request to withdraw funds.
// WithdrawFundsRequest 提取资金请求。
type WithdrawFundsRequest struct {
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	RequestID string  `json:"request_id"`
}

// ListProgramSpendingAccountsResponse represents a response listing program spending accounts.
// ListProgramSpendingAccountsResponse 支出账户列表响应。
type ListProgramSpendingAccountsResponse struct {
	Items      []ProgramSpendingAccount `json:"items"`
	PageAfter  string                   `json:"page_after,omitempty"`
	PageBefore string                   `json:"page_before,omitempty"`
}

// ListProgramTransactionsRequest represents query parameters for listing program transactions.
// ListProgramTransactionsRequest 交易列表查询参数。
type ListProgramTransactionsRequest struct {
	ConnectedAccountID string `json:"connected_account_id,omitempty"`
	Currency           string `json:"currency,omitempty"`
	FromPostAt         string `json:"from_post_at,omitempty"`
	ToPostAt           string `json:"to_post_at,omitempty"`
	Page               string `json:"page,omitempty"`
	PageSize           int32  `json:"page_size,omitempty"`
}

// ListProgramTransactionsResponse represents a response listing program transactions.
// ListProgramTransactionsResponse 交易列表响应。
type ListProgramTransactionsResponse struct {
	Items      []ProgramTransaction `json:"items,omitempty"`
	PageAfter  string               `json:"page_after,omitempty"`
	PageBefore string               `json:"page_before,omitempty"`
}

// GetProgram retrieves a platform liquidity program by ID.
// GetProgram 获取程序详情。
// 官方文档: https://www.airwallex.com/docs/api/scale/platform_liquidity_programs/retrieve.md
func (s *Service) GetProgram(ctx context.Context, id string, opts ...sdk.RequestOption) (*Program, error) {
	var resp Program
	err := s.doer.Do(ctx, "GET", "/api/v1/platform_liquidity_programs/"+id, nil, &resp, opts...)
	return &resp, err
}

// DepositFunds deposits funds into a program.
// DepositFunds 存入资金到程序。
// 官方文档: https://www.airwallex.com/docs/api/scale/platform_liquidity_programs/deposit.md
func (s *Service) DepositFunds(ctx context.Context, id string, req *DepositFundsRequest, opts ...sdk.RequestOption) (*Program, error) {
	var resp Program
	err := s.doer.Do(ctx, "POST", "/api/v1/platform_liquidity_programs/"+id+"/deposit", req, &resp, opts...)
	return &resp, err
}

// WithdrawFunds withdraws funds from a program.
// WithdrawFunds 从程序提取资金。
// 官方文档: https://www.airwallex.com/docs/api/scale/platform_liquidity_programs/withdraw.md
func (s *Service) WithdrawFunds(ctx context.Context, id string, req *WithdrawFundsRequest, opts ...sdk.RequestOption) (*Program, error) {
	var resp Program
	err := s.doer.Do(ctx, "POST", "/api/v1/platform_liquidity_programs/"+id+"/withdraw", req, &resp, opts...)
	return &resp, err
}

// ListProgramSpendingAccounts lists program spending accounts.
// ListProgramSpendingAccounts 列出程序支出账户。
// 官方文档: https://www.airwallex.com/docs/api/scale/platform_liquidity_programs/program_spending_accounts.md
func (s *Service) ListProgramSpendingAccounts(ctx context.Context, id string, opts ...sdk.RequestOption) (*ListProgramSpendingAccountsResponse, error) {
	var resp ListProgramSpendingAccountsResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/platform_liquidity_programs/"+id+"/program_spending_accounts", nil, &resp, opts...)
	return &resp, err
}

// ListProgramTransactions lists program transaction records.
// ListProgramTransactions 列出程序交易记录。
// 官方文档: https://www.airwallex.com/docs/api/scale/platform_liquidity_programs/transactions.md
func (s *Service) ListProgramTransactions(ctx context.Context, id string, req *ListProgramTransactionsRequest, opts ...sdk.RequestOption) (*ListProgramTransactionsResponse, error) {
	var resp ListProgramTransactionsResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/platform_liquidity_programs/"+id+"/transactions", req, &resp, opts...)
	return &resp, err
}

// --- Deprecated ---
// CreatePlatformLiquidityProgram / GetPlatformLiquidityProgram / ListPlatformLiquidityPrograms 已移除。
