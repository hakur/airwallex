package pa

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PaymentIntentStatus represents a payment intent status.
// PaymentIntentStatus 支付意图状态枚举。
type PaymentIntentStatus = string

const (
	// PaymentIntentStatusRequiresPaymentMethod indicates a payment method is required.
	// PaymentIntentStatusRequiresPaymentMethod 需要支付方式。
	PaymentIntentStatusRequiresPaymentMethod PaymentIntentStatus = "REQUIRES_PAYMENT_METHOD"
	// PaymentIntentStatusRequiresCustomerAction indicates customer action is required.
	// PaymentIntentStatusRequiresCustomerAction 需要客户操作。
	PaymentIntentStatusRequiresCustomerAction PaymentIntentStatus = "REQUIRES_CUSTOMER_ACTION"
	// PaymentIntentStatusRequiresCapture indicates capture is required.
	// PaymentIntentStatusRequiresCapture 需要捕获。
	PaymentIntentStatusRequiresCapture PaymentIntentStatus = "REQUIRES_CAPTURE"
	// PaymentIntentStatusProcessing indicates the payment is processing.
	// PaymentIntentStatusProcessing 处理中。
	PaymentIntentStatusProcessing PaymentIntentStatus = "PROCESSING"
	// PaymentIntentStatusSucceeded indicates the payment succeeded.
	// PaymentIntentStatusSucceeded 成功。
	PaymentIntentStatusSucceeded PaymentIntentStatus = "SUCCEEDED"
	// PaymentIntentStatusCancelled indicates the payment was cancelled.
	// PaymentIntentStatusCancelled 已取消。
	PaymentIntentStatusCancelled PaymentIntentStatus = "CANCELLED"
)

// PrepaymentModel represents the prepayment model for an order.
// PrepaymentModel 订单预支付模式枚举。
type PrepaymentModel = string

const (
	// PrepaymentModelFull indicates full prepayment.
	// PrepaymentModelFull 全额预付。
	PrepaymentModelFull PrepaymentModel = "FULL"
	// PrepaymentModelPartial indicates partial prepayment.
	// PrepaymentModelPartial 部分预付。
	PrepaymentModelPartial PrepaymentModel = "PARTIAL"
)

// PaymentConsentTrigger represents who triggers subsequent payments.
// PaymentConsentTrigger 后续支付触发方枚举。
type PaymentConsentTrigger = string

const (
	// PaymentConsentTriggerMerchant indicates merchant triggers subsequent payments.
	// PaymentConsentTriggerMerchant 商户触发后续支付。
	PaymentConsentTriggerMerchant PaymentConsentTrigger = "merchant"
	// PaymentConsentTriggerCustomer indicates customer triggers subsequent payments.
	// PaymentConsentTriggerCustomer 客户触发后续支付。
	PaymentConsentTriggerCustomer PaymentConsentTrigger = "customer"
)

// PaymentAmountType represents the payment amount agreement type.
// PaymentAmountType 支付金额协议类型枚举。
type PaymentAmountType = string

const (
	// PaymentAmountTypeFixed indicates fixed payment amount.
	// PaymentAmountTypeFixed 固定金额。
	PaymentAmountTypeFixed PaymentAmountType = "FIXED"
	// PaymentAmountTypeVariable indicates variable payment amount.
	// PaymentAmountTypeVariable 可变金额。
	PaymentAmountTypeVariable PaymentAmountType = "VARIABLE"
)

// PaymentPeriodUnit represents billing frequency period unit.
// PaymentPeriodUnit 计费周期单位枚举。
type PaymentPeriodUnit = string

const (
	// PaymentPeriodUnitDay indicates daily billing.
	// PaymentPeriodUnitDay 按天。
	PaymentPeriodUnitDay PaymentPeriodUnit = "DAY"
	// PaymentPeriodUnitWeek indicates weekly billing.
	// PaymentPeriodUnitWeek 按周。
	PaymentPeriodUnitWeek PaymentPeriodUnit = "WEEK"
	// PaymentPeriodUnitMonth indicates monthly billing.
	// PaymentPeriodUnitMonth 按月。
	PaymentPeriodUnitMonth PaymentPeriodUnit = "MONTH"
	// PaymentPeriodUnitYear indicates yearly billing.
	// PaymentPeriodUnitYear 按年。
	PaymentPeriodUnitYear PaymentPeriodUnit = "YEAR"
)

