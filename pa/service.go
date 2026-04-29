package pa

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Payment Acceptance API.
// Service 提供 Payment Acceptance API 的访问。
type Service struct {
	// doer is the HTTP request executor. Required.
	// doer HTTP 请求执行器。必填。
	doer sdk.Doer
}

// New creates a Payment Acceptance service instance.
// New 创建 Payment Acceptance 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
