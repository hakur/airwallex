// Package events provides typed webhook event structures for the transfers domain.
// Transfers 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/transfers.md
//
// 事件映射表:
//
//	payout.transfer.in_approval             → TransferInApprovalEvent          (Data: TransferEventData)
//	payout.transfer.approval_recalled       → TransferApprovalRecalledEvent    (Data: TransferEventData)
//	payout.transfer.approval_rejected       → TransferApprovalRejectedEvent    (Data: TransferEventData)
//	payout.transfer.approval_blocked        → TransferApprovalBlockedEvent     (Data: TransferEventData)
//	payout.transfer.scheduled               → TransferScheduledEvent           (Data: TransferEventData)
//	payout.transfer.overdue                 → TransferOverdueEvent             (Data: TransferEventData)
//	payout.transfer.processing              → TransferProcessingEvent          (Data: TransferEventData)
//	payout.transfer.sent                    → TransferSentEvent                (Data: TransferEventData)
//	payout.transfer.paid                    → TransferPaidEvent                (Data: TransferEventData)
//	payout.transfer.failed                  → TransferFailedEvent              (Data: TransferEventData)
//	payout.transfer.cancellation_requested  → TransferCancellationRequestedEvent (Data: TransferEventData)
//	payout.transfer.cancelled               → TransferCancelledEvent           (Data: TransferEventData)
//	payout.transfer.funding.requires_funding_confirmation → TransferFundingRequiresFundingConfirmationEvent (Data: TransferEventData)
//	payout.transfer.funding.scheduled       → TransferFundingScheduledEvent    (Data: TransferEventData)
//	payout.transfer.funding.processing      → TransferFundingProcessingEvent   (Data: TransferEventData)
//	payout.transfer.funding.funded          → TransferFundingFundedEvent       (Data: TransferEventData)
//	payout.transfer.funding.failed          → TransferFundingFailedEvent       (Data: TransferEventData)
//	payout.transfer.funding.cancelled       → TransferFundingCancelledEvent    (Data: TransferEventData)
//	payout.transfer.funding.reversed        → TransferFundingReversedEvent     (Data: TransferEventData)
//	payout.batch_transfers.drafting         → BatchTransfersDraftingEvent      (Data: BatchTransferEventData)
//	payout.batch_transfers.in_approval      → BatchTransfersInApprovalEvent    (Data: BatchTransferEventData)
//	payout.batch_transfers.approval_recalled → BatchTransfersApprovalRecalledEvent (Data: BatchTransferEventData)
//	payout.batch_transfers.approval_rejected → BatchTransfersApprovalRejectedEvent (Data: BatchTransferEventData)
//	payout.batch_transfers.approval_blocked → BatchTransfersApprovalBlockedEvent (Data: BatchTransferEventData)
//	payout.batch_transfers.scheduled        → BatchTransfersScheduledEvent     (Data: BatchTransferEventData)
//	payout.batch_transfers.overdue          → BatchTransfersOverdueEvent       (Data: BatchTransferEventData)
//	payout.batch_transfers.booking          → BatchTransfersBookingEvent       (Data: BatchTransferEventData)
//	payout.batch_transfers.partially_booked → BatchTransfersPartiallyBookedEvent (Data: BatchTransferEventData)
//	payout.batch_transfers.booked           → BatchTransfersBookedEvent        (Data: BatchTransferEventData)
//	payout.batch_transfers.failed           → BatchTransfersFailedEvent        (Data: BatchTransferEventData)
//	payout.batch_transfers.cancellation_requested → BatchTransfersCancellationRequestedEvent (Data: BatchTransferEventData)
//	payout.batch_transfers.cancelled        → BatchTransfersCancelledEvent     (Data: BatchTransferEventData)
//	payout.batch_transfers.funding.scheduled → BatchTransfersFundingScheduledEvent (Data: BatchTransferEventData)
//	payout.batch_transfers.funding.processing → BatchTransfersFundingProcessingEvent (Data: BatchTransferEventData)
//	payout.batch_transfers.funding.funded   → BatchTransfersFundingFundedEvent (Data: BatchTransferEventData)
//	payout.batch_transfers.funding.cancelled → BatchTransfersFundingCancelledEvent (Data: BatchTransferEventData)
//	payout.batch_transfers.funding.failed   → BatchTransfersFundingFailedEvent (Data: BatchTransferEventData)
//	payout.batch_transfers.funding.reversed → BatchTransfersFundingReversedEvent (Data: BatchTransferEventData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Transfer Events ---

