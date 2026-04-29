package airwallex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hakur/airwallex/billing"
	"github.com/hakur/airwallex/capability"
	"github.com/hakur/airwallex/confirmation"
	"github.com/hakur/airwallex/core"
	"github.com/hakur/airwallex/finance"
	"github.com/hakur/airwallex/fx"
	"github.com/hakur/airwallex/issuing"
	"github.com/hakur/airwallex/pa"
	"github.com/hakur/airwallex/payouts"
	"github.com/hakur/airwallex/risk"
	"github.com/hakur/airwallex/scale"
	"github.com/hakur/airwallex/sdk"
	"github.com/hakur/airwallex/simulation"
	"github.com/hakur/airwallex/spend"
	"github.com/hakur/airwallex/supporting"
	"github.com/hakur/airwallex/webhook"
	"github.com/joho/godotenv"
	"resty.dev/v3"
)

// NewFromEnv 从环境变量创建 Airwallex 客户端。
// 如果 envFilename 不为空，会先加载该 .env 文件（.env 中的值覆盖系统环境变量），
// 然后读取 AIRWALLEX_CLIENT_ID 和 AIRWALLEX_API_KEY。
func NewFromEnv(envFilename string, opts ...sdk.Option) (*Client, error) {
	if envFilename != "" {
		_ = godotenv.Load(envFilename)
	}

	if sdk.GetEnvClientID() == "" || sdk.GetEnvAPIKey() == "" {
		return nil, errors.New("AIRWALLEX_CLIENT_ID and AIRWALLEX_API_KEY must be set")
	}

	return New(sdk.GetEnvClientID(), sdk.GetEnvAPIKey(), opts...), nil
}

// New 创建一个 Airwallex 客户端。
func New(clientID, apiKey string, opts ...sdk.Option) *Client {
	opt := sdk.ClientOptions{}
	for _, o := range opts {
		o(&opt)
	}
	if opt.BaseURL == "" {
		opt.BaseURL = sdk.ServerURL
	}
	if opt.APIVersion == "" {
		opt.APIVersion = sdk.APIVersion
	}

	c := &Client{
		opt: opt,
	}

	// 如果没有替换 HTTP 客户端，创建默认的
	if opt.HTTPClient != nil {
		c.rc = opt.HTTPClient
	} else {
		c.rc = resty.New().SetBaseURL(opt.BaseURL)
	}

	// 初始化认证器
	c.auth = NewAuthenticator(clientID, apiKey, opt.BaseURL, opt.LoginAs)

	return c
}

// Client 是 Airwallex SDK 的根客户端，实现 sdk.Doer 接口。
type Client struct {
	rc   *resty.Client
	auth *Authenticator
	opt  sdk.ClientOptions
}

// Do 执行 HTTP 请求，实现 sdk.Doer 接口。
// 自动注入认证 header，处理 401 重试，解析 API 错误。
func (c *Client) Do(ctx context.Context, method, path string, req, resp any, opts ...sdk.RequestOption) error {
	return c.doWithRetry(ctx, method, path, req, resp, opts, true)
}

