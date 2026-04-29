package sdk

import "resty.dev/v3"

// RequestOption 用于单次请求级别的配置，如覆盖 on-behalf-of header。
type RequestOption func(*resty.Request)

// WithRequestOnBehalfOf 返回一个 RequestOption，为单次请求设置 x-on-behalf-of header。
func WithRequestOnBehalfOf(accountID string) RequestOption {
	return func(r *resty.Request) {
		r.SetHeader("x-on-behalf-of", accountID)
	}
}

// WithRequestHeader sets a custom HTTP header for a single request.
// WithRequestHeader 为单次请求设置自定义 HTTP header。
func WithRequestHeader(key, value string) RequestOption {
	return func(r *resty.Request) {
		r.SetHeader(key, value)
	}
}
