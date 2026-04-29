package pa

import (
	"context"
	"net/url"
	"strconv"

	"github.com/hakur/airwallex/sdk"
)

// PaymentMethodTypeConfig represents a payment method type configuration.
// PaymentMethodTypeConfig 表示支付方式类型配置。
type PaymentMethodTypeConfig struct {
	// Active indicates whether the config is active. Required.
	// Active 是否激活。必填。
	Active bool `json:"active"`
	// Flows is the list of supported payment flows. Required.
	// Flows 支持的支付流程列表。必填。
	Flows []string `json:"flows"`
	// Name is the payment method name. Required.
	// Name 支付方式名称。必填。
	Name string `json:"name"`
	// TransactionCurrencies is the list of supported transaction currencies. Required.
	// TransactionCurrencies 支持的交易货币列表。必填。
	TransactionCurrencies []string `json:"transaction_currencies"`
	// TransactionMode is the transaction mode. Required.
	// TransactionMode 交易模式。必填。
	TransactionMode string `json:"transaction_mode"`
}

// GetPaymentMethodTypesRequest is the request to query payment method types.
// GetPaymentMethodTypesRequest 查询支付方式类型请求参数。
type GetPaymentMethodTypesRequest struct {
	// Active indicates whether to return only active payment methods. Optional (query param, zero value false not sent).
	// Active 是否只返回激活的支付方式。可选（查询参数，无 omitempty，但零值 false 不发送）。
	Active bool `json:"-"`
	// CountryCode is the country code. Optional.
	// CountryCode 国家代码。可选。
	CountryCode string `json:"-"`
	// PageNum is the page number. Optional.
	// PageNum 页码。可选。
	PageNum int32 `json:"-"`
	// PageSize is the page size. Optional.
	// PageSize 每页数量。可选。
	PageSize int32 `json:"-"`
	// TransactionCurrency is the transaction currency. Optional.
	// TransactionCurrency 交易货币。可选。
	TransactionCurrency string `json:"-"`
	// TransactionMode is the transaction mode. Optional.
	// TransactionMode 交易模式。可选。
	TransactionMode string `json:"-"`
}

// GetPaymentMethodTypes retrieves available payment method types.
// GetPaymentMethodTypes 查询可用的支付方式类型。
// 官方文档: https://www.airwallex.com/docs/api/payments/config/payment_method_types.md
func (s *Service) GetPaymentMethodTypes(ctx context.Context, req *GetPaymentMethodTypesRequest, opts ...sdk.RequestOption) (*sdk.ListResult[PaymentMethodTypeConfig], error) {
	path := "/api/v1/pa/config/payment_method_types"
	if req != nil {
		q := url.Values{}
		if req.Active {
			q.Set("active", "true")
		}
		if req.CountryCode != "" {
			q.Set("country_code", req.CountryCode)
		}
		if req.PageNum > 0 {
			q.Set("page_num", strconv.Itoa(int(req.PageNum)))
		}
		if req.PageSize > 0 {
			q.Set("page_size", strconv.Itoa(int(req.PageSize)))
		}
		if req.TransactionCurrency != "" {
			q.Set("transaction_currency", req.TransactionCurrency)
		}
		if req.TransactionMode != "" {
			q.Set("transaction_mode", req.TransactionMode)
		}
		if len(q) > 0 {
			path += "?" + q.Encode()
		}
	}
	var resp sdk.ListResult[PaymentMethodTypeConfig]
	err := s.doer.Do(ctx, "GET", path, nil, &resp, opts...)
	return &resp, err
}

// BankConfig represents bank configuration.
// BankConfig 表示银行配置信息。
type BankConfig struct {
	// BankName is the bank name. Required.
	// BankName 银行名称。必填。
	BankName string `json:"bank_name"`
	// DisplayName is the display name. Required.
	// DisplayName 显示名称。必填。
	DisplayName string `json:"display_name"`
	// Resources is the bank resource information. Required.
	// Resources 银行资源信息。必填。
	Resources BankResources `json:"resources"`
}

