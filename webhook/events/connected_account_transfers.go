// Package events provides typed webhook event structures for the connected account transfers domain.
// Connected Account Transfers 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/connected-account-transfers.md
//
// 事件映射表:
//
//	connected_account_transfer.new       → ConnectedAccountTransferNewEvent       (Data: ConnectedAccountTransfer)
//	connected_account_transfer.pending   → ConnectedAccountTransferPendingEvent   (Data: ConnectedAccountTransfer)
//	connected_account_transfer.settled   → ConnectedAccountTransferSettledEvent   (Data: ConnectedAccountTransfer)
//	connected_account_transfer.suspended → ConnectedAccountTransferSuspendedEvent (Data: ConnectedAccountTransfer)
//	connected_account_transfer.failed    → ConnectedAccountTransferFailedEvent    (Data: ConnectedAccountTransfer)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Connected Account Transfer Events ---

// ConnectedAccountTransferNewEvent represents the connected_account_transfer.new webhook event.
// ConnectedAccountTransferNewEvent 表示 connected_account_transfer.new webhook 事件。
type ConnectedAccountTransferNewEvent struct {
	Event
	Data ConnectedAccountTransfer `json:"data"`
}

// ConnectedAccountTransferPendingEvent represents the connected_account_transfer.pending webhook event.
// ConnectedAccountTransferPendingEvent 表示 connected_account_transfer.pending webhook 事件。
type ConnectedAccountTransferPendingEvent struct {
	Event
	Data ConnectedAccountTransfer `json:"data"`
}

// ConnectedAccountTransferSettledEvent represents the connected_account_transfer.settled webhook event.
// ConnectedAccountTransferSettledEvent 表示 connected_account_transfer.settled webhook 事件。
type ConnectedAccountTransferSettledEvent struct {
	Event
	Data ConnectedAccountTransfer `json:"data"`
}

// ConnectedAccountTransferSuspendedEvent represents the connected_account_transfer.suspended webhook event.
// ConnectedAccountTransferSuspendedEvent 表示 connected_account_transfer.suspended webhook 事件。
type ConnectedAccountTransferSuspendedEvent struct {
	Event
	Data ConnectedAccountTransfer `json:"data"`
}

// ConnectedAccountTransferFailedEvent represents the connected_account_transfer.failed webhook event.
// ConnectedAccountTransferFailedEvent 表示 connected_account_transfer.failed webhook 事件。
type ConnectedAccountTransferFailedEvent struct {
	Event
	Data ConnectedAccountTransfer `json:"data"`
}

// --- Event Data Structures ---

// ConnectedAccountTransfer represents the payload data for connected account transfer webhook events.
// ConnectedAccountTransfer 表示 connected account transfer webhook 事件的载荷数据。
// 官方文档说明 data 为 Get a connected account transfer by ID 的响应体。
type ConnectedAccountTransfer struct {
	Amount      int    `json:"amount"`
	CreatedAt   string `json:"created_at"`
	Currency    string `json:"currency"`
	Destination string `json:"destination"`
	Fee         int    `json:"fee"`
	ID          string `json:"id"`
	Reason      string `json:"reason,omitempty"`
	Reference   string `json:"reference,omitempty"`
	RequestID   string `json:"request_id"`
	Status      string `json:"status"`
	UpdatedAt   string `json:"updated_at"`
}
