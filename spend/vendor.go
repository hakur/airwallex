package spend

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// VendorStatus represents the vendor status.
// VendorStatus 供应商状态。
type VendorStatus = string

const (
	VendorStatusDraft            VendorStatus = "DRAFT"
	VendorStatusAwaitingApproval VendorStatus = "AWAITING_APPROVAL"
	VendorStatusActive           VendorStatus = "ACTIVE"
	VendorStatusArchived         VendorStatus = "ARCHIVED"
)

// VendorSyncStatus represents the vendor sync status.
// VendorSyncStatus 供应商同步状态。
type VendorSyncStatus = string

const (
	VendorSyncStatusNotSynced  VendorSyncStatus = "NOT_SYNCED"
	VendorSyncStatusSynced     VendorSyncStatus = "SYNCED"
	VendorSyncStatusSyncFailed VendorSyncStatus = "SYNC_FAILED"
)

// Vendor represents a vendor.
// Vendor 表示供应商。
type Vendor struct {
	ID                         string           `json:"id"`
	LegalEntityIDs             []string         `json:"legal_entity_ids"`
	Name                       string           `json:"name"`
	BusinessName               string           `json:"business_name,omitempty"`
	CountryCode                string           `json:"country_code,omitempty"`
	Address                    *Address         `json:"address,omitempty"`
	Status                     VendorStatus     `json:"status"`
	SyncStatus                 VendorSyncStatus `json:"sync_status"`
	SyncErrorMessage           string           `json:"sync_error_message,omitempty"`
	Approvers                  []string         `json:"approvers"`
	OwnerEmail                 string           `json:"owner_email,omitempty"`
	BusinessRegistrationNumber string           `json:"business_registration_number,omitempty"`
	ExternalID                 string           `json:"external_id,omitempty"`
	Contacts                   []VendorContact  `json:"contacts"`
	Attachments                []Attachment     `json:"attachments"`
	Comments                   []Comment        `json:"comments"`
	CreatedAt                  string           `json:"created_at"`
	UpdatedAt                  string           `json:"updated_at"`
}

// VendorContact represents a vendor contact.
// VendorContact 供应商联系人。
type VendorContact struct {
	ContactName string `json:"contact_name,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

// CreateVendorRequest represents a request to create a vendor.
// CreateVendorRequest 创建供应商请求。
type CreateVendorRequest struct {
	RequestID                  string           `json:"request_id"`
	ExternalID                 string           `json:"external_id"`
	Name                       string           `json:"name"`
	LegalEntityIDs             []string         `json:"legal_entity_ids"`
	Status                     VendorStatus     `json:"status"`
	SyncStatus                 VendorSyncStatus `json:"sync_status"`
	OwnerEmail                 string           `json:"owner_email,omitempty"`
	BusinessRegistrationNumber string           `json:"business_registration_number,omitempty"`
	BusinessName               string           `json:"business_name,omitempty"`
	CountryCode                string           `json:"country_code,omitempty"`
	Address                    *Address         `json:"address,omitempty"`
	Contacts                   []VendorContact  `json:"contacts,omitempty"`
}

// SyncVendorRequest represents a request to sync a vendor.
// SyncVendorRequest 同步供应商请求。
type SyncVendorRequest struct {
	SyncStatus       VendorSyncStatus `json:"sync_status"`
	SyncErrorMessage string           `json:"sync_error_message,omitempty"`
}

// ListVendorsRequest represents query parameters for listing vendors.
// ListVendorsRequest 供应商列表查询参数。
type ListVendorsRequest struct {
	Page          string `json:"page,omitempty"`
	FromCreatedAt string `json:"from_created_at,omitempty"`
	ToCreatedAt   string `json:"to_created_at,omitempty"`
	Status        string `json:"status,omitempty"`
	SyncStatus    string `json:"sync_status,omitempty"`
	LegalEntityID string `json:"legal_entity_id,omitempty"`
	ExternalID    string `json:"external_id,omitempty"`
}

// ListVendorsResponse represents a vendor list response with cursor pagination.
// ListVendorsResponse 供应商列表响应（cursor 分页）。
type ListVendorsResponse struct {
	Items      []Vendor `json:"items"`
	PageAfter  string   `json:"page_after,omitempty"`
	PageBefore string   `json:"page_before,omitempty"`
}

// CreateVendor creates a vendor.
// CreateVendor 创建供应商。
// 官方文档: https://www.airwallex.com/docs/api/spend/vendors/create.md
func (s *Service) CreateVendor(ctx context.Context, req *CreateVendorRequest, opts ...sdk.RequestOption) (*Vendor, error) {
	var resp Vendor
	err := s.doer.Do(ctx, "POST", "/api/v1/spend/vendors/create", req, &resp, opts...)
	return &resp, err
}

// GetVendor retrieves vendor details.
// GetVendor 获取供应商详情。
// 官方文档: https://www.airwallex.com/docs/api/spend/vendors/retrieve.md
func (s *Service) GetVendor(ctx context.Context, id string, opts ...sdk.RequestOption) (*Vendor, error) {
	var resp Vendor
	err := s.doer.Do(ctx, "GET", "/api/v1/spend/vendors/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListVendors lists all vendors.
// ListVendors 列出供应商。
// 官方文档: https://www.airwallex.com/docs/api/spend/vendors/list.md
func (s *Service) ListVendors(ctx context.Context, req *ListVendorsRequest, opts ...sdk.RequestOption) (*ListVendorsResponse, error) {
	var resp ListVendorsResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/spend/vendors", req, &resp, opts...)
	return &resp, err
}

// SyncVendor updates the vendor sync status.
// SyncVendor 更新供应商同步状态。
// 官方文档: https://www.airwallex.com/docs/api/spend/vendors/sync.md
func (s *Service) SyncVendor(ctx context.Context, id string, req *SyncVendorRequest, opts ...sdk.RequestOption) (*Vendor, error) {
	var resp Vendor
	err := s.doer.Do(ctx, "POST", "/api/v1/spend/vendors/"+id+"/sync", req, &resp, opts...)
	return &resp, err
}

// --- Deprecated ---
// UpdateVendor 已移除。官方 API 不支持更新供应商。
