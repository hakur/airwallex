package pa

import (
	"context"
	"net/url"
	"strconv"

	"github.com/hakur/airwallex/sdk"
)

// TerminalStatus represents a terminal status.
// TerminalStatus 终端状态。
type TerminalStatus = string

const (
	// TerminalStatusActive indicates the terminal is active.
	// TerminalStatusActive 已激活。
	TerminalStatusActive TerminalStatus = "ACTIVE"
	// TerminalStatusInactive indicates the terminal is inactive.
	// TerminalStatusInactive 未激活。
	TerminalStatusInactive TerminalStatus = "INACTIVE"
	// TerminalStatusTerminated indicates the terminal is terminated.
	// TerminalStatusTerminated 已终止。
	TerminalStatusTerminated TerminalStatus = "TERMINATED"
)

// TerminalPasswordStatus represents a terminal password status.
// TerminalPasswordStatus 终端密码状态。
type TerminalPasswordStatus = string

const (
	// TerminalPasswordStatusActive indicates the password is active.
	// TerminalPasswordStatusActive 密码已激活。
	TerminalPasswordStatusActive TerminalPasswordStatus = "ACTIVE"
	// TerminalPasswordStatusLocked indicates the password is locked.
	// TerminalPasswordStatusLocked 密码已锁定。
	TerminalPasswordStatusLocked TerminalPasswordStatus = "LOCKED"
	// TerminalPasswordStatusResetRequested indicates a password reset has been requested.
	// TerminalPasswordStatusResetRequested 已请求重置密码。
	TerminalPasswordStatusResetRequested TerminalPasswordStatus = "RESET_REQUESTED"
	// TerminalPasswordStatusOptOut indicates the user has opted out of passwords.
	// TerminalPasswordStatusOptOut 已选择不使用密码。
	TerminalPasswordStatusOptOut TerminalPasswordStatus = "OPT_OUT"
)

// TerminalModel represents a terminal model.
// TerminalModel 终端型号。
type TerminalModel = string

const (
	// TerminalModelMorefunM90 indicates the Morefun M90 model.
	// TerminalModelMorefunM90 Morefun M90 型号。
	TerminalModelMorefunM90 TerminalModel = "morefun_m90"
)

// Terminal represents a POS terminal device.
// Terminal 表示 POS 终端设备信息。
type Terminal struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Nickname 昵称。可选。
	Nickname string `json:"nickname,omitempty"`
	// SerialNumber 序列号。可选。
	SerialNumber string `json:"serial_number,omitempty"`
	// Status 终端状态。可选。
	Status TerminalStatus `json:"status,omitempty"`
	// Model 终端型号。可选。
	Model TerminalModel `json:"model,omitempty"`
	// AdminPasswordStatus 管理员密码状态。可选。
	AdminPasswordStatus TerminalPasswordStatus `json:"admin_password_status,omitempty"`
	// RefundPasswordStatus 退款密码状态。可选。
	RefundPasswordStatus TerminalPasswordStatus `json:"refund_password_status,omitempty"`
	// ConnectedAccountID 关联账户唯一标识符。可选。
	ConnectedAccountID string `json:"connected_account_id,omitempty"`
}

// CreateTerminalRequest is the request to create a terminal.
// CreateTerminalRequest 创建终端设备请求。
type CreateTerminalRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// ActivationCode 激活码。必填。
	ActivationCode string `json:"activation_code"`
	// Nickname 昵称。可选。
	Nickname string `json:"nickname,omitempty"`
	// ConnectedAccountID 关联账户唯一标识符。可选。
	ConnectedAccountID string `json:"connected_account_id,omitempty"`
}

// UpdateTerminalRequest is the request to update a terminal.
// UpdateTerminalRequest 更新终端设备请求。
type UpdateTerminalRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Nickname 昵称。可选。
	Nickname string `json:"nickname,omitempty"`
}

// ListTerminalsRequest is the request to query terminal list.
// ListTerminalsRequest 查询终端设备列表请求参数。
type ListTerminalsRequest struct {
	// DeviceModel 设备型号。可选（查询参数）。
	DeviceModel string `json:"-"`
	// Nickname 昵称。可选（查询参数）。
	Nickname string `json:"-"`
	// Page 页码。可选（查询参数）。
	Page string `json:"-"`
	// PageSize 每页数量。可选（查询参数）。
	PageSize int32 `json:"-"`
	// SerialNumber 序列号。可选（查询参数）。
	SerialNumber string `json:"-"`
	// Status 终端状态。可选（查询参数）。
	Status TerminalStatus `json:"-"`
}

// ActivateTerminalRequest is the request to activate a terminal.
// ActivateTerminalRequest 激活终端设备请求。
type ActivateTerminalRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
}

// DeactivateTerminalRequest is the request to deactivate a terminal.
// DeactivateTerminalRequest 停用终端设备请求。
type DeactivateTerminalRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
}

// TerminateTerminalRequest is the request to terminate a terminal.
// TerminateTerminalRequest 终止终端设备请求。
type TerminateTerminalRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
}

// ResetTerminalPasswordRequest is the request to reset terminal password.
// ResetTerminalPasswordRequest 重置终端密码请求。
type ResetTerminalPasswordRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// PasswordType 密码类型（admin 或 refund）。可选。
	PasswordType string `json:"password_type,omitempty"`
}

// CancelTerminalOperationRequest is the request to cancel current terminal operation.
// CancelTerminalOperationRequest 取消终端当前操作请求。
type CancelTerminalOperationRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
}

