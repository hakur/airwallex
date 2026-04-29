package simulation

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// RFIType represents the RFI type enum.
// RFIType RFI 类型枚举。
type RFIType = string

const (
	RFITypeKYC               RFIType = "KYC"
	RFITypeKYCOngoing        RFIType = "KYC_ONGOING"
	RFITypeCardholder        RFIType = "CARDHOLDER"
	RFITypeTransaction       RFIType = "TRANSACTION"
	RFITypePaymentEnablement RFIType = "PAYMENT_ENABLEMENT"
	RFITypeMerchantRisk      RFIType = "MERCHANT_RISK"
)

// RFIQuestionAnswerType represents the question answer type enum.
// RFIQuestionAnswerType 问题回答类型枚举。
type RFIQuestionAnswerType = string

const (
	RFIQuestionAnswerTypeAddress          RFIQuestionAnswerType = "ADDRESS"
	RFIQuestionAnswerTypeAttachment       RFIQuestionAnswerType = "ATTACHMENT"
	RFIQuestionAnswerTypeConfirmation     RFIQuestionAnswerType = "CONFIRMATION"
	RFIQuestionAnswerTypeIdentityDocument RFIQuestionAnswerType = "IDENTITY_DOCUMENT"
	RFIQuestionAnswerTypeLiveness         RFIQuestionAnswerType = "LIVENESS"
	RFIQuestionAnswerTypeText             RFIQuestionAnswerType = "TEXT"
)

// RFIQuestion represents an RFI question.
// RFIQuestion 表示 RFI 问题。
type RFIQuestion struct {
	// Answer is the answer type. Required.
	// Answer 回答类型。必填。
	Answer struct {
		// Type is the answer type. Required.
		// Type 回答类型。必填。
		Type RFIQuestionAnswerType `json:"type"`
	} `json:"answer"`
}

// SimulateRFICreateRequest represents a request to create an RFI.
// SimulateRFICreateRequest 创建 RFI 请求。
type SimulateRFICreateRequest struct {
	// Type is the RFI type. Required.
	// Type RFI 类型。必填。
	Type RFIType `json:"type"`
	// Questions is the list of questions. Optional.
	// Questions 问题列表。可选。
	Questions []RFIQuestion `json:"questions,omitempty"`
}

// SimulateRFIFollowUpRequest represents a request to follow up on an RFI.
// SimulateRFIFollowUpRequest 跟进 RFI 请求。
type SimulateRFIFollowUpRequest struct {
	// 空请求体
}

// SimulateRFICreate creates an RFI simulation.
// SimulateRFICreate 创建 RFI。
// 官方文档: https://www.airwallex.com/docs/api/simulation/rfis/create.md
func (s *Service) SimulateRFICreate(ctx context.Context, req *SimulateRFICreateRequest, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/rfis/create", req, nil, opts...)
}

// SimulateRFIClose closes an RFI simulation.
// SimulateRFIClose 关闭 RFI。
// 官方文档: https://www.airwallex.com/docs/api/simulation/rfis/close.md
func (s *Service) SimulateRFIClose(ctx context.Context, rfiID string, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/rfis/"+rfiID+"/close", nil, nil, opts...)
}

// SimulateRFIFollowUp follows up on an RFI simulation.
// SimulateRFIFollowUp 跟进 RFI。
// 官方文档: https://www.airwallex.com/docs/api/simulation/rfis/follow_up.md
func (s *Service) SimulateRFIFollowUp(ctx context.Context, rfiID string, req *SimulateRFIFollowUpRequest, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/rfis/"+rfiID+"/follow_up", req, nil, opts...)
}
