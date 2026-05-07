// Package events provides typed webhook event structures for the account domain.
// Account 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/account.md
//
// 事件映射表:
//
//	account.active         → AccountActiveEvent         (Data: AccountEventData)
//	account.connected      → AccountConnectedEvent      (Data: AccountConnectedEventData)
//	account.suspended      → AccountSuspendedEvent      (Data: AccountEventData)
//	account.action_required → AccountActionRequiredEvent (Data: AccountActionRequiredEventData)
//	account.submitted      → AccountSubmittedEvent      (Data: AccountSubmittedEventData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Account Events ---

// AccountActiveEvent represents the account.active webhook event.
// AccountActiveEvent 表示 account.active webhook 事件。
type AccountActiveEvent struct {
	Event
	Data AccountEventData `json:"data"`
}

// AccountConnectedEvent represents the account.connected webhook event.
// AccountConnectedEvent 表示 account.connected webhook 事件。
type AccountConnectedEvent struct {
	Event
	Data AccountConnectedEventData `json:"data"`
}

// AccountSuspendedEvent represents the account.suspended webhook event.
// AccountSuspendedEvent 表示 account.suspended webhook 事件。
type AccountSuspendedEvent struct {
	Event
	Data AccountEventData `json:"data"`
}

// AccountActionRequiredEvent represents the account.action_required webhook event.
// AccountActionRequiredEvent 表示 account.action_required webhook 事件。
type AccountActionRequiredEvent struct {
	Event
	Data AccountActionRequiredEventData `json:"data"`
}

// AccountSubmittedEvent represents the account.submitted webhook event.
// AccountSubmittedEvent 表示 account.submitted webhook 事件。
type AccountSubmittedEvent struct {
	Event
	Data AccountSubmittedEventData `json:"data"`
}

// --- Event Data Structures ---

// AccountEventData represents the payload data for account.active and account.suspended events.
// AccountEventData 表示 account.active 和 account.suspended 事件的载荷数据。
// 官方文档说明 data 为 Get account by ID 的响应体。
type AccountEventData struct {
	AccountDetails     *AccountDetails     `json:"accountDetails,omitempty"`
	AccountUsage       *AccountUsage       `json:"accountUsage,omitempty"`
	CreatedAt          string              `json:"createdAt,omitempty"`
	CustomerAgreements *CustomerAgreements `json:"customerAgreements,omitempty"`
	ID                 string              `json:"id,omitempty"`
	Identifier         string              `json:"identifier,omitempty"`
	Metadata           map[string]any      `json:"metadata,omitempty"`
	NextAction         *NextAction         `json:"nextAction,omitempty"`
	PrimaryContact     *PrimaryContact     `json:"primaryContact,omitempty"`
	Requirements       *Requirements       `json:"requirements,omitempty"`
	Status             string              `json:"status,omitempty"`
}

// AccountConnectedEventData represents the payload data for account.connected event.
// AccountConnectedEventData 表示 account.connected 事件的载荷数据。
type AccountConnectedEventData struct {
	ConnectedAccountID         string `json:"connectedAccountId,omitempty"`
	ConnectedAccountIdentifier string `json:"connectedAccountIdentifier,omitempty"`
	ConnectedAccountName       string `json:"connectedAccountName,omitempty"`
	CreatedAt                  string `json:"createdAt,omitempty"`
	ConnectedAccountStatus     string `json:"connectedAccountStatus,omitempty"`
	PlatformAccountID          string `json:"platformAccountId,omitempty"`
	PlatformName               string `json:"platformName,omitempty"`
}

// AccountActionRequiredEventData represents the payload data for account.action_required event.
// AccountActionRequiredEventData 表示 account.action_required 事件的载荷数据。
type AccountActionRequiredEventData struct {
	AgreedToTermsAndConditionsRequired *TermsAndConditionsRequirement `json:"agreedToTermsAndConditionsRequired,omitempty"`
	PhotoFileIDRequired                *PhotoFileIDRequirement        `json:"photoFileIdRequired,omitempty"`
	PrimaryIdentificationRequired      []IdentificationRequirement    `json:"primaryIdentificationRequired,omitempty"`
	SecondaryIdentificationRequired    []IdentificationRequirement    `json:"secondaryIdentificationRequired,omitempty"`
}