// TransferInApprovalEvent represents the payout.transfer.in_approval webhook event.
// TransferInApprovalEvent 表示 payout.transfer.in_approval webhook 事件。
type TransferInApprovalEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferApprovalRecalledEvent represents the payout.transfer.approval_recalled webhook event.
// TransferApprovalRecalledEvent 表示 payout.transfer.approval_recalled webhook 事件。
type TransferApprovalRecalledEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferApprovalRejectedEvent represents the payout.transfer.approval_rejected webhook event.
// TransferApprovalRejectedEvent 表示 payout.transfer.approval_rejected webhook 事件。
type TransferApprovalRejectedEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferApprovalBlockedEvent represents the payout.transfer.approval_blocked webhook event.
// TransferApprovalBlockedEvent 表示 payout.transfer.approval_blocked webhook 事件。
type TransferApprovalBlockedEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferScheduledEvent represents the payout.transfer.scheduled webhook event.
// TransferScheduledEvent 表示 payout.transfer.scheduled webhook 事件。
type TransferScheduledEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferOverdueEvent represents the payout.transfer.overdue webhook event.
// TransferOverdueEvent 表示 payout.transfer.overdue webhook 事件。
type TransferOverdueEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferProcessingEvent represents the payout.transfer.processing webhook event.
// TransferProcessingEvent 表示 payout.transfer.processing webhook 事件。
type TransferProcessingEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferSentEvent represents the payout.transfer.sent webhook event.
// TransferSentEvent 表示 payout.transfer.sent webhook 事件。
type TransferSentEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferPaidEvent represents the payout.transfer.paid webhook event.
// TransferPaidEvent 表示 payout.transfer.paid webhook 事件。
type TransferPaidEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferFailedEvent represents the payout.transfer.failed webhook event.
// TransferFailedEvent 表示 payout.transfer.failed webhook 事件。
type TransferFailedEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferCancellationRequestedEvent represents the payout.transfer.cancellation_requested webhook event.
// TransferCancellationRequestedEvent 表示 payout.transfer.cancellation_requested webhook 事件。
type TransferCancellationRequestedEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferCancelledEvent represents the payout.transfer.cancelled webhook event.
// TransferCancelledEvent 表示 payout.transfer.cancelled webhook 事件。
type TransferCancelledEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferFundingRequiresFundingConfirmationEvent represents the payout.transfer.funding.requires_funding_confirmation webhook event.
// TransferFundingRequiresFundingConfirmationEvent 表示 payout.transfer.funding.requires_funding_confirmation webhook 事件。
type TransferFundingRequiresFundingConfirmationEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferFundingScheduledEvent represents the payout.transfer.funding.scheduled webhook event.
// TransferFundingScheduledEvent 表示 payout.transfer.funding.scheduled webhook 事件。
type TransferFundingScheduledEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferFundingProcessingEvent represents the payout.transfer.funding.processing webhook event.
// TransferFundingProcessingEvent 表示 payout.transfer.funding.processing webhook 事件。
type TransferFundingProcessingEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferFundingFundedEvent represents the payout.transfer.funding.funded webhook event.
// TransferFundingFundedEvent 表示 payout.transfer.funding.funded webhook 事件。
type TransferFundingFundedEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferFundingFailedEvent represents the payout.transfer.funding.failed webhook event.
// TransferFundingFailedEvent 表示 payout.transfer.funding.failed webhook 事件。
type TransferFundingFailedEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferFundingCancelledEvent represents the payout.transfer.funding.cancelled webhook event.
// TransferFundingCancelledEvent 表示 payout.transfer.funding.cancelled webhook 事件。
type TransferFundingCancelledEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// TransferFundingReversedEvent represents the payout.transfer.funding.reversed webhook event.
// TransferFundingReversedEvent 表示 payout.transfer.funding.reversed webhook 事件。
type TransferFundingReversedEvent struct {
	Event
	Data TransferEventData `json:"data"`
}

