package core

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// AccountType represents the global account type.
// AccountType 全局账户类型。
type AccountType string

const (
	// AccountTypeStandard is a standard account.
	// AccountTypeStandard 标准账户
	AccountTypeStandard AccountType = "STANDARD"
	// AccountTypeVirtual is a virtual account.
	// AccountTypeVirtual 虚拟账户
	AccountTypeVirtual AccountType = "VIRTUAL"
)

// GlobalAccountStatus represents the global account status.
// GlobalAccountStatus 全局账户状态。
type GlobalAccountStatus string

const (
	// GlobalAccountStatusActive indicates the account is active.
	// GlobalAccountStatusActive 活跃状态
	GlobalAccountStatusActive GlobalAccountStatus = "ACTIVE"
	// GlobalAccountStatusPending indicates the account is pending.
	// GlobalAccountStatusPending 待处理
	GlobalAccountStatusPending GlobalAccountStatus = "PENDING"
	// GlobalAccountStatusSuspended indicates the account is suspended.
	// GlobalAccountStatusSuspended 已暂停
	GlobalAccountStatusSuspended GlobalAccountStatus = "SUSPENDED"
	// GlobalAccountStatusProcessing indicates the account is processing.
	// GlobalAccountStatusProcessing 处理中
	GlobalAccountStatusProcessing GlobalAccountStatus = "PROCESSING"
	// GlobalAccountStatusClosed indicates the account is closed.
	// GlobalAccountStatusClosed 已关闭
	GlobalAccountStatusClosed GlobalAccountStatus = "CLOSED"
	// GlobalAccountStatusFailed indicates the account failed.
	// GlobalAccountStatusFailed 失败
	GlobalAccountStatusFailed GlobalAccountStatus = "FAILED"
)

// Institution represents financial institution information.
// Institution 金融机构信息。
type Institution struct {
	// Name is the institution name. Optional.
	// Name 机构名称。可选。
	Name string `json:"name,omitempty"`
	// CountryCode is the country code. Optional.
	// CountryCode 国家代码。可选。
	CountryCode string `json:"country_code,omitempty"`
	// BankCode is the bank code. Optional.
	// BankCode 银行代码。可选。
	BankCode string `json:"bank_code,omitempty"`
	// BranchCode is the branch code. Optional.
	// BranchCode 支行代码。可选。
	BranchCode string `json:"branch_code,omitempty"`
}

// AlternateAccountIdentifier represents an alternate account identifier.
// AlternateAccountIdentifier 备用账户标识。
type AlternateAccountIdentifier struct {
	// Type is the identifier type. Optional.
	// Type 标识类型。可选。
	Type string `json:"type,omitempty"`
	// Value is the identifier value. Optional.
	// Value 标识值。可选。
	Value string `json:"value,omitempty"`
}

