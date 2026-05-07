// Package events provides typed webhook event structures for the deposits domain.
// Deposits 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/deposits.md
//
// 事件映射表:
//
//	deposit.pending                     → DepositPendingEvent              (Data: Deposit)
//	deposit.settled                     → DepositSettledEvent              (Data: Deposit)
//	deposit.rejected                    → DepositRejectedEvent             (Data: Deposit)
//	deposit.reversed                    → DepositReversedEvent             (Data: Deposit)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Deposit Events ---

// DepositPendingEvent represents the deposit.pending webhook event.
type DepositPendingEvent struct {
	Event
	Data Deposit `json:"data"`
}

// DepositSettledEvent represents the deposit.settled webhook event.
type DepositSettledEvent struct {
	Event
	Data Deposit `json:"data"`
}

// DepositRejectedEvent represents the deposit.rejected webhook event.
type DepositRejectedEvent struct {
	Event
	Data Deposit `json:"data"`
}

// DepositReversedEvent represents the deposit.reversed webhook event.
type DepositReversedEvent struct {
	Event
	Data Deposit `json:"data"`
}

// Deposit represents a deposit object in webhook payload.
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/deposits.md
type Deposit struct {
	ID                    string                 `json:"id"`
	Type                  string                 `json:"type"`
	Status                string                 `json:"status"`
	CreatedAt             string                 `json:"created_at"`
	Currency              string                 `json:"currency"`
	Amount                float64                `json:"amount"`
	GlobalAccountID       string                 `json:"global_account_id,omitempty"`
	Payer                 *DepositPayer          `json:"payer,omitempty"`
	FundingSourceID       string                 `json:"funding_source_id,omitempty"`
	SettledAt             string                 `json:"settled_at,omitempty"`
	Fee                   *DepositFee            `json:"fee,omitempty"`
	FailureDetails        *DepositFailureDetails `json:"failure_details,omitempty"`
	Reference             string                 `json:"reference"`
	ProviderTransactionID string                 `json:"provider_transaction_id"`
}

// DepositPayer represents payer information in a deposit event.
type DepositPayer struct {
	Name        string              `json:"name"`
	CountryCode string              `json:"country_code"`
	BankAccount *DepositBankAccount `json:"bank_account,omitempty"`
}

// DepositBankAccount represents bank account details in a deposit event.
type DepositBankAccount struct {
	Type        string              `json:"type"`
	AUBSB       *DepositAUBSB       `json:"au_bsb,omitempty"`
	Name        string              `json:"name"`
	Institution *DepositInstitution `json:"institution,omitempty"`
}

// DepositAUBSB represents Australian BSB account details.
type DepositAUBSB struct {
	BSB           string `json:"bsb"`
	AccountNumber string `json:"account_number"`
}

// DepositInstitution represents the bank institution details.
type DepositInstitution struct {
	SWIFTCode   string `json:"swift_code"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
}

// DepositFee represents fee details in a settled deposit event.
type DepositFee struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

// DepositFailureDetails represents failure details in a rejected/reversed deposit event.
type DepositFailureDetails struct {
	Code                   string                         `json:"code"`
	ISOCode                string                         `json:"iso_code"`
	ProviderFailureDetails *DepositProviderFailureDetails `json:"provider_failure_details,omitempty"`
}

// DepositProviderFailureDetails represents provider-specific failure details.
type DepositProviderFailureDetails struct {
	Code                string `json:"code"`
	LocalClearingSystem string `json:"local_clearing_system"`
	Message             string `json:"message"`
}
