// Package events provides typed webhook event structures for the PSP agnostic domain.
// PSP Agnostic 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/psp-agnostic.md
//
// 事件映射表:
//
//	psp_settlement_intent.new           → PSPSettlementIntentNewEvent      (Data: PSPSettlementIntent)
//	psp_settlement_intent.cancelled     → PSPSettlementIntentCancelledEvent (Data: PSPSettlementIntent)
//	psp_settlement_intent.submitted     → PSPSettlementIntentSubmittedEvent (Data: PSPSettlementIntent)
//	psp_settlement_intent.action_required → PSPSettlementIntentActionRequiredEvent (Data: PSPSettlementIntent)
//	psp_settlement_intent.matched       → PSPSettlementIntentMatchedEvent  (Data: PSPSettlementIntent)
//	psp_settlement_intent.settled       → PSPSettlementIntentSettledEvent  (Data: PSPSettlementIntent)
//	psp_settlement_deposit.new          → PSPSettlementDepositNewEvent     (Data: PSPSettlementDeposit)
//	psp_settlement_deposit.action_required → PSPSettlementDepositActionRequiredEvent (Data: PSPSettlementDeposit)
//	psp_settlement_deposit.matched      → PSPSettlementDepositMatchedEvent (Data: PSPSettlementDeposit)
//	psp_settlement_deposit.settled      → PSPSettlementDepositSettledEvent (Data: PSPSettlementDeposit)
//	psp_settlement_split.new            → PSPSettlementSplitNewEvent       (Data: PSPSettlementSplit)
//	psp_settlement_split.create_failed  → PSPSettlementSplitCreateFailedEvent (Data: PSPSettlementSplit)
//	psp_settlement_split.cancelled      → PSPSettlementSplitCancelledEvent (Data: PSPSettlementSplit)
//	psp_settlement_split.matched        → PSPSettlementSplitMatchedEvent   (Data: PSPSettlementSplit)
//	psp_settlement_split.pending        → PSPSettlementSplitPendingEvent   (Data: PSPSettlementSplit)
//	psp_settlement_split.failed         → PSPSettlementSplitFailedEvent    (Data: PSPSettlementSplit)
//	psp_settlement_split.settled        → PSPSettlementSplitSettledEvent   (Data: PSPSettlementSplit)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- PSP Settlement Intent Events ---

// PSPSettlementIntentNewEvent represents the psp_settlement_intent.new webhook event.
type PSPSettlementIntentNewEvent struct {
	Event
	Data PSPSettlementIntent `json:"data"`
}

// PSPSettlementIntentCancelledEvent represents the psp_settlement_intent.cancelled webhook event.
type PSPSettlementIntentCancelledEvent struct {
	Event
	Data PSPSettlementIntent `json:"data"`
}

// PSPSettlementIntentSubmittedEvent represents the psp_settlement_intent.submitted webhook event.
type PSPSettlementIntentSubmittedEvent struct {
	Event
	Data PSPSettlementIntent `json:"data"`
}

// PSPSettlementIntentActionRequiredEvent represents the psp_settlement_intent.action_required webhook event.
type PSPSettlementIntentActionRequiredEvent struct {
	Event
	Data PSPSettlementIntent `json:"data"`
}

// PSPSettlementIntentMatchedEvent represents the psp_settlement_intent.matched webhook event.
type PSPSettlementIntentMatchedEvent struct {
	Event
	Data PSPSettlementIntent `json:"data"`
}

// PSPSettlementIntentSettledEvent represents the psp_settlement_intent.settled webhook event.
type PSPSettlementIntentSettledEvent struct {
	Event
	Data PSPSettlementIntent `json:"data"`
}

// --- PSP Settlement Deposit Events ---

// PSPSettlementDepositNewEvent represents the psp_settlement_deposit.new webhook event.
type PSPSettlementDepositNewEvent struct {
	Event
	Data PSPSettlementDeposit `json:"data"`
}

// PSPSettlementDepositActionRequiredEvent represents the psp_settlement_deposit.action_required webhook event.
type PSPSettlementDepositActionRequiredEvent struct {
	Event
	Data PSPSettlementDeposit `json:"data"`
}

