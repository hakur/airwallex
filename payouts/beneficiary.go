package payouts

import (
	"context"
	"net/url"

	"github.com/hakur/airwallex/sdk"
)

// BeneficiaryType represents a beneficiary type.
// BeneficiaryType 收款人类型。
type BeneficiaryType = string

const (
	// BeneficiaryTypeBankAccount indicates a bank account beneficiary.
	// BeneficiaryTypeBankAccount 银行账户。
	BeneficiaryTypeBankAccount BeneficiaryType = "BANK_ACCOUNT"
	// BeneficiaryTypeDigitalWallet indicates a digital wallet beneficiary.
	// BeneficiaryTypeDigitalWallet 数字钱包。
	BeneficiaryTypeDigitalWallet BeneficiaryType = "DIGITAL_WALLET"
	// BeneficiaryTypeCard indicates a card beneficiary.
	// BeneficiaryTypeCard 卡片。
	BeneficiaryTypeCard BeneficiaryType = "CARD"
)

// BankDetails 银行详情。
type BankDetails struct {
	// AccountCurrency 账户货币。可选。
	AccountCurrency sdk.Currency `json:"account_currency,omitempty"`
	// AccountName 账户名称。可选。
	AccountName string `json:"account_name,omitempty"`
	// AccountNameAlias 账户名称别名。可选。
	AccountNameAlias string `json:"account_name_alias,omitempty"`
	// AccountNumber 账户号码。可选。
	AccountNumber string `json:"account_number,omitempty"`
	// AccountRoutingType1 账户路由类型1。可选。
	AccountRoutingType1 string `json:"account_routing_type1,omitempty"`
	// AccountRoutingValue1 账户路由值1。可选。
	AccountRoutingValue1 string `json:"account_routing_value1,omitempty"`
	// AccountRoutingType2 账户路由类型2。可选。
	AccountRoutingType2 string `json:"account_routing_type2,omitempty"`
	// AccountRoutingValue2 账户路由值2。可选。
	AccountRoutingValue2 string `json:"account_routing_value2,omitempty"`
	// BankAccountCategory 银行账户类别。可选。
	BankAccountCategory string `json:"bank_account_category,omitempty"`
	// BankBranch 银行支行。可选。
	BankBranch string `json:"bank_branch,omitempty"`
	// BankCountryCode 银行国家代码。可选。
	BankCountryCode sdk.CountryCode `json:"bank_country_code,omitempty"`
	// BankName 银行名称。可选。
	BankName string `json:"bank_name,omitempty"`
	// BankState 银行所在州/省。可选。
	BankState string `json:"bank_state,omitempty"`
	// BankStreetAddress 银行街道地址。可选。
	BankStreetAddress string `json:"bank_street_address,omitempty"`
	// BindingMobileNumber 绑定手机号码。可选。
	BindingMobileNumber string `json:"binding_mobile_number,omitempty"`
	// BranchCode 支行代码。可选。
	BranchCode string `json:"branch_code,omitempty"`
	// City 城市。可选。
	City string `json:"city,omitempty"`
	// CountryCode 国家代码。可选。
	CountryCode sdk.CountryCode `json:"country_code,omitempty"`
	// Fingerprint 银行账户指纹。可选。
	Fingerprint string `json:"fingerprint,omitempty"`
	// IBAN 国际银行账号。可选。
	IBAN string `json:"iban,omitempty"`
	// IntermediaryBankName intermediary 银行名称。可选。
	IntermediaryBankName string `json:"intermediary_bank_name,omitempty"`
	// IntermediaryBankSwiftCode intermediary 银行 SWIFT 代码。可选。
	IntermediaryBankSwiftCode string `json:"intermediary_bank_swift_code,omitempty"`
	// LocalClearingSystem 本地清算系统。可选。
	LocalClearingSystem string `json:"local_clearing_system,omitempty"`
	// SwiftCode SWIFT代码。可选。
	SwiftCode string `json:"swift_code,omitempty"`
}

