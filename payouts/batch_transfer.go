package payouts

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// BatchTransferFundingSource represents a batch transfer funding source.
// BatchTransferFundingSource 批量转账资金来源。
type BatchTransferFundingSource struct {
	// DepositType 存款类型（DIRECT_DEBIT或FASTER_DIRECT_DEBIT）。
	DepositType string `json:"deposit_type,omitempty"`
	// ID 资金来源ID（仅支持关联账户）。
	ID string `json:"id,omitempty"`
	// Reference 用户指定的参考号（显示在Direct Debit银行对账单上）。
	Reference string `json:"reference,omitempty"`
}

// BatchTransferQuote represents a batch transfer quote.
// BatchTransferQuote 批量转账报价。
type BatchTransferQuote struct {
	// AmountBeneficiaryReceives 收款人收到的总金额。
	AmountBeneficiaryReceives float64 `json:"amount_beneficiary_receives,omitempty"`
	// AmountPayerPays 付款人支付的总金额。
	AmountPayerPays float64 `json:"amount_payer_pays,omitempty"`
	// ClientRate 转换汇率（纯转账则为null）。
	ClientRate float64 `json:"client_rate,omitempty"`
	// CurrencyPair 货币对（纯转账则为null）。
	CurrencyPair string `json:"currency_pair,omitempty"`
	// FeeAmount 总费用金额。
	FeeAmount float64 `json:"fee_amount,omitempty"`
	// FeeCurrency 费用货币。
	FeeCurrency sdk.Currency `json:"fee_currency,omitempty"`
	// SourceCurrency 源货币（付款人用于资助转账的货币）。
	SourceCurrency sdk.Currency `json:"source_currency,omitempty"`
	// TransferCurrency 转账货币（收款人收到的货币）。
	TransferCurrency sdk.Currency `json:"transfer_currency,omitempty"`
}

// BatchTransferQuoteSummary represents a batch transfer quote summary.
// BatchTransferQuoteSummary 批量转账报价摘要。
type BatchTransferQuoteSummary struct {
	// ExpiresAt 报价过期时间。
	ExpiresAt string `json:"expires_at,omitempty"`
	// LastQuotedAt 最后报价时间。
	LastQuotedAt string `json:"last_quoted_at,omitempty"`
	// Quotes 报价列表。
	Quotes []BatchTransferQuote `json:"quotes,omitempty"`
	// Validity 报价有效期。
	Validity string `json:"validity,omitempty"`
}

// BatchTransferFunding represents batch transfer funding information.
// BatchTransferFunding 批量转账资金信息。
type BatchTransferFunding struct {
	// DepositType 存款类型。
	DepositType string `json:"deposit_type,omitempty"`
	// FailureReason 失败原因（仅在错误状态时非空）。
	FailureReason string `json:"failure_reason,omitempty"`
	// FundingSourceID 资金来源ID（使用钱包时为null）。
	FundingSourceID string `json:"funding_source_id,omitempty"`
	// Reference 用户指定的参考号。
	Reference string `json:"reference,omitempty"`
	// Status 资金状态。
	Status string `json:"status,omitempty"`
}

// BatchTransfer represents a batch transfer.
// BatchTransfer 表示批量转账信息。
type BatchTransfer struct {
	// ID 系统生成的批量转账ID。
	ID string `json:"id,omitempty"`
	// Status 批量转账状态。
	Status string `json:"status,omitempty"`
	// Funding 资金信息。
	Funding *BatchTransferFunding `json:"funding,omitempty"`
	// Metadata 元数据。
	Metadata map[string]any `json:"metadata,omitempty"`
	// Name 批量转账名称。
	Name string `json:"name,omitempty"`
	// QuoteSummary 报价摘要。
	QuoteSummary *BatchTransferQuoteSummary `json:"quote_summary,omitempty"`
	// Remarks 备注（不会传递给收款人）。
	Remarks string `json:"remarks,omitempty"`
	// RequestID 客户端指定的唯一批量转账参考号。
	RequestID string `json:"request_id,omitempty"`
	// ShortReferenceID 系统生成的短参考ID。
	ShortReferenceID string `json:"short_reference_id,omitempty"`
	// TotalItemCount 批量转账中的总项目数。
	TotalItemCount int `json:"total_item_count,omitempty"`
	// TransferDate 计划处理日期。
	TransferDate string `json:"transfer_date,omitempty"`
	// UpdatedAt 最后更新时间。
	UpdatedAt string `json:"updated_at,omitempty"`
	// ValidItemCount 有效项目数。
	ValidItemCount int `json:"valid_item_count,omitempty"`
}

