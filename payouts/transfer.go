package payouts

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// TransferStatus represents a transfer status.
// TransferStatus 转账状态枚举。
type TransferStatus = string

const (
	TransferStatusPending    TransferStatus = "PENDING"
	TransferStatusInitiated  TransferStatus = "INITIATED"
	TransferStatusProcessing TransferStatus = "PROCESSING"
	TransferStatusSent       TransferStatus = "SENT"
	TransferStatusPaid       TransferStatus = "PAID"
	TransferStatusCancelled  TransferStatus = "CANCELLED"
)

// TransferMethod represents a transfer method.
// TransferMethod 转账方式枚举。
type TransferMethod = string

const (
	TransferMethodLocal TransferMethod = "LOCAL"
	TransferMethodSwift TransferMethod = "SWIFT"
	TransferMethodWire  TransferMethod = "WIRE"
)

// TransferPayerAddress represents a transfer payer address.
// TransferPayerAddress 转账付款人地址。
type TransferPayerAddress struct {
	// City 城市。可选。
	City string `json:"city,omitempty"`
	// CountryCode 国家代码。必填。
	CountryCode sdk.CountryCode `json:"country_code"`
	// Postcode 邮政编码。可选。
	Postcode string `json:"postcode,omitempty"`
	// State 州/省。可选。
	State string `json:"state,omitempty"`
	// StreetAddress 街道地址。可选。
	StreetAddress string `json:"street_address,omitempty"`
}

// TransferPayerAdditionalInfo represents additional transfer payer information.
// TransferPayerAdditionalInfo 转账付款人附加信息。
type TransferPayerAdditionalInfo struct {
	// BusinessRegistrationNumber 工商注册号。可选。
	BusinessRegistrationNumber string `json:"business_registration_number,omitempty"`
	// BusinessRegistrationType 工商注册类型。可选。
	BusinessRegistrationType string `json:"business_registration_type,omitempty"`
	// ExternalID 外部ID。可选。
	ExternalID string `json:"external_id,omitempty"`
	// PersonalEmail 个人邮箱。可选。
	PersonalEmail string `json:"personal_email,omitempty"`
	// PersonalIDNumber 个人证件号。可选。
	PersonalIDNumber string `json:"personal_id_number,omitempty"`
}

// TransferPayer represents the payer in a transfer.
// TransferPayer 转账中的付款人信息。
type TransferPayer struct {
	// AdditionalInfo 附加信息。可选。
	AdditionalInfo *TransferPayerAdditionalInfo `json:"additional_info,omitempty"`
	// Address 地址。可选。
	Address *TransferPayerAddress `json:"address,omitempty"`
	// CompanyName 公司名称。可选。
	CompanyName string `json:"company_name,omitempty"`
	// DateOfBirth 出生日期。可选。
	DateOfBirth string `json:"date_of_birth,omitempty"`
	// EntityType 实体类型。必填。
	EntityType string `json:"entity_type"`
	// FirstName 名。可选。
	FirstName string `json:"first_name,omitempty"`
	// LastName 姓。可选。
	LastName string `json:"last_name,omitempty"`
}

// Funding 资金来源信息。
type Funding struct {
	// DepositType 存款类型。可选。
	DepositType string `json:"deposit_type,omitempty"`
	// FailureReason 失败原因。可选。
	FailureReason string `json:"failure_reason,omitempty"`
	// FundingSourceID 资金来源ID。可选。
	FundingSourceID string `json:"funding_source_id,omitempty"`
	// Reference 参考号。可选。
	Reference string `json:"reference,omitempty"`
	// Status 资金状态。必填。
	Status string `json:"status"`
}

