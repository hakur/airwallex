package scale

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// InvitationLink represents an invitation link.
// InvitationLink 表示邀请链接。
type InvitationLink struct {
	ID                string                      `json:"id,omitempty"`
	AccountID         string                      `json:"account_id,omitempty"`
	Mode              string                      `json:"mode,omitempty"`
	Identifier        string                      `json:"identifier,omitempty"`
	Metadata          map[string]any              `json:"metadata,omitempty"`
	URL               string                      `json:"url,omitempty"`
	CreatedAt         string                      `json:"created_at,omitempty"`
	ExpireAt          string                      `json:"expire_at,omitempty"`
	OAuth2            *InvitationLinkOAuth2       `json:"oauth2,omitempty"`
	ScaleConnect      *InvitationLinkScaleConnect `json:"scale_connect,omitempty"`
	PrefilledFormData []PrefilledFormField        `json:"prefilled_formdata,omitempty"`
}

// InvitationLinkOAuth2 represents OAuth2 mode configuration.
// InvitationLinkOAuth2 OAuth2 模式配置。
type InvitationLinkOAuth2 struct {
	RedirectURI  string   `json:"redirect_uri"`
	ResponseType string   `json:"response_type"`
	Scope        []string `json:"scope"`
	State        string   `json:"state,omitempty"`
}

// InvitationLinkScaleConnect represents Scale Connect mode configuration.
// InvitationLinkScaleConnect Scale Connect 模式配置。
type InvitationLinkScaleConnect struct {
	RedirectURI string `json:"redirect_uri"`
}

// PrefilledFormField represents a prefilled form field.
// PrefilledFormField 预填表单字段。
type PrefilledFormField struct {
	Field    string `json:"field"`
	Value    string `json:"value"`
	Verified bool   `json:"verified,omitempty"`
	Editable bool   `json:"editable,omitempty"`
}

// CreateInvitationLinkRequest represents a request to create an invitation link.
// CreateInvitationLinkRequest 创建邀请链接请求。
type CreateInvitationLinkRequest struct {
	AccountID         string                      `json:"account_id,omitempty"`
	Mode              string                      `json:"mode"`
	OAuth2            *InvitationLinkOAuth2       `json:"oauth2,omitempty"`
	ScaleConnect      *InvitationLinkScaleConnect `json:"scale_connect,omitempty"`
	Metadata          map[string]any              `json:"metadata,omitempty"`
	Identifier        string                      `json:"identifier,omitempty"`
	PrefilledFormData []PrefilledFormField        `json:"prefilled_formdata,omitempty"`
}

// CreateInvitationLink creates an invitation link.
// CreateInvitationLink 创建邀请链接。
// 官方文档: https://www.airwallex.com/docs/api/scale/invitation_links/create.md
func (s *Service) CreateInvitationLink(ctx context.Context, req *CreateInvitationLinkRequest, opts ...sdk.RequestOption) (*InvitationLink, error) {
	var resp InvitationLink
	err := s.doer.Do(ctx, "POST", "/api/v1/accounts/invitation_links/create", req, &resp, opts...)
	return &resp, err
}

// GetInvitationLink retrieves the details of an invitation link.
// GetInvitationLink 获取邀请链接详情。
// 官方文档: https://www.airwallex.com/docs/api/scale/invitation_links/retrieve.md
func (s *Service) GetInvitationLink(ctx context.Context, id string, opts ...sdk.RequestOption) (*InvitationLink, error) {
	var resp InvitationLink
	err := s.doer.Do(ctx, "GET", "/api/v1/accounts/invitation_links/"+id, nil, &resp, opts...)
	return &resp, err
}

// --- Deprecated ---
// ListInvitationLinks 已移除。官方 API 不支持列出邀请链接。