// --- Batch Transfer Events ---

// BatchTransfersDraftingEvent represents the payout.batch_transfers.drafting webhook event.
// BatchTransfersDraftingEvent 表示 payout.batch_transfers.drafting webhook 事件。
type BatchTransfersDraftingEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersInApprovalEvent represents the payout.batch_transfers.in_approval webhook event.
// BatchTransfersInApprovalEvent 表示 payout.batch_transfers.in_approval webhook 事件。
type BatchTransfersInApprovalEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersApprovalRecalledEvent represents the payout.batch_transfers.approval_recalled webhook event.
// BatchTransfersApprovalRecalledEvent 表示 payout.batch_transfers.approval_recalled webhook 事件。
type BatchTransfersApprovalRecalledEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersApprovalRejectedEvent represents the payout.batch_transfers.approval_rejected webhook event.
// BatchTransfersApprovalRejectedEvent 表示 payout.batch_transfers.approval_rejected webhook 事件。
type BatchTransfersApprovalRejectedEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersApprovalBlockedEvent represents the payout.batch_transfers.approval_blocked webhook event.
// BatchTransfersApprovalBlockedEvent 表示 payout.batch_transfers.approval_blocked webhook 事件。
type BatchTransfersApprovalBlockedEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersScheduledEvent represents the payout.batch_transfers.scheduled webhook event.
// BatchTransfersScheduledEvent 表示 payout.batch_transfers.scheduled webhook 事件。
type BatchTransfersScheduledEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersOverdueEvent represents the payout.batch_transfers.overdue webhook event.
// BatchTransfersOverdueEvent 表示 payout.batch_transfers.overdue webhook 事件。
type BatchTransfersOverdueEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersBookingEvent represents the payout.batch_transfers.booking webhook event.
// BatchTransfersBookingEvent 表示 payout.batch_transfers.booking webhook 事件。
type BatchTransfersBookingEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersPartiallyBookedEvent represents the payout.batch_transfers.partially_booked webhook event.
// BatchTransfersPartiallyBookedEvent 表示 payout.batch_transfers.partially_booked webhook 事件。
type BatchTransfersPartiallyBookedEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersBookedEvent represents the payout.batch_transfers.booked webhook event.
// BatchTransfersBookedEvent 表示 payout.batch_transfers.booked webhook 事件。
type BatchTransfersBookedEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersFailedEvent represents the payout.batch_transfers.failed webhook event.
// BatchTransfersFailedEvent 表示 payout.batch_transfers.failed webhook 事件。
type BatchTransfersFailedEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersCancellationRequestedEvent represents the payout.batch_transfers.cancellation_requested webhook event.
// BatchTransfersCancellationRequestedEvent 表示 payout.batch_transfers.cancellation_requested webhook 事件。
type BatchTransfersCancellationRequestedEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersCancelledEvent represents the payout.batch_transfers.cancelled webhook event.
// BatchTransfersCancelledEvent 表示 payout.batch_transfers.cancelled webhook 事件。
type BatchTransfersCancelledEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersFundingScheduledEvent represents the payout.batch_transfers.funding.scheduled webhook event.
// BatchTransfersFundingScheduledEvent 表示 payout.batch_transfers.funding.scheduled webhook 事件。
type BatchTransfersFundingScheduledEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersFundingProcessingEvent represents the payout.batch_transfers.funding.processing webhook event.
// BatchTransfersFundingProcessingEvent 表示 payout.batch_transfers.funding.processing webhook 事件。
type BatchTransfersFundingProcessingEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersFundingFundedEvent represents the payout.batch_transfers.funding.funded webhook event.
// BatchTransfersFundingFundedEvent 表示 payout.batch_transfers.funding.funded webhook 事件。
type BatchTransfersFundingFundedEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersFundingCancelledEvent represents the payout.batch_transfers.funding.cancelled webhook event.
// BatchTransfersFundingCancelledEvent 表示 payout.batch_transfers.funding.cancelled webhook 事件。
type BatchTransfersFundingCancelledEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersFundingFailedEvent represents the payout.batch_transfers.funding.failed webhook event.
// BatchTransfersFundingFailedEvent 表示 payout.batch_transfers.funding.failed webhook 事件。
type BatchTransfersFundingFailedEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// BatchTransfersFundingReversedEvent represents the payout.batch_transfers.funding.reversed webhook event.
// BatchTransfersFundingReversedEvent 表示 payout.batch_transfers.funding.reversed webhook 事件。
type BatchTransfersFundingReversedEvent struct {
	Event
	Data BatchTransferEventData `json:"data"`
}