// PSPSettlementDepositMatchedEvent represents the psp_settlement_deposit.matched webhook event.
type PSPSettlementDepositMatchedEvent struct {
	Event
	Data PSPSettlementDeposit `json:"data"`
}

// PSPSettlementDepositSettledEvent represents the psp_settlement_deposit.settled webhook event.
type PSPSettlementDepositSettledEvent struct {
	Event
	Data PSPSettlementDeposit `json:"data"`
}

// --- PSP Settlement Split Events ---

// PSPSettlementSplitNewEvent represents the psp_settlement_split.new webhook event.
type PSPSettlementSplitNewEvent struct {
	Event
	Data PSPSettlementSplit `json:"data"`
}

// PSPSettlementSplitCreateFailedEvent represents the psp_settlement_split.create_failed webhook event.
type PSPSettlementSplitCreateFailedEvent struct {
	Event
	Data PSPSettlementSplit `json:"data"`
}

// PSPSettlementSplitCancelledEvent represents the psp_settlement_split.cancelled webhook event.
type PSPSettlementSplitCancelledEvent struct {
	Event
	Data PSPSettlementSplit `json:"data"`
}

// PSPSettlementSplitMatchedEvent represents the psp_settlement_split.matched webhook event.
type PSPSettlementSplitMatchedEvent struct {
	Event
	Data PSPSettlementSplit `json:"data"`
}

// PSPSettlementSplitPendingEvent represents the psp_settlement_split.pending webhook event.
type PSPSettlementSplitPendingEvent struct {
	Event
	Data PSPSettlementSplit `json:"data"`
}

// PSPSettlementSplitFailedEvent represents the psp_settlement_split.failed webhook event.
type PSPSettlementSplitFailedEvent struct {
	Event
	Data PSPSettlementSplit `json:"data"`
}

// PSPSettlementSplitSettledEvent represents the psp_settlement_split.settled webhook event.
type PSPSettlementSplitSettledEvent struct {
	Event
	Data PSPSettlementSplit `json:"data"`
}

// --- Event Data Structures ---

// PSPSettlementIntent represents a PSP settlement intent in webhook payload.
type PSPSettlementIntent struct {
	ID                     string         `json:"id,omitempty"`
	RequestID              string         `json:"request_id,omitempty"`
	Currency               string         `json:"currency,omitempty"`
	GlobalAccountID        string         `json:"global_account_id,omitempty"`
	SettlementReference    string         `json:"settlement_reference,omitempty"`
	ExpectedSettlementDate string         `json:"expected_settlement_date,omitempty"`
	Status                 string         `json:"status,omitempty"`
	Metadata               map[string]any `json:"metadata,omitempty"`
	AssociateByDepositIDs  []string       `json:"associate_by_deposit_ids,omitempty"`
	AssociatedDepositIDs   []string       `json:"associated_deposit_ids,omitempty"`
	Requirements           map[string]any `json:"requirements,omitempty"`
	CreatedAt              string         `json:"created_at,omitempty"`
}

// PSPSettlementDeposit represents a PSP settlement deposit in webhook payload.
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

// PSPSettlementDepositPayer represents the payer of a PSP settlement deposit in webhook payload.
type PSPSettlementDepositPayer struct {
	Name string `json:"name,omitempty"`
}

// PSPSettlementDepositRequirements represents requirements for a PSP settlement deposit in webhook payload.
type PSPSettlementDepositRequirements struct {
	MatchingPSPSettlementRequired   *RequirementMessage `json:"matching_psp_settlement_required,omitempty"`
	PSPSettlementIntentRequired     *RequirementMessage `json:"psp_settlement_intent_required,omitempty"`
	ReferenceDisambiguationRequired *RequirementMessage `json:"reference_disambiguation_required,omitempty"`
}

// RequirementMessage represents a requirement message in webhook payload.
type RequirementMessage struct {
	Message string `json:"message,omitempty"`
}

// PSPSettlementSplit represents a PSP settlement split in webhook payload.
type PSPSettlementSplit struct {
	ID              string `json:"id,omitempty"`
	Amount          string `json:"amount,omitempty"`
	SettlementType  string `json:"settlement_type,omitempty"`
	Status          string `json:"status,omitempty"`
	TargetAccountID string `json:"target_account_id,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
}
