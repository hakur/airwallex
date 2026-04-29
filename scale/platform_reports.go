package scale

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PlatformReport represents a platform report.
// PlatformReport 表示平台报告。
type PlatformReport struct {
	ID                   string `json:"id,omitempty"`
	Status               string `json:"status,omitempty"`
	Type                 string `json:"type,omitempty"`
	FileFormat           string `json:"file_format,omitempty"`
	FileName             string `json:"file_name,omitempty"`
	DownloadURL          string `json:"download_url,omitempty"`
	DownloadURLExpiresAt string `json:"download_url_expires_at,omitempty"`
	FailedReason         string `json:"failed_reason,omitempty"`
	FromCreatedAt        string `json:"from_created_at,omitempty"`
	ToCreatedAt          string `json:"to_created_at,omitempty"`
	FromUpdatedAt        string `json:"from_updated_at,omitempty"`
	ToUpdatedAt          string `json:"to_updated_at,omitempty"`
}

// CreatePlatformReportRequest represents a request to create a platform report.
// CreatePlatformReportRequest 创建平台报告请求。
type CreatePlatformReportRequest struct {
	FileFormat    string `json:"file_format"`
	Type          string `json:"type"`
	FileName      string `json:"file_name,omitempty"`
	FromCreatedAt string `json:"from_created_at,omitempty"`
	ToCreatedAt   string `json:"to_created_at,omitempty"`
	FromUpdatedAt string `json:"from_updated_at,omitempty"`
	ToUpdatedAt   string `json:"to_updated_at,omitempty"`
}

// CreatePlatformReportResponse represents the response for creating a platform report (returns ID only).
// CreatePlatformReportResponse 创建平台报告响应（仅返回 ID）。
type CreatePlatformReportResponse struct {
	ID string `json:"id"`
}

// CreatePlatformReport creates a new platform report.
// CreatePlatformReport 创建平台报告。
// 官方文档: https://www.airwallex.com/docs/api/scale/platform_reports/create.md
func (s *Service) CreatePlatformReport(ctx context.Context, req *CreatePlatformReportRequest, opts ...sdk.RequestOption) (*CreatePlatformReportResponse, error) {
	var resp CreatePlatformReportResponse
	err := s.doer.Do(ctx, "POST", "/api/v1/platform_reports/create", req, &resp, opts...)
	return &resp, err
}

// GetPlatformReport retrieves a platform report details and download URL.
// GetPlatformReport 获取平台报告详情及下载链接。
// 官方文档: https://www.airwallex.com/docs/api/scale/platform_reports/retrieve.md
func (s *Service) GetPlatformReport(ctx context.Context, id string, opts ...sdk.RequestOption) (*PlatformReport, error) {
	var resp PlatformReport
	err := s.doer.Do(ctx, "GET", "/api/v1/platform_reports/"+id, nil, &resp, opts...)
	return &resp, err
}

// --- Deprecated ---
// ListPlatformReports 已移除。官方 API 不支持列出平台报告。
