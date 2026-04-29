package finance

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// FinancialReportStatus represents a financial report status.
// FinancialReportStatus 财务报告状态。
type FinancialReportStatus = string

const (
	// FinancialReportStatusPending 待处理。
	FinancialReportStatusPending FinancialReportStatus = "PENDING"
	// FinancialReportStatusCompleted 已完成。
	FinancialReportStatusCompleted FinancialReportStatus = "COMPLETED"
	// FinancialReportStatusFailed 失败。
	FinancialReportStatusFailed FinancialReportStatus = "FAILED"
)

// FinancialReportType represents a financial report type.
// FinancialReportType 财务报告类型。
type FinancialReportType = string

const (
	// FinancialReportTypeAccountStatement 账户对账单。
	FinancialReportTypeAccountStatement FinancialReportType = "ACCOUNT_STATEMENT_REPORT"
	// FinancialReportTypeBalanceActivity 余额活动报告。
	FinancialReportTypeBalanceActivity FinancialReportType = "BALANCE_ACTIVITY_REPORT"
	// FinancialReportTypeTransactionRecon 交易对账报告。
	FinancialReportTypeTransactionRecon FinancialReportType = "TRANSACTION_RECON_REPORT"
	// FinancialReportTypeOnlinePayments 在线支付交易报告。
	FinancialReportTypeOnlinePayments FinancialReportType = "ONLINE_PAYMENTS_TRANSACTION_REPORT"
	// FinancialReportTypeSettlement 结算报告。
	FinancialReportTypeSettlement FinancialReportType = "SETTLEMENT_REPORT"
	// FinancialReportTypeAggregatedSettlement 聚合结算报告。
	FinancialReportTypeAggregatedSettlement FinancialReportType = "AGGREGATED_SETTLEMENT_REPORT"
)

// FinancialReportFileFormat represents a financial report file format.
// FinancialReportFileFormat 财务报告文件格式。
type FinancialReportFileFormat = string

const (
	// FinancialReportFileFormatCSV CSV格式。
	FinancialReportFileFormatCSV FinancialReportFileFormat = "CSV"
	// FinancialReportFileFormatExcel Excel格式。
	FinancialReportFileFormatExcel FinancialReportFileFormat = "EXCEL"
	// FinancialReportFileFormatPDF PDF格式。
	FinancialReportFileFormatPDF FinancialReportFileFormat = "PDF"
)

// FinancialReportParameters represents financial report parameters.
// FinancialReportParameters 财务报告参数。
type FinancialReportParameters struct {
	// Currencies is the list of currencies to include. Optional.
	// Currencies 包含的货币列表。可选。
	Currencies []string `json:"currencies,omitempty"`
	// FromDate is the report start date. Optional.
	// FromDate 报告开始日期。可选。
	FromDate string `json:"from_date,omitempty"`
	// TimeZone is the time zone. Optional.
	// TimeZone 时区。可选。
	TimeZone string `json:"time_zone,omitempty"`
	// ToDate is the report end date. Optional.
	// ToDate 报告结束日期。可选。
	ToDate string `json:"to_date,omitempty"`
	// TransactionTypes is the list of transaction types to include. Optional.
	TransactionTypes []string `json:"transaction_types,omitempty"`
}

// FinancialReport represents financial report information.
// FinancialReport 表示财务报告信息。
type FinancialReport struct {
	// FileFormat 文件格式。可选。
	FileFormat FinancialReportFileFormat `json:"file_format,omitempty"`
	// FileName 文件名。可选。
	FileName string `json:"file_name,omitempty"`
	// ID 报告唯一标识符。可选。
	ID string `json:"id,omitempty"`
	// ReportExpiresAt 报告过期时间。可选。
	ReportExpiresAt string `json:"report_expires_at,omitempty"`
	// ReportParameters 报告参数。可选。
	ReportParameters *FinancialReportParameters `json:"report_parameters,omitempty"`
	// ReportVersion 报告版本。可选。
	ReportVersion string `json:"report_version,omitempty"`
	// Status 报告状态。可选。
	Status FinancialReportStatus `json:"status,omitempty"`
	// Type 报告类型。可选。
	Type FinancialReportType `json:"type,omitempty"`
}

