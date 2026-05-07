// Package events provides typed webhook event structures for the global accounts domain.
// Global Accounts 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/global-accounts.md
//
// 事件映射表:
//
//	global_account.active               → GlobalAccountActiveEvent         (Data: GlobalAccountEventData)
//	global_account.closed               → GlobalAccountClosedEvent         (Data: GlobalAccountEventData)
//	global_account.failed               → GlobalAccountFailedEvent         (Data: GlobalAccountEventData)
//	ga.new                              → GADepositNewEvent                (Data: GADepositNewEventData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Global Account Status Events ---

// GlobalAccountActiveEvent represents the global_account.active webhook event.
// GlobalAccountActiveEvent 表示 global_account.active webhook 事件。
type GlobalAccountActiveEvent struct {
	Event
	Data GlobalAccountEventData `json:"data"`
}

// GlobalAccountClosedEvent represents the global_account.closed webhook event.
// GlobalAccountClosedEvent 表示 global_account.closed webhook 事件。
type GlobalAccountClosedEvent struct {
	Event
	Data GlobalAccountEventData `json:"data"`
}

// GlobalAccountFailedEvent represents the global_account.failed webhook event.
// GlobalAccountFailedEvent 表示 global_account.failed webhook 事件。
type GlobalAccountFailedEvent struct {
	Event
	Data GlobalAccountEventData `json:"data"`
}

// GlobalAccountEventData represents the payload data for global account status events.
// GlobalAccountEventData 表示 global account 状态事件的载荷数据。
// 官方文档说明 data 为 Get global account by ID 的响应体。
type GlobalAccountEventData struct {
	AccountName                 string                      `json:"account_name"`
	AccountNumber               string                      `json:"account_number"`
	AccountType                 string                      `json:"account_type"`
	AlternateAccountIdentifiers AlternateAccountIdentifiers `json:"alternate_account_identifiers,omitempty"`
	CountryCode                 string                      `json:"country_code"`
	IBAN                        string                      `json:"iban,omitempty"`
	ID                          string                      `json:"id"`
	Institution                 Institution                 `json:"institution,omitempty"`
	NickName                    string                      `json:"nick_name,omitempty"`
	RequestID                   string                      `json:"request_id,omitempty"`
	RequiredFeatures            []RequiredFeature           `json:"required_features,omitempty"`
	Status                      string                      `json:"status"`
	CloseReason                 string                      `json:"close_reason,omitempty"`
	FailureReason               string                      `json:"failure_reason,omitempty"`
	SupportedFeatures           []SupportedFeature          `json:"supported_features,omitempty"`
	SwiftCode                   string                      `json:"swift_code,omitempty"`
}

// AlternateAccountIdentifiers represents alternate account identifiers in webhook payload.
// AlternateAccountIdentifiers 表示 webhook payload 中的备用账户标识。
type AlternateAccountIdentifiers struct {
	Email string `json:"email,omitempty"`
}

// Institution represents financial institution information in webhook payload.
// Institution 表示 webhook payload 中的金融机构信息。
type Institution struct {
	Address string `json:"address,omitempty"`
	City    string `json:"city,omitempty"`
	Name    string `json:"name,omitempty"`
	ZipCode string `json:"zip_code,omitempty"`
}

// RequiredFeature represents a required feature for the global account.
// RequiredFeature 表示全球账户的必需功能。
type RequiredFeature struct {
	Currency       string `json:"currency"`
	TransferMethod string `json:"transfer_method"`
}

// SupportedFeature represents a supported feature for the global account.
// SupportedFeature 表示全球账户的支持功能。
type SupportedFeature struct {
	Currency            string        `json:"currency"`
	LocalClearingSystem string        `json:"local_clearing_system,omitempty"`
	RoutingCodes        []RoutingCode `json:"routing_codes,omitempty"`
	TransferMethod      string        `json:"transfer_method"`
	Type                string        `json:"type"`
}

// RoutingCode represents a routing code for a supported feature.
// RoutingCode 表示支持功能的路由代码。
type RoutingCode struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// --- Global Account Deposit Events ---

// GADepositNewEvent represents the ga.new webhook event.
// GADepositNewEvent 表示 ga.new webhook 事件。
type GADepositNewEvent struct {
	Event
	Data GADepositNewEventData `json:"data"`
}

// GADepositNewEventData represents the payload data for ga.new event.
// GADepositNewEventData 表示 ga.new 事件的载荷数据。
// 官方文档说明 data 为 Get global account transactions 的响应体。
type GADepositNewEventData struct {
	HasMore bool                           `json:"has_more"`
	Items   []GlobalAccountTransactionItem `json:"items"`
}

// GlobalAccountTransactionItem represents a transaction item in the ga.new event.
// GlobalAccountTransactionItem 表示 ga.new 事件中的交易项。
type GlobalAccountTransactionItem struct {
	Amount       float64 `json:"amount"`
	CreateTime   string  `json:"create_time,omitempty"`
	Currency     string  `json:"currency"`
	Description  string  `json:"description,omitempty"`
	FeeAmount    float64 `json:"fee_amount,omitempty"`
	FeeCurrency  string  `json:"fee_currency,omitempty"`
	ID           string  `json:"id,omitempty"`
	PayerCountry string  `json:"payer_country,omitempty"`
	PayerName    string  `json:"payer_name,omitempty"`
	Status       string  `json:"status,omitempty"`
	Type         string  `json:"type,omitempty"`
}
