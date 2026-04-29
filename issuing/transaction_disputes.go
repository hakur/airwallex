package issuing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// DisputeStatus represents the dispute status.
// DisputeStatus 争议状态。
type DisputeStatus = string

const (
	// DisputeStatusDraft is a draft dispute.
	// DisputeStatusDraft 草稿
	DisputeStatusDraft DisputeStatus = "DRAFT"
	// DisputeStatusSubmitted is a submitted dispute.
	// DisputeStatusSubmitted 已提交
	DisputeStatusSubmitted DisputeStatus = "SUBMITTED"
	// DisputeStatusRejected is a rejected dispute.
	// DisputeStatusRejected 已拒绝
	DisputeStatusRejected DisputeStatus = "REJECTED"
	// DisputeStatusCanceled is a canceled dispute.
	// DisputeStatusCanceled 已取消
	DisputeStatusCanceled DisputeStatus = "CANCELED"
	// DisputeStatusInProgress is a dispute in progress.
	// DisputeStatusInProgress 处理中
	DisputeStatusInProgress DisputeStatus = "IN_PROGRESS"
	// DisputeStatusWon is a won dispute.
	// DisputeStatusWon 已获胜
	DisputeStatusWon DisputeStatus = "WON"
	// DisputeStatusLost is a lost dispute.
	// DisputeStatusLost 已失败
	DisputeStatusLost DisputeStatus = "LOST"
	// DisputeStatusExpired is an expired dispute.
	// DisputeStatusExpired 已过期
	DisputeStatusExpired DisputeStatus = "EXPIRED"
)

// DisputeReason represents the dispute reason enum.
// DisputeReason 争议原因枚举。
type DisputeReason = string

const (
	// DisputeReasonSuspectedFraud is a suspected fraud dispute.
	// DisputeReasonSuspectedFraud 疑似欺诈
	DisputeReasonSuspectedFraud DisputeReason = "SUSPECTED_FRAUD"
	// DisputeReasonUnauthorizedTransaction is an unauthorized transaction dispute.
	// DisputeReasonUnauthorizedTransaction 未授权交易
	DisputeReasonUnauthorizedTransaction DisputeReason = "UNAUTHORIZED_TRANSACTION"
	// DisputeReasonDuplicatedTransaction is a duplicated transaction dispute.
	// DisputeReasonDuplicatedTransaction 重复交易
	DisputeReasonDuplicatedTransaction DisputeReason = "DUPLICATED_TRANSACTION"
	// DisputeReasonPaidByOtherMeans indicates the transaction was paid by other means.
	// DisputeReasonPaidByOtherMeans 其他方式支付
	DisputeReasonPaidByOtherMeans DisputeReason = "PAID_BY_OTHER_MEANS"
	// DisputeReasonGoodsServiceNotAsDescribed indicates goods/services not as described.
	// DisputeReasonGoodsServiceNotAsDescribed 商品/服务与描述不符
	DisputeReasonGoodsServiceNotAsDescribed DisputeReason = "GOODS_SERVICE_NOT_AS_DESCRIBED"
	// DisputeReasonGoodsDamaged indicates damaged goods.
	// DisputeReasonGoodsDamaged 商品损坏
	DisputeReasonGoodsDamaged DisputeReason = "GOODS_DAMAGED"
	// DisputeReasonGoodsServiceNotReceived indicates goods/services not received.
	// DisputeReasonGoodsServiceNotReceived 未收到商品/服务
	DisputeReasonGoodsServiceNotReceived DisputeReason = "GOODS_SERVICE_NOT_RECEIVED"
	// DisputeReasonRefundUnprocessed indicates an unprocessed refund.
	// DisputeReasonRefundUnprocessed 退款未处理
	DisputeReasonRefundUnprocessed DisputeReason = "REFUND_UNPROCESSED"
	// DisputeReasonGoodsServiceCanceled indicates canceled goods/services.
	// DisputeReasonGoodsServiceCanceled 商品/服务已取消
	DisputeReasonGoodsServiceCanceled DisputeReason = "GOODS_SERVICE_CANCELED"
	// DisputeReasonRecurringCanceled indicates a canceled recurring transaction.
	// DisputeReasonRecurringCanceled 定期交易已取消
	DisputeReasonRecurringCanceled DisputeReason = "RECURRING_CANCELED"
	// DisputeReasonOther is an other reason.
	// DisputeReasonOther 其他
	DisputeReasonOther DisputeReason = "OTHER"
)

// DisputeDetailedStatus represents the detailed dispute status.
// DisputeDetailedStatus 争议详细状态。
type DisputeDetailedStatus = string

