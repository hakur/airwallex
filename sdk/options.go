package sdk

import "resty.dev/v3"

// Option 用于配置 Client 的初始化选项。
type Option func(*ClientOptions)

// ClientOptions 是 Client 的内部配置结构，供 Option 修改。
type ClientOptions struct {
	BaseURL    string
	HTTPClient *resty.Client
	OnBehalfOf string
	APIVersion string
	LoginAs    string
	Debug      bool // 启用 Debug 模式，打印 HTTP 请求/响应详情
}

// WithBaseURL 设置 API 基础地址，默认为 Production。
// Sandbox 使用示例：WithBaseURL(sdk.SandboxURL)
func WithBaseURL(baseURL string) Option {
	return func(o *ClientOptions) {
		o.BaseURL = baseURL
	}
}

// WithHTTPClient 替换默认的 HTTP 客户端，用于自定义超时、代理等。
func WithHTTPClient(httpClient *resty.Client) Option {
	return func(o *ClientOptions) {
		o.HTTPClient = httpClient
	}
}

// WithOnBehalfOf 设置默认的 x-on-behalf-of header，用于平台多商户场景。
// 所有后续请求都会自动携带此 header，单次请求可通过 RequestOption 覆盖。
func WithOnBehalfOf(accountID string) Option {
	return func(o *ClientOptions) {
		o.OnBehalfOf = accountID
	}
}

func WithAPIVersion(apiVersion string) Option {
	return func(o *ClientOptions) {
		o.APIVersion = apiVersion
	}
}

// WithLoginAs 设置认证请求中的 x-login-as header，用于 scoped API key 的多账户场景。
func WithLoginAs(accountID string) Option {
	return func(o *ClientOptions) {
		o.LoginAs = accountID
	}
}

// WithDebug 启用 Debug 模式，打印 HTTP 请求和响应详情，用于排查问题。
func WithDebug(debug bool) Option {
	return func(o *ClientOptions) {
		o.Debug = debug
	}
}
