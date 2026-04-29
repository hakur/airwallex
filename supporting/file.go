package supporting

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// UploadFileResponse represents the response for uploading a file.
// UploadFileResponse 文件上传响应。
type UploadFileResponse struct {
	// FileID is the unique identifier of the uploaded file.
	// FileID 文件唯一标识符。
	FileID string `json:"file_id,omitempty"`
	// Filename is the name of the uploaded file.
	// Filename 文件名。
	Filename string `json:"filename,omitempty"`
	// Size is the file size in bytes.
	// Size 文件大小（字节）。
	Size int64 `json:"size,omitempty"`
	// Notes is the optional notes for the file.
	// Notes 文件备注。
	Notes string `json:"notes,omitempty"`
	// ObjectType is the type of the associated object.
	// ObjectType 关联对象类型。
	ObjectType string `json:"object_type,omitempty"`
	// Created is the creation timestamp in milliseconds.
	// Created 创建时间戳（毫秒）。
	Created int64 `json:"created,omitempty"`
}

// DownloadLinkRequest represents a request for file download links.
// DownloadLinkRequest 下载链接请求。
type DownloadLinkRequest struct {
	// FileIDs is the list of file IDs to generate download links for.
	// FileIDs 文件 ID 列表。
	FileIDs []string `json:"file_ids,omitempty"`
}

// FileDownloadLink represents a file download link.
// FileDownloadLink 文件下载链接。
type FileDownloadLink struct {
	// FileID is the unique identifier of the file.
	// FileID 文件唯一标识符。
	FileID string `json:"file_id,omitempty"`
	// Filename is the name of the file.
	// Filename 文件名。
	Filename string `json:"filename,omitempty"`
	// URL is the download URL for the file.
	// URL 文件下载 URL。
	URL string `json:"url,omitempty"`
	// ContentType is the MIME type of the file.
	// ContentType 文件 MIME 类型。
	ContentType string `json:"content_type,omitempty"`
	// Size is the file size in bytes.
	// Size 文件大小（字节）。
	Size int64 `json:"size,omitempty"`
	// DownloadLinkValidUntil is the expiration time of the download link.
	// DownloadLinkValidUntil 下载链接过期时间。
	DownloadLinkValidUntil string `json:"download_link_valid_until,omitempty"`
}

// DownloadLinksResponse represents the response for download link requests.
// DownloadLinksResponse 下载链接响应。
type DownloadLinksResponse struct {
	// AbsentFiles is the list of file IDs that were not found.
	// AbsentFiles 未找到的文件 ID 列表。
	AbsentFiles []string `json:"absent_files,omitempty"`
	// Files is the list of file download links.
	// Files 文件下载链接列表。
	Files []FileDownloadLink `json:"files,omitempty"`
}

// UploadFile uploads a file.
// UploadFile 上传文件。
// 官方文档: https://www.airwallex.com/docs/api/supporting_services/file_service/upload_files.md
func (s *Service) UploadFile(ctx context.Context, req *UploadFileRequest, opts ...sdk.RequestOption) (*UploadFileResponse, error) {
	var resp UploadFileResponse
	err := s.doer.Do(ctx, "POST", "/api/v1/files/upload", req, &resp, opts...)
	return &resp, err
}

// UploadFileRequest represents an upload file request (multipart form, using the file field).
// UploadFileRequest 上传文件请求（multipart form，使用 file 字段上传）。
type UploadFileRequest struct {
	// File is the multipart form file data (caller constructs the form body).
	// File multipart form 文件数据（由调用方构造 form body）。
	// Notes is the optional notes for the file.
	// Notes 文件备注。可选。
	Notes string `json:"notes,omitempty"`
}

// GetDownloadLinks gets download links for files.
// GetDownloadLinks 获取文件下载链接。
// 官方文档: https://www.airwallex.com/docs/api/supporting_services/file_service/download_links_files.md
func (s *Service) GetDownloadLinks(ctx context.Context, req *DownloadLinkRequest, opts ...sdk.RequestOption) (*DownloadLinksResponse, error) {
	var resp DownloadLinksResponse
	err := s.doer.Do(ctx, "POST", "/api/v1/files/download_links", req, &resp, opts...)
	return &resp, err
}

// --- Deprecated ---
// CreateFile / GetFile / UpdateFile / ListFiles 已移除。官方 File Service API 仅支持 Upload 和 DownloadLinks。
