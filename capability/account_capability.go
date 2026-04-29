package capability

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// CapabilityID represents an account capability identifier.
// CapabilityID 账户能力 ID。
type CapabilityID = string

// CapabilityStatus represents the status of an account capability.
// CapabilityStatus 账户能力状态。
type CapabilityStatus = string

// EntityType represents the declarant entity type.
// EntityType 申报实体类型。
type EntityType = string

// Account capability status constants.
// 账户能力状态常量。
const (
	CapabilityStatusEnabled  CapabilityStatus = "ENABLED"
	CapabilityStatusDisabled CapabilityStatus = "DISABLED"
	CapabilityStatusPending  CapabilityStatus = "PENDING"
)

// Declarant entity type constants.
// 申报实体类型常量。
const (
	EntityTypeIndividual EntityType = "INDIVIDUAL"
	EntityTypeBusiness   EntityType = "BUSINESS"
)

// Transfer capability constants.
// 转账能力常量。
const (
	CapabilityTransferCNYLocal CapabilityID = "transfer_cny_local"
)

// Payment capability constants - bank cards.
// 支付能力常量 - 银行卡。
const (
	CapabilityPaymentsMastercard      CapabilityID = "payments_mastercard"
	CapabilityPaymentsVisa            CapabilityID = "payments_visa"
	CapabilityPaymentsAmex            CapabilityID = "payments_amex"
	CapabilityPaymentsJCB             CapabilityID = "payments_jcb"
	CapabilityPaymentsDiners          CapabilityID = "payments_diners"
	CapabilityPaymentsDiscover        CapabilityID = "payments_discover"
	CapabilityPaymentsUnionPay        CapabilityID = "payments_unionpay"
	CapabilityPaymentsKoreanLocalCard CapabilityID = "payments_korean_local_card"
)

// Payment capability constants - e-wallets.
// 支付能力常量 - 电子钱包。
const (
	CapabilityPaymentsGooglePay     CapabilityID = "payments_googlepay"
	CapabilityPaymentsApplePay      CapabilityID = "payments_applepay"
	CapabilityPaymentsWeChat        CapabilityID = "payments_wechat"
	CapabilityPaymentsAlipayCN      CapabilityID = "payments_alipaycn"
	CapabilityPaymentsAlipayHK      CapabilityID = "payments_alipayhk"
	CapabilityPaymentsDana          CapabilityID = "payments_dana"
	CapabilityPaymentsGCash         CapabilityID = "payments_gcash"
	CapabilityPaymentsKakaoPay      CapabilityID = "payments_kakaopay"
	CapabilityPaymentsGrabPayMY     CapabilityID = "payments_grabpay_my"
	CapabilityPaymentsGrabPaySG     CapabilityID = "payments_grabpay_sg"
	CapabilityPaymentsBoost         CapabilityID = "payments_boost"
	CapabilityPaymentsBlik          CapabilityID = "payments_blik"
	CapabilityPaymentsTouchNGo      CapabilityID = "payments_tng"
	CapabilityPaymentsTrueMoney     CapabilityID = "payments_truemoney"
	CapabilityPaymentsRabbitLinePay CapabilityID = "payments_rabbit_line_pay"
	CapabilityPaymentsShopeePay     CapabilityID = "payments_shopee_pay"
	CapabilityPaymentsSamsungPay    CapabilityID = "payments_samsung_pay"
	CapabilityPaymentsPayMe         CapabilityID = "payments_payme"
	CapabilityPaymentsJKOPay        CapabilityID = "payments_jkopay"
	CapabilityPaymentsLinePay       CapabilityID = "payments_line_pay"
	CapabilityPaymentsPayPay        CapabilityID = "payments_paypay"
	CapabilityPaymentsOVO           CapabilityID = "payments_ovo"
	CapabilityPaymentsLinkAja       CapabilityID = "payments_linkaja"
	CapabilityPaymentsGoPay         CapabilityID = "payments_go_pay"
	CapabilityPaymentsDOKUEWallet   CapabilityID = "payments_doku_ewallet"
	CapabilityPaymentsNaverPay      CapabilityID = "payments_naver_pay"
	CapabilityPaymentsPayco         CapabilityID = "payments_payco"
	CapabilityPaymentsSkrill        CapabilityID = "payments_skrill"
	CapabilityPaymentsPayPal        CapabilityID = "payments_paypal"
	CapabilityPaymentsVenmo         CapabilityID = "payments_venmo"
	CapabilityPaymentsCashAppPay    CapabilityID = "payments_cash_app_pay"
	CapabilityPaymentsLumi          CapabilityID = "payments_lumi"
	CapabilityPaymentsBirdPay       CapabilityID = "payments_birdpay"
)

