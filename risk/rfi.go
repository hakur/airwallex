package risk

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// RFIStatus represents the RFI status.
// RFIStatus RFI状态。
type RFIStatus = string

const (
	// RFIStatusActionRequired indicates Airwallex requires your response to the RFI request.
	// RFIStatusActionRequired Airwallex要求您响应RFI请求。
	RFIStatusActionRequired RFIStatus = "ACTION_REQUIRED"
	// RFIStatusAnswered indicates you have responded and Airwallex is reviewing.
	// RFIStatusAnswered 您已响应RFI请求，正在由Airwallex审核。
	RFIStatusAnswered RFIStatus = "ANSWERED"
	// RFIStatusClosed indicates the review is complete with no further information needed.
	// RFIStatusClosed 审核完成，无需进一步信息。
	RFIStatusClosed RFIStatus = "CLOSED"
)

// RFIType represents the RFI type.
// RFIType RFI类型。
type RFIType = string

const (
	// RFITypeKYC is the KYC review type.
	// RFITypeKYC KYC审核。
	RFITypeKYC RFIType = "KYC"
	// RFITypeKYCOngoing is the ongoing KYC review type.
	// RFITypeKYCOngoing 持续KYC审核。
	RFITypeKYCOngoing RFIType = "KYC_ONGOING"
	// RFITypeCardholder is the cardholder-related type.
	// RFITypeCardholder 持卡人相关。
	RFITypeCardholder RFIType = "CARDHOLDER"
	// RFITypeTransaction is the transaction-related type.
	// RFITypeTransaction 交易相关。
	RFITypeTransaction RFIType = "TRANSACTION"
	// RFITypePaymentEnablement is the payment enablement type.
	// RFITypePaymentEnablement 支付功能开通。
	RFITypePaymentEnablement RFIType = "PAYMENT_ENABLEMENT"
	// RFITypeMerchantRisk is the merchant risk type.
	// RFITypeMerchantRisk 商户风险。
	RFITypeMerchantRisk RFIType = "MERCHANT_RISK"
)

// RFISubType represents the RFI sub-type.
// RFISubType RFI子类型。
type RFISubType = string

const (
	// RFISubTypeTransactionMonitoring is the internal transaction monitoring sub-type.
	// RFISubTypeTransactionMonitoring 内部交易监控。
	RFISubTypeTransactionMonitoring RFISubType = "TRANSACTION_MONITORING"
	// RFISubTypeTransactionScreening is the internal transaction screening sub-type.
	// RFISubTypeTransactionScreening 内部交易筛查。
	RFISubTypeTransactionScreening RFISubType = "TRANSACTION_SCREENING"
	// RFISubTypeBankingPartnerInquiry is the external banking partner inquiry sub-type.
	// RFISubTypeBankingPartnerInquiry 外部银行合作伙伴查询。
	RFISubTypeBankingPartnerInquiry RFISubType = "BANKING_PARTNER_INQUIRY"
)

// RFIQuestionAnswerType represents the question answer type.
// RFIQuestionAnswerType 问题答案类型。
type RFIQuestionAnswerType = string

const (
	// RFIQuestionAnswerTypeText is a text reply.
	// RFIQuestionAnswerTypeText 文本回复。
	RFIQuestionAnswerTypeText RFIQuestionAnswerType = "TEXT"
	// RFIQuestionAnswerTypeAttachment is an attachment reply.
	// RFIQuestionAnswerTypeAttachment 附件回复。
	RFIQuestionAnswerTypeAttachment RFIQuestionAnswerType = "ATTACHMENT"
	// RFIQuestionAnswerTypeIdentityDocument is an identity document reply.
	// RFIQuestionAnswerTypeIdentityDocument 身份证件回复。
	RFIQuestionAnswerTypeIdentityDocument RFIQuestionAnswerType = "IDENTITY_DOCUMENT"
	// RFIQuestionAnswerTypeConfirmation is a confirmation reply.
	// RFIQuestionAnswerTypeConfirmation 确认回复。
	RFIQuestionAnswerTypeConfirmation RFIQuestionAnswerType = "CONFIRMATION"
	// RFIQuestionAnswerTypeLiveness is a liveness verification reply.
	// RFIQuestionAnswerTypeLiveness 活体认证回复。
	RFIQuestionAnswerTypeLiveness RFIQuestionAnswerType = "LIVENESS"
	// RFIQuestionAnswerTypeAddress is an address reply.
	// RFIQuestionAnswerTypeAddress 地址回复。
	RFIQuestionAnswerTypeAddress RFIQuestionAnswerType = "ADDRESS"
)

