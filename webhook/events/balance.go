// Package events provides typed webhook event structures for the balance domain.
// Balance 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/balance.md
//
// 事件映射表:
//
//	balance.va.top_up  → BalanceVATopUpEvent    (Data: BalanceVATopUpData)
//	balance.ga.top_up  → BalanceGATopUpEvent    (Data: BalanceGATopUpData)
//	balance.adjustment → BalanceAdjustmentEvent (Data: BalanceAdjustmentData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Balance: Virtual Account Top Up ---

// BalanceVATopUpEvent represents the balance.va.top_up webhook event.
// 当通过虚拟账户收到存款导致钱包余额增加时触发。
type BalanceVATopUpEvent struct {
	Event
	Data BalanceVATopUpData `json:"data"`
}

// BalanceVATopUpData is the payload data for balance.va.top_up event.
type BalanceVATopUpData struct {
	Amount               float64                       `json:"amount"`
	Balance              float64                       `json:"balance"`
	Ccy                  string                        `json:"ccy"`
	CreatedAt            string                        `json:"createdAt"`
	TransactionReference BalanceVATransactionReference `json:"transactionReference"`
	Type                 string                        `json:"type"`
}

// BalanceVATransactionReference represents the transaction reference for balance.va.top_up.
type BalanceVATransactionReference struct {
	DepositID string `json:"depositId"`
}

// --- Balance: Global Account Top Up ---

// BalanceGATopUpEvent represents the balance.ga.top_up webhook event.
// 当通过全球账户收到存款导致钱包余额增加时触发。
type BalanceGATopUpEvent struct {
	Event
	Data BalanceGATopUpData `json:"data"`
}

// BalanceGATopUpData is the payload data for balance.ga.top_up event.
type BalanceGATopUpData struct {
	Amount               float64                       `json:"amount"`
	Balance              float64                       `json:"balance"`
	Ccy                  string                        `json:"ccy"`
	CreatedAt            string                        `json:"createdAt"`
	Reference            string                        `json:"reference"`
	TransactionReference BalanceGATransactionReference `json:"transactionReference"`
	Type                 string                        `json:"type"`
}

// BalanceGATransactionReference represents the transaction reference for balance.ga.top_up.
type BalanceGATransactionReference struct {
	DepositID       string `json:"depositId"`
	GlobalAccountID string `json:"globalAccountId"`
	TransactionID   string `json:"transactionId"`
	PayerName       string `json:"payerName"`
	PayerBankName   string `json:"payerBankName"`
}

// --- Balance: Adjustment ---

// BalanceAdjustmentEvent represents the balance.adjustment webhook event.
// 当余额因调整而发生变化时触发。
type BalanceAdjustmentEvent struct {
	Event
	Data BalanceAdjustmentData `json:"data"`
}

// BalanceAdjustmentData is the payload data for balance.adjustment event.
type BalanceAdjustmentData struct {
	Amount               float64                               `json:"amount"`
	Balance              float64                               `json:"balance"`
	Ccy                  string                                `json:"ccy"`
	CreatedAt            string                                `json:"createdAt"`
	TransactionReference BalanceAdjustmentTransactionReference `json:"transactionReference"`
	Type                 string                                `json:"type"`
}

// BalanceAdjustmentTransactionReference represents the transaction reference for balance.adjustment.
type BalanceAdjustmentTransactionReference struct {
	AdjustmentID string `json:"adjustmentId"`
	Reason       string `json:"reason"`
}
