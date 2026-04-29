package confirmation

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Confirmation API.
// Service 提供 Confirmation API 的访问。
type Service struct {
	doer sdk.Doer
}

// New creates a new Confirmation service instance.
// New 创建 Confirmation 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
