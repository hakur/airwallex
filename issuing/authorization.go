package issuing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// Merchant represents merchant information.
// Merchant 表示商户信息。
type Merchant struct {
	// Name is the merchant name. Optional.
	// Name 商户名称。可选。
	Name string `json:"name,omitempty"`
	// City is the merchant city. Optional.
	// City 商户所在城市。可选。
	City string `json:"city,omitempty"`
	// Country is the merchant country. Optional.
	// Country 商户所在国家。可选。
	Country string `json:"country,omitempty"`
	// MCC is the merchant category code. Optional.
	// MCC 商户类别代码。可选。
	MCC string `json:"mcc,omitempty"`
	// ID is the unique merchant identifier. Optional.
	// ID 商户唯一标识符。可选。
	ID string `json:"id,omitempty"`
	// URL is the merchant URL. Optional.
	// URL 商户网址。可选。
	URL string `json:"url,omitempty"`
	// AcquirerID is the acquirer identifier. Optional.
	// AcquirerID 收单机构标识符。可选。
	AcquirerID string `json:"acquirer_id,omitempty"`
}

// Authorization represents authorization information.
// Authorization 表示授权信息。
type Authorization struct {
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
	// BillingAmount is the billing amount. Optional.
	// BillingAmount 账单金额。可选。
	BillingAmount float64 `json:"billing_amount,omitempty"`
	// BillingCurrency is the billing currency. Optional.
	// BillingCurrency 账单币种。可选。
	BillingCurrency sdk.Currency `json:"billing_currency,omitempty"`
	// CreateTime is the creation time. Optional.
	// CreateTime 创建时间。可选。
	CreateTime string `json:"create_time,omitempty"`
	// ExpiryDate is the expiry date. Optional.
	// ExpiryDate 到期日期。可选。
	ExpiryDate string `json:"expiry_date,omitempty"`
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
	// NetworkTransactionID is the network transaction ID. Optional.
	// NetworkTransactionID 网络交易ID。可选。
	NetworkTransactionID string `json:"network_transaction_id,omitempty"`
	// RetrievalRef is the retrieval reference number. Optional.
	// RetrievalRef 检索参考号。可选。
	RetrievalRef string `json:"retrieval_ref,omitempty"`
	// RiskDetails contains risk details. Optional.
	// RiskDetails 风险详情。可选。
	RiskDetails map[string]any `json:"risk_details,omitempty"`
	// UpdatedByTransaction is the transaction that triggered the update. Optional.
	// UpdatedByTransaction 更新来源交易。可选。
	UpdatedByTransaction string `json:"updated_by_transaction,omitempty"`
	// AcquiringInstitutionIdentifier is the acquiring institution identifier. Optional.
	// AcquiringInstitutionIdentifier 收单机构标识符。可选。
	AcquiringInstitutionIdentifier string `json:"acquiring_institution_identifier,omitempty"`
}

// GetAuthorization retrieves an authorization by ID.
// 官方文档: https://www.airwallex.com/docs/api/issuing/authorizations/retrieve.md
// GetAuthorization 根据 ID 获取授权。
func (s *Service) GetAuthorization(ctx context.Context, id string, opts ...sdk.RequestOption) (*Authorization, error) {
	var resp Authorization
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/authorizations/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListAuthorizations lists authorizations.
// 官方文档: https://www.airwallex.com/docs/api/issuing/authorizations/list.md
// ListAuthorizations 列出授权。
func (s *Service) ListAuthorizations(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[Authorization], error) {
	var resp sdk.ListResult[Authorization]
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/authorizations", nil, &resp, opts...)
	return &resp, err
}