// Payment capability constants - BNPL/installments.
// 支付能力常量 - BNPL/分期。
const (
	CapabilityPaymentsKlarna   CapabilityID = "payments_klarna"
	CapabilityPaymentsAfterpay CapabilityID = "payments_afterpay"
	CapabilityPaymentsAtome    CapabilityID = "payments_atome"
	CapabilityPaymentsZip      CapabilityID = "payments_zip"
)

// Payment capability constants - bank transfers / direct debits.
// 支付能力常量 - 银行转账/直接扣款。
const (
	CapabilityPaymentsBacsDirectDebit CapabilityID = "payments_bacs_direct_debit"
	CapabilityPaymentsACHDirectDebit  CapabilityID = "payments_ach_direct_debit"
	CapabilityPaymentsSEPADirectDebit CapabilityID = "payments_sepa_direct_debit"
	CapabilityPaymentsEFTDirectDebit  CapabilityID = "payments_eft_direct_debit"
	CapabilityPaymentsBankTransfer    CapabilityID = "payments_bank_transfer"
	CapabilityPaymentsBankTransferID  CapabilityID = "payments_bank_transfer_id"
	CapabilityPaymentsBankTransferKR  CapabilityID = "payments_bank_transfer_kr"
	CapabilityPaymentsFPS             CapabilityID = "payments_fps"
	CapabilityPaymentsPayNow          CapabilityID = "payments_pay_now"
	CapabilityPaymentsPromptPay       CapabilityID = "payments_prompt_pay"
	CapabilityPaymentsDuitNow         CapabilityID = "payments_duit_now"
	CapabilityPaymentsPayTo           CapabilityID = "payments_payto"
	CapabilityPaymentsPIX             CapabilityID = "payments_pix"
	CapabilityPaymentsSPEI            CapabilityID = "payments_spei"
)

// Payment capability constants - convenience stores / cash payments.
// 支付能力常量 - 便利店/现金支付。
const (
	CapabilityPaymentsAlfamart        CapabilityID = "payments_alfamart"
	CapabilityPaymentsIndomaret       CapabilityID = "payments_indomaret"
	CapabilityPaymentsFamilyMart      CapabilityID = "payments_family_mart"
	CapabilityPaymentsHiLife          CapabilityID = "payments_hi_life"
	CapabilityPaymentsSevenElevenMY   CapabilityID = "payments_seven_eleven_my"
	CapabilityPaymentsSevenElevenTW   CapabilityID = "payments_seven_eleven_tw"
	CapabilityPaymentsSevenElevenJP   CapabilityID = "payments_seven_eleven_jp"
	CapabilityPaymentsSamKiosk        CapabilityID = "payments_sam_kiosk"
	CapabilityPaymentsAXSKiosk        CapabilityID = "payments_axs_kiosk"
	CapabilityPaymentsTescoLotus      CapabilityID = "payments_tesco_lotus"
	CapabilityPaymentsNarvesen        CapabilityID = "payments_narvesen"
	CapabilityPaymentsPerlasTerminals CapabilityID = "payments_perlas_terminals"
	CapabilityPaymentsPermataATM      CapabilityID = "payments_permata_atm"
	CapabilityPaymentsPayeasyATM      CapabilityID = "payments_payeasy_atm"
	CapabilityPaymentsPayPost         CapabilityID = "payments_paypost"
	CapabilityPaymentsKonbini         CapabilityID = "payments_konbini"
)

// Payment capability constants - online banking.
// 支付能力常量 - 在线银行。
const (
	CapabilityPaymentsOnlineBankingKR CapabilityID = "payments_online_banking_kr"
	CapabilityPaymentsOnlineBankingPH CapabilityID = "payments_online_banking_ph"
	CapabilityPaymentsOnlineBankingTH CapabilityID = "payments_online_banking_th"
	CapabilityPaymentsOnlineBankingEE CapabilityID = "payments_online_banking_ee"
	CapabilityPaymentsOnlineBankingLT CapabilityID = "payments_online_banking_lt"
	CapabilityPaymentsOnlineBankingLV CapabilityID = "payments_online_banking_lv"
	CapabilityPaymentsMyBank          CapabilityID = "payments_mybank"
	CapabilityPaymentsFPX             CapabilityID = "payments_fpx"
	CapabilityPaymentsTrustly         CapabilityID = "payments_trustly"
	CapabilityPaymentsPaybybankapp    CapabilityID = "payments_paybybankapp"
)

