package simulation

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// IssuingTransactionType represents the issuing transaction type.
// IssuingTransactionType 发卡交易类型。
type IssuingTransactionType = string

const (
	// IssuingTransactionTypeAuthorization is an authorization type.
	IssuingTransactionTypeAuthorization IssuingTransactionType = "AUTHORIZATION"
	// IssuingTransactionTypeClearing is a clearing type.
	IssuingTransactionTypeClearing IssuingTransactionType = "CLEARING"
	// IssuingTransactionTypeRefund is a refund type.
	IssuingTransactionTypeRefund IssuingTransactionType = "REFUND"
	// IssuingTransactionTypeReversal is a reversal type.
	IssuingTransactionTypeReversal IssuingTransactionType = "REVERSAL"
	// IssuingTransactionTypeOriginalCredit is an original credit type.
	IssuingTransactionTypeOriginalCredit IssuingTransactionType = "ORIGINAL_CREDIT"
)

// IssuingTransactionStatus represents the issuing transaction status.
// IssuingTransactionStatus 发卡交易状态。
type IssuingTransactionStatus = string

const (
	// IssuingTransactionStatusApproved is the approved status.
	IssuingTransactionStatusApproved IssuingTransactionStatus = "APPROVED"
	// IssuingTransactionStatusPending is the pending status.
	IssuingTransactionStatusPending IssuingTransactionStatus = "PENDING"
	// IssuingTransactionStatusFailed is the failed status.
	IssuingTransactionStatusFailed IssuingTransactionStatus = "FAILED"
)

// IssuingTransaction represents an issuing transaction response.
// IssuingTransaction 表示发卡交易响应。
type IssuingTransaction struct {
	// TransactionID is the unique transaction identifier. Optional.
	// TransactionID 交易唯一标识符。可选。
	TransactionID string `json:"transaction_id,omitempty"`
	// TransactionType is the transaction type. Optional.
	// TransactionType 交易类型。可选。
	TransactionType IssuingTransactionType `json:"transaction_type,omitempty"`
	// Status is the transaction status. Optional.
	// Status 状态。可选。
	Status IssuingTransactionStatus `json:"status,omitempty"`
	// AuthCode is the authorization code. Optional.
	// AuthCode 授权码。可选。
	AuthCode string `json:"auth_code,omitempty"`
	// TransactionAmount is the transaction amount. Optional.
	// TransactionAmount 交易金额。可选。
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
	// TransactionCurrency is the transaction currency. Optional.
	// TransactionCurrency 交易币种。可选。
	TransactionCurrency string `json:"transaction_currency,omitempty"`
	// BillingAmount is the billing amount. Optional.
	// BillingAmount 账单金额。可选。
	BillingAmount float64 `json:"billing_amount,omitempty"`
	// BillingCurrency is the billing currency. Optional.
	// BillingCurrency 账单币种。可选。
	BillingCurrency string `json:"billing_currency,omitempty"`
	// CardID is the unique card identifier. Optional.
	// CardID 卡片唯一标识符。可选。
	CardID string `json:"card_id,omitempty"`
	// CardNickname is the card nickname. Optional.
	// CardNickname 卡片昵称。可选。
	CardNickname string `json:"card_nickname,omitempty"`
	// MaskedCardNumber is the masked card number. Optional.
	// MaskedCardNumber 掩码卡号。可选。
	MaskedCardNumber string `json:"masked_card_number,omitempty"`
	// ClientData is the client data. Optional.
	// ClientData 客户端数据。可选。
	ClientData string `json:"client_data,omitempty"`
	// FailureReason is the failure reason. Optional.
	// FailureReason 失败原因。可选。
	FailureReason string `json:"failure_reason,omitempty"`
	// NetworkTransactionID is the network transaction ID. Optional.
	// NetworkTransactionID 网络交易ID。可选。
	NetworkTransactionID string `json:"network_transaction_id,omitempty"`
	// RetrievalRef is the retrieval reference number. Optional.
	// RetrievalRef 检索参考号。可选。
	RetrievalRef string `json:"retrieval_ref,omitempty"`
	// PostedDate is the posting date. Optional.
	// PostedDate 入账日期。可选。
	PostedDate string `json:"posted_date,omitempty"`
	// TransactionDate is the transaction date. Optional.
	// TransactionDate 交易日期。可选。
	TransactionDate string `json:"transaction_date,omitempty"`
	// Merchant is the merchant information. Optional.
	// Merchant 商户信息。可选。
	Merchant *Merchant `json:"merchant,omitempty"`
}

