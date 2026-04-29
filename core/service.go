package core

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Core Resources API.
// Service 提供 Core Resources API 的访问。
type Service struct {
	// doer is the HTTP request executor. / doer HTTP 请求执行器。
	doer sdk.Doer
}

// New creates a new Core Resources service instance.
// New 创建 Core Resources 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
