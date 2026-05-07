// Package events provides typed webhook event structures for the direct debit payouts domain.
// Direct Debit Payouts 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/direct-debit-payouts.md
//
// 事件映射表:
//
//	direct_debit.created                → DirectDebitCreatedEvent          (Data: DirectDebit)
//	direct_debit.in_review              → DirectDebitInReviewEvent         (Data: DirectDebit)
//	direct_debit.pending                → DirectDebitPendingEvent          (Data: DirectDebit)
//	direct_debit.rejected               → DirectDebitRejectedEvent         (Data: DirectDebit)
//	direct_debit.settled                → DirectDebitSettledEvent          (Data: DirectDebit)
//	direct_debit.returned               → DirectDebitReturnedEvent         (Data: DirectDebit)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- DirectDebit Events ---

// DirectDebitCreatedEvent represents the direct_debit.created webhook event.
// DirectDebitCreatedEvent 表示 direct_debit.created webhook 事件。
type DirectDebitCreatedEvent struct {
	Event
	Data DirectDebit `json:"data"`
}

// DirectDebitInReviewEvent represents the direct_debit.in_review webhook event.
// DirectDebitInReviewEvent 表示 direct_debit.in_review webhook 事件。
type DirectDebitInReviewEvent struct {
	Event
	Data DirectDebit `json:"data"`
}

// DirectDebitPendingEvent represents the direct_debit.pending webhook event.
// DirectDebitPendingEvent 表示 direct_debit.pending webhook 事件。
type DirectDebitPendingEvent struct {
	Event
	Data DirectDebit `json:"data"`
}

// DirectDebitRejectedEvent represents the direct_debit.rejected webhook event.
// DirectDebitRejectedEvent 表示 direct_debit.rejected webhook 事件。
type DirectDebitRejectedEvent struct {
	Event
	Data DirectDebit `json:"data"`
}

// DirectDebitSettledEvent represents the direct_debit.settled webhook event.
// DirectDebitSettledEvent 表示 direct_debit.settled webhook 事件。
type DirectDebitSettledEvent struct {
	Event
	Data DirectDebit `json:"data"`
}

// DirectDebitReturnedEvent represents the direct_debit.returned webhook event.
// DirectDebitReturnedEvent 表示 direct_debit.returned webhook 事件。
type DirectDebitReturnedEvent struct {
	Event
	Data DirectDebit `json:"data"`
}

// --- Event Data Structures ---

// DirectDebit represents the payload data for direct debit webhook events.
// DirectDebit 表示 direct debit webhook 事件的载荷数据。
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/direct-debit-payouts.md
type DirectDebit struct {
	AccountID       string              `json:"accountId"`
	DirectDebitID   string              `json:"directDebitId"`
	TransactionType string              `json:"transactionType"`
	DebtorAccount   *DirectDebitAccount `json:"debtorAccount,omitempty"`
	CreditorAccount *DirectDebitAccount `json:"creditorAccount,omitempty"`
	Amount          float64             `json:"amount"`
	Currency        string              `json:"currency"`
	Reference       string              `json:"reference"`
	Status          string              `json:"status"`
	RejectReason    *string             `json:"rejectReason,omitempty"`
	CreatedAt       string              `json:"createdAt"`
	LastUpdatedAt   string              `json:"lastUpdatedAt"`
}

// DirectDebitAccount represents a debtor or creditor account in a direct debit event.
// DirectDebitAccount 表示 direct debit 事件中的借方或贷方账户。
type DirectDebitAccount struct {
	ID     *string `json:"id,omitempty"`
	Name   string  `json:"name"`
	Type   *string `json:"type,omitempty"`
	Number string  `json:"number"`
}
