package finance

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// FinancialTransaction represents financial transaction information.
// FinancialTransaction 表示财务交易信息。
type FinancialTransaction struct {
	// Amount is the transaction amount. Optional.
	// Amount 交易金额。可选。
	Amount float64 `json:"amount,omitempty"`
	// BatchID is the settlement batch ID. Optional.
	// BatchID 结算批次ID。可选。
	BatchID string `json:"batch_id,omitempty"`
	// ClientRate is the client exchange rate. Optional.
	// ClientRate 客户汇率。可选。
	ClientRate float64 `json:"client_rate,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// Currency is the currency code. Optional.
	// Currency 货币代码。可选。
	Currency sdk.Currency `json:"currency,omitempty"`
	// CurrencyPair is the currency pair. Optional.
	// CurrencyPair 货币对。可选。
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Description is the description. Optional.
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// EstimatedSettledAt is the estimated settlement time. Optional.
	// EstimatedSettledAt 预计结算时间。可选。
	EstimatedSettledAt string `json:"estimated_settled_at,omitempty"`
	// Fee is the transaction fee. Optional.
	// Fee 手续费。可选。
	Fee float64 `json:"fee,omitempty"`
	// FundingSourceID is the funding source ID. Optional.
	// FundingSourceID 资金来源ID。可选。
	FundingSourceID string `json:"funding_source_id,omitempty"`
	// ID is the unique transaction identifier. Optional.
	// ID 交易唯一标识符。可选。
	ID string `json:"id,omitempty"`
	// Net is the net amount. Optional.
	// Net 净额。可选。
	Net float64 `json:"net,omitempty"`
	// SettledAt is the actual settlement time. Optional.
	// SettledAt 实际结算时间。可选。
	SettledAt string `json:"settled_at,omitempty"`
	// SourceID is the source transaction ID. Optional.
	// SourceID 源交易ID。可选。
	SourceID string `json:"source_id,omitempty"`
	// SourceType is the source transaction type. Optional.
	// SourceType 源交易类型。可选。
	SourceType string `json:"source_type,omitempty"`
	// Status is the transaction status. Optional.
	// Status 交易状态。可选。
	Status string `json:"status,omitempty"`
	// TransactionType is the transaction type. Optional.
	TransactionType string `json:"transaction_type,omitempty"`
}

// ListFinancialTransactions lists financial transactions.
// ListFinancialTransactions 列出财务交易。
// 官方文档: https://www.airwallex.com/docs/api/finance/financial_transactions/list.md
func (s *Service) ListFinancialTransactions(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[FinancialTransaction], error) {
	var resp sdk.ListResult[FinancialTransaction]
	err := s.doer.Do(ctx, "GET", "/api/v1/financial_transactions", nil, &resp, opts...)
	return &resp, err
}

// GetFinancialTransaction retrieves a financial transaction by ID.
// GetFinancialTransaction 根据 ID 获取财务交易。
// 官方文档: https://www.airwallex.com/docs/api/finance/financial_transactions/retrieve.md
func (s *Service) GetFinancialTransaction(ctx context.Context, id string, opts ...sdk.RequestOption) (*FinancialTransaction, error) {
	var resp FinancialTransaction
	err := s.doer.Do(ctx, "GET", "/api/v1/financial_transactions/"+id, nil, &resp, opts...)
	return &resp, err
}