// ApplicationFeeOption represents an application fee option.
// ApplicationFeeOption 应用费用选项。
type ApplicationFeeOption struct {
	// Amount 金额。可选。
	Amount string `json:"amount,omitempty"`
	// Currency 货币。可选。
	Currency sdk.Currency `json:"currency,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// Percentage 百分比。可选。
	Percentage string `json:"percentage,omitempty"`
	// SourceType 费用来源类型。可选。
	SourceType string `json:"source_type,omitempty"`
	// Type 费用类型。可选。
	Type string `json:"type,omitempty"`
}

// ApplicationFee represents an application fee.
// ApplicationFee 应用费用。
type ApplicationFee struct {
	// Amount 金额。可选。
	Amount string `json:"amount,omitempty"`
	// Currency 货币。可选。
	Currency sdk.Currency `json:"currency,omitempty"`
	// SourceType 费用来源类型。可选。
	SourceType string `json:"source_type,omitempty"`
}

// Conversion 兑换信息。
type Conversion struct {
	// CurrencyPair 货币对。必填。
	CurrencyPair string `json:"currency_pair"`
	// Rate 汇率。必填。
	Rate float64 `json:"rate"`
}

// DispatchInfo 调度信息。
type DispatchInfo struct {
	// ExternalReference 外部参考号。可选。
	ExternalReference string `json:"external_reference,omitempty"`
	// ExternalReferenceType 外部参考类型。可选。
	ExternalReferenceType string `json:"external_reference_type,omitempty"`
}

// Prepayment 预付款信息。
type Prepayment struct {
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// Currency 货币。必填。
	Currency sdk.Currency `json:"currency"`
}

// Transfer represents a transfer.
// Transfer 表示转账信息。
type Transfer struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// BeneficiaryID 收款人唯一标识符。必填。
	BeneficiaryID string `json:"beneficiary_id"`
	// Beneficiary 收款人信息。可选。
	Beneficiary *BeneficiaryDetails `json:"beneficiary,omitempty"`
	// TransferAmount 转账金额。可选。
	TransferAmount float64 `json:"transfer_amount,omitempty"`
	// TransferCurrency 转账货币。可选。
	TransferCurrency sdk.Currency `json:"transfer_currency,omitempty"`
	// TransferDate 转账日期。可选。
	TransferDate string `json:"transfer_date,omitempty"`
	// TransferMethod 转账方式。可选。
	TransferMethod TransferMethod `json:"transfer_method,omitempty"`
	// SourceAmount 源金额。可选。
	SourceAmount float64 `json:"source_amount,omitempty"`
	// SourceCurrency 源货币。可选。
	SourceCurrency string `json:"source_currency,omitempty"`
	// AmountBeneficiaryReceives 收款人实际收到金额。可选。
	AmountBeneficiaryReceives float64 `json:"amount_beneficiary_receives,omitempty"`
	// AmountPayerPays 付款人支付金额。可选。
	AmountPayerPays float64 `json:"amount_payer_pays,omitempty"`
	// FeeAmount 手续费金额。可选。
	FeeAmount float64 `json:"fee_amount,omitempty"`
	// FeeCurrency 手续费货币。可选。
	FeeCurrency sdk.Currency `json:"fee_currency,omitempty"`
	// FeePaidBy 手续费承担方。可选。
	FeePaidBy string `json:"fee_paid_by,omitempty"`
	// Payer 付款人信息。可选。
	Payer *TransferPayer `json:"payer,omitempty"`
	// Funding 资金来源信息。可选。
	Funding *Funding `json:"funding,omitempty"`
	// Status 转账状态。必填。
	Status TransferStatus `json:"status"`
	// Reference 参考号。可选。
	Reference string `json:"reference,omitempty"`
	// Reason 原因。可选。
	Reason string `json:"reason,omitempty"`
	// Remarks 备注。可选。
	Remarks string `json:"remarks,omitempty"`
	// ShortReferenceID 短参考标识符。可选。
	ShortReferenceID string `json:"short_reference_id,omitempty"`
	// SwiftChargeOption SWIFT费用承担选项。可选。
	SwiftChargeOption string `json:"swift_charge_option,omitempty"`
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
	// ApplicationFeeOptions 应用费用选项列表。可选。
	ApplicationFeeOptions []ApplicationFeeOption `json:"application_fee_options,omitempty"`
	// ApplicationFees 应用费用列表。可选。
	ApplicationFees []ApplicationFee `json:"application_fees,omitempty"`
	// BatchTransferID 批量转账ID。可选。
	BatchTransferID string `json:"batch_transfer_id,omitempty"`
	// Conversion 兑换信息。可选。
	Conversion *Conversion `json:"conversion,omitempty"`
	// DispatchDate 实际调度日期。可选。
	DispatchDate string `json:"dispatch_date,omitempty"`
	// DispatchInfo 调度信息。可选。
	DispatchInfo *DispatchInfo `json:"dispatch_info,omitempty"`
	// FailureReason 失败原因。可选。
	FailureReason string `json:"failure_reason,omitempty"`
	// FailureType 失败类型。可选。
	FailureType string `json:"failure_type,omitempty"`
	// LockRateOnCreate 创建时是否锁定汇率。可选。
	LockRateOnCreate bool `json:"lock_rate_on_create,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// PayerID 付款人ID。可选。
	PayerID string `json:"payer_id,omitempty"`
	// Prepayment 预付款信息。可选。
	Prepayment *Prepayment `json:"prepayment,omitempty"`
}