// MerchantTriggerReason represents the reason for merchant-triggered payments.
// MerchantTriggerReason 商户触发支付原因枚举。
type MerchantTriggerReason = string

const (
	// MerchantTriggerReasonScheduled indicates scheduled payments.
	// MerchantTriggerReasonScheduled 预定的支付。
	MerchantTriggerReasonScheduled MerchantTriggerReason = "scheduled"
	// MerchantTriggerReasonUnscheduled indicates unscheduled payments.
	// MerchantTriggerReasonUnscheduled 非预定的支付。
	MerchantTriggerReasonUnscheduled MerchantTriggerReason = "unscheduled"
	// MerchantTriggerReasonInstallments indicates installment payments.
	// MerchantTriggerReasonInstallments 分期付款。
	MerchantTriggerReasonInstallments MerchantTriggerReason = "installments"
)

// CreateAddressRequest is the request to provide address information.
// CreateAddressRequest 提供地址信息。
type CreateAddressRequest struct {
	// City is the city. Maximum 100 characters.
	// City 城市。最多100个字符。
	City string `json:"city,omitempty"`
	// CountryCode is the two-letter ISO 3166-1 alpha-2 country code. Required.
	// CountryCode 两位 ISO 3166-1 alpha-2 国家代码。必填。
	CountryCode string `json:"country_code"`
	// Postcode is the postcode. Maximum 10 characters.
	// Postcode 邮编。最多10个字符。
	Postcode string `json:"postcode,omitempty"`
	// State is the state or province. Maximum 100 characters.
	// State 州或省。最多100个字符。
	State string `json:"state,omitempty"`
	// Street is the street address. Maximum 1000 characters.
	// Street 街道地址。最多1000个字符。
	Street string `json:"street,omitempty"`
}

// CreateShippingRequest is the request to provide shipping information.
// CreateShippingRequest 提供配送信息。
type CreateShippingRequest struct {
	// PhoneNumber is the recipient phone number. Maximum 50 characters.
	// PhoneNumber 收件人电话。最多50个字符。
	PhoneNumber string `json:"phone_number,omitempty"`
	// ShippingCompany is the shipping company name. Maximum 100 characters.
	// ShippingCompany 物流公司名。最多100个字符。
	ShippingCompany string `json:"shipping_company,omitempty"`
	// ShippingDelayedAt is the delayed delivery time in ISO8601 format.
	// ShippingDelayedAt 延迟发货时间，ISO8601 格式。
	ShippingDelayedAt string `json:"shipping_delayed_at,omitempty"`
	// ShippingMethod is the shipping method. Maximum 128 characters.
	// ShippingMethod 配送方式。最多128个字符。
	ShippingMethod string `json:"shipping_method,omitempty"`
	// TrackingNumber is the tracking number. Maximum 100 characters.
	// TrackingNumber 运单号。最多100个字符。
	TrackingNumber string `json:"tracking_number,omitempty"`
	// TrackingURL is the tracking URL. Maximum 1024 characters.
	// TrackingURL 追踪链接。最多1024个字符。
	TrackingURL string `json:"tracking_url,omitempty"`
}

// CreateBillingRequest is the request to provide billing information.
// CreateBillingRequest 提供账单信息。
type CreateBillingRequest struct {
	// Address is the billing address.
	// Address 账单地址。
	Address *CreateAddressRequest `json:"address,omitempty"`
	// Email is the customer email address.
	// Email 客户邮箱。
	Email string `json:"email,omitempty"`
	// FirstName is the customer first name. Maximum 128 characters.
	// FirstName 客户名。最多128个字符。
	FirstName string `json:"first_name,omitempty"`
	// LastName is the customer last name. Maximum 128 characters.
	// LastName 客户姓。最多128个字符。
	LastName string `json:"last_name,omitempty"`
	// PhoneNumber is the customer phone number.
	// PhoneNumber 客户电话。
	PhoneNumber string `json:"phone_number,omitempty"`
}

