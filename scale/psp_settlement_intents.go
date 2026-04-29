package scale

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PSPSettlementIntentStatus represents a PSP settlement intent status.
// PSPSettlementIntentStatus PSP结算意图状态。
type PSPSettlementIntentStatus = string

const (
	PSPSettlementIntentStatusNew            PSPSettlementIntentStatus = "NEW"
	PSPSettlementIntentStatusSubmitted      PSPSettlementIntentStatus = "SUBMITTED"
	PSPSettlementIntentStatusActionRequired PSPSettlementIntentStatus = "ACTION_REQUIRED"
	PSPSettlementIntentStatusMatched        PSPSettlementIntentStatus = "MATCHED"
	PSPSettlementIntentStatusSettled        PSPSettlementIntentStatus = "SETTLED"
)

// PSPSettlementIntent represents a PSP settlement intent.
// PSPSettlementIntent 表示PSP结算意图信息。
type PSPSettlementIntent struct {
	// ID PSP结算意图唯一标识符。
	ID string `json:"id,omitempty"`
	// RequestID 请求唯一标识符。
	RequestID string `json:"request_id,omitempty"`
	// Currency 结算货币（3位ISO-4217代码）。
	Currency sdk.Currency `json:"currency,omitempty"`
	// GlobalAccountID 接受存款的全球账户ID。
	GlobalAccountID string `json:"global_account_id,omitempty"`
	// SettlementReference PSP结算参考值。
	SettlementReference string `json:"settlement_reference,omitempty"`
	// ExpectedSettlementDate 预期结算日期（YYYY-MM-DD格式，UTC时区）。
	ExpectedSettlementDate string `json:"expected_settlement_date,omitempty"`
	// Status 结算意图状态。
	Status PSPSettlementIntentStatus `json:"status,omitempty"`
	// Metadata 元数据。
	Metadata map[string]any `json:"metadata,omitempty"`
	// AssociateByDepositIDs 应关联的存款ID列表。
	AssociateByDepositIDs []string `json:"associate_by_deposit_ids,omitempty"`
	// AssociatedDepositIDs 当前已关联的存款ID列表。
	AssociatedDepositIDs []string `json:"associated_deposit_ids,omitempty"`
	// Requirements 要求信息。
	Requirements map[string]any `json:"requirements,omitempty"`
	// CreatedAt 创建时间。
	CreatedAt string `json:"created_at,omitempty"`
}

// CreatePSPSettlementIntentRequest represents a request to create a PSP settlement intent.
// CreatePSPSettlementIntentRequest 创建PSP结算意图请求。
type CreatePSPSettlementIntentRequest struct {
	// RequestID 请求唯一标识符（最多50字符）。必填。
	RequestID string `json:"request_id"`
	// Currency 结算货币（3位ISO-4217代码）。必填。
	Currency sdk.Currency `json:"currency"`
	// GlobalAccountID 接受存款的全球账户ID。必填。
	GlobalAccountID string `json:"global_account_id"`
	// SettlementReference PSP结算参考值（最少5字符，最多255字符）。必填。
	SettlementReference string `json:"settlement_reference"`
	// ExpectedSettlementDate 预期结算日期（YYYY-MM-DD格式，UTC时区）。必填。
	ExpectedSettlementDate string `json:"expected_settlement_date"`
	// Metadata 元数据。
	Metadata map[string]any `json:"metadata,omitempty"`
	// AssociateByDepositIDs 应关联的存款ID列表。
	AssociateByDepositIDs []string `json:"associate_by_deposit_ids,omitempty"`
}

