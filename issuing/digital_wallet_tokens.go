package issuing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// DigitalWalletTokenStatus represents the digital wallet token status.
// DigitalWalletTokenStatus 数字钱包令牌状态。
type DigitalWalletTokenStatus = string

const (
	// DigitalWalletTokenStatusActive indicates the token is active.
	// DigitalWalletTokenStatusActive 活跃状态
	DigitalWalletTokenStatusActive DigitalWalletTokenStatus = "ACTIVE"
	// DigitalWalletTokenStatusInactive indicates the token is inactive.
	// DigitalWalletTokenStatusInactive 非活跃状态
	DigitalWalletTokenStatusInactive DigitalWalletTokenStatus = "INACTIVE"
	// DigitalWalletTokenStatusSuspended indicates the token is suspended.
	// DigitalWalletTokenStatusSuspended 已暂停
	DigitalWalletTokenStatusSuspended DigitalWalletTokenStatus = "SUSPENDED"
)

// DigitalWalletTokenType represents the digital wallet token type.
// DigitalWalletTokenType 数字钱包令牌类型。
type DigitalWalletTokenType = string

const (
	// DigitalWalletTokenTypeApplePay is an Apple Pay token.
	// DigitalWalletTokenTypeApplePay Apple Pay
	DigitalWalletTokenTypeApplePay DigitalWalletTokenType = "APPLE_PAY"
	// DigitalWalletTokenTypeGooglePay is a Google Pay token.
	// DigitalWalletTokenTypeGooglePay Google Pay
	DigitalWalletTokenTypeGooglePay DigitalWalletTokenType = "GOOGLE_PAY"
)

// DigitalWalletToken represents digital wallet token information.
// DigitalWalletToken 表示数字钱包令牌信息。
type DigitalWalletToken struct {
	// TokenID is the unique token identifier. Required.
	// TokenID 令牌唯一标识符。必填。
	TokenID string `json:"token_id"`
	// CardID is the unique card identifier. Required.
	// CardID 卡片唯一标识符。必填。
	CardID string `json:"card_id"`
	// TokenStatus is the token status. Required.
	// TokenStatus 令牌状态。必填。
	TokenStatus DigitalWalletTokenStatus `json:"token_status"`
	// TokenType is the token type. Required.
	// TokenType 令牌类型。必填。
	TokenType DigitalWalletTokenType `json:"token_type"`
	// CardholderID is the unique cardholder identifier. Optional.
	// CardholderID 持卡人唯一标识符。可选。
	CardholderID string `json:"cardholder_id,omitempty"`
	// CreateTime is the creation time. Optional.
	// CreateTime 创建时间。可选。
	CreateTime string `json:"create_time,omitempty"`
	// DeviceInformation contains device information. Optional.
	// DeviceInformation 设备信息。可选。
	DeviceInformation map[string]any `json:"device_information,omitempty"`
	// ExpiryMonth is the expiry month. Optional.
	// ExpiryMonth 到期月份。可选。
	ExpiryMonth string `json:"expiry_month,omitempty"`
	// ExpiryYear is the expiry year. Optional.
	// ExpiryYear 到期年份。可选。
	ExpiryYear string `json:"expiry_year,omitempty"`
	// MaskedCardNumber is the masked card number. Optional.
	// MaskedCardNumber 掩码卡号。可选。
	MaskedCardNumber string `json:"masked_card_number,omitempty"`
	// PanReferenceID is the primary account number reference ID. Optional.
	// PanReferenceID 主账号引用ID。可选。
	PanReferenceID string `json:"pan_reference_id,omitempty"`
	// RiskInformation contains risk information. Optional.
	// RiskInformation 风险信息。可选。
	RiskInformation map[string]any `json:"risk_information,omitempty"`
	// TokenReferenceID is the token reference ID. Optional.
	// TokenReferenceID 令牌引用ID。可选。
	TokenReferenceID string `json:"token_reference_id,omitempty"`
}

// ListDigitalWalletTokens lists digital wallet tokens.
// 官方文档: https://www.airwallex.com/docs/api/issuing/digital_wallet_tokens/list.md
// ListDigitalWalletTokens 列出数字钱包令牌。
func (s *Service) ListDigitalWalletTokens(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[DigitalWalletToken], error) {
	var resp sdk.ListResult[DigitalWalletToken]
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/digital_wallet_tokens", nil, &resp, opts...)
	return &resp, err
}

// GetDigitalWalletToken retrieves a digital wallet token by ID.
// 官方文档: https://www.airwallex.com/docs/api/issuing/digital_wallet_tokens/retrieve.md
// GetDigitalWalletToken 根据 ID 获取数字钱包令牌。
func (s *Service) GetDigitalWalletToken(ctx context.Context, id string, opts ...sdk.RequestOption) (*DigitalWalletToken, error) {
	var resp DigitalWalletToken
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/digital_wallet_tokens/"+id, nil, &resp, opts...)
	return &resp, err
}
