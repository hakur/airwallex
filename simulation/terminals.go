package simulation

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// TerminalPaymentScenario represents a terminal payment scenario.
// TerminalPaymentScenario 表示终端支付场景。
type TerminalPaymentScenario struct {
	// ID 场景唯一标识符。可选。
	ID string `json:"id,omitempty"`
	// Name 场景名称。可选。
	Name string `json:"name,omitempty"`
	// Description 场景描述。可选。
	Description string `json:"description,omitempty"`
}

// TerminalPaymentScenariosResponse represents a response containing payment scenarios.
// TerminalPaymentScenariosResponse 表示支付场景列表响应。
type TerminalPaymentScenariosResponse struct {
	// Items 场景列表。可选。
	Items []TerminalPaymentScenario `json:"items,omitempty"`
}

// SimulateTerminalConfirmPaymentIntentRequest represents a request to confirm a terminal payment intent.
// SimulateTerminalConfirmPaymentIntentRequest 终端确认支付请求。
type SimulateTerminalConfirmPaymentIntentRequest struct {
	// PaymentIntentID 支付意图ID。可选。
	PaymentIntentID string `json:"payment_intent_id,omitempty"`
	// ScenarioID 场景ID。可选。
	ScenarioID string `json:"scenario_id,omitempty"`
}

// SimulateTerminalGenerateActivationCodeRequest represents a request to generate an activation code.
// SimulateTerminalGenerateActivationCodeRequest 生成激活码请求。
type SimulateTerminalGenerateActivationCodeRequest struct {
	// 空请求体
}

// SimulateTerminalTurnOffRequest represents a request to turn off a terminal.
// SimulateTerminalTurnOffRequest 关闭终端请求。
type SimulateTerminalTurnOffRequest struct {
	// 空请求体
}

// SimulateTerminalTurnOnRequest represents a request to turn on a terminal.
// SimulateTerminalTurnOnRequest 打开终端请求。
type SimulateTerminalTurnOnRequest struct {
	// 空请求体
}

// SimulateTerminalConfirmPaymentIntent confirms a terminal payment intent.
// SimulateTerminalConfirmPaymentIntent 终端确认支付。
// 官方文档: https://www.airwallex.com/docs/api/simulation/terminals/confirm_payment_intent.md
func (s *Service) SimulateTerminalConfirmPaymentIntent(ctx context.Context, req *SimulateTerminalConfirmPaymentIntentRequest, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/pa/pos/terminals/confirm_payment_intent", req, nil, opts...)
}

// SimulateTerminalGenerateActivationCode generates a terminal activation code.
// SimulateTerminalGenerateActivationCode 生成终端激活码。
// 官方文档: https://www.airwallex.com/docs/api/simulation/terminals/generate_activation_code.md
func (s *Service) SimulateTerminalGenerateActivationCode(ctx context.Context, req *SimulateTerminalGenerateActivationCodeRequest, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/pa/pos/terminals/generate_activation_code", req, nil, opts...)
}

// SimulateTerminalPaymentScenarios retrieves terminal payment scenario list.
// SimulateTerminalPaymentScenarios 获取终端支付场景列表。
// 官方文档: https://www.airwallex.com/docs/api/simulation/terminals/payment_scenarios.md
func (s *Service) SimulateTerminalPaymentScenarios(ctx context.Context, opts ...sdk.RequestOption) (*TerminalPaymentScenariosResponse, error) {
	var resp TerminalPaymentScenariosResponse
	err := s.doer.Do(ctx, "GET", "/api/v1/simulation/pa/pos/terminals/payment_scenarios", nil, &resp, opts...)
	return &resp, err
}

// SimulateTerminalTurnOff turns off a terminal.
// SimulateTerminalTurnOff 关闭终端。
// 官方文档: https://www.airwallex.com/docs/api/simulation/terminals/turn_off.md
func (s *Service) SimulateTerminalTurnOff(ctx context.Context, req *SimulateTerminalTurnOffRequest, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/pa/pos/terminals/turn_off", req, nil, opts...)
}

// SimulateTerminalTurnOn turns on a terminal.
// SimulateTerminalTurnOn 打开终端。
// 官方文档: https://www.airwallex.com/docs/api/simulation/terminals/turn_on.md
func (s *Service) SimulateTerminalTurnOn(ctx context.Context, req *SimulateTerminalTurnOnRequest, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/pa/pos/terminals/turn_on", req, nil, opts...)
}
