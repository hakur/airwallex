package billing

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Billing API.
// Service 提供 Billing API 的访问。
type Service struct {
	doer sdk.Doer
}

// New creates a new Billing service instance.
// New 创建 Billing 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
