package billing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// CreditNoteDiscountAmount 折扣金额。
type CreditNoteDiscountAmount struct {
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// AppliedDiscountID 已应用折扣唯一标识符。必填。
	AppliedDiscountID string `json:"applied_discount_id"`
}

// CreditNoteLineItemDiscountAmount 行项目折扣金额。
type CreditNoteLineItemDiscountAmount struct {
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// AppliedDiscountID 已应用折扣唯一标识符。必填。
	AppliedDiscountID string `json:"applied_discount_id"`
}

// CreditNoteLineItem 贷项通知单行项目。
type CreditNoteLineItem struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// CreditNoteID 贷项通知单唯一标识符。必填。
	CreditNoteID string `json:"credit_note_id"`
	// InvoiceLineItemID 发票行项唯一标识符。可选。
	InvoiceLineItemID string `json:"invoice_line_item_id,omitempty"`
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// Currency 货币代码。必填。
	Currency string `json:"currency"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// Quantity 数量。可选。
	Quantity int32 `json:"quantity,omitempty"`
	// UnitAmount 单价。可选。
	UnitAmount float64 `json:"unit_amount,omitempty"`
	// DiscountAmounts 折扣金额列表。必填。
	DiscountAmounts []CreditNoteLineItemDiscountAmount `json:"discount_amounts"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// UpdatedAt 更新时间。必填。
	UpdatedAt string `json:"updated_at"`
}

// CreditNote represents a credit note.
// CreditNote 贷项通知单。
type CreditNote struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// BillingCustomerID 账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// InvoiceID 发票唯一标识符。必填。
	InvoiceID string `json:"invoice_id"`
	// Currency 货币代码。必填。
	Currency string `json:"currency"`
	// Status 状态。必填。
	Status CreditNoteStatus `json:"status"`
	// Type 类型。必填。
	Type CreditNoteType `json:"type"`
	// Reason 原因。必填。
	Reason CreditNoteReason `json:"reason"`
	// CustomReason 自定义原因。可选。
	CustomReason string `json:"custom_reason,omitempty"`
	// Memo 备注。可选。
	Memo string `json:"memo,omitempty"`
	// Number 编号。必填。
	Number string `json:"number"`
	// TotalAmount 总金额。必填。
	TotalAmount float64 `json:"total_amount"`
	// AdjustmentAmount 调整金额。必填。
	AdjustmentAmount float64 `json:"adjustment_amount"`
	// RefundAmount 退款金额。必填。
	RefundAmount float64 `json:"refund_amount"`
	// OutOfBandAmount 外部支付金额。必填。
	OutOfBandAmount float64 `json:"out_of_band_amount"`
	// TotalTaxAmount 总税额。必填。
	TotalTaxAmount float64 `json:"total_tax_amount"`
	// TotalDiscountAmounts 总折扣金额列表。必填。
	TotalDiscountAmounts []CreditNoteDiscountAmount `json:"total_discount_amounts"`
	// RefundStatus 退款状态。可选。
	RefundStatus CreditNoteRefundStatus `json:"refund_status,omitempty"`
	// PDFURL PDF 地址。可选。
	PDFURL string `json:"pdf_url,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// UpdatedAt 更新时间。必填。
	UpdatedAt string `json:"updated_at"`
	// FinalizedAt 完成时间。可选。
	FinalizedAt string `json:"finalized_at,omitempty"`
	// VoidedAt 作废时间。可选。
	VoidedAt string `json:"voided_at,omitempty"`
}

// CreateCreditNoteRequest 创建请求。
type CreateCreditNoteRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// InvoiceID 发票唯一标识符。必填。
	InvoiceID string `json:"invoice_id"`
	// Type 类型。必填。
	Type CreditNoteType `json:"type"`
	// Reason 原因。必填。
	Reason CreditNoteReason `json:"reason"`
	// CustomReason 自定义原因。可选。
	CustomReason string `json:"custom_reason,omitempty"`
	// Memo 备注。可选。
	Memo string `json:"memo,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
}

// UpdateCreditNoteRequest 更新请求。
type UpdateCreditNoteRequest struct {
	// Reason 原因。可选。
	Reason string `json:"reason,omitempty"`
	// CustomReason 自定义原因。可选。
	CustomReason string `json:"custom_reason,omitempty"`
	// Memo 备注。可选。
	Memo string `json:"memo,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
}

// FinalizeCreditNoteRequest 完成请求。
type FinalizeCreditNoteRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// RefundAmount 退款金额。可选。
	RefundAmount float64 `json:"refund_amount,omitempty"`
	// OutOfBandAmount 外部支付金额。可选。
	OutOfBandAmount float64 `json:"out_of_band_amount,omitempty"`
}

// PreviewCreditNoteRequest 预览请求。
type PreviewCreditNoteRequest struct {
	// InvoiceID 发票唯一标识符。必填。
	InvoiceID string `json:"invoice_id"`
	// Type 类型。必填。
	Type CreditNoteType `json:"type"`
	// Reason 原因。必填。
	Reason CreditNoteReason `json:"reason"`
	// CustomReason 自定义原因。可选。
	CustomReason string `json:"custom_reason,omitempty"`
	// Memo 备注。可选。
	Memo string `json:"memo,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// LineItems 行项目列表。必填。
	LineItems []CreditNoteLineItemInput `json:"line_items"`
}