// CreateTransferRequest is the request to create a transfer.
// CreateTransferRequest 创建转账请求。
type CreateTransferRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// BeneficiaryID 收款人唯一标识符。可选（若提供 beneficiary 对象则可不填）。
	BeneficiaryID string `json:"beneficiary_id,omitempty"`
	// Beneficiary 收款人信息。可选（若提供 beneficiary_id 则可不填）。
	Beneficiary *BeneficiaryInput `json:"beneficiary,omitempty"`
	// TransferAmount 转账金额。可选。
	TransferAmount float64 `json:"transfer_amount,omitempty"`
	// TransferCurrency 转账货币。可选。
	TransferCurrency sdk.Currency `json:"transfer_currency,omitempty"`
	// TransferMethod 转账方式。可选。
	TransferMethod TransferMethod `json:"transfer_method,omitempty"`
	// SourceAmount 源金额。可选。
	SourceAmount float64 `json:"source_amount,omitempty"`
	// SourceCurrency 源货币。可选。
	SourceCurrency string `json:"source_currency,omitempty"`
	// TransferDate 转账日期。可选。
	TransferDate string `json:"transfer_date,omitempty"`
	// Reference 参考号。可选。
	Reference string `json:"reference,omitempty"`
	// Reason 原因。可选。
	Reason string `json:"reason,omitempty"`
	// Remarks 备注。可选。
	Remarks string `json:"remarks,omitempty"`
	// FeePaidBy 手续费承担方。可选。
	FeePaidBy string `json:"fee_paid_by,omitempty"`
	// LockRateOnCreate 创建时是否锁定汇率。可选。
	LockRateOnCreate bool `json:"lock_rate_on_create,omitempty"`
	// SwiftChargeOption SWIFT费用承担选项。可选。
	SwiftChargeOption string `json:"swift_charge_option,omitempty"`
	// ClientData 客户端数据。可选。
	ClientData string `json:"client_data,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// Payer 付款人信息。可选。
	Payer *TransferPayer `json:"payer,omitempty"`
	// PayerID 付款人ID。可选。
	PayerID string `json:"payer_id,omitempty"`
	// QuoteID 报价ID。可选。
	QuoteID string `json:"quote_id,omitempty"`
	// ApplicationFeeOptions 应用费用选项。可选。
	ApplicationFeeOptions []ApplicationFeeOption `json:"application_fee_options,omitempty"`
}

// ValidateTransferRequest is the request to validate a transfer.
// ValidateTransferRequest 验证转账请求。
type ValidateTransferRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// BeneficiaryID 收款人唯一标识符。可选。
	BeneficiaryID string `json:"beneficiary_id,omitempty"`
	// Beneficiary 收款人信息。可选。
	Beneficiary *BeneficiaryInput `json:"beneficiary,omitempty"`
	// TransferAmount 转账金额。可选。
	TransferAmount float64 `json:"transfer_amount,omitempty"`
	// TransferCurrency 转账货币。可选。
	TransferCurrency sdk.Currency `json:"transfer_currency,omitempty"`
	// TransferMethod 转账方式。可选。
	TransferMethod TransferMethod `json:"transfer_method,omitempty"`
	// SourceAmount 源金额。可选。
	SourceAmount float64 `json:"source_amount,omitempty"`
	// SourceCurrency 源货币。可选。
	SourceCurrency string `json:"source_currency,omitempty"`
	// TransferDate 转账日期。可选。
	TransferDate string `json:"transfer_date,omitempty"`
	// Reference 参考号。可选。
	Reference string `json:"reference,omitempty"`
	// Reason 原因。可选。
	Reason string `json:"reason,omitempty"`
	// Remarks 备注。可选。
	Remarks string `json:"remarks,omitempty"`
	// FeePaidBy 手续费承担方。可选。
	FeePaidBy string `json:"fee_paid_by,omitempty"`
	// SwiftChargeOption SWIFT费用承担选项。可选。
	SwiftChargeOption string `json:"swift_charge_option,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// Payer 付款人信息。可选。
	Payer *TransferPayer `json:"payer,omitempty"`
	// PayerID 付款人ID。可选。
	PayerID string `json:"payer_id,omitempty"`
	// QuoteID 报价ID。可选。
	QuoteID string `json:"quote_id,omitempty"`
	// ApplicationFeeOptions 应用费用选项。可选。
	ApplicationFeeOptions []ApplicationFeeOption `json:"application_fee_options,omitempty"`
}

