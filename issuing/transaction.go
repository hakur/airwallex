package issuing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// Transaction represents transaction information.
// Transaction 表示交易信息。
type Transaction struct {
	// TransactionID is the unique transaction identifier. Required.
	// TransactionID 交易唯一标识符。必填。
	TransactionID string `json:"transaction_id"`
	// CardID is the unique card identifier. Required.
	// CardID 卡片唯一标识符。必填。
	CardID string `json:"card_id"`
	// Status is the status. Required.
	// Status 状态。必填。
	Status string `json:"status"`
	// TransactionAmount is the transaction amount. Required.
	// TransactionAmount 交易金额。必填。
	TransactionAmount float64 `json:"transaction_amount"`
	// TransactionCurrency is the transaction currency. Required.
	// TransactionCurrency 交易币种。必填。
	TransactionCurrency sdk.Currency `json:"transaction_currency"`
	// TransactionType is the transaction type. Optional.
	// TransactionType 交易类型。可选。
	TransactionType string `json:"transaction_type,omitempty"`
	// TransactionDate is the transaction date. Optional.
	// TransactionDate 交易日期。可选。
	TransactionDate string `json:"transaction_date,omitempty"`
	// BillingAmount is the billing amount. Optional.
	// BillingAmount 账单金额。可选。
	BillingAmount float64 `json:"billing_amount,omitempty"`
	// BillingCurrency is the billing currency. Optional.
	// BillingCurrency 账单币种。可选。
	BillingCurrency sdk.Currency `json:"billing_currency,omitempty"`
	// Merchant contains merchant information. Optional.
	// Merchant 商户信息。可选。
	Merchant *Merchant `json:"merchant,omitempty"`
	// AuthCode is the authorization code. Optional.
	// AuthCode 授权码。可选。
	AuthCode string `json:"auth_code,omitempty"`
	// CardNickname is the card nickname. Optional.
	// CardNickname 卡片昵称。可选。
	CardNickname string `json:"card_nickname,omitempty"`
	// ClientData contains client data. Optional.
	// ClientData 客户端数据。可选。
	ClientData map[string]any `json:"client_data,omitempty"`
	// DigitalWalletTokenID is the digital wallet token ID. Optional.
	// DigitalWalletTokenID 数字钱包令牌ID。可选。
	DigitalWalletTokenID string `json:"digital_wallet_token_id,omitempty"`
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
	// MatchedAuthorizations is the list of matched authorizations. Optional.
	// MatchedAuthorizations 匹配的授权列表。可选。
	MatchedAuthorizations []string `json:"matched_authorizations,omitempty"`
	// NetworkTransactionID is the network transaction ID. Optional.
	// NetworkTransactionID 网络交易ID。可选。
	NetworkTransactionID string `json:"network_transaction_id,omitempty"`
	// PostedDate is the posting date. Optional.
	// PostedDate 入账日期。可选。
	PostedDate string `json:"posted_date,omitempty"`
	// RetrievalRef is the retrieval reference number. Optional.
	// RetrievalRef 检索参考号。可选。
	RetrievalRef string `json:"retrieval_ref,omitempty"`
	// RiskDetails contains risk details. Optional.
	// RiskDetails 风险详情。可选。
	RiskDetails map[string]any `json:"risk_details,omitempty"`
	// AcquiringInstitutionIdentifier is the acquiring institution identifier. Optional.
	// AcquiringInstitutionIdentifier 收单机构标识符。可选。
	AcquiringInstitutionIdentifier string `json:"acquiring_institution_identifier,omitempty"`
}

// ListTransactions lists transactions.
// 官方文档: https://www.airwallex.com/docs/api/issuing/transactions/list.md
// ListTransactions 列出交易。
func (s *Service) ListTransactions(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[Transaction], error) {
	var resp sdk.ListResult[Transaction]
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/transactions", nil, &resp, opts...)
	return &resp, err
}

// GetTransaction retrieves a transaction by ID.
// 官方文档: https://www.airwallex.com/docs/api/issuing/transactions/retrieve.md
// GetTransaction 根据 ID 获取交易。
func (s *Service) GetTransaction(ctx context.Context, id string, opts ...sdk.RequestOption) (*Transaction, error) {
	var resp Transaction
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/transactions/"+id, nil, &resp, opts...)
	return &resp, err
}
