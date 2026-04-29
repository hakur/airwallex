package scale

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// ConnectedAccountTransferStatus represents a connected account transfer status.
// ConnectedAccountTransferStatus 连接账户转账状态。
type ConnectedAccountTransferStatus = string

const (
	ConnectedAccountTransferStatusNew       ConnectedAccountTransferStatus = "NEW"
	ConnectedAccountTransferStatusSettled   ConnectedAccountTransferStatus = "SETTLED"
	ConnectedAccountTransferStatusPending   ConnectedAccountTransferStatus = "PENDING"
	ConnectedAccountTransferStatusSuspended ConnectedAccountTransferStatus = "SUSPENDED"
	ConnectedAccountTransferStatusFailed    ConnectedAccountTransferStatus = "FAILED"
)

// ConnectedAccountTransfer represents a connected account transfer.
// ConnectedAccountTransfer 表示连接账户转账信息。
type ConnectedAccountTransfer struct {
	// ID 转账唯一标识符。
	ID string `json:"id,omitempty"`
	// Destination 目标Airwallex账户ID。
	Destination string `json:"destination,omitempty"`
	// Amount 转账金额。
	Amount float64 `json:"amount,omitempty"`
	// Currency 转账货币（3位ISO-4217代码）。
	Currency sdk.Currency `json:"currency,omitempty"`
	// Status 转账状态。
	Status ConnectedAccountTransferStatus `json:"status,omitempty"`
	// Description 转账描述。
	Description string `json:"description,omitempty"`
	// Reason 转账原因。
	Reason string `json:"reason,omitempty"`
	// Reference 用户指定的参考信息。
	Reference string `json:"reference,omitempty"`
	// RequestID 请求唯一标识符。
	RequestID string `json:"request_id,omitempty"`
	// ShortReferenceID 短参考ID（用于客服支持）。
	ShortReferenceID string `json:"short_reference_id,omitempty"`
	// Fee 转账费用。
	Fee float64 `json:"fee,omitempty"`
	// FailureReasons 失败原因列表（当状态为FAILED时）。
	FailureReasons []string `json:"failure_reasons,omitempty"`
	// AdditionalInfo 附加信息。
	AdditionalInfo map[string]any `json:"additional_info,omitempty"`
	// CreatedAt 创建时间。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt 更新时间。
	UpdatedAt string `json:"updated_at,omitempty"`
}

// CreateConnectedAccountTransferRequest represents a request to create a connected account transfer.
// CreateConnectedAccountTransferRequest 创建连接账户转账请求。
type CreateConnectedAccountTransferRequest struct {
	// RequestID 请求唯一标识符（1-50字符）。必填。
	RequestID string `json:"request_id"`
	// Destination 目标Airwallex账户ID。必填。
	Destination string `json:"destination"`
	// Amount 转账金额（字符串类型）。必填。
	Amount string `json:"amount"`
	// Currency 转账货币（3位ISO-4217代码）。必填。
	Currency sdk.Currency `json:"currency"`
	// Reason 转账原因。必填。
	Reason string `json:"reason"`
	// Reference 用户指定的参考信息（1-140字符）。必填。
	Reference string `json:"reference"`
	// Description 转账描述。
	Description string `json:"description,omitempty"`
	// AdditionalInfo 附加信息。
	AdditionalInfo map[string]any `json:"additional_info,omitempty"`
}

// CreateConnectedAccountTransfer creates a connected account transfer.
// CreateConnectedAccountTransfer 创建连接账户转账。
// 官方文档: https://www.airwallex.com/docs/api/scale/connected_account_transfers/create.md
func (s *Service) CreateConnectedAccountTransfer(ctx context.Context, req *CreateConnectedAccountTransferRequest, opts ...sdk.RequestOption) (*ConnectedAccountTransfer, error) {
	var resp ConnectedAccountTransfer
	err := s.doer.Do(ctx, "POST", "/api/v1/connected_account_transfers/create", req, &resp, opts...)
	return &resp, err
}

// GetConnectedAccountTransfer retrieves a connected account transfer by ID.
// GetConnectedAccountTransfer 根据ID获取连接账户转账。
// 官方文档: https://www.airwallex.com/docs/api/scale/connected_account_transfers/retrieve.md
func (s *Service) GetConnectedAccountTransfer(ctx context.Context, id string, opts ...sdk.RequestOption) (*ConnectedAccountTransfer, error) {
	var resp ConnectedAccountTransfer
	err := s.doer.Do(ctx, "GET", "/api/v1/connected_account_transfers/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListConnectedAccountTransfers lists connected account transfers.
// ListConnectedAccountTransfers 列出连接账户转账。
// 官方文档: https://www.airwallex.com/docs/api/scale/connected_account_transfers/list.md
func (s *Service) ListConnectedAccountTransfers(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[ConnectedAccountTransfer], error) {
	var resp sdk.ListResult[ConnectedAccountTransfer]
	err := s.doer.Do(ctx, "GET", "/api/v1/connected_account_transfers", nil, &resp, opts...)
	return &resp, err
}
