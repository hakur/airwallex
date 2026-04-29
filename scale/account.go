package scale

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// AccountAddress represents account address information.
// AccountAddress 账户地址信息。
type AccountAddress struct {
	// City 城市。
	City string `json:"city,omitempty"`
	// CountryCode 国家代码（2位ISO 3166-2代码）。
	CountryCode string `json:"country_code,omitempty"`
	// Postcode 邮政编码。
	Postcode string `json:"postcode,omitempty"`
	// State 州/省。
	State string `json:"state,omitempty"`
	// Street 街道。
	Street string `json:"street,omitempty"`
}

// AccountDetails represents account details for a connected account.
// AccountDetails 账户详情。
type AccountDetails struct {
	// BusinessName 企业名称。
	BusinessName string `json:"business_name,omitempty"`
	// RegistrationNumber 注册号码。
	RegistrationNumber string `json:"registration_number,omitempty"`
	// TaxID 税号。
	TaxID string `json:"tax_id,omitempty"`
	// IndustryCode 行业代码。
	IndustryCode string `json:"industry_code,omitempty"`
	// Website 网站。
	Website string `json:"website,omitempty"`
	// BusinessType 企业类型。
	BusinessType string `json:"business_type,omitempty"`
	// DateOfIncorporation 成立日期。
	DateOfIncorporation string `json:"date_of_incorporation,omitempty"`
	// Address 地址。
	Address *AccountAddress `json:"address,omitempty"`
	// PhoneNumber 电话号码。
	PhoneNumber string `json:"phone_number,omitempty"`
	// Email 邮箱。
	Email string `json:"email,omitempty"`
}

// BusinessDetails represents business details for a connected account.
// BusinessDetails 企业详情。
type BusinessDetails struct {
	// LegalName 法律名称。
	LegalName string `json:"legal_name,omitempty"`
	// TradingName 经营名称。
	TradingName string `json:"trading_name,omitempty"`
	// RegistrationNumber 注册号码。
	RegistrationNumber string `json:"registration_number,omitempty"`
	// TaxID 税号。
	TaxID string `json:"tax_id,omitempty"`
	// IndustryCode 行业代码。
	IndustryCode string `json:"industry_code,omitempty"`
	// Website 网站。
	Website string `json:"website,omitempty"`
	// BusinessType 企业类型。
	BusinessType string `json:"business_type,omitempty"`
	// DateOfIncorporation 成立日期。
	DateOfIncorporation string `json:"date_of_incorporation,omitempty"`
	// Address 地址。
	Address *AccountAddress `json:"address,omitempty"`
	// PhoneNumber 电话号码。
	PhoneNumber string `json:"phone_number,omitempty"`
	// Email 邮箱。
	Email string `json:"email,omitempty"`
}

// IndividualDetails represents individual details for a connected account.
// IndividualDetails 个人详情。
type IndividualDetails struct {
	// FirstName 名字。
	FirstName string `json:"first_name,omitempty"`
	// LastName 姓氏。
	LastName string `json:"last_name,omitempty"`
	// DateOfBirth 出生日期。
	DateOfBirth string `json:"date_of_birth,omitempty"`
	// Address 地址。
	Address *AccountAddress `json:"address,omitempty"`
	// PhoneNumber 电话号码。
	PhoneNumber string `json:"phone_number,omitempty"`
	// Email 邮箱。
	Email string `json:"email,omitempty"`
	// Nationality 国籍。
	Nationality string `json:"nationality,omitempty"`
}

// PrimaryContact represents a primary contact for a connected account.
// PrimaryContact 主要联系人。
type PrimaryContact struct {
	// FirstName 名字。
	FirstName string `json:"first_name,omitempty"`
	// LastName 姓氏。
	LastName string `json:"last_name,omitempty"`
	// Email 邮箱。
	Email string `json:"email,omitempty"`
	// PhoneNumber 电话号码。
	PhoneNumber string `json:"phone_number,omitempty"`
	// JobTitle 职位。
	JobTitle string `json:"job_title,omitempty"`
}

