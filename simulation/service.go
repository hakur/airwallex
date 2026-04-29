package simulation

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Simulation API.
// Service 提供 Simulation API 的访问。
type Service struct {
	doer sdk.Doer
}

// New creates a new Simulation service instance.
// New 创建 Simulation 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