// --- Transfer Event Data Structures ---

// TransferEventData represents the payload data for transfer webhook events.
// TransferEventData 表示 transfer webhook 事件的载荷数据。
// 官方文档说明 data 为 Get transfer by ID 的响应体。
type TransferEventData struct {
	AmountBeneficiaryReceives float64              `json:"amount_beneficiary_receives"`
	AmountPayerPays           float64              `json:"amount_payer_pays"`
	Beneficiary               *TransferBeneficiary `json:"beneficiary,omitempty"`
	Conversion                *TransferConversion  `json:"conversion,omitempty"`
	CreatedAt                 string               `json:"created_at"`
	FeeAmount                 float64              `json:"fee_amount"`
	FeeCurrency               string               `json:"fee_currency"`
	FeePaidBy                 string               `json:"fee_paid_by"`
	Funding                   *TransferFunding     `json:"funding,omitempty"`
	ID                        string               `json:"id"`
	LockRateOnCreate          bool                 `json:"lock_rate_on_create"`
	Payer                     *TransferPayer       `json:"payer,omitempty"`
	Prepayment                *TransferPrepayment  `json:"prepayment,omitempty"`
	Reason                    string               `json:"reason"`
	Reference                 string               `json:"reference"`
	Remarks                   string               `json:"remarks"`
	RequestID                 string               `json:"request_id"`
	ShortReferenceID          string               `json:"short_reference_id"`
	SourceCurrency            string               `json:"source_currency"`
	Status                    string               `json:"status"`
	SwiftChargeOption         string               `json:"swift_charge_option"`
	TransferAmount            float64              `json:"transfer_amount"`
	TransferCurrency          string               `json:"transfer_currency"`
	TransferDate              string               `json:"transfer_date"`
	TransferMethod            string               `json:"transfer_method"`
	UpdatedAt                 string               `json:"updated_at"`
}

// TransferBeneficiary represents the beneficiary information in a transfer.
// TransferBeneficiary 表示转账中的收款人信息。
type TransferBeneficiary struct {
	AdditionalInfo *TransferBeneficiaryAdditionalInfo `json:"additional_info,omitempty"`
	Address        *TransferAddress                   `json:"address,omitempty"`
	BankDetails    *TransferBankDetails               `json:"bank_details,omitempty"`
	CompanyName    string                             `json:"company_name"`
	EntityType     string                             `json:"entity_type"`
	Type           string                             `json:"type"`
}

// TransferBeneficiaryAdditionalInfo represents additional info for a transfer beneficiary.
// TransferBeneficiaryAdditionalInfo 表示转账收款人的附加信息。
type TransferBeneficiaryAdditionalInfo struct {
	PersonalEmail string `json:"personal_email"`
}

// TransferAddress represents an address in a transfer payload.
// TransferAddress 表示转账 payload 中的地址信息。
type TransferAddress struct {
	City          string `json:"city"`
	CountryCode   string `json:"country_code"`
	Postcode      string `json:"postcode"`
	State         string `json:"state"`
	StreetAddress string `json:"street_address"`
}

// TransferBankDetails represents bank account details in a transfer payload.
// TransferBankDetails 表示转账 payload 中的银行账户详情。
type TransferBankDetails struct {
	AccountCurrency      string `json:"account_currency"`
	AccountName          string `json:"account_name"`
	AccountNumber        string `json:"account_number"`
	AccountRoutingType1  string `json:"account_routing_type1"`
	AccountRoutingValue1 string `json:"account_routing_value1"`
	BankAccountCategory  string `json:"bank_account_category"`
	BankCountryCode      string `json:"bank_country_code"`
	BankName             string `json:"bank_name"`
	LocalClearingSystem  string `json:"local_clearing_system"`
}

