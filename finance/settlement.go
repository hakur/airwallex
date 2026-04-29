package finance

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// SettlementStatus represents a settlement status.
// SettlementStatus 结算状态。
type SettlementStatus = string

const (
	// SettlementStatusPending 待处理。
	SettlementStatusPending SettlementStatus = "PENDING"
	// SettlementStatusSettled 已结算。
	SettlementStatusSettled SettlementStatus = "SETTLED"
)

// Settlement represents settlement information.
// Settlement 表示结算信息。
type Settlement struct {
	// Amount is the settlement amount. Optional.
	// Amount 结算金额。可选。
	Amount float64 `json:"amount,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// Currency is the currency code. Optional.
	// Currency 货币代码。可选。
	Currency string `json:"currency,omitempty"`
	// EstimatedSettledAt is the estimated settlement time. Optional.
	// EstimatedSettledAt 预计结算时间。可选。
	EstimatedSettledAt string `json:"estimated_settled_at,omitempty"`
	// Fee is the settlement fee. Optional.
	// Fee 手续费。可选。
	Fee float64 `json:"fee,omitempty"`
	// ID is the settlement batch ID. Optional.
	// ID 结算批次ID。可选。
	ID string `json:"id,omitempty"`
	// SettledAt is the actual settlement time. Optional.
	// SettledAt 实际结算时间。可选。
	SettledAt string `json:"settled_at,omitempty"`
	// Status is the settlement status. Optional.
	Status SettlementStatus `json:"status,omitempty"`
}

// ListSettlementsRequest represents a request to list settlements.
// ListSettlementsRequest 列出结算请求。
type ListSettlementsRequest struct {
	// Currency is the settlement currency code. Required.
	// Currency 结算货币代码。必填。
	Currency string `json:"currency"`
	// FromSettledAt is the start of settlement time (ISO8601 inclusive). Required.
	// FromSettledAt 结算时间起始（ISO8601 inclusive）。必填。
	FromSettledAt string `json:"from_settled_at"`
	// PageNum is the page number, starting from 0. Optional.
	// PageNum 页码，从0开始。可选。
	PageNum int32 `json:"page_num,omitempty"`
	// PageSize is the number of items per page. Optional.
	// PageSize 每页数量。可选。
	PageSize int32 `json:"page_size,omitempty"`
	// Status is the settlement status. Required.
	Status SettlementStatus `json:"status"`
	// ToSettledAt is the end of settlement time (ISO8601 inclusive). Required.
	// ToSettledAt 结算时间截止（ISO8601 inclusive）。必填。
	ToSettledAt string `json:"to_settled_at"`
}

// GetSettlementReportRequest represents a request to get a settlement report.
// GetSettlementReportRequest 获取结算报告请求。
type GetSettlementReportRequest struct {
	// FileFormat is the report file format. Optional.
	// FileFormat 报告文件格式。可选。
	FileFormat string `json:"file_format,omitempty"`
	// Version is the report version. Optional.
	// Version 报告版本。可选。
	Version string `json:"version,omitempty"`
}

// SettlementReport represents a settlement report.
// SettlementReport 结算报告。
type SettlementReport struct {
	// CreatedAt is the report creation time. Optional.
	// CreatedAt 报告创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// ID is the settlement batch ID. Optional.
	// ID 结算批次ID。可选。
	ID string `json:"id,omitempty"`
	// ReportURL is the report download URL. Optional.
	// ReportURL 报告下载地址。可选。
	ReportURL string `json:"report_url,omitempty"`
}

// GetSettlement retrieves a settlement by ID.
// GetSettlement 根据 ID 获取结算。
// 官方文档: https://www.airwallex.com/docs/api/finance/settlements/retrieve.md
func (s *Service) GetSettlement(ctx context.Context, id string, opts ...sdk.RequestOption) (*Settlement, error) {
	var resp Settlement
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/financial/settlements/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListSettlements lists settlements with the specified filters.
// ListSettlements 列出结算。
// 官方文档: https://www.airwallex.com/docs/api/finance/settlements/list.md
func (s *Service) ListSettlements(ctx context.Context, req *ListSettlementsRequest, opts ...sdk.RequestOption) (*sdk.ListResult[Settlement], error) {
	var resp sdk.ListResult[Settlement]
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/financial/settlements", req, &resp, opts...)
	return &resp, err
}

// GetSettlementReport retrieves a settlement report.
// GetSettlementReport 获取结算报告。
// 官方文档: https://www.airwallex.com/docs/api/finance/settlements/report.md
func (s *Service) GetSettlementReport(ctx context.Context, id string, req *GetSettlementReportRequest, opts ...sdk.RequestOption) (*SettlementReport, error) {
	var resp SettlementReport
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/financial/settlements/"+id+"/report", req, &resp, opts...)
	return &resp, err
}
