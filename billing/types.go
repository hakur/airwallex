package billing

// Address represents a physical address.
// Address 地址。
type Address struct {
	// City 城市。可选。
	City string `json:"city,omitempty"`
	// CountryCode 国家代码。必填。
	CountryCode string `json:"country_code"`
	// Postcode 邮政编码。可选。
	Postcode string `json:"postcode,omitempty"`
	// State 州/省。可选。
	State string `json:"state,omitempty"`
	// Street 街道。可选。
	Street string `json:"street,omitempty"`
}

// Recurring represents recurring configuration.
// Recurring 周期性配置。
type Recurring struct {
	// Period 周期数。必填。
	Period int32 `json:"period"`
	// PeriodUnit 周期单位。必填。
	PeriodUnit PeriodUnit `json:"period_unit"`
}

// Duration represents a time duration.
// Duration 持续时间。
type Duration struct {
	// Period is the number of periods. Required.
	// Period 周期数。必填。
	Period int32 `json:"period"`
	// PeriodUnit 周期单位。必填。
	PeriodUnit PeriodUnit `json:"period_unit"`
}

// PaymentOptions represents payment configuration options.
// PaymentOptions 支付选项。
type PaymentOptions struct {
	// PaymentMethodTypes 支付方式类型列表。可选。
	PaymentMethodTypes []string `json:"payment_method_types,omitempty"`
}

// CouponRef represents a coupon reference.
// CouponRef 优惠券引用。
type CouponRef struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
}

// DiscountInput represents a discount input for requests.
// DiscountInput 折扣请求输入。
type DiscountInput struct {
	// Coupon 优惠券引用。必填。
	Coupon *CouponRef `json:"coupon"`
	// Type 折扣类型。必填。
	Type DiscountType `json:"type"`
}