// CreateFinancialReportRequest represents a request to create a financial report.
// CreateFinancialReportRequest 创建财务报告请求。
type CreateFinancialReportRequest struct {
	// Currencies 包含的货币列表。可选。
	Currencies []string `json:"currencies,omitempty"`
	// FileFormat 文件格式。必填。
	FileFormat FinancialReportFileFormat `json:"file_format"`
	// FileName 文件名。可选。
	FileName string `json:"file_name,omitempty"`
	// FromDate 报告开始日期（inclusive）。必填。
	FromDate string `json:"from_date"`
	// ReportOptions 报告选项。可选。
	ReportOptions *FinancialReportOptions `json:"report_options,omitempty"`
	// ReportVersion 报告版本。可选。
	ReportVersion string `json:"report_version,omitempty"`
	// SettlementCurrencies 结算货币列表。可选。
	SettlementCurrencies []string `json:"settlement_currencies,omitempty"`
	// Statuses 交易状态列表。可选。
	Statuses []string `json:"statuses,omitempty"`
	// TimeZone 时区。可选。
	TimeZone string `json:"time_zone,omitempty"`
	// ToDate 报告结束日期（inclusive）。必填。
	ToDate string `json:"to_date"`
	// TransactionCurrencies 交易货币列表。可选。
	TransactionCurrencies []string `json:"transaction_currencies,omitempty"`
	// TransactionTypes 交易类型列表。可选。
	TransactionTypes []string `json:"transaction_types,omitempty"`
	// Type 报告类型。必填。
	Type FinancialReportType `json:"type"`
}

// FinancialReportOptions represents financial report options.
// FinancialReportOptions 财务报告选项。
type FinancialReportOptions struct {
	// IncludeReservations 是否包含预留交易。可选。
	IncludeReservations *bool `json:"include_reservations,omitempty"`
}

// GetFinancialReport retrieves a financial report by ID.
// GetFinancialReport 根据 ID 获取财务报告。
// 官方文档: https://www.airwallex.com/docs/api/finance/financial_reports/retrieve.md
func (s *Service) GetFinancialReport(ctx context.Context, id string, opts ...sdk.RequestOption) (*FinancialReport, error) {
	var resp FinancialReport
	err := s.doer.Do(ctx, "GET", "/api/v1/finance/financial_reports/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListFinancialReports lists financial reports.
// ListFinancialReports 列出财务报告。
// 官方文档: https://www.airwallex.com/docs/api/finance/financial_reports/list.md
func (s *Service) ListFinancialReports(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[FinancialReport], error) {
	var resp sdk.ListResult[FinancialReport]
	err := s.doer.Do(ctx, "GET", "/api/v1/finance/financial_reports", nil, &resp, opts...)
	return &resp, err
}

// CreateFinancialReport creates a new financial report.
// CreateFinancialReport 创建财务报告。
// 官方文档: https://www.airwallex.com/docs/api/finance/financial_reports/create.md
func (s *Service) CreateFinancialReport(ctx context.Context, req *CreateFinancialReportRequest, opts ...sdk.RequestOption) (*FinancialReport, error) {
	var resp FinancialReport
	err := s.doer.Do(ctx, "POST", "/api/v1/finance/financial_reports/create", req, &resp, opts...)
	return &resp, err
}

// GetFinancialReportContent retrieves the content of a financial report.
// GetFinancialReportContent 获取财务报告内容。
// 官方文档: https://www.airwallex.com/docs/api/finance/financial_reports/content.md
func (s *Service) GetFinancialReportContent(ctx context.Context, id string, opts ...sdk.RequestOption) ([]byte, error) {
	var resp []byte
	err := s.doer.Do(ctx, "GET", "/api/v1/finance/financial_reports/"+id+"/content", nil, &resp, opts...)
	return resp, err
}
