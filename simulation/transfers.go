package simulation

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// SimulatedTransfer represents a simulated transfer response.
// SimulatedTransfer 表示模拟转账响应。
type SimulatedTransfer struct {
	ID                        string                            `json:"id,omitempty"`
	Status                    string                            `json:"status,omitempty"`
	RequestID                 string                            `json:"request_id,omitempty"`
	SourceAmount              float64                           `json:"source_amount,omitempty"`
	SourceCurrency            string                            `json:"source_currency,omitempty"`
	TransferAmount            float64                           `json:"transfer_amount,omitempty"`
	TransferCurrency          string                            `json:"transfer_currency,omitempty"`
	TransferDate              string                            `json:"transfer_date,omitempty"`
	TransferMethod            string                            `json:"transfer_method,omitempty"`
	Reason                    string                            `json:"reason,omitempty"`
	Reference                 string                            `json:"reference,omitempty"`
	Remarks                   string                            `json:"remarks,omitempty"`
	ShortReferenceID          string                            `json:"short_reference_id,omitempty"`
	CreatedAt                 string                            `json:"created_at,omitempty"`
	UpdatedAt                 string                            `json:"updated_at,omitempty"`
	FailureReason             string                            `json:"failure_reason,omitempty"`
	FailureType               string                            `json:"failure_type,omitempty"`
	AmountBeneficiaryReceives float64                           `json:"amount_beneficiary_receives,omitempty"`
	AmountPayerPays           float64                           `json:"amount_payer_pays,omitempty"`
	FeeAmount                 float64                           `json:"fee_amount,omitempty"`
	FeeCurrency               string                            `json:"fee_currency,omitempty"`
	FeePaidBy                 string                            `json:"fee_paid_by,omitempty"`
	BatchTransferID           string                            `json:"batch_transfer_id,omitempty"`
	BeneficiaryID             string                            `json:"beneficiary_id,omitempty"`
	PayerID                   string                            `json:"payer_id,omitempty"`
	LockRateOnCreate          bool                              `json:"lock_rate_on_create,omitempty"`
	SwiftChargeOption         string                            `json:"swift_charge_option,omitempty"`
	DispatchDate              string                            `json:"dispatch_date,omitempty"`
	Metadata                  map[string]any                    `json:"metadata,omitempty"`
	Beneficiary               *SimulatedTransferBeneficiary     `json:"beneficiary,omitempty"`
	Payer                     *SimulatedTransferPayer           `json:"payer,omitempty"`
	Funding                   *SimulatedTransferFunding         `json:"funding,omitempty"`
	Conversion                *SimulatedTransferConversion      `json:"conversion,omitempty"`
	Prepayment                *SimulatedTransferPrepayment      `json:"prepayment,omitempty"`
	DispatchInfo              *SimulatedTransferDispatchInfo    `json:"dispatch_info,omitempty"`
	ApplicationFees           []SimulatedTransferApplicationFee `json:"application_fees,omitempty"`
}

// SimulatedTransferBeneficiary represents beneficiary information in a simulated transfer response.
// SimulatedTransferBeneficiary 模拟转账响应中的收款人信息。
type SimulatedTransferBeneficiary struct {
	Type           string                          `json:"type,omitempty"`
	FirstName      string                          `json:"first_name,omitempty"`
	LastName       string                          `json:"last_name,omitempty"`
	CompanyName    string                          `json:"company_name,omitempty"`
	DateOfBirth    string                          `json:"date_of_birth,omitempty"`
	EntityType     string                          `json:"entity_type,omitempty"`
	Address        *SimulatedTransferAddress       `json:"address,omitempty"`
	BankDetails    *SimulatedTransferBankDetails   `json:"bank_details,omitempty"`
	DigitalWallet  *SimulatedTransferDigitalWallet `json:"digital_wallet,omitempty"`
	AdditionalInfo map[string]any                  `json:"additional_info,omitempty"`
}

