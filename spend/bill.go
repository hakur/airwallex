package spend

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// BillStatus represents the bill status.
// BillStatus 账单状态。
type BillStatus = string

const (
	BillStatusDraft             BillStatus = "DRAFT"
	BillStatusAwaitingApproval  BillStatus = "AWAITING_APPROVAL"
	BillStatusAwaitingPayment   BillStatus = "AWAITING_PAYMENT"
	BillStatusPaymentInProgress BillStatus = "PAYMENT_IN_PROGRESS"
	BillStatusPaid              BillStatus = "PAID"
	BillStatusRejected          BillStatus = "REJECTED"
	BillStatusScheduled         BillStatus = "SCHEDULED"
)

// BillSyncStatus represents the bill sync status with the accounting system.
// BillSyncStatus 账单会计系统同步状态。
type BillSyncStatus = string

const (
	BillSyncStatusNotSynced  BillSyncStatus = "NOT_SYNCED"
	BillSyncStatusSynced     BillSyncStatus = "SYNCED"
	BillSyncStatusSyncFailed BillSyncStatus = "SYNC_FAILED"
)

// BillTaxStatus represents the bill tax status.
// BillTaxStatus 账单税务状态。
type BillTaxStatus = string

const (
	BillTaxStatusExclusive   BillTaxStatus = "TAX_EXCLUSIVE"
	BillTaxStatusInclusive   BillTaxStatus = "TAX_INCLUSIVE"
	BillTaxStatusUnspecified BillTaxStatus = "UNSPECIFIED"
)

// Bill represents a bill.
// Bill 表示账单。
type Bill struct {
	ID               string         `json:"id"`
	LegalEntityID    string         `json:"legal_entity_id"`
	InvoiceNumber    string         `json:"invoice_number,omitempty"`
	BillingAmount    string         `json:"billing_amount"`
	BillingCurrency  string         `json:"billing_currency"`
	Status           BillStatus     `json:"status"`
	SyncStatus       BillSyncStatus `json:"sync_status"`
	SyncErrorMessage string         `json:"sync_error_message,omitempty"`
	TaxStatus        BillTaxStatus  `json:"tax_status"`
	Description      string         `json:"description,omitempty"`
	Approvers        []string       `json:"approvers"`
	IssuedDate       string         `json:"issued_date,omitempty"`
	DueDate          string         `json:"due_date,omitempty"`
	VendorID         string         `json:"vendor_id,omitempty"`
	PurchaseOrderID  string         `json:"purchase_order_id,omitempty"`
	BillPayments     []BillPayment  `json:"bill_payments"`
	ExternalID       string         `json:"external_id,omitempty"`
	Attachments      []Attachment   `json:"attachments"`
	LineItems        []BillLineItem `json:"line_items"`
	CreatedAt        string         `json:"created_at"`
	UpdatedAt        string         `json:"updated_at"`
}

// BillPayment represents a bill payment record.
// BillPayment 账单付款记录。
type BillPayment struct {
	ID              string                      `json:"id"`
	CreatedAt       string                      `json:"created_at"`
	Type            string                      `json:"type"`
	Amount          string                      `json:"amount"`
	Currency        string                      `json:"currency"`
	Transfer        *BillTransferPayment        `json:"transfer,omitempty"`
	CardTransaction *BillCardTransactionPayment `json:"card_transaction,omitempty"`
}

// BillTransferPayment represents transfer payment details.
// BillTransferPayment 转账付款详情。
type BillTransferPayment struct {
	TransferID     string `json:"transfer_id"`
	AccountID      string `json:"account_id"`
	MultiBill      bool   `json:"multi_bill"`
	SourceAmount   string `json:"source_amount,omitempty"`
	SourceCurrency string `json:"source_currency,omitempty"`
	TransferDate   string `json:"transfer_date"`
}