// DigitalWalletDetails represents digital wallet details.
// DigitalWalletDetails 数字钱包详情。
type DigitalWalletDetails struct {
	// AccountName 账户持有人姓名。必填。
	AccountName string `json:"account_name"`
	// CountryCode 数字钱包国家代码。可选。
	CountryCode sdk.CountryCode `json:"country_code,omitempty"`
	// IDType 账户标识符类型。必填。
	IDType string `json:"id_type"`
	// IDValue 账户标识符值。必填。
	IDValue string `json:"id_value"`
	// Provider 数字钱包服务提供商。必填。
	Provider string `json:"provider"`
}

// CardDetails represents card details.
// CardDetails 卡片详情。
type CardDetails struct {
	// Brand 卡品牌。可选。
	Brand string `json:"brand,omitempty"`
	// CardholderName 持卡人姓名。必填。
	CardholderName string `json:"cardholder_name"`
	// Currency 卡货币。可选。
	Currency sdk.Currency `json:"currency,omitempty"`
	// ExpiryMonth 到期月份。可选。
	ExpiryMonth string `json:"expiry_month,omitempty"`
	// ExpiryYear 到期年份。可选。
	ExpiryYear string `json:"expiry_year,omitempty"`
	// Number 卡号。可选。
	Number string `json:"number,omitempty"`
	// Token 令牌化卡号。必填。
	Token string `json:"token"`
}

// WalletDetails 钱包详情（保留用于向后兼容）。
type WalletDetails struct {
	// WalletType 钱包类型。必填。
	WalletType string `json:"wallet_type"`
	// AccountID 账户唯一标识符。必填。
	AccountID string `json:"account_id"`
}

// BeneficiaryAddress represents a beneficiary address.
// BeneficiaryAddress 收款人地址。
type BeneficiaryAddress struct {
	// CountryCode 国家代码。必填。
	CountryCode sdk.CountryCode `json:"country_code"`
	// State 州/省。可选。
	State string `json:"state,omitempty"`
	// City 城市。可选。
	City string `json:"city,omitempty"`
	// StreetAddress 街道地址。可选。
	StreetAddress string `json:"street_address,omitempty"`
	// Postcode 邮政编码。可选。
	Postcode string `json:"postcode,omitempty"`
}

// BeneficiaryDetails represents beneficiary details.
// BeneficiaryDetails 表示收款人详情。
type BeneficiaryDetails struct {
	// Type 收款人类型。必填。
	Type BeneficiaryType `json:"type"`
	// EntityType 实体类型。可选。
	EntityType string `json:"entity_type,omitempty"`
	// CompanyName 公司名称。可选。
	CompanyName string `json:"company_name,omitempty"`
	// FirstName 名。可选。
	FirstName string `json:"first_name,omitempty"`
	// LastName 姓。可选。
	LastName string `json:"last_name,omitempty"`
	// DateOfBirth 出生日期。可选。
	DateOfBirth string `json:"date_of_birth,omitempty"`
	// Name 姓名。可选。
	Name string `json:"name,omitempty"`
	// Email 电子邮箱。可选。
	Email string `json:"email,omitempty"`
	// Address 收款人地址。可选。
	Address *BeneficiaryAddress `json:"address,omitempty"`
	// BankDetails 银行详情。可选。
	BankDetails *BankDetails `json:"bank_details,omitempty"`
	// DigitalWallet 数字钱包详情。可选。
	DigitalWallet *DigitalWalletDetails `json:"digital_wallet,omitempty"`
	// WalletDetails 钱包详情（保留用于向后兼容）。可选。
	WalletDetails *WalletDetails `json:"wallet_details,omitempty"`
	// Card 卡片详情。可选。
	Card *CardDetails `json:"card,omitempty"`
	// AdditionalInfo 附加信息。可选。
	AdditionalInfo map[string]any `json:"additional_info,omitempty"`
}

