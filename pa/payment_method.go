package pa

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PaymentMethodType represents a payment method type.
// PaymentMethodType 支付方式类型枚举。
type PaymentMethodType = string

const (
	// PaymentMethodTypeCard indicates card payment method.
	// PaymentMethodTypeCard 银行卡支付方式。
	PaymentMethodTypeCard PaymentMethodType = "card"
	// PaymentMethodTypeWechatPay indicates WeChat Pay payment method.
	// PaymentMethodTypeWechatPay 微信支付方式。
	PaymentMethodTypeWechatPay PaymentMethodType = "wechatpay"
	// PaymentMethodTypeAlipay indicates Alipay payment method.
	// PaymentMethodTypeAlipay 支付宝支付方式。
	PaymentMethodTypeAlipay PaymentMethodType = "alipay"
	// PaymentMethodTypeGooglePay indicates Google Pay payment method.
	// PaymentMethodTypeGooglePay Google Pay 支付方式。
	PaymentMethodTypeGooglePay PaymentMethodType = "googlepay"
	// PaymentMethodTypeApplePay indicates Apple Pay payment method.
	// PaymentMethodTypeApplePay Apple Pay 支付方式。
	PaymentMethodTypeApplePay PaymentMethodType = "applepay"
)

// PaymentMethodStatus represents a payment method status.
// PaymentMethodStatus 支付方式状态枚举。
type PaymentMethodStatus = string

const (
	// PaymentMethodStatusCreated indicates the payment method has been created.
	// PaymentMethodStatusCreated 已创建。
	PaymentMethodStatusCreated PaymentMethodStatus = "CREATED"
	// PaymentMethodStatusDisabled indicates the payment method has been disabled.
	// PaymentMethodStatusDisabled 已禁用。
	PaymentMethodStatusDisabled PaymentMethodStatus = "DISABLED"
)

