package webhook

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// Webhook represents a webhook configuration.
// Webhook 表示 webhook 配置。
type Webhook struct {
	// ID is the unique identifier.
	// ID 唯一标识符。
	ID string `json:"id"`
	// URL is the target endpoint for event delivery. Required.
	// URL 事件发送目标端点。必填。
	URL string `json:"url"`
	// Secret is used to verify events originate from Airwallex.
	// Secret 用于验证事件来自 Airwallex 的密钥。
	Secret string `json:"secret,omitempty"`
	// Version is the API version controlling event payload structure. Format: YYYY-MM-DD.
	// Version API 版本，控制事件载荷结构。格式 YYYY-MM-DD。
	Version string `json:"version"`
	// Events is the list of subscribed events. Immutable after creation.
	// Events 订阅的事件列表。创建后不可修改。
	Events []string `json:"events"`
	// CreatedAt is the creation time.
	// CreatedAt 创建时间。
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the last update time.
	// UpdatedAt 最后更新时间。
	UpdatedAt string `json:"updated_at"`
	// RequestID is the client-specified idempotency request ID.
	// RequestID 客户端指定的幂等请求 ID。
	RequestID string `json:"request_id,omitempty"`
}

// CreateWebhookRequest represents a request to create a webhook.
// CreateWebhookRequest 创建 webhook 请求。
type CreateWebhookRequest struct {
	// RequestID is the client-specified idempotency request ID. Required, max 50 chars.
	// RequestID 客户端指定的幂等请求 ID。必填，最长 50 字符。
	RequestID string `json:"request_id"`
	// URL is the target endpoint for event delivery. Required.
	// URL 事件发送目标端点。必填。
	URL string `json:"url"`
	// Version is the API version for event payload structure. Required. Format: YYYY-MM-DD.
	// Version API 版本，控制事件载荷结构。必填，格式 YYYY-MM-DD。
	Version string `json:"version"`
	// Events is the list of events to subscribe to. Required.
	// Events 订阅的事件列表。必填。
	Events []string `json:"events"`
}

// UpdateWebhookRequest represents a request to update a webhook. Only URL can be updated.
// UpdateWebhookRequest 更新 webhook 请求。仅 URL 可更新。
type UpdateWebhookRequest struct {
	// URL is the target endpoint for event delivery. Required.
	// URL 事件发送目标端点。必填。
	URL string `json:"url"`
}

// DeleteWebhookResponse represents the response for deleting a webhook.
// DeleteWebhookResponse 删除 webhook 响应。
type DeleteWebhookResponse struct {
	// ID is the webhook object ID.
	// ID Webhook 对象 ID。
	ID string `json:"id"`
	// Deleted indicates whether the deletion was successful.
	// Deleted 是否成功删除。
	Deleted bool `json:"deleted"`
}

// ListWebhooksResponse represents the webhook list response (cursor pagination).
// ListWebhooksResponse webhook 列表响应（cursor 分页）。
type ListWebhooksResponse struct {
	// Items is the list of webhooks on the current page.
	// Items 当前页结果。
	Items []Webhook `json:"items"`
	// PageBefore is the cursor for the previous page.
	// PageBefore 上一页 cursor。
	PageBefore string `json:"page_before,omitempty"`
	// PageAfter is the cursor for the next page.
	// PageAfter 下一页 cursor。
	PageAfter string `json:"page_after,omitempty"`
}

// CreateWebhook creates a webhook.
// CreateWebhook 创建 webhook。
// 官方文档: https://www.airwallex.com/docs/api/webhook/webhooks/create.md
func (s *Service) CreateWebhook(ctx context.Context, req *CreateWebhookRequest, opts ...sdk.RequestOption) (*Webhook, error) {
	var resp Webhook
	err := s.doer.Do(ctx, "POST", "/api/v1/webhooks/create", req, &resp, opts...)
	return &resp, err
}

// GetWebhook retrieves a webhook by ID.
// GetWebhook 根据 ID 获取 webhook。
// 官方文档: https://www.airwallex.com/docs/api/webhook/webhooks/retrieve.md
func (s *Service) GetWebhook(ctx context.Context, id string, opts ...sdk.RequestOption) (*Webhook, error) {
	var resp Webhook
	err := s.doer.Do(ctx, "GET", "/api/v1/webhooks/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateWebhook updates a webhook (only URL can be updated).
// UpdateWebhook 更新 webhook（仅 URL 可更新）。
// 官方文档: https://www.airwallex.com/docs/api/webhook/webhooks/update.md
func (s *Service) UpdateWebhook(ctx context.Context, id string, req *UpdateWebhookRequest, opts ...sdk.RequestOption) (*Webhook, error) {
	var resp Webhook
	err := s.doer.Do(ctx, "POST", "/api/v1/webhooks/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// DeleteWebhook deletes a webhook.
// DeleteWebhook 删除 webhook。
// 官方文档: https://www.airwallex.com/docs/api/webhook/webhooks/delete.md
func (s *Service) DeleteWebhook(ctx context.Context, id string, opts ...sdk.RequestOption) (*DeleteWebhookResponse, error) {
	var resp DeleteWebhookResponse
	err := s.doer.Do(ctx, "POST", "/api/v1/webhooks/"+id+"/delete", nil, &resp, opts...)
	return &resp, err
}

// ListWebhooks lists webhooks.
// ListWebhooks 列出 webhook。
// 官方文档: https://www.airwallex.com/docs/api/webhook/webhooks/list.md
func (s *Service) ListWebhooks(ctx context.Context, req *ListWebhooksRequest, opts ...sdk.RequestOption) (*ListWebhooksResponse, error) {
	var resp ListWebhooksResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/webhooks", req, &resp, opts...)
	return &resp, err
}

// ListWebhooksRequest represents query parameters for listing webhooks.
// ListWebhooksRequest webhook 列表查询参数。
type ListWebhooksRequest struct {
	// Page is the cursor for pagination. Use page_before for previous, page_after for next.
	// Page cursor，取上一页用 page_before，下一页用 page_after。
	Page string `json:"page,omitempty"`
	// PageSize is the number of items per page. Defaults to 20.
	// PageSize 每页数量，默认 20。
	PageSize int32 `json:"page_size,omitempty"`
}
