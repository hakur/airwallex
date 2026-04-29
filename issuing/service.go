package issuing

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Issuing API.
// Service 提供 Issuing API 的访问。
type Service struct {
	// doer is the HTTP request executor. Required.
	// doer HTTP请求执行器。必填。
	doer sdk.Doer
}

// New creates a new Issuing service instance.
// New 创建 Issuing 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
