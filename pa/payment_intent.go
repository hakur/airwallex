package pa

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PaymentIntentStatus represents a payment intent status.
// PaymentIntentStatus 支付意图状态枚举。
type PaymentIntentStatus = string

const (
	// PaymentIntentStatusRequiresPaymentMethod indicates a payment method is required.
	// PaymentIntentStatusRequiresPaymentMethod 需要支付方式。
	PaymentIntentStatusRequiresPaymentMethod PaymentIntentStatus = "REQUIRES_PAYMENT_METHOD"
	// PaymentIntentStatusRequiresCustomerAction indicates customer action is required.
	// PaymentIntentStatusRequiresCustomerAction 需要客户操作。
	PaymentIntentStatusRequiresCustomerAction PaymentIntentStatus = "REQUIRES_CUSTOMER_ACTION"
	// PaymentIntentStatusRequiresCapture indicates capture is required.
	// PaymentIntentStatusRequiresCapture 需要捕获。
	PaymentIntentStatusRequiresCapture PaymentIntentStatus = "REQUIRES_CAPTURE"
	// PaymentIntentStatusProcessing indicates the payment is processing.
	// PaymentIntentStatusProcessing 处理中。
	PaymentIntentStatusProcessing PaymentIntentStatus = "PROCESSING"
	// PaymentIntentStatusSucceeded indicates the payment succeeded.
	// PaymentIntentStatusSucceeded 成功。
	PaymentIntentStatusSucceeded PaymentIntentStatus = "SUCCEEDED"
	// PaymentIntentStatusCancelled indicates the payment was cancelled.
	// PaymentIntentStatusCancelled 已取消。
	PaymentIntentStatusCancelled PaymentIntentStatus = "CANCELLED"
)

