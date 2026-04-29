package payouts

import "github.com/hakur/airwallex/sdk"

// Service provides access to the Payouts API.
// Service 提供 Payouts API 的访问。
type Service struct {
	// doer is the HTTP request executor.
	// doer HTTP请求执行器。
	doer sdk.Doer
}

// New creates a Payouts service instance.
// New 创建 Payouts 服务实例。
func New(doer sdk.Doer) *Service {
	return &Service{doer: doer}
}
