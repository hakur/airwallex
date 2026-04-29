package scale

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// HostedFlowStatus represents a hosted flow status.
// HostedFlowStatus 托管流程状态。
type HostedFlowStatus string

const (
	HostedFlowStatusInit       HostedFlowStatus = "INIT"
	HostedFlowStatusProcessing HostedFlowStatus = "PROCESSING"
	HostedFlowStatusFinished   HostedFlowStatus = "FINISHED"
)

// HostedFlow represents a hosted flow instance.
// HostedFlow 表示托管流程实例。
type HostedFlow struct {
	ID       string           `json:"id,omitempty"`
	Metadata map[string]any   `json:"metadata,omitempty"`
	Result   map[string]any   `json:"result,omitempty"`
	Status   HostedFlowStatus `json:"status,omitempty"`
	URL      string           `json:"url,omitempty"`
}

// CreateHostedFlowRequest represents a request to create a hosted flow.
// CreateHostedFlowRequest 创建托管流程请求。
type CreateHostedFlowRequest struct {
	AccountID string         `json:"account_id"`
	Template  string         `json:"template"`
	ReturnURL string         `json:"return_url,omitempty"`
	ErrorURL  string         `json:"error_url,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

// AuthorizeHostedFlowRequest represents a request to authorize a hosted flow.
// AuthorizeHostedFlowRequest 授权托管流程请求。
type AuthorizeHostedFlowRequest struct {
	Identity string `json:"identity,omitempty"`
}

// AuthorizeHostedFlowResponse represents the response for authorizing a hosted flow.
// AuthorizeHostedFlowResponse 授权托管流程响应。
type AuthorizeHostedFlowResponse struct {
	ID                string           `json:"id,omitempty"`
	AuthorizationCode string           `json:"authorization_code,omitempty"`
	Metadata          map[string]any   `json:"metadata,omitempty"`
	Result            map[string]any   `json:"result,omitempty"`
	Status            HostedFlowStatus `json:"status,omitempty"`
	URL               string           `json:"url,omitempty"`
}

// CreateHostedFlow creates a new hosted flow.
// CreateHostedFlow 创建托管流程。
// 官方文档: https://www.airwallex.com/docs/api/scale/hosted_flows/create.md
func (s *Service) CreateHostedFlow(ctx context.Context, req *CreateHostedFlowRequest, opts ...sdk.RequestOption) (*HostedFlow, error) {
	var resp HostedFlow
	err := s.doer.Do(ctx, "POST", "/api/v1/hosted_flows/create", req, &resp, opts...)
	return &resp, err
}

// GetHostedFlow retrieves the details of a hosted flow.
// GetHostedFlow 获取托管流程详情。
// 官方文档: https://www.airwallex.com/docs/api/scale/hosted_flows/retrieve.md
func (s *Service) GetHostedFlow(ctx context.Context, id string, opts ...sdk.RequestOption) (*HostedFlow, error) {
	var resp HostedFlow
	err := s.doer.Do(ctx, "GET", "/api/v1/hosted_flows/"+id, nil, &resp, opts...)
	return &resp, err
}

// AuthorizeHostedFlow authorizes a user to access a hosted flow.
// AuthorizeHostedFlow 授权用户访问托管流程。
// 官方文档: https://www.airwallex.com/docs/api/scale/hosted_flows/authorize.md
func (s *Service) AuthorizeHostedFlow(ctx context.Context, id string, req *AuthorizeHostedFlowRequest, opts ...sdk.RequestOption) (*AuthorizeHostedFlowResponse, error) {
	var resp AuthorizeHostedFlowResponse
	err := s.doer.Do(ctx, "POST", "/api/v1/hosted_flows/"+id+"/authorize", req, &resp, opts...)
	return &resp, err
}

// --- Deprecated ---
// ListHostedFlows 已移除。官方 API 不支持列出托管流程。