// CreateCustomerDataRequest is the request to provide customer information inline.
// CreateCustomerDataRequest 内联提供客户信息。
type CreateCustomerDataRequest struct {
	// Address is the customer address.
	// Address 客户地址。
	Address *CreateAddressRequest `json:"address,omitempty"`
	// BusinessName is the business name of the customer.
	// BusinessName 客户企业名称。
	BusinessName string `json:"business_name,omitempty"`
	// Email is the email address.
	// Email 邮箱地址。
	Email string `json:"email,omitempty"`
	// FirstName is the first name.
	// FirstName 名。
	FirstName string `json:"first_name,omitempty"`
	// LastName is the last name.
	// LastName 姓。
	LastName string `json:"last_name,omitempty"`
	// MerchantCustomerID is the unique identifier of the customer in merchant's system. Maximum 64 characters.
	// MerchantCustomerID 商户系统中的客户唯一标识符。最多64个字符。
	MerchantCustomerID string `json:"merchant_customer_id,omitempty"`
	// PhoneNumber is the phone number.
	// PhoneNumber 电话号码。
	PhoneNumber string `json:"phone_number,omitempty"`
}

// CreateSurchargeRequest is the request to provide surcharge information.
// CreateSurchargeRequest 提供附加费信息。
type CreateSurchargeRequest struct {
	// Amount is the surcharge amount (included in total). Required.
	// Amount 附加费金额（已包含在总额中）。必填。
	Amount float64 `json:"amount"`
	// Percent is the surcharge rate in percentage (e.g. 5 = 5%).
	// Percent 附加费率百分比（如 5 = 5%）。
	Percent float64 `json:"percent,omitempty"`
}

// CreateTipRequest is the request to provide tip information.
// CreateTipRequest 提供小费信息。
type CreateTipRequest struct {
	// Amount is the tip amount (included in total). Required.
	// Amount 小费金额（已包含在总额中）。必填。
	Amount float64 `json:"amount"`
}

// CreatePaymentScheduleRequest is the request to provide payment schedule.
// CreatePaymentScheduleRequest 提供支付周期信息。
type CreatePaymentScheduleRequest struct {
	// Period is the number of period units between billing cycles.
	// Period 计费周期间的周期数。
	Period int `json:"period,omitempty"`
	// PeriodUnit is the billing frequency unit. One of DAY, WEEK, MONTH, YEAR.
	// PeriodUnit 计费频率单位。取值为 PaymentPeriodUnitDay / Week / Month / Year。
	PeriodUnit PaymentPeriodUnit `json:"period_unit,omitempty"`
}

// CreatePaymentConsentTermsRequest is the request to provide payment consent terms of use.
// CreatePaymentConsentTermsRequest 提供支付授权使用条款。
type CreatePaymentConsentTermsRequest struct {
	// BillingCycleChargeDay is the charge day per billing cycle.
	// BillingCycleChargeDay 每计费周期的扣款日。
	BillingCycleChargeDay int `json:"billing_cycle_charge_day,omitempty"`
	// EndDate is the end date to expect payment request.
	// EndDate 预期支付请求的结束日期。
	EndDate string `json:"end_date,omitempty"`
	// FirstPaymentAmount is the first payment amount.
	// FirstPaymentAmount 首次支付金额。
	FirstPaymentAmount float64 `json:"first_payment_amount,omitempty"`
	// FixedPaymentAmount is the fixed payment amount. Required if PaymentAmountType is FIXED.
	// FixedPaymentAmount 固定支付金额。PaymentAmountType 为 FIXED 时必填。
	FixedPaymentAmount float64 `json:"fixed_payment_amount,omitempty"`
	// MaxPaymentAmount is the maximum payment amount per single charge.
	// MaxPaymentAmount 单次最高支付金额。
	MaxPaymentAmount float64 `json:"max_payment_amount,omitempty"`
	// MinPaymentAmount is the minimum payment amount per single charge.
	// MinPaymentAmount 单次最低支付金额。
	MinPaymentAmount float64 `json:"min_payment_amount,omitempty"`
	// PaymentAmountType is the agreed type of amounts. Required. One of FIXED or VARIABLE.
	// PaymentAmountType 约定的金额类型。必填。取值为 PaymentAmountTypeFixed 或 PaymentAmountTypeVariable。
	PaymentAmountType PaymentAmountType `json:"payment_amount_type"`
	// PaymentCurrency is the currency of this payment.
	// PaymentCurrency 支付币种。
	PaymentCurrency string `json:"payment_currency,omitempty"`
	// PaymentSchedule is the payment schedule configuration.
	// PaymentSchedule 支付周期配置。
	PaymentSchedule *CreatePaymentScheduleRequest `json:"payment_schedule,omitempty"`
	// StartDate is the start date to expect payment request.
	// StartDate 预期支付请求的开始日期。
	StartDate string `json:"start_date,omitempty"`
	// TotalBillingCycles is the total number of billing cycles. Null means indefinite.
	// TotalBillingCycles 总计费周期数。不传表示无限期。
	TotalBillingCycles int `json:"total_billing_cycles,omitempty"`
}

