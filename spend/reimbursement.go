package spend

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// ReimbursementReportStatus represents the reimbursement report status.
// ReimbursementReportStatus 报销报告状态。
type ReimbursementReportStatus = string

const (
	ReimbursementReportStatusDraft             ReimbursementReportStatus = "DRAFT"
	ReimbursementReportStatusAwaitingApproval  ReimbursementReportStatus = "AWAITING_APPROVAL"
	ReimbursementReportStatusAwaitingPayment   ReimbursementReportStatus = "AWAITING_PAYMENT"
	ReimbursementReportStatusPaymentInProgress ReimbursementReportStatus = "PAYMENT_IN_PROGRESS"
	ReimbursementReportStatusPaid              ReimbursementReportStatus = "PAID"
	ReimbursementReportStatusRejected          ReimbursementReportStatus = "REJECTED"
	ReimbursementReportStatusMarkedAsPaid      ReimbursementReportStatus = "MARKED_AS_PAID"
	ReimbursementReportStatusDeleted           ReimbursementReportStatus = "DELETED"
)

// ReimbursementReportSyncStatus represents the reimbursement report sync status with the accounting system.
// ReimbursementReportSyncStatus 报销报告会计系统同步状态。
type ReimbursementReportSyncStatus = string

const (
	ReimbursementReportSyncStatusNotSynced  ReimbursementReportSyncStatus = "NOT_SYNCED"
	ReimbursementReportSyncStatusSynced     ReimbursementReportSyncStatus = "SYNCED"
	ReimbursementReportSyncStatusSyncFailed ReimbursementReportSyncStatus = "SYNC_FAILED"
)

// ReimbursementReport represents a reimbursement report.
// ReimbursementReport 表示报销报告。
type ReimbursementReport struct {
	ID                             string                        `json:"id"`
	LegalEntityID                  string                        `json:"legal_entity_id"`
	Name                           string                        `json:"name,omitempty"`
	Status                         ReimbursementReportStatus     `json:"status"`
	SyncStatus                     ReimbursementReportSyncStatus `json:"sync_status"`
	BillingCurrency                string                        `json:"billing_currency,omitempty"`
	Approvers                      []string                      `json:"approvers"`
	CreatedBy                      string                        `json:"created_by,omitempty"`
	CreatedAt                      string                        `json:"created_at"`
	UpdatedAt                      string                        `json:"updated_at"`
	BeneficiaryID                  string                        `json:"beneficiary_id,omitempty"`
	Comments                       []Comment                     `json:"comments"`
	ReimbursementReportTransferIDs []string                      `json:"reimbursement_report_transfer_ids"`
	AccountingFieldSelections      []AccountingFieldSelection    `json:"accounting_field_selections"`
}

// ListReimbursementReportsRequest represents query parameters for listing reimbursement reports.
// ListReimbursementReportsRequest 报销报告列表查询参数。
type ListReimbursementReportsRequest struct {
	Page          string   `json:"page,omitempty"`
	FromUpdatedAt string   `json:"from_updated_at,omitempty"`
	ToUpdatedAt   string   `json:"to_updated_at,omitempty"`
	Status        []string `json:"status,omitempty"`
	SyncStatus    []string `json:"sync_status,omitempty"`
	LegalEntityID string   `json:"legal_entity_id,omitempty"`
	Approver      string   `json:"approver,omitempty"`
}

// ListReimbursementReportsResponse represents a reimbursement report list response with cursor pagination.
// ListReimbursementReportsResponse 报销报告列表响应（cursor 分页）。
type ListReimbursementReportsResponse struct {
	Items      []ReimbursementReport `json:"items"`
	PageAfter  string                `json:"page_after,omitempty"`
	PageBefore string                `json:"page_before,omitempty"`
}

// ListReimbursementReports lists all reimbursement reports.
// ListReimbursementReports 列出报销报告。
// 官方文档: https://www.airwallex.com/docs/api/spend/reimbursements/reimbursement_reports.md
func (s *Service) ListReimbursementReports(ctx context.Context, req *ListReimbursementReportsRequest, opts ...sdk.RequestOption) (*ListReimbursementReportsResponse, error) {
	var resp ListReimbursementReportsResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/spend/reimbursement_reports", req, &resp, opts...)
	return &resp, err
}

// GetReimbursementReport retrieves reimbursement report details.
// GetReimbursementReport 获取报销报告详情。
// 官方文档: https://www.airwallex.com/docs/api/spend/reimbursements/retrieve_reimbursement_reports.md
func (s *Service) GetReimbursementReport(ctx context.Context, id string, opts ...sdk.RequestOption) (*ReimbursementReport, error) {
	var resp ReimbursementReport
	err := s.doer.Do(ctx, "GET", "/api/v1/spend/reimbursement_reports/"+id, nil, &resp, opts...)
	return &resp, err
}

// MarkReimbursementReportAsPaid marks a reimbursement report as paid outside Airwallex.
// MarkReimbursementReportAsPaid 标记报销报告为 Airwallex 外已付款。
// 官方文档: https://www.airwallex.com/docs/api/spend/reimbursements/mark_as_paid_reimbursement_reports.md
func (s *Service) MarkReimbursementReportAsPaid(ctx context.Context, id string, opts ...sdk.RequestOption) (*ReimbursementReport, error) {
	var resp ReimbursementReport
	err := s.doer.Do(ctx, "POST", "/api/v1/spend/reimbursement_reports/"+id+"/mark_as_paid", nil, &resp, opts...)
	return &resp, err
}

// --- Deprecated 说明 ---
// 旧版 CreateReimbursement/GetReimbursement/UpdateReimbursement/ListReimbursements 已移除。
// 对应端点 POST/GET /api/v1/spend/reimbursements/* 在官方 API 文档中不存在。
// 请使用 ReimbursementReport 相关函数。
