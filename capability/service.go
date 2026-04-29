package capability

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Capability API.
// Service 提供 Capability API 的访问。
type Service struct {
	doer sdk.Doer
}

// New creates a new Capability service instance.
// New 创建 Capability 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