// CreatePIConsentRequest is the request to provide payment consent configuration.
// CreatePIConsentRequest 提供支付授权配置。
type CreatePIConsentRequest struct {
	// MerchantTriggerReason indicates whether subsequent payments are scheduled. One of scheduled, unscheduled, installments.
	// MerchantTriggerReason 后续支付是否预定。取值为 MerchantTriggerReasonScheduled / Unscheduled / Installments。
	MerchantTriggerReason MerchantTriggerReason `json:"merchant_trigger_reason,omitempty"`
	// NextTriggeredBy is the party to trigger subsequent payments. Required. One of merchant or customer.
	// NextTriggeredBy 触发后续支付的一方。必填。取值为 PaymentConsentTriggerMerchant 或 PaymentConsentTriggerCustomer。
	NextTriggeredBy PaymentConsentTrigger `json:"next_triggered_by"`
	// TermsOfUse is the terms of use for the payment consent.
	// TermsOfUse 支付授权使用条款。
	TermsOfUse *CreatePaymentConsentTermsRequest `json:"terms_of_use,omitempty"`
}

// PaymentIntent represents a payment intent.
// PaymentIntent 表示支付意图信息。
type PaymentIntent struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// RequestID is the unique request identifier. Optional.
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// Amount is the amount. Required.
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// MerchantOrderID is the merchant order ID. Optional.
	// MerchantOrderID 商户订单号。可选。
	MerchantOrderID string `json:"merchant_order_id,omitempty"`
	// Descriptor is the transaction descriptor. Optional.
	// Descriptor 交易描述。可选。
	Descriptor string `json:"descriptor,omitempty"`
	// CustomerID is the customer unique identifier. Optional.
	// CustomerID 客户唯一标识符。可选。
	CustomerID string `json:"customer_id,omitempty"`
	// Status is the payment intent status. Required.
	// Status 支付意图状态。必填。
	Status PaymentIntentStatus `json:"status"`
	// CapturedAmount is the captured amount. Optional.
	// CapturedAmount 已捕获金额。可选。
	CapturedAmount float64 `json:"captured_amount,omitempty"`
	// OriginalAmount is the original amount. Optional.
	// OriginalAmount 原始金额。可选。
	OriginalAmount float64 `json:"original_amount,omitempty"`
	// OriginalCurrency is the original currency code. Optional.
	// OriginalCurrency 原始货币代码。可选。
	OriginalCurrency sdk.Currency `json:"original_currency,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt is the update time. Optional.
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
	// ClientSecret is the client secret. Optional.
	// ClientSecret 客户端密钥。可选。
	ClientSecret string `json:"client_secret,omitempty"`
	// NextAction is the next action information. Optional.
	// NextAction 下一步操作信息。可选。
	NextAction map[string]any `json:"next_action,omitempty"`
	// PaymentMethod is the payment method information. Optional.
	// PaymentMethod 支付方式信息。可选。
	PaymentMethod map[string]any `json:"payment_method,omitempty"`
	// LatestPaymentAttempt is the latest payment attempt information. Optional.
	// LatestPaymentAttempt 最新支付尝试信息。可选。
	LatestPaymentAttempt map[string]any `json:"latest_payment_attempt,omitempty"`
	// AdditionalInfo 附加信息（渠道特定数据统一字段）。可选。
	AdditionalInfo map[string]any `json:"additional_info,omitempty"`
	// Customer is the customer information returned by the API. Optional.
	// Customer 客户信息。可选。
	Customer *CustomerResponse `json:"customer,omitempty"`
}

// CustomerResponse represents customer information returned in PaymentIntent response.
// CustomerResponse 表示 PaymentIntent 响应中返回的客户信息。
type CustomerResponse struct {
	// Address is the customer address.
	// Address 客户地址。
	Address *CreateAddressRequest `json:"address,omitempty"`
	// BusinessName is the business name.
	// BusinessName 企业名称。
	BusinessName string `json:"business_name,omitempty"`
	// Email is the email address.
	// Email 邮箱地址。
	Email string `json:"email,omitempty"`
	// FirstName is the first name.
	// FirstName 名。
	FirstName string `json:"first_name,omitempty"`
	// LastName is the last name.
	// LastName 姓。
	LastName string `json:"last_name,omitempty"`
	// MerchantCustomerID is the unique identifier of the customer in merchant's system.
	// MerchantCustomerID 商户系统中的客户唯一标识符。
	MerchantCustomerID string `json:"merchant_customer_id,omitempty"`
	// PhoneNumber is the phone number.
	// PhoneNumber 电话号码。
	PhoneNumber string `json:"phone_number,omitempty"`
}

// CreateOrderRequest is the request to create order information for a payment intent.
// CreateOrderRequest 创建支付意图时的订单信息。
type CreateOrderRequest struct {
	// Cancellable indicates if the booking is cancellable. Required by API.
	// Cancellable 订单是否可取消。API 必填。
	Cancellable bool `json:"cancellable"`
	// CreatedAt is the time the order was made in ISO8601 format. Required by API.
	// CreatedAt 订单创建时间，ISO8601 格式（必须带时区）。API 必填。
	CreatedAt string `json:"created_at"`
	// PrepaymentModel is the model of payment. Optional. One of FULL or PARTIAL.
	// PrepaymentModel 支付模式。可选。取值为 PrepaymentModelFull 或 PrepaymentModelPartial。
	PrepaymentModel PrepaymentModel `json:"prepayment_model,omitempty"`
}

// CreatePaymentIntentRequest is the request to create a payment intent.
// CreatePaymentIntentRequest 创建支付意图请求。
type CreatePaymentIntentRequest struct {
	// RequestID is the unique request identifier. Optional.
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// MerchantOrderID is the merchant order ID. Optional.
	// MerchantOrderID 商户订单号。可选。
	MerchantOrderID string `json:"merchant_order_id,omitempty"`
	// Amount is the amount. Required.
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// OrderType is the industry category of the order. Maximum 128 characters.
	// OrderType 订单行业类别。最多128个字符。
	OrderType string `json:"type,omitempty"`
	// Descriptor is the transaction descriptor. Optional.
	// Descriptor 交易描述。可选。
	Descriptor string `json:"descriptor,omitempty"`
	// CustomerID is the customer unique identifier. Optional.
	// CustomerID 客户唯一标识符。可选。
	CustomerID string `json:"customer_id,omitempty"`
	// PaymentMethod is the payment method details. Optional.
	// PaymentMethod 支付方式详情。可选。
	PaymentMethod *PaymentMethodInput `json:"payment_method,omitempty"`
	// PaymentMethodOptions are the payment method options. Optional.
	// PaymentMethodOptions 支付方式选项。可选。
	PaymentMethodOptions map[string]any `json:"payment_method_options,omitempty"`
	// ReturnURL 返回地址。可选。
	ReturnURL string `json:"return_url,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// CaptureMethod 捕获方式。可选。
	CaptureMethod string `json:"capture_method,omitempty"`
	// SetupFuturePayment 是否设置未来支付。可选。
	SetupFuturePayment bool `json:"setup_future_payment,omitempty"`
	// NotificationURL 通知地址。可选。
	NotificationURL string `json:"notification_url,omitempty"`
	// PaymentType 支付类型。可选。
	PaymentType string `json:"payment_type,omitempty"`
	// PlatformInitiated indicates whether initiated by platform. Optional.
	// PlatformInitiated 是否平台发起。可选。
	PlatformInitiated bool `json:"platform_initiated,omitempty"`
	// Force3DS indicates whether to force 3D Secure. Optional.
	// Force3DS 是否强制 3D Secure。可选。
	Force3DS bool `json:"force_3ds,omitempty"`
	// PaymentMethodTypes are the supported payment method types. Optional.
	// PaymentMethodTypes 支持的支付方式类型列表。可选。
	PaymentMethodTypes []string `json:"payment_method_types,omitempty"`
	// PaymentConsentID is the related PaymentConsent ID. Optional.
	// PaymentConsentID 关联的 PaymentConsent ID。可选。
	PaymentConsentID string `json:"payment_consent_id,omitempty"`
	// PaymentLinkID is the related PaymentLink ID. Optional.
	// PaymentLinkID 关联的 PaymentLink ID。可选。
	PaymentLinkID string `json:"payment_link_id,omitempty"`
	// Surcharge is the surcharge information. Optional.
	// Surcharge 附加费信息。可选。
	Surcharge *CreateSurchargeRequest `json:"surcharge,omitempty"`
	// Tip is the tip information. Optional.
	// Tip 小费信息。可选。
	Tip *CreateTipRequest `json:"tip,omitempty"`
	// PaymentConsent is the payment consent configuration. Optional.
	// PaymentConsent 支付授权配置。可选。
	PaymentConsent *CreatePIConsentRequest `json:"payment_consent,omitempty"`
	// CustomerData is the customer data. Optional.
	// CustomerData 客户数据。可选。
	CustomerData *CreateCustomerDataRequest `json:"customer_data,omitempty"`
	// Order is the order information. Optional.
	// Order 订单信息。可选。
	Order *CreateOrderRequest `json:"order,omitempty"`
	// Shipping is the shipping information. Optional.
	// Shipping 配送信息。可选。
	Shipping *CreateShippingRequest `json:"shipping,omitempty"`
	// Billing is the billing information. Optional.
	// Billing 账单信息。可选。
	Billing *CreateBillingRequest `json:"billing,omitempty"`
	// AdditionalInfo is additional information (unified field for channel-specific data). Optional.
	// AdditionalInfo 附加信息（渠道特定数据统一字段）。可选。
	AdditionalInfo map[string]any `json:"additional_info,omitempty"`
}