// TransferConversion represents the currency conversion details in a transfer.
// TransferConversion 表示转账中的货币兑换详情。
type TransferConversion struct {
	CurrencyPair string  `json:"currency_pair"`
	Rate         float64 `json:"rate"`
}

// TransferFunding represents the funding status in a transfer.
// TransferFunding 表示转账中的资金状态。
type TransferFunding struct {
	Status string `json:"status"`
}

// TransferPayer represents the payer information in a transfer.
// TransferPayer 表示转账中的付款人信息。
type TransferPayer struct {
	AdditionalInfo *TransferPayerAdditionalInfo `json:"additional_info,omitempty"`
	Address        *TransferAddress             `json:"address,omitempty"`
	CompanyName    string                       `json:"company_name"`
	EntityType     string                       `json:"entity_type"`
}

// TransferPayerAdditionalInfo represents additional info for a transfer payer.
// TransferPayerAdditionalInfo 表示转账付款人的附加信息。
type TransferPayerAdditionalInfo struct {
	BusinessIncorporationDate  string `json:"business_incorporation_date"`
	BusinessRegistrationNumber string `json:"business_registration_number"`
	BusinessRegistrationType   string `json:"business_registration_type"`
}

// TransferPrepayment represents the prepayment details in a transfer.
// TransferPrepayment 表示转账中的预付款详情。
type TransferPrepayment struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// --- Batch Transfer Event Data Structures ---

// BatchTransferEventData represents the payload data for batch transfer webhook events.
// BatchTransferEventData 表示 batch transfer webhook 事件的载荷数据。
type BatchTransferEventData struct {
	Funding          *BatchTransferFunding      `json:"funding,omitempty"`
	ID               string                     `json:"id"`
	Metadata         *BatchTransferMetadata     `json:"metadata,omitempty"`
	Name             string                     `json:"name"`
	QuoteSummary     *BatchTransferQuoteSummary `json:"quote_summary,omitempty"`
	Remarks          string                     `json:"remarks"`
	RequestID        string                     `json:"request_id"`
	ShortReferenceID string                     `json:"short_reference_id"`
	Status           string                     `json:"status"`
	TotalItemCount   int                        `json:"total_item_count"`
	TransferDate     string                     `json:"transfer_date"`
	UpdatedAt        string                     `json:"updated_at"`
	ValidItemCount   int                        `json:"valid_item_count"`
}

// BatchTransferFunding represents the funding details in a batch transfer.
// BatchTransferFunding 表示批量转账中的资金详情。
type BatchTransferFunding struct {
	DepositType     string `json:"deposit_type"`
	FundingSourceID string `json:"funding_source_id"`
	Status          string `json:"status"`
}

// BatchTransferMetadata represents the metadata in a batch transfer.
// BatchTransferMetadata 表示批量转账中的元数据。
type BatchTransferMetadata struct {
	ID string `json:"id"`
}

// BatchTransferQuoteSummary represents the quote summary in a batch transfer.
// BatchTransferQuoteSummary 表示批量转账中的报价摘要。
type BatchTransferQuoteSummary struct {
	ExpiresAt    string               `json:"expires_at"`
	LastQuotedAt string               `json:"last_quoted_at"`
	Quotes       []BatchTransferQuote `json:"quotes"`
	Validity     string               `json:"validity"`
}

// BatchTransferQuote represents an individual quote in a batch transfer quote summary.
// BatchTransferQuote 表示批量转账报价摘要中的单个报价。
type BatchTransferQuote struct {
	AmountBeneficiaryReceives float64 `json:"amount_beneficiary_receives"`
	AmountPayerPays           float64 `json:"amount_payer_pays"`
	ClientRate                float64 `json:"client_rate,omitempty"`
	CurrencyPair              string  `json:"currency_pair,omitempty"`
	FeeAmount                 float64 `json:"fee_amount"`
	FeeCurrency               string  `json:"fee_currency"`
	TransferCurrency          string  `json:"transfer_currency"`
	SourceCurrency            string  `json:"source_currency"`
}
