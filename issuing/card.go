package issuing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// CardStatus represents the card status.
// CardStatus 卡片状态。
type CardStatus = string

const (
	// CardStatusActive indicates the card is active.
	// CardStatusActive 活跃状态
	CardStatusActive CardStatus = "ACTIVE"
	// CardStatusInactive indicates the card is inactive.
	// CardStatusInactive 非活跃状态
	CardStatusInactive CardStatus = "INACTIVE"
	// CardStatusBlocked indicates the card is blocked.
	// CardStatusBlocked 已冻结
	CardStatusBlocked CardStatus = "BLOCKED"
)

// CardFormFactor represents the card form factor.
// CardFormFactor 卡片形态。
type CardFormFactor = string

const (
	// CardFormFactorVirtual is a virtual card.
	// CardFormFactorVirtual 虚拟卡
	CardFormFactorVirtual CardFormFactor = "VIRTUAL"
	// CardFormFactorPhysical is a physical card.
	// CardFormFactorPhysical 实体卡
	CardFormFactorPhysical CardFormFactor = "PHYSICAL"
)

// CardBrand represents the card brand.
// CardBrand 卡片品牌。
type CardBrand = string

const (
	// CardBrandVisa is a Visa card.
	// CardBrandVisa Visa
	CardBrandVisa CardBrand = "VISA"
	// CardBrandMastercard is a Mastercard card.
	// CardBrandMastercard Mastercard
	CardBrandMastercard CardBrand = "MASTERCARD"
)

// CardProgram represents card program configuration.
// CardProgram 卡片程序配置。
type CardProgram struct {
	// Purpose is the purpose. Required.
	// Purpose 用途。必填。
	Purpose string `json:"purpose"`
	// Type is the type. Optional.
	// Type 类型。可选。
	Type string `json:"type,omitempty"`
	// SubType is the sub type. Optional.
	// SubType 子类型。可选。
	SubType string `json:"sub_type,omitempty"`
	// InterchangePercent is the interchange fee percentage. Optional.
	// InterchangePercent 交换费百分比。可选。
	InterchangePercent float64 `json:"interchange_percent,omitempty"`
}

// AuthorizationControls represents authorization control configuration.
// AuthorizationControls 授权控制配置。
type AuthorizationControls struct {
	// ActiveFrom is the effective start time. Optional.
	// ActiveFrom 生效起始时间。可选。
	ActiveFrom string `json:"active_from,omitempty"`
	// ActiveTo is the effective end time. Optional.
	// ActiveTo 生效截止时间。可选。
	ActiveTo string `json:"active_to,omitempty"`
	// AllowedCategories is the list of allowed categories. Optional.
	// AllowedCategories 允许的类别列表。可选。
	AllowedCategories []string `json:"allowed_categories,omitempty"`
	// BlockedCategories is the list of blocked categories. Optional.
	// BlockedCategories 阻止的类别列表。可选。
	BlockedCategories []string `json:"blocked_categories,omitempty"`
	// AllowedCurrencies is the list of allowed currencies. Optional.
	// AllowedCurrencies 允许的币种列表。可选。
	AllowedCurrencies []sdk.Currency `json:"allowed_currencies,omitempty"`
	// AllowedMerchantCategories is the list of allowed merchant category codes. Optional.
	// AllowedMerchantCategories 允许的商户类别代码列表。可选。
	AllowedMerchantCategories []string `json:"allowed_merchant_categories,omitempty"`
	// AllowedMerchants is the list of allowed merchants. Optional.
	// AllowedMerchants 允许的商户列表。可选。
	AllowedMerchants []string `json:"allowed_merchants,omitempty"`
	// BlockedMerchants is the list of blocked merchants. Optional.
	// BlockedMerchants 阻止的商户列表。可选。
	BlockedMerchants []string `json:"blocked_merchants,omitempty"`
	// AllowedTransactionCount is the allowed number of transactions. Optional.
	// AllowedTransactionCount 允许的交易次数。可选。
	AllowedTransactionCount int `json:"allowed_transaction_count,omitempty"`
	// BlockedTransactionUsages is the list of blocked transaction usages. Optional.
	// BlockedTransactionUsages 被阻止的交易用途列表。可选。
	BlockedTransactionUsages []string `json:"blocked_transaction_usages,omitempty"`
	// MaxTransactionAmount is the maximum transaction amount. Optional.
	// MaxTransactionAmount 单笔交易最大金额。可选。
	MaxTransactionAmount float64 `json:"max_transaction_amount,omitempty"`
	// DailyTransactionLimit is the daily transaction limit. Optional.
	// DailyTransactionLimit 日交易限额。可选。
	DailyTransactionLimit float64 `json:"daily_transaction_limit,omitempty"`
	// WeeklyTransactionLimit is the weekly transaction limit. Optional.
	// WeeklyTransactionLimit 周交易限额。可选。
	WeeklyTransactionLimit float64 `json:"weekly_transaction_limit,omitempty"`
	// MonthlyTransactionLimit is the monthly transaction limit. Optional.
	// MonthlyTransactionLimit 月交易限额。可选。
	MonthlyTransactionLimit float64 `json:"monthly_transaction_limit,omitempty"`
	// TransactionLimits is the list of transaction limit configurations. Optional.
	// TransactionLimits 交易限额配置列表。可选。
	TransactionLimits []map[string]any `json:"transaction_limits,omitempty"`
}