// AccountSubmittedEventData represents the payload data for account.submitted event.
// AccountSubmittedEventData 表示 account.submitted 事件的载荷数据。
type AccountSubmittedEventData struct{}

// --- Account Details ---

// AccountDetails represents the account details in webhook payload.
// AccountDetails 表示 webhook payload 中的账户详情。
type AccountDetails struct {
	AuthorisedPersonDetails *AuthorisedPersonDetails `json:"authorisedPersonDetails,omitempty"`
	BeneficialOwners        []BeneficialOwner        `json:"beneficialOwners,omitempty"`
	BusinessDetails         *BusinessDetails         `json:"businessDetails,omitempty"`
	DirectorDetails         []DirectorDetails        `json:"directorDetails,omitempty"`
	IndividualDetails       *IndividualDetails       `json:"individualDetails,omitempty"`
	LegalEntityType         string                   `json:"legalEntityType,omitempty"`
	LegalRepDetails         *LegalRepDetails         `json:"legalRepDetails,omitempty"`
	TrusteeDetails          *TrusteeDetails          `json:"trusteeDetails,omitempty"`
}

// AuthorisedPersonDetails represents authorised person details.
// AuthorisedPersonDetails 表示授权人详情。
type AuthorisedPersonDetails struct {
	Attachments          *Attachments `json:"attachments,omitempty"`
	Email                string       `json:"email,omitempty"`
	FillingAs            string       `json:"fillingAs,omitempty"`
	FirstName            string       `json:"firstName,omitempty"`
	FirstNameEnglish     string       `json:"firstNameEnglish,omitempty"`
	IdentificationNumber string       `json:"identificationNumber,omitempty"`
	LastName             string       `json:"lastName,omitempty"`
	LastNameEnglish      string       `json:"lastNameEnglish,omitempty"`
	Nationality          string       `json:"nationality,omitempty"`
}

// Attachments represents file attachments.
// Attachments 表示文件附件。
type Attachments struct {
	IdentityFiles []IdentityFile `json:"identityFiles,omitempty"`
}

// IdentityFile represents an identity file attachment.
// IdentityFile 表示身份证明文件附件。
type IdentityFile struct {
	Description string `json:"description,omitempty"`
	FileID      string `json:"fileId,omitempty"`
	Tag         string `json:"tag,omitempty"`
}

// BeneficialOwner represents a beneficial owner.
// BeneficialOwner 表示受益所有人。
type BeneficialOwner struct {
	Address              *Address     `json:"address,omitempty"`
	Attachments          *Attachments `json:"attachments,omitempty"`
	DateOfBirth          string       `json:"dateOfBirth,omitempty"`
	FirstName            string       `json:"firstName,omitempty"`
	FirstNameEnglish     string       `json:"firstNameEnglish,omitempty"`
	IdentificationNumber string       `json:"identificationNumber,omitempty"`
	IdentificationType   string       `json:"identificationType,omitempty"`
	LastName             string       `json:"lastName,omitempty"`
	LastNameEnglish      string       `json:"lastNameEnglish,omitempty"`
	Nationality          string       `json:"nationality,omitempty"`
}

// Address represents an address.
// Address 表示地址。
type Address struct {
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	CountryCode  string `json:"countryCode,omitempty"`
	Postcode     string `json:"postcode,omitempty"`
	State        string `json:"state,omitempty"`
	Suburb       string `json:"suburb,omitempty"`
}

// BusinessDetails represents business details.
// BusinessDetails 表示企业详情。
type BusinessDetails struct {
	Address                      *Address             `json:"address,omitempty"`
	AddressEnglish               *Address             `json:"addressEnglish,omitempty"`
	AsTrustee                    bool                 `json:"asTrustee,omitempty"`
	Attachments                  *BusinessAttachments `json:"attachments,omitempty"`
	BusinessAddress              *Address             `json:"businessAddress,omitempty"`
	BusinessName                 string               `json:"businessName,omitempty"`
	BusinessNameEnglish          string               `json:"businessNameEnglish,omitempty"`
	BusinessRegistrationNumber   string               `json:"businessRegistrationNumber,omitempty"`
	BusinessStructure            string               `json:"businessStructure,omitempty"`
	ContactNumber                string               `json:"contactNumber,omitempty"`
	DescriptionOfGoodsOrServices string               `json:"descriptionOfGoodsOrServices,omitempty"`
	DescriptionOfIndustry        string               `json:"descriptionOfIndustry,omitempty"`
	IndustryCategory             string               `json:"industryCategory,omitempty"`
	IndustryCategoryLevel3       string               `json:"industryCategoryLevel3,omitempty"`
	IndustrySubCategory          string               `json:"industrySubCategory,omitempty"`
	OperatingCountry             []string             `json:"operatingCountry,omitempty"`
	Purpose                      string               `json:"purpose,omitempty"`
	TrustName                    string               `json:"trustName,omitempty"`
	URL                          string               `json:"url,omitempty"`
	VATNumbers                   []VATNumber          `json:"vatNumbers,omitempty"`
}

