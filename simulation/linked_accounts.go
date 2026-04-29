package simulation

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// SimulateLinkedAccountFailMicrodeposits simulates a linked account micro-deposit verification failure.
// SimulateLinkedAccountFailMicrodeposits 模拟关联账户微存款验证失败。
// 官方文档: https://www.airwallex.com/docs/api/simulation/linked_accounts/fail_microdeposits.md
func (s *Service) SimulateLinkedAccountFailMicrodeposits(ctx context.Context, linkedAccountID string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/linked_accounts/"+linkedAccountID+"/fail_microdeposits", nil, nil, opts...)
}

// SimulateLinkedAccountMandateAccept accepts a linked account mandate.
// SimulateLinkedAccountMandateAccept 接受授权。
// 官方文档: https://www.airwallex.com/docs/api/simulation/linked_accounts/accept_mandate.md
func (s *Service) SimulateLinkedAccountMandateAccept(ctx context.Context, linkedAccountID string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/linked_accounts/"+linkedAccountID+"/mandate/accept", nil, nil, opts...)
}

// SimulateLinkedAccountMandateCancel cancels a linked account mandate.
// SimulateLinkedAccountMandateCancel 取消授权。
// 官方文档: https://www.airwallex.com/docs/api/simulation/linked_accounts/cancel_mandate.md
func (s *Service) SimulateLinkedAccountMandateCancel(ctx context.Context, linkedAccountID string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/linked_accounts/"+linkedAccountID+"/mandate/cancel", nil, nil, opts...)
}

// SimulateLinkedAccountMandateReject rejects a linked account mandate.
// SimulateLinkedAccountMandateReject 拒绝授权。
// 官方文档: https://www.airwallex.com/docs/api/simulation/linked_accounts/reject_mandate.md
func (s *Service) SimulateLinkedAccountMandateReject(ctx context.Context, linkedAccountID string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/linked_accounts/"+linkedAccountID+"/mandate/reject", nil, nil, opts...)
}
