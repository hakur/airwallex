package fx

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Transactional FX API.
// Service 提供 Transactional FX API 的访问。
type Service struct {
	// doer is the HTTP request executor. Required.
	// doer HTTP 请求执行器。必填。
	doer sdk.Doer
}

// New creates a new Transactional FX service instance.
// New 创建 Transactional FX 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
