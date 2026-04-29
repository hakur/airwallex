package spend

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Spend API.
// Service 提供 Spend API 的访问。
type Service struct {
	doer sdk.Doer
}

// New creates a new Spend service instance.
// New 创建 Spend 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
