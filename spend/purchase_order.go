package spend

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// POStatus represents the purchase order status.
// POStatus 采购订单状态。
type POStatus = string

const (
	POStatusOpen      POStatus = "OPEN"
	POStatusCompleted POStatus = "COMPLETED"
	POStatusArchived  POStatus = "ARCHIVED"
	POStatusClosed    POStatus = "CLOSED"
)

// POSyncStatus represents the purchase order sync status.
// POSyncStatus 采购订单同步状态。
type POSyncStatus = string

const (
	POSyncStatusNotSynced  POSyncStatus = "NOT_SYNCED"
	POSyncStatusSynced     POSyncStatus = "SYNCED"
	POSyncStatusSyncFailed POSyncStatus = "SYNC_FAILED"
)

// PurchaseOrder represents a purchase order.
// PurchaseOrder 表示采购订单。
type PurchaseOrder struct {
	ID                    string        `json:"id"`
	LegalEntityID         string        `json:"legal_entity_id"`
	VendorID              string        `json:"vendor_id"`
	BillingCurrency       string        `json:"billing_currency"`
	RecurringAmount       string        `json:"recurring_amount,omitempty"`
	TotalAmount           string        `json:"total_amount"`
	BilledAmount          string        `json:"billed_amount"`
	PurchaseOrderNumber   string        `json:"purchase_order_number"`
	ExternalID            string        `json:"external_id,omitempty"`
	Status                POStatus      `json:"status"`
	SyncStatus            POSyncStatus  `json:"sync_status"`
	SyncErrorMessage      string        `json:"sync_error_message,omitempty"`
	Recurrence            *PORecurrence `json:"recurrence,omitempty"`
	Description           string        `json:"description,omitempty"`
	Note                  string        `json:"note,omitempty"`
	NetPaymentTermsInDays int32         `json:"net_payment_terms_in_days,omitempty"`
	PromiseDate           string        `json:"promise_date,omitempty"`
	Approvers             []string      `json:"approvers"`
	OwnerEmail            string        `json:"owner_email,omitempty"`
	Contacts              *POContacts   `json:"contacts,omitempty"`
	Attachments           []Attachment  `json:"attachments"`
	LineItems             []POLineItem  `json:"line_items"`
	Comments              []Comment     `json:"comments"`
	CreatedAt             string        `json:"created_at"`
	UpdatedAt             string        `json:"updated_at"`
}

// PORecurrence represents a purchase order recurrence.
// PORecurrence 采购订单重复周期。
type PORecurrence struct {
	Frequency  string `json:"frequency"`
	Iterations int32  `json:"iterations"`
}

// POContacts represents purchase order contacts.
// POContacts 采购订单联系人。
type POContacts struct {
	VendorContact   *POContact `json:"vendor_contact,omitempty"`
	ShippingContact *POContact `json:"shipping_contact,omitempty"`
	BillingContact  *POContact `json:"billing_contact,omitempty"`
}

// POContact represents contact details.
// POContact 联系人详情。
type POContact struct {
	ContactName string   `json:"contact_name,omitempty"`
	CompanyName string   `json:"company_name,omitempty"`
	Email       string   `json:"email,omitempty"`
	PhoneNumber string   `json:"phone_number,omitempty"`
	Address     *Address `json:"address,omitempty"`
}

// POLineItem represents a purchase order line item (response).
// POLineItem 采购订单行项目（响应）。
type POLineItem struct {
	ID                        string                     `json:"id"`
	Description               string                     `json:"description,omitempty"`
	Quantity                  string                     `json:"quantity"`
	UnitPrice                 string                     `json:"unit_price"`
	AccountingFieldSelections []AccountingFieldSelection `json:"accounting_field_selections"`
}

// CreatePOLineItem represents a purchase order line item creation request (different accounting field format).
// CreatePOLineItem 创建采购订单行项目（请求，accounting field 格式不同）。
type CreatePOLineItem struct {
	Description               string                             `json:"description,omitempty"`
	Quantity                  string                             `json:"quantity"`
	UnitPrice                 string                             `json:"unit_price"`
	AccountingFieldSelections []CreatePOAccountingFieldSelection `json:"accounting_field_selections,omitempty"`
}

// CreatePOAccountingFieldSelection represents an accounting field selection in creation requests (different from response format).
// CreatePOAccountingFieldSelection 创建请求中的会计字段选择（响应格式不同）。
type CreatePOAccountingFieldSelection struct {
	IdentifierType string `json:"identifier_type"`
	FieldID        string `json:"field_id"`
	FieldValueID   string `json:"field_value_id"`
}