// Beneficiary represents a beneficiary.
// Beneficiary 表示收款人信息。
type Beneficiary struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// PayerEntityType 付款人实体类型。可选。
	PayerEntityType string `json:"payer_entity_type,omitempty"`
	// Nickname 收款人昵称。可选。
	Nickname string `json:"nickname,omitempty"`
	// SCAExemptible 是否可豁免SCA。可选。
	SCAExemptible bool `json:"sca_exemptible,omitempty"`
	// TransferMethods 转账方式列表。可选。
	TransferMethods []string `json:"transfer_methods,omitempty"`
	// Beneficiary 收款人详情。可选。
	Beneficiary *BeneficiaryDetails `json:"beneficiary,omitempty"`
	// Type 收款人类型。可选。
	Type BeneficiaryType `json:"type,omitempty"`
	// Status 状态。可选。
	Status string `json:"status,omitempty"`
}

// BeneficiaryInput is the beneficiary input for create/update requests.
// BeneficiaryInput 用于创建/更新请求的收款人输入结构。
type BeneficiaryInput struct {
	// Type 收款人类型。必填。
	Type BeneficiaryType `json:"type"`
	// EntityType 实体类型。必填。
	EntityType string `json:"entity_type"`
	// CompanyName 公司名称。可选。
	CompanyName string `json:"company_name,omitempty"`
	// FirstName 名。可选。
	FirstName string `json:"first_name,omitempty"`
	// LastName 姓。可选。
	LastName string `json:"last_name,omitempty"`
	// DateOfBirth 出生日期。可选。
	DateOfBirth string `json:"date_of_birth,omitempty"`
	// Name 姓名。可选。
	Name string `json:"name,omitempty"`
	// Email 电子邮箱。可选。
	Email string `json:"email,omitempty"`
	// Address 收款人地址。可选。
	Address *BeneficiaryAddress `json:"address,omitempty"`
	// BankDetails 银行详情。可选。
	BankDetails *BankDetails `json:"bank_details,omitempty"`
	// DigitalWallet 数字钱包详情。可选。
	DigitalWallet *DigitalWalletDetails `json:"digital_wallet,omitempty"`
	// WalletDetails 钱包详情（保留用于向后兼容）。可选。
	WalletDetails *WalletDetails `json:"wallet_details,omitempty"`
	// Card 卡片详情。可选。
	Card *CardDetails `json:"card,omitempty"`
	// AdditionalInfo 附加信息。可选。
	AdditionalInfo map[string]any `json:"additional_info,omitempty"`
}

// CreateBeneficiaryRequest is the request to create a beneficiary.
// CreateBeneficiaryRequest 创建收款人请求。
type CreateBeneficiaryRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Beneficiary 收款人信息。必填。
	Beneficiary *BeneficiaryInput `json:"beneficiary"`
	// Nickname 收款人昵称。可选。
	Nickname string `json:"nickname,omitempty"`
	// PayerEntityType 付款人实体类型。可选。
	PayerEntityType string `json:"payer_entity_type,omitempty"`
	// SCAExemptible 是否可豁免SCA。可选。
	SCAExemptible bool `json:"sca_exemptible,omitempty"`
	// TransferMethods 转账方式列表。必填。
	TransferMethods []string `json:"transfer_methods"`
	// TransferReason 转账原因。可选。
	TransferReason string `json:"transfer_reason,omitempty"`
}

// UpdateBeneficiaryRequest is the request to update a beneficiary.
// UpdateBeneficiaryRequest 更新收款人请求。
type UpdateBeneficiaryRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Beneficiary 收款人信息。可选。
	Beneficiary *BeneficiaryInput `json:"beneficiary,omitempty"`
	// Nickname 收款人昵称。可选。
	Nickname string `json:"nickname,omitempty"`
	// PayerEntityType 付款人实体类型。可选。
	PayerEntityType string `json:"payer_entity_type,omitempty"`
	// SCAExemptible 是否可豁免SCA。可选。
	SCAExemptible bool `json:"sca_exemptible,omitempty"`
	// TransferMethods 转账方式列表。可选。
	TransferMethods []string `json:"transfer_methods,omitempty"`
	// TransferReason 转账原因。可选。
	TransferReason string `json:"transfer_reason,omitempty"`
}

