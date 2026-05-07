// Package events provides typed webhook event structures for the platform domain.
// Platform 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/platform.md
//
// 事件映射表:
//
//	platform_report.completed           → PlatformReportCompletedEvent     (Data: PlatformReport)
//	platform_report.failed              → PlatformReportFailedEvent        (Data: PlatformReport)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Platform Report Events ---

// PlatformReportCompletedEvent represents the platform_report.completed webhook event.
// PlatformReportCompletedEvent 表示平台报告已完成事件。
type PlatformReportCompletedEvent struct {
	Event
	Data PlatformReport `json:"data"`
}

// PlatformReportFailedEvent represents the platform_report.failed webhook event.
// PlatformReportFailedEvent 表示平台报告生成失败事件。
type PlatformReportFailedEvent struct {
	Event
	Data PlatformReport `json:"data"`
}

// PlatformReport represents a platform report in webhook payloads.
// PlatformReport 表示 webhook payload 中的平台报告对象。
//
// 字段精确匹配官方文档 payload 示例。
type PlatformReport struct {
	// ID is the unique identifier of the platform report.
	ID string `json:"id,omitempty"`
	// Status is the current status of the report (e.g., COMPLETED, FAILED).
	Status string `json:"status,omitempty"`
	// Type is the report type (e.g., TRANSACTION_RECON_REPORT).
	Type string `json:"type,omitempty"`
	// FileFormat is the format of the generated report file (e.g., CSV).
	FileFormat string `json:"file_format,omitempty"`
	// FileName is the name of the report file.
	FileName string `json:"file_name,omitempty"`
	// DownloadURL is the signed URL to download the completed report.
	// Only present when status is COMPLETED.
	DownloadURL string `json:"download_url,omitempty"`
	// DownloadURLExpiresAt is the expiration time of the download URL.
	// Only present when status is COMPLETED.
	DownloadURLExpiresAt string `json:"download_url_expires_at,omitempty"`
	// FailedReason is the reason why report generation failed.
	// Only present when status is FAILED.
	FailedReason string `json:"failed_reason,omitempty"`
	// FromCreatedAt is the start of the date range for the report.
	FromCreatedAt string `json:"from_created_at,omitempty"`
	// ToCreatedAt is the end of the date range for the report.
	ToCreatedAt string `json:"to_created_at,omitempty"`
}
