package billing

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hakur/airwallex/sdk"
)

// BillingTransaction represents a billing transaction.
// BillingTransaction 账单交易。
type BillingTransaction struct {
	// ID is the unique identifier for the billing transaction. Required.
	// ID 账单交易唯一标识符。必填。
	ID string `json:"id"`
	// Amount is the total transaction amount. Required.
	// Amount 交易总金额。必填。
	Amount float64 `json:"amount"`
	// BillingCustomerID is the unique identifier of the billing customer who initiated the transaction. Required.
	// BillingCustomerID 发起交易的账单客户唯一标识符。必填。
	BillingCustomerID string `json:"billing_customer_id"`
	// CancelledAt is the transaction cancellation time. Optional.
	// CancelledAt 交易取消时间。可选。
	CancelledAt string `json:"cancelled_at,omitempty"`
	// CreatedAt is the transaction creation time. Required.
	// CreatedAt 交易创建时间。必填。
	CreatedAt string `json:"created_at"`
	// Currency is the transaction currency code (3-letter ISO-4217 format). Required.
	// Currency 交易货币代码（3位ISO-4217格式）。必填。
	Currency string `json:"currency"`
	// ExternalID is the unique identifier of the associated payment intent. Optional.
	// ExternalID 关联支付意向的唯一标识符。可选。
	ExternalID string `json:"external_id,omitempty"`
	// InvoiceID is the unique identifier of the associated invoice. Optional.
	// InvoiceID 关联发票的唯一标识符。可选。
	InvoiceID string `json:"invoice_id,omitempty"`
	// LinkedPaymentAccountID is the unique identifier of the linked payment account. Optional.
	// LinkedPaymentAccountID 关联收款支付账户唯一标识符。可选。
	LinkedPaymentAccountID string `json:"linked_payment_account_id,omitempty"`
	// Metadata is a set of key-value pairs for storing additional information. Optional.
	// Metadata 字符串键值对元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// OutOfBand indicates whether the transaction was manually marked as successful. Optional.
	// OutOfBand 是否手动标记为成功。可选。
	OutOfBand bool `json:"out_of_band,omitempty"`
	// PaymentSourceID is the unique identifier of the associated payment source. Optional.
	// PaymentSourceID 关联支付来源的唯一标识符。可选。
	PaymentSourceID string `json:"payment_source_id,omitempty"`
	// Status is the transaction status. Required.
	// Status 交易状态。必填。
	Status BillingTransactionStatus `json:"status"`
	// SucceededAt is the transaction success time. Optional.
	// SucceededAt 交易成功时间。可选。
	SucceededAt string `json:"succeeded_at,omitempty"`
	// Type is the transaction type. Optional.
	// Type 交易类型。可选。
	Type BillingTransactionType `json:"type,omitempty"`
	// UpdatedAt is the transaction update time. Required.
	// UpdatedAt 交易更新时间。必填。
	UpdatedAt string `json:"updated_at"`
}

// ListBillingTransactionsRequest represents a request to list billing transactions.
// ListBillingTransactionsRequest 列出账单交易请求。
type ListBillingTransactionsRequest struct {
	// InvoiceID is the unique identifier of the associated invoice. Required.
	// InvoiceID 关联发票唯一标识符。必填。
	InvoiceID string `json:"invoice_id"`
	// FromCreatedAt is the start of creation time (ISO8601 inclusive). Optional.
	// FromCreatedAt 创建时间起始（ISO8601 inclusive）。可选。
	FromCreatedAt string `json:"from_created_at,omitempty"`
	// ToCreatedAt is the end of creation time (ISO8601 exclusive). Optional.
	// ToCreatedAt 创建时间截止（ISO8601 exclusive）。可选。
	ToCreatedAt string `json:"to_created_at,omitempty"`
	// Page is the pagination cursor. Optional.
	// Page 分页游标。可选。
	Page string `json:"page,omitempty"`
	// PageSize is the number of items per page, defaults to 20. Optional.
	// PageSize 每页数量，默认20。可选。
	PageSize int32 `json:"page_size,omitempty"`
}

// GetBillingTransaction retrieves a single billing transaction by ID.
// GetBillingTransaction 获取单个账单交易。
// 官方文档: https://www.airwallex.com/docs/api/billing/billing_transactions/retrieve.md
func (s *Service) GetBillingTransaction(ctx context.Context, id string, opts ...sdk.RequestOption) (*BillingTransaction, error) {
	var resp BillingTransaction
	err := s.doer.Do(ctx, "GET", "/api/v1/billing_transactions/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListBillingTransactions lists billing transactions with optional filters.
// ListBillingTransactions 列出账单交易。
// 官方文档: https://www.airwallex.com/docs/api/billing/billing_transactions/list.md
func (s *Service) ListBillingTransactions(ctx context.Context, req *ListBillingTransactionsRequest, opts ...sdk.RequestOption) (*ListResult[BillingTransaction], error) {
	var resp ListResult[BillingTransaction]
	path := "/api/v1/billing_transactions"
	if req != nil {
		query := url.Values{}
		if req.InvoiceID != "" {
			query.Set("invoice_id", req.InvoiceID)
		}
		if req.FromCreatedAt != "" {
			query.Set("from_created_at", req.FromCreatedAt)
		}
		if req.ToCreatedAt != "" {
			query.Set("to_created_at", req.ToCreatedAt)
		}
		if req.Page != "" {
			query.Set("page", req.Page)
		}
		if req.PageSize > 0 {
			query.Set("page_size", fmt.Sprintf("%d", req.PageSize))
		}
		if len(query) > 0 {
			path = path + "?" + query.Encode()
		}
	}
	err := s.doer.Do(ctx, "GET", path, nil, &resp, opts...)
	return &resp, err
}
