package simulation

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PayoutPaymentStatus represents the payout payment status enum.
// PayoutPaymentStatus 支付状态枚举。
type PayoutPaymentStatus = string

const (
	PayoutPaymentStatusPending    PayoutPaymentStatus = "PENDING"
	PayoutPaymentStatusProcessing PayoutPaymentStatus = "PROCESSING"
	PayoutPaymentStatusSent       PayoutPaymentStatus = "SENT"
	PayoutPaymentStatusPaid       PayoutPaymentStatus = "PAID"
	PayoutPaymentStatusFailed     PayoutPaymentStatus = "FAILED"
	PayoutPaymentStatusCancelled  PayoutPaymentStatus = "CANCELLED"
)

// SimulatePayoutPaymentTransitionRequest represents a request to transition a payment status.
// SimulatePayoutPaymentTransitionRequest 支付状态转换请求。
type SimulatePayoutPaymentTransitionRequest struct {
	// NextStatus is the target status. Optional.
	// NextStatus 目标状态。可选。
	NextStatus PayoutPaymentStatus `json:"next_status,omitempty"`
	// FailureType is the failure type. Optional.
	// FailureType 失败类型。可选。
	FailureType string `json:"failure_type,omitempty"`
}

// SimulatePayoutPaymentTransition simulates a payment status transition.
// SimulatePayoutPaymentTransition 模拟支付状态转换。
// 官方文档: https://www.airwallex.com/docs/api/simulation/payments/transition.md
func (s *Service) SimulatePayoutPaymentTransition(ctx context.Context, paymentID string, req *SimulatePayoutPaymentTransitionRequest, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/payments/"+paymentID+"/transition", req, nil, opts...)
}
