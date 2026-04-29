package simulation

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// GlobalAccountDepositStatus represents the global account deposit status.
// GlobalAccountDepositStatus 全球账户存款状态。
type GlobalAccountDepositStatus = string

const (
	// GlobalAccountDepositStatusPending is the pending status.
	GlobalAccountDepositStatusPending GlobalAccountDepositStatus = "PENDING"
	// GlobalAccountDepositStatusSettled is the settled status.
	GlobalAccountDepositStatusSettled GlobalAccountDepositStatus = "SETTLED"
	// GlobalAccountDepositStatusRejected is the rejected status.
	GlobalAccountDepositStatusRejected GlobalAccountDepositStatus = "REJECTED"
)

// DepositType represents the deposit type.
// DepositType 存款类型。
type DepositType = string

const (
	// DepositTypeDebit is a debit deposit type.
	DepositTypeDebit DepositType = "DEBIT"
	// DepositTypeCredit is a credit deposit type.
	DepositTypeCredit DepositType = "CREDIT"
)

// GlobalAccountDeposit represents a global account deposit response.
// GlobalAccountDeposit 表示全球账户存款响应。
type GlobalAccountDeposit struct {
	// ID is the unique deposit identifier. Optional.
	// ID 存款唯一标识符。可选。
	ID string `json:"id,omitempty"`
	// Amount is the deposit amount. Optional.
	// Amount 金额。可选。
	Amount float64 `json:"amount,omitempty"`
	// Currency is the currency code. Optional.
	// Currency 币种。可选。
	Currency string `json:"currency,omitempty"`
	// Status is the deposit status. Optional.
	// Status 状态。可选。
	Status GlobalAccountDepositStatus `json:"status,omitempty"`
	// DepositType is the deposit type. Optional.
	// DepositType 存款类型。可选。
	DepositType DepositType `json:"deposit_type,omitempty"`
	// CreateTime is the creation time. Optional.
	// CreateTime 创建时间。可选。
	CreateTime int64 `json:"create_time,omitempty"`
	// PayerName is the payer name. Optional.
	// PayerName 付款人名称。可选。
	PayerName string `json:"payer_name,omitempty"`
	// PayerCountry is the payer's country. Optional.
	// PayerCountry 付款人国家。可选。
	PayerCountry string `json:"payer_country,omitempty"`
	// Reference is the reference number. Optional.
	// Reference 引用编号。可选。
	Reference string `json:"reference,omitempty"`
	// StatementRef is the bank statement reference. Optional.
	// StatementRef 银行对账单引用。可选。
	StatementRef string `json:"statement_ref,omitempty"`
	// FeeAmount is the fee amount. Optional.
	// FeeAmount 手续费金额。可选。
	FeeAmount float64 `json:"fee_amount,omitempty"`
	// FeeCurrency is the fee currency details. Optional.
	// FeeCurrency 手续费币种详情。可选。
	FeeCurrency *CurrencyDetail `json:"fee_currency,omitempty"`
}

// CurrencyDetail represents currency details.
// CurrencyDetail 表示货币详情。
// CurrencyCode is the currency code.
// CurrencyCode 货币代码。
type CurrencyDetail struct {
	CurrencyCode string `json:"currency_code,omitempty"`
	// Precision is the currency precision.
	// Precision 货币精度。
	Precision int32 `json:"precision,omitempty"`
	// Value is the currency value.
	// Value 货币值。
	Value string `json:"value,omitempty"`
}

// SimulateGlobalAccountDepositRequest represents a request to simulate a global account deposit.
// SimulateGlobalAccountDepositRequest 模拟全球账户存款请求。
type SimulateGlobalAccountDepositRequest struct {
	// GlobalAccountID is the global account ID. Required.
	GlobalAccountID string `json:"global_account_id"`
	// Amount is the deposit amount. Required.
	Amount float64 `json:"amount"`
	// PayerBankName is the payer's bank name. Optional.
	PayerBankName string `json:"payer_bankname,omitempty"`
	// PayerCountry is the payer's country. Optional.
	PayerCountry string `json:"payer_country,omitempty"`
	// PayerName is the payer name. Optional.
	PayerName string `json:"payer_name,omitempty"`
	// Reference is the reference number. Optional.
	Reference string `json:"reference,omitempty"`
	// StatementRef is the bank statement reference. Optional.
	StatementRef string `json:"statement_ref,omitempty"`
	// Status is the deposit status. Optional.
	Status GlobalAccountDepositStatus `json:"status,omitempty"`
}

// SimulateGlobalAccountDeposit simulates a global account deposit.
// SimulateGlobalAccountDeposit 模拟全球账户存款。
// 官方文档: https://www.airwallex.com/docs/api/simulation/deposits/create.md
func (s *Service) SimulateGlobalAccountDeposit(ctx context.Context, req *SimulateGlobalAccountDepositRequest, opts ...sdk.RequestOption) (*GlobalAccountDeposit, error) {
	var resp GlobalAccountDeposit
	err := s.doer.Do(ctx, "POST", "/api/v1/simulation/deposit/create", req, &resp, opts...)
	return &resp, err
}

// SimulateDirectDebitReject simulates a direct debit rejection.
// SimulateDirectDebitReject 模拟直接借记拒绝。
// 官方文档: https://www.airwallex.com/docs/api/simulation/deposits/reject.md
func (s *Service) SimulateDirectDebitReject(ctx context.Context, depositID string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/deposits/"+depositID+"/reject", nil, nil, opts...)
}

// SimulateDirectDebitReverse simulates a direct debit reversal.
// SimulateDirectDebitReverse 模拟直接借记冲正。
// 官方文档: https://www.airwallex.com/docs/api/simulation/deposits/reverse.md
func (s *Service) SimulateDirectDebitReverse(ctx context.Context, depositID string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/deposits/"+depositID+"/reverse", nil, nil, opts...)
}

// SimulateDirectDebitSettle simulates a direct debit settlement.
// SimulateDirectDebitSettle 模拟直接借记结算。
// 官方文档: https://www.airwallex.com/docs/api/simulation/deposits/settle.md
func (s *Service) SimulateDirectDebitSettle(ctx context.Context, depositID string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/deposits/"+depositID+"/settle", nil, nil, opts...)
}
