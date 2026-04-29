package core

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// DepositType represents the type of deposit.
// DepositType 存款类型。
type DepositType = string

const (
	// DepositTypeBankTransfer is a bank transfer deposit.
	// DepositTypeBankTransfer 银行转账存款
	DepositTypeBankTransfer DepositType = "BANK_TRANSFER"
	// DepositTypeDirectDebit is a direct debit deposit.
	// DepositTypeDirectDebit 直接扣款存款
	DepositTypeDirectDebit DepositType = "DIRECT_DEBIT"
	// DepositTypeDigitalWalletTransfer is a digital wallet transfer deposit.
	// DepositTypeDigitalWalletTransfer 数字钱包转账存款
	DepositTypeDigitalWalletTransfer DepositType = "DIGITAL_WALLET_TRANSFER"
	// DepositTypeAggregateFunds is an aggregate funds deposit.
	// DepositTypeAggregateFunds 聚合资金存款
	DepositTypeAggregateFunds DepositType = "AGGREGATE_FUNDS"
)

// DepositProviderFailureDetails represents provider failure details.
// DepositProviderFailureDetails 提供商失败详情。
type DepositProviderFailureDetails struct {
	Code                string `json:"code,omitempty"`
	LocalClearingSystem string `json:"local_clearing_system,omitempty"`
	Message             string `json:"message,omitempty"`
}

// DepositFailureDetails represents deposit failure details.
// DepositFailureDetails 存款失败详情。
type DepositFailureDetails struct {
	Code                   string                         `json:"code,omitempty"`
	ISOCode                string                         `json:"iso_code,omitempty"`
	ProviderFailureDetails *DepositProviderFailureDetails `json:"provider_failure_details,omitempty"`
}

// DepositFee represents deposit fee information.
// DepositFee 存款费用。
type DepositFee struct {
	Amount   float64      `json:"amount,omitempty"`
	Currency sdk.Currency `json:"currency,omitempty"`
}

// DigitalWallet represents digital wallet information.
// DigitalWallet 数字钱包信息。
type DigitalWallet struct {
	AccountName string `json:"account_name"`
	IDType      string `json:"id_type,omitempty"`
	IDValue     string `json:"id_value,omitempty"`
	Provider    string `json:"provider"`
}

// DepositPayer represents the deposit payer information.
// DepositPayer 存款付款人信息。
type DepositPayer struct {
	BankAccount   map[string]any `json:"bank_account,omitempty"`
	CountryCode   string         `json:"country_code,omitempty"`
	DigitalWallet *DigitalWallet `json:"digital_wallet,omitempty"`
	Name          string         `json:"name,omitempty"`
}

// Deposit represents deposit information.
// Deposit 表示存款信息。
type Deposit struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Amount is the amount. Required.
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// Status is the status. Required.
	// Status 状态。必填。
	Status string `json:"status"`
	// Type is the deposit type. Optional.
	// Type 存款类型。可选。
	Type DepositType `json:"type,omitempty"`
	// CreatedAt is the creation time. Required.
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// EstimatedSettledAt is the estimated settlement time. Optional.
	// EstimatedSettledAt 预计结算时间。可选。
	EstimatedSettledAt string `json:"estimated_settled_at,omitempty"`
	// SettledAt is the actual settlement time. Optional.
	// SettledAt 实际结算时间。可选。
	SettledAt string `json:"settled_at,omitempty"`
	// Reference is the user-specified reference number. Optional.
	// Reference 用户指定的参考号。可选。
	Reference string `json:"reference,omitempty"`
	// ProviderTransactionID is the provider transaction ID. Optional.
	// ProviderTransactionID 提供商交易ID。可选。
	ProviderTransactionID string `json:"provider_transaction_id,omitempty"`
	// FundingSourceID is the funding source ID. Optional.
	// FundingSourceID 资金来源ID。可选。
	FundingSourceID string `json:"funding_source_id,omitempty"`
	// GlobalAccountID is the global account ID. Optional.
	// GlobalAccountID 全局账户ID。可选。
	GlobalAccountID string `json:"global_account_id,omitempty"`
	// FailureDetails contains failure details. Optional.
	// FailureDetails 失败详情。可选。
	FailureDetails *DepositFailureDetails `json:"failure_details,omitempty"`
	// Fee contains fee information. Optional.
	// Fee 费用信息。可选。
	Fee *DepositFee `json:"fee,omitempty"`
	// Payer contains payer information. Optional.
	// Payer 付款人信息。可选。
	Payer *DepositPayer `json:"payer,omitempty"`
}

// CreateDepositRequest represents the request to create a deposit.
// CreateDepositRequest 创建存款请求。
type CreateDepositRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Amount is the amount. Required.
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// DepositType is the deposit type. Optional.
	// DepositType 存款类型。可选。
	DepositType DepositType `json:"deposit_type,omitempty"`
	// FundingSourceID is the funding source ID. Required.
	// FundingSourceID 资金来源ID。必填。
	FundingSourceID string `json:"funding_source_id"`
	// Reference is the user-specified reference number. Optional.
	// Reference 用户指定的参考号。可选。
	Reference string `json:"reference,omitempty"`
}

// ListDepositsRequest represents the request to list deposits.
// ListDepositsRequest 列出存款请求。
type ListDepositsRequest struct {
	// FromCreatedAt is the start creation time (ISO8601). Optional.
	// FromCreatedAt 创建时间起始（ISO8601）。可选。
	FromCreatedAt string `json:"from_created_at,omitempty"`
	// PageNum is the page number, starting from 0. Optional.
	// PageNum 页码，从0开始。可选。
	PageNum int32 `json:"page_num,omitempty"`
	// PageSize is the number of items per page. Optional.
	// PageSize 每页数量。可选。
	PageSize int32 `json:"page_size,omitempty"`
	// ToCreatedAt is the end creation time (ISO8601). Optional.
	// ToCreatedAt 创建时间截止（ISO8601）。可选。
	ToCreatedAt string `json:"to_created_at,omitempty"`
}

// CreateDeposit creates a deposit.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/deposits/create.md
// CreateDeposit 创建存款。
func (s *Service) CreateDeposit(ctx context.Context, req *CreateDepositRequest, opts ...sdk.RequestOption) (*Deposit, error) {
	var resp Deposit
	err := s.doer.Do(ctx, "POST", "/api/v1/deposits/create", req, &resp, opts...)
	return &resp, err
}

// GetDeposit retrieves a deposit by ID.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/deposits/retrieve.md
// GetDeposit 根据 ID 获取存款。
func (s *Service) GetDeposit(ctx context.Context, id string, opts ...sdk.RequestOption) (*Deposit, error) {
	var resp Deposit
	err := s.doer.Do(ctx, "GET", "/api/v1/deposits/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListDeposits lists deposits.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/deposits/list.md
// ListDeposits 列出存款。
func (s *Service) ListDeposits(ctx context.Context, req *ListDepositsRequest, opts ...sdk.RequestOption) ([]Deposit, error) {
	var resp []Deposit
	err := s.doer.Do(ctx, "GET", "/api/v1/deposits", req, &resp, opts...)
	return resp, err
}