// BusinessAttachments represents business document attachments.
// BusinessAttachments 表示企业文件附件。
type BusinessAttachments struct {
	BusinessDocuments []BusinessDocument `json:"businessDocuments,omitempty"`
}

// BusinessDocument represents a business document.
// BusinessDocument 表示企业文件。
type BusinessDocument struct {
	Description string `json:"description,omitempty"`
	FileID      string `json:"fileId,omitempty"`
	Tag         string `json:"tag,omitempty"`
}

// VATNumber represents a VAT number.
// VATNumber 表示增值税号。
type VATNumber struct {
	CountryCode string `json:"countryCode,omitempty"`
	VATNumber   string `json:"vatNumber,omitempty"`
}

// DirectorDetails represents director details.
// DirectorDetails 表示董事详情。
type DirectorDetails struct {
	Address              *Address     `json:"address,omitempty"`
	Attachments          *Attachments `json:"attachments,omitempty"`
	DateOfBirth          string       `json:"dateOfBirth,omitempty"`
	FirstName            string       `json:"firstName,omitempty"`
	FirstNameEnglish     string       `json:"firstNameEnglish,omitempty"`
	IdentificationNumber string       `json:"identificationNumber,omitempty"`
	IdentificationType   string       `json:"identificationType,omitempty"`
	LastName             string       `json:"lastName,omitempty"`
	LastNameEnglish      string       `json:"lastNameEnglish,omitempty"`
	Nationality          string       `json:"nationality,omitempty"`
}

// IndividualDetails represents individual account holder details.
// IndividualDetails 表示个人账户持有人详情。
type IndividualDetails struct {
	Address                 *Address               `json:"address,omitempty"`
	AddressEnglish          *Address               `json:"addressEnglish,omitempty"`
	Attachments             *IndividualAttachments `json:"attachments,omitempty"`
	DateOfBirth             string                 `json:"dateOfBirth,omitempty"`
	FirstName               string                 `json:"firstName,omitempty"`
	FirstNameEnglish        string                 `json:"firstNameEnglish,omitempty"`
	LastName                string                 `json:"lastName,omitempty"`
	LastNameEnglish         string                 `json:"lastNameEnglish,omitempty"`
	MiddleName              string                 `json:"middleName,omitempty"`
	Nationality             string                 `json:"nationality,omitempty"`
	PhotoFileID             string                 `json:"photoFileId,omitempty"`
	PrimaryIdentification   *Identification        `json:"primaryIdentification,omitempty"`
	SecondaryIdentification *Identification        `json:"secondaryIdentification,omitempty"`
	UserIDOnPlatform        string                 `json:"userIdOnPlatform,omitempty"`
}

// IndividualAttachments represents individual document attachments.
// IndividualAttachments 表示个人文件附件。
type IndividualAttachments struct {
	IndividualDocuments []IdentityFile `json:"individualDocuments,omitempty"`
}

// Identification represents identification documents.
// Identification 表示身份证明文件。
type Identification struct {
	DriversLicense     *DriversLicense `json:"driversLicense,omitempty"`
	IdentificationType string          `json:"identificationType,omitempty"`
	IssuingCountryCode string          `json:"issuingCountryCode,omitempty"`
	MedicareCard       *MedicareCard   `json:"medicareCard,omitempty"`
	Passport           *Passport       `json:"passport,omitempty"`
	PersonalID         *PersonalID     `json:"personalId,omitempty"`
}

// DriversLicense represents a driver's license.
// DriversLicense 表示驾照。
type DriversLicense struct {
	BackFileID    string `json:"backFileId,omitempty"`
	EffectiveAt   string `json:"effectiveAt,omitempty"`
	ExpireAt      string `json:"expireAt,omitempty"`
	FrontFileID   string `json:"frontFileId,omitempty"`
	Gender        string `json:"gender,omitempty"`
	IssuingState  string `json:"issuingState,omitempty"`
	LicenseNumber string `json:"licenseNumber,omitempty"`
	Version       string `json:"version,omitempty"`
}

