package risk

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// CardholderDecision represents the cardholder authorization decision type.
// CardholderDecision 持卡人授权决定类型。
type CardholderDecision = string

const (
	// CardholderDecisionAuthorized indicates the cardholder authorized the transaction.
	// CardholderDecisionAuthorized 持卡人已授权。
	CardholderDecisionAuthorized CardholderDecision = "AUTHORIZED_BY_CARDHOLDER"
	// CardholderDecisionNotAuthorized indicates the cardholder did not authorize the transaction.
	// CardholderDecisionNotAuthorized 持卡人未授权。
	CardholderDecisionNotAuthorized CardholderDecision = "NOT_AUTHORIZED_BY_CARDHOLDER"
)

// FraudFeedback represents issuing fraud feedback information.
// FraudFeedback 表示发卡欺诈反馈信息。
type FraudFeedback struct {
	// AuthorizationID is the unique authorization identifier.
	// AuthorizationID 授权唯一标识符。
	AuthorizationID string `json:"authorization_id,omitempty"`
	// CardID is the unique card identifier.
	// CardID 卡片唯一标识符。
	CardID string `json:"card_id,omitempty"`
	// CardholderDecision is the cardholder's authorization decision.
	// CardholderDecision 持卡人授权决定。
	CardholderDecision CardholderDecision `json:"cardholder_decision,omitempty"`
}

// CreateFraudFeedbackRequest represents a request to create fraud feedback.
// CreateFraudFeedbackRequest 创建发卡欺诈反馈请求。
type CreateFraudFeedbackRequest struct {
	// AuthorizationID is the unique authorization identifier. Required.
	// AuthorizationID 授权唯一标识符。必填。
	AuthorizationID string `json:"authorization_id"`
	// CardholderDecision is the cardholder's authorization decision.
	// CardholderDecision 持卡人授权决定。
	CardholderDecision CardholderDecision `json:"cardholder_decision,omitempty"`
}

// CreateFraudFeedback creates issuing fraud feedback.
// CreateFraudFeedback 创建发卡欺诈反馈。
// 官方文档: https://www.airwallex.com/docs/api/risk/fraud_feedback_issuing/fraud_feedback_issuing.md
func (s *Service) CreateFraudFeedback(ctx context.Context, req *CreateFraudFeedbackRequest, opts ...sdk.RequestOption) (*FraudFeedback, error) {
	var resp FraudFeedback
	err := s.doer.Do(ctx, "POST", "/api/v1/risk/issuing/fraud_feedback", req, &resp, opts...)
	return &resp, err
}

// GetFraudFeedback retrieves fraud feedback by authorization ID.
// GetFraudFeedback 根据授权ID获取发卡欺诈反馈。
// 官方文档: https://www.airwallex.com/docs/api/risk/fraud_feedback_issuing/retrieve.md
func (s *Service) GetFraudFeedback(ctx context.Context, authorizationID string, opts ...sdk.RequestOption) (*FraudFeedback, error) {
	var resp FraudFeedback
	err := s.doer.Do(ctx, "GET", "/api/v1/risk/issuing/fraud_feedback/"+authorizationID, nil, &resp, opts...)
	return &resp, err
}

// ListFraudFeedback lists all issuing fraud feedback.
// ListFraudFeedback 列出发卡欺诈反馈。
// 官方文档: https://www.airwallex.com/docs/api/risk/fraud_feedback_issuing/list.md
func (s *Service) ListFraudFeedback(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[FraudFeedback], error) {
	var resp sdk.ListResult[FraudFeedback]
	err := s.doer.Do(ctx, "GET", "/api/v1/risk/issuing/fraud_feedback", nil, &resp, opts...)
	return &resp, err
}
