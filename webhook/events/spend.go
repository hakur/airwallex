// Package events provides typed webhook event structures for the spend domain.
// Spend 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/spend.md
//
// 事件映射表:
//
//	spend.expense.draft                     → SpendExpenseDraftEvent           (Data: SpendExpenseEventData)
//	spend.expense.awaiting_approval         → SpendExpenseAwaitingApprovalEvent (Data: SpendExpenseEventData)
//	spend.expense.updated                   → SpendExpenseUpdatedEvent         (Data: SpendExpenseEventData)
//	spend.expense.rejected                  → SpendExpenseRejectedEvent        (Data: SpendExpenseEventData)
//	spend.expense.approved                  → SpendExpenseApprovedEvent        (Data: SpendExpenseEventData)
//	spend.expense.archived                  → SpendExpenseArchivedEvent        (Data: SpendExpenseEventData)
//	spend.expense.deleted                   → SpendExpenseDeletedEvent         (Data: SpendExpenseEventData)
//	spend.reimbursement_report.draft        → SpendReimbursementReportDraftEvent (Data: SpendReimbursementReportEventData)
//	spend.reimbursement_report.awaiting_approval → SpendReimbursementReportAwaitingApprovalEvent (Data: SpendReimbursementReportEventData)
//	spend.reimbursement_report.awaiting_payment → SpendReimbursementReportAwaitingPaymentEvent (Data: SpendReimbursementReportEventData)
//	spend.reimbursement_report.rejected     → SpendReimbursementReportRejectedEvent (Data: SpendReimbursementReportEventData)
//	spend.reimbursement_report.payment_in_progress → SpendReimbursementReportPaymentInProgressEvent (Data: SpendReimbursementReportEventData)
//	spend.reimbursement_report.paid         → SpendReimbursementReportPaidEvent (Data: SpendReimbursementReportEventData)
//	spend.reimbursement_report.mark_as_paid → SpendReimbursementReportMarkAsPaidEvent (Data: SpendReimbursementReportEventData)
//	spend.reimbursement_report.deleted      → SpendReimbursementReportDeletedEvent (Data: SpendReimbursementReportEventData)
//	spend.reimbursement_report.updated      → SpendReimbursementReportUpdatedEvent (Data: SpendReimbursementReportEventData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Card Expense Events ---

// SpendExpenseDraftEvent represents the spend.expense.draft webhook event.
type SpendExpenseDraftEvent struct {
	Event
	Data SpendExpenseEventData `json:"data"`
}

// SpendExpenseAwaitingApprovalEvent represents the spend.expense.awaiting_approval webhook event.
type SpendExpenseAwaitingApprovalEvent struct {
	Event
	Data SpendExpenseEventData `json:"data"`
}

// SpendExpenseUpdatedEvent represents the spend.expense.updated webhook event.
type SpendExpenseUpdatedEvent struct {
	Event
	Data SpendExpenseEventData `json:"data"`
}

// SpendExpenseRejectedEvent represents the spend.expense.rejected webhook event.
type SpendExpenseRejectedEvent struct {
	Event
	Data SpendExpenseEventData `json:"data"`
}

// SpendExpenseApprovedEvent represents the spend.expense.approved webhook event.
type SpendExpenseApprovedEvent struct {
	Event
	Data SpendExpenseEventData `json:"data"`
}

// SpendExpenseArchivedEvent represents the spend.expense.archived webhook event.
type SpendExpenseArchivedEvent struct {
	Event
	Data SpendExpenseEventData `json:"data"`
}

// SpendExpenseDeletedEvent represents the spend.expense.deleted webhook event.
type SpendExpenseDeletedEvent struct {
	Event
	Data SpendExpenseEventData `json:"data"`
}

