package core

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// LinkedAccountType represents the linked account type.
// LinkedAccountType 关联账户类型。
type LinkedAccountType string

const (
	// LinkedAccountTypeAUBank is an Australian bank account.
	// LinkedAccountTypeAUBank 澳大利亚银行
	LinkedAccountTypeAUBank LinkedAccountType = "AU_BANK"
	// LinkedAccountTypeNZBank is a New Zealand bank account.
	// LinkedAccountTypeNZBank 新西兰银行
	LinkedAccountTypeNZBank LinkedAccountType = "NZ_BANK"
	// LinkedAccountTypeUSBank is a US bank account.
	// LinkedAccountTypeUSBank 美国银行
	LinkedAccountTypeUSBank LinkedAccountType = "US_BANK"
	// LinkedAccountTypeGBBank is a UK bank account.
	// LinkedAccountTypeGBBank 英国银行
	LinkedAccountTypeGBBank LinkedAccountType = "GB_BANK"
	// LinkedAccountTypeEUBank is an EU bank account.
	// LinkedAccountTypeEUBank 欧洲银行
	LinkedAccountTypeEUBank LinkedAccountType = "EU_BANK"
	// LinkedAccountTypeHKBank is a Hong Kong bank account.
	// LinkedAccountTypeHKBank 香港银行
	LinkedAccountTypeHKBank LinkedAccountType = "HK_BANK"
	// LinkedAccountTypeSGBank is a Singapore bank account.
	// LinkedAccountTypeSGBank 新加坡银行
	LinkedAccountTypeSGBank LinkedAccountType = "SG_BANK"
	// LinkedAccountTypeCABank is a Canadian bank account.
	// LinkedAccountTypeCABank 加拿大银行
	LinkedAccountTypeCABank LinkedAccountType = "CA_BANK"
	// LinkedAccountTypeAUPayID is an Australian PayID account.
	// LinkedAccountTypeAUPayID 澳大利亚PayID
	LinkedAccountTypeAUPayID LinkedAccountType = "AU_PAYID"
)

// LinkedAccountStatus represents the linked account status.
// LinkedAccountStatus 关联账户状态。
type LinkedAccountStatus string

const (
	// LinkedAccountStatusRequiresAction indicates action is required.
	// LinkedAccountStatusRequiresAction 需要操作
	LinkedAccountStatusRequiresAction LinkedAccountStatus = "REQUIRES_ACTION"
	// LinkedAccountStatusProcessing indicates the account is processing.
	// LinkedAccountStatusProcessing 处理中
	LinkedAccountStatusProcessing LinkedAccountStatus = "PROCESSING"
	// LinkedAccountStatusSucceeded indicates the account was created successfully.
	// LinkedAccountStatusSucceeded 成功
	LinkedAccountStatusSucceeded LinkedAccountStatus = "SUCCEEDED"
	// LinkedAccountStatusFailed indicates the account creation failed.
	// LinkedAccountStatusFailed 失败
	LinkedAccountStatusFailed LinkedAccountStatus = "FAILED"
	// LinkedAccountStatusSuspended indicates the account is suspended.
	// LinkedAccountStatusSuspended 已暂停
	LinkedAccountStatusSuspended LinkedAccountStatus = "SUSPENDED"
)

// AUBank represents Australian bank account information.
// AUBank 澳大利亚银行账户信息。
type AUBank struct {
	AccountName   string       `json:"account_name"`
	AccountNumber string       `json:"account_number"`
	BSB           string       `json:"bsb"`
	Currency      sdk.Currency `json:"currency"`
	EntityType    string       `json:"entity_type"`
}

// AUPayID represents Australian PayID account information.
// AUPayID 澳大利亚PayID账户信息。
type AUPayID struct {
	AccountIdentifierType string       `json:"account_identifier_type"`
	Currency              sdk.Currency `json:"currency"`
	EntityType            string       `json:"entity_type"`
	Payid                 string       `json:"payid"`
}

// CABank represents Canadian bank account information.
// CABank 加拿大银行账户信息。
type CABank struct {
	AccountName       string       `json:"account_name,omitempty"`
	AccountNumber     string       `json:"account_number,omitempty"`
	Currency          sdk.Currency `json:"currency"`
	EntityType        string       `json:"entity_type"`
	InstitutionNumber string       `json:"institution_number,omitempty"`
	TransitNumber     string       `json:"transit_number,omitempty"`
}

