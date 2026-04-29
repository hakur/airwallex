package pa

import (
	"context"
	"net/url"
	"strconv"

	"github.com/hakur/airwallex/sdk"
)

// CustomsDeclarationStatus represents a customs declaration status.
// CustomsDeclarationStatus 表示海关申报状态。
type CustomsDeclarationStatus = string

const (
	// CustomsDeclarationStatusUndeclared indicates the declaration is undeclared.
	// CustomsDeclarationStatusUndeclared 未申报。
	CustomsDeclarationStatusUndeclared CustomsDeclarationStatus = "UNDECLARED"
	// CustomsDeclarationStatusProcessing indicates the declaration is processing.
	// CustomsDeclarationStatusProcessing 处理中。
	CustomsDeclarationStatusProcessing CustomsDeclarationStatus = "PROCESSING"
	// CustomsDeclarationStatusSuccess indicates the declaration succeeded.
	// CustomsDeclarationStatusSuccess 申报成功。
	CustomsDeclarationStatusSuccess CustomsDeclarationStatus = "SUCCESS"
	// CustomsDeclarationStatusFail indicates the declaration failed.
	// CustomsDeclarationStatusFail 申报失败。
	CustomsDeclarationStatusFail CustomsDeclarationStatus = "FAIL"
	// CustomsDeclarationStatusDeclared indicates the declaration is declared.
	// CustomsDeclarationStatusDeclared 已申报。
	CustomsDeclarationStatusDeclared CustomsDeclarationStatus = "DECLARED"
	// CustomsDeclarationStatusRedeclared indicates the declaration is redeclared.
	// CustomsDeclarationStatusRedeclared 已重新申报。
	CustomsDeclarationStatusRedeclared CustomsDeclarationStatus = "REDECLARED"
)