// SpendExpenseEventData contains card expense information.
type SpendExpenseEventData struct {
	ID                        string                       `json:"id"`
	AccountID                 string                       `json:"account_id"`
	LegalEntityID             string                       `json:"legal_entity_id"`
	CardID                    string                       `json:"card_id"`
	BillingAmount             string                       `json:"billing_amount"`
	BillingCurrency           string                       `json:"billing_currency"`
	Merchant                  string                       `json:"merchant,omitempty"`
	Status                    string                       `json:"status"`
	SyncStatus                string                       `json:"sync_status"`
	Description               string                       `json:"description,omitempty"`
	CreatedAt                 string                       `json:"created_at"`
	UpdatedAt                 string                       `json:"updated_at"`
	SettledAt                 string                       `json:"settled_at,omitempty"`
	AccountingFieldSelections []any                        `json:"accounting_field_selections,omitempty"`
	Approvers                 []string                     `json:"approvers,omitempty"`
	Attachments               []string                     `json:"attachments,omitempty"`
	Comments                  []string                     `json:"comments,omitempty"`
	LineItems                 []SpendExpenseLineItem       `json:"line_items,omitempty"`
	CardTransaction           *SpendExpenseCardTransaction `json:"card_transaction,omitempty"`
}

// SpendExpenseLineItem represents a line item within a card expense.
type SpendExpenseLineItem struct {
	ID                        string `json:"id"`
	Description               string `json:"description,omitempty"`
	TransactionAmount         string `json:"transaction_amount"`
	AccountingFieldSelections []any  `json:"accounting_field_selections,omitempty"`
}

// SpendExpenseCardTransaction represents card transaction details in a spend expense event.
type SpendExpenseCardTransaction struct {
	Status   string `json:"status"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// --- Reimbursement Report Events ---

// SpendReimbursementReportDraftEvent represents the spend.reimbursement_report.draft webhook event.
type SpendReimbursementReportDraftEvent struct {
	Event
	Data SpendReimbursementReportEventData `json:"data"`
}

// SpendReimbursementReportAwaitingApprovalEvent represents the spend.reimbursement_report.awaiting_approval webhook event.
type SpendReimbursementReportAwaitingApprovalEvent struct {
	Event
	Data SpendReimbursementReportEventData `json:"data"`
}

// SpendReimbursementReportAwaitingPaymentEvent represents the spend.reimbursement_report.awaiting_payment webhook event.
type SpendReimbursementReportAwaitingPaymentEvent struct {
	Event
	Data SpendReimbursementReportEventData `json:"data"`
}

// SpendReimbursementReportRejectedEvent represents the spend.reimbursement_report.rejected webhook event.
type SpendReimbursementReportRejectedEvent struct {
	Event
	Data SpendReimbursementReportEventData `json:"data"`
}

// SpendReimbursementReportPaymentInProgressEvent represents the spend.reimbursement_report.payment_in_progress webhook event.
type SpendReimbursementReportPaymentInProgressEvent struct {
	Event
	Data SpendReimbursementReportEventData `json:"data"`
}

// SpendReimbursementReportPaidEvent represents the spend.reimbursement_report.paid webhook event.
type SpendReimbursementReportPaidEvent struct {
	Event
	Data SpendReimbursementReportEventData `json:"data"`
}

// SpendReimbursementReportMarkAsPaidEvent represents the spend.reimbursement_report.mark_as_paid webhook event.
type SpendReimbursementReportMarkAsPaidEvent struct {
	Event
	Data SpendReimbursementReportEventData `json:"data"`
}

// SpendReimbursementReportDeletedEvent represents the spend.reimbursement_report.deleted webhook event.
type SpendReimbursementReportDeletedEvent struct {
	Event
	Data SpendReimbursementReportEventData `json:"data"`
}

// SpendReimbursementReportUpdatedEvent represents the spend.reimbursement_report.updated webhook event.
type SpendReimbursementReportUpdatedEvent struct {
	Event
	Data SpendReimbursementReportEventData `json:"data"`
}

// SpendReimbursementReportEventData contains reimbursement report information.
type SpendReimbursementReportEventData struct {
	ID                        string   `json:"id"`
	Name                      string   `json:"name"`
	LegalEntityID             string   `json:"legal_entity_id"`
	BeneficiaryID             string   `json:"beneficiary_id,omitempty"`
	BillingCurrency           string   `json:"billing_currency"`
	Status                    string   `json:"status"`
	SyncStatus                string   `json:"sync_status"`
	CreatedBy                 string   `json:"created_by,omitempty"`
	CreatedAt                 string   `json:"created_at"`
	UpdatedAt                 string   `json:"updated_at"`
	AccountingFieldSelections []any    `json:"accounting_field_selections,omitempty"`
	Approvers                 []string `json:"approvers,omitempty"`
	Comments                  []string `json:"comments,omitempty"`
}
