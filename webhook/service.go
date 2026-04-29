package webhook

import "github.com/hakur/airwallex/sdk"

// Service provides webhook endpoint management functions.
// Service 提供 Webhook 端点管理功能。
type Service struct {
	doer sdk.Doer
}

// New creates a new Webhook service instance.
// New 创建 Webhook 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
