package billing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// InvoiceLineItemPriceInput 发票行项价格输入。
type InvoiceLineItemPriceInput struct {
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// FlatAmount 固定金额。可选。
	FlatAmount float64 `json:"flat_amount,omitempty"`
	// PricingModel 定价模型。必填。
	PricingModel PricingModel `json:"pricing_model"`
	// ProductID 产品唯一标识符。必填。
	ProductID string `json:"product_id"`
	// UnitAmount 单价。可选。
	UnitAmount float64 `json:"unit_amount,omitempty"`
}

// InvoiceLineItemInput 发票行项请求输入。
type InvoiceLineItemInput struct {
	// Discounts 折扣列表。可选。
	Discounts []DiscountInput `json:"discounts,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// Price 价格。可选。
	Price *InvoiceLineItemPriceInput `json:"price,omitempty"`
	// PriceID 价格唯一标识符。可选。
	PriceID string `json:"price_id,omitempty"`
	// Quantity 数量。可选。
	Quantity int32 `json:"quantity,omitempty"`
}

// CreateInvoiceRequest 创建发票请求。
type CreateInvoiceRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// CollectionMethod 收款方式。可选。
	CollectionMethod CollectionMethod `json:"collection_method,omitempty"`
	// Currency 货币代码。必填。
	Currency string `json:"currency"`
	// DaysUntilDue 到期天数。可选。
	DaysUntilDue int32 `json:"days_until_due,omitempty"`
	// DefaultTaxPercent 默认税率百分比。可选。
	DefaultTaxPercent float64 `json:"default_tax_percent,omitempty"`
	// Discounts 折扣列表。可选。
	Discounts []DiscountInput `json:"discounts,omitempty"`
	// DueAt 到期时间。可选。
	DueAt string `json:"due_at,omitempty"`
	// LegalEntityID 法律实体唯一标识符。可选。
	LegalEntityID string `json:"legal_entity_id,omitempty"`
	// LinkedPaymentAccountID 关联支付账户唯一标识符。可选。
	LinkedPaymentAccountID string `json:"linked_payment_account_id,omitempty"`
	// Memo 备注。可选。
	Memo string `json:"memo,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// PaymentOptions 支付选项。可选。
	PaymentOptions *PaymentOptions `json:"payment_options,omitempty"`
	// PaymentSourceID 支付来源唯一标识符。可选。
	PaymentSourceID string `json:"payment_source_id,omitempty"`
}

