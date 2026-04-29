package simulation

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// AccountStatus represents the account status enum.
// AccountStatus 账户状态枚举。
type AccountStatus = string

const (
	AccountStatusActionRequired AccountStatus = "ACTION_REQUIRED"
	AccountStatusActive         AccountStatus = "ACTIVE"
	AccountStatusSuspended      AccountStatus = "SUSPENDED"
)

// SimulateAccountUpdateStatusRequest represents a request to update a connected account status.
// SimulateAccountUpdateStatusRequest 更新连接账户状态请求。
type SimulateAccountUpdateStatusRequest struct {
	// NextStatus is the target status. Required.
	// NextStatus 目标状态。必填。
	NextStatus AccountStatus `json:"next_status"`
	// Force indicates whether to force the update. Optional.
	// Force 是否强制更新。可选。
	Force bool `json:"force,omitempty"`
}

// SimulateAccountUpdateStatus updates a connected account's status.
// SimulateAccountUpdateStatus 更新连接账户状态。
// 官方文档: https://www.airwallex.com/docs/api/simulation/accounts/update_status.md
func (s *Service) SimulateAccountUpdateStatus(ctx context.Context, accountID string, req *SimulateAccountUpdateStatusRequest, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/accounts/"+accountID+"/update_status", req, nil, opts...)
}

// SimulateAccountAmendmentApprove approves an account amendment.
// SimulateAccountAmendmentApprove 审批账户修改。
// 官方文档: https://www.airwallex.com/docs/api/simulation/accounts/approve_amendments.md
func (s *Service) SimulateAccountAmendmentApprove(ctx context.Context, amendmentID string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/account/amendments/"+amendmentID+"/approve", nil, nil, opts...)
}

// SimulateAccountAmendmentReject rejects an account amendment.
// SimulateAccountAmendmentReject 拒绝账户修改。
// 官方文档: https://www.airwallex.com/docs/api/simulation/accounts/reject_amendments.md
func (s *Service) SimulateAccountAmendmentReject(ctx context.Context, amendmentID string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/account/amendments/"+amendmentID+"/reject", nil, nil, opts...)
}
