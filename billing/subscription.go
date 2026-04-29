package billing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// SubscriptionItemInput 订阅项请求输入。
type SubscriptionItemInput struct {
	// Discounts 折扣列表。可选。
	Discounts []DiscountInput `json:"discounts,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// PriceID 价格唯一标识符。必填。
	PriceID string `json:"price_id"`
	// Quantity 数量。可选。
	Quantity int32 `json:"quantity,omitempty"`
}

// UpdateSubscriptionItemInput 更新订阅项输入。
type UpdateSubscriptionItemInput struct {
	// Deleted 是否删除。可选。
	Deleted bool `json:"deleted,omitempty"`
	// Discounts 折扣列表。可选。
	Discounts []DiscountInput `json:"discounts,omitempty"`
	// ID 订阅项唯一标识符。可选。
	ID string `json:"id,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// PriceID 价格唯一标识符。可选。
	PriceID string `json:"price_id,omitempty"`
	// ProrationMode 按比例退款行为。可选。
	ProrationMode ProrationBehavior `json:"proration_mode,omitempty"`
	// Quantity 数量。可选。
	Quantity int32 `json:"quantity,omitempty"`
}

// CreateSubscriptionRequest 创建订阅请求。
type CreateSubscriptionRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// BillingCycleAnchorAt 账单周期锚点时间。可选。
	BillingCycleAnchorAt string `json:"billing_cycle_anchor_at,omitempty"`
	// CollectionMethod 收款方式。必填。
	CollectionMethod CollectionMethod `json:"collection_method"`
	// Currency 货币代码。可选。
	Currency string `json:"currency,omitempty"`
	// DaysUntilDue 到期天数。可选。
	DaysUntilDue int32 `json:"days_until_due,omitempty"`
	// DefaultInvoiceTemplate 默认发票模板。可选。
	DefaultInvoiceTemplate *DefaultInvoiceTemplate `json:"default_invoice_template,omitempty"`
	// DefaultTaxPercent 默认税率百分比。可选。
	DefaultTaxPercent float64 `json:"default_tax_percent,omitempty"`
	// Discounts 折扣列表。可选。
	Discounts []DiscountInput `json:"discounts,omitempty"`
	// Duration 持续时间。可选。
	Duration *Duration `json:"duration,omitempty"`
	// EndsAt 结束时间。可选。
	EndsAt string `json:"ends_at,omitempty"`
	// Items 订阅项列表。必填。
	Items []SubscriptionItemInput `json:"items"`
	// LegalEntityID 法律实体唯一标识符。可选。
	LegalEntityID string `json:"legal_entity_id,omitempty"`
	// LinkedPaymentAccountID 关联支付账户唯一标识符。可选。
	LinkedPaymentAccountID string `json:"linked_payment_account_id,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// PaymentOptions 支付选项。可选。
	PaymentOptions *PaymentOptions `json:"payment_options,omitempty"`
	// PaymentSourceID 支付来源唯一标识符。可选。
	PaymentSourceID string `json:"payment_source_id,omitempty"`
	// StartsAt 开始时间。可选。
	StartsAt string `json:"starts_at,omitempty"`
	// TrialEndsAt 试用结束时间。可选。
	TrialEndsAt string `json:"trial_ends_at,omitempty"`
}

// UpdateSubscriptionRequest 更新订阅请求。
type UpdateSubscriptionRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// BillingAction 账单处理方式。可选。
	BillingAction BillingAction `json:"billing_action,omitempty"`
	// CancelAtPeriodEnd 是否在周期结束时取消。可选。
	CancelAtPeriodEnd bool `json:"cancel_at_period_end,omitempty"`
	// CollectionMethod 收款方式。可选。
	CollectionMethod CollectionMethod `json:"collection_method,omitempty"`
	// DaysUntilDue 到期天数。可选。
	DaysUntilDue int32 `json:"days_until_due,omitempty"`
	// DefaultInvoiceTemplate 默认发票模板。可选。
	DefaultInvoiceTemplate *DefaultInvoiceTemplate `json:"default_invoice_template,omitempty"`
	// DefaultProrationMode 默认按比例退款行为。可选。
	DefaultProrationMode ProrationBehavior `json:"default_proration_mode,omitempty"`
	// DefaultTaxPercent 默认税率百分比。可选。
	DefaultTaxPercent float64 `json:"default_tax_percent,omitempty"`
	// Discounts 折扣列表。可选。
	Discounts []DiscountInput `json:"discounts,omitempty"`
	// Duration 持续时间。可选。
	Duration *Duration `json:"duration,omitempty"`
	// EndsAt 结束时间。可选。
	EndsAt string `json:"ends_at,omitempty"`
	// Items 订阅项列表。可选。
	Items []UpdateSubscriptionItemInput `json:"items,omitempty"`
	// LegalEntityID 法律实体唯一标识符。可选。
	LegalEntityID string `json:"legal_entity_id,omitempty"`
	// LinkedPaymentAccountID 关联支付账户唯一标识符。可选。
	LinkedPaymentAccountID string `json:"linked_payment_account_id,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// PaymentOptions 支付选项。可选。
	PaymentOptions *PaymentOptions `json:"payment_options,omitempty"`
	// PaymentSourceID 支付来源唯一标识符。可选。
	PaymentSourceID string `json:"payment_source_id,omitempty"`
	// RemainingDuration 剩余持续时间。可选。
	RemainingDuration *Duration `json:"remaining_duration,omitempty"`
	// TrialEndsAt 试用结束时间。可选。
	TrialEndsAt string `json:"trial_ends_at,omitempty"`
}