// UpdateInvoiceRequest 更新发票请求。
type UpdateInvoiceRequest struct {
	// BillingCustomerID 账单客户唯一标识符。可选。
	BillingCustomerID string `json:"billing_customer_id,omitempty"`
	// CollectionMethod 收款方式。可选。
	CollectionMethod CollectionMethod `json:"collection_method,omitempty"`
	// Currency 货币代码。可选。
	Currency string `json:"currency,omitempty"`
	// DaysUntilDue 到期天数。可选。
	DaysUntilDue int32 `json:"days_until_due,omitempty"`
	// DefaultTaxPercent 默认税率百分比。可选。
	DefaultTaxPercent float64 `json:"default_tax_percent,omitempty"`
	// Discounts 折扣列表。可选。
	Discounts []DiscountInput `json:"discounts,omitempty"`
	// DueAt 到期时间。可选。
	DueAt string `json:"due_at,omitempty"`
	// LegalEntityID 法律实体唯一标识符。可选。
	LegalEntityID string `json:"legal_entity_id,omitempty"`
	// LinkedPaymentAccountID 关联支付账户唯一标识符。可选。
	LinkedPaymentAccountID string `json:"linked_payment_account_id,omitempty"`
	// Memo 备注。可选。
	Memo string `json:"memo,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// PaymentOptions 支付选项。可选。
	PaymentOptions *PaymentOptions `json:"payment_options,omitempty"`
	// PaymentSourceID 支付来源唯一标识符。可选。
	PaymentSourceID string `json:"payment_source_id,omitempty"`
}

// AddInvoiceLineItemsRequest 添加发票行项请求。
type AddInvoiceLineItemsRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// LineItems 行项目列表。必填。
	LineItems []InvoiceLineItemInput `json:"line_items"`
}

// UpdateInvoiceLineItem 更新发票行项。
type UpdateInvoiceLineItem struct {
	// Discounts 折扣列表。可选。
	Discounts []DiscountInput `json:"discounts,omitempty"`
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// Quantity 数量。可选。
	Quantity int32 `json:"quantity,omitempty"`
}

// UpdateInvoiceLineItemsRequest 更新发票行项请求。
type UpdateInvoiceLineItemsRequest struct {
	// LineItems 行项目列表。必填。
	LineItems []UpdateInvoiceLineItem `json:"line_items"`
}

// DeleteInvoiceLineItemsRequest 删除发票行项请求。
type DeleteInvoiceLineItemsRequest struct {
	// LineItemIDs 行项目唯一标识符列表。必填。
	LineItemIDs []string `json:"line_item_ids"`
}

// PreviewInvoiceRequest 预览发票请求。
type PreviewInvoiceRequest struct {
	// BillingCustomerID 账单客户唯一标识符。可选。
	BillingCustomerID string `json:"billing_customer_id,omitempty"`
	// BillingCycleAnchorAt 账单周期锚点时间。可选。
	BillingCycleAnchorAt string `json:"billing_cycle_anchor_at,omitempty"`
	// DefaultTaxPercent 默认税率百分比。可选。
	DefaultTaxPercent float64 `json:"default_tax_percent,omitempty"`
	// Discounts 折扣列表。可选。
	Discounts []DiscountInput `json:"discounts,omitempty"`
	// EndsAt 结束时间。可选。
	EndsAt string `json:"ends_at,omitempty"`
	// Items 订阅项列表。可选。
	Items []SubscriptionItemInput `json:"items,omitempty"`
	// LegalEntityID 法律实体唯一标识符。可选。
	LegalEntityID string `json:"legal_entity_id,omitempty"`
	// StartsAt 开始时间。可选。
	StartsAt string `json:"starts_at,omitempty"`
	// SubscriptionID 订阅唯一标识符。可选。
	SubscriptionID string `json:"subscription_id,omitempty"`
	// TrialEndsAt 试用结束时间。可选。
	TrialEndsAt string `json:"trial_ends_at,omitempty"`
}

// InvoiceLineItem 发票行项响应。
type InvoiceLineItem struct {
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// AppliedDiscounts 已应用折扣列表。必填。
	AppliedDiscounts []Discount `json:"applied_discounts"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// Currency 货币代码。必填。
	Currency string `json:"currency"`
	// DiscountAmounts 折扣金额列表。必填。
	DiscountAmounts []DiscountAmount `json:"discount_amounts"`
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// InvoiceID 发票唯一标识符。必填。
	InvoiceID string `json:"invoice_id"`
	// Label 标签。可选。
	Label string `json:"label,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// PeriodEndsAt 周期结束时间。可选。
	PeriodEndsAt string `json:"period_ends_at,omitempty"`
	// PeriodStartsAt 周期开始时间。可选。
	PeriodStartsAt string `json:"period_starts_at,omitempty"`
	// Price 价格。必填。
	Price Price `json:"price"`
	// Quantity 数量。可选。
	Quantity int32 `json:"quantity,omitempty"`
	// UpdatedAt 更新时间。必填。
	UpdatedAt string `json:"updated_at"`
}

// PreviewInvoiceItem 预览发票行项。
type PreviewInvoiceItem struct {
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// AppliedDiscounts 已应用折扣列表。必填。
	AppliedDiscounts []Discount `json:"applied_discounts"`
	// Currency 货币代码。必填。
	Currency string `json:"currency"`
	// DiscountAmounts 折扣金额列表。必填。
	DiscountAmounts []DiscountAmount `json:"discount_amounts"`
	// PeriodEndsAt 周期结束时间。可选。
	PeriodEndsAt string `json:"period_ends_at,omitempty"`
	// PeriodStartsAt 周期开始时间。可选。
	PeriodStartsAt string `json:"period_starts_at,omitempty"`
	// Price 价格。必填。
	Price Price `json:"price"`
	// Quantity 数量。必填。
	Quantity int32 `json:"quantity"`
}

// Invoice represents an invoice response.
// Invoice 发票响应。
type Invoice struct {
	// AppliedDiscounts 已应用折扣列表。必填。
	AppliedDiscounts []Discount `json:"applied_discounts"`
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// CollectionMethod 收款方式。可选。
	CollectionMethod CollectionMethod `json:"collection_method,omitempty"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// CreditNoteAmountAfterPayment 支付后贷项通知单金额。可选。
	CreditNoteAmountAfterPayment float64 `json:"credit_note_amount_after_payment,omitempty"`
	// CreditNoteAmountBeforePayment 支付前贷项通知单金额。可选。
	CreditNoteAmountBeforePayment float64 `json:"credit_note_amount_before_payment,omitempty"`
	// Currency 货币代码。必填。
	Currency string `json:"currency"`
	// DaysUntilDue 到期天数。可选。
	DaysUntilDue int32 `json:"days_until_due,omitempty"`
	// DefaultTaxPercent 默认税率百分比。可选。
	DefaultTaxPercent float64 `json:"default_tax_percent,omitempty"`
	// DueAt 到期时间。可选。
	DueAt string `json:"due_at,omitempty"`
	// FinalizedAt 完成时间。可选。
	FinalizedAt string `json:"finalized_at,omitempty"`
	// HostedURL 托管地址。可选。
	HostedURL string `json:"hosted_url,omitempty"`
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// LastPaymentAttemptAt 上次支付尝试时间。可选。
	LastPaymentAttemptAt string `json:"last_payment_attempt_at,omitempty"`
	// LegalEntityID 法律实体唯一标识符。可选。
	LegalEntityID string `json:"legal_entity_id,omitempty"`
	// LinkedPaymentAccountID 关联支付账户唯一标识符。可选。
	LinkedPaymentAccountID string `json:"linked_payment_account_id,omitempty"`
	// Memo 备注。可选。
	Memo string `json:"memo,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// NextPaymentAttemptAt 下次支付尝试时间。可选。
	NextPaymentAttemptAt string `json:"next_payment_attempt_at,omitempty"`
	// Number 编号。可选。
	Number string `json:"number,omitempty"`
	// PaidAt 支付时间。可选。
	PaidAt string `json:"paid_at,omitempty"`
	// PaidOutOfBand 是否外部支付。必填。
	PaidOutOfBand bool `json:"paid_out_of_band"`
	// PastPaymentAttemptCount 历史支付尝试次数。可选。
	PastPaymentAttemptCount int32 `json:"past_payment_attempt_count,omitempty"`
	// PaymentOptions 支付选项。可选。
	PaymentOptions *PaymentOptions `json:"payment_options,omitempty"`
	// PaymentSourceID 支付来源唯一标识符。可选。
	PaymentSourceID string `json:"payment_source_id,omitempty"`
	// PaymentStatus 支付状态。必填。
	PaymentStatus PaymentStatus `json:"payment_status"`
	// PDFURL PDF 地址。可选。
	PDFURL string `json:"pdf_url,omitempty"`
	// RemainingPaymentAttemptCount 剩余支付尝试次数。可选。
	RemainingPaymentAttemptCount int32 `json:"remaining_payment_attempt_count,omitempty"`
	// Status 发票状态。必填。
	Status InvoiceStatus `json:"status"`
	// SubscriptionID 订阅唯一标识符。可选。
	SubscriptionID string `json:"subscription_id,omitempty"`
	// TotalAmount 总金额。必填。
	TotalAmount float64 `json:"total_amount"`
	// TotalDiscountAmounts 总折扣金额列表。必填。
	TotalDiscountAmounts []DiscountAmount `json:"total_discount_amounts"`
	// TotalTaxAmount 总税额。必填。
	TotalTaxAmount float64 `json:"total_tax_amount"`
	// UpdatedAt 更新时间。必填。
	UpdatedAt string `json:"updated_at"`
	// VoidedAt 作废时间。可选。
	VoidedAt string `json:"voided_at,omitempty"`
}

// PreviewInvoiceResponse 预览发票响应。
type PreviewInvoiceResponse struct {
	// AppliedDiscounts 已应用折扣列表。必填。
	AppliedDiscounts []Discount `json:"applied_discounts"`
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// Currency 货币代码。必填。
	Currency string `json:"currency"`
	// DefaultTaxPercent 默认税率百分比。可选。
	DefaultTaxPercent float64 `json:"default_tax_percent,omitempty"`
	// Items 行项目列表。必填。
	Items []PreviewInvoiceItem `json:"items"`
	// LegalEntityID 法律实体唯一标识符。可选。
	LegalEntityID string `json:"legal_entity_id,omitempty"`
	// SubscriptionID 订阅唯一标识符。可选。
	SubscriptionID string `json:"subscription_id,omitempty"`
	// TotalAmount 总金额。必填。
	TotalAmount float64 `json:"total_amount"`
	// TotalDiscountAmounts 总折扣金额列表。必填。
	TotalDiscountAmounts []DiscountAmount `json:"total_discount_amounts"`
}

// DeleteInvoiceResponse 删除发票响应。
type DeleteInvoiceResponse struct {
	// Deleted 是否已删除。必填。
	Deleted bool `json:"deleted"`
	// ID 唯一标识符。必填。
	ID string `json:"id"`
}

// ListInvoicesRequest 获取发票列表请求。
type ListInvoicesRequest struct {
	// BillingCustomerID 账单客户唯一标识符。可选。
	BillingCustomerID string `json:"billing_customer_id,omitempty"`
	// SubscriptionID 订阅唯一标识符。可选。
	SubscriptionID string `json:"subscription_id,omitempty"`
	// Status 发票状态。可选。
	Status InvoiceStatus `json:"status,omitempty"`
	// PaymentStatus 支付状态。可选。
	PaymentStatus PaymentStatus `json:"payment_status,omitempty"`
	// FromCreatedAt 创建时间起始。可选。
	FromCreatedAt string `json:"from_created_at,omitempty"`
	// ToCreatedAt 创建时间截止。可选。
	ToCreatedAt string `json:"to_created_at,omitempty"`
	// Page 分页游标。可选。
	Page string `json:"page,omitempty"`
	// PageSize 每页数量。可选。
	PageSize int32 `json:"page_size,omitempty"`
}

// CreateInvoice creates a new invoice.
// CreateInvoice 创建发票。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/create.md
func (s *Service) CreateInvoice(ctx context.Context, req *CreateInvoiceRequest, opts ...sdk.RequestOption) (*Invoice, error) {
	var resp Invoice
	err := s.doer.Do(ctx, "POST", "/api/v1/invoices/create", req, &resp, opts...)
	return &resp, err
}

// PreviewInvoice previews an invoice before creation.
// PreviewInvoice 预览发票。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/preview.md
func (s *Service) PreviewInvoice(ctx context.Context, req *PreviewInvoiceRequest, opts ...sdk.RequestOption) (*PreviewInvoiceResponse, error) {
	var resp PreviewInvoiceResponse
	err := s.doer.Do(ctx, "POST", "/api/v1/invoices/preview", req, &resp, opts...)
	return &resp, err
}

// GetInvoice retrieves an invoice by ID.
// GetInvoice 根据 ID 获取发票。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/retrieve.md
func (s *Service) GetInvoice(ctx context.Context, id string, opts ...sdk.RequestOption) (*Invoice, error) {
	var resp Invoice
	err := s.doer.Do(ctx, "GET", "/api/v1/invoices/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateInvoice updates an existing invoice.
// UpdateInvoice 更新发票。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/update.md
func (s *Service) UpdateInvoice(ctx context.Context, id string, req *UpdateInvoiceRequest, opts ...sdk.RequestOption) (*Invoice, error) {
	var resp Invoice
	err := s.doer.Do(ctx, "POST", "/api/v1/invoices/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// FinalizeInvoice finalizes a draft invoice.
// FinalizeInvoice 完成发票。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/finalize.md
func (s *Service) FinalizeInvoice(ctx context.Context, id string, opts ...sdk.RequestOption) (*Invoice, error) {
	var resp Invoice
	err := s.doer.Do(ctx, "POST", "/api/v1/invoices/"+id+"/finalize", nil, &resp, opts...)
	return &resp, err
}

// VoidInvoice voids an invoice.
// VoidInvoice 作废发票。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/void.md
func (s *Service) VoidInvoice(ctx context.Context, id string, opts ...sdk.RequestOption) (*Invoice, error) {
	var resp Invoice
	err := s.doer.Do(ctx, "POST", "/api/v1/invoices/"+id+"/void", nil, &resp, opts...)
	return &resp, err
}

// MarkInvoiceAsPaid marks an invoice as paid.
// MarkInvoiceAsPaid 标记发票为已支付。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/mark_as_paid.md
func (s *Service) MarkInvoiceAsPaid(ctx context.Context, id string, opts ...sdk.RequestOption) (*Invoice, error) {
	var resp Invoice
	err := s.doer.Do(ctx, "POST", "/api/v1/invoices/"+id+"/mark_as_paid", nil, &resp, opts...)
	return &resp, err
}

// DeleteInvoice deletes a draft invoice.
// DeleteInvoice 删除草稿发票。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/delete.md
func (s *Service) DeleteInvoice(ctx context.Context, id string, opts ...sdk.RequestOption) (*DeleteInvoiceResponse, error) {
	var resp DeleteInvoiceResponse
	err := s.doer.Do(ctx, "POST", "/api/v1/invoices/"+id+"/delete", nil, &resp, opts...)
	return &resp, err
}

// AddInvoiceLineItems adds line items to an invoice.
// AddInvoiceLineItems 添加发票行项。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/add_line_items.md
func (s *Service) AddInvoiceLineItems(ctx context.Context, id string, req *AddInvoiceLineItemsRequest, opts ...sdk.RequestOption) (*Invoice, error) {
	var resp Invoice
	err := s.doer.Do(ctx, "POST", "/api/v1/invoices/"+id+"/add_line_items", req, &resp, opts...)
	return &resp, err
}

// UpdateInvoiceLineItems updates line items on an invoice.
// UpdateInvoiceLineItems 更新发票行项。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/update_line_items.md
func (s *Service) UpdateInvoiceLineItems(ctx context.Context, id string, req *UpdateInvoiceLineItemsRequest, opts ...sdk.RequestOption) (*Invoice, error) {
	var resp Invoice
	err := s.doer.Do(ctx, "POST", "/api/v1/invoices/"+id+"/update_line_items", req, &resp, opts...)
	return &resp, err
}

// DeleteInvoiceLineItems deletes line items from an invoice.
// DeleteInvoiceLineItems 删除发票行项。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/delete_line_items.md
func (s *Service) DeleteInvoiceLineItems(ctx context.Context, id string, req *DeleteInvoiceLineItemsRequest, opts ...sdk.RequestOption) (*Invoice, error) {
	var resp Invoice
	err := s.doer.Do(ctx, "POST", "/api/v1/invoices/"+id+"/delete_line_items", req, &resp, opts...)
	return &resp, err
}

// ListInvoices lists invoices with optional filters.
// ListInvoices 列出发票。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/list.md
func (s *Service) ListInvoices(ctx context.Context, req *ListInvoicesRequest, opts ...sdk.RequestOption) (*ListResult[Invoice], error) {
	var resp ListResult[Invoice]
	err := s.doer.Do(ctx, "GET", "/api/v1/invoices", req, &resp, opts...)
	return &resp, err
}

// ListInvoiceLineItems lists line items for an invoice.
// ListInvoiceLineItems 获取发票行项列表。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/list.md
func (s *Service) ListInvoiceLineItems(ctx context.Context, invoiceID string, opts ...sdk.RequestOption) (*ListResult[InvoiceLineItem], error) {
	var resp ListResult[InvoiceLineItem]
	err := s.doer.Do(ctx, "GET", "/api/v1/invoices/"+invoiceID+"/line_items", nil, &resp, opts...)
	return &resp, err
}

// GetInvoiceLineItem retrieves a line item detail for an invoice.
// GetInvoiceLineItem 获取发票行项详情。
// 官方文档: https://www.airwallex.com/docs/api/billing/invoices/retrieve.md
func (s *Service) GetInvoiceLineItem(ctx context.Context, invoiceID, itemID string, opts ...sdk.RequestOption) (*InvoiceLineItem, error) {
	var resp InvoiceLineItem
	err := s.doer.Do(ctx, "GET", "/api/v1/invoices/"+invoiceID+"/line_items/"+itemID, nil, &resp, opts...)
	return &resp, err
}