// MedicareCard represents a Medicare card.
// MedicareCard 表示医保卡。
type MedicareCard struct {
	BackFileID      string `json:"backFileId,omitempty"`
	CardNumber      string `json:"cardNumber,omitempty"`
	Color           string `json:"color,omitempty"`
	EffectiveAt     string `json:"effectiveAt,omitempty"`
	ExpireAt        string `json:"expireAt,omitempty"`
	FrontFileID     string `json:"frontFileId,omitempty"`
	ReferenceNumber string `json:"referenceNumber,omitempty"`
}

// Passport represents a passport.
// Passport 表示护照。
type Passport struct {
	EffectiveAt    string `json:"effectiveAt,omitempty"`
	ExpireAt       string `json:"expireAt,omitempty"`
	FrontFileID    string `json:"frontFileId,omitempty"`
	MRZLine1       string `json:"mrzLine1,omitempty"`
	MRZLine2       string `json:"mrzLine2,omitempty"`
	PassportNumber string `json:"passportNumber,omitempty"`
}

// PersonalID represents a personal ID.
// PersonalID 表示个人身份证。
type PersonalID struct {
	BackFileID  string `json:"backFileId,omitempty"`
	EffectiveAt string `json:"effectiveAt,omitempty"`
	ExpireAt    string `json:"expireAt,omitempty"`
	FrontFileID string `json:"frontFileId,omitempty"`
	IDNumber    string `json:"idNumber,omitempty"`
}

// LegalRepDetails represents legal representative details.
// LegalRepDetails 表示法定代表人详情。
type LegalRepDetails struct {
	Address              *Address     `json:"address,omitempty"`
	Attachments          *Attachments `json:"attachments,omitempty"`
	DateOfBirth          string       `json:"dateOfBirth,omitempty"`
	FirstName            string       `json:"firstName,omitempty"`
	FirstNameEnglish     string       `json:"firstNameEnglish,omitempty"`
	IdentificationNumber string       `json:"identificationNumber,omitempty"`
	IdentificationType   string       `json:"identificationType,omitempty"`
	LastName             string       `json:"lastName,omitempty"`
	LastNameEnglish      string       `json:"lastNameEnglish,omitempty"`
	Nationality          string       `json:"nationality,omitempty"`
}

// TrusteeDetails represents trustee details.
// TrusteeDetails 表示受托人详情。
type TrusteeDetails struct {
	BusinessDetails   *TrusteeBusinessDetails    `json:"businessDetails,omitempty"`
	IndividualDetails []TrusteeIndividualDetails `json:"individualDetails,omitempty"`
	LegalEntityType   string                     `json:"legalEntityType,omitempty"`
}

// TrusteeBusinessDetails represents trustee business details.
// TrusteeBusinessDetails 表示受托人企业详情。
type TrusteeBusinessDetails struct {
	Address                      *Address             `json:"address,omitempty"`
	AddressEnglish               *Address             `json:"addressEnglish,omitempty"`
	Attachments                  *BusinessAttachments `json:"attachments,omitempty"`
	BeneficialOwners             []BeneficialOwner    `json:"beneficialOwners,omitempty"`
	BusinessName                 string               `json:"businessName,omitempty"`
	BusinessNameEnglish          string               `json:"businessNameEnglish,omitempty"`
	BusinessRegistrationNumber   string               `json:"businessRegistrationNumber,omitempty"`
	BusinessStructure            string               `json:"businessStructure,omitempty"`
	ContactNumber                string               `json:"contactNumber,omitempty"`
	DescriptionOfGoodsOrServices string               `json:"descriptionOfGoodsOrServices,omitempty"`
	DescriptionOfIndustry        string               `json:"descriptionOfIndustry,omitempty"`
	DirectorDetails              []DirectorDetails    `json:"directorDetails,omitempty"`
	IndustryCategory             string               `json:"industryCategory,omitempty"`
	IndustryCategoryLevel3       string               `json:"industryCategoryLevel3,omitempty"`
	IndustrySubCategory          string               `json:"industrySubCategory,omitempty"`
	OperatingCountry             []string             `json:"operatingCountry,omitempty"`
	URL                          string               `json:"url,omitempty"`
}

