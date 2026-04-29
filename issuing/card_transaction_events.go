package issuing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// CardTransactionEventType represents the card transaction event type.
// CardTransactionEventType 卡片交易事件类型。
type CardTransactionEventType = string

const (
	// CardTransactionEventTypeAuthorization is an authorization event.
	// CardTransactionEventTypeAuthorization 授权事件
	CardTransactionEventTypeAuthorization CardTransactionEventType = "AUTHORIZATION"
	// CardTransactionEventTypeClearing is a clearing event.
	// CardTransactionEventTypeClearing 清算事件
	CardTransactionEventTypeClearing CardTransactionEventType = "CLEARING"
	// CardTransactionEventTypeReversalAuth is a reversal authorization event.
	// CardTransactionEventTypeReversalAuth 逆向授权事件
	CardTransactionEventTypeReversalAuth CardTransactionEventType = "REVERSAL_AUTH"
)

// CardTransactionEventSubtype represents the card transaction event subtype.
// CardTransactionEventSubtype 卡片交易事件子类型。
type CardTransactionEventSubtype = string

const (
	// CardTransactionEventSubtypeAuthorization is an authorization subtype.
	// CardTransactionEventSubtypeAuthorization 授权
	CardTransactionEventSubtypeAuthorization CardTransactionEventSubtype = "AUTHORIZATION"
	// CardTransactionEventSubtypeIncrementalAuthorization is an incremental authorization subtype.
	// CardTransactionEventSubtypeIncrementalAuthorization 增量授权
	CardTransactionEventSubtypeIncrementalAuthorization CardTransactionEventSubtype = "INCREMENTAL_AUTHORIZATION"
	// CardTransactionEventSubtypeVerification is a verification subtype.
	// CardTransactionEventSubtypeVerification 验证
	CardTransactionEventSubtypeVerification CardTransactionEventSubtype = "VERIFICATION"
	// CardTransactionEventSubtypeReversal is a reversal subtype.
	// CardTransactionEventSubtypeReversal 撤销
	CardTransactionEventSubtypeReversal CardTransactionEventSubtype = "REVERSAL"
	// CardTransactionEventSubtypePartialReversal is a partial reversal subtype.
	// CardTransactionEventSubtypePartialReversal 部分撤销
	CardTransactionEventSubtypePartialReversal CardTransactionEventSubtype = "PARTIAL_REVERSAL"
	// CardTransactionEventSubtypeExpiredAuthorization is an expired authorization subtype.
	// CardTransactionEventSubtypeExpiredAuthorization 过期授权
	CardTransactionEventSubtypeExpiredAuthorization CardTransactionEventSubtype = "EXPIRED_AUTHORIZATION"
	// CardTransactionEventSubtypeClearing is a clearing subtype.
	// CardTransactionEventSubtypeClearing 清算
	CardTransactionEventSubtypeClearing CardTransactionEventSubtype = "CLEARING"
	// CardTransactionEventSubtypePartialClearing is a partial clearing subtype.
	// CardTransactionEventSubtypePartialClearing 部分清算
	CardTransactionEventSubtypePartialClearing CardTransactionEventSubtype = "PARTIAL_CLEARING"
	// CardTransactionEventSubtypeClearingReversal is a clearing reversal subtype.
	// CardTransactionEventSubtypeClearingReversal 清算撤销
	CardTransactionEventSubtypeClearingReversal CardTransactionEventSubtype = "CLEARING_REVERSAL"
)

// CardTransactionEventProcessResult represents the card transaction event process result.
// CardTransactionEventProcessResult 卡片交易事件处理结果。
type CardTransactionEventProcessResult = string

const (
	// CardTransactionEventProcessResultApproved indicates the event was approved.
	// CardTransactionEventProcessResultApproved 批准
	CardTransactionEventProcessResultApproved CardTransactionEventProcessResult = "APPROVED"
	// CardTransactionEventProcessResultDeclined indicates the event was declined.
	// CardTransactionEventProcessResultDeclined 拒绝
	CardTransactionEventProcessResultDeclined CardTransactionEventProcessResult = "DECLINED"
	// CardTransactionEventProcessResultPending indicates the event is pending.
	// CardTransactionEventProcessResultPending 待处理
	CardTransactionEventProcessResultPending CardTransactionEventProcessResult = "PENDING"
)