// Merchant represents merchant information.
// Merchant 表示商户信息。
type Merchant struct {
	// CategoryCode is the merchant category code. Optional.
	// CategoryCode 商户类别代码。可选。
	CategoryCode string `json:"category_code,omitempty"`
	// City is the city. Optional.
	// City 城市。可选。
	City string `json:"city,omitempty"`
	// Country is the country. Optional.
	// Country 国家。可选。
	Country string `json:"country,omitempty"`
	// Name is the merchant name. Optional.
	// Name 商户名称。可选。
	Name string `json:"name,omitempty"`
}

// SimulateIssuingTransactionCreateRequest represents a request to create an issuing transaction.
// SimulateIssuingTransactionCreateRequest 创建发卡交易请求。
type SimulateIssuingTransactionCreateRequest struct {
	CardID                   string  `json:"card_id,omitempty"`
	CardNumber               string  `json:"card_number,omitempty"`
	TransactionAmount        float64 `json:"transaction_amount"`
	TransactionCurrency      string  `json:"transaction_currency"`
	AuthCode                 string  `json:"auth_code,omitempty"`
	MerchantCategoryCode     string  `json:"merchant_category_code,omitempty"`
	MerchantInfo             string  `json:"merchant_info,omitempty"`
	SinglePhase              bool    `json:"single_phase,omitempty"`
	TransactionFailureReason string  `json:"transaction_failure_reason,omitempty"`
	TransactionID            string  `json:"transaction_id,omitempty"`
}

// SimulateIssuingTransactionRefundRequest represents a request to refund an issuing transaction.
// SimulateIssuingTransactionRefundRequest 退款发卡交易请求。
type SimulateIssuingTransactionRefundRequest struct {
	CardID               string  `json:"card_id,omitempty"`
	CardNumber           string  `json:"card_number,omitempty"`
	TransactionAmount    float64 `json:"transaction_amount"`
	TransactionCurrency  string  `json:"transaction_currency"`
	MerchantCategoryCode string  `json:"merchant_category_code,omitempty"`
	MerchantInfo         string  `json:"merchant_info,omitempty"`
	TransactionID        string  `json:"transaction_id,omitempty"`
}

// SimulateIssuingTransactionCaptureRequest represents a request to capture an issuing transaction.
// SimulateIssuingTransactionCaptureRequest 捕获发卡交易请求。
type SimulateIssuingTransactionCaptureRequest struct {
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
	MerchantInfo      string  `json:"merchant_info,omitempty"`
}

// SimulateIssuingTransactionReverseRequest represents a request to reverse an issuing transaction.
// SimulateIssuingTransactionReverseRequest 撤销发卡交易请求。
type SimulateIssuingTransactionReverseRequest struct {
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
}

// SimulateIssuingThreedsNotifyRequest represents a 3DS notification request.
// SimulateIssuingThreedsNotifyRequest 3DS 通知请求。
type SimulateIssuingThreedsNotifyRequest struct {
	CardNumber   string                     `json:"card_number"`
	MerchantInfo *ThreedsNotifyMerchantInfo `json:"merchant_info,omitempty"`
}