// Discount represents a discount in responses.
// Discount 折扣（响应）。
type Discount struct {
	// AmountOff 固定金额减免。可选。
	AmountOff float64 `json:"amount_off,omitempty"`
	// AppliedTo 应用对象。必填。
	AppliedTo DiscountAppliedTo `json:"applied_to"`
	// AppliedToID 应用对象唯一标识符。必填。
	AppliedToID string `json:"applied_to_id"`
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// CheckoutID 结账唯一标识符。可选。
	CheckoutID string `json:"checkout_id,omitempty"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// Currency 货币代码。可选。
	Currency string `json:"currency,omitempty"`
	// DiscountModel 折扣模型。必填。
	DiscountModel DiscountModel `json:"discount_model"`
	// Duration 持续时间。可选。
	Duration *Duration `json:"duration,omitempty"`
	// DurationType 持续时间类型。必填。
	DurationType DiscountDurationType `json:"duration_type"`
	// EndsAt 结束时间。可选。
	EndsAt string `json:"ends_at,omitempty"`
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Name 名称。必填。
	Name string `json:"name"`
	// PercentageOff 百分比减免。可选。
	PercentageOff float64 `json:"percentage_off,omitempty"`
	// Source 折扣来源。必填。
	Source DiscountSource `json:"source"`
	// SourceID 来源唯一标识符。可选。
	SourceID string `json:"source_id,omitempty"`
	// StartsAt 开始时间。必填。
	StartsAt string `json:"starts_at"`
	// UpdatedAt 更新时间。必填。
	UpdatedAt string `json:"updated_at"`
}

// DiscountAmount represents a discount amount.
// DiscountAmount 折扣金额。
type DiscountAmount struct {
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// AppliedDiscountID 已应用折扣唯一标识符。必填。
	AppliedDiscountID string `json:"applied_discount_id"`
}

// DefaultInvoiceTemplate represents a default invoice template.
// DefaultInvoiceTemplate 默认发票模板。
type DefaultInvoiceTemplate struct {
	// InvoiceMemo 发票备注。可选。
	InvoiceMemo string `json:"invoice_memo,omitempty"`
}

// PriceTier represents a pricing tier.
// PriceTier 价格层级。
type PriceTier struct {
	// FlatAmount 固定金额。可选。
	FlatAmount float64 `json:"flat_amount,omitempty"`
	// UnitAmount 单价。可选。
	UnitAmount float64 `json:"unit_amount,omitempty"`
	// UpperBound 上限。可选。
	UpperBound float64 `json:"upper_bound,omitempty"`
}

// PageResponse represents a pagination response.
// PageResponse 分页响应。
type PageResponse struct {
	// PageAfter 下一页游标。可选。
	PageAfter string `json:"page_after,omitempty"`
	// PageBefore 上一页游标。可选。
	PageBefore string `json:"page_before,omitempty"`
}

// ListResult is a generic list response wrapper for the Billing domain.
// Uses cursor-based pagination (page_before / page_after).
// ListResult 是 Billing 领域的泛型列表响应包装器。
// 适用于采用 cursor 分页（page_before / page_after）的 Billing API。
type ListResult[T any] struct {
	Items      []T    `json:"items"`
	PageBefore string `json:"page_before,omitempty"`
	PageAfter  string `json:"page_after,omitempty"`
	TotalCount int    `json:"total_count"`
}

// CustomerType represents a billing customer type.
// CustomerType 客户类型。
type CustomerType = string

const (
	// CustomerTypeBusiness 企业客户。
	CustomerTypeBusiness CustomerType = "BUSINESS"
	// CustomerTypeIndividual 个人客户。
	CustomerTypeIndividual CustomerType = "INDIVIDUAL"
)

// CollectionMethod represents a collection method for payments.
// CollectionMethod 收款方式。
type CollectionMethod = string

const (
	// CollectionMethodAutoCharge 自动扣款。
	CollectionMethodAutoCharge CollectionMethod = "AUTO_CHARGE"
	// CollectionMethodChargeOnCheckout 结账时扣款。
	CollectionMethodChargeOnCheckout CollectionMethod = "CHARGE_ON_CHECKOUT"
	// CollectionMethodOutOfBand 外部扣款。
	CollectionMethodOutOfBand CollectionMethod = "OUT_OF_BAND"
)

// BillingType represents a billing type (in advance or in arrears).
// BillingType 账单类型。
type BillingType = string

const (
	// BillingTypeInAdvance 预付费。
	BillingTypeInAdvance BillingType = "IN_ADVANCE"
	// BillingTypeInArrears 后付费。
	BillingTypeInArrears BillingType = "IN_ARREARS"
)

// PricingModel represents a pricing model.
// PricingModel 定价模型。
type PricingModel = string

const (
	// PricingModelFlat 固定价格。
	PricingModelFlat PricingModel = "FLAT"
	// PricingModelPerUnit 按单位计费。
	PricingModelPerUnit PricingModel = "PER_UNIT"
	// PricingModelVolume 阶梯价格。
	PricingModelVolume PricingModel = "VOLUME"
	// PricingModelGraduated 梯度价格。
	PricingModelGraduated PricingModel = "GRADUATED"
)

// PriceType represents a price type (one-off or recurring).
// PriceType 价格类型。
type PriceType = string

const (
	// PriceTypeOneOff 一次性。
	PriceTypeOneOff PriceType = "ONE_OFF"
	// PriceTypeRecurring 周期性。
	PriceTypeRecurring PriceType = "RECURRING"
)

// PeriodUnit represents a period unit type.
// PeriodUnit 周期单位。
type PeriodUnit = string

const (
	// PeriodUnitDay represents day unit.
	// PeriodUnitDay 天。
	PeriodUnitDay PeriodUnit = "DAY"
	// PeriodUnitWeek represents week unit.
	// PeriodUnitWeek 周。
	PeriodUnitWeek PeriodUnit = "WEEK"
	// PeriodUnitMonth represents month unit.
	// PeriodUnitMonth 月。
	PeriodUnitMonth PeriodUnit = "MONTH"
	// PeriodUnitYear represents year unit.
	// PeriodUnitYear 年。
	PeriodUnitYear PeriodUnit = "YEAR"
)

// SubscriptionStatus represents a subscription status.
// SubscriptionStatus 订阅状态。
type SubscriptionStatus = string

const (
	// SubscriptionStatusPending 待处理。
	SubscriptionStatusPending SubscriptionStatus = "PENDING"
	// SubscriptionStatusInTrial 试用中。
	SubscriptionStatusInTrial SubscriptionStatus = "IN_TRIAL"
	// SubscriptionStatusActive 活跃。
	SubscriptionStatusActive SubscriptionStatus = "ACTIVE"
	// SubscriptionStatusUnpaid 未支付。
	SubscriptionStatusUnpaid SubscriptionStatus = "UNPAID"
	// SubscriptionStatusCancelled 已取消。
	SubscriptionStatusCancelled SubscriptionStatus = "CANCELLED"
)

// InvoiceStatus represents an invoice status.
// InvoiceStatus 发票状态。
type InvoiceStatus = string

const (
	// InvoiceStatusDraft 草稿。
	InvoiceStatusDraft InvoiceStatus = "DRAFT"
	// InvoiceStatusFinalized 已完成。
	InvoiceStatusFinalized InvoiceStatus = "FINALIZED"
	// InvoiceStatusVoided 已作废。
	InvoiceStatusVoided InvoiceStatus = "VOIDED"
)

// PaymentStatus represents a payment status.
// PaymentStatus 支付状态。
type PaymentStatus = string

const (
	// PaymentStatusUnpaid 未支付。
	PaymentStatusUnpaid PaymentStatus = "UNPAID"
	// PaymentStatusPaid 已支付。
	PaymentStatusPaid PaymentStatus = "PAID"
)

// DiscountModel represents a discount model type.
// DiscountModel 折扣模型。
type DiscountModel = string

const (
	// DiscountModelFlat 固定金额。
	DiscountModelFlat DiscountModel = "FLAT"
	// DiscountModelPercentage 百分比。
	DiscountModelPercentage DiscountModel = "PERCENTAGE"
)

// DiscountAppliedTo represents the target type of a discount.
// DiscountAppliedTo 折扣应用对象。
type DiscountAppliedTo = string

const (
	// DiscountAppliedToSubscription 订阅。
	DiscountAppliedToSubscription DiscountAppliedTo = "SUBSCRIPTION"
	// DiscountAppliedToSubscriptionItem 订阅项。
	DiscountAppliedToSubscriptionItem DiscountAppliedTo = "SUBSCRIPTION_ITEM"
	// DiscountAppliedToInvoice 发票。
	DiscountAppliedToInvoice DiscountAppliedTo = "INVOICE"
	// DiscountAppliedToInvoiceLineItem 发票行项。
	DiscountAppliedToInvoiceLineItem DiscountAppliedTo = "INVOICE_LINE_ITEM"
)

// DiscountDurationType represents the duration type of a discount.
// DiscountDurationType 折扣持续时间类型。
type DiscountDurationType = string

const (
	// DiscountDurationTypeOnce 一次。
	DiscountDurationTypeOnce DiscountDurationType = "ONCE"
	// DiscountDurationTypeCustom 自定义。
	DiscountDurationTypeCustom DiscountDurationType = "CUSTOM"
	// DiscountDurationTypeIndefinitely 无限期。
	DiscountDurationTypeIndefinitely DiscountDurationType = "INDEFINITELY"
)

// DiscountSource represents a discount source.
// DiscountSource 折扣来源。
type DiscountSource = string

const (
	// DiscountSourceCoupon 优惠券。
	DiscountSourceCoupon DiscountSource = "COUPON"
)

// DiscountType represents the acquisition method of a discount.
// DiscountType 折扣获取方式。
type DiscountType = string

const (
	// DiscountTypeCoupon 优惠券。
	DiscountTypeCoupon DiscountType = "COUPON"
)

// ProrationBehavior represents the proration behavior for refunds.
// ProrationBehavior 按比例退款行为。
type ProrationBehavior = string

const (
	// ProrationBehaviorAll 全部。
	ProrationBehaviorAll ProrationBehavior = "ALL"
	// ProrationBehaviorProrated 按比例。
	ProrationBehaviorProrated ProrationBehavior = "PRORATED"
	// ProrationBehaviorNone 无。
	ProrationBehaviorNone ProrationBehavior = "NONE"
)

// BillingAction represents the billing action when updating a subscription.
// BillingAction 更新订阅时的账单处理方式。
type BillingAction = string

const (
	// BillingActionDeferChargeAndKeepCycle defers charge and keeps the billing cycle.
	// BillingActionDeferChargeAndKeepCycle 延期扣款并保持周期。
	BillingActionDeferChargeAndKeepCycle BillingAction = "DEFER_CHARGE_AND_KEEP_CYCLE"
	// BillingActionImmediateChargeAndKeepCycle 立即扣款并保持周期。
	BillingActionImmediateChargeAndKeepCycle BillingAction = "IMMEDIATE_CHARGE_AND_KEEP_CYCLE"
	// BillingActionImmediateChargeAndResetCycle 立即扣款并重置周期。
	BillingActionImmediateChargeAndResetCycle BillingAction = "IMMEDIATE_CHARGE_AND_RESET_CYCLE"
)

// CheckoutMode represents a checkout mode.
// CheckoutMode 结账模式。
type CheckoutMode = string

const (
	// CheckoutModePayment 支付。
	CheckoutModePayment CheckoutMode = "PAYMENT"
	// CheckoutModeSubscription 订阅。
	CheckoutModeSubscription CheckoutMode = "SUBSCRIPTION"
	// CheckoutModeSetup 设置。
	CheckoutModeSetup CheckoutMode = "SETUP"
)

// CheckoutStatus represents a checkout status.
// CheckoutStatus 结账状态。
type CheckoutStatus = string

const (
	// CheckoutStatusActive 活跃。
	CheckoutStatusActive CheckoutStatus = "ACTIVE"
	// CheckoutStatusCompleted 已完成。
	CheckoutStatusCompleted CheckoutStatus = "COMPLETED"
	// CheckoutStatusCancelled 已取消。
	CheckoutStatusCancelled CheckoutStatus = "CANCELLED"
	// CheckoutStatusExpired 已过期。
	CheckoutStatusExpired CheckoutStatus = "EXPIRED"
)

// CreditNoteType represents a credit note type.
// CreditNoteType 贷项通知单类型。
type CreditNoteType = string

const (
	// CreditNoteTypeBeforePayment 支付前。
	CreditNoteTypeBeforePayment CreditNoteType = "BEFORE_PAYMENT"
	// CreditNoteTypeAfterPayment 支付后。
	CreditNoteTypeAfterPayment CreditNoteType = "AFTER_PAYMENT"
)

// CreditNoteReason represents a reason for a credit note.
// CreditNoteReason 贷项通知单原因。
type CreditNoteReason = string

const (
	// CreditNoteReasonProductReturn 产品退货。
	CreditNoteReasonProductReturn CreditNoteReason = "PRODUCT_RETURN"
	// CreditNoteReasonOrderChange 订单变更。
	CreditNoteReasonOrderChange CreditNoteReason = "ORDER_CHANGE"
	// CreditNoteReasonProductOrServiceUnsatisfactory 产品或服务不满意。
	CreditNoteReasonProductOrServiceUnsatisfactory CreditNoteReason = "PRODUCT_OR_SERVICE_UNSATISFACTORY"
	// CreditNoteReasonBillingError 账单错误。
	CreditNoteReasonBillingError CreditNoteReason = "BILLING_ERROR"
	// CreditNoteReasonGoodwillGesture 善意姿态。
	CreditNoteReasonGoodwillGesture CreditNoteReason = "GOODWILL_GESTURE"
	// CreditNoteReasonOther 其他。
	CreditNoteReasonOther CreditNoteReason = "OTHER"
)

// CreditNoteStatus represents a credit note status.
// CreditNoteStatus 贷项通知单状态。
type CreditNoteStatus = string

const (
	// CreditNoteStatusDraft 草稿。
	CreditNoteStatusDraft CreditNoteStatus = "DRAFT"
	// CreditNoteStatusFinalized 已完成。
	CreditNoteStatusFinalized CreditNoteStatus = "FINALIZED"
	// CreditNoteStatusVoided 已作废。
	CreditNoteStatusVoided CreditNoteStatus = "VOIDED"
)

// CreditNoteRefundStatus represents a credit note refund status.
// CreditNoteRefundStatus 贷项通知单退款状态。
type CreditNoteRefundStatus = string

const (
	// CreditNoteRefundStatusCreated 已创建。
	CreditNoteRefundStatusCreated CreditNoteRefundStatus = "CREATED"
	// CreditNoteRefundStatusSucceeded 已成功。
	CreditNoteRefundStatusSucceeded CreditNoteRefundStatus = "SUCCEEDED"
	// CreditNoteRefundStatusFailed 失败。
	CreditNoteRefundStatusFailed CreditNoteRefundStatus = "FAILED"
)

// MeterAggregationMethod represents a meter aggregation method.
// MeterAggregationMethod 计量器聚合方法。
type MeterAggregationMethod = string

const (
	// MeterAggregationMethodCount 计数。
	MeterAggregationMethodCount MeterAggregationMethod = "COUNT"
	// MeterAggregationMethodLast 最新值。
	MeterAggregationMethodLast MeterAggregationMethod = "LAST"
	// MeterAggregationMethodMax 最大值。
	MeterAggregationMethodMax MeterAggregationMethod = "MAX"
	// MeterAggregationMethodSum 求和。
	MeterAggregationMethodSum MeterAggregationMethod = "SUM"
	// MeterAggregationMethodUnique 唯一值。
	MeterAggregationMethodUnique MeterAggregationMethod = "UNIQUE"
)

// BillingTransactionStatus represents a billing transaction status.
// BillingTransactionStatus 交易状态。
type BillingTransactionStatus = string

const (
	// BillingTransactionStatusCreated 已创建。
	BillingTransactionStatusCreated BillingTransactionStatus = "CREATED"
	// BillingTransactionStatusSucceeded 已成功。
	BillingTransactionStatusSucceeded BillingTransactionStatus = "SUCCEEDED"
	// BillingTransactionStatusCancelled 已取消。
	BillingTransactionStatusCancelled BillingTransactionStatus = "CANCELLED"
	// BillingTransactionStatusFailed 失败。
	BillingTransactionStatusFailed BillingTransactionStatus = "FAILED"
)

// BillingTransactionType represents a billing transaction type.
// BillingTransactionType 交易类型。
type BillingTransactionType = string

const (
	// BillingTransactionTypePayment 支付。
	BillingTransactionTypePayment BillingTransactionType = "PAYMENT"
	// BillingTransactionTypeRefund 退款。
	BillingTransactionTypeRefund BillingTransactionType = "REFUND"
)

// listParams 用于构建 query string 的通用接口。