// BillCardTransactionPayment represents card transaction payment details.
// BillCardTransactionPayment 银行卡付款详情。
type BillCardTransactionPayment struct {
	CardTransactionID string `json:"card_transaction_id"`
	CardID            string `json:"card_id"`
	AccountID         string `json:"account_id"`
	CardFundingType   string `json:"card_funding_type"`
	SourceAmount      string `json:"source_amount"`
	SourceCurrency    string `json:"source_currency"`
	TransactedAt      string `json:"transacted_at"`
}

// BillLineItem represents a bill line item.
// BillLineItem 账单行项目。
type BillLineItem struct {
	ID                        string                     `json:"id"`
	UnitPrice                 string                     `json:"unit_price,omitempty"`
	Quantity                  string                     `json:"quantity,omitempty"`
	Description               string                     `json:"description,omitempty"`
	TotalAmount               string                     `json:"total_amount"`
	TaxAmount                 string                     `json:"tax_amount,omitempty"`
	AccountingFieldSelections []AccountingFieldSelection `json:"accounting_field_selections"`
	PurchaseOrderLineItemID   string                     `json:"purchase_order_line_item_id,omitempty"`
}

// Attachment represents a file attachment.
// Attachment 文件附件。
type Attachment struct {
	ID          string `json:"id"`
	ContentType string `json:"content_type"`
	FileName    string `json:"file_name"`
	FileURL     string `json:"file_url"`
	CreatedAt   string `json:"created_at"`
}

// SyncBillRequest represents a request to update bill sync status.
// SyncBillRequest 更新账单同步状态请求。
type SyncBillRequest struct {
	SyncStatus       BillSyncStatus `json:"sync_status"`
	SyncErrorMessage string         `json:"sync_error_message,omitempty"`
}

// ListBillsRequest represents query parameters for listing bills.
// ListBillsRequest 账单列表查询参数。
type ListBillsRequest struct {
	Page            string `json:"page,omitempty"`
	FromCreatedAt   string `json:"from_created_at,omitempty"`
	ToCreatedAt     string `json:"to_created_at,omitempty"`
	Status          string `json:"status,omitempty"`
	SyncStatus      string `json:"sync_status,omitempty"`
	LegalEntityID   string `json:"legal_entity_id,omitempty"`
	PurchaseOrderID string `json:"purchase_order_id,omitempty"`
}

// ListBillsResponse represents a bill list response with cursor pagination.
// ListBillsResponse 账单列表响应（cursor 分页）。
type ListBillsResponse struct {
	Items      []Bill `json:"items"`
	PageAfter  string `json:"page_after,omitempty"`
	PageBefore string `json:"page_before,omitempty"`
}

// ListBills lists all bills.
// ListBills 列出账单。
// 官方文档: https://www.airwallex.com/docs/api/spend/bills/list.md
func (s *Service) ListBills(ctx context.Context, req *ListBillsRequest, opts ...sdk.RequestOption) (*ListBillsResponse, error) {
	var resp ListBillsResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/spend/bills", req, &resp, opts...)
	return &resp, err
}

// GetBill retrieves bill details.
// GetBill 获取账单详情。
// 官方文档: https://www.airwallex.com/docs/api/spend/bills/retrieve.md
func (s *Service) GetBill(ctx context.Context, id string, opts ...sdk.RequestOption) (*Bill, error) {
	var resp Bill
	err := s.doer.Do(ctx, "GET", "/api/v1/spend/bills/"+id, nil, &resp, opts...)
	return &resp, err
}

// SyncBill updates the bill sync status.
// SyncBill 更新账单同步状态。
// 官方文档: https://www.airwallex.com/docs/api/spend/bills/sync.md
func (s *Service) SyncBill(ctx context.Context, id string, req *SyncBillRequest, opts ...sdk.RequestOption) (*Bill, error) {
	var resp Bill
	err := s.doer.Do(ctx, "POST", "/api/v1/spend/bills/"+id+"/sync", req, &resp, opts...)
	return &resp, err
}

// --- Deprecated ---
// CreateBill / UpdateBill 已移除。官方 API 仅支持 List/Get/Sync，不支持创建和更新账单。