// UpdatePSPSettlementIntentRequest represents a request to update a PSP settlement intent.
// UpdatePSPSettlementIntentRequest 更新PSP结算意图请求。
type UpdatePSPSettlementIntentRequest struct {
	// Currency 结算货币（3位ISO-4217代码）。
	Currency sdk.Currency `json:"currency,omitempty"`
	// GlobalAccountID 接受存款的全球账户ID。
	GlobalAccountID string `json:"global_account_id,omitempty"`
	// SettlementReference PSP结算参考值。
	SettlementReference string `json:"settlement_reference,omitempty"`
	// ExpectedSettlementDate 预期结算日期（YYYY-MM-DD格式，UTC时区）。
	ExpectedSettlementDate string `json:"expected_settlement_date,omitempty"`
	// Metadata 元数据。
	Metadata map[string]any `json:"metadata,omitempty"`
	// AssociateByDepositIDs 应关联的存款ID列表。
	AssociateByDepositIDs []string `json:"associate_by_deposit_ids,omitempty"`
}

// CreatePSPSettlementIntent creates a new PSP settlement intent.
// CreatePSPSettlementIntent 创建PSP结算意图。
// 官方文档: https://www.airwallex.com/docs/api/scale/psp_settlement_intents/create.md
func (s *Service) CreatePSPSettlementIntent(ctx context.Context, req *CreatePSPSettlementIntentRequest, opts ...sdk.RequestOption) (*PSPSettlementIntent, error) {
	var resp PSPSettlementIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/psp_settlement_intents/create", req, &resp, opts...)
	return &resp, err
}

// GetPSPSettlementIntent retrieves a PSP settlement intent by ID.
// GetPSPSettlementIntent 根据ID获取PSP结算意图。
// 官方文档: https://www.airwallex.com/docs/api/scale/psp_settlement_intents/retrieve.md
func (s *Service) GetPSPSettlementIntent(ctx context.Context, id string, opts ...sdk.RequestOption) (*PSPSettlementIntent, error) {
	var resp PSPSettlementIntent
	err := s.doer.Do(ctx, "GET", "/api/v1/psp_settlement_intents/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListPSPSettlementIntents lists PSP settlement intents.
// ListPSPSettlementIntents 列出PSP结算意图。
// 官方文档: https://www.airwallex.com/docs/api/scale/psp_settlement_intents/list.md
func (s *Service) ListPSPSettlementIntents(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[PSPSettlementIntent], error) {
	var resp sdk.ListResult[PSPSettlementIntent]
	err := s.doer.Do(ctx, "GET", "/api/v1/psp_settlement_intents", nil, &resp, opts...)
	return &resp, err
}

// CancelPSPSettlementIntent cancels a PSP settlement intent.
// CancelPSPSettlementIntent 取消PSP结算意图。
// 官方文档: https://www.airwallex.com/docs/api/scale/psp_settlement_intents/cancel.md
func (s *Service) CancelPSPSettlementIntent(ctx context.Context, id string, opts ...sdk.RequestOption) (*PSPSettlementIntent, error) {
	var resp PSPSettlementIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/psp_settlement_intents/"+id+"/cancel", nil, &resp, opts...)
	return &resp, err
}

// SubmitPSPSettlementIntent submits a PSP settlement intent for processing.
// SubmitPSPSettlementIntent 提交PSP结算意图。
// 官方文档: https://www.airwallex.com/docs/api/scale/psp_settlement_intents/submit.md
func (s *Service) SubmitPSPSettlementIntent(ctx context.Context, id string, opts ...sdk.RequestOption) (*PSPSettlementIntent, error) {
	var resp PSPSettlementIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/psp_settlement_intents/"+id+"/submit", nil, &resp, opts...)
	return &resp, err
}

// UpdatePSPSettlementIntent updates an existing PSP settlement intent.
// UpdatePSPSettlementIntent 更新PSP结算意图。
// 官方文档: https://www.airwallex.com/docs/api/scale/psp_settlement_intents/update.md
func (s *Service) UpdatePSPSettlementIntent(ctx context.Context, id string, req *UpdatePSPSettlementIntentRequest, opts ...sdk.RequestOption) (*PSPSettlementIntent, error) {
	var resp PSPSettlementIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/psp_settlement_intents/"+id+"/update", req, &resp, opts...)
	return &resp, err
}
