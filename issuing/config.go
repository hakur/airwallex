package issuing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// Config represents issuing configuration information.
// Config 表示发卡配置信息。
type Config struct {
	// BlockedTransactionUsages is the list of blocked transaction usages. Optional.
	// BlockedTransactionUsages 被阻止的交易用途列表。可选。
	BlockedTransactionUsages []string `json:"blocked_transaction_usages,omitempty"`
	// EnableAutoConversion indicates whether auto conversion is enabled. Optional.
	// EnableAutoConversion 是否启用自动转换。可选。
	EnableAutoConversion bool `json:"enable_auto_conversion,omitempty"`
	// PrimaryCurrency is the primary currency. Optional.
	// PrimaryCurrency 主币种。可选。
	PrimaryCurrency sdk.Currency `json:"primary_currency,omitempty"`
	// RemoteAuthSettings contains remote authentication settings. Optional.
	// RemoteAuthSettings 远程认证设置。可选。
	RemoteAuthSettings map[string]any `json:"remote_auth_settings,omitempty"`
	// RemoteCallConfig contains remote call configuration. Optional.
	// RemoteCallConfig 远程通话配置。可选。
	RemoteCallConfig map[string]any `json:"remote_call_config,omitempty"`
	// RemoteProvisioningConfig contains remote provisioning configuration. Optional.
	// RemoteProvisioningConfig 远程配发配置。可选。
	RemoteProvisioningConfig map[string]any `json:"remote_provisioning_config,omitempty"`
	// SpendingLimitSettings contains spending limit settings. Optional.
	// SpendingLimitSettings 消费限额设置。可选。
	SpendingLimitSettings map[string]any `json:"spending_limit_settings,omitempty"`
}

// GetConfig retrieves issuing configuration.
// 官方文档: https://www.airwallex.com/docs/api/issuing/config/retrieve.md
// GetConfig 获取发卡配置。
func (s *Service) GetConfig(ctx context.Context, opts ...sdk.RequestOption) (*Config, error) {
	var resp Config
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/config", nil, &resp, opts...)
	return &resp, err
}

// UpdateConfigRequest represents the request to update issuing configuration.
// UpdateConfigRequest 更新发卡配置请求。
type UpdateConfigRequest struct {
	// EnableAutoConversion indicates whether auto conversion is enabled. Optional.
	// EnableAutoConversion 是否启用自动转换。可选。
	EnableAutoConversion bool `json:"enable_auto_conversion,omitempty"`
	// PrimaryCurrency is the primary currency. Optional.
	// PrimaryCurrency 主币种。可选。
	PrimaryCurrency sdk.Currency `json:"primary_currency,omitempty"`
	// RemoteAuth contains remote authentication configuration. Optional.
	// RemoteAuth 远程认证配置。可选。
	RemoteAuth map[string]any `json:"remote_auth,omitempty"`
	// RemoteCallConfig contains remote call configuration. Optional.
	// RemoteCallConfig 远程通话配置。可选。
	RemoteCallConfig map[string]any `json:"remote_call_config,omitempty"`
	// RemoteProvisioningConfig contains remote provisioning configuration. Optional.
	// RemoteProvisioningConfig 远程配发配置。可选。
	RemoteProvisioningConfig map[string]any `json:"remote_provisioning_config,omitempty"`
}

// UpdateConfig updates issuing configuration.
// 官方文档: https://www.airwallex.com/docs/api/issuing/config/update.md
// UpdateConfig 更新发卡配置。
func (s *Service) UpdateConfig(ctx context.Context, req *UpdateConfigRequest, opts ...sdk.RequestOption) (*Config, error) {
	var resp Config
	err := s.doer.Do(ctx, "POST", "/api/v1/issuing/config/update", req, &resp, opts...)
	return &resp, err
}
