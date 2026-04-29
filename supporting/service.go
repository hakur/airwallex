package supporting

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Supporting API.
// Service 提供 Supporting API 的访问。
type Service struct {
	doer sdk.Doer
}

// New creates a new Supporting service instance.
// New 创建 Supporting 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