// BankResources represents bank resource information.
// BankResources 银行资源信息。
type BankResources struct {
	// LogoURL is the bank logo URL. Optional.
	// LogoURL 银行 Logo 地址。可选。
	LogoURL string `json:"logo_url,omitempty"`
}

// GetBanksRequest is the request to query bank list.
// GetBanksRequest 查询银行列表请求参数。
type GetBanksRequest struct {
	// CountryCode is the country code. Optional.
	// CountryCode 国家代码。可选。
	CountryCode string `json:"-"`
	// PageNum is the page number. Optional.
	// PageNum 页码。可选。
	PageNum int32 `json:"-"`
	// PageSize is the page size. Optional.
	// PageSize 每页数量。可选。
	PageSize int32 `json:"-"`
	// PaymentMethodType is the payment method type. Optional.
	// PaymentMethodType 支付方式类型。可选。
	PaymentMethodType string `json:"-"`
}

// GetBanks retrieves available bank list.
// GetBanks 查询可用的银行名称列表。
// 官方文档: https://www.airwallex.com/docs/api/payments/config/banks.md
func (s *Service) GetBanks(ctx context.Context, req *GetBanksRequest, opts ...sdk.RequestOption) (*sdk.ListResult[BankConfig], error) {
	path := "/api/v1/pa/config/banks"
	if req != nil {
		q := url.Values{}
		if req.CountryCode != "" {
			q.Set("country_code", req.CountryCode)
		}
		if req.PageNum > 0 {
			q.Set("page_num", strconv.Itoa(int(req.PageNum)))
		}
		if req.PageSize > 0 {
			q.Set("page_size", strconv.Itoa(int(req.PageSize)))
		}
		if req.PaymentMethodType != "" {
			q.Set("payment_method_type", req.PaymentMethodType)
		}
		if len(q) > 0 {
			path += "?" + q.Encode()
		}
	}
	var resp sdk.ListResult[BankConfig]
	err := s.doer.Do(ctx, "GET", path, nil, &resp, opts...)
	return &resp, err
}

// GetConvertibleShopperCurrencies retrieves shopper currencies available for currency conversion.
// GetConvertibleShopperCurrencies 查询可用于支付货币转换的购物货币列表。
// 官方文档: https://www.airwallex.com/docs/api/payments/config/convertible_shopper_currencies.md
func (s *Service) GetConvertibleShopperCurrencies(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[sdk.Currency], error) {
	var resp sdk.ListResult[sdk.Currency]
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/config/convertible_shopper_currencies", nil, &resp, opts...)
	return &resp, err
}

// ReservePlan represents a reserve plan configuration.
// ReservePlan 表示准备金计划配置。
type ReservePlan struct {
	// Type is the reserve plan type. Required.
	// Type 准备金计划类型。必填。
	Type string `json:"type"`
	// Delayed is the delayed reserve plan. Optional.
	// Delayed 延迟准备金计划。可选。
	Delayed *ReservePlanDelayed `json:"delayed,omitempty"`
	// DelayedRolling is the delayed rolling reserve plan. Optional.
	// DelayedRolling 延迟滚动准备金计划。可选。
	DelayedRolling *ReservePlanDelayedRolling `json:"delayed_rolling,omitempty"`
	// Rolling is the rolling reserve plan. Optional.
	// Rolling 滚动准备金计划。可选。
	Rolling *ReservePlanRolling `json:"rolling,omitempty"`
}

// ReservePlanDelayed represents a delayed reserve plan.
// ReservePlanDelayed 延迟准备金计划。
type ReservePlanDelayed struct {
	// DelayInDays is the delay in days. Required.
	// DelayInDays 延迟天数。必填。
	DelayInDays int32 `json:"delay_in_days"`
}