// GlobalAccount represents global account information.
// GlobalAccount 表示全球账户信息。
type GlobalAccount struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// CountryCode is the country code. Required.
	// CountryCode 国家代码。必填。
	CountryCode sdk.CountryCode `json:"country_code"`
	// AccountType is the account type. Required.
	// AccountType 账户类型。必填。
	AccountType AccountType `json:"account_type"`
	// Status is the status. Required.
	// Status 状态。必填。
	Status GlobalAccountStatus `json:"status"`
	// AccountName is the account name. Required.
	// AccountName 账户名称。必填。
	AccountName string `json:"account_name"`
	// NickName is the nickname. Optional.
	// NickName 昵称。可选。
	NickName string `json:"nick_name,omitempty"`
	// AccountNumber is the account number. Required.
	// AccountNumber 账号。必填。
	AccountNumber string `json:"account_number"`
	// BankName is the bank name. Required.
	// BankName 银行名称。必填。
	BankName string `json:"bank_name"`
	// BankSwiftCode is the bank SWIFT code. Optional.
	// BankSwiftCode 银行 SWIFT 代码。可选。
	BankSwiftCode string `json:"bank_swift_code,omitempty"`
	// IBAN is the International Bank Account Number. Optional.
	// IBAN 国际银行账号。可选。
	IBAN string `json:"iban,omitempty"`
	// Institution is the financial institution information. Optional.
	// Institution 金融机构信息。可选。
	Institution *Institution `json:"institution,omitempty"`
	// SupportedFeatures is the list of supported features. Optional.
	// SupportedFeatures 支持的功能列表。可选。
	SupportedFeatures []map[string]any `json:"supported_features,omitempty"`
	// AlternateAccountIdentifiers is the list of alternate account identifiers. Optional.
	// AlternateAccountIdentifiers 备用账户标识列表。可选。
	AlternateAccountIdentifiers []AlternateAccountIdentifier `json:"alternate_account_identifiers,omitempty"`
	// CloseReason is the reason for closing. Optional.
	// CloseReason 关闭原因。可选。
	CloseReason string `json:"close_reason,omitempty"`
	// FailureReason is the reason for failure. Optional.
	// FailureReason 失败原因。可选。
	FailureReason string `json:"failure_reason,omitempty"`
	// RequiredFeatures is the list of required features. Optional.
	// RequiredFeatures 所需功能列表。可选。
	RequiredFeatures []map[string]any `json:"required_features,omitempty"`
	// DepositConversionCurrency is the deposit conversion currency code. Optional.
	// DepositConversionCurrency 存款转换货币代码。可选。
	DepositConversionCurrency sdk.Currency `json:"deposit_conversion_currency,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt is the update time. Optional.
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
	// RequestID is the request identifier. Optional.
	// RequestID 请求标识符。可选。
	RequestID string `json:"request_id,omitempty"`
}

// CreateGlobalAccountRequest represents the request to create a global account.
// CreateGlobalAccountRequest 创建全球账户请求。
type CreateGlobalAccountRequest struct {
	// RequestID is the request identifier. Required.
	// RequestID 请求标识符。必填。
	RequestID string `json:"request_id"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// CountryCode is the country code. Required.
	// CountryCode 国家代码。必填。
	CountryCode sdk.CountryCode `json:"country_code"`
	// AccountType is the account type. Optional.
	// AccountType 账户类型。可选。
	AccountType AccountType `json:"account_type,omitempty"`
	// NickName is the nickname. Optional.
	// NickName 昵称。可选。
	NickName string `json:"nick_name,omitempty"`
	// DepositConversionCurrency is the deposit conversion currency code. Optional.
	// DepositConversionCurrency 存款转换货币代码。可选。
	DepositConversionCurrency sdk.Currency `json:"deposit_conversion_currency,omitempty"`
}

// UpdateGlobalAccountRequest represents the request to update a global account.
// UpdateGlobalAccountRequest 更新全球账户请求。
type UpdateGlobalAccountRequest struct {
	// NickName is the nickname. Optional.
	// NickName 昵称。可选。
	NickName string `json:"nick_name,omitempty"`
}

// CreateGlobalAccount creates a global account.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/global_accounts/create.md
// CreateGlobalAccount 创建全球账户。
func (s *Service) CreateGlobalAccount(ctx context.Context, req *CreateGlobalAccountRequest, opts ...sdk.RequestOption) (*GlobalAccount, error) {
	var resp GlobalAccount
	err := s.doer.Do(ctx, "POST", "/api/v1/global_accounts/create", req, &resp, opts...)
	return &resp, err
}