// CancelSubscriptionRequest 取消订阅请求。
type CancelSubscriptionRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// ProrationBehavior 按比例退款行为。必填。
	ProrationBehavior ProrationBehavior `json:"proration_behavior"`
}

// SubscriptionItem 订阅项响应。
type SubscriptionItem struct {
	// AppliedDiscounts 已应用折扣列表。必填。
	AppliedDiscounts []Discount `json:"applied_discounts"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// CurrentPeriodEndsAt 当前周期结束时间。可选。
	CurrentPeriodEndsAt string `json:"current_period_ends_at,omitempty"`
	// CurrentPeriodStartsAt 当前周期开始时间。可选。
	CurrentPeriodStartsAt string `json:"current_period_starts_at,omitempty"`
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// Price 价格。必填。
	Price Price `json:"price"`
	// Quantity 数量。可选。
	Quantity int32 `json:"quantity,omitempty"`
	// SubscriptionID 订阅唯一标识符。必填。
	SubscriptionID string `json:"subscription_id"`
	// UpdatedAt 更新时间。必填。
	UpdatedAt string `json:"updated_at"`
}

// Subscription 订阅响应。
type Subscription struct {
	// AppliedDiscounts 已应用折扣列表。必填。
	AppliedDiscounts []Discount `json:"applied_discounts"`
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// BillingCycleAnchorAt 账单周期锚点时间。可选。
	BillingCycleAnchorAt string `json:"billing_cycle_anchor_at,omitempty"`
	// CancelAtPeriodEnd 是否在周期结束时取消。必填。
	CancelAtPeriodEnd bool `json:"cancel_at_period_end"`
	// CancelRequestedAt 取消请求时间。可选。
	CancelRequestedAt string `json:"cancel_requested_at,omitempty"`
	// CollectionMethod 收款方式。必填。
	CollectionMethod CollectionMethod `json:"collection_method"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// Currency 货币代码。必填。
	Currency string `json:"currency"`
	// CurrentPeriodEndsAt 当前周期结束时间。可选。
	CurrentPeriodEndsAt string `json:"current_period_ends_at,omitempty"`
	// CurrentPeriodStartsAt 当前周期开始时间。可选。
	CurrentPeriodStartsAt string `json:"current_period_starts_at,omitempty"`
	// DaysUntilDue 到期天数。必填。
	DaysUntilDue int32 `json:"days_until_due"`
	// DefaultInvoiceTemplate 默认发票模板。可选。
	DefaultInvoiceTemplate *DefaultInvoiceTemplate `json:"default_invoice_template,omitempty"`
	// DefaultTaxPercent 默认税率百分比。可选。
	DefaultTaxPercent float64 `json:"default_tax_percent,omitempty"`
	// Duration 持续时间。可选。
	Duration *Duration `json:"duration,omitempty"`
	// EndsAt 结束时间。可选。
	EndsAt string `json:"ends_at,omitempty"`
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// LatestInvoiceID 最新发票唯一标识符。可选。
	LatestInvoiceID string `json:"latest_invoice_id,omitempty"`
	// LegalEntityID 法律实体唯一标识符。可选。
	LegalEntityID string `json:"legal_entity_id,omitempty"`
	// LinkedPaymentAccountID 关联支付账户唯一标识符。可选。
	LinkedPaymentAccountID string `json:"linked_payment_account_id,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// NextBillingAt 下次账单时间。可选。
	NextBillingAt string `json:"next_billing_at,omitempty"`
	// PaymentOptions 支付选项。可选。
	PaymentOptions *PaymentOptions `json:"payment_options,omitempty"`
	// PaymentSourceID 支付来源唯一标识符。可选。
	PaymentSourceID string `json:"payment_source_id,omitempty"`
	// Recurring 周期性配置。必填。
	Recurring Recurring `json:"recurring"`
	// StartsAt 开始时间。必填。
	StartsAt string `json:"starts_at"`
	// Status 订阅状态。必填。
	Status SubscriptionStatus `json:"status"`
	// TrialEndsAt 试用结束时间。可选。
	TrialEndsAt string `json:"trial_ends_at,omitempty"`
	// TrialStartsAt 试用开始时间。可选。
	TrialStartsAt string `json:"trial_starts_at,omitempty"`
	// UpdatedAt 更新时间。必填。
	UpdatedAt string `json:"updated_at"`
}

// ListSubscriptionsRequest 获取订阅列表请求。
type ListSubscriptionsRequest struct {
	// BillingCustomerID 账单客户唯一标识符。可选。
	BillingCustomerID string `json:"billing_customer_id,omitempty"`
	// Status 订阅状态。可选。
	Status SubscriptionStatus `json:"status,omitempty"`
	// FromCreatedAt 创建时间起始。可选。
	FromCreatedAt string `json:"from_created_at,omitempty"`
	// ToCreatedAt 创建时间截止。可选。
	ToCreatedAt string `json:"to_created_at,omitempty"`
	// RecurringPeriod 周期数。可选。
	RecurringPeriod int32 `json:"recurring_period,omitempty"`
	// RecurringPeriodUnit 周期单位。可选。
	RecurringPeriodUnit PeriodUnit `json:"recurring_period_unit,omitempty"`
	// Page 分页游标。可选。
	Page string `json:"page,omitempty"`
	// PageSize 每页数量。可选。
	PageSize int32 `json:"page_size,omitempty"`
}

// CreateSubscription creates a new subscription.
// CreateSubscription 创建订阅。
// 官方文档: https://www.airwallex.com/docs/api/billing/subscriptions/create.md
func (s *Service) CreateSubscription(ctx context.Context, req *CreateSubscriptionRequest, opts ...sdk.RequestOption) (*Subscription, error) {
	var resp Subscription
	err := s.doer.Do(ctx, "POST", "/api/v1/subscriptions/create", req, &resp, opts...)
	return &resp, err
}

// GetSubscription retrieves a subscription by ID.
// GetSubscription 根据 ID 获取订阅。
// 官方文档: https://www.airwallex.com/docs/api/billing/subscriptions/retrieve.md
func (s *Service) GetSubscription(ctx context.Context, id string, opts ...sdk.RequestOption) (*Subscription, error) {
	var resp Subscription
	err := s.doer.Do(ctx, "GET", "/api/v1/subscriptions/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateSubscription updates an existing subscription.
// UpdateSubscription 更新订阅。
// 官方文档: https://www.airwallex.com/docs/api/billing/subscriptions/update.md
func (s *Service) UpdateSubscription(ctx context.Context, id string, req *UpdateSubscriptionRequest, opts ...sdk.RequestOption) (*Subscription, error) {
	var resp Subscription
	err := s.doer.Do(ctx, "POST", "/api/v1/subscriptions/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// CancelSubscription cancels a subscription.
// CancelSubscription 取消订阅。
// 官方文档: https://www.airwallex.com/docs/api/billing/subscriptions/cancel.md
func (s *Service) CancelSubscription(ctx context.Context, id string, req *CancelSubscriptionRequest, opts ...sdk.RequestOption) (*Subscription, error) {
	var resp Subscription
	err := s.doer.Do(ctx, "POST", "/api/v1/subscriptions/"+id+"/cancel", req, &resp, opts...)
	return &resp, err
}

// ListSubscriptions lists subscriptions with optional filters.
// ListSubscriptions 列出订阅。
// 官方文档: https://www.airwallex.com/docs/api/billing/subscriptions/list.md
func (s *Service) ListSubscriptions(ctx context.Context, req *ListSubscriptionsRequest, opts ...sdk.RequestOption) (*ListResult[Subscription], error) {
	var resp ListResult[Subscription]
	err := s.doer.Do(ctx, "GET", "/api/v1/subscriptions", req, &resp, opts...)
	return &resp, err
}

// ListSubscriptionItems lists items for a subscription.
// ListSubscriptionItems 列出订阅项。
// 官方文档: https://www.airwallex.com/docs/api/billing/subscriptions/list.md
func (s *Service) ListSubscriptionItems(ctx context.Context, subscriptionID string, opts ...sdk.RequestOption) (*ListResult[SubscriptionItem], error) {
	var resp ListResult[SubscriptionItem]
	err := s.doer.Do(ctx, "GET", "/api/v1/subscriptions/"+subscriptionID+"/items", nil, &resp, opts...)
	return &resp, err
}

// GetSubscriptionItem retrieves a subscription item detail.
// GetSubscriptionItem 获取订阅项详情。
// 官方文档: https://www.airwallex.com/docs/api/billing/subscriptions/retrieve.md
func (s *Service) GetSubscriptionItem(ctx context.Context, subscriptionID, itemID string, opts ...sdk.RequestOption) (*SubscriptionItem, error) {
	var resp SubscriptionItem
	err := s.doer.Do(ctx, "GET", "/api/v1/subscriptions/"+subscriptionID+"/items/"+itemID, nil, &resp, opts...)
	return &resp, err
}