// ReservePlanDelayedRolling represents a delayed rolling reserve plan.
// ReservePlanDelayedRolling 延迟滚动准备金计划。
type ReservePlanDelayedRolling struct {
	// DelayInDays is the delay in days. Required.
	// DelayInDays 延迟天数。必填。
	DelayInDays int32 `json:"delay_in_days"`
	// RollingPercentage is the rolling percentage. Required.
	// RollingPercentage 滚动百分比。必填。
	RollingPercentage float64 `json:"rolling_percentage"`
	// RollingWindowInDays is the rolling window in days. Required.
	// RollingWindowInDays 滚动窗口天数。必填。
	RollingWindowInDays int32 `json:"rolling_window_in_days"`
}

// ReservePlanRolling represents a rolling reserve plan.
// ReservePlanRolling 滚动准备金计划。
type ReservePlanRolling struct {
	// RollingPercentage is the rolling percentage. Required.
	// RollingPercentage 滚动百分比。必填。
	RollingPercentage float64 `json:"rolling_percentage"`
	// RollingWindowInDays is the rolling window in days. Required.
	// RollingWindowInDays 滚动窗口天数。必填。
	RollingWindowInDays int32 `json:"rolling_window_in_days"`
}

// GetReservePlan retrieves the current reserve plan.
// GetReservePlan 查询当前使用的准备金计划。
// 官方文档: https://www.airwallex.com/docs/api/payments/config/reserve_plan.md
func (s *Service) GetReservePlan(ctx context.Context, opts ...sdk.RequestOption) (*ReservePlan, error) {
	var resp ReservePlan
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/config/reserve_plan", nil, &resp, opts...)
	return &resp, err
}

// RegisteredDomains represents a list of registered domains.
// RegisteredDomains 表示注册的域名列表。
type RegisteredDomains struct {
	// Items is the list of items. Required.
	// Items 域名列表。必填。
	Items []string `json:"items"`
}

// GetApplePayDomains retrieves Apple Pay registered domains.
// GetApplePayDomains 查询 Apple Pay 注册域名。
// 官方文档: https://www.airwallex.com/docs/api/payments/config/registered_domains.md
func (s *Service) GetApplePayDomains(ctx context.Context, opts ...sdk.RequestOption) (*RegisteredDomains, error) {
	var resp RegisteredDomains
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/config/applepay/registered_domains", nil, &resp, opts...)
	return &resp, err
}

// AddApplePayDomainsRequest is the request to add Apple Pay domains.
// AddApplePayDomainsRequest 添加 Apple Pay 域名请求。
type AddApplePayDomainsRequest struct {
	// Items is the list of domains to add. Required.
	// Items 要添加的域名列表。必填。
	Items []string `json:"items"`
}

// AddApplePayDomains adds Apple Pay registered domains.
// AddApplePayDomains 添加 Apple Pay 注册域名。
// 官方文档: https://www.airwallex.com/docs/api/payments/config/add_items_registered_domains.md
func (s *Service) AddApplePayDomains(ctx context.Context, req *AddApplePayDomainsRequest, opts ...sdk.RequestOption) (*RegisteredDomains, error) {
	var resp RegisteredDomains
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/config/applepay/registered_domains/add_items", req, &resp, opts...)
	return &resp, err
}

// RemoveApplePayDomainsRequest is the request to remove Apple Pay domains.
// RemoveApplePayDomainsRequest 移除 Apple Pay 域名请求。
type RemoveApplePayDomainsRequest struct {
	// Items is the list of domains to remove. Required.
	// Items 要移除的域名列表。必填。
	Items []string `json:"items"`
	// Reason is the removal reason. Optional.
	// Reason 移除原因。可选。
	Reason string `json:"reason,omitempty"`
}

// RemoveApplePayDomains removes Apple Pay registered domains.
// RemoveApplePayDomains 移除 Apple Pay 注册域名。
// 官方文档: https://www.airwallex.com/docs/api/payments/config/remove_items_registered_domains.md
func (s *Service) RemoveApplePayDomains(ctx context.Context, req *RemoveApplePayDomainsRequest, opts ...sdk.RequestOption) (*RegisteredDomains, error) {
	var resp RegisteredDomains
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/config/applepay/registered_domains/remove_items", req, &resp, opts...)
	return &resp, err
}