// TrusteeIndividualDetails represents trustee individual details.
// TrusteeIndividualDetails 表示受托人个人详情。
type TrusteeIndividualDetails struct {
	Address              *Address     `json:"address,omitempty"`
	Attachments          *Attachments `json:"attachments,omitempty"`
	DateOfBirth          string       `json:"dateOfBirth,omitempty"`
	FirstName            string       `json:"firstName,omitempty"`
	FirstNameEnglish     string       `json:"firstNameEnglish,omitempty"`
	IdentificationNumber string       `json:"identificationNumber,omitempty"`
	IdentificationType   string       `json:"identificationType,omitempty"`
	LastName             string       `json:"lastName,omitempty"`
	LastNameEnglish      string       `json:"lastNameEnglish,omitempty"`
	Nationality          string       `json:"nationality,omitempty"`
}

// --- Account Usage ---

// AccountUsage represents account usage information.
// AccountUsage 表示账户使用信息。
type AccountUsage struct {
	CardUsage                        []string                          `json:"cardUsage,omitempty"`
	CollectionCountryCodes           []string                          `json:"collectionCountryCodes,omitempty"`
	CollectionFrom                   string                            `json:"collectionFrom,omitempty"`
	ExpectedMonthlyTransactionVolume *ExpectedMonthlyTransactionVolume `json:"expectedMonthlyTransactionVolume,omitempty"`
	PayoutCountryCodes               []string                          `json:"payoutCountryCodes,omitempty"`
	PayoutTo                         string                            `json:"payoutTo,omitempty"`
}

// ExpectedMonthlyTransactionVolume represents expected monthly transaction volume.
// ExpectedMonthlyTransactionVolume 表示预期月交易量。
type ExpectedMonthlyTransactionVolume struct {
	Amount string `json:"amount,omitempty"`
}

// --- Customer Agreements ---

// CustomerAgreements represents customer agreement statuses.
// CustomerAgreements 表示客户协议状态。
type CustomerAgreements struct {
	AgreedToDataUsage          bool `json:"agreedToDataUsage,omitempty"`
	AgreedToTermsAndConditions bool `json:"agreedToTermsAndConditions,omitempty"`
	OptInForMarketing          bool `json:"optInForMarketing,omitempty"`
}

// --- Next Action ---

// NextAction represents the next action required.
// NextAction 表示所需的下个操作。
type NextAction struct {
	Type string `json:"type,omitempty"`
}

// --- Primary Contact ---

// PrimaryContact represents primary contact information.
// PrimaryContact 表示主要联系人信息。
type PrimaryContact struct {
	Email  string `json:"email,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}

// --- Requirements ---

// Requirements represents account requirements.
// Requirements 表示账户要求。
type Requirements struct {
	AgreementToTermsAndConditionsRequired *TermsAndConditionsRequirement `json:"agreementToTermsAndConditionsRequired,omitempty"`
	PhotoFileIDRequired                   *PhotoFileIDRequirement        `json:"photoFileIdRequired,omitempty"`
	PrimaryIdentificationRequired         []IdentificationRequirement    `json:"primaryIdentificationRequired,omitempty"`
	SecondaryIdentificationRequired       []IdentificationRequirement    `json:"secondaryIdentificationRequired,omitempty"`
}

// TermsAndConditionsRequirement represents terms and conditions requirement.
// TermsAndConditionsRequirement 表示条款和条件要求。
type TermsAndConditionsRequirement struct {
	Endpoint string `json:"endpoint,omitempty"`
	Message  string `json:"message,omitempty"`
	Method   string `json:"method,omitempty"`
}

// PhotoFileIDRequirement represents photo file ID requirement.
// PhotoFileIDRequirement 表示照片文件 ID 要求。
type PhotoFileIDRequirement struct {
	Message     string `json:"message,omitempty"`
	PhotoFileID string `json:"photoFileId,omitempty"`
}

// IdentificationRequirement represents identification requirement.
// IdentificationRequirement 表示身份证明要求。
type IdentificationRequirement struct {
	BackFileID  string `json:"backFileId,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	FrontFileID string `json:"frontFileId,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Message     string `json:"message,omitempty"`
	MiddleName  string `json:"middleName,omitempty"`
}