// UpdatePaymentIntentRequest is the request to update a payment intent.
// UpdatePaymentIntentRequest 更新支付意图请求。
type UpdatePaymentIntentRequest struct {
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// Amount is the amount. Optional.
	// Amount 金额。可选。
	Amount float64 `json:"amount,omitempty"`
	// Currency is the currency code. Optional.
	// Currency 货币代码。可选。
	Currency sdk.Currency `json:"currency,omitempty"`
	// OrderType is the industry category of the order. Maximum 128 characters.
	// OrderType 订单行业类别。最多128个字符。
	OrderType string `json:"type,omitempty"`
	// Descriptor is the transaction descriptor. Optional.
	// Descriptor 交易描述。可选。
	Descriptor string `json:"descriptor,omitempty"`
	// CustomerID is the customer unique identifier. Optional.
	// CustomerID 客户唯一标识符。可选。
	CustomerID string `json:"customer_id,omitempty"`
	// PaymentMethod is the payment method details. Optional.
	// PaymentMethod 支付方式详情。可选。
	PaymentMethod *PaymentMethodInput `json:"payment_method,omitempty"`
	// PaymentMethodOptions are the payment method options. Optional.
	// PaymentMethodOptions 支付方式选项。可选。
	PaymentMethodOptions map[string]any `json:"payment_method_options,omitempty"`
	// ReturnURL is the return URL. Optional.
	// ReturnURL 返回地址。可选。
	ReturnURL string `json:"return_url,omitempty"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// CaptureMethod is the capture method. Optional.
	// CaptureMethod 捕获方式。可选。
	CaptureMethod string `json:"capture_method,omitempty"`
	// SetupFuturePayment indicates whether to set up future payment. Optional.
	// SetupFuturePayment 是否设置未来支付。可选。
	SetupFuturePayment bool `json:"setup_future_payment,omitempty"`
	// NotificationURL is the notification URL. Optional.
	// NotificationURL 通知地址。可选。
	NotificationURL string `json:"notification_url,omitempty"`
	// PaymentType is the payment type. Optional.
	// PaymentType 支付类型。可选。
	PaymentType string `json:"payment_type,omitempty"`
	// Force3DS indicates whether to force 3D Secure. Optional.
	// Force3DS 是否强制 3D Secure。可选。
	Force3DS bool `json:"force_3ds,omitempty"`
	// PaymentMethodTypes are the supported payment method types. Optional.
	// PaymentMethodTypes 支持的支付方式类型列表。可选。
	PaymentMethodTypes []string `json:"payment_method_types,omitempty"`
	// PaymentConsentID is the related PaymentConsent ID. Optional.
	// PaymentConsentID 关联的 PaymentConsent ID。可选。
	PaymentConsentID string `json:"payment_consent_id,omitempty"`
	// PaymentLinkID is the related PaymentLink ID. Optional.
	// PaymentLinkID 关联的 PaymentLink ID。可选。
	PaymentLinkID string `json:"payment_link_id,omitempty"`
	// Surcharge is the surcharge information. Optional.
	// Surcharge 附加费信息。可选。
	Surcharge *CreateSurchargeRequest `json:"surcharge,omitempty"`
	// Tip is the tip information. Optional.
	// Tip 小费信息。可选。
	Tip *CreateTipRequest `json:"tip,omitempty"`
	// PaymentConsent is the payment consent configuration. Optional.
	// PaymentConsent 支付授权配置。可选。
	PaymentConsent *CreatePIConsentRequest `json:"payment_consent,omitempty"`
	// CustomerData is the customer data. Optional.
	// CustomerData 客户数据。可选。
	CustomerData *CreateCustomerDataRequest `json:"customer_data,omitempty"`
	// Order is the order information. Optional.
	// Order 订单信息。可选。
	Order *CreateOrderRequest `json:"order,omitempty"`
	// Shipping is the shipping information. Optional.
	// Shipping 配送信息。可选。
	Shipping *CreateShippingRequest `json:"shipping,omitempty"`
	// Billing is the billing information. Optional.
	// Billing 账单信息。可选。
	Billing *CreateBillingRequest `json:"billing,omitempty"`
	// AdditionalInfo is additional information (unified field for channel-specific data). Optional.
	// AdditionalInfo 附加信息（渠道特定数据统一字段）。可选。
	AdditionalInfo map[string]any `json:"additional_info,omitempty"`
}

// ConfirmPaymentIntentRequest is the request to confirm a payment intent.
// ConfirmPaymentIntentRequest 确认支付意图请求。
type ConfirmPaymentIntentRequest struct {
	// RequestID is the unique request identifier. Optional.
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// PaymentMethod is the payment method details. Optional.
	// PaymentMethod 支付方式详情。可选。
	PaymentMethod *PaymentMethodInput `json:"payment_method,omitempty"`
	// PaymentMethodOptions are the payment method options. Optional.
	// PaymentMethodOptions 支付方式选项。可选。
	PaymentMethodOptions map[string]any `json:"payment_method_options,omitempty"`
	// ReturnURL is the return URL. Optional.
	// ReturnURL 返回地址。可选。
	ReturnURL string `json:"return_url,omitempty"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// CaptureMethod is the capture method. Optional.
	// CaptureMethod 捕获方式。可选。
	CaptureMethod string `json:"capture_method,omitempty"`
	// SetupFuturePayment indicates whether to set up future payment. Optional.
	// SetupFuturePayment 是否设置未来支付。可选。
	SetupFuturePayment bool `json:"setup_future_payment,omitempty"`
	// NotificationURL is the notification URL. Optional.
	// NotificationURL 通知地址。可选。
	NotificationURL string `json:"notification_url,omitempty"`
	// PaymentType is the payment type. Optional.
	// PaymentType 支付类型。可选。
	PaymentType string `json:"payment_type,omitempty"`
	// OrderType is the industry category of the order. Maximum 128 characters.
	// OrderType 订单行业类别。最多128个字符。
	OrderType string `json:"type,omitempty"`
	// Force3DS indicates whether to force 3D Secure. Optional.
	// Force3DS 是否强制 3D Secure。可选。
	Force3DS bool `json:"force_3ds,omitempty"`
	// PaymentConsentID is the related PaymentConsent ID. Optional.
	// PaymentConsentID 关联的 PaymentConsent ID。可选。
	PaymentConsentID string `json:"payment_consent_id,omitempty"`
	// PaymentLinkID is the related PaymentLink ID. Optional.
	// PaymentLinkID 关联的 PaymentLink ID。可选。
	PaymentLinkID string `json:"payment_link_id,omitempty"`
	// Surcharge is the surcharge information. Optional.
	// Surcharge 附加费信息。可选。
	Surcharge *CreateSurchargeRequest `json:"surcharge,omitempty"`
	// Tip is the tip information. Optional.
	// Tip 小费信息。可选。
	Tip *CreateTipRequest `json:"tip,omitempty"`
	// PaymentConsent is the payment consent configuration. Optional.
	// PaymentConsent 支付授权配置。可选。
	PaymentConsent *CreatePIConsentRequest `json:"payment_consent,omitempty"`
	// CustomerData is the customer data. Optional.
	// CustomerData 客户数据。可选。
	CustomerData *CreateCustomerDataRequest `json:"customer_data,omitempty"`
	// Order is the order information. Optional.
	// Order 订单信息。可选。
	Order *CreateOrderRequest `json:"order,omitempty"`
	// Shipping is the shipping information. Optional.
	// Shipping 配送信息。可选。
	Shipping *CreateShippingRequest `json:"shipping,omitempty"`
	// Billing is the billing information. Optional.
	// Billing 账单信息。可选。
	Billing *CreateBillingRequest `json:"billing,omitempty"`
	// AdditionalInfo is additional information (unified field for channel-specific data). Optional.
	// AdditionalInfo 附加信息（渠道特定数据统一字段）。可选。
	AdditionalInfo map[string]any `json:"additional_info,omitempty"`
}