// DeliveryDetails represents card delivery details.
// DeliveryDetails 卡片配送详情。
type DeliveryDetails struct {
	// Method is the delivery method. Optional.
	// Method 配送方式。可选。
	Method string `json:"method,omitempty"`
	// Address is the delivery address. Optional.
	// Address 配送地址。可选。
	Address map[string]any `json:"address,omitempty"`
	// RecipientName is the recipient name. Optional.
	// RecipientName 收件人姓名。可选。
	RecipientName string `json:"recipient_name,omitempty"`
	// PhoneNumber is the recipient phone number. Optional.
	// PhoneNumber 收件人电话号码。可选。
	PhoneNumber string `json:"phone_number,omitempty"`
	// TrackingNumber is the tracking number. Optional.
	// TrackingNumber 物流追踪号。可选。
	TrackingNumber string `json:"tracking_number,omitempty"`
	// Status is the delivery status. Optional.
	// Status 配送状态。可选。
	Status string `json:"status,omitempty"`
	// DeliveryMode is the delivery mode. Optional.
	// DeliveryMode 配送模式。可选。
	DeliveryMode string `json:"delivery_mode,omitempty"`
	// DeliveryVendor is the delivery vendor. Optional.
	// DeliveryVendor 配送供应商。可选。
	DeliveryVendor string `json:"delivery_vendor,omitempty"`
	// MobileNumber is the mobile number. Optional.
	// MobileNumber 手机号。可选。
	MobileNumber string `json:"mobile_number,omitempty"`
	// PreferredDeliveryMode is the preferred delivery mode. Optional.
	// PreferredDeliveryMode 首选配送模式。可选。
	PreferredDeliveryMode string `json:"preferred_delivery_mode,omitempty"`
	// StatusDescription is the status description. Optional.
	// StatusDescription 状态描述。可选。
	StatusDescription string `json:"status_description,omitempty"`
	// Tracked indicates whether the delivery is trackable. Optional.
	// Tracked 是否可追踪。可选。
	Tracked bool `json:"tracked,omitempty"`
	// TrackingLink is the tracking link. Optional.
	// TrackingLink 追踪链接。可选。
	TrackingLink string `json:"tracking_link,omitempty"`
	// UpdatedAt is the update time. Optional.
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Card represents card information.
// Card 表示卡片信息。
type Card struct {
	// ID is the unique card identifier. Required.
	// ID 卡片唯一标识符。必填。
	ID string `json:"card_id"`
	// CardholderID is the unique cardholder identifier. Required.
	// CardholderID 持卡人唯一标识符。必填。
	CardholderID string `json:"cardholder_id"`
	// Status is the card status. Required.
	// Status 卡片状态。必填。
	Status CardStatus `json:"card_status"`
	// CardNumber is the card number. Optional.
	// CardNumber 卡号。可选。
	CardNumber string `json:"card_number,omitempty"`
	// Brand is the brand. Optional.
	// Brand 品牌。可选。
	Brand string `json:"brand,omitempty"`
	// NameOnCard is the name on the card. Optional.
	// NameOnCard 卡片上的名称。可选。
	NameOnCard string `json:"name_on_card,omitempty"`
	// AuthorizationControls contains authorization control configuration. Optional.
	// AuthorizationControls 授权控制配置。可选。
	AuthorizationControls *AuthorizationControls `json:"authorization_controls,omitempty"`
	// DeliveryDetails contains delivery details. Optional.
	// DeliveryDetails 配送详情。可选。
	DeliveryDetails *DeliveryDetails `json:"delivery_details,omitempty"`
	// FundingSourceID is the funding source unique identifier. Optional.
	// FundingSourceID 资金来源唯一标识符。可选。
	FundingSourceID string `json:"funding_source_id,omitempty"`
	// Metadata is the metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// Program is the card program configuration. Optional.
	// Program 卡片程序配置。可选。
	Program CardProgram `json:"program,omitempty"`
	// IsPersonalized indicates whether the card is personalized. Optional.
	// IsPersonalized 是否个性化。可选。
	IsPersonalized bool `json:"is_personalized,omitempty"`
	// FormFactor is the card form factor. Optional.
	// FormFactor 卡形态。可选。
	FormFactor string `json:"form_factor,omitempty"`
	// CreatedBy is the creator. Optional.
	// CreatedBy 创建者。可选。
	CreatedBy string `json:"created_by,omitempty"`
	// Purpose is the purpose. Optional.
	// Purpose 用途。可选。
	Purpose string `json:"purpose,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt is the update time. Optional.
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
	// ActivateOnIssue indicates whether to activate on issuance. Optional.
	// ActivateOnIssue 发卡时自动激活。可选。
	ActivateOnIssue bool `json:"activate_on_issue,omitempty"`
	// AdditionalCardholderIDs is the list of additional cardholder IDs. Optional.
	// AdditionalCardholderIDs 附加持卡人ID列表。可选。
	AdditionalCardholderIDs []string `json:"additional_cardholder_ids,omitempty"`
	// AlertSettings contains alert settings. Optional.
	// AlertSettings 提醒设置。可选。
	AlertSettings map[string]any `json:"alert_settings,omitempty"`
	// AllCardVersions contains all card versions. Optional.
	// AllCardVersions 所有卡片版本。可选。
	AllCardVersions []map[string]any `json:"all_card_versions,omitempty"`
	// CardVersion is the card version. Optional.
	// CardVersion 卡片版本。可选。
	CardVersion int `json:"card_version,omitempty"`
	// ClientData contains client data. Optional.
	// ClientData 客户端数据。可选。
	ClientData map[string]any `json:"client_data,omitempty"`
	// IssueTo is the issue target. Optional.
	// IssueTo 发卡对象。可选。
	IssueTo string `json:"issue_to,omitempty"`
	// NickName is the nickname. Optional.
	// NickName 昵称。可选。
	NickName string `json:"nick_name,omitempty"`
	// Note is the note. Optional.
	// Note 备注。可选。
	Note string `json:"note,omitempty"`
	// PostalAddress is the postal address. Optional.
	// PostalAddress 邮寄地址。可选。
	PostalAddress *CardholderAddress `json:"postal_address,omitempty"`
	// PrimaryContactDetails contains primary contact details. Optional.
	// PrimaryContactDetails 主要联系方式。可选。
	PrimaryContactDetails map[string]any `json:"primary_contact_details,omitempty"`
	// RequestID is the unique request identifier. Optional.
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
}

// CreateCardRequest represents the request to create a card.
// CreateCardRequest 创建卡片请求。
type CreateCardRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// CardholderID is the unique cardholder identifier. Required.
	// CardholderID 持卡人唯一标识符。必填。
	CardholderID string `json:"cardholder_id"`
	// Program is the card program configuration. Required.
	// Program 卡片程序配置。必填。
	Program CardProgram `json:"program"`
	// IsPersonalized indicates whether the card is personalized. Required.
	// IsPersonalized 是否个性化。必填。
	IsPersonalized bool `json:"is_personalized"`
	// FormFactor is the card form factor. Required.
	// FormFactor 卡形态。必填。
	FormFactor string `json:"form_factor"`
	// CreatedBy is the creator. Required.
	// CreatedBy 创建者。必填。
	CreatedBy string `json:"created_by"`
	// Purpose is the purpose. Optional.
	// Purpose 用途。可选。
	Purpose string `json:"purpose,omitempty"`
	// ActivateOnIssue indicates whether to activate on issuance. Optional.
	// ActivateOnIssue 发卡时自动激活。可选。
	ActivateOnIssue bool `json:"activate_on_issue,omitempty"`
	// AdditionalCardholderIDs is the list of additional cardholder IDs. Optional.
	// AdditionalCardholderIDs 附加持卡人ID列表。可选。
	AdditionalCardholderIDs []string `json:"additional_cardholder_ids,omitempty"`
	// AlertSettings contains alert settings. Optional.
	// AlertSettings 提醒设置。可选。
	AlertSettings map[string]any `json:"alert_settings,omitempty"`
	// AuthorizationControls contains authorization control configuration. Optional.
	// AuthorizationControls 授权控制配置。可选。
	AuthorizationControls *AuthorizationControls `json:"authorization_controls,omitempty"`
	// Brand is the brand. Optional.
	// Brand 品牌。可选。
	Brand string `json:"brand,omitempty"`
	// ClientData contains client data. Optional.
	// ClientData 客户端数据。可选。
	ClientData map[string]any `json:"client_data,omitempty"`
	// DeliveryDetails contains delivery details. Optional.
	// DeliveryDetails 配送详情。可选。
	DeliveryDetails *DeliveryDetails `json:"delivery_details,omitempty"`
	// FundingSourceID is the funding source unique identifier. Optional.
	// FundingSourceID 资金来源唯一标识符。可选。
	FundingSourceID string `json:"funding_source_id,omitempty"`
	// Metadata is the metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// NickName is the nickname. Optional.
	// NickName 昵称。可选。
	NickName string `json:"nick_name,omitempty"`
	// Note is the note. Optional.
	// Note 备注。可选。
	Note string `json:"note,omitempty"`
	// PostalAddress is the postal address. Optional.
	// PostalAddress 邮寄地址。可选。
	PostalAddress *CardholderAddress `json:"postal_address,omitempty"`
}

// CreateCard creates a card.
// 官方文档: https://www.airwallex.com/docs/api/issuing/cards/create.md
// CreateCard 创建卡片。
func (s *Service) CreateCard(ctx context.Context, req *CreateCardRequest, opts ...sdk.RequestOption) (*Card, error) {
	var resp Card
	err := s.doer.Do(ctx, "POST", "/api/v1/issuing/cards/create", req, &resp, opts...)
	return &resp, err
}

// GetCard retrieves a card by ID.
// 官方文档: https://www.airwallex.com/docs/api/issuing/cards/retrieve.md
// GetCard 根据 ID 获取卡片。
func (s *Service) GetCard(ctx context.Context, id string, opts ...sdk.RequestOption) (*Card, error) {
	var resp Card
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/cards/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateCard updates a card.
// 官方文档: https://www.airwallex.com/docs/api/issuing/cards/update.md
// UpdateCard 更新卡片。
func (s *Service) UpdateCard(ctx context.Context, id string, req map[string]any, opts ...sdk.RequestOption) (*Card, error) {
	var resp Card
	err := s.doer.Do(ctx, "POST", "/api/v1/issuing/cards/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ListCards lists cards.
// 官方文档: https://www.airwallex.com/docs/api/issuing/cards/list.md
// ListCards 列出卡片。
func (s *Service) ListCards(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[Card], error) {
	var resp sdk.ListResult[Card]
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/cards", nil, &resp, opts...)
	return &resp, err
}

// CardDetails represents sensitive card details.
// CardDetails 表示卡片敏感详情。
type CardDetails struct {
	// CardNumber is the card number. Required.
	// CardNumber 卡号。必填。
	CardNumber string `json:"card_number"`
	// CVV is the card verification value. Required.
	// CVV 卡片验证码。必填。
	CVV string `json:"cvv"`
	// ExpiryMonth is the expiry month. Required.
	// ExpiryMonth 到期月份。必填。
	ExpiryMonth int32 `json:"expiry_month"`
	// ExpiryYear is the expiry year. Required.
	// ExpiryYear 到期年份。必填。
	ExpiryYear int32 `json:"expiry_year"`
	// NameOnCard is the name on the card. Required.
	// NameOnCard 卡片上的名称。必填。
	NameOnCard string `json:"name_on_card"`
}

// GetCardDetails retrieves sensitive card details.
// 官方文档: https://www.airwallex.com/docs/api/issuing/cards/details.md
// GetCardDetails 获取卡片敏感详情。
func (s *Service) GetCardDetails(ctx context.Context, id string, opts ...sdk.RequestOption) (*CardDetails, error) {
	var resp CardDetails
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/cards/"+id+"/details", nil, &resp, opts...)
	return &resp, err
}

// CardLimit represents card limit details.
// CardLimit 表示卡片限额详情。
type CardLimit struct {
	// Amount is the limit amount. Required.
	// Amount 限额金额。必填。
	Amount float64 `json:"amount"`
	// Interval is the limit interval. Required.
	// Interval 限额周期。必填。
	Interval string `json:"interval"`
	// Remaining is the remaining limit. Required.
	// Remaining 剩余限额。必填。
	Remaining float64 `json:"remaining"`
}

// CardLimits represents the remaining card limits response.
// CardLimits 表示卡片剩余限额响应。
type CardLimits struct {
	// CashWithdrawalLimits is the list of cash withdrawal limits. Optional.
	// CashWithdrawalLimits 取现限额列表。可选。
	CashWithdrawalLimits []CardLimit `json:"cash_withdrawal_limits,omitempty"`
	// Currency is the limit currency. Optional.
	// Currency 限额币种。可选。
	Currency string `json:"currency,omitempty"`
	// Limits is the list of limits. Optional.
	// Limits 限额列表。可选。
	Limits []CardLimit `json:"limits,omitempty"`
}

// GetCardLimits retrieves remaining card limits.
// 官方文档: https://www.airwallex.com/docs/api/issuing/cards/limits.md
// GetCardLimits 获取卡片剩余限额。
func (s *Service) GetCardLimits(ctx context.Context, id string, opts ...sdk.RequestOption) (*CardLimits, error) {
	var resp CardLimits
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/cards/"+id+"/limits", nil, &resp, opts...)
	return &resp, err
}

// ActivateCard activates a physical card.
// 官方文档: https://www.airwallex.com/docs/api/issuing/cards/activate.md
// ActivateCard 激活实体卡片。
func (s *Service) ActivateCard(ctx context.Context, id string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/issuing/cards/"+id+"/activate", nil, nil, opts...)
}
