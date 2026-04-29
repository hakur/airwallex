package issuing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// CardholderType represents the cardholder type enum.
// CardholderType 持卡人类型枚举。
type CardholderType = string

const (
	// CardholderTypeIndividual is an individual cardholder.
	// CardholderTypeIndividual 个人持卡人
	CardholderTypeIndividual CardholderType = "INDIVIDUAL"
	// CardholderTypeDelegate is a delegate cardholder.
	// CardholderTypeDelegate 委派持卡人
	CardholderTypeDelegate CardholderType = "DELEGATE"
)

// CardholderAddress represents the cardholder address.
// CardholderAddress 持卡人地址。
type CardholderAddress struct {
	// City is the city. Optional.
	// City 城市。可选。
	City string `json:"city,omitempty"`
	// Country is the country. Optional.
	// Country 国家。可选。
	Country string `json:"country,omitempty"`
	// Postcode is the postal code. Optional.
	// Postcode 邮政编码。可选。
	Postcode string `json:"postcode,omitempty"`
	// State is the state/province. Optional.
	// State 州/省。可选。
	State string `json:"state,omitempty"`
	// Line1 is address line 1. Optional.
	// Line1 地址行1。可选。
	Line1 string `json:"line1,omitempty"`
	// Line2 is address line 2. Optional.
	// Line2 地址行2。可选。
	Line2 string `json:"line2,omitempty"`
}

// CardholderName represents the cardholder name information.
// CardholderName 持卡人姓名信息。
type CardholderName struct {
	// FirstName is the first name. Optional.
	// FirstName 名。可选。
	FirstName string `json:"first_name,omitempty"`
	// MiddleName is the middle name. Optional.
	// MiddleName 中间名。可选。
	MiddleName string `json:"middle_name,omitempty"`
	// LastName is the last name. Optional.
	// LastName 姓。可选。
	LastName string `json:"last_name,omitempty"`
	// Title is the title. Optional.
	// Title 称谓。可选。
	Title string `json:"title,omitempty"`
}

// CardholderIdentification represents the cardholder identification document information.
// CardholderIdentification 持卡人身份证明文件信息。
type CardholderIdentification struct {
	// Country is the issuing country of the document. Optional.
	// Country 证件签发国家。可选。
	Country string `json:"country,omitempty"`
	// DocumentBackFileID is the back of the document file ID. Optional.
	// DocumentBackFileID 证件背面文件ID。可选。
	DocumentBackFileID string `json:"document_back_file_id,omitempty"`
	// DocumentFrontFileID is the front of the document file ID. Optional.
	// DocumentFrontFileID 证件正面文件ID。可选。
	DocumentFrontFileID string `json:"document_front_file_id,omitempty"`
	// ExpiryDate is the document expiry date. Optional.
	// ExpiryDate 证件到期日。可选。
	ExpiryDate string `json:"expiry_date,omitempty"`
	// Gender is the gender. Optional.
	// Gender 性别。可选。
	Gender string `json:"gender,omitempty"`
	// Number is the document number. Optional.
	// Number 证件号码。可选。
	Number string `json:"number,omitempty"`
	// State is the issuing state/province of the document. Optional.
	// State 证件签发州/省。可选。
	State string `json:"state,omitempty"`
	// Type is the document type. Optional.
	// Type 证件类型。可选。
	Type string `json:"type,omitempty"`
}

// CardholderIndividual represents the cardholder individual information.
// CardholderIndividual 持卡人个人信息。
type CardholderIndividual struct {
	// Name contains name information. Optional.
	// Name 姓名信息。可选。
	Name *CardholderName `json:"name,omitempty"`
	// DateOfBirth is the date of birth. Optional.
	// DateOfBirth 出生日期。可选。
	DateOfBirth string `json:"date_of_birth,omitempty"`
	// Address is the address. Optional.
	// Address 地址。可选。
	Address *CardholderAddress `json:"address,omitempty"`
	// PhoneNumber is the phone number. Optional.
	// PhoneNumber 电话号码。可选。
	PhoneNumber string `json:"phone_number,omitempty"`
	// Email is the email address. Optional.
	// Email 电子邮箱。可选。
	Email string `json:"email,omitempty"`
	// Nationality is the nationality. Optional.
	// Nationality 国籍。可选。
	Nationality string `json:"nationality,omitempty"`
	// Identification contains identification document information. Optional.
	// Identification 身份证明文件信息。可选。
	Identification *CardholderIdentification `json:"identification,omitempty"`
	// CardholderAgreementTermsConsentObtained indicates whether consent was obtained for cardholder agreement terms. Optional.
	// CardholderAgreementTermsConsentObtained 是否已获取持卡人协议条款同意。可选。
	CardholderAgreementTermsConsentObtained bool `json:"cardholder_agreement_terms_consent_obtained,omitempty"`
	// PaperlessNotificationConsentObtained indicates whether paperless notification consent was obtained. Optional.
	// PaperlessNotificationConsentObtained 是否已获取无纸化通知同意。可选。
	PaperlessNotificationConsentObtained bool `json:"paperless_notification_consent_obtained,omitempty"`
	// PrivacyPolicyTermsConsentObtained indicates whether privacy policy consent was obtained. Optional.
	// PrivacyPolicyTermsConsentObtained 是否已获取隐私政策条款同意。可选。
	PrivacyPolicyTermsConsentObtained bool `json:"privacy_policy_terms_consent_obtained,omitempty"`
	// TaxIdentificationNumber is the tax identification number. Optional.
	// TaxIdentificationNumber 纳税人识别号。可选。
	TaxIdentificationNumber string `json:"tax_identification_number,omitempty"`
	// ExpressConsentObtained indicates whether express consent was obtained. Optional.
	// ExpressConsentObtained 是否已获取持卡人明确同意。可选。
	ExpressConsentObtained string `json:"express_consent_obtained,omitempty"`
}