// Account represents a connected account.
// Account 表示连接账户信息。
type Account struct {
	// ID 账户唯一标识符。
	ID string `json:"id,omitempty"`
	// Email 账户邮箱。
	Email string `json:"email,omitempty"`
	// Status 账户状态。
	Status string `json:"status,omitempty"`
	// LegalEntityID 法律实体ID。
	LegalEntityID string `json:"legal_entity_id,omitempty"`
	// LegalEntityType 法律实体类型（BUSINESS/INDIVIDUAL）。
	LegalEntityType string `json:"legal_entity_type,omitempty"`
	// AccountType 账户类型。
	AccountType string `json:"account_type,omitempty"`
	// AccountDetails 账户详情。
	AccountDetails *AccountDetails `json:"account_details,omitempty"`
	// BusinessDetails 企业详情。
	BusinessDetails *BusinessDetails `json:"business_details,omitempty"`
	// IndividualDetails 个人详情。
	IndividualDetails *IndividualDetails `json:"individual_details,omitempty"`
	// PrimaryContact 主要联系人。
	PrimaryContact *PrimaryContact `json:"primary_contact,omitempty"`
	// CreatedAt 创建时间。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt 更新时间。
	UpdatedAt string `json:"updated_at,omitempty"`
	// KYCStatus KYC状态。
	KYCStatus string `json:"kyc_status,omitempty"`
	// KYCDocuments KYC文档列表。
	KYCDocuments []map[string]any `json:"kyc_documents,omitempty"`
	// BankAccounts 银行账户列表。
	BankAccounts []map[string]any `json:"bank_accounts,omitempty"`
	// Settings 账户设置。
	Settings map[string]any `json:"settings,omitempty"`
	// Metadata 元数据。
	Metadata map[string]any `json:"metadata,omitempty"`
	// BusinessPersonDetails 企业相关人员详情。
	BusinessPersonDetails []map[string]any `json:"business_person_details,omitempty"`
	// TrusteeDetails 受托人详情。
	TrusteeDetails map[string]any `json:"trustee_details,omitempty"`
	// StoreDetails 店铺详情。
	StoreDetails map[string]any `json:"store_details,omitempty"`
}

// CreateAccountRequest represents a request to create a connected account.
// CreateAccountRequest 创建连接账户请求。
type CreateAccountRequest struct {
	// RequestID 请求唯一标识符。
	RequestID string `json:"request_id,omitempty"`
	// Email 账户邮箱。必填。
	Email string `json:"email"`
	// AccountType 账户类型。
	AccountType string `json:"account_type,omitempty"`
	// AccountDetails 账户详情。
	AccountDetails *AccountDetails `json:"account_details,omitempty"`
	// BusinessDetails 企业详情。
	BusinessDetails *BusinessDetails `json:"business_details,omitempty"`
	// IndividualDetails 个人详情。
	IndividualDetails *IndividualDetails `json:"individual_details,omitempty"`
	// PrimaryContact 主要联系人。
	PrimaryContact *PrimaryContact `json:"primary_contact,omitempty"`
	// Metadata 元数据。
	Metadata map[string]any `json:"metadata,omitempty"`
	// StoreDetails 店铺详情。
	StoreDetails map[string]any `json:"store_details,omitempty"`
}

// UpdateAccountRequest represents a request to update a connected account.
// UpdateAccountRequest 更新连接账户请求。
type UpdateAccountRequest struct {
	// AccountDetails 账户详情。
	AccountDetails *AccountDetails `json:"account_details,omitempty"`
	// BusinessDetails 企业详情。
	BusinessDetails *BusinessDetails `json:"business_details,omitempty"`
	// IndividualDetails 个人详情。
	IndividualDetails *IndividualDetails `json:"individual_details,omitempty"`
	// PrimaryContact 主要联系人。
	PrimaryContact *PrimaryContact `json:"primary_contact,omitempty"`
	// BusinessPersonDetails 企业相关人员详情。
	BusinessPersonDetails []map[string]any `json:"business_person_details,omitempty"`
	// StoreDetails 店铺详情。
	StoreDetails map[string]any `json:"store_details,omitempty"`
}