// ValidateBeneficiaryRequest is the request to validate a beneficiary.
// ValidateBeneficiaryRequest 验证收款人请求。
type ValidateBeneficiaryRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Beneficiary 收款人信息。必填。
	Beneficiary *BeneficiaryInput `json:"beneficiary"`
	// Nickname 收款人昵称。可选。
	Nickname string `json:"nickname,omitempty"`
	// PayerEntityType 付款人实体类型。可选。
	PayerEntityType string `json:"payer_entity_type,omitempty"`
	// SCAExemptible 是否可豁免SCA。可选。
	SCAExemptible bool `json:"sca_exemptible,omitempty"`
	// TransferMethods 转账方式列表。必填。
	TransferMethods []string `json:"transfer_methods"`
	// TransferReason 转账原因。可选。
	TransferReason string `json:"transfer_reason,omitempty"`
}

// VerifyAccountRequest is the request to verify an account.
// VerifyAccountRequest 验证账户请求。
type VerifyAccountRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Beneficiary 收款人信息。必填。
	Beneficiary *BeneficiaryInput `json:"beneficiary"`
}

// VerifyAccountResponse is the response for account verification.
// VerifyAccountResponse 验证账户响应。
type VerifyAccountResponse struct {
	// Valid 账户是否有效。
	Valid bool `json:"valid,omitempty"`
	// Errors 错误信息列表。
	Errors []map[string]any `json:"errors,omitempty"`
}

// SupportedFinancialInstitution represents a supported financial institution.
// SupportedFinancialInstitution 支持的金融机构。
type SupportedFinancialInstitution struct {
	// Name 机构名称。
	Name string `json:"name,omitempty"`
	// Code 机构代码。
	Code string `json:"code,omitempty"`
	// CountryCode 国家代码。
	CountryCode sdk.CountryCode `json:"country_code,omitempty"`
	// Currency 支持货币。
	Currency sdk.Currency `json:"currency,omitempty"`
}

// GetSupportedFinancialInstitutionsRequest is the request to get supported financial institutions.
// GetSupportedFinancialInstitutionsRequest 获取支持的金融机构请求。
type GetSupportedFinancialInstitutionsRequest struct {
	// CountryCode 国家代码。必填。
	CountryCode sdk.CountryCode `json:"country_code"`
	// Currency 货币。可选。
	Currency sdk.Currency `json:"currency,omitempty"`
	// LocalClearingSystem 本地清算系统。可选。
	LocalClearingSystem string `json:"local_clearing_system,omitempty"`
}

// CreateBeneficiary creates a beneficiary.
// CreateBeneficiary 创建收款人。
// 官方文档: https://www.airwallex.com/docs/api/payouts/beneficiaries/create.md
func (s *Service) CreateBeneficiary(ctx context.Context, req *CreateBeneficiaryRequest, opts ...sdk.RequestOption) (*Beneficiary, error) {
	var resp Beneficiary
	err := s.doer.Do(ctx, "POST", "/api/v1/beneficiaries/create", req, &resp, opts...)
	return &resp, err
}

