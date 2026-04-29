package sdk

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joho/godotenv"
)

const (
	EnvAPIKey      = "AIRWALLEX_API_KEY"
	EnvClientIDKey = "AIRWALLEX_CLIENT_ID"
)

var (
	// SandboxURL 沙箱地址 单元测试专用
	SandboxURL = "https://api-demo.airwallex.com"
	// ServerURL 正式服务器地址
	ServerURL = "https://api.airwallex.com"
	// APIVersion api 版本，不管是沙箱还是正式环境都是这个
	APIVersion = "2026-02-27"
)

// ListResult 是泛型列表响应包装器。
type ListResult[T any] struct {
	Items      []T    `json:"items"`
	HasMore    bool   `json:"has_more,omitempty"`
	TotalCount int    `json:"total_count,omitempty"`
	PageAfter  string `json:"page_after,omitempty"`
	PageBefore string `json:"page_before,omitempty"`
}

func GetEnvClientID() string {
	return os.Getenv(EnvClientIDKey)
}

func GetEnvAPIKey() string {
	return os.Getenv(EnvAPIKey)
}

// Doer 定义 HTTP 请求抽象接口，所有业务子包通过此接口执行 API 调用。
// 根包的 Client 实现此接口，测试环境可注入自定义实现。
type Doer interface {
	Do(ctx context.Context, method, path string, req, resp any, opts ...RequestOption) error
}

// LoadEnv 显式加载 .env 文件。
// 如果不传路径，默认加载当前目录的 .env。
// 返回的 error 可用于判断文件是否存在。
func LoadEnv(paths ...string) error {
	if len(paths) == 0 {
		return godotenv.Load()
	}
	return godotenv.Load(paths...)
}

// ResolveEnvPath 从调用者所在目录向上查找 .env 文件，返回其绝对路径。
// 如果未找到，返回空字符串。
func ResolveEnvPath() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	dir := filepath.Dir(filename)
	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}

// IsUnauthorized 判断错误是否为 Airwallex unauthorized 响应。
func IsUnauthorized(err error) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == "unauthorized"
	}
	return false
}

// IsNotFound 判断错误是否为 404 Not Found 响应。
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeNotFound) || (apiErr.Code == "" && apiErr.Message == "Not Found")
	}
	return false
}

// GenerateRequestID 生成一个 UUIDv4 格式的请求 ID，用于幂等性控制。
func GenerateRequestID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		// 理论上不会发生，但以防万一返回时间戳格式的 fallback
		return fmt.Sprintf("req-%d", time.Now().UnixNano())
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant 10
	return hex.EncodeToString(b[:4]) + "-" +
		hex.EncodeToString(b[4:6]) + "-" +
		hex.EncodeToString(b[6:8]) + "-" +
		hex.EncodeToString(b[8:10]) + "-" +
		hex.EncodeToString(b[10:16])
}