// BatchTransferItem represents a batch transfer item.
// BatchTransferItem 批量转账项目。
type BatchTransferItem struct {
	// ID 系统生成的批量转账项目ID。
	ID string `json:"id,omitempty"`
	// RequestID 用户原始的请求ID。
	RequestID string `json:"request_id,omitempty"`
	// Status 项目状态。
	Status string `json:"status,omitempty"`
	// TransferID 系统生成的交易ID（未预订则返回null）。
	TransferID string `json:"transfer_id,omitempty"`
	// TransferDraft 转账草稿信息。
	TransferDraft map[string]any `json:"transfer_draft,omitempty"`
	// Errors 错误详情（当状态为VALIDATION_FAILED或BOOKING_FAILED时）。
	Errors []map[string]any `json:"errors,omitempty"`
	// UpdatedAt 最后更新时间。
	UpdatedAt string `json:"updated_at,omitempty"`
}

// CreateBatchTransferRequest is the request to create a batch transfer.
// CreateBatchTransferRequest 创建批量转账请求。
type CreateBatchTransferRequest struct {
	// RequestID 客户端指定的唯一批量转账参考号。必填。
	RequestID string `json:"request_id"`
	// FundingSource 资金来源。
	FundingSource *BatchTransferFundingSource `json:"funding_source,omitempty"`
	// Metadata 元数据。
	Metadata map[string]any `json:"metadata,omitempty"`
	// Name 批量转账名称。
	Name string `json:"name,omitempty"`
	// Remarks 备注。
	Remarks string `json:"remarks,omitempty"`
	// TransferDate 转账日期（ISO 8601格式）。
	TransferDate string `json:"transfer_date,omitempty"`
}

// AddBatchTransferItemsRequest is the request to add items to a batch transfer.
// AddBatchTransferItemsRequest 添加批量转账项目请求。
type AddBatchTransferItemsRequest struct {
	// Items 批量项目列表（每次最多100个，每批最多1000个）。必填。
	Items []CreateTransferRequest `json:"items"`
}

// DeleteBatchTransferItemsRequest is the request to delete items from a batch transfer.
// DeleteBatchTransferItemsRequest 删除批量转账项目请求。
type DeleteBatchTransferItemsRequest struct {
	// ItemIDs 要删除的项目ID列表。必填。
	ItemIDs []string `json:"item_ids"`
}

// QuoteBatchTransferRequest is the request to quote a batch transfer.
// QuoteBatchTransferRequest 批量转账报价请求。
type QuoteBatchTransferRequest struct {
	// Validity 报价有效期。
	Validity string `json:"validity,omitempty"`
}

// CreateBatchTransfer creates a batch transfer.
// CreateBatchTransfer 创建批量转账。
// 官方文档: https://www.airwallex.com/docs/api/payouts/batch_transfers/create.md
func (s *Service) CreateBatchTransfer(ctx context.Context, req *CreateBatchTransferRequest, opts ...sdk.RequestOption) (*BatchTransfer, error) {
	var resp BatchTransfer
	err := s.doer.Do(ctx, "POST", "/api/v1/batch_transfers/create", req, &resp, opts...)
	return &resp, err
}

