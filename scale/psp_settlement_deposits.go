package scale

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PSPSettlementDeposit represents a PSP settlement deposit.
// PSPSettlementDeposit 表示 PSP 结算存款。
type PSPSettlementDeposit struct {
	Amount                string                            `json:"amount,omitempty"`
	Currency              string                            `json:"currency,omitempty"`
	DepositID             string                            `json:"deposit_id,omitempty"`
	DepositedAt           string                            `json:"deposited_at,omitempty"`
	GlobalAccountID       string                            `json:"global_account_id,omitempty"`
	HoldingAccountID      string                            `json:"holding_account_id,omitempty"`
	Payer                 *PSPSettlementDepositPayer        `json:"payer,omitempty"`
	PSPSettlementIntentID string                            `json:"psp_settlement_intent_id,omitempty"`
	Reference             string                            `json:"reference,omitempty"`
	Requirements          *PSPSettlementDepositRequirements `json:"requirements,omitempty"`
	Status                string                            `json:"status,omitempty"`
}

// PSPSettlementDepositPayer represents the payer of a settlement deposit.
// PSPSettlementDepositPayer 存款付款人。
type PSPSettlementDepositPayer struct {
	Name string `json:"name,omitempty"`
}

// PSPSettlementDepositRequirements represents requirements for a settlement deposit.
// PSPSettlementDepositRequirements 存款要求。
type PSPSettlementDepositRequirements struct {
	MatchingPSPSettlementRequired   *RequirementMessage `json:"matching_psp_settlement_required,omitempty"`
	PSPSettlementIntentRequired     *RequirementMessage `json:"psp_settlement_intent_required,omitempty"`
	ReferenceDisambiguationRequired *RequirementMessage `json:"reference_disambiguation_required,omitempty"`
}

// RequirementMessage represents a requirement message.
// RequirementMessage 要求信息。
type RequirementMessage struct {
	Message string `json:"message,omitempty"`
}

// ListPSPSettlementDepositsRequest represents parameters for listing PSP settlement deposits.
// ListPSPSettlementDepositsRequest 结算存款列表查询参数。
type ListPSPSettlementDepositsRequest struct {
	DepositID        string `json:"deposit_id,omitempty"`
	FromDepositedAt  string `json:"from_deposited_at,omitempty"`
	GlobalAccountID  string `json:"global_account_id,omitempty"`
	HoldingAccountID string `json:"holding_account_id,omitempty"`
	PageNum          int32  `json:"page_num,omitempty"`
	PageSize         int32  `json:"page_size,omitempty"`
	Status           string `json:"status,omitempty"`
	ToDepositedAt    string `json:"to_deposited_at,omitempty"`
}

// ListPSPSettlementDepositsResponse represents a response listing PSP settlement deposits.
// ListPSPSettlementDepositsResponse 结算存款列表响应。
type ListPSPSettlementDepositsResponse struct {
	Items      []PSPSettlementDeposit `json:"items"`
	PageAfter  string                 `json:"page_after,omitempty"`
	PageBefore string                 `json:"page_before,omitempty"`
}

// ListPSPSettlementDeposits lists PSP settlement deposits.
// ListPSPSettlementDeposits 列出 PSP 结算存款。
// 官方文档: https://www.airwallex.com/docs/api/scale/psp_settlement_deposits/list.md
func (s *Service) ListPSPSettlementDeposits(ctx context.Context, req *ListPSPSettlementDepositsRequest, opts ...sdk.RequestOption) (*ListPSPSettlementDepositsResponse, error) {
	var resp ListPSPSettlementDepositsResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/psp_settlement_deposits", req, &resp, opts...)
	return &resp, err
}

// --- Deprecated ---
// CreatePSPSettlementDeposit / GetPSPSettlementDeposit 已移除。
// 官方 API 仅支持 List 端点。PSP Settlement Deposit 由系统自动创建，不可通过 API 手动创建。