const (
	// DisputeDetailedStatusDisputeFiled indicates the dispute has been filed.
	// DisputeDetailedStatusDisputeFiled 争议已提交
	DisputeDetailedStatusDisputeFiled DisputeDetailedStatus = "DISPUTE_FILED"
	// DisputeDetailedStatusPreArbReceived indicates pre-arbitration has been received.
	// DisputeDetailedStatusPreArbReceived 预仲裁已接收
	DisputeDetailedStatusPreArbReceived DisputeDetailedStatus = "PRE_ARB_RECEIVED"
	// DisputeDetailedStatusPreArbDeclinedByIssuer indicates pre-arbitration declined by issuer.
	// DisputeDetailedStatusPreArbDeclinedByIssuer 发卡行拒绝预仲裁
	DisputeDetailedStatusPreArbDeclinedByIssuer DisputeDetailedStatus = "PRE_ARB_DECLINED_BY_ISSUER"
	// DisputeDetailedStatusArbitrationReceived indicates arbitration has been received.
	// DisputeDetailedStatusArbitrationReceived 仲裁已接收
	DisputeDetailedStatusArbitrationReceived DisputeDetailedStatus = "ARBITRATION_RECEIVED"
	// DisputeDetailedStatusDisputeDeclined indicates the dispute has been declined.
	// DisputeDetailedStatusDisputeDeclined 争议被拒绝
	DisputeDetailedStatusDisputeDeclined DisputeDetailedStatus = "DISPUTE_DECLINED"
	// DisputeDetailedStatusPreArbDelivered indicates pre-arbitration has been delivered.
	// DisputeDetailedStatusPreArbDelivered 预仲裁已送达
	DisputeDetailedStatusPreArbDelivered DisputeDetailedStatus = "PRE_ARB_DELIVERED"
	// DisputeDetailedStatusPreArbDeclinedByAcquirer indicates pre-arbitration declined by acquirer.
	// DisputeDetailedStatusPreArbDeclinedByAcquirer 收单行拒绝预仲裁
	DisputeDetailedStatusPreArbDeclinedByAcquirer DisputeDetailedStatus = "PRE_ARB_DECLINED_BY_ACQUIRER"
	// DisputeDetailedStatusArbDelivered indicates arbitration has been delivered.
	// DisputeDetailedStatusArbDelivered 仲裁已送达
	DisputeDetailedStatusArbDelivered DisputeDetailedStatus = "ARB_DELIVERED"
	// DisputeDetailedStatusWon indicates the dispute was won.
	// DisputeDetailedStatusWon 获胜
	DisputeDetailedStatusWon DisputeDetailedStatus = "WON"
	// DisputeDetailedStatusLost indicates the dispute was lost.
	// DisputeDetailedStatusLost 失败
	DisputeDetailedStatusLost DisputeDetailedStatus = "LOST"
)

// DisputeUpdateHistoryEntry represents a dispute update history record.
// DisputeUpdateHistoryEntry 争议更新历史记录。
type DisputeUpdateHistoryEntry struct {
	// EvidenceFiles is the list of evidence file IDs. Optional.
	// EvidenceFiles 证据文件ID列表。可选。
	EvidenceFiles []string `json:"evidence_files,omitempty"`
	// Note is the note. Optional.
	// Note 备注。可选。
	Note string `json:"note,omitempty"`
	// UpdatedAt is the update time. Optional.
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
	// UpdatedBy is the updater. Optional.
	// UpdatedBy 更新者。可选。
	UpdatedBy string `json:"updated_by,omitempty"`
}