// CapturePaymentIntentRequest is the request to capture a payment intent.
// CapturePaymentIntentRequest 捕获支付意图请求。
type CapturePaymentIntentRequest struct {
	// RequestID is the unique request identifier. Optional.
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// Amount is the capture amount. Optional.
	// Amount 捕获金额。可选。
	Amount float64 `json:"amount,omitempty"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// CreatePaymentIntent creates a payment intent.
// CreatePaymentIntent 创建支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/create.md
func (s *Service) CreatePaymentIntent(ctx context.Context, req *CreatePaymentIntentRequest, opts ...sdk.RequestOption) (*PaymentIntent, error) {
	var resp PaymentIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_intents/create", req, &resp, opts...)
	return &resp, err
}

// GetPaymentIntent retrieves a payment intent by ID.
// GetPaymentIntent 根据 ID 获取支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/retrieve.md
func (s *Service) GetPaymentIntent(ctx context.Context, id string, opts ...sdk.RequestOption) (*PaymentIntent, error) {
	var resp PaymentIntent
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_intents/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdatePaymentIntent updates a payment intent.
// UpdatePaymentIntent 更新支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/update.md
func (s *Service) UpdatePaymentIntent(ctx context.Context, id string, req *UpdatePaymentIntentRequest, opts ...sdk.RequestOption) (*PaymentIntent, error) {
	var resp PaymentIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_intents/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ConfirmPaymentIntent confirms a payment intent.
// ConfirmPaymentIntent 确认支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/confirm.md
func (s *Service) ConfirmPaymentIntent(ctx context.Context, id string, req *ConfirmPaymentIntentRequest, opts ...sdk.RequestOption) (*PaymentIntent, error) {
	var resp PaymentIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_intents/"+id+"/confirm", req, &resp, opts...)
	return &resp, err
}

// CapturePaymentIntent captures a payment intent.
// CapturePaymentIntent 捕获支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/capture.md
func (s *Service) CapturePaymentIntent(ctx context.Context, id string, req *CapturePaymentIntentRequest, opts ...sdk.RequestOption) (*PaymentIntent, error) {
	var resp PaymentIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_intents/"+id+"/capture", req, &resp, opts...)
	return &resp, err
}

// CancelPaymentIntentRequest is the request to cancel a payment intent.
// CancelPaymentIntentRequest 取消支付意图请求。
type CancelPaymentIntentRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
}

// CancelPaymentIntent cancels a payment intent.
// CancelPaymentIntent 取消支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/cancel.md
func (s *Service) CancelPaymentIntent(ctx context.Context, id string, req *CancelPaymentIntentRequest, opts ...sdk.RequestOption) (*PaymentIntent, error) {
	var resp PaymentIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_intents/"+id+"/cancel", req, &resp, opts...)
	return &resp, err
}

// ListPaymentIntents lists payment intents.
// ListPaymentIntents 列出支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/list.md
func (s *Service) ListPaymentIntents(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[PaymentIntent], error) {
	var resp sdk.ListResult[PaymentIntent]
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_intents", nil, &resp, opts...)
	return &resp, err
}