// EUBank represents European bank account information.
// EUBank 欧洲银行账户信息。
type EUBank struct {
	AccountName string       `json:"account_name"`
	Currency    sdk.Currency `json:"currency"`
	EntityType  string       `json:"entity_type"`
	IBAN        string       `json:"iban"`
	SwiftCode   string       `json:"swift_code,omitempty"`
}

// GBBank represents UK bank account information.
// GBBank 英国银行账户信息。
type GBBank struct {
	AccountName   string       `json:"account_name"`
	AccountNumber string       `json:"account_number"`
	Currency      sdk.Currency `json:"currency"`
	EntityType    string       `json:"entity_type"`
	SortCode      string       `json:"sort_code"`
}

// HKBank represents Hong Kong bank account information.
// HKBank 香港银行账户信息。
type HKBank struct {
	AccountName   string       `json:"account_name"`
	AccountNumber string       `json:"account_number"`
	BankCode      string       `json:"bank_code"`
	Currency      sdk.Currency `json:"currency"`
	EntityType    string       `json:"entity_type"`
}

// NZBank represents New Zealand bank account information.
// NZBank 新西兰银行账户信息。
type NZBank struct {
	AccountName   string       `json:"account_name"`
	AccountNumber string       `json:"account_number"`
	BankCode      string       `json:"bank_code,omitempty"`
	BranchCode    string       `json:"branch_code,omitempty"`
	Currency      sdk.Currency `json:"currency"`
	EntityType    string       `json:"entity_type"`
}

// SGBank represents Singapore bank account information.
// SGBank 新加坡银行账户信息。
type SGBank struct {
	AccountName   string       `json:"account_name"`
	AccountNumber string       `json:"account_number"`
	Currency      sdk.Currency `json:"currency"`
	EntityType    string       `json:"entity_type"`
	SwiftCode     string       `json:"swift_code,omitempty"`
}

// USBank represents US bank account information.
// USBank 美国银行账户信息。
type USBank struct {
	AccountName   string       `json:"account_name,omitempty"`
	AccountNumber string       `json:"account_number,omitempty"`
	AccountType   string       `json:"account_type,omitempty"`
	ACH           string       `json:"ach,omitempty"`
	Currency      sdk.Currency `json:"currency"`
	EntityType    string       `json:"entity_type"`
	Fedwire       string       `json:"fedwire,omitempty"`
}

// LinkedAccountCapabilities represents linked account capabilities.
// LinkedAccountCapabilities 关联账户能力。
type LinkedAccountCapabilities struct {
	BalanceCheck       bool `json:"balance_check,omitempty"`
	DirectDebitDeposit bool `json:"direct_debit_deposit,omitempty"`
}

// ProviderFailureDetails represents provider failure details.
// ProviderFailureDetails 提供商失败详情。
type ProviderFailureDetails struct {
	Code                string `json:"code,omitempty"`
	LocalClearingSystem string `json:"local_clearing_system,omitempty"`
	Message             string `json:"message,omitempty"`
}

// FailureDetails represents failure details.
// FailureDetails 失败详情。
type FailureDetails struct {
	Code                   string                  `json:"code,omitempty"`
	ISOCode                string                  `json:"iso_code,omitempty"`
	ProviderFailureDetails *ProviderFailureDetails `json:"provider_failure_details,omitempty"`
}

// MandateControls represents mandate control information.
// MandateControls 授权控制信息。
type MandateControls struct {
	AmountType                string  `json:"amount_type,omitempty"`
	EndDate                   string  `json:"end_date,omitempty"`
	FixedAmountPerTransaction float64 `json:"fixed_amount_per_transaction,omitempty"`
	MaxAmountPerTransaction   float64 `json:"max_amount_per_transaction,omitempty"`
}

// Mandate represents mandate information.
// Mandate 授权信息。
type Mandate struct {
	Controls              *MandateControls `json:"controls,omitempty"`
	Email                 string           `json:"email,omitempty"`
	PreferredReference    string           `json:"preferred_reference,omitempty"`
	ReuseMandateReference string           `json:"reuse_mandate_reference,omitempty"`
	Signatory             string           `json:"signatory,omitempty"`
	Type                  string           `json:"type"`
	Version               string           `json:"version,omitempty"`
}

// NextAction represents next action information.
// NextAction 下一步操作信息。
type NextAction struct {
	MicroDepositCount int32  `json:"micro_deposit_count,omitempty"`
	RemainingAttempts int32  `json:"remaining_attempts,omitempty"`
	Type              string `json:"type,omitempty"`
	VerificationType  string `json:"verification_type,omitempty"`
}

