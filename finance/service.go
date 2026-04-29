package finance

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Finance API.
// Service 提供 Finance API 的访问。
type Service struct {
	// doer HTTP 请求执行器。
	doer sdk.Doer
}

// New creates a new Finance service instance.
// New 创建 Finance 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