// PaymentIntent represents a payment intent.
// PaymentIntent 表示支付意图信息。
type PaymentIntent struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// RequestID is the unique request identifier. Optional.
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// Amount is the amount. Required.
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// MerchantOrderID is the merchant order ID. Optional.
	// MerchantOrderID 商户订单号。可选。
	MerchantOrderID string `json:"merchant_order_id,omitempty"`
	// Descriptor is the transaction descriptor. Optional.
	// Descriptor 交易描述。可选。
	Descriptor string `json:"descriptor,omitempty"`
	// CustomerID is the customer unique identifier. Optional.
	// CustomerID 客户唯一标识符。可选。
	CustomerID string `json:"customer_id,omitempty"`
	// Status is the payment intent status. Required.
	// Status 支付意图状态。必填。
	Status PaymentIntentStatus `json:"status"`
	// CapturedAmount is the captured amount. Optional.
	// CapturedAmount 已捕获金额。可选。
	CapturedAmount float64 `json:"captured_amount,omitempty"`
	// OriginalAmount is the original amount. Optional.
	// OriginalAmount 原始金额。可选。
	OriginalAmount float64 `json:"original_amount,omitempty"`
	// OriginalCurrency is the original currency code. Optional.
	// OriginalCurrency 原始货币代码。可选。
	OriginalCurrency sdk.Currency `json:"original_currency,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt is the update time. Optional.
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
	// ClientSecret is the client secret. Optional.
	// ClientSecret 客户端密钥。可选。
	ClientSecret string `json:"client_secret,omitempty"`
	// NextAction is the next action information. Optional.
	// NextAction 下一步操作信息。可选。
	NextAction map[string]any `json:"next_action,omitempty"`
	// PaymentMethod is the payment method information. Optional.
	// PaymentMethod 支付方式信息。可选。
	PaymentMethod map[string]any `json:"payment_method,omitempty"`
	// LatestPaymentAttempt is the latest payment attempt information. Optional.
	// LatestPaymentAttempt 最新支付尝试信息。可选。
	LatestPaymentAttempt map[string]any `json:"latest_payment_attempt,omitempty"`
	// AdditionalInfo 附加信息（渠道特定数据统一字段）。可选。
	AdditionalInfo map[string]any `json:"additional_info,omitempty"`
}

// CreatePaymentIntentRequest is the request to create a payment intent.
// CreatePaymentIntentRequest 创建支付意图请求。
type CreatePaymentIntentRequest struct {
	// RequestID is the unique request identifier. Optional.
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// MerchantOrderID is the merchant order ID. Optional.
	// MerchantOrderID 商户订单号。可选。
	MerchantOrderID string `json:"merchant_order_id,omitempty"`
	// Amount is the amount. Required.
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency sdk.Currency `json:"currency"`
	// Descriptor is the transaction descriptor. Optional.
	// Descriptor 交易描述。可选。
	Descriptor string `json:"descriptor,omitempty"`
	// CustomerID is the customer unique identifier. Optional.
	// CustomerID 客户唯一标识符。可选。
	CustomerID string `json:"customer_id,omitempty"`
	// PaymentMethod is the payment method details. Optional.
	// PaymentMethod 支付方式详情。可选。
	PaymentMethod *PaymentMethodInput `json:"payment_method,omitempty"`
	// PaymentMethodOptions are the payment method options. Optional.
	// PaymentMethodOptions 支付方式选项。可选。
	PaymentMethodOptions map[string]any `json:"payment_method_options,omitempty"`
	// ReturnURL 返回地址。可选。
	ReturnURL string `json:"return_url,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// CaptureMethod 捕获方式。可选。
	CaptureMethod string `json:"capture_method,omitempty"`
	// SetupFuturePayment 是否设置未来支付。可选。
	SetupFuturePayment bool `json:"setup_future_payment,omitempty"`
	// NotificationURL 通知地址。可选。
	NotificationURL string `json:"notification_url,omitempty"`
	// PaymentType 支付类型。可选。
	PaymentType string `json:"payment_type,omitempty"`
	// PlatformInitiated indicates whether initiated by platform. Optional.
	// PlatformInitiated 是否平台发起。可选。
	PlatformInitiated bool `json:"platform_initiated,omitempty"`
	// Force3DS indicates whether to force 3D Secure. Optional.
	// Force3DS 是否强制 3D Secure。可选。
	Force3DS bool `json:"force_3ds,omitempty"`
	// PaymentMethodTypes are the supported payment method types. Optional.
	// PaymentMethodTypes 支持的支付方式类型列表。可选。
	PaymentMethodTypes []string `json:"payment_method_types,omitempty"`
	// CustomerData is the customer data. Optional.
	// CustomerData 客户数据。可选。
	CustomerData map[string]any `json:"customer_data,omitempty"`
	// Order is the order information. Optional.
	// Order 订单信息。可选。
	Order map[string]any `json:"order,omitempty"`
	// Shipping is the shipping information. Optional.
	// Shipping 配送信息。可选。
	Shipping map[string]any `json:"shipping,omitempty"`
	// Billing is the billing information. Optional.
	// Billing 账单信息。可选。
	Billing map[string]any `json:"billing,omitempty"`
	// AdditionalInfo is additional information (unified field for channel-specific data). Optional.
	// AdditionalInfo 附加信息（渠道特定数据统一字段）。可选。
	AdditionalInfo map[string]any `json:"additional_info,omitempty"`
}