// Cardholder represents cardholder information.
// Cardholder 表示持卡人信息。
type Cardholder struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"cardholder_id"`
	// Email is the email address. Optional.
	// Email 电子邮箱。可选。
	Email string `json:"email,omitempty"`
	// Status is the status. Required.
	// Status 状态。必填。
	Status string `json:"status"`
	// Type is the type. Optional.
	// Type 类型。可选。
	Type CardholderType `json:"type,omitempty"`
	// Individual contains individual information. Optional.
	// Individual 个人信息。可选。
	Individual *CardholderIndividual `json:"individual,omitempty"`
	// MobileNumber is the mobile phone number. Optional.
	// MobileNumber 手机号码。可选。
	MobileNumber string `json:"mobile_number,omitempty"`
	// PostalAddress is the postal address. Optional.
	// PostalAddress 邮寄地址。可选。
	PostalAddress *CardholderAddress `json:"postal_address,omitempty"`
	// Employers is the list of employer information. Optional.
	// Employers 雇主信息列表。可选。
	Employers []map[string]any `json:"employers,omitempty"`
}

// CreateCardholderRequest represents the request to create a cardholder.
// CreateCardholderRequest 创建持卡人请求。
type CreateCardholderRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Type is the type. Required.
	// Type 类型。必填。
	Type CardholderType `json:"type"`
	// Email is the email address. Optional.
	// Email 电子邮箱。可选。
	Email string `json:"email,omitempty"`
	// Individual contains individual information. Optional.
	// Individual 个人信息。可选。
	Individual *CardholderIndividual `json:"individual,omitempty"`
	// MobileNumber is the mobile phone number. Optional.
	// MobileNumber 手机号码。可选。
	MobileNumber string `json:"mobile_number,omitempty"`
	// PostalAddress is the postal address. Optional.
	// PostalAddress 邮寄地址。可选。
	PostalAddress *CardholderAddress `json:"postal_address,omitempty"`
}

// UpdateCardholderRequest represents the request to update a cardholder.
// UpdateCardholderRequest 更新持卡人请求。
type UpdateCardholderRequest struct {
	// Type is the type. Optional.
	// Type 类型。可选。
	Type CardholderType `json:"type,omitempty"`
	// Email is the email address. Optional.
	// Email 电子邮箱。可选。
	Email string `json:"email,omitempty"`
	// Individual contains individual information. Optional.
	// Individual 个人信息。可选。
	Individual *CardholderIndividual `json:"individual,omitempty"`
	// MobileNumber is the mobile phone number. Optional.
	// MobileNumber 手机号码。可选。
	MobileNumber string `json:"mobile_number,omitempty"`
	// PostalAddress is the postal address. Optional.
	// PostalAddress 邮寄地址。可选。
	PostalAddress *CardholderAddress `json:"postal_address,omitempty"`
}

// CreateCardholder creates a cardholder.
// 官方文档: https://www.airwallex.com/docs/api/issuing/cardholders/create.md
// CreateCardholder 创建持卡人。
func (s *Service) CreateCardholder(ctx context.Context, req *CreateCardholderRequest, opts ...sdk.RequestOption) (*Cardholder, error) {
	var resp Cardholder
	err := s.doer.Do(ctx, "POST", "/api/v1/issuing/cardholders/create", req, &resp, opts...)
	return &resp, err
}

// GetCardholder retrieves a cardholder by ID.
// 官方文档: https://www.airwallex.com/docs/api/issuing/cardholders/retrieve.md
// GetCardholder 根据 ID 获取持卡人。
func (s *Service) GetCardholder(ctx context.Context, id string, opts ...sdk.RequestOption) (*Cardholder, error) {
	var resp Cardholder
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/cardholders/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateCardholder updates a cardholder.
// 官方文档: https://www.airwallex.com/docs/api/issuing/cardholders/update.md
// UpdateCardholder 更新持卡人。
func (s *Service) UpdateCardholder(ctx context.Context, id string, req *UpdateCardholderRequest, opts ...sdk.RequestOption) (*Cardholder, error) {
	var resp Cardholder
	err := s.doer.Do(ctx, "POST", "/api/v1/issuing/cardholders/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ListCardholders lists cardholders.
// 官方文档: https://www.airwallex.com/docs/api/issuing/cardholders/list.md
// ListCardholders 列出持卡人。
func (s *Service) ListCardholders(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[Cardholder], error) {
	var resp sdk.ListResult[Cardholder]
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/cardholders", nil, &resp, opts...)
	return &resp, err
}

// DeleteCardholderResponse represents the response to a delete cardholder request.
// DeleteCardholderResponse 删除持卡人响应。
type DeleteCardholderResponse struct {
	// CardholderID is the unique cardholder identifier. Optional.
	// CardholderID 持卡人唯一标识符。可选。
	CardholderID string `json:"cardholder_id,omitempty"`
	// Deleted indicates whether the deletion was successful. Optional.
	// Deleted 是否删除成功。可选。
	Deleted bool `json:"deleted,omitempty"`
}

// DeleteCardholder deletes a cardholder.
// 官方文档: https://www.airwallex.com/docs/api/issuing/cardholders/delete.md
// DeleteCardholder 删除持卡人。
func (s *Service) DeleteCardholder(ctx context.Context, id string, opts ...sdk.RequestOption) (*DeleteCardholderResponse, error) {
	var resp DeleteCardholderResponse
	err := s.doer.Do(ctx, "POST", "/api/v1/issuing/cardholders/"+id+"/delete", nil, &resp, opts...)
	return &resp, err
}