// PlaidInfo represents Plaid connection information.
// PlaidInfo Plaid连接信息。
type PlaidInfo struct {
	PublicToken string `json:"public_token"`
}

// TrueLayerInfo represents TrueLayer connection information.
// TrueLayerInfo TrueLayer连接信息。
type TrueLayerInfo struct {
	Code  string `json:"code"`
	State int64  `json:"state"`
}

// EGiroInfo represents EGiro connection information.
// EGiroInfo EGiro连接信息。
type EGiroInfo struct {
	AuthID        int64  `json:"auth_id"`
	AuthReference string `json:"auth_reference"`
}

// LinkedAccount represents linked account information.
// LinkedAccount 表示关联账户信息。
type LinkedAccount struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// AccountName is the account name. Optional.
	// AccountName 账户名称。可选。
	AccountName string `json:"account_name,omitempty"`
	// Currency is the currency code. Optional.
	// Currency 货币代码。可选。
	Currency sdk.Currency `json:"currency,omitempty"`
	// Status is the status. Optional.
	// Status 状态。可选。
	Status string `json:"status,omitempty"`
	// Type is the linked account type. Required.
	// Type 关联账户类型。必填。
	Type LinkedAccountType `json:"type"`
	// Reason is the status reason. Optional.
	// Reason 状态说明。可选。
	Reason string `json:"reason,omitempty"`
	// AUBank is the Australian bank account information. Optional.
	// AUBank 澳大利亚银行账户信息。可选。
	AUBank *AUBank `json:"au_bank,omitempty"`
	// AUPayID is the Australian PayID account information. Optional.
	// AUPayID 澳大利亚PayID账户信息。可选。
	AUPayID *AUPayID `json:"au_payid,omitempty"`
	// CABank is the Canadian bank account information. Optional.
	// CABank 加拿大银行账户信息。可选。
	CABank *CABank `json:"ca_bank,omitempty"`
	// EUBank is the European bank account information. Optional.
	// EUBank 欧洲银行账户信息。可选。
	EUBank *EUBank `json:"eu_bank,omitempty"`
	// GBBank is the UK bank account information. Optional.
	// GBBank 英国银行账户信息。可选。
	GBBank *GBBank `json:"gb_bank,omitempty"`
	// HKBank is the Hong Kong bank account information. Optional.
	// HKBank 香港银行账户信息。可选。
	HKBank *HKBank `json:"hk_bank,omitempty"`
	// NZBank is the New Zealand bank account information. Optional.
	// NZBank 新西兰银行账户信息。可选。
	NZBank *NZBank `json:"nz_bank,omitempty"`
	// SGBank is the Singapore bank account information. Optional.
	// SGBank 新加坡银行账户信息。可选。
	SGBank *SGBank `json:"sg_bank,omitempty"`
	// USBank is the US bank account information. Optional.
	// USBank 美国银行账户信息。可选。
	USBank *USBank `json:"us_bank,omitempty"`
	// Capabilities are the account capabilities. Optional.
	// Capabilities 账户能力。可选。
	Capabilities *LinkedAccountCapabilities `json:"capabilities,omitempty"`
	// FailureDetails contains failure details. Optional.
	// FailureDetails 失败详情。可选。
	FailureDetails *FailureDetails `json:"failure_details,omitempty"`
	// Mandate contains mandate information. Optional.
	// Mandate 授权信息。可选。
	Mandate *Mandate `json:"mandate,omitempty"`
	// NextAction is the next action to take. Optional.
	// NextAction 下一步操作。可选。
	NextAction *NextAction `json:"next_action,omitempty"`
	// Plaid contains Plaid connection information. Optional.
	// Plaid Plaid连接信息。可选。
	Plaid *PlaidInfo `json:"plaid,omitempty"`
	// TrueLayer contains TrueLayer connection information. Optional.
	// TrueLayer TrueLayer连接信息。可选。
	TrueLayer *TrueLayerInfo `json:"truelayer,omitempty"`
	// EGiro contains EGiro connection information. Optional.
	// EGiro EGiro连接信息。可选。
	EGiro *EGiroInfo `json:"egiro,omitempty"`
}