// RFIQuestionKey represents the question key.
// RFIQuestionKey 问题关键字。
type RFIQuestionKey = string

const (
	// RFIQuestionKeyProofOfAddress is the proof of address key.
	// RFIQuestionKeyProofOfAddress 地址证明。
	RFIQuestionKeyProofOfAddress RFIQuestionKey = "PROOF_OF_ADDRESS"
	// RFIQuestionKeyIDCopy is the ID copy key.
	// RFIQuestionKeyIDCopy 身份证件复印件。
	RFIQuestionKeyIDCopy RFIQuestionKey = "ID_COPY"
)

// RFILivenessStatus represents the liveness verification status.
// RFILivenessStatus 活体认证状态。
type RFILivenessStatus = string

const (
	// RFILivenessStatusNotStarted indicates the liveness check has not started.
	// RFILivenessStatusNotStarted 未开始。
	RFILivenessStatusNotStarted RFILivenessStatus = "NOT_STARTED"
	// RFILivenessStatusStarted indicates the liveness check has started.
	// RFILivenessStatusStarted 已开始。
	RFILivenessStatusStarted RFILivenessStatus = "STARTED"
	// RFILivenessStatusSubmitted indicates the liveness check has been submitted.
	// RFILivenessStatusSubmitted 已提交。
	RFILivenessStatusSubmitted RFILivenessStatus = "SUBMITTED"
	// RFILivenessStatusSucceeded indicates the liveness check succeeded.
	// RFILivenessStatusSucceeded 成功。
	RFILivenessStatusSucceeded RFILivenessStatus = "SUCCEEDED"
	// RFILivenessStatusFailed indicates the liveness check failed.
	// RFILivenessStatusFailed 失败。
	RFILivenessStatusFailed RFILivenessStatus = "FAILED"
)

// RFISourceType represents the source type.
// RFISourceType 来源类型。
type RFISourceType = string

const (
	// RFISourceTypePerson is a person source.
	// RFISourceTypePerson 人员。
	RFISourceTypePerson RFISourceType = "PERSON"
	// RFISourceTypeBusiness is a business source.
	// RFISourceTypeBusiness 企业。
	RFISourceTypeBusiness RFISourceType = "BUSINESS"
	// RFISourceTypeIndividual is an individual source.
	// RFISourceTypeIndividual 个人。
	RFISourceTypeIndividual RFISourceType = "INDIVIDUAL"
	// RFISourceTypeCardholder is a cardholder source.
	// RFISourceTypeCardholder 持卡人。
	RFISourceTypeCardholder RFISourceType = "CARDHOLDER"
	// RFISourceTypePaymentIntent is a payment intent source.
	// RFISourceTypePaymentIntent 支付意图。
	RFISourceTypePaymentIntent RFISourceType = "PAYMENT_INTENT"
	// RFISourceTypeCharge is a charge source.
	// RFISourceTypeCharge 费用。
	RFISourceTypeCharge RFISourceType = "CHARGE"
	// RFISourceTypeTransfer is a transfer source.
	// RFISourceTypeTransfer 转账。
	RFISourceTypeTransfer RFISourceType = "TRANSFER"
	// RFISourceTypeDirectDebit is a direct debit source.
	// RFISourceTypeDirectDebit 直接扣款。
	RFISourceTypeDirectDebit RFISourceType = "DIRECT_DEBIT"
	// RFISourceTypeDeposit is a deposit source.
	// RFISourceTypeDeposit 存款。
	RFISourceTypeDeposit RFISourceType = "DEPOSIT"
	// RFISourceTypePayment is a payment source.
	// RFISourceTypePayment 支付。
	RFISourceTypePayment RFISourceType = "PAYMENT"
)

// IdentityDocumentType represents the identity document type.
// IdentityDocumentType 身份证件类型。
type IdentityDocumentType = string

