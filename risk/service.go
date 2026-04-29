package risk

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Risk API.
// Service 提供 Risk API 的访问。
type Service struct {
	doer sdk.Doer
}

// New creates a new Risk service instance.
// New 创建 Risk 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