// TransactionDispute represents transaction dispute information.
// TransactionDispute 表示交易争议信息。
type TransactionDispute struct {
	// ID is the unique identifier. Optional.
	// ID 唯一标识符。可选。
	ID string `json:"id,omitempty"`
	// TransactionID is the unique transaction identifier. Optional.
	// TransactionID 交易唯一标识符。可选。
	TransactionID string `json:"transaction_id,omitempty"`
	// Status is the status. Optional.
	// Status 状态。可选。
	Status DisputeStatus `json:"status,omitempty"`
	// Amount is the amount. Optional.
	// Amount 金额。可选。
	Amount float64 `json:"amount,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// DetailedStatus is the detailed status. Optional.
	// DetailedStatus 详细状态。可选。
	DetailedStatus DisputeDetailedStatus `json:"detailed_status,omitempty"`
	// Notes are the notes. Optional.
	// Notes 备注。可选。
	Notes string `json:"notes,omitempty"`
	// Reason is the dispute reason. Optional.
	// Reason 争议原因。可选。
	Reason DisputeReason `json:"reason,omitempty"`
	// Reference is the reference number. Optional.
	// Reference 引用编号。可选。
	Reference string `json:"reference,omitempty"`
	// UpdateHistory is the update history. Optional.
	// UpdateHistory 更新历史。可选。
	UpdateHistory []DisputeUpdateHistoryEntry `json:"update_history,omitempty"`
	// UpdatedAt is the update time. Optional.
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
	// UpdatedBy is the updater. Optional.
	// UpdatedBy 更新者。可选。
	UpdatedBy string `json:"updated_by,omitempty"`
}

// CreateTransactionDisputeRequest represents the request to create a transaction dispute.
// CreateTransactionDisputeRequest 创建交易争议请求。
type CreateTransactionDisputeRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// TransactionID is the unique transaction identifier. Required.
	// TransactionID 交易唯一标识符。必填。
	TransactionID string `json:"transaction_id"`
	// Amount is the amount. Required.
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// Reason is the dispute reason. Required.
	// Reason 争议原因。必填。
	Reason DisputeReason `json:"reason"`
	// EvidenceFiles is the list of evidence files. Optional.
	// EvidenceFiles 证据文件列表。可选。
	EvidenceFiles []string `json:"evidence_files,omitempty"`
	// Notes are the notes. Optional.
	// Notes 备注。可选。
	Notes string `json:"notes,omitempty"`
	// Reference is the reference number. Optional.
	// Reference 引用编号。可选。
	Reference string `json:"reference,omitempty"`
}

// UpdateTransactionDisputeRequest represents the request to update a transaction dispute.
// UpdateTransactionDisputeRequest 更新交易争议请求。
type UpdateTransactionDisputeRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Amount is the dispute amount. Optional.
	// Amount 争议金额。可选。
	Amount float64 `json:"amount,omitempty"`
	// EvidenceFiles is the list of evidence files. Optional.
	// EvidenceFiles 证据文件列表。可选。
	EvidenceFiles []string `json:"evidence_files,omitempty"`
	// Notes are the notes. Optional.
	// Notes 备注。可选。
	Notes string `json:"notes,omitempty"`
	// Reason is the dispute reason. Optional.
	// Reason 争议原因。可选。
	Reason DisputeReason `json:"reason,omitempty"`
}

// CreateTransactionDispute creates a transaction dispute.
// 官方文档: https://www.airwallex.com/docs/api/issuing/transaction_disputes/create.md
// CreateTransactionDispute 创建交易争议。
func (s *Service) CreateTransactionDispute(ctx context.Context, req *CreateTransactionDisputeRequest, opts ...sdk.RequestOption) (*TransactionDispute, error) {
	var resp TransactionDispute
	err := s.doer.Do(ctx, "POST", "/api/v1/issuing/transaction_disputes/create", req, &resp, opts...)
	return &resp, err
}

// GetTransactionDispute retrieves a transaction dispute by ID.
// 官方文档: https://www.airwallex.com/docs/api/issuing/transaction_disputes/retrieve.md
// GetTransactionDispute 根据 ID 获取交易争议。
func (s *Service) GetTransactionDispute(ctx context.Context, id string, opts ...sdk.RequestOption) (*TransactionDispute, error) {
	var resp TransactionDispute
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/transaction_disputes/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListTransactionDisputes lists transaction disputes.
// 官方文档: https://www.airwallex.com/docs/api/issuing/transaction_disputes/list.md
// ListTransactionDisputes 列出交易争议。
func (s *Service) ListTransactionDisputes(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[TransactionDispute], error) {
	var resp sdk.ListResult[TransactionDispute]
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/transaction_disputes", nil, &resp, opts...)
	return &resp, err
}

// UpdateTransactionDispute updates a transaction dispute.
// 官方文档: https://www.airwallex.com/docs/api/issuing/transaction_disputes/update.md
// UpdateTransactionDispute 更新交易争议。
func (s *Service) UpdateTransactionDispute(ctx context.Context, id string, req *UpdateTransactionDisputeRequest, opts ...sdk.RequestOption) (*TransactionDispute, error) {
	var resp TransactionDispute
	err := s.doer.Do(ctx, "POST", "/api/v1/issuing/transaction_disputes/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// CancelTransactionDispute cancels a transaction dispute.
// 官方文档: https://www.airwallex.com/docs/api/issuing/transaction_disputes/cancel.md
// CancelTransactionDispute 取消交易争议。
func (s *Service) CancelTransactionDispute(ctx context.Context, id string, opts ...sdk.RequestOption) (*TransactionDispute, error) {
	var resp TransactionDispute
	err := s.doer.Do(ctx, "POST", "/api/v1/issuing/transaction_disputes/"+id+"/cancel", nil, &resp, opts...)
	return &resp, err
}

// SubmitTransactionDispute submits a transaction dispute.
// 官方文档: https://www.airwallex.com/docs/api/issuing/transaction_disputes/submit.md
// SubmitTransactionDispute 提交交易争议。
func (s *Service) SubmitTransactionDispute(ctx context.Context, id string, opts ...sdk.RequestOption) (*TransactionDispute, error) {
	var resp TransactionDispute
	err := s.doer.Do(ctx, "POST", "/api/v1/issuing/transaction_disputes/"+id+"/submit", nil, &resp, opts...)
	return &resp, err
}