// doWithRetry 内部请求方法，支持一次 401 重试。
func (c *Client) doWithRetry(ctx context.Context, method, path string, req, resp any, opts []sdk.RequestOption, canRetry bool) error {
	token, err := c.auth.Token(ctx)
	if err != nil {
		return fmt.Errorf("get auth token: %w", err)
	}

	r := c.rc.R().SetContext(ctx).
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("x-api-version", c.opt.APIVersion).
		SetHeader("Accept", "application/json")

	if c.opt.OnBehalfOf != "" {
		r.SetHeader("x-on-behalf-of", c.opt.OnBehalfOf)
	}

	for _, opt := range opts {
		opt(r)
	}

	if req != nil {
		r.SetBody(req)
		if c.opt.Debug {
			// r.EnableDebug()
			if reqJSON, err := json.Marshal(req); err == nil {
				fmt.Printf("[Airwallex DEBUG] Request Body: %s\n", string(reqJSON))
			}
		}
	} else if method != "GET" && method != "DELETE" {
		// POST/PUT/PATCH 需要 body，避免 Content-Type 问题
		r.SetBody(map[string]any{})
	}

	// 如果启用了 Debug，禁用自动 JSON 解析，以便获取原始响应体
	if !c.opt.Debug && resp != nil {
		r.SetResult(resp)
	}

	var httpResp *resty.Response
	switch strings.ToUpper(method) {
	case http.MethodGet:
		httpResp, err = r.Get(path)
	case http.MethodPost:
		httpResp, err = r.Post(path)
	case http.MethodPut:
		httpResp, err = r.Put(path)
	case http.MethodPatch:
		httpResp, err = r.Patch(path)
	case http.MethodDelete:
		httpResp, err = r.Delete(path)
	default:
		return fmt.Errorf("unsupported http method: %s", method)
	}

	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}

	// 401 时尝试刷新 token 并重试一次
	if httpResp.StatusCode() == http.StatusUnauthorized && canRetry {
		c.auth.Invalidate()
		return c.doWithRetry(ctx, method, path, req, resp, opts, false)
	}

	if c.opt.Debug && httpResp != nil {
		fmt.Printf("[Airwallex DEBUG] Response Status: %d\n", httpResp.StatusCode())
		if body := httpResp.Bytes(); len(body) > 0 {
			fmt.Printf("[Airwallex DEBUG] Response Body: %s\n", string(body))
		}
	}

	if httpResp.IsError() {
		return c.parseError(httpResp)
	}

	// Debug 模式下手动解析 JSON（因为禁用了 resty 的自动解析）
	if c.opt.Debug && resp != nil {
		if body := httpResp.Bytes(); len(body) > 0 {
			if err := json.Unmarshal(body, resp); err != nil {
				return fmt.Errorf("unmarshal response: %w", err)
			}
		}
	}

	return nil
}

// parseError 将 Airwallex 错误响应解析为 sdk.APIError。
func (c *Client) parseError(resp *resty.Response) error {
	// 404 且响应不是 JSON 时，返回标准的 not_found 错误
	if resp.StatusCode() == http.StatusNotFound {
		var apiErr sdk.APIError
		if err := json.Unmarshal(resp.Bytes(), &apiErr); err != nil || apiErr.Code == "" {
			return &sdk.APIError{Code: string(sdk.ErrorCodeNotFound), Message: "Not Found"}
		}
		return &apiErr
	}

	var apiErr sdk.APIError
	if err := json.Unmarshal(resp.Bytes(), &apiErr); err != nil {
		return fmt.Errorf("http error status=%d body=%s", resp.StatusCode(), resp.String())
	}
	return &apiErr
}

// PA 返回 Payment Acceptance 服务。
func (c *Client) PA() *pa.Service {
	return pa.New(c)
}

// Payouts 返回 Payouts 服务。
func (c *Client) Payouts() *payouts.Service {
	return payouts.New(c)
}

// Core 返回 Core Resources 服务。
func (c *Client) Core() *core.Service {
	return core.New(c)
}

// Issuing 返回 Issuing 服务。
func (c *Client) Issuing() *issuing.Service {
	return issuing.New(c)
}

// FX 返回 Transactional FX 服务。
func (c *Client) FX() *fx.Service {
	return fx.New(c)
}

// Billing 返回 Billing 服务。
func (c *Client) Billing() *billing.Service {
	return billing.New(c)
}

// Finance 返回 Finance 服务。
func (c *Client) Finance() *finance.Service {
	return finance.New(c)
}

// Scale 返回 Scale 服务。
func (c *Client) Scale() *scale.Service {
	return scale.New(c)
}

// Risk 返回 Risk 服务。
func (c *Client) Risk() *risk.Service {
	return risk.New(c)
}

// Simulation 返回 Simulation 服务。
func (c *Client) Simulation() *simulation.Service {
	return simulation.New(c)
}

// Webhook 返回 Webhook 服务。
func (c *Client) Webhook() *webhook.Service {
	return webhook.New(c)
}

// Spend 返回 Spend 服务。
func (c *Client) Spend() *spend.Service {
	return spend.New(c)
}

// Supporting 返回 Supporting 服务。
func (c *Client) Supporting() *supporting.Service {
	return supporting.New(c)
}

// Capability 返回 Capability 服务。
func (c *Client) Capability() *capability.Service {
	return capability.New(c)
}

// Confirmation 返回 Confirmation 服务。
func (c *Client) Confirmation() *confirmation.Service {
	return confirmation.New(c)
}