// GetBatchTransfer retrieves a batch transfer by ID.
// GetBatchTransfer 根据ID获取批量转账。
// 官方文档: https://www.airwallex.com/docs/api/payouts/batch_transfers/retrieve.md
func (s *Service) GetBatchTransfer(ctx context.Context, id string, opts ...sdk.RequestOption) (*BatchTransfer, error) {
	var resp BatchTransfer
	err := s.doer.Do(ctx, "GET", "/api/v1/batch_transfers/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListBatchTransfers lists batch transfers.
// ListBatchTransfers 列出批量转账。
// 官方文档: https://www.airwallex.com/docs/api/payouts/batch_transfers/list.md
func (s *Service) ListBatchTransfers(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[BatchTransfer], error) {
	var resp sdk.ListResult[BatchTransfer]
	err := s.doer.Do(ctx, "GET", "/api/v1/batch_transfers", nil, &resp, opts...)
	return &resp, err
}

// DeleteBatchTransfer deletes a batch transfer.
// DeleteBatchTransfer 删除批量转账。
// 官方文档: https://www.airwallex.com/docs/api/payouts/batch_transfers/delete.md
func (s *Service) DeleteBatchTransfer(ctx context.Context, id string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/batch_transfers/"+id+"/delete", map[string]any{}, nil, opts...)
}

// AddBatchTransferItems adds items to a batch transfer.
// AddBatchTransferItems 添加项目到批量转账。
// 官方文档: https://www.airwallex.com/docs/api/payouts/batch_transfers/add_items.md
func (s *Service) AddBatchTransferItems(ctx context.Context, id string, req *AddBatchTransferItemsRequest, opts ...sdk.RequestOption) (*BatchTransfer, error) {
	var resp BatchTransfer
	err := s.doer.Do(ctx, "POST", "/api/v1/batch_transfers/"+id+"/add_items", req, &resp, opts...)
	return &resp, err
}

// DeleteBatchTransferItems deletes items from a batch transfer.
// DeleteBatchTransferItems 删除批量转账中的项目。
// 官方文档: https://www.airwallex.com/docs/api/payouts/batch_transfers/delete_items.md
func (s *Service) DeleteBatchTransferItems(ctx context.Context, id string, req *DeleteBatchTransferItemsRequest, opts ...sdk.RequestOption) (*BatchTransfer, error) {
	var resp BatchTransfer
	err := s.doer.Do(ctx, "POST", "/api/v1/batch_transfers/"+id+"/delete_items", req, &resp, opts...)
	return &resp, err
}

// ListBatchTransferItems lists all items in a batch transfer.
// ListBatchTransferItems 列出批量转账中的所有项目。
// 官方文档: https://www.airwallex.com/docs/api/payouts/batch_transfers/items.md
func (s *Service) ListBatchTransferItems(ctx context.Context, id string, opts ...sdk.RequestOption) (*sdk.ListResult[BatchTransferItem], error) {
	var resp sdk.ListResult[BatchTransferItem]
	err := s.doer.Do(ctx, "GET", "/api/v1/batch_transfers/"+id+"/items", nil, &resp, opts...)
	return &resp, err
}

// QuoteBatchTransfer gets a quote for a batch transfer.
// QuoteBatchTransfer 获取批量转账报价。
// 官方文档: https://www.airwallex.com/docs/api/payouts/batch_transfers/quote.md
func (s *Service) QuoteBatchTransfer(ctx context.Context, id string, req *QuoteBatchTransferRequest, opts ...sdk.RequestOption) (*BatchTransfer, error) {
	var resp BatchTransfer
	err := s.doer.Do(ctx, "POST", "/api/v1/batch_transfers/"+id+"/quote", req, &resp, opts...)
	return &resp, err
}

// SubmitBatchTransfer submits a batch transfer.
// SubmitBatchTransfer 提交批量转账。
// 官方文档: https://www.airwallex.com/docs/api/payouts/batch_transfers/submit.md
func (s *Service) SubmitBatchTransfer(ctx context.Context, id string, opts ...sdk.RequestOption) (*BatchTransfer, error) {
	var resp BatchTransfer
	err := s.doer.Do(ctx, "POST", "/api/v1/batch_transfers/"+id+"/submit", nil, &resp, opts...)
	return &resp, err
}