// PaymentMethod represents a payment method.
// PaymentMethod 表示支付方式信息。
type PaymentMethod struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Type is the payment method type. Required.
	// Type 支付方式类型。必填。
	Type PaymentMethodType `json:"type"`
	// Status is the payment method status. Required.
	// Status 支付方式状态。必填。
	Status PaymentMethodStatus `json:"status"`
	// Card is the card details. Optional.
	// Card 银行卡详情。可选。
	Card *CardDetails `json:"card,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt is the update time. Optional.
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
	// CustomerID is the customer unique identifier. Optional.
	// CustomerID 客户唯一标识符。可选。
	CustomerID string `json:"customer_id,omitempty"`
}

// CardDetails represents card details.
// CardDetails 银行卡详情。
type CardDetails struct {
	// Brand is the card brand. Optional.
	// Brand 卡组织品牌。可选。
	Brand string `json:"brand,omitempty"`
	// Country is the issuing country code. Optional.
	// Country 发卡国家代码。可选。
	Country string `json:"country,omitempty"`
	// ExpiryMonth is the expiry month. Optional.
	// ExpiryMonth 过期月份。可选。
	ExpiryMonth string `json:"expiry_month,omitempty"`
	// ExpiryYear is the expiry year. Optional.
	// ExpiryYear 过期年份。可选。
	ExpiryYear string `json:"expiry_year,omitempty"`
	// LastFour is the last four digits of the card number. Optional.
	// LastFour 卡号后四位。可选。
	LastFour string `json:"last_four,omitempty"`
	// Name is the cardholder name. Optional.
	// Name 持卡人姓名。可选。
	Name string `json:"name,omitempty"`
	// Bin is the bank identification number (first 6 digits). Optional.
	// Bin 卡号前六位（Bank Identification Number）。可选。
	Bin string `json:"bin,omitempty"`
	// Fingerprint is the card fingerprint. Optional.
	// Fingerprint 卡片指纹。可选。
	Fingerprint string `json:"fingerprint,omitempty"`
	// IssuerName is the issuer bank name. Optional.
	// IssuerName 发卡行名称。可选。
	IssuerName string `json:"issuer_name,omitempty"`
	// IsCommercial indicates whether the card is commercial. Optional.
	// IsCommercial 是否为商业卡。可选。
	IsCommercial bool `json:"is_commercial,omitempty"`
	// NumberType is the card number type. Optional.
	// NumberType 卡号类型。可选。
	NumberType string `json:"number_type,omitempty"`
}

// PaymentMethodInput represents payment method details for Create/Confirm PaymentIntent requests.
// PaymentMethodInput 表示支付方式详情对象，用于 Create/Confirm PaymentIntent 请求。
type PaymentMethodInput struct {
	// Type is the payment method type. Required.
	// Type 支付方式类型。必填。
	Type PaymentMethodType `json:"type"`
	// Card is the card payment method details. Optional.
	// Card 银行卡支付方式详情。可选。
	Card *CardPaymentMethod `json:"card,omitempty"`
	// WechatPay is the WeChat Pay payment method details. Optional.
	// WechatPay 微信支付方式详情。可选。
	WechatPay *WechatPayPaymentMethod `json:"wechatpay,omitempty"`
	// AlipayCN is the Alipay China payment method details. Optional.
	// AlipayCN 中国大陆支付宝支付方式详情。可选。
	AlipayCN *AlipayPaymentMethod `json:"alipaycn,omitempty"`
	// AlipayHK is the Alipay Hong Kong payment method details. Optional.
	// AlipayHK 中国香港支付宝支付方式详情。可选。
	AlipayHK *AlipayPaymentMethod `json:"alipayhk,omitempty"`
	// AirwallexPay is the Airwallex Pay payment method details. Optional.
	// AirwallexPay Airwallex Pay 支付方式详情。可选。
	AirwallexPay *AirwallexPayPaymentMethod `json:"airwallex_pay,omitempty"`
	// KakaoPay is the Kakao Pay payment method details. Optional.
	// KakaoPay Kakao Pay 支付方式详情。可选。
	KakaoPay *KakaoPayPaymentMethod `json:"kakaopay,omitempty"`
	// Visa is the Visa payment method details. Optional.
	// Visa Visa 支付方式详情。可选。
	Visa *VisaPaymentMethod `json:"visa,omitempty"`
}

// CardPaymentMethod represents card payment method details.
// CardPaymentMethod 表示银行卡支付方式详情。
type CardPaymentMethod struct {
	// Number is the card number. Required.
	// Number 银行卡号。必填。
	Number string `json:"number"`
	// ExpiryMonth 过期月份。必填。
	ExpiryMonth string `json:"expiry_month"`
	// ExpiryYear 过期年份。必填。
	ExpiryYear string `json:"expiry_year"`
	// CVC is the card security code. Optional.
	// CVC 安全码。可选。
	CVC string `json:"cvc,omitempty"`
	// Name is the cardholder name. Optional.
	// Name 持卡人姓名。可选。
	Name string `json:"name,omitempty"`
}

// WechatPayPaymentMethod represents WeChat Pay payment method details.
// WechatPayPaymentMethod 表示微信支付方式详情。
type WechatPayPaymentMethod struct {
	// Flow 支付流程类型。必填。
	Flow string `json:"flow"`
	// Channel is the payment channel. Optional.
	// Channel 支付渠道。可选。
	Channel string `json:"channel,omitempty"`
	// OpenID is the user OpenID. Optional.
	// OpenID 用户 OpenID。可选。
	OpenID string `json:"open_id,omitempty"`
}

// AlipayPaymentMethod represents Alipay payment method details.
// AlipayPaymentMethod 表示支付宝支付方式详情。
type AlipayPaymentMethod struct {
	// Flow is the payment flow type. Required.
	// Flow 支付流程类型。必填。
	Flow string `json:"flow"`
	// OSType is the operating system type. Optional.
	// OSType 操作系统类型。可选。
	OSType string `json:"os_type,omitempty"`
}

// AirwallexPayPaymentMethod represents Airwallex Pay payment method details.
// AirwallexPayPaymentMethod 表示 Airwallex Pay 支付方式详情。
type AirwallexPayPaymentMethod struct {
	// PayerName is the payer name. Optional.
	// PayerName 付款人姓名。可选。
	PayerName string `json:"payer_name,omitempty"`
}

// KakaoPayPaymentMethod represents Kakao Pay payment method details.
// KakaoPayPaymentMethod 表示 Kakao Pay 支付方式详情。
type KakaoPayPaymentMethod struct {
	// Flow is the payment flow type. Optional.
	// Flow 支付流程类型。可选。
	Flow string `json:"flow,omitempty"`
}

// VisaPaymentMethod represents Visa payment method details.
// VisaPaymentMethod 表示 Visa 支付方式详情。
type VisaPaymentMethod struct {
	// 当前保留为空结构体，如需扩展可添加相关字段
}

// 注：以下支付方式尚未实现结构体定义，如有需要请参照上述模式扩展：
// applepay, googlepay, ideal, tng, truemoney, dana, gcash, grabpay_my, grabpay_sg,
// pay_now, fps, fpx, giropay, eps, p24, multibanco, mybank,
// bancontact, blik, trustly, twint, unionpay, ach_direct_debit,
// bacs_direct_debit, becs_direct_debit, sepa_direct_debit,
// eft_direct_debit, afterpay, klarna, affirm, atome, zip,
// tabby, spaylater, venmo, paysafecard, bitpay, ovo, go_pay,
// linkaja, jenius_pay, jkopay, payco, samsung_pay, naver_pay,
// toss_pay, korean_local_card, rabbit_line_pay, family_mart,
// hi_life, shopee_pay, alfamart, indomaret, konbini,
// online_banking, doku_ewallet, easypaisa, paypost,
// wechatpay_score, amex, paypal, enets, esun, axs_kiosk,
// easy_atm, perlas_terminals, afecash, narvesen, verkkopankki,
// bybankapp, boleto, pix, spei, oxxo, rapipago, pagofacil,
// webpay, multicaja, servipag, khipu, efecty, pse, nequi,
// addi, nupay, mercadopago, redcompra, bank_transfer, wero, etc.

// CreatePaymentMethodRequest is the request to create a payment method.
// CreatePaymentMethodRequest 创建支付方式请求。
type CreatePaymentMethodRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Type 支付方式类型。必填。可选值: card, applepay, googlepay。
	Type PaymentMethodType `json:"type"`
	// CustomerID 客户唯一标识符。必填。
	CustomerID string `json:"customer_id"`
	// Card 银行卡详情。当 type=card 时必填。
	Card *CreateCardPaymentMethod `json:"card,omitempty"`
	// Metadata 附加元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// CreateCardPaymentMethod represents card details when creating a payment method.
// CreateCardPaymentMethod 创建支付方式时的银行卡详情。
type CreateCardPaymentMethod struct {
	// Number 卡号。必填。
	Number string `json:"number"`
	// ExpiryMonth 过期月份。必填。
	ExpiryMonth string `json:"expiry_month"`
	// ExpiryYear 过期年份。必填。
	ExpiryYear string `json:"expiry_year"`
	// Name 持卡人姓名。可选。
	Name string `json:"name,omitempty"`
	// NumberType 卡号类型。可选。PAN, EXTERNAL_NETWORK_TOKEN。
	NumberType string `json:"number_type,omitempty"`
	// Billing 账单信息。可选。
	Billing *BillingInfo `json:"billing,omitempty"`
	// AdditionalInfo 附加信息。可选。
	AdditionalInfo *CardAdditionalInfo `json:"additional_info,omitempty"`
}

// BillingInfo represents billing information.
// BillingInfo 账单信息。
type BillingInfo struct {
	// FirstName 名。可选。
	FirstName string `json:"first_name,omitempty"`
	// LastName 姓。可选。
	LastName string `json:"last_name,omitempty"`
	// Email 邮箱。可选。
	Email string `json:"email,omitempty"`
	// PhoneNumber 电话号码。可选。
	PhoneNumber string `json:"phone_number,omitempty"`
	// Address 地址。可选。
	Address *BillingAddress `json:"address,omitempty"`
}

// BillingAddress represents a billing address.
// BillingAddress 账单地址。
type BillingAddress struct {
	// City 城市。可选。
	City string `json:"city,omitempty"`
	// CountryCode 国家代码（ISO 3166-2）。必填。
	CountryCode string `json:"country_code"`
	// State 州/省。可选。
	State string `json:"state,omitempty"`
	// Street 街道地址。可选。
	Street string `json:"street,omitempty"`
	// Postcode 邮政编码。可选。
	Postcode string `json:"postcode,omitempty"`
}

// CardAdditionalInfo represents additional card information.
// CardAdditionalInfo 银行卡附加信息。
type CardAdditionalInfo struct {
	// MerchantVerificationValue 商户验证值。可选。
	MerchantVerificationValue string `json:"merchant_verification_value,omitempty"`
	// TokenRequestorID Token 请求者 ID。可选。
	TokenRequestorID string `json:"token_requestor_id,omitempty"`
}

// DisablePaymentMethodRequest is the request to disable a payment method.
// DisablePaymentMethodRequest 禁用支付方式请求。
type DisablePaymentMethodRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
}

// UpdatePaymentMethodRequest is the request to update a payment method.
// UpdatePaymentMethodRequest 更新支付方式请求。
type UpdatePaymentMethodRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Metadata 附加元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// Billing 账单信息。可选。
	Billing *BillingInfo `json:"billing,omitempty"`
}

// CreatePaymentMethod creates a payment method.
// CreatePaymentMethod 创建支付方式。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_methods/create.md
func (s *Service) CreatePaymentMethod(ctx context.Context, req *CreatePaymentMethodRequest, opts ...sdk.RequestOption) (*PaymentMethod, error) {
	var resp PaymentMethod
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_methods/create", req, &resp, opts...)
	return &resp, err
}

// GetPaymentMethod retrieves a payment method by ID.
// GetPaymentMethod 根据 ID 获取支付方式。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_methods/retrieve.md
func (s *Service) GetPaymentMethod(ctx context.Context, id string, opts ...sdk.RequestOption) (*PaymentMethod, error) {
	var resp PaymentMethod
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_methods/"+id, nil, &resp, opts...)
	return &resp, err
}

// DisablePaymentMethod disables a payment method.
// DisablePaymentMethod 禁用支付方式。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_methods/disable.md
func (s *Service) DisablePaymentMethod(ctx context.Context, id string, req *DisablePaymentMethodRequest, opts ...sdk.RequestOption) (*PaymentMethod, error) {
	var resp PaymentMethod
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_methods/"+id+"/disable", req, &resp, opts...)
	return &resp, err
}

// UpdatePaymentMethod updates a payment method.
// UpdatePaymentMethod 更新支付方式。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_methods/update.md
func (s *Service) UpdatePaymentMethod(ctx context.Context, id string, req *UpdatePaymentMethodRequest, opts ...sdk.RequestOption) (*PaymentMethod, error) {
	var resp PaymentMethod
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_methods/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ListPaymentMethods lists payment methods.
// ListPaymentMethods 列出支付方式。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_methods/list.md
func (s *Service) ListPaymentMethods(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[PaymentMethod], error) {
	var resp sdk.ListResult[PaymentMethod]
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_methods", nil, &resp, opts...)
	return &resp, err
}
