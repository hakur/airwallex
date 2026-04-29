package billing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// CheckoutLineItemPrice 结账行项目价格。
type CheckoutLineItemPrice struct {
	// Currency 货币代码。可选。
	Currency string `json:"currency,omitempty"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// FlatAmount 固定金额。可选。
	FlatAmount float64 `json:"flat_amount,omitempty"`
	// PricingModel 定价模型。可选。
	PricingModel string `json:"pricing_model,omitempty"`
	// ProductID 产品唯一标识符。可选。
	ProductID string `json:"product_id,omitempty"`
	// UnitAmount 单价。可选。
	UnitAmount float64 `json:"unit_amount,omitempty"`
}

// CheckoutLineItemDiscount 结账行项目折扣。
type CheckoutLineItemDiscount struct {
	// CouponID 优惠券唯一标识符。可选。
	CouponID string `json:"coupon_id,omitempty"`
	// Type 类型。可选。
	Type string `json:"type,omitempty"`
}

// CheckoutLineItem 结账行项目。
type CheckoutLineItem struct {
	// Price 价格。可选。
	Price *CheckoutLineItemPrice `json:"price,omitempty"`
	// PriceID 价格唯一标识符。必填。
	PriceID string `json:"price_id"`
	// Quantity 数量。必填。
	Quantity int32 `json:"quantity"`
	// Discounts 折扣列表。可选。
	Discounts []CheckoutLineItemDiscount `json:"discounts,omitempty"`
}

// CheckoutCustomerData 客户数据。
type CheckoutCustomerData struct {
	// Email 电子邮箱。可选。
	Email string `json:"email,omitempty"`
	// Name 名称。可选。
	Name string `json:"name,omitempty"`
	// Type 类型。可选。
	Type string `json:"type,omitempty"`
}

// CheckoutCustomerDataCollection 客户数据收集配置。
type CheckoutCustomerDataCollection struct {
	// Enabled 是否启用。可选。
	Enabled bool `json:"enabled,omitempty"`
}

// CheckoutDiscountCoupon 折扣优惠券。
type CheckoutDiscountCoupon struct {
	// ID 唯一标识符。可选。
	ID string `json:"id,omitempty"`
}

// CheckoutDiscount 折扣。
type CheckoutDiscount struct {
	// Coupon 优惠券。可选。
	Coupon *CheckoutDiscountCoupon `json:"coupon,omitempty"`
	// Type 类型。必填。
	Type string `json:"type"`
}

// CheckoutInvoiceData 发票数据。
type CheckoutInvoiceData struct {
	// DaysUntilDue 到期天数。可选。
	DaysUntilDue int32 `json:"days_until_due,omitempty"`
	// DefaultTaxPercent 默认税率百分比。可选。
	DefaultTaxPercent float64 `json:"default_tax_percent,omitempty"`
	// DueAt 到期时间。可选。
	DueAt string `json:"due_at,omitempty"`
	// Memo 备注。可选。
	Memo string `json:"memo,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// CheckoutSubscriptionData 订阅数据。
type CheckoutSubscriptionData struct {
	// DaysUntilDue 到期天数。必填。
	DaysUntilDue int32 `json:"days_until_due"`
	// DefaultTaxPercent 默认税率百分比。可选。
	DefaultTaxPercent float64 `json:"default_tax_percent,omitempty"`
	// DefaultInvoiceTemplate 默认发票模板。可选。
	DefaultInvoiceTemplate *DefaultInvoiceTemplate `json:"default_invoice_template,omitempty"`
	// Duration 持续时间。可选。
	Duration *Duration `json:"duration,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// Recurring 周期性配置。可选。
	Recurring *Recurring `json:"recurring,omitempty"`
	// StartsAt 开始时间。可选。
	StartsAt string `json:"starts_at,omitempty"`
	// TrialEndsAt 试用结束时间。可选。
	TrialEndsAt string `json:"trial_ends_at,omitempty"`
}

// CheckoutPaymentMethodSave 支付方式保存。
type CheckoutPaymentMethodSave struct {
	// Mode 模式。必填。
	Mode string `json:"mode"`
	// NextTriggeredBy 下次触发方式。可选。
	NextTriggeredBy string `json:"next_triggered_by,omitempty"`
}

// CheckoutPaymentOptions 支付选项。
type CheckoutPaymentOptions struct {
	// PayByInvoiceEnabled 是否启用发票支付。可选。
	PayByInvoiceEnabled bool `json:"pay_by_invoice_enabled,omitempty"`
	// PaymentMethodTypes 支付方式类型列表。可选。
	PaymentMethodTypes []string `json:"payment_method_types,omitempty"`
	// PaymentMethodSave 支付方式保存配置。可选。
	PaymentMethodSave *CheckoutPaymentMethodSave `json:"payment_method_save,omitempty"`
}

// Checkout represents a checkout object.
// Checkout 结账对象。
type Checkout struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Mode 结账模式。必填。
	Mode CheckoutMode `json:"mode"`
	// Status 结账状态。必填。
	Status CheckoutStatus `json:"status"`
	// SuccessURL 成功跳转地址。必填。
	SuccessURL string `json:"success_url"`
	// BackURL 返回跳转地址。可选。
	BackURL string `json:"back_url,omitempty"`
	// URL 结账地址。可选。
	URL string `json:"url,omitempty"`
	// BillingCustomerID 账单客户唯一标识符。可选。
	BillingCustomerID string `json:"billing_customer_id,omitempty"`
	// LegalEntityID 法律实体唯一标识符。可选。
	LegalEntityID string `json:"legal_entity_id,omitempty"`
	// LinkedPaymentAccountID 关联支付账户唯一标识符。必填。
	LinkedPaymentAccountID string `json:"linked_payment_account_id"`
	// Currency 货币代码。可选。
	Currency string `json:"currency,omitempty"`
	// Locale 地区语言。可选。
	Locale string `json:"locale,omitempty"`
	// CustomerData 客户数据。可选。
	CustomerData *CheckoutCustomerData `json:"customer_data,omitempty"`
	// CustomerDataCollection 客户数据收集配置。可选。
	CustomerDataCollection *CheckoutCustomerDataCollection `json:"customer_data_collection,omitempty"`
	// LineItems 行项目列表。必填。
	LineItems []CheckoutLineItem `json:"line_items"`
	// Discounts 折扣列表。必填。
	Discounts []CheckoutDiscount `json:"discounts"`
	// PaymentOptions 支付选项。可选。
	PaymentOptions *CheckoutPaymentOptions `json:"payment_options,omitempty"`
	// PaymentSourceID 支付来源唯一标识符。可选。
	PaymentSourceID string `json:"payment_source_id,omitempty"`
	// InvoiceData 发票数据。可选。
	InvoiceData *CheckoutInvoiceData `json:"invoice_data,omitempty"`
	// InvoiceID 发票唯一标识符。可选。
	InvoiceID string `json:"invoice_id,omitempty"`
	// SubscriptionData 订阅数据。可选。
	SubscriptionData *CheckoutSubscriptionData `json:"subscription_data,omitempty"`
	// SubscriptionID 订阅唯一标识符。可选。
	SubscriptionID string `json:"subscription_id,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// UpdatedAt 更新时间。必填。
	UpdatedAt string `json:"updated_at"`
	// ExpiresAt 过期时间。必填。
	ExpiresAt string `json:"expires_at"`
	// CancelledAt 取消时间。可选。
	CancelledAt string `json:"cancelled_at,omitempty"`
	// CompletedAt 完成时间。可选。
	CompletedAt string `json:"completed_at,omitempty"`
}

// CreateCheckoutRequest 创建结账请求。
type CreateCheckoutRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Mode 结账模式。必填。
	Mode CheckoutMode `json:"mode"`
	// SuccessURL 成功跳转地址。必填。
	SuccessURL string `json:"success_url"`
	// BackURL 返回跳转地址。可选。
	BackURL string `json:"back_url,omitempty"`
	// BillingCustomerID 账单客户唯一标识符。可选。
	BillingCustomerID string `json:"billing_customer_id,omitempty"`
	// Currency 货币代码。可选。
	Currency string `json:"currency,omitempty"`
	// LegalEntityID 法律实体唯一标识符。可选。
	LegalEntityID string `json:"legal_entity_id,omitempty"`
	// LinkedPaymentAccountID 关联支付账户唯一标识符。可选。
	LinkedPaymentAccountID string `json:"linked_payment_account_id,omitempty"`
	// CustomerData 客户数据。可选。
	CustomerData *CheckoutCustomerData `json:"customer_data,omitempty"`
	// CustomerDataCollection 客户数据收集配置。可选。
	CustomerDataCollection *CheckoutCustomerDataCollection `json:"customer_data_collection,omitempty"`
	// LineItems 行项目列表。可选。
	LineItems []CheckoutLineItem `json:"line_items,omitempty"`
	// Discounts 折扣列表。可选。
	Discounts []CheckoutDiscount `json:"discounts,omitempty"`
	// PaymentOptions 支付选项。可选。
	PaymentOptions *CheckoutPaymentOptions `json:"payment_options,omitempty"`
	// PaymentSourceID 支付来源唯一标识符。可选。
	PaymentSourceID string `json:"payment_source_id,omitempty"`
	// InvoiceData 发票数据。可选。
	InvoiceData *CheckoutInvoiceData `json:"invoice_data,omitempty"`
	// SubscriptionData 订阅数据。可选。
	SubscriptionData *CheckoutSubscriptionData `json:"subscription_data,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// UpdateCheckoutRequest 更新结账请求。
type UpdateCheckoutRequest struct {
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// ListCheckoutsRequest 列出结账请求。
type ListCheckoutsRequest struct {
	// Page 分页游标。可选。
	Page string `json:"page,omitempty"`
	// PageSize 每页数量。可选。
	PageSize int32 `json:"page_size,omitempty"`
}

// CreateCheckout creates a new checkout session.
// CreateCheckout 创建结账。
// 官方文档: https://www.airwallex.com/docs/api/billing/billing_checkouts/create.md
func (s *Service) CreateCheckout(ctx context.Context, req *CreateCheckoutRequest, opts ...sdk.RequestOption) (*Checkout, error) {
	var resp Checkout
	err := s.doer.Do(ctx, "POST", "/api/v1/billing_checkouts/create", req, &resp, opts...)
	return &resp, err
}

// GetCheckout retrieves a checkout by ID.
// GetCheckout 根据 ID 获取结账。
// 官方文档: https://www.airwallex.com/docs/api/billing/billing_checkouts/retrieve.md
func (s *Service) GetCheckout(ctx context.Context, id string, opts ...sdk.RequestOption) (*Checkout, error) {
	var resp Checkout
	err := s.doer.Do(ctx, "GET", "/api/v1/billing_checkouts/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateCheckout updates an existing checkout.
// UpdateCheckout 更新结账。
// 官方文档: https://www.airwallex.com/docs/api/billing/billing_checkouts/update.md
func (s *Service) UpdateCheckout(ctx context.Context, id string, req *UpdateCheckoutRequest, opts ...sdk.RequestOption) (*Checkout, error) {
	var resp Checkout
	err := s.doer.Do(ctx, "POST", "/api/v1/billing_checkouts/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// CancelCheckout cancels a checkout session.
// CancelCheckout 取消结账。
// 官方文档: https://www.airwallex.com/docs/api/billing/billing_checkouts/cancel.md
func (s *Service) CancelCheckout(ctx context.Context, id string, opts ...sdk.RequestOption) (*Checkout, error) {
	var resp Checkout
	err := s.doer.Do(ctx, "POST", "/api/v1/billing_checkouts/"+id+"/cancel", nil, &resp, opts...)
	return &resp, err
}

// ListCheckouts lists checkout sessions with optional filters.
// ListCheckouts 列出结账。
// 官方文档: https://www.airwallex.com/docs/api/billing/billing_checkouts/list.md
func (s *Service) ListCheckouts(ctx context.Context, req *ListCheckoutsRequest, opts ...sdk.RequestOption) (*ListResult[Checkout], error) {
	var resp ListResult[Checkout]
	err := s.doer.Do(ctx, "GET", "/api/v1/billing_checkouts", req, &resp, opts...)
	return &resp, err
}