// CardTransactionEvent represents card transaction event information.
// CardTransactionEvent 表示卡片交易事件信息。
type CardTransactionEvent struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Type is the event type. Required.
	// Type 事件类型。必填。
	Type CardTransactionEventType `json:"type"`
	// ProcessResult is the process result. Required.
	// ProcessResult 处理结果。必填。
	ProcessResult CardTransactionEventProcessResult `json:"process_result"`
	// AcquiringInstitutionIdentifier is the acquiring institution identifier. Optional.
	// AcquiringInstitutionIdentifier 收单机构标识符。可选。
	AcquiringInstitutionIdentifier string `json:"acquiring_institution_identifier,omitempty"`
	// AuthCode is the authorization code. Optional.
	// AuthCode 授权码。可选。
	AuthCode string `json:"auth_code,omitempty"`
	// BillingAmount is the billing amount. Optional.
	// BillingAmount 账单金额。可选。
	BillingAmount float64 `json:"billing_amount,omitempty"`
	// BillingCurrency is the billing currency. Optional.
	// BillingCurrency 账单币种。可选。
	BillingCurrency sdk.Currency `json:"billing_currency,omitempty"`
	// CardTransactionID is the card transaction ID. Optional.
	// CardTransactionID 卡片交易ID。可选。
	CardTransactionID string `json:"card_transaction_id,omitempty"`
	// ConversionDetails contains conversion details. Optional.
	// ConversionDetails 转换详情。可选。
	ConversionDetails map[string]any `json:"conversion_details,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// DigitalWalletToken contains digital wallet token information. Optional.
	// DigitalWalletToken 数字钱包令牌信息。可选。
	DigitalWalletToken map[string]any `json:"digital_wallet_token,omitempty"`
	// ExpiryDate is the expiry date. Optional.
	// ExpiryDate 到期日期。可选。
	ExpiryDate string `json:"expiry_date,omitempty"`
	// FailureReason is the failure reason. Optional.
	// FailureReason 失败原因。可选。
	FailureReason string `json:"failure_reason,omitempty"`
	// FeeDetails contains fee details. Optional.
	// FeeDetails 费用详情。可选。
	FeeDetails []map[string]any `json:"fee_details,omitempty"`
	// LifecycleID is the lifecycle ID. Optional.
	// LifecycleID 生命周期ID。可选。
	LifecycleID string `json:"lifecycle_id,omitempty"`
	// MaskedCardNumber is the masked card number. Optional.
	// MaskedCardNumber 掩码卡号。可选。
	MaskedCardNumber string `json:"masked_card_number,omitempty"`
	// Merchant contains merchant information. Optional.
	// Merchant 商户信息。可选。
	Merchant *Merchant `json:"merchant,omitempty"`
	// NetworkTransactionID is the network transaction ID. Optional.
	// NetworkTransactionID 网络交易ID。可选。
	NetworkTransactionID string `json:"network_transaction_id,omitempty"`
	// PosEntryMode is the POS entry mode. Optional.
	// PosEntryMode POS输入模式。可选。
	PosEntryMode string `json:"pos_entry_mode,omitempty"`
	// RetrievalRef is the retrieval reference number. Optional.
	// RetrievalRef 检索参考号。可选。
	RetrievalRef string `json:"retrieval_ref,omitempty"`
	// RiskDetails contains risk details. Optional.
	// RiskDetails 风险详情。可选。
	RiskDetails map[string]any `json:"risk_details,omitempty"`
	// Subtype is the event subtype. Optional.
	// Subtype 子类型。可选。
	Subtype CardTransactionEventSubtype `json:"subtype,omitempty"`
	// TransactionAmount is the transaction amount. Optional.
	// TransactionAmount 交易金额。可选。
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
	// TransactionCategory is the transaction category. Optional.
	// TransactionCategory 交易类别。可选。
	TransactionCategory string `json:"transaction_category,omitempty"`
	// TransactionCurrency is the transaction currency. Optional.
	// TransactionCurrency 交易币种。可选。
	TransactionCurrency sdk.Currency `json:"transaction_currency,omitempty"`
	// TransactionDateTime is the transaction date and time. Optional.
	// TransactionDateTime 交易日期时间。可选。
	TransactionDateTime string `json:"transaction_date_time,omitempty"`
	// UpdatedAt is the update time. Optional.
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
}

// ListCardTransactionEvents lists card transaction events.
// 官方文档: https://www.airwallex.com/docs/api/issuing/card_transaction_events/list.md
// ListCardTransactionEvents 列出卡片交易事件。
func (s *Service) ListCardTransactionEvents(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[CardTransactionEvent], error) {
	var resp sdk.ListResult[CardTransactionEvent]
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/card_transaction_events", nil, &resp, opts...)
	return &resp, err
}

// GetCardTransactionEvent retrieves a card transaction event by ID.
// 官方文档: https://www.airwallex.com/docs/api/issuing/card_transaction_events/retrieve.md
// GetCardTransactionEvent 根据 ID 获取卡片交易事件。
func (s *Service) GetCardTransactionEvent(ctx context.Context, id string, opts ...sdk.RequestOption) (*CardTransactionEvent, error) {
	var resp CardTransactionEvent
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/card_transaction_events/"+id, nil, &resp, opts...)
	return &resp, err
}
