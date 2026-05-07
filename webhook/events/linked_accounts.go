// Package events provides typed webhook event structures for the linked accounts domain.
// Linked Accounts 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/linked-accounts.md
//
// 事件映射表:
//
//	linked_account.requires_action      → LinkedAccountRequiresActionEvent (Data: LinkedAccountEventData)
//	linked_account.processing           → LinkedAccountProcessingEvent     (Data: LinkedAccountEventData)
//	linked_account.succeeded            → LinkedAccountSucceededEvent      (Data: LinkedAccountEventData)
//	linked_account.failed               → LinkedAccountFailedEvent         (Data: LinkedAccountEventData)
//	linked_account.suspended            → LinkedAccountSuspendedEvent      (Data: LinkedAccountEventData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Linked Account Status Events ---

// LinkedAccountRequiresActionEvent represents the linked_account.requires_action webhook event.
// LinkedAccountRequiresActionEvent 表示 linked_account.requires_action webhook 事件。
type LinkedAccountRequiresActionEvent struct {
	Event
	Data LinkedAccountEventData `json:"data"`
}

// LinkedAccountProcessingEvent represents the linked_account.processing webhook event.
// LinkedAccountProcessingEvent 表示 linked_account.processing webhook 事件。
type LinkedAccountProcessingEvent struct {
	Event
	Data LinkedAccountEventData `json:"data"`
}

// LinkedAccountSucceededEvent represents the linked_account.succeeded webhook event.
// LinkedAccountSucceededEvent 表示 linked_account.succeeded webhook 事件。
type LinkedAccountSucceededEvent struct {
	Event
	Data LinkedAccountEventData `json:"data"`
}

// LinkedAccountFailedEvent represents the linked_account.failed webhook event.
// LinkedAccountFailedEvent 表示 linked_account.failed webhook 事件。
type LinkedAccountFailedEvent struct {
	Event
	Data LinkedAccountEventData `json:"data"`
}

// LinkedAccountSuspendedEvent represents the linked_account.suspended webhook event.
// LinkedAccountSuspendedEvent 表示 linked_account.suspended webhook 事件。
type LinkedAccountSuspendedEvent struct {
	Event
	Data LinkedAccountEventData `json:"data"`
}

// LinkedAccountEventData represents the payload data for linked account status events.
// LinkedAccountEventData 表示 linked account 状态事件的载荷数据。
type LinkedAccountEventData struct {
	AUBank              AUBank                  `json:"au_bank,omitempty"`
	Capabilities        Capabilities            `json:"capabilities,omitempty"`
	ID                  string                  `json:"id"`
	NextAction          LinkedAccountNextAction `json:"next_action,omitempty"`
	Status              string                  `json:"status"`
	SupportedCurrencies []string                `json:"supported_currencies,omitempty"`
	Type                string                  `json:"type"`
}

// AUBank represents Australian bank account details in the linked account payload.
// AUBank 表示 linked account payload 中的澳大利亚银行账户详情。
type AUBank struct {
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	BSB           string `json:"bsb"`
	Currency      string `json:"currency"`
	EntityType    string `json:"entity_type"`
}

// Capabilities represents the capabilities of a linked account.
// Capabilities 表示 linked account 的功能。
type Capabilities struct {
	BalanceCheck       bool `json:"balance_check"`
	DirectDebitDeposit bool `json:"direct_debit_deposit"`
}

// LinkedAccountNextAction represents the next action required for a linked account.
// LinkedAccountNextAction 表示 linked account 所需的下一步操作。
type LinkedAccountNextAction struct {
	Type string `json:"type"`
}