// CustomsDeclaration represents a customs declaration.
// CustomsDeclaration 表示海关申报信息。
type CustomsDeclaration struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Status is the declaration status. Required.
	// Status 申报状态。必填。
	Status string `json:"status"`
	// PaymentMethodType is the payment method type. Optional.
	// PaymentMethodType 支付方式类型。可选。
	PaymentMethodType string `json:"payment_method_type,omitempty"`
	// ProviderTransactionID is the provider transaction ID. Optional.
	// ProviderTransactionID 支付提供商交易号。可选。
	ProviderTransactionID string `json:"provider_transaction_id,omitempty"`
	// AWXRequestID is the Airwallex request unique identifier. Optional.
	// AWXRequestID Airwallex 请求唯一标识符。可选。
	AWXRequestID string `json:"awx_request_id,omitempty"`
	// VerificationDepartmentCode is the verification department code. Optional.
	// VerificationDepartmentCode 核验机构代码。可选。
	VerificationDepartmentCode string `json:"verification_department_code,omitempty"`
	// VerificationDepartmentTransactionID is the verification department transaction ID. Optional.
	// VerificationDepartmentTransactionID 核验机构交易号。可选。
	VerificationDepartmentTransactionID string `json:"verification_department_transaction_id,omitempty"`
	// CustomsDetails is the customs declaration details. Optional.
	// CustomsDetails 海关申报详情。可选。
	CustomsDetails *CustomsDetails `json:"customs_details,omitempty"`
	// SubOrder is the sub-order information. Optional.
	// SubOrder 子订单信息。可选。
	SubOrder *SubOrder `json:"sub_order,omitempty"`
	// ShopperIdentityCheckResult is the shopper identity check result. Optional.
	// ShopperIdentityCheckResult 购物者身份核验结果。可选。
	ShopperIdentityCheckResult string `json:"shopper_identity_check_result,omitempty"`
	// CustomsStatusMessage is the customs status message. Optional.
	// CustomsStatusMessage 海关状态消息。可选。
	CustomsStatusMessage string `json:"customs_status_message,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt is the update time. Optional.
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
}

// CustomsDetails represents customs declaration details.
// CustomsDetails 表示海关申报详情。
type CustomsDetails struct {
	// CustomsCode is the customs code. Required.
	// CustomsCode 海关代码。必填。
	CustomsCode string `json:"customs_code"`
	// MerchantCustomsName is the merchant customs name. Required.
	// MerchantCustomsName 商户海关名称。必填。
	MerchantCustomsName string `json:"merchant_customs_name"`
	// MerchantCustomsNumber is the merchant customs number. Required.
	// MerchantCustomsNumber 商户海关编号。必填。
	MerchantCustomsNumber string `json:"merchant_customs_number"`
}

// SubOrder represents sub-order information.
// SubOrder 表示子订单信息。
type SubOrder struct {
	// OrderNumber is the order number. Required.
	// OrderNumber 订单编号。必填。
	OrderNumber string `json:"order_number"`
	// ProviderOrderID is the provider order ID. Optional.
	// ProviderOrderID 支付提供商订单号。可选。
	ProviderOrderID string `json:"provider_order_id,omitempty"`
	// Amount is the amount. Required.
	// Amount 金额。必填。
	Amount float64 `json:"amount"`
	// ShippingFee is the shipping fee. Required.
	// ShippingFee 运费。必填。
	ShippingFee float64 `json:"shipping_fee"`
	// Currency is the currency code. Required.
	// Currency 货币代码。必填。
	Currency string `json:"currency"`
}

// ShopperDetails represents shopper identity information.
// ShopperDetails 表示购物者身份信息。
type ShopperDetails struct {
	// ShopperID is the shopper unique identifier. Required.
	// ShopperID 购物者唯一标识符。必填。
	ShopperID string `json:"shopper_id"`
	// ShopperName is the shopper name. Required.
	// ShopperName 购物者姓名。必填。
	ShopperName string `json:"shopper_name"`
}

// CreateCustomsDeclarationRequest is the request to create a customs declaration.
// CreateCustomsDeclarationRequest 创建海关申报请求。
type CreateCustomsDeclarationRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// PaymentIntentID is the payment intent unique identifier. Required.
	// PaymentIntentID 支付意图唯一标识符。必填。
	PaymentIntentID string `json:"payment_intent_id"`
	// CustomsDetails is the customs declaration details. Required.
	// CustomsDetails 海关申报详情。必填。
	CustomsDetails *CustomsDetails `json:"customs_details"`
	// SubOrder is the sub-order information. Required.
	// SubOrder 子订单信息。必填。
	SubOrder *SubOrder `json:"sub_order"`
	// ShopperDetails is the shopper identity information. Optional.
	// ShopperDetails 购物者身份信息。可选。
	ShopperDetails *ShopperDetails `json:"shopper_details,omitempty"`
	// VerificationDepartmentCode is the verification department code. Optional.
	// VerificationDepartmentCode 核验机构代码。可选。
	VerificationDepartmentCode string `json:"verification_department_code,omitempty"`
	// VerificationDepartmentTransactionID is the verification department transaction ID. Optional.
	// VerificationDepartmentTransactionID 核验机构交易号。可选。
	VerificationDepartmentTransactionID string `json:"verification_department_transaction_id,omitempty"`
}

// UpdateCustomsDeclarationRequest is the request to update a customs declaration.
// UpdateCustomsDeclarationRequest 更新海关申报请求。
// 注意：Update 的请求体使用扁平字段而非嵌套对象。
type UpdateCustomsDeclarationRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// CustomsCode is the customs code. Optional.
	// CustomsCode 海关代码。可选。
	CustomsCode string `json:"customs_code,omitempty"`
	// MerchantCustomsName is the merchant customs name. Optional.
	// MerchantCustomsName 商户海关名称。可选。
	MerchantCustomsName string `json:"merchant_customs_name,omitempty"`
	// MerchantCustomsNumber is the merchant customs number. Optional.
	// MerchantCustomsNumber 商户海关编号。可选。
	MerchantCustomsNumber string `json:"merchant_customs_number,omitempty"`
	// ShopperDetails 购物者身份信息。可选。
	ShopperDetails *ShopperDetails `json:"shopper_details,omitempty"`
	// SubOrderAmount is the sub-order amount. Optional.
	// SubOrderAmount 子订单金额。可选。
	SubOrderAmount float64 `json:"sub_order_amount,omitempty"`
}

// RedeclareCustomsDeclarationRequest is the request to redeclare a customs declaration.
// RedeclareCustomsDeclarationRequest 重新申报海关申报请求。
type RedeclareCustomsDeclarationRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
}

// ListCustomsDeclarationsRequest is the request to list customs declarations.
// ListCustomsDeclarationsRequest 列出海关申报请求参数。
type ListCustomsDeclarationsRequest struct {
	// Page is the page number. Optional (query param).
	// Page 页码。可选（查询参数）。
	Page int32 `json:"-"`
	// PageSize is the page size. Optional (query param).
	// PageSize 每页数量。可选（查询参数）。
	PageSize int32 `json:"-"`
	// Status is the declaration status. Optional (query param).
	// Status 申报状态。可选（查询参数）。
	Status string `json:"-"`
}

// CreateCustomsDeclaration creates a customs declaration.
// CreateCustomsDeclaration 创建海关申报。
// 官方文档: https://www.airwallex.com/docs/api/payments/customs_declarations/create.md
func (s *Service) CreateCustomsDeclaration(ctx context.Context, req *CreateCustomsDeclarationRequest, opts ...sdk.RequestOption) (*CustomsDeclaration, error) {
	var resp CustomsDeclaration
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/customs_declarations/create", req, &resp, opts...)
	return &resp, err
}

// GetCustomsDeclaration retrieves a customs declaration by ID.
// GetCustomsDeclaration 根据 ID 获取海关申报。
// 官方文档: https://www.airwallex.com/docs/api/payments/customs_declarations/retrieve.md
func (s *Service) GetCustomsDeclaration(ctx context.Context, id string, opts ...sdk.RequestOption) (*CustomsDeclaration, error) {
	var resp CustomsDeclaration
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/customs_declarations/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateCustomsDeclaration updates a customs declaration.
// UpdateCustomsDeclaration 更新海关申报。
// 官方文档: https://www.airwallex.com/docs/api/payments/customs_declarations/update.md
func (s *Service) UpdateCustomsDeclaration(ctx context.Context, id string, req *UpdateCustomsDeclarationRequest, opts ...sdk.RequestOption) (*CustomsDeclaration, error) {
	var resp CustomsDeclaration
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/customs_declarations/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// RedeclareCustomsDeclaration redeclares a customs declaration.
// RedeclareCustomsDeclaration 重新申报海关申报。
// 官方文档: https://www.airwallex.com/docs/api/payments/customs_declarations/redeclare.md
func (s *Service) RedeclareCustomsDeclaration(ctx context.Context, id string, req *RedeclareCustomsDeclarationRequest, opts ...sdk.RequestOption) (*CustomsDeclaration, error) {
	var resp CustomsDeclaration
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/customs_declarations/"+id+"/redeclare", req, &resp, opts...)
	return &resp, err
}

// ListCustomsDeclarations lists customs declarations.
// ListCustomsDeclarations 列出海关申报。
// 官方文档: https://www.airwallex.com/docs/api/payments/customs_declarations/list.md
func (s *Service) ListCustomsDeclarations(ctx context.Context, req *ListCustomsDeclarationsRequest, opts ...sdk.RequestOption) (*sdk.ListResult[CustomsDeclaration], error) {
	path := "/api/v1/pa/customs_declarations"
	if req != nil {
		q := url.Values{}
		if req.Page > 0 {
			q.Set("page", strconv.Itoa(int(req.Page)))
		}
		if req.PageSize > 0 {
			q.Set("page_size", strconv.Itoa(int(req.PageSize)))
		}
		if req.Status != "" {
			q.Set("status", req.Status)
		}
		if len(q) > 0 {
			path += "?" + q.Encode()
		}
	}
	var resp sdk.ListResult[CustomsDeclaration]
	err := s.doer.Do(ctx, "GET", path, nil, &resp, opts...)
	return &resp, err
}
