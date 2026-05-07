// Package events provides typed webhook event structures for the charges domain.
// Charges 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/charges.md
//
// 事件映射表:
//
//	charge.new       → ChargeNewEvent      (Data: ChargeEventData)
//	charge.pending   → ChargePendingEvent  (Data: ChargeEventData)
//	charge.settled   → ChargeSettledEvent  (Data: ChargeEventData)
//	charge.suspended → ChargeSuspendedEvent (Data: ChargeEventData)
//	charge.failed    → ChargeFailedEvent   (Data: ChargeEventData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Charge Events ---

// ChargeNewEvent represents the charge.new webhook event.
// ChargeNewEvent 表示 charge.new webhook 事件。
type ChargeNewEvent struct {
	Event
	Data ChargeEventData `json:"data"`
}

// ChargePendingEvent represents the charge.pending webhook event.
// ChargePendingEvent 表示 charge.pending webhook 事件。
type ChargePendingEvent struct {
	Event
	Data ChargeEventData `json:"data"`
}

// ChargeSettledEvent represents the charge.settled webhook event.
// ChargeSettledEvent 表示 charge.settled webhook 事件。
type ChargeSettledEvent struct {
	Event
	Data ChargeEventData `json:"data"`
}

// ChargeSuspendedEvent represents the charge.suspended webhook event.
// ChargeSuspendedEvent 表示 charge.suspended webhook 事件。
type ChargeSuspendedEvent struct {
	Event
	Data ChargeEventData `json:"data"`
}

// ChargeFailedEvent represents the charge.failed webhook event.
// ChargeFailedEvent 表示 charge.failed webhook 事件。
type ChargeFailedEvent struct {
	Event
	Data ChargeEventData `json:"data"`
}

// --- Event Data Structures ---

// ChargeEventData represents the payload data for charge webhook events.
// ChargeEventData 表示 charge webhook 事件的载荷数据。
// 官方文档说明 data 为 Get a charge by ID 的响应体。
type ChargeEventData struct {
	Amount    int    `json:"amount"`
	CreatedAt string `json:"createdAt"`
	Currency  string `json:"currency"`
	Fee       int    `json:"fee"`
	ID        string `json:"id"`
	Reason    string `json:"reason,omitempty"`
	Reference string `json:"reference,omitempty"`
	RequestID string `json:"requestId"`
	Source    string `json:"source,omitempty"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updatedAt"`
}