// Payment capability constants - other payment methods.
// 支付能力常量 - 其他支付方式。
const (
	CapabilityPaymentsAirwallexPay CapabilityID = "payments_airwallex_pay"
	CapabilityPaymentsGiropay      CapabilityID = "payments_giropay"
	CapabilityPaymentsiDEAL        CapabilityID = "payments_ideal"
	CapabilityPaymentsBancontact   CapabilityID = "payments_bancontact"
	CapabilityPaymentsSofort       CapabilityID = "payments_sofort"
	CapabilityPaymentsEPS          CapabilityID = "payments_eps"
	CapabilityPaymentsP24          CapabilityID = "payments_p24"
	CapabilityPaymentsDragonpay    CapabilityID = "payments_dragonpay"
	CapabilityPaymentsBitPay       CapabilityID = "payments_bitpay"
	CapabilityPaymentsEnets        CapabilityID = "payments_enets"
	CapabilityPaymentsESUN         CapabilityID = "payments_esun"
	CapabilityPaymentsJeniusPay    CapabilityID = "payments_jenius_pay"
	CapabilityPaymentsMaxima       CapabilityID = "payments_maxima"
	CapabilityPaymentsMultibanco   CapabilityID = "payments_multibanco"
	CapabilityPaymentsPayeasy      CapabilityID = "payments_payeasy"
	CapabilityPaymentsPaysafecard  CapabilityID = "payments_paysafecard"
	CapabilityPaymentsPaysafecash  CapabilityID = "payments_paysafecash"
	CapabilityPaymentsPaysera      CapabilityID = "payments_paysera"
	CapabilityPaymentsPayU         CapabilityID = "payments_payu"
	CapabilityPaymentsSatispay     CapabilityID = "payments_satispay"
	CapabilityPaymentsTossPay      CapabilityID = "payments_toss_pay"
	CapabilityPaymentsTwint        CapabilityID = "payments_twint"
	CapabilityPaymentsUPI          CapabilityID = "payments_upi"
)

// AccountCapability represents account capability information.
// AccountCapability 表示账户能力信息。
type AccountCapability struct {
	// Comment is the optional remark for the capability.
	// Comment 能力备注说明。可选。
	Comment string `json:"comment"`
	// Details is the optional capability details.
	// Details 能力详情。可选。
	Details *CapabilityDetails `json:"details,omitempty"`
	// EntityType is the optional declarant entity type.
	// EntityType 申报实体类型。可选。
	EntityType EntityType `json:"entity_type,omitempty"`
	// ID is the unique identifier of the account capability. Required.
	// ID 账户能力唯一标识符。必填。
	ID CapabilityID `json:"id"`
	// Status is the status of the account capability. Required.
	// Status 账户能力状态。必填。
	Status CapabilityStatus `json:"status"`
	// UpdatedAt is the last update time. Required.
	// UpdatedAt 更新时间。必填。
	UpdatedAt string `json:"updated_at"`
}

// CapabilityDetails represents capability details.
// CapabilityDetails 能力详情。
type CapabilityDetails struct {
	// ReasonCodes is the list of reason codes. Optional.
	// ReasonCodes 原因代码列表。可选。
	ReasonCodes []string `json:"reason_codes,omitempty"`
}

// BankDetails represents bank account details.
// BankDetails 银行详情。
type BankDetails struct {
	// AccountNumber is the bank account number. Required.
	// AccountNumber 银行账号。必填。
	AccountNumber string `json:"account_number"`
	// BankName is the bank name. Required.
	// BankName 银行名称。必填。
	BankName string `json:"bank_name"`
	// CNAPSCode is the China National Advanced Payment System code. Optional.
	// CNAPSCode 中国现代化支付系统代码。可选。
	CNAPSCode string `json:"cnaps_code,omitempty"`
	// MobileNumber is the mobile phone number. Optional.
	// MobileNumber 手机号码。可选。
	MobileNumber string `json:"mobile_number,omitempty"`
}

// BusinessDeclarantInfo represents business declarant information.
// BusinessDeclarantInfo 企业申报人信息。
type BusinessDeclarantInfo struct {
	// BankDetails is the optional bank details.
	// BankDetails 银行详情。可选。
	BankDetails *BankDetails `json:"bank_details,omitempty"`
	// ContactEmail is the optional contact email.
	// ContactEmail 联系邮箱。可选。
	ContactEmail string `json:"contact_email,omitempty"`
	// EconomicCategoryCode is the optional economic category code.
	// EconomicCategoryCode 经济类别代码。可选。
	EconomicCategoryCode string `json:"economic_category_code,omitempty"`
	// IsSpecialEconomicZone indicates whether it is a special economic zone. Optional.
	// IsSpecialEconomicZone 是否经济特区。可选。
	IsSpecialEconomicZone *bool `json:"is_special_economic_zone,omitempty"`
	// SpecialEconomicZoneBusinessType is the optional business type in the special economic zone.
	// SpecialEconomicZoneBusinessType 经济特区业务类型。可选。
	SpecialEconomicZoneBusinessType string `json:"special_economic_zone_business_type,omitempty"`
}