// ProcessPaymentIntentInTerminalRequest is the request to process payment intent in terminal.
// ProcessPaymentIntentInTerminalRequest 在终端处理支付意图请求。
type ProcessPaymentIntentInTerminalRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// PaymentIntentID 支付意图唯一标识符。必填。
	PaymentIntentID string `json:"payment_intent_id"`
	// PaymentMethodOptions 支付方式选项。可选。
	PaymentMethodOptions map[string]any `json:"payment_method_options,omitempty"`
}

// CreateTerminal creates a terminal device.
// CreateTerminal 创建终端设备。
// 官方文档: https://www.airwallex.com/docs/api/payments/terminals/create.md
func (s *Service) CreateTerminal(ctx context.Context, req *CreateTerminalRequest, opts ...sdk.RequestOption) (*Terminal, error) {
	var resp Terminal
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/pos/terminals/create", req, &resp, opts...)
	return &resp, err
}

// GetTerminal retrieves a terminal by ID.
// GetTerminal 根据 ID 获取终端设备。
// 官方文档: https://www.airwallex.com/docs/api/payments/terminals/retrieve.md
func (s *Service) GetTerminal(ctx context.Context, id string, opts ...sdk.RequestOption) (*Terminal, error) {
	var resp Terminal
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/pos/terminals/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateTerminal updates a terminal device.
// UpdateTerminal 更新终端设备。
// 官方文档: https://www.airwallex.com/docs/api/payments/terminals/update.md
func (s *Service) UpdateTerminal(ctx context.Context, id string, req *UpdateTerminalRequest, opts ...sdk.RequestOption) (*Terminal, error) {
	var resp Terminal
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/pos/terminals/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ListTerminals lists terminal devices.
// ListTerminals 列出终端设备。
// 官方文档: https://www.airwallex.com/docs/api/payments/terminals/list.md
func (s *Service) ListTerminals(ctx context.Context, req *ListTerminalsRequest, opts ...sdk.RequestOption) (*sdk.ListResult[Terminal], error) {
	path := "/api/v1/pa/pos/terminals"
	if req != nil {
		q := url.Values{}
		if req.DeviceModel != "" {
			q.Set("device_model", req.DeviceModel)
		}
		if req.Nickname != "" {
			q.Set("nickname", req.Nickname)
		}
		if req.Page != "" {
			q.Set("page", req.Page)
		}
		if req.PageSize > 0 {
			q.Set("page_size", strconv.Itoa(int(req.PageSize)))
		}
		if req.SerialNumber != "" {
			q.Set("serial_number", req.SerialNumber)
		}
		if req.Status != "" {
			q.Set("status", string(req.Status))
		}
		if len(q) > 0 {
			path += "?" + q.Encode()
		}
	}
	var resp sdk.ListResult[Terminal]
	err := s.doer.Do(ctx, "GET", path, nil, &resp, opts...)
	return &resp, err
}

// ActivateTerminal activates a terminal device.
// ActivateTerminal 激活终端设备。
// 官方文档: https://www.airwallex.com/docs/api/payments/terminals/activate.md
func (s *Service) ActivateTerminal(ctx context.Context, id string, req *ActivateTerminalRequest, opts ...sdk.RequestOption) (*Terminal, error) {
	var resp Terminal
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/pos/terminals/"+id+"/activate", req, &resp, opts...)
	return &resp, err
}

// DeactivateTerminal deactivates a terminal device.
// DeactivateTerminal 停用终端设备。
// 官方文档: https://www.airwallex.com/docs/api/payments/terminals/deactivate.md
func (s *Service) DeactivateTerminal(ctx context.Context, id string, req *DeactivateTerminalRequest, opts ...sdk.RequestOption) (*Terminal, error) {
	var resp Terminal
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/pos/terminals/"+id+"/deactivate", req, &resp, opts...)
	return &resp, err
}

// TerminateTerminal terminates a terminal device.
// TerminateTerminal 终止终端设备。
// 官方文档: https://www.airwallex.com/docs/api/payments/terminals/terminate.md
func (s *Service) TerminateTerminal(ctx context.Context, id string, req *TerminateTerminalRequest, opts ...sdk.RequestOption) (*Terminal, error) {
	var resp Terminal
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/pos/terminals/"+id+"/terminate", req, &resp, opts...)
	return &resp, err
}

// ResetTerminalPassword resets the terminal password.
// ResetTerminalPassword 重置终端密码。
// 官方文档: https://www.airwallex.com/docs/api/payments/terminals/reset_password.md
func (s *Service) ResetTerminalPassword(ctx context.Context, id string, req *ResetTerminalPasswordRequest, opts ...sdk.RequestOption) (*Terminal, error) {
	var resp Terminal
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/pos/terminals/"+id+"/reset_password", req, &resp, opts...)
	return &resp, err
}

// CancelCurrentOperation cancels the current terminal operation.
// CancelCurrentOperation 取消终端当前操作。
// 官方文档: https://www.airwallex.com/docs/api/payments/terminals/cancel_current_operation.md
func (s *Service) CancelCurrentOperation(ctx context.Context, id string, req *CancelTerminalOperationRequest, opts ...sdk.RequestOption) (*Terminal, error) {
	var resp Terminal
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/pos/terminals/"+id+"/cancel_current_operation", req, &resp, opts...)
	return &resp, err
}

// ProcessPaymentIntentInTerminal processes a payment intent in the terminal.
// ProcessPaymentIntentInTerminal 在终端处理支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/terminals/process_payment_intent.md
func (s *Service) ProcessPaymentIntentInTerminal(ctx context.Context, id string, req *ProcessPaymentIntentInTerminalRequest, opts ...sdk.RequestOption) (*Terminal, error) {
	var resp Terminal
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/pos/terminals/"+id+"/process_payment_intent", req, &resp, opts...)
	return &resp, err
}