// CreateLinkedAccountRequest represents the request to create a linked account.
// CreateLinkedAccountRequest 创建关联账户请求。
type CreateLinkedAccountRequest struct {
	// AccountOwnerIdentifier is the account owner identifier. Optional.
	// AccountOwnerIdentifier 账户所有者标识符。可选。
	AccountOwnerIdentifier string `json:"account_owner_identifier,omitempty"`
	// AccountName is the account name. Optional.
	// AccountName 账户名称。可选。
	AccountName string `json:"account_name,omitempty"`
	// Currency is the currency code. Optional.
	// Currency 货币代码。可选。
	Currency sdk.Currency `json:"currency,omitempty"`
	// Type is the linked account type. Required.
	// Type 关联账户类型。必填。
	Type LinkedAccountType `json:"type"`
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// PreferredVerificationType is the preferred verification type. Optional.
	// PreferredVerificationType 首选验证类型。可选。
	PreferredVerificationType string `json:"preferred_verification_type,omitempty"`
	// AUBank is the Australian bank account information. Optional.
	// AUBank 澳大利亚银行账户信息。可选。
	AUBank *AUBank `json:"au_bank,omitempty"`
	// AUPayID is the Australian PayID account information. Optional.
	// AUPayID 澳大利亚PayID账户信息。可选。
	AUPayID *AUPayID `json:"au_payid,omitempty"`
	// CABank is the Canadian bank account information. Optional.
	// CABank 加拿大银行账户信息。可选。
	CABank *CABank `json:"ca_bank,omitempty"`
	// EUBank is the European bank account information. Optional.
	// EUBank 欧洲银行账户信息。可选。
	EUBank *EUBank `json:"eu_bank,omitempty"`
	// GBBank is the UK bank account information. Optional.
	// GBBank 英国银行账户信息。可选。
	GBBank *GBBank `json:"gb_bank,omitempty"`
	// HKBank is the Hong Kong bank account information. Optional.
	// HKBank 香港银行账户信息。可选。
	HKBank *HKBank `json:"hk_bank,omitempty"`
	// NZBank is the New Zealand bank account information. Optional.
	// NZBank 新西兰银行账户信息。可选。
	NZBank *NZBank `json:"nz_bank,omitempty"`
	// SGBank is the Singapore bank account information. Optional.
	// SGBank 新加坡银行账户信息。可选。
	SGBank *SGBank `json:"sg_bank,omitempty"`
	// USBank is the US bank account information. Optional.
	// USBank 美国银行账户信息。可选。
	USBank *USBank `json:"us_bank,omitempty"`
	// Mandate contains mandate information. Optional.
	// Mandate 授权信息。可选。
	Mandate *Mandate `json:"mandate,omitempty"`
	// Plaid contains Plaid connection information. Optional.
	// Plaid Plaid连接信息。可选。
	Plaid *PlaidInfo `json:"plaid,omitempty"`
	// TrueLayer contains TrueLayer connection information. Optional.
	// TrueLayer TrueLayer连接信息。可选。
	TrueLayer *TrueLayerInfo `json:"truelayer,omitempty"`
	// EGiro contains EGiro connection information. Optional.
	// EGiro EGiro连接信息。可选。
	EGiro *EGiroInfo `json:"egiro,omitempty"`
}

// CreateLinkedAccount creates a linked account.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/linked_accounts/create.md
// CreateLinkedAccount 创建关联账户。
func (s *Service) CreateLinkedAccount(ctx context.Context, req *CreateLinkedAccountRequest, opts ...sdk.RequestOption) (*LinkedAccount, error) {
	var resp LinkedAccount
	err := s.doer.Do(ctx, "POST", "/api/v1/linked_accounts/create", req, &resp, opts...)
	return &resp, err
}

// GetLinkedAccount retrieves a linked account by ID.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/linked_accounts/retrieve.md
// GetLinkedAccount 根据 ID 获取关联账户。
func (s *Service) GetLinkedAccount(ctx context.Context, id string, opts ...sdk.RequestOption) (*LinkedAccount, error) {
	var resp LinkedAccount
	err := s.doer.Do(ctx, "GET", "/api/v1/linked_accounts/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListLinkedAccounts lists linked accounts.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/linked_accounts/list.md
// ListLinkedAccounts 列出关联账户。
func (s *Service) ListLinkedAccounts(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[LinkedAccount], error) {
	var resp sdk.ListResult[LinkedAccount]
	err := s.doer.Do(ctx, "GET", "/api/v1/linked_accounts", nil, &resp, opts...)
	return &resp, err
}