// CreateAccount creates a new connected account.
// CreateAccount 创建连接账户。
// 官方文档: https://www.airwallex.com/docs/api/scale/accounts/create.md
func (s *Service) CreateAccount(ctx context.Context, req *CreateAccountRequest, opts ...sdk.RequestOption) (*Account, error) {
	var resp Account
	err := s.doer.Do(ctx, "POST", "/api/v1/accounts/create", req, &resp, opts...)
	return &resp, err
}

// GetAccount retrieves a connected account by ID.
// GetAccount 根据ID获取连接账户。
// 官方文档: https://www.airwallex.com/docs/api/scale/accounts/retrieve.md
func (s *Service) GetAccount(ctx context.Context, id string, opts ...sdk.RequestOption) (*Account, error) {
	var resp Account
	err := s.doer.Do(ctx, "GET", "/api/v1/accounts/"+id, nil, &resp, opts...)
	return &resp, err
}

// GetCurrentAccount retrieves the current account details.
// GetCurrentAccount 获取当前账户详情。
// 官方文档: https://www.airwallex.com/docs/api/scale/accounts/detail.md
func (s *Service) GetCurrentAccount(ctx context.Context, opts ...sdk.RequestOption) (*Account, error) {
	var resp Account
	err := s.doer.Do(ctx, "GET", "/api/v1/account", nil, &resp, opts...)
	return &resp, err
}

// ListAccounts lists connected accounts.
// ListAccounts 列出连接账户。
// 官方文档: https://www.airwallex.com/docs/api/scale/accounts/list.md
func (s *Service) ListAccounts(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[Account], error) {
	var resp sdk.ListResult[Account]
	err := s.doer.Do(ctx, "GET", "/api/v1/accounts", nil, &resp, opts...)
	return &resp, err
}

// UpdateAccount updates a connected account.
// UpdateAccount 更新连接账户。
// 官方文档: https://www.airwallex.com/docs/api/scale/accounts/update.md
func (s *Service) UpdateAccount(ctx context.Context, id string, req *UpdateAccountRequest, opts ...sdk.RequestOption) (*Account, error) {
	var resp Account
	err := s.doer.Do(ctx, "POST", "/api/v1/accounts/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ReactivateAccount reactivates a suspended connected account.
// ReactivateAccount 重新激活连接账户。
// 官方文档: https://www.airwallex.com/docs/api/scale/accounts/reactivate.md
func (s *Service) ReactivateAccount(ctx context.Context, id string, opts ...sdk.RequestOption) (*Account, error) {
	var resp Account
	err := s.doer.Do(ctx, "POST", "/api/v1/accounts/"+id+"/reactivate", nil, &resp, opts...)
	return &resp, err
}

// SubmitAccount submits a connected account for activation.
// SubmitAccount 提交连接账户以激活。
// 官方文档: https://www.airwallex.com/docs/api/scale/accounts/submit.md
func (s *Service) SubmitAccount(ctx context.Context, id string, opts ...sdk.RequestOption) (*Account, error) {
	var resp Account
	err := s.doer.Do(ctx, "POST", "/api/v1/accounts/"+id+"/submit", nil, &resp, opts...)
	return &resp, err
}

// SuspendAccount suspends a connected account.
// SuspendAccount 暂停连接账户。
// 官方文档: https://www.airwallex.com/docs/api/scale/accounts/suspend.md
func (s *Service) SuspendAccount(ctx context.Context, id string, opts ...sdk.RequestOption) (*Account, error) {
	var resp Account
	err := s.doer.Do(ctx, "POST", "/api/v1/accounts/"+id+"/suspend", nil, &resp, opts...)
	return &resp, err
}

// AgreeToTermsAndConditions agrees to terms and conditions for a connected account.
// AgreeToTermsAndConditions 同意连接账户的条款和条件。
// 官方文档: https://www.airwallex.com/docs/api/scale/accounts/agree_terms_and_conditions.md
func (s *Service) AgreeToTermsAndConditions(ctx context.Context, id string, opts ...sdk.RequestOption) (*Account, error) {
	var resp Account
	err := s.doer.Do(ctx, "POST", "/api/v1/accounts/"+id+"/terms_and_conditions/agree", nil, &resp, opts...)
	return &resp, err
}