// ConfirmTransferFundingRequest is the request to confirm transfer funding source.
// ConfirmTransferFundingRequest 确认转账资金来源请求。
type ConfirmTransferFundingRequest struct {
	// FundingSourceID 资金来源ID。可选。
	FundingSourceID string `json:"funding_source_id,omitempty"`
	// DepositType 存款类型（DIRECT_DEBIT或FASTER_DIRECT_DEBIT）。可选。
	DepositType string `json:"deposit_type,omitempty"`
	// Reference 用户指定的参考号（显示在Direct Debit银行对账单上）。可选。
	Reference string `json:"reference,omitempty"`
}

// CreateTransfer creates a transfer.
// CreateTransfer 创建转账。
// 官方文档: https://www.airwallex.com/docs/api/payouts/transfers/create.md
func (s *Service) CreateTransfer(ctx context.Context, req *CreateTransferRequest, opts ...sdk.RequestOption) (*Transfer, error) {
	var resp Transfer
	err := s.doer.Do(ctx, "POST", "/api/v1/transfers/create", req, &resp, opts...)
	return &resp, err
}

// GetTransfer retrieves a transfer by ID.
// GetTransfer 根据 ID 获取转账。
// 官方文档: https://www.airwallex.com/docs/api/payouts/transfers/retrieve.md
func (s *Service) GetTransfer(ctx context.Context, id string, opts ...sdk.RequestOption) (*Transfer, error) {
	var resp Transfer
	err := s.doer.Do(ctx, "GET", "/api/v1/transfers/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListTransfers lists transfers.
// ListTransfers 列出转账。
// 官方文档: https://www.airwallex.com/docs/api/payouts/transfers/list.md
func (s *Service) ListTransfers(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[Transfer], error) {
	var resp sdk.ListResult[Transfer]
	err := s.doer.Do(ctx, "GET", "/api/v1/transfers", nil, &resp, opts...)
	return &resp, err
}

// CancelTransfer cancels a transfer.
// CancelTransfer 取消转账。
// 官方文档: https://www.airwallex.com/docs/api/payouts/transfers/cancel.md
func (s *Service) CancelTransfer(ctx context.Context, id string, opts ...sdk.RequestOption) (bool, error) {
	var resp string
	err := s.doer.Do(ctx, "POST", "/api/v1/transfers/"+id+"/cancel", nil, &resp, opts...)
	if err != nil {
		return false, err
	}
	return resp == "OK", nil
}

// ValidateTransfer validates a transfer request.
// ValidateTransfer 验证转账请求。
// 官方文档: https://www.airwallex.com/docs/api/payouts/transfers/validate.md
func (s *Service) ValidateTransfer(ctx context.Context, req *ValidateTransferRequest, opts ...sdk.RequestOption) (bool, error) {
	var resp string
	err := s.doer.Do(ctx, "POST", "/api/v1/transfers/validate", req, &resp, opts...)
	if err != nil {
		return false, err
	}
	return resp == "OK", nil
}

// ConfirmTransferFunding confirms the transfer funding source.
// ConfirmTransferFunding 确认转账资金来源。
// 官方文档: https://www.airwallex.com/docs/api/payouts/transfers/confirm_funding.md
func (s *Service) ConfirmTransferFunding(ctx context.Context, id string, req *ConfirmTransferFundingRequest, opts ...sdk.RequestOption) (*Transfer, error) {
	var resp Transfer
	err := s.doer.Do(ctx, "POST", "/api/v1/transfers/"+id+"/confirm_funding", req, &resp, opts...)
	return &resp, err
}