// ThreedsNotifyMerchantInfo represents 3DS notification merchant information.
// ThreedsNotifyMerchantInfo 3DS 通知商户信息。
type ThreedsNotifyMerchantInfo struct {
	AcquirerID           string `json:"acquirer_id,omitempty"`
	MerchantCategoryCode string `json:"merchant_category_code,omitempty"`
	MerchantCountryCode  string `json:"merchant_country_code,omitempty"`
	MerchantID           string `json:"merchant_id,omitempty"`
	MerchantName         string `json:"merchant_name,omitempty"`
	MerchantURL          string `json:"merchant_url,omitempty"`
}

// SimulateIssuingTransactionCreate creates an issuing transaction.
// SimulateIssuingTransactionCreate 创建发卡交易。
// 官方文档: https://www.airwallex.com/docs/api/simulation/issuing/create.md
func (s *Service) SimulateIssuingTransactionCreate(ctx context.Context, req *SimulateIssuingTransactionCreateRequest, opts ...sdk.RequestOption) (*IssuingTransaction, error) {
	var resp IssuingTransaction
	err := s.doer.Do(ctx, "POST", "/api/v1/simulation/issuing/create", req, &resp, opts...)
	return &resp, err
}

// SimulateIssuingTransactionRefund refunds an issuing transaction.
// SimulateIssuingTransactionRefund 退款发卡交易。
// 官方文档: https://www.airwallex.com/docs/api/simulation/issuing/refund.md
func (s *Service) SimulateIssuingTransactionRefund(ctx context.Context, req *SimulateIssuingTransactionRefundRequest, opts ...sdk.RequestOption) (*IssuingTransaction, error) {
	var resp IssuingTransaction
	err := s.doer.Do(ctx, "POST", "/api/v1/simulation/issuing/refund", req, &resp, opts...)
	return &resp, err
}

// SimulateIssuingTransactionCapture captures an issuing transaction.
// SimulateIssuingTransactionCapture 捕获发卡交易。
// 官方文档: https://www.airwallex.com/docs/api/simulation/issuing/capture.md
func (s *Service) SimulateIssuingTransactionCapture(ctx context.Context, transactionID string, req *SimulateIssuingTransactionCaptureRequest, opts ...sdk.RequestOption) (*IssuingTransaction, error) {
	var resp IssuingTransaction
	err := s.doer.Do(ctx, "POST", "/api/v1/simulation/issuing/"+transactionID+"/capture", req, &resp, opts...)
	return &resp, err
}

// SimulateIssuingTransactionReverse reverses an issuing transaction.
// SimulateIssuingTransactionReverse 撤销发卡交易。
// 官方文档: https://www.airwallex.com/docs/api/simulation/issuing/reverse.md
func (s *Service) SimulateIssuingTransactionReverse(ctx context.Context, transactionID string, req *SimulateIssuingTransactionReverseRequest, opts ...sdk.RequestOption) (*IssuingTransaction, error) {
	var resp IssuingTransaction
	err := s.doer.Do(ctx, "POST", "/api/v1/simulation/issuing/"+transactionID+"/reverse", req, &resp, opts...)
	return &resp, err
}

// SimulateIssuingThreedsNotify sends a 3DS notification.
// SimulateIssuingThreedsNotify 发送 3DS 通知。
// 官方文档: https://www.airwallex.com/docs/api/simulation/issuing/notify_threeds.md
func (s *Service) SimulateIssuingThreedsNotify(ctx context.Context, req *SimulateIssuingThreedsNotifyRequest, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/issuing/threeds/notify", req, nil, opts...)
}

// SimulateIssuingCardholderPassReview passes a cardholder review.
// SimulateIssuingCardholderPassReview 持卡人审核通过。
// 官方文档: https://www.airwallex.com/docs/api/simulation/issuing/pass_review_cardholders.md
func (s *Service) SimulateIssuingCardholderPassReview(ctx context.Context, cardholderID string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/issuing/cardholders/"+cardholderID+"/pass_review", nil, nil, opts...)
}