// SimulatedTransferAddress represents address information.
// SimulatedTransferAddress 地址信息。
type SimulatedTransferAddress struct {
	City          string `json:"city,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	Postcode      string `json:"postcode,omitempty"`
	State         string `json:"state,omitempty"`
	StreetAddress string `json:"street_address,omitempty"`
}

// SimulatedTransferBankDetails represents bank account details.
// SimulatedTransferBankDetails 银行账户详情。
type SimulatedTransferBankDetails struct {
	AccountCurrency      string `json:"account_currency,omitempty"`
	AccountName          string `json:"account_name,omitempty"`
	AccountNumber        string `json:"account_number,omitempty"`
	AccountRoutingType1  string `json:"account_routing_type1,omitempty"`
	AccountRoutingValue1 string `json:"account_routing_value1,omitempty"`
	BankCountryCode      string `json:"bank_country_code,omitempty"`
	BankName             string `json:"bank_name,omitempty"`
	Iban                 string `json:"iban,omitempty"`
	SwiftCode            string `json:"swift_code,omitempty"`
	LocalClearingSystem  string `json:"local_clearing_system,omitempty"`
	BankAccountCategory  string `json:"bank_account_category,omitempty"`
}

// SimulatedTransferDigitalWallet represents digital wallet details.
// SimulatedTransferDigitalWallet 数字钱包详情。
type SimulatedTransferDigitalWallet struct {
	AccountName string `json:"account_name,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	IDType      string `json:"id_type,omitempty"`
	IDValue     string `json:"id_value,omitempty"`
	Provider    string `json:"provider,omitempty"`
}

// SimulatedTransferPayer represents payer information in a simulated transfer response.
// SimulatedTransferPayer 模拟转账响应中的付款人信息。
type SimulatedTransferPayer struct {
	EntityType     string                    `json:"entity_type,omitempty"`
	FirstName      string                    `json:"first_name,omitempty"`
	LastName       string                    `json:"last_name,omitempty"`
	CompanyName    string                    `json:"company_name,omitempty"`
	DateOfBirth    string                    `json:"date_of_birth,omitempty"`
	Address        *SimulatedTransferAddress `json:"address,omitempty"`
	AdditionalInfo map[string]any            `json:"additional_info,omitempty"`
}

// SimulatedTransferFunding represents funding source information.
// SimulatedTransferFunding 资金来源信息。
type SimulatedTransferFunding struct {
	DepositType     string `json:"deposit_type,omitempty"`
	FailureReason   string `json:"failure_reason,omitempty"`
	FundingSourceID string `json:"funding_source_id,omitempty"`
	Status          string `json:"status,omitempty"`
}

// SimulatedTransferConversion represents conversion information.
// SimulatedTransferConversion 兑换信息。
type SimulatedTransferConversion struct {
	CurrencyPair string  `json:"currency_pair,omitempty"`
	Rate         float64 `json:"rate,omitempty"`
}

// SimulatedTransferPrepayment represents prepayment information.
// SimulatedTransferPrepayment 预付款信息。
type SimulatedTransferPrepayment struct {
	Amount   float64 `json:"amount,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

// SimulatedTransferDispatchInfo represents dispatch information.
// SimulatedTransferDispatchInfo 调度信息。
type SimulatedTransferDispatchInfo struct {
	ExternalReference     string `json:"external_reference,omitempty"`
	ExternalReferenceType string `json:"external_reference_type,omitempty"`
}

// SimulatedTransferApplicationFee represents application fee information.
// SimulatedTransferApplicationFee 应用费用信息。
type SimulatedTransferApplicationFee struct {
	Amount     string `json:"amount,omitempty"`
	Currency   string `json:"currency,omitempty"`
	SourceType string `json:"source_type,omitempty"`
}

// TransferStatus represents the transfer status enum.
// TransferStatus 转账状态枚举。
type TransferStatus = string

const (
	TransferStatusOverdue    TransferStatus = "OVERDUE"
	TransferStatusProcessing TransferStatus = "PROCESSING"
	TransferStatusSent       TransferStatus = "SENT"
	TransferStatusPaid       TransferStatus = "PAID"
	TransferStatusFailed     TransferStatus = "FAILED"
	TransferStatusCancelled  TransferStatus = "CANCELLED"
)

// SimulateTransferTransitionRequest represents a request to transition a transfer status.
// SimulateTransferTransitionRequest 模拟转账状态转换请求。
type SimulateTransferTransitionRequest struct {
	// NextStatus is the target status. Required.
	// NextStatus 目标状态。必填。
	NextStatus TransferStatus `json:"next_status"`
	// FailureType is the failure type. Optional.
	// FailureType 失败类型。可选。
	FailureType string `json:"failure_type,omitempty"`
}

// SimulateTransferTransition simulates a transfer status transition.
// SimulateTransferTransition 模拟转账状态转换。
// 官方文档: https://www.airwallex.com/docs/api/simulation/transfers/transition.md
func (s *Service) SimulateTransferTransition(ctx context.Context, id string, req *SimulateTransferTransitionRequest, opts ...sdk.RequestOption) (*SimulatedTransfer, error) {
	var resp SimulatedTransfer
	err := s.doer.Do(ctx, "POST", "/api/v1/simulation/transfers/"+id+"/transition", req, &resp, opts...)
	return &resp, err
}