// GetGlobalAccount retrieves a global account by ID.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/global_accounts/retrieve.md
// GetGlobalAccount 根据 ID 获取全球账户。
func (s *Service) GetGlobalAccount(ctx context.Context, id string, opts ...sdk.RequestOption) (*GlobalAccount, error) {
	var resp GlobalAccount
	err := s.doer.Do(ctx, "GET", "/api/v1/global_accounts/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListGlobalAccounts lists global accounts.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/global_accounts/list.md
// ListGlobalAccounts 列出全球账户。
func (s *Service) ListGlobalAccounts(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[GlobalAccount], error) {
	var resp sdk.ListResult[GlobalAccount]
	err := s.doer.Do(ctx, "GET", "/api/v1/global_accounts", nil, &resp, opts...)
	return &resp, err
}

// UpdateGlobalAccount updates a global account.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/global_accounts/update.md
// UpdateGlobalAccount 更新全球账户。
func (s *Service) UpdateGlobalAccount(ctx context.Context, id string, req *UpdateGlobalAccountRequest, opts ...sdk.RequestOption) (*GlobalAccount, error) {
	var resp GlobalAccount
	err := s.doer.Do(ctx, "POST", "/api/v1/global_accounts/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// DeleteGlobalAccount closes a global account.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/global_accounts/close.md
// DeleteGlobalAccount 关闭全球账户。
func (s *Service) DeleteGlobalAccount(ctx context.Context, id string, opts ...sdk.RequestOption) (*GlobalAccount, error) {
	var resp GlobalAccount
	err := s.doer.Do(ctx, "POST", "/api/v1/global_accounts/"+id+"/close", nil, &resp, opts...)
	return &resp, err
}

// GetGlobalAccountBankDetails retrieves bank details for a global account.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/global_accounts/retrieve.md
// GetGlobalAccountBankDetails 获取全球账户银行详情。
func (s *Service) GetGlobalAccountBankDetails(ctx context.Context, id string, opts ...sdk.RequestOption) (map[string]any, error) {
	var resp map[string]any
	err := s.doer.Do(ctx, "GET", "/api/v1/global_accounts/"+id+"/bank_details", nil, &resp, opts...)
	return resp, err
}

// GlobalAccountTransaction represents a global account transaction record.
// GlobalAccountTransaction 全球账户交易记录。
type GlobalAccountTransaction struct {
	// Amount is the transaction amount. Required.
	// Amount 交易金额。必填。
	Amount float64 `json:"amount"`
	// CreateTime is the transaction creation time. Optional.
	// CreateTime 交易创建时间。可选。
	CreateTime string `json:"create_time,omitempty"`
	// Currency is the transaction currency code. Required.
	// Currency 交易货币代码。必填。
	Currency string `json:"currency"`
	// Description is the transaction description. Optional.
	// Description 交易描述。可选。
	Description string `json:"description,omitempty"`
	// FeeAmount is the fee amount. Optional.
	// FeeAmount 手续费金额。可选。
	FeeAmount float64 `json:"fee_amount,omitempty"`
	// FeeCurrency is the fee currency code. Optional.
	// FeeCurrency 手续费货币代码。可选。
	FeeCurrency string `json:"fee_currency,omitempty"`
	// ID is the unique transaction identifier. Optional.
	// ID 交易唯一标识符。可选。
	ID string `json:"id,omitempty"`
	// PayerCountry is the payer country code. Optional.
	// PayerCountry 付款人国家代码。可选。
	PayerCountry string `json:"payer_country,omitempty"`
	// PayerName is the payer name. Optional.
	// PayerName 付款人名称。可选。
	PayerName string `json:"payer_name,omitempty"`
	// Status is the transaction status. Optional.
	// Status 交易状态。可选。
	Status string `json:"status,omitempty"`
	// Type is the transaction type. Optional.
	// Type 交易类型。可选。
	Type string `json:"type,omitempty"`
}

// GetGlobalAccountTransactionsRequest represents the request to query global account transactions.
// GetGlobalAccountTransactionsRequest 查询全球账户交易请求。
type GetGlobalAccountTransactionsRequest struct {
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

// GetGlobalAccountTransactions retrieves global account transaction records.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/global_accounts/transactions.md
// GetGlobalAccountTransactions 获取全球账户交易记录。
func (s *Service) GetGlobalAccountTransactions(ctx context.Context, id string, req *GetGlobalAccountTransactionsRequest, opts ...sdk.RequestOption) (*sdk.ListResult[GlobalAccountTransaction], error) {
	var resp sdk.ListResult[GlobalAccountTransaction]
	err := s.doer.Do(ctx, "GET", "/api/v1/global_accounts/"+id+"/transactions", req, &resp, opts...)
	return &resp, err
}
