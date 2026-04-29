package scale

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Scale API.
// Service 提供 Scale API 的访问。
type Service struct {
	doer sdk.Doer
}

// New creates a new Scale service instance.
// New 创建 Scale 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
