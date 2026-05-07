// Package events provides typed webhook event structures for the rfi domain.
// RFI 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/rfi.md
//
// 事件映射表:
//
//	rfi.action_required                     → RFIActionRequiredEvent           (Data: RFIActionRequiredEventData)
//	rfi.answered                            → RFIAnsweredEvent                 (Data: RFIAnsweredEventData)
//	rfi.closed                              → RFIClosedEvent                   (Data: RFIClosedEventData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/rfi.md
// RFI webhook events for request for information status changes.

// RFIActionRequiredEvent occurs when RFI is pending answers.
// RFIActionRequiredEvent 在 RFI 等待回答时触发。
type RFIActionRequiredEvent struct {
	Event
	Data RFIActionRequiredEventData `json:"data"`
}

// RFIActionRequiredEventData contains the payload for rfi.action_required event.
// RFIActionRequiredEventData 包含 rfi.action_required 事件的 payload。
type RFIActionRequiredEventData struct {
	ID string `json:"id"`
}

// RFIAnsweredEvent occurs when RFI is answered.
// RFIAnsweredEvent 在 RFI 已回答时触发。
type RFIAnsweredEvent struct {
	Event
	Data RFIAnsweredEventData `json:"data"`
}

// RFIAnsweredEventData contains the payload for rfi.answered event.
// RFIAnsweredEventData 包含 rfi.answered 事件的 payload。
type RFIAnsweredEventData struct {
	ID string `json:"id"`
}

// RFIClosedEvent occurs when RFI is finished and closed.
// RFIClosedEvent 在 RFI 完成并关闭时触发。
type RFIClosedEvent struct {
	Event
	Data RFIClosedEventData `json:"data"`
}

// RFIClosedEventData contains the payload for rfi.closed event.
// RFIClosedEventData 包含 rfi.closed 事件的 payload。
type RFIClosedEventData struct {
	ID string `json:"id"`
}