// CreditNoteLineItemInput 行项目输入。
type CreditNoteLineItemInput struct {
	// Amount 金额。可选。
	Amount float64 `json:"amount,omitempty"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// InvoiceLineItemID 发票行项唯一标识符。可选。
	InvoiceLineItemID string `json:"invoice_line_item_id,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// Quantity 数量。可选。
	Quantity int32 `json:"quantity,omitempty"`
	// UnitAmount 单价。可选。
	UnitAmount float64 `json:"unit_amount,omitempty"`
}

// AddCreditNoteLineItemsRequest 添加行项目请求。
type AddCreditNoteLineItemsRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// LineItems 行项目列表。必填。
	LineItems []CreditNoteLineItemInput `json:"line_items"`
}

// UpdateCreditNoteLineItemInput 更新行项目输入。
type UpdateCreditNoteLineItemInput struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Amount 金额。可选。
	Amount float64 `json:"amount,omitempty"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// Quantity 数量。可选。
	Quantity int32 `json:"quantity,omitempty"`
	// UnitAmount 单价。可选。
	UnitAmount float64 `json:"unit_amount,omitempty"`
}

// UpdateCreditNoteLineItemsRequest 更新行项目请求。
type UpdateCreditNoteLineItemsRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// LineItems 行项目列表。必填。
	LineItems []UpdateCreditNoteLineItemInput `json:"line_items"`
}

// DeleteCreditNoteLineItemsRequest 删除行项目请求。
type DeleteCreditNoteLineItemsRequest struct {
	// LineItemIDs 行项目唯一标识符列表。必填。
	LineItemIDs []string `json:"line_item_ids"`
}

// DeleteCreditNoteResponse 删除响应。
type DeleteCreditNoteResponse struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Deleted 是否已删除。必填。
	Deleted bool `json:"deleted"`
}

// ListCreditNotesRequest 列出请求。
type ListCreditNotesRequest struct {
	// InvoiceID 发票唯一标识符。可选。
	InvoiceID string `json:"invoice_id,omitempty"`
	// BillingCustomerID 账单客户唯一标识符。可选。
	BillingCustomerID string `json:"billing_customer_id,omitempty"`
	// Type 类型。可选。
	Type CreditNoteType `json:"type,omitempty"`
	// Status 状态。可选。
	Status CreditNoteStatus `json:"status,omitempty"`
	// FromCreatedAt 创建时间起始。可选。
	FromCreatedAt string `json:"from_created_at,omitempty"`
	// ToCreatedAt 创建时间截止。可选。
	ToCreatedAt string `json:"to_created_at,omitempty"`
	// Page 分页游标。可选。
	Page string `json:"page,omitempty"`
	// PageSize 每页数量。可选。
	PageSize int32 `json:"page_size,omitempty"`
}

// CreateCreditNote creates a new credit note.
// CreateCreditNote 创建贷项通知单。
// 官方文档: https://www.airwallex.com/docs/api/billing/credit_notes/create.md
func (s *Service) CreateCreditNote(ctx context.Context, req *CreateCreditNoteRequest, opts ...sdk.RequestOption) (*CreditNote, error) {
	var resp CreditNote
	err := s.doer.Do(ctx, "POST", "/api/v1/billing/credit_notes/create", req, &resp, opts...)
	return &resp, err
}

// PreviewCreditNote previews a credit note before creation.
// PreviewCreditNote 预览贷项通知单。
// 官方文档: https://www.airwallex.com/docs/api/billing/credit_notes/preview.md
func (s *Service) PreviewCreditNote(ctx context.Context, req *PreviewCreditNoteRequest, opts ...sdk.RequestOption) (*CreditNote, error) {
	var resp CreditNote
	err := s.doer.Do(ctx, "POST", "/api/v1/billing/credit_notes/preview", req, &resp, opts...)
	return &resp, err
}

// GetCreditNote retrieves a credit note by ID.
// GetCreditNote 获取贷项通知单。
// 官方文档: https://www.airwallex.com/docs/api/billing/credit_notes/retrieve.md
func (s *Service) GetCreditNote(ctx context.Context, id string, opts ...sdk.RequestOption) (*CreditNote, error) {
	var resp CreditNote
	err := s.doer.Do(ctx, "GET", "/api/v1/billing/credit_notes/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateCreditNote updates an existing credit note.
// UpdateCreditNote 更新贷项通知单。
// 官方文档: https://www.airwallex.com/docs/api/billing/credit_notes/update.md
func (s *Service) UpdateCreditNote(ctx context.Context, id string, req *UpdateCreditNoteRequest, opts ...sdk.RequestOption) (*CreditNote, error) {
	var resp CreditNote
	err := s.doer.Do(ctx, "POST", "/api/v1/billing/credit_notes/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// DeleteCreditNote deletes a draft credit note.
// DeleteCreditNote 删除贷项通知单。
// 官方文档: https://www.airwallex.com/docs/api/billing/credit_notes/delete.md
func (s *Service) DeleteCreditNote(ctx context.Context, id string, opts ...sdk.RequestOption) (*DeleteCreditNoteResponse, error) {
	var resp DeleteCreditNoteResponse
	err := s.doer.Do(ctx, "POST", "/api/v1/billing/credit_notes/"+id+"/delete", nil, &resp, opts...)
	return &resp, err
}

// FinalizeCreditNote finalizes a credit note.
// FinalizeCreditNote 完成贷项通知单。
// 官方文档: https://www.airwallex.com/docs/api/billing/credit_notes/finalize.md
func (s *Service) FinalizeCreditNote(ctx context.Context, id string, req *FinalizeCreditNoteRequest, opts ...sdk.RequestOption) (*CreditNote, error) {
	var resp CreditNote
	err := s.doer.Do(ctx, "POST", "/api/v1/billing/credit_notes/"+id+"/finalize", req, &resp, opts...)
	return &resp, err
}

// VoidCreditNote voids a credit note.
// VoidCreditNote 作废贷项通知单。
// 官方文档: https://www.airwallex.com/docs/api/billing/credit_notes/void.md
func (s *Service) VoidCreditNote(ctx context.Context, id string, opts ...sdk.RequestOption) (*CreditNote, error) {
	var resp CreditNote
	err := s.doer.Do(ctx, "POST", "/api/v1/billing/credit_notes/"+id+"/void", nil, &resp, opts...)
	return &resp, err
}

// ListCreditNotes lists credit notes with optional filters.
// ListCreditNotes 列出贷项通知单。
// 官方文档: https://www.airwallex.com/docs/api/billing/credit_notes/list.md
func (s *Service) ListCreditNotes(ctx context.Context, req *ListCreditNotesRequest, opts ...sdk.RequestOption) (*ListResult[CreditNote], error) {
	var resp ListResult[CreditNote]
	err := s.doer.Do(ctx, "GET", "/api/v1/billing/credit_notes", req, &resp, opts...)
	return &resp, err
}

// AddCreditNoteLineItems adds line items to a credit note.
// AddCreditNoteLineItems 添加行项目。
// 官方文档: https://www.airwallex.com/docs/api/billing/credit_notes/add_line_items.md
func (s *Service) AddCreditNoteLineItems(ctx context.Context, id string, req *AddCreditNoteLineItemsRequest, opts ...sdk.RequestOption) (*CreditNote, error) {
	var resp CreditNote
	err := s.doer.Do(ctx, "POST", "/api/v1/billing/credit_notes/"+id+"/add_line_items", req, &resp, opts...)
	return &resp, err
}

// UpdateCreditNoteLineItems updates line items on a credit note.
// UpdateCreditNoteLineItems 更新行项目。
// 官方文档: https://www.airwallex.com/docs/api/billing/credit_notes/update_line_items.md
func (s *Service) UpdateCreditNoteLineItems(ctx context.Context, id string, req *UpdateCreditNoteLineItemsRequest, opts ...sdk.RequestOption) (*CreditNote, error) {
	var resp CreditNote
	err := s.doer.Do(ctx, "POST", "/api/v1/billing/credit_notes/"+id+"/update_line_items", req, &resp, opts...)
	return &resp, err
}

// DeleteCreditNoteLineItems deletes line items from a credit note.
// DeleteCreditNoteLineItems 删除行项目。
// 官方文档: https://www.airwallex.com/docs/api/billing/credit_notes/delete_line_items.md
func (s *Service) DeleteCreditNoteLineItems(ctx context.Context, id string, req *DeleteCreditNoteLineItemsRequest, opts ...sdk.RequestOption) (*CreditNote, error) {
	var resp CreditNote
	err := s.doer.Do(ctx, "POST", "/api/v1/billing/credit_notes/"+id+"/delete_line_items", req, &resp, opts...)
	return &resp, err
}

// ListCreditNoteLineItems lists line items for a credit note.
// ListCreditNoteLineItems 列出行项目。
// 官方文档: https://www.airwallex.com/docs/api/billing/credit_notes/list.md
func (s *Service) ListCreditNoteLineItems(ctx context.Context, creditNoteID string, opts ...sdk.RequestOption) (*ListResult[CreditNoteLineItem], error) {
	var resp ListResult[CreditNoteLineItem]
	err := s.doer.Do(ctx, "GET", "/api/v1/billing/credit_notes/"+creditNoteID+"/line_items", nil, &resp, opts...)
	return &resp, err
}

// GetCreditNoteLineItem retrieves a line item detail for a credit note.
// GetCreditNoteLineItem 获取行项目详情。
// 官方文档: https://www.airwallex.com/docs/api/billing/credit_notes/retrieve.md
func (s *Service) GetCreditNoteLineItem(ctx context.Context, creditNoteID, lineItemID string, opts ...sdk.RequestOption) (*CreditNoteLineItem, error) {
	var resp CreditNoteLineItem
	err := s.doer.Do(ctx, "GET", "/api/v1/billing/credit_notes/"+creditNoteID+"/line_items/"+lineItemID, nil, &resp, opts...)
	return &resp, err
}
