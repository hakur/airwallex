package webhook

import (
	"encoding/json"
	"fmt"

	"github.com/hakur/airwallex/webhook/events"
)

// UnmarshalDataObject extracts the "object" field from Event.Data and unmarshals it into a value of the specified type.
// UnmarshalDataObject 从 Event.Data 中提取 "object" 字段并解析为指定类型。
// This is the recommended way to parse events from domains that wrap data in {"object": {...}},
// such as online-payments, account, billing, etc.
// 这是解析 data 被包装在 {"object": {...}} 中的事件的推荐方式，如 online-payments、account、billing 等领域。
func UnmarshalDataObject[T any](evt *events.Event) (*T, error) {
	var wrapper struct {
		Object json.RawMessage `json:"object"`
	}
	if err := json.Unmarshal(evt.Data, &wrapper); err != nil {
		return nil, fmt.Errorf("webhook: unmarshal data wrapper: %w", err)
	}
	var v T
	if err := json.Unmarshal(wrapper.Object, &v); err != nil {
		return nil, fmt.Errorf("webhook: unmarshal event data object: %w", err)
	}
	return &v, nil
}

// UnmarshalData unmarshals Event.Data directly into a value of the specified type.
// UnmarshalData 将 Event.Data 直接解析为指定类型的值。
// For backward compatibility, this function first attempts to extract "object" from data,
// and falls back to direct parsing if extraction fails.
// This is suitable for domains with flat data structures like issuing.
// 为保持向后兼容，此函数先尝试从 data 中提取 "object"，失败则直接解析。
// 适用于 issuing 等使用平铺数据结构的领域。
func UnmarshalData[T any](evt *events.Event) (*T, error) {
	// Try to extract object first (for object-wrapped data)
	var wrapper struct {
		Object json.RawMessage `json:"object"`
	}
	if err := json.Unmarshal(evt.Data, &wrapper); err == nil && len(wrapper.Object) > 0 {
		var v T
		if err := json.Unmarshal(wrapper.Object, &v); err == nil {
			return &v, nil
		}
	}

	// Fallback: direct parse (for flat data)
	var v T
	if err := json.Unmarshal(evt.Data, &v); err != nil {
		return nil, fmt.Errorf("webhook: unmarshal event data: %w", err)
	}
	return &v, nil
}

// ParseEvent parses a webhook payload into an Event struct.
// ParseEvent 将 webhook payload 解析为 Event 结构体。
func ParseEvent(payload []byte) (*events.Event, error) {
	var evt events.Event
	if err := json.Unmarshal(payload, &evt); err != nil {
		return nil, fmt.Errorf("webhook: unmarshal event: %w", err)
	}
	return &evt, nil
}