const (
	// IdentityDocumentTypeDrivingLicense is a driving license.
	// IdentityDocumentTypeDrivingLicense 驾驶证。
	IdentityDocumentTypeDrivingLicense IdentityDocumentType = "DRIVING_LICENSE"
	// IdentityDocumentTypePassport is a passport.
	// IdentityDocumentTypePassport 护照。
	IdentityDocumentTypePassport IdentityDocumentType = "PASSPORT"
	// IdentityDocumentTypeIdentityCard is an identity card.
	// IdentityDocumentTypeIdentityCard 身份证。
	IdentityDocumentTypeIdentityCard IdentityDocumentType = "IDENTITY_CARD"
	// IdentityDocumentTypeMedicareCard is a medicare card.
	// IdentityDocumentTypeMedicareCard 医保卡。
	IdentityDocumentTypeMedicareCard IdentityDocumentType = "MEDICARE_CARD"
	// IdentityDocumentTypeSSN is a Social Security Number.
	// IdentityDocumentTypeSSN 社会安全号码。
	IdentityDocumentTypeSSN IdentityDocumentType = "SSN"
	// IdentityDocumentTypeITIN is an Individual Taxpayer Identification Number.
	// IdentityDocumentTypeITIN 个人纳税人识别号。
	IdentityDocumentTypeITIN IdentityDocumentType = "ITIN"
)

// LocalizedContent represents localized content.
// LocalizedContent 本地化内容。
type LocalizedContent struct {
	// EN is the English content.
	// EN 英文内容。
	EN string `json:"en,omitempty"`
	// ZH is the Chinese content.
	// ZH 中文内容。
	ZH string `json:"zh,omitempty"`
}

// RFIAddressAnswer represents an address answer.
// RFIAddressAnswer 地址答案。
type RFIAddressAnswer struct {
	// AddressLine1 is the first line of the address (street number and building name).
	// AddressLine1 地址第一行（街道号码和建筑名称）。
	AddressLine1 string `json:"address_line1,omitempty"`
	// AddressLine2 is the second line of the address (apartment/suite or floor).
	// AddressLine2 地址第二行（公寓/套房号码或楼层）。
	AddressLine2 string `json:"address_line2,omitempty"`
	// CountryCode is the 2-letter ISO 3166-2 country code.
	// CountryCode 国家代码（2位ISO 3166-2代码）。
	CountryCode string `json:"country_code,omitempty"`
	// Postcode is the postal code.
	// Postcode 邮政编码。
	Postcode string `json:"postcode,omitempty"`
	// State is the state/province name.
	// State 州/省名称。
	State string `json:"state,omitempty"`
	// Suburb is the city/suburb name.
	// Suburb 城市/郊区名称。
	Suburb string `json:"suburb,omitempty"`
}

// RFIAttachment represents an attachment.
// RFIAttachment 附件。
type RFIAttachment struct {
	// FileID is the file ID returned by the File Service API.
	// FileID 文件服务API返回的文件ID。
	FileID string `json:"file_id,omitempty"`
}

// IdentityDocument represents an identity document.
// IdentityDocument 身份证件。
type IdentityDocument struct {
	// BackFileID is the file ID for the back of the document.
	// BackFileID 证件背面照片文件ID。
	BackFileID string `json:"back_file_id,omitempty"`
	// FrontFileID is the file ID for the front of the document.
	// FrontFileID 证件正面照片文件ID。
	FrontFileID string `json:"front_file_id,omitempty"`
	// IssuingCountry is the issuing country code (2-letter ISO 3166-2).
	// IssuingCountry 证件签发国家（2位ISO 3166-2代码）。
	IssuingCountry string `json:"issuing_country,omitempty"`
	// Number is the document number.
	// Number 证件号码。
	Number string `json:"number,omitempty"`
	// Type is the identity document type.
	// Type 身份证件类型。
	Type IdentityDocumentType `json:"type,omitempty"`
}

