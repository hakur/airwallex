package simulation

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// SimulateBillingFailNextAutocharge simulates a payment source's next autocharge failure.
// SimulateBillingFailNextAutocharge 模拟账单支付来源的下次自动扣款失败。
// 官方文档: https://www.airwallex.com/docs/api/simulation/billing/fail_next_autocharge_payment_sources.md
func (s *Service) SimulateBillingFailNextAutocharge(ctx context.Context, paymentSourceID string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/billing/payment_sources/"+paymentSourceID+"/fail_next_autocharge", nil, nil, opts...)
}
