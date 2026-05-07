// Package events provides typed webhook event structures for the account capability domain.
// Account Capability 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/account-capability.md
//
// 事件映射表:
//
//	account_capability.enabled  → AccountCapabilityEnabledEvent  (Data: AccountCapabilityEventData)
//	account_capability.disabled → AccountCapabilityDisabledEvent (Data: AccountCapabilityEventData)
//	account_capability.pending  → AccountCapabilityPendingEvent  (Data: AccountCapabilityEventData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Account Capability Events ---

// AccountCapabilityEnabledEvent represents the account_capability.enabled webhook event.
// AccountCapabilityEnabledEvent 表示 account_capability.enabled webhook 事件。
type AccountCapabilityEnabledEvent struct {
	Event
	Data AccountCapabilityEventData `json:"data"`
}

// AccountCapabilityDisabledEvent represents the account_capability.disabled webhook event.
// AccountCapabilityDisabledEvent 表示 account_capability.disabled webhook 事件。
type AccountCapabilityDisabledEvent struct {
	Event
	Data AccountCapabilityEventData `json:"data"`
}

// AccountCapabilityPendingEvent represents the account_capability.pending webhook event.
// AccountCapabilityPendingEvent 表示 account_capability.pending webhook 事件。
type AccountCapabilityPendingEvent struct {
	Event
	Data AccountCapabilityEventData `json:"data"`
}

// --- Event Data Structures ---

// AccountCapabilityEventData represents the payload data for account capability status events.
// AccountCapabilityEventData 表示账户能力状态事件的载荷数据。
// 官方文档说明 data 为 Get account capability by ID 的响应体。
type AccountCapabilityEventData struct {
	// Comment is the optional remark for the capability.
	// Comment 能力备注说明。
	Comment string `json:"comment,omitempty"`
	// Details is the optional capability details.
	// Details 能力详情。
	Details *AccountCapabilityDetails `json:"details,omitempty"`
	// EntityType is the optional declarant entity type.
	// EntityType 申报实体类型。
	EntityType string `json:"entity_type,omitempty"`
	// ID is the unique identifier of the account capability.
	// ID 账户能力唯一标识符。
	ID string `json:"id"`
	// Status is the status of the account capability.
	// Status 账户能力状态。
	Status string `json:"status"`
	// UpdatedAt is the last update time.
	// UpdatedAt 最后更新时间。
	UpdatedAt string `json:"updated_at"`
}

// AccountCapabilityDetails represents capability details in webhook payload.
// AccountCapabilityDetails 表示 webhook payload 中的能力详情。
type AccountCapabilityDetails struct {
	// ReasonCodes is the list of reason codes.
	// ReasonCodes 原因代码列表。
	ReasonCodes []string `json:"reason_codes,omitempty"`
}