// RFIQuestionAnswer represents a question answer.
// RFIQuestionAnswer 问题答案。
type RFIQuestionAnswer struct {
	// Address is the address-type answer.
	// Address 地址类型答案。
	Address *RFIAddressAnswer `json:"address,omitempty"`
	// Attachments are the attachment-type answers.
	// Attachments 附件类型答案。
	Attachments []RFIAttachment `json:"attachments,omitempty"`
	// Comment is the reply comment.
	// Comment 回复备注。
	Comment string `json:"comment,omitempty"`
	// Confirmed is the confirmation-type answer.
	// Confirmed 确认类型答案。
	Confirmed bool `json:"confirmed,omitempty"`
	// IdentityDocument is the identity document-type answer.
	// IdentityDocument 身份证件类型答案。
	IdentityDocument *IdentityDocument `json:"identity_document,omitempty"`
	// Text is the text-type answer.
	// Text 文本类型答案。
	Text string `json:"text,omitempty"`
	// Type is the answer type.
	// Type 答案类型。
	Type RFIQuestionAnswerType `json:"type,omitempty"`
}

// RFILiveness represents a liveness verification.
// RFILiveness 活体认证。
type RFILiveness struct {
	// Status is the liveness verification status.
	// Status 活体认证状态。
	Status RFILivenessStatus `json:"status,omitempty"`
	// URL is the liveness verification URL (to be opened on a mobile device).
	// URL 活体认证URL（需在移动设备上打开）。
	URL string `json:"url,omitempty"`
}

// RFIPerson represents a person's information.
// RFIPerson 人员信息。
type RFIPerson struct {
	// FirstName is the first name.
	// FirstName 名字。
	FirstName string `json:"first_name,omitempty"`
	// LastName is the last name.
	// LastName 姓氏。
	LastName string `json:"last_name,omitempty"`
}

// RFISource represents source information.
// RFISource 来源信息。
type RFISource struct {
	// ID is the unique source identifier.
	// ID 来源唯一标识符。
	ID string `json:"id,omitempty"`
	// Person is the person information.
	// Person 人员信息。
	Person *RFIPerson `json:"person,omitempty"`
	// Type is the source type.
	// Type 来源类型。
	Type RFISourceType `json:"type,omitempty"`
}

// RFIUploadedOrder represents an uploaded order.
// RFIUploadedOrder 上传的订单。
type RFIUploadedOrder struct {
	// OrderID is the unique order identifier.
	// OrderID 订单唯一标识符。
	OrderID string `json:"order_id,omitempty"`
}

// RFIQuestion represents an RFI question.
// RFIQuestion RFI问题。
type RFIQuestion struct {
	// Answer is the question answer.
	// Answer 问题答案。
	Answer *RFIQuestionAnswer `json:"answer,omitempty"`
	// Attachments are the supporting documents for the question.
	// Attachments 问题的支持文档。
	Attachments []RFIAttachment `json:"attachments,omitempty"`
	// Comment is the additional note for the question.
	// Comment 问题的附加说明。
	Comment string `json:"comment,omitempty"`
	// Description is the localized question description.
	// Description 问题描述（本地化）。
	Description *LocalizedContent `json:"description,omitempty"`
	// ID is the unique question identifier.
	// ID 问题唯一标识符。
	ID string `json:"id,omitempty"`
	// Key is the unique code for the question scenario (supports automated responses).
	// Key 问题场景的唯一编码值（支持自动化响应）。
	Key RFIQuestionKey `json:"key,omitempty"`
	// Liveness is the liveness verification information.
	// Liveness 活体认证信息。
	Liveness *RFILiveness `json:"liveness,omitempty"`
	// Sources are the list of sources associated with the question.
	// Sources 问题关联的来源列表。
	Sources []RFISource `json:"sources,omitempty"`
	// Title is the localized question title.
	// Title 问题标题（本地化）。
	Title *LocalizedContent `json:"title,omitempty"`
	// UploadedOrders are the associated uploaded orders (for TRANSACTION type RFI only).
	// UploadedOrders 上传的关联订单列表（仅适用于TRANSACTION类型RFI）。
	UploadedOrders []RFIUploadedOrder `json:"uploaded_orders,omitempty"`
}

// RFIRequest represents an RFI request.
// RFIRequest RFI请求信息。
type RFIRequest struct {
	// CreatedAt is the RFI request creation timestamp.
	// CreatedAt RFI请求创建时间戳。
	CreatedAt string `json:"created_at,omitempty"`
	// ExpiresAt is the RFI request expiry timestamp (status becomes CLOSED after expiry).
	// ExpiresAt RFI请求过期时间戳（过期后状态转为CLOSED）。
	ExpiresAt string `json:"expires_at,omitempty"`
	// Questions is the list of questions.
	// Questions 问题列表。
	Questions []RFIQuestion `json:"questions,omitempty"`
	// UpdatedAt is the RFI request last update timestamp.
	// UpdatedAt RFI请求最后更新时间戳。
	UpdatedAt string `json:"updated_at,omitempty"`
}