// UpdatePaymentIntentRequest is the request to update a payment intent.
// UpdatePaymentIntentRequest 更新支付意图请求。
type UpdatePaymentIntentRequest struct {
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// Amount is the amount. Optional.
	// Amount 金额。可选。
	Amount float64 `json:"amount,omitempty"`
	// Currency is the currency code. Optional.
	// Currency 货币代码。可选。
	Currency sdk.Currency `json:"currency,omitempty"`
	// Descriptor is the transaction descriptor. Optional.
	// Descriptor 交易描述。可选。
	Descriptor string `json:"descriptor,omitempty"`
	// CustomerID is the customer unique identifier. Optional.
	// CustomerID 客户唯一标识符。可选。
	CustomerID string `json:"customer_id,omitempty"`
	// PaymentMethod is the payment method details. Optional.
	// PaymentMethod 支付方式详情。可选。
	PaymentMethod *PaymentMethodInput `json:"payment_method,omitempty"`
	// PaymentMethodOptions are the payment method options. Optional.
	// PaymentMethodOptions 支付方式选项。可选。
	PaymentMethodOptions map[string]any `json:"payment_method_options,omitempty"`
	// ReturnURL is the return URL. Optional.
	// ReturnURL 返回地址。可选。
	ReturnURL string `json:"return_url,omitempty"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// CaptureMethod is the capture method. Optional.
	// CaptureMethod 捕获方式。可选。
	CaptureMethod string `json:"capture_method,omitempty"`
	// SetupFuturePayment indicates whether to set up future payment. Optional.
	// SetupFuturePayment 是否设置未来支付。可选。
	SetupFuturePayment bool `json:"setup_future_payment,omitempty"`
	// NotificationURL is the notification URL. Optional.
	// NotificationURL 通知地址。可选。
	NotificationURL string `json:"notification_url,omitempty"`
	// PaymentType is the payment type. Optional.
	// PaymentType 支付类型。可选。
	PaymentType string `json:"payment_type,omitempty"`
	// Force3DS indicates whether to force 3D Secure. Optional.
	// Force3DS 是否强制 3D Secure。可选。
	Force3DS bool `json:"force_3ds,omitempty"`
	// PaymentMethodTypes are the supported payment method types. Optional.
	// PaymentMethodTypes 支持的支付方式类型列表。可选。
	PaymentMethodTypes []string `json:"payment_method_types,omitempty"`
	// CustomerData is the customer data. Optional.
	// CustomerData 客户数据。可选。
	CustomerData map[string]any `json:"customer_data,omitempty"`
	// Order is the order information. Optional.
	// Order 订单信息。可选。
	Order map[string]any `json:"order,omitempty"`
	// Shipping is the shipping information. Optional.
	// Shipping 配送信息。可选。
	Shipping map[string]any `json:"shipping,omitempty"`
	// Billing is the billing information. Optional.
	// Billing 账单信息。可选。
	Billing map[string]any `json:"billing,omitempty"`
	// AdditionalInfo is additional information (unified field for channel-specific data). Optional.
	// AdditionalInfo 附加信息（渠道特定数据统一字段）。可选。
	AdditionalInfo map[string]any `json:"additional_info,omitempty"`
}

// ConfirmPaymentIntentRequest is the request to confirm a payment intent.
// ConfirmPaymentIntentRequest 确认支付意图请求。
type ConfirmPaymentIntentRequest struct {
	// RequestID is the unique request identifier. Optional.
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// PaymentMethod is the payment method details. Optional.
	// PaymentMethod 支付方式详情。可选。
	PaymentMethod *PaymentMethodInput `json:"payment_method,omitempty"`
	// PaymentMethodOptions are the payment method options. Optional.
	// PaymentMethodOptions 支付方式选项。可选。
	PaymentMethodOptions map[string]any `json:"payment_method_options,omitempty"`
	// ReturnURL is the return URL. Optional.
	// ReturnURL 返回地址。可选。
	ReturnURL string `json:"return_url,omitempty"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// CaptureMethod is the capture method. Optional.
	// CaptureMethod 捕获方式。可选。
	CaptureMethod string `json:"capture_method,omitempty"`
	// SetupFuturePayment indicates whether to set up future payment. Optional.
	// SetupFuturePayment 是否设置未来支付。可选。
	SetupFuturePayment bool `json:"setup_future_payment,omitempty"`
	// NotificationURL is the notification URL. Optional.
	// NotificationURL 通知地址。可选。
	NotificationURL string `json:"notification_url,omitempty"`
	// PaymentType is the payment type. Optional.
	// PaymentType 支付类型。可选。
	PaymentType string `json:"payment_type,omitempty"`
	// Force3DS indicates whether to force 3D Secure. Optional.
	// Force3DS 是否强制 3D Secure。可选。
	Force3DS bool `json:"force_3ds,omitempty"`
	// CustomerData is the customer data. Optional.
	// CustomerData 客户数据。可选。
	CustomerData map[string]any `json:"customer_data,omitempty"`
	// Order is the order information. Optional.
	// Order 订单信息。可选。
	Order map[string]any `json:"order,omitempty"`
	// Shipping is the shipping information. Optional.
	// Shipping 配送信息。可选。
	Shipping map[string]any `json:"shipping,omitempty"`
	// Billing is the billing information. Optional.
	// Billing 账单信息。可选。
	Billing map[string]any `json:"billing,omitempty"`
	// AdditionalInfo is additional information (unified field for channel-specific data). Optional.
	// AdditionalInfo 附加信息（渠道特定数据统一字段）。可选。
	AdditionalInfo map[string]any `json:"additional_info,omitempty"`
}

