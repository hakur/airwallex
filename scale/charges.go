package scale

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// ChargeStatus represents a charge status.
// ChargeStatus 收费状态。
type ChargeStatus = string

const (
	ChargeStatusNew       ChargeStatus = "NEW"
	ChargeStatusPending   ChargeStatus = "PENDING"
	ChargeStatusSettled   ChargeStatus = "SETTLED"
	ChargeStatusSuspended ChargeStatus = "SUSPENDED"
	ChargeStatusFailed    ChargeStatus = "FAILED"
)

// Charge represents a charge.
// Charge 表示收费信息。
type Charge struct {
	// ID 收费唯一标识符。
	ID string `json:"id,omitempty"`
	// Source 来源Airwallex账户ID。
	Source string `json:"source,omitempty"`
	// Amount 收费金额。
	Amount float64 `json:"amount,omitempty"`
	// Currency 收费货币（3位ISO-4217代码）。
	Currency sdk.Currency `json:"currency,omitempty"`
	// Status 收费状态。
	Status ChargeStatus `json:"status,omitempty"`
	// Description 收费描述。
	Description string `json:"description,omitempty"`
	// Reason 收费原因。
	Reason string `json:"reason,omitempty"`
	// Reference 用户指定的参考信息。
	Reference string `json:"reference,omitempty"`
	// RequestID 请求唯一标识符。
	RequestID string `json:"request_id,omitempty"`
	// ShortReferenceID 短参考ID（用于客服支持）。
	ShortReferenceID string `json:"short_reference_id,omitempty"`
	// Fee 收费费用。
	Fee float64 `json:"fee,omitempty"`
	// CreatedAt 创建时间。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt 更新时间。
	UpdatedAt string `json:"updated_at,omitempty"`
}

// CreateChargeRequest represents a request to create a charge.
// CreateChargeRequest 创建收费请求。
type CreateChargeRequest struct {
	// RequestID 请求唯一标识符（1-50字符）。必填。
	RequestID string `json:"request_id"`
	// Source 来源Airwallex账户ID。必填。
	Source string `json:"source"`
	// Amount 收费金额（字符串类型）。必填。
	Amount string `json:"amount"`
	// Currency 收费货币（3位ISO-4217代码）。必填。
	Currency sdk.Currency `json:"currency"`
	// Reason 收费原因。必填。
	Reason string `json:"reason"`
	// Reference 用户指定的参考信息（1-140字符）。必填。
	Reference string `json:"reference"`
	// Description 收费描述。
	Description string `json:"description,omitempty"`
}

// CreateCharge creates a new charge.
// CreateCharge 创建收费。
// 官方文档: https://www.airwallex.com/docs/api/scale/charges/create.md
func (s *Service) CreateCharge(ctx context.Context, req *CreateChargeRequest, opts ...sdk.RequestOption) (*Charge, error) {
	var resp Charge
	err := s.doer.Do(ctx, "POST", "/api/v1/charges/create", req, &resp, opts...)
	return &resp, err
}

// GetCharge retrieves a charge by ID.
// GetCharge 根据ID获取收费。
// 官方文档: https://www.airwallex.com/docs/api/scale/charges/retrieve.md
func (s *Service) GetCharge(ctx context.Context, id string, opts ...sdk.RequestOption) (*Charge, error) {
	var resp Charge
	err := s.doer.Do(ctx, "GET", "/api/v1/charges/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListCharges lists charges.
// ListCharges 列出收费。
// 官方文档: https://www.airwallex.com/docs/api/scale/charges/list.md
func (s *Service) ListCharges(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[Charge], error) {
	var resp sdk.ListResult[Charge]
	err := s.doer.Do(ctx, "GET", "/api/v1/charges", nil, &resp, opts...)
	return &resp, err
}