// RFI represents a Request for Information.
// RFI 表示信息请求（Request for Information）。
type RFI struct {
	// AccountID is the Airwallex account ID.
	// AccountID Airwallex账户ID。
	AccountID string `json:"account_id,omitempty"`
	// ActiveRequest is the currently active RFI request.
	// ActiveRequest 当前活跃的RFI请求。
	ActiveRequest *RFIRequest `json:"active_request,omitempty"`
	// AnsweredRequests are the answered RFI requests.
	// AnsweredRequests 已响应的RFI请求列表。
	AnsweredRequests []RFIRequest `json:"answered_requests,omitempty"`
	// CreatedAt is the RFI creation timestamp.
	// CreatedAt RFI创建时间戳。
	CreatedAt string `json:"created_at,omitempty"`
	// ID is the unique RFI identifier.
	// ID RFI唯一标识符。
	ID string `json:"id,omitempty"`
	// Status is the RFI status.
	// Status RFI状态。
	Status RFIStatus `json:"status,omitempty"`
	// SubType is the RFI sub-type.
	// SubType RFI子类型。
	SubType RFISubType `json:"sub_type,omitempty"`
	// Type is the RFI type.
	// Type RFI类型。
	Type RFIType `json:"type,omitempty"`
	// UpdatedAt is the RFI last update timestamp.
	// UpdatedAt RFI最后更新时间戳。
	UpdatedAt string `json:"updated_at,omitempty"`
}

// RespondRFIRequest represents a request to respond to an RFI.
// RespondRFIRequest 回复RFI请求。
type RespondRFIRequest struct {
	// Questions is the list of questions (with answers).
	// Questions 问题列表（包含答案）。
	Questions []RFIQuestion `json:"questions,omitempty"`
}

// CreateRFIRequest represents a request to create an RFI.
// CreateRFIRequest 创建RFI请求。
type CreateRFIRequest struct {
	// Type is the RFI type. Required.
	// Type RFI类型。必填。
	Type RFIType `json:"type"`
	// Questions is the list of questions. Optional.
	// Questions 问题列表。可选。
	Questions []RFIQuestion `json:"questions,omitempty"`
}

// CreateRFI creates a new Request for Information.
// CreateRFI 创建信息请求。
// 官方文档: https://www.airwallex.com/docs/api/risk/request_for_information_rfi/create.md
func (s *Service) CreateRFI(ctx context.Context, req *CreateRFIRequest, opts ...sdk.RequestOption) (*RFI, error) {
	var resp RFI
	err := s.doer.Do(ctx, "POST", "/api/v1/rfis/create", req, &resp, opts...)
	return &resp, err
}

// GetRFI retrieves an RFI by ID.
// GetRFI 根据ID获取RFI。
// 官方文档: https://www.airwallex.com/docs/api/risk/request_for_information_rfi/retrieve.md
func (s *Service) GetRFI(ctx context.Context, id string, opts ...sdk.RequestOption) (*RFI, error) {
	var resp RFI
	err := s.doer.Do(ctx, "GET", "/api/v1/rfis/"+id, nil, &resp, opts...)
	return &resp, err
}

// RespondRFI responds to a Request for Information.
// RespondRFI 回复信息请求。
// 官方文档: https://www.airwallex.com/docs/api/risk/request_for_information_rfi/respond_rfis.md
func (s *Service) RespondRFI(ctx context.Context, id string, req *RespondRFIRequest, opts ...sdk.RequestOption) (*RFI, error) {
	var resp RFI
	err := s.doer.Do(ctx, "POST", "/api/v1/rfis/"+id+"/respond", req, &resp, opts...)
	return &resp, err
}

// ListRFIs lists all Requests for Information.
// ListRFIs 列出信息请求。
// 官方文档: https://www.airwallex.com/docs/api/risk/request_for_information_rfi/list.md
func (s *Service) ListRFIs(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[RFI], error) {
	var resp sdk.ListResult[RFI]
	err := s.doer.Do(ctx, "GET", "/api/v1/rfis", nil, &resp, opts...)
	return &resp, err
}