// CapturePaymentIntentRequest is the request to capture a payment intent.
// CapturePaymentIntentRequest 捕获支付意图请求。
type CapturePaymentIntentRequest struct {
	// RequestID is the unique request identifier. Optional.
	// RequestID 请求唯一标识符。可选。
	RequestID string `json:"request_id,omitempty"`
	// Amount is the capture amount. Optional.
	// Amount 捕获金额。可选。
	Amount float64 `json:"amount,omitempty"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// CreatePaymentIntent creates a payment intent.
// CreatePaymentIntent 创建支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/create.md
func (s *Service) CreatePaymentIntent(ctx context.Context, req *CreatePaymentIntentRequest, opts ...sdk.RequestOption) (*PaymentIntent, error) {
	var resp PaymentIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_intents/create", req, &resp, opts...)
	return &resp, err
}

// GetPaymentIntent retrieves a payment intent by ID.
// GetPaymentIntent 根据 ID 获取支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/retrieve.md
func (s *Service) GetPaymentIntent(ctx context.Context, id string, opts ...sdk.RequestOption) (*PaymentIntent, error) {
	var resp PaymentIntent
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_intents/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdatePaymentIntent updates a payment intent.
// UpdatePaymentIntent 更新支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/update.md
func (s *Service) UpdatePaymentIntent(ctx context.Context, id string, req *UpdatePaymentIntentRequest, opts ...sdk.RequestOption) (*PaymentIntent, error) {
	var resp PaymentIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_intents/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ConfirmPaymentIntent confirms a payment intent.
// ConfirmPaymentIntent 确认支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/confirm.md
func (s *Service) ConfirmPaymentIntent(ctx context.Context, id string, req *ConfirmPaymentIntentRequest, opts ...sdk.RequestOption) (*PaymentIntent, error) {
	var resp PaymentIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_intents/"+id+"/confirm", req, &resp, opts...)
	return &resp, err
}

// CapturePaymentIntent captures a payment intent.
// CapturePaymentIntent 捕获支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/capture.md
func (s *Service) CapturePaymentIntent(ctx context.Context, id string, req *CapturePaymentIntentRequest, opts ...sdk.RequestOption) (*PaymentIntent, error) {
	var resp PaymentIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_intents/"+id+"/capture", req, &resp, opts...)
	return &resp, err
}

// CancelPaymentIntentRequest is the request to cancel a payment intent.
// CancelPaymentIntentRequest 取消支付意图请求。
type CancelPaymentIntentRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
}

// CancelPaymentIntent cancels a payment intent.
// CancelPaymentIntent 取消支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/cancel.md
func (s *Service) CancelPaymentIntent(ctx context.Context, id string, req *CancelPaymentIntentRequest, opts ...sdk.RequestOption) (*PaymentIntent, error) {
	var resp PaymentIntent
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/payment_intents/"+id+"/cancel", req, &resp, opts...)
	return &resp, err
}

// ListPaymentIntents lists payment intents.
// ListPaymentIntents 列出支付意图。
// 官方文档: https://www.airwallex.com/docs/api/payments/payment_intents/list.md
func (s *Service) ListPaymentIntents(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[PaymentIntent], error) {
	var resp sdk.ListResult[PaymentIntent]
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/payment_intents", nil, &resp, opts...)
	return &resp, err
}
