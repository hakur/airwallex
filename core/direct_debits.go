package core

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// DirectDebit represents direct debit information.
// DirectDebit 表示直接扣款信息。
type DirectDebit struct {
	// TransactionID is the unique transaction identifier. Required.
	// TransactionID 交易唯一标识符。必填。
	TransactionID string `json:"transaction_id"`
	// Amount is the amount. Required.
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// Status is the status. Required.
	// Status 状态。必填。
	Status string `json:"status"`
	// CreatedAt is the creation time. Required.
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// GlobalAccountID is the global account ID. Required.
	// GlobalAccountID 全局账户ID。必填。
	GlobalAccountID string `json:"global_account_id"`
	// DebtorName is the debtor name. Optional.
	// DebtorName 债务人名称。可选。
	DebtorName string `json:"debtor_name,omitempty"`
	// StatementRef is the statement reference number. Optional.
	// StatementRef 对账单参考号。可选。
	StatementRef string `json:"statement_ref,omitempty"`
	// MandateID is the related mandate ID. Optional.
	// MandateID 相关授权ID。可选。
	MandateID string `json:"mandate_id,omitempty"`
}

// GetDirectDebit retrieves a direct debit by ID.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/direct_debits/retrieve.md
// GetDirectDebit 根据 ID 获取直接扣款。
func (s *Service) GetDirectDebit(ctx context.Context, id string, opts ...sdk.RequestOption) (*DirectDebit, error) {
	var resp DirectDebit
	err := s.doer.Do(ctx, "GET", "/api/v1/direct_debits/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListDirectDebits lists direct debits.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/direct_debits/list.md
// ListDirectDebits 列出直接扣款。
func (s *Service) ListDirectDebits(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[DirectDebit], error) {
	var resp sdk.ListResult[DirectDebit]
	err := s.doer.Do(ctx, "GET", "/api/v1/direct_debits", nil, &resp, opts...)
	return &resp, err
}

// CancelDirectDebit cancels a direct debit.
// 官方文档: https://www.airwallex.com/docs/api/core_resources/direct_debits/cancel.md
// CancelDirectDebit 取消直接扣款。
func (s *Service) CancelDirectDebit(ctx context.Context, transactionID string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/direct_debits/"+transactionID+"/cancel", nil, nil, opts...)
}