// IndividualDeclarantInfo represents individual declarant information.
// IndividualDeclarantInfo 个人申报人信息。
type IndividualDeclarantInfo struct {
	// BankDetails is the optional bank details.
	// BankDetails 银行详情。可选。
	BankDetails *BankDetails `json:"bank_details,omitempty"`
	// ContactEmail is the optional contact email.
	// ContactEmail 联系邮箱。可选。
	ContactEmail string `json:"contact_email,omitempty"`
	// PersonID is the optional personal identification number.
	// PersonID 个人身份证件号码。可选。
	PersonID string `json:"person_id,omitempty"`
}

// EnableAccountCapabilityRequest represents a request to enable an account capability.
// EnableAccountCapabilityRequest 启用账户能力请求。
type EnableAccountCapabilityRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// ID is the unique account capability identifier. Required.
	// ID 账户能力唯一标识符。必填。
	ID CapabilityID `json:"id"`
	// EnrollSMEProgram indicates whether to enroll in the SME program. Optional.
	// EnrollSMEProgram 是否加入中小企业计划。可选。
	EnrollSMEProgram *bool `json:"enroll_sme_program,omitempty"`
	// EntityType is the optional declarant entity type.
	// EntityType 申报实体类型。可选。
	EntityType EntityType `json:"entity_type,omitempty"`
	// BusinessDeclarantInfo is the optional business declarant information.
	// BusinessDeclarantInfo 企业申报人信息。可选。
	BusinessDeclarantInfo *BusinessDeclarantInfo `json:"business_declarant_info,omitempty"`
	// IndividualDeclarantInfo is the optional individual declarant information.
	// IndividualDeclarantInfo 个人申报人信息。可选。
	IndividualDeclarantInfo *IndividualDeclarantInfo `json:"individual_declarant_info,omitempty"`
}

// FundingLimit represents a funding limit.
// FundingLimit 资金限额。
type FundingLimit struct {
	// Availables is the map of available amounts per currency. Optional.
	// Availables 各货币可用额度映射。可选。
	Availables map[string]float64 `json:"availables,omitempty"`
	// Currency is the limit currency code. Required.
	// Currency 限额货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// EffectiveAt is the effective time. Optional.
	// EffectiveAt 生效时间。可选。
	EffectiveAt string `json:"effective_at,omitempty"`
	// Limit is the limit amount. Required.
	// Limit 限额金额。必填。
	Limit float64 `json:"limit"`
	// MandateType is the mandate type. Optional.
	// MandateType 授权类型。可选。
	MandateType string `json:"mandate_type,omitempty"`
	// RequestedLimit is the requested limit amount. Optional.
	// RequestedLimit 请求的限额金额。可选。
	RequestedLimit float64 `json:"requested_limit,omitempty"`
	// Status is the limit status. Required.
	// Status 限额状态。必填。
	Status string `json:"status"`
	// Type is the limit type. Required.
	// Type 限额类型。必填。
	Type string `json:"type"`
	// UpdatedAt is the update time. Optional.
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
}

// FundingLimitsResponse represents the funding limits response.
// FundingLimitsResponse 资金限额响应。
type FundingLimitsResponse struct {
	// HasMore indicates whether there are more results. Optional.
	// HasMore 是否有更多结果。可选。
	HasMore bool `json:"has_more,omitempty"`
	// Items is the list of funding limits. Optional.
	// Items 资金限额列表。可选。
	Items []FundingLimit `json:"items,omitempty"`
}

// GetAccountCapability retrieves an account capability by ID.
// GetAccountCapability 根据 ID 获取账户能力。
// 官方文档: https://www.airwallex.com/docs/api/account_capability/account_capability/retrieve.md
func (s *Service) GetAccountCapability(ctx context.Context, id CapabilityID, opts ...sdk.RequestOption) (*AccountCapability, error) {
	var resp AccountCapability
	err := s.doer.Do(ctx, "GET", "/api/v1/account_capabilities/"+id, nil, &resp, opts...)
	return &resp, err
}

// EnableAccountCapability enables an account capability.
// EnableAccountCapability 启用账户能力。
// 官方文档: https://www.airwallex.com/docs/api/account_capability/account_capability/enable.md
func (s *Service) EnableAccountCapability(ctx context.Context, id CapabilityID, req *EnableAccountCapabilityRequest, opts ...sdk.RequestOption) (*AccountCapability, error) {
	var resp AccountCapability
	err := s.doer.Do(ctx, "POST", "/api/v1/account_capabilities/"+id+"/enable", req, &resp, opts...)
	return &resp, err
}

// GetFundingLimits retrieves funding limits.
// GetFundingLimits 获取资金限额。
// 官方文档: https://www.airwallex.com/docs/api/account_capability/account_capability/funding_limits.md
func (s *Service) GetFundingLimits(ctx context.Context, opts ...sdk.RequestOption) (*FundingLimitsResponse, error) {
	var resp FundingLimitsResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/account_capabilities/funding_limits", nil, &resp, opts...)
	return &resp, err
}