// CreatePurchaseOrderRequest represents a request to create a purchase order.
// CreatePurchaseOrderRequest 创建采购订单请求。
type CreatePurchaseOrderRequest struct {
	RequestID             string             `json:"request_id"`
	ExternalID            string             `json:"external_id"`
	LegalEntityID         string             `json:"legal_entity_id"`
	VendorID              string             `json:"vendor_id"`
	PurchaseOrderNumber   string             `json:"purchase_order_number"`
	SyncStatus            POSyncStatus       `json:"sync_status"`
	BillingCurrency       string             `json:"billing_currency"`
	OwnerEmail            string             `json:"owner_email,omitempty"`
	Recurrence            *PORecurrence      `json:"recurrence,omitempty"`
	LineItems             []CreatePOLineItem `json:"line_items"`
	Description           string             `json:"description,omitempty"`
	Note                  string             `json:"note,omitempty"`
	NetPaymentTermsInDays int32              `json:"net_payment_terms_in_days,omitempty"`
	PromiseDate           string             `json:"promise_date,omitempty"`
	Contacts              *POContacts        `json:"contacts,omitempty"`
}

// SyncPurchaseOrderRequest represents a request to sync a purchase order.
// SyncPurchaseOrderRequest 同步采购订单请求。
type SyncPurchaseOrderRequest struct {
	SyncStatus       POSyncStatus `json:"sync_status"`
	SyncErrorMessage string       `json:"sync_error_message,omitempty"`
}

// ListPurchaseOrdersRequest represents query parameters for listing purchase orders.
// ListPurchaseOrdersRequest 采购订单列表查询参数。
type ListPurchaseOrdersRequest struct {
	Page          string   `json:"page,omitempty"`
	FromCreatedAt string   `json:"from_created_at,omitempty"`
	ToCreatedAt   string   `json:"to_created_at,omitempty"`
	Status        []string `json:"status,omitempty"`
	SyncStatus    []string `json:"sync_status,omitempty"`
	LegalEntityID string   `json:"legal_entity_id,omitempty"`
}

// ListPurchaseOrdersResponse represents a purchase order list response with cursor pagination.
// ListPurchaseOrdersResponse 采购订单列表响应（cursor 分页）。
type ListPurchaseOrdersResponse struct {
	Items      []PurchaseOrder `json:"items"`
	PageAfter  string          `json:"page_after,omitempty"`
	PageBefore string          `json:"page_before,omitempty"`
}

// CreatePurchaseOrder creates a purchase order.
// CreatePurchaseOrder 创建采购订单。
// 官方文档: https://www.airwallex.com/docs/api/spend/purchase_orders/create.md
func (s *Service) CreatePurchaseOrder(ctx context.Context, req *CreatePurchaseOrderRequest, opts ...sdk.RequestOption) (*PurchaseOrder, error) {
	var resp PurchaseOrder
	err := s.doer.Do(ctx, "POST", "/api/v1/spend/purchase_orders/create", req, &resp, opts...)
	return &resp, err
}

// GetPurchaseOrder retrieves purchase order details.
// GetPurchaseOrder 获取采购订单详情。
// 官方文档: https://www.airwallex.com/docs/api/spend/purchase_orders/retrieve.md
func (s *Service) GetPurchaseOrder(ctx context.Context, id string, opts ...sdk.RequestOption) (*PurchaseOrder, error) {
	var resp PurchaseOrder
	err := s.doer.Do(ctx, "GET", "/api/v1/spend/purchase_orders/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListPurchaseOrders lists all purchase orders.
// ListPurchaseOrders 列出采购订单。
// 官方文档: https://www.airwallex.com/docs/api/spend/purchase_orders/list.md
func (s *Service) ListPurchaseOrders(ctx context.Context, req *ListPurchaseOrdersRequest, opts ...sdk.RequestOption) (*ListPurchaseOrdersResponse, error) {
	var resp ListPurchaseOrdersResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/spend/purchase_orders", req, &resp, opts...)
	return &resp, err
}

// SyncPurchaseOrder updates the purchase order sync status.
// SyncPurchaseOrder 更新采购订单同步状态。
// 官方文档: https://www.airwallex.com/docs/api/spend/purchase_orders/sync.md
func (s *Service) SyncPurchaseOrder(ctx context.Context, id string, req *SyncPurchaseOrderRequest, opts ...sdk.RequestOption) (*PurchaseOrder, error) {
	var resp PurchaseOrder
	err := s.doer.Do(ctx, "POST", "/api/v1/spend/purchase_orders/"+id+"/sync", req, &resp, opts...)
	return &resp, err
}

// --- Deprecated ---
// UpdatePurchaseOrder 已移除。官方 API 不支持更新采购订单。