// GetBeneficiary retrieves a beneficiary by ID.
// GetBeneficiary 根据 ID 获取收款人。
// 官方文档: https://www.airwallex.com/docs/api/payouts/beneficiaries/retrieve.md
func (s *Service) GetBeneficiary(ctx context.Context, id string, opts ...sdk.RequestOption) (*Beneficiary, error) {
	var resp Beneficiary
	err := s.doer.Do(ctx, "GET", "/api/v1/beneficiaries/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateBeneficiary updates a beneficiary.
// UpdateBeneficiary 更新收款人。
// 官方文档: https://www.airwallex.com/docs/api/payouts/beneficiaries/update.md
func (s *Service) UpdateBeneficiary(ctx context.Context, id string, req *UpdateBeneficiaryRequest, opts ...sdk.RequestOption) (*Beneficiary, error) {
	var resp Beneficiary
	err := s.doer.Do(ctx, "POST", "/api/v1/beneficiaries/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ListBeneficiaries lists beneficiaries.
// ListBeneficiaries 列出收款人。
// 官方文档: https://www.airwallex.com/docs/api/payouts/beneficiaries/list.md
func (s *Service) ListBeneficiaries(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[Beneficiary], error) {
	var resp sdk.ListResult[Beneficiary]
	err := s.doer.Do(ctx, "GET", "/api/v1/beneficiaries", nil, &resp, opts...)
	return &resp, err
}

// DeleteBeneficiary deletes a beneficiary.
// DeleteBeneficiary 删除收款人。
// 官方文档: https://www.airwallex.com/docs/api/payouts/beneficiaries/delete.md
func (s *Service) DeleteBeneficiary(ctx context.Context, id string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/beneficiaries/"+id+"/delete", map[string]any{}, nil, opts...)
}

// ValidateBeneficiary validates beneficiary information.
// ValidateBeneficiary 验证收款人信息。
// 官方文档: https://www.airwallex.com/docs/api/payouts/beneficiaries/validate.md
func (s *Service) ValidateBeneficiary(ctx context.Context, req *ValidateBeneficiaryRequest, opts ...sdk.RequestOption) (bool, error) {
	var resp string
	err := s.doer.Do(ctx, "POST", "/api/v1/beneficiaries/validate", req, &resp, opts...)
	if err != nil {
		return false, err
	}
	return resp == "OK", nil
}

// VerifyAccount verifies bank account information.
// VerifyAccount 验证银行账户信息。
// 官方文档: https://www.airwallex.com/docs/api/payouts/beneficiaries/verify_account.md
func (s *Service) VerifyAccount(ctx context.Context, req *VerifyAccountRequest, opts ...sdk.RequestOption) (*VerifyAccountResponse, error) {
	var resp VerifyAccountResponse
	err := s.doer.Do(ctx, "POST", "/api/v1/beneficiaries/verify_account", req, &resp, opts...)
	return &resp, err
}

// GenerateAPISchema generates beneficiary API schema.
// GenerateAPISchema 生成收款人API模式。
// 官方文档: https://www.airwallex.com/docs/api/payouts/beneficiaries/generate_api_schema.md
func (s *Service) GenerateAPISchema(ctx context.Context, opts ...sdk.RequestOption) (map[string]any, error) {
	var resp map[string]any
	err := s.doer.Do(ctx, "POST", "/api/v1/beneficiaries/generate_api_schema", nil, &resp, opts...)
	return resp, err
}

// GenerateFormSchema generates beneficiary form schema.
// GenerateFormSchema 生成收款人表单模式。
// 官方文档: https://www.airwallex.com/docs/api/payouts/beneficiaries/generate_form_schema.md
func (s *Service) GenerateFormSchema(ctx context.Context, opts ...sdk.RequestOption) (map[string]any, error) {
	var resp map[string]any
	err := s.doer.Do(ctx, "POST", "/api/v1/beneficiaries/generate_form_schema", nil, &resp, opts...)
	return resp, err
}

// GetSupportedFinancialInstitutions retrieves supported financial institutions.
// GetSupportedFinancialInstitutions 获取支持的金融机构列表。
// 官方文档: https://www.airwallex.com/docs/api/payouts/beneficiaries/supported_financial_institutions.md
func (s *Service) GetSupportedFinancialInstitutions(ctx context.Context, req *GetSupportedFinancialInstitutionsRequest, opts ...sdk.RequestOption) (*sdk.ListResult[SupportedFinancialInstitution], error) {
	path := "/api/v1/beneficiaries/supported_financial_institutions"
	if req != nil {
		q := url.Values{}
		if req.CountryCode != "" {
			q.Set("country_code", string(req.CountryCode))
		}
		if req.Currency != "" {
			q.Set("currency", string(req.Currency))
		}
		if req.LocalClearingSystem != "" {
			q.Set("local_clearing_system", req.LocalClearingSystem)
		}
		if len(q) > 0 {
			path += "?" + q.Encode()
		}
	}
	var resp sdk.ListResult[SupportedFinancialInstitution]
	err := s.doer.Do(ctx, "GET", path, nil, &resp, opts...)
	return &resp, err
}
