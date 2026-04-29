package airwallex

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sasha-s/go-deadlock"
	"resty.dev/v3"
)

// authResponse 是 authentication/login 的响应结构。
type authResponse struct {
	Token string `json:"token"`
}

// Authenticator 管理 Bearer Token 的生命周期。
type Authenticator struct {
	clientID string
	apiKey   string
	baseURL  string
	loginAs  string

	mu     deadlock.RWMutex
	token  string
	expiry time.Time
}

// NewAuthenticator 创建一个 Authenticator。
func NewAuthenticator(clientID, apiKey, baseURL, loginAs string) *Authenticator {
	return &Authenticator{
		clientID: clientID,
		apiKey:   apiKey,
		baseURL:  baseURL,
		loginAs:  loginAs,
	}
}

// Token 返回当前有效的 Bearer Token。
// 如果 token 将在 5 分钟内过期，会自动刷新。
func (a *Authenticator) Token(ctx context.Context) (string, error) {
	a.mu.RLock()
	if time.Until(a.expiry) > 5*time.Minute {
		t := a.token
		a.mu.RUnlock()
		return t, nil
	}
	a.mu.RUnlock()

	a.mu.Lock()
	defer a.mu.Unlock()

	// 双重检查，避免多个 goroutine 同时刷新
	if time.Until(a.expiry) > 5*time.Minute {
		return a.token, nil
	}

	return a.refresh(ctx)
}

// Invalidate 强制清除缓存的 token，下次请求会重新获取。
func (a *Authenticator) Invalidate() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.token = ""
	a.expiry = time.Time{}
}

// refresh 调用 Airwallex authentication/login 端点获取新 token。
func (a *Authenticator) refresh(ctx context.Context) (string, error) {
	req := resty.New().R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("x-client-id", a.clientID).
		SetHeader("x-api-key", a.apiKey)

	if a.loginAs != "" {
		req.SetHeader("x-login-as", a.loginAs)
	}

	resp, err := req.Post(a.baseURL + "/api/v1/authentication/login")
	if err != nil {
		return "", fmt.Errorf("auth login request failed: %w", err)
	}

	if resp.IsError() {
		return "", fmt.Errorf("auth login failed: status=%d body=%s", resp.StatusCode(), resp.String())
	}

	var authResp authResponse
	if err := json.Unmarshal(resp.Bytes(), &authResp); err != nil {
		return "", fmt.Errorf("auth login response parse failed: %w", err)
	}

	if authResp.Token == "" {
		return "", fmt.Errorf("auth login returned empty token")
	}

	// Airwallex token 默认有效期约 30 分钟，保守设置 25 分钟
	a.token = authResp.Token
	a.expiry = time.Now().Add(25 * time.Minute)

	return a.token, nil
}
