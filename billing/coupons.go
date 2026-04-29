package billing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// CouponDuration 优惠券有效期。
type CouponDuration struct {
	// Period 周期数。必填。
	Period int32 `json:"period"`
	// PeriodUnit 周期单位。必填。
	PeriodUnit PeriodUnit `json:"period_unit"`
}

// Coupon represents a coupon object.
// Coupon 优惠券对象。
type Coupon struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Name 名称。必填。
	Name string `json:"name"`
	// Active 是否活跃。必填。
	Active bool `json:"active"`
	// DiscountModel 折扣模型。必填。
	DiscountModel DiscountModel `json:"discount_model"`
	// AmountOff 固定金额减免。可选。
	AmountOff float64 `json:"amount_off,omitempty"`
	// PercentageOff 百分比减免。可选。
	PercentageOff float64 `json:"percentage_off,omitempty"`
	// Currency 货币代码。可选。
	Currency string `json:"currency,omitempty"`
	// DurationType 持续时间类型。必填。
	DurationType DiscountDurationType `json:"duration_type"`
	// Duration 有效期。可选。
	Duration *CouponDuration `json:"duration,omitempty"`
	// ExpiresAt 过期时间。可选。
	ExpiresAt string `json:"expires_at,omitempty"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// TimesRedeemed 已兑换次数。必填。
	TimesRedeemed int32 `json:"times_redeemed"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// UpdatedAt 更新时间。必填。
	UpdatedAt string `json:"updated_at"`
}

// CreateCouponRequest 创建优惠券请求。
type CreateCouponRequest struct {
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Name 名称。必填。
	Name string `json:"name"`
	// DiscountModel 折扣模型。必填。
	DiscountModel DiscountModel `json:"discount_model"`
	// AmountOff 固定金额减免。可选。
	AmountOff float64 `json:"amount_off,omitempty"`
	// PercentageOff 百分比减免。可选。
	PercentageOff float64 `json:"percentage_off,omitempty"`
	// Currency 货币代码。可选。
	Currency string `json:"currency,omitempty"`
	// DurationType 持续时间类型。必填。
	DurationType DiscountDurationType `json:"duration_type"`
	// Duration 有效期。可选。
	Duration *CouponDuration `json:"duration,omitempty"`
	// ExpiresAt 过期时间。可选。
	ExpiresAt string `json:"expires_at,omitempty"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
}

// UpdateCouponRequest 更新优惠券请求。
type UpdateCouponRequest struct {
	// Active 是否活跃。可选。
	Active bool `json:"active,omitempty"`
	// Name 名称。可选。
	Name string `json:"name,omitempty"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// ExpiresAt 过期时间。可选。
	ExpiresAt string `json:"expires_at,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// AmountOff 固定金额减免。可选。
	AmountOff float64 `json:"amount_off,omitempty"`
	// PercentageOff 百分比减免。可选。
	PercentageOff float64 `json:"percentage_off,omitempty"`
	// Currency 货币代码。可选。
	Currency string `json:"currency,omitempty"`
	// DiscountModel 折扣模型。可选。
	DiscountModel DiscountModel `json:"discount_model,omitempty"`
	// DurationType 持续时间类型。可选。
	DurationType DiscountDurationType `json:"duration_type,omitempty"`
	// Duration 有效期。可选。
	Duration *CouponDuration `json:"duration,omitempty"`
}

// ListCouponsRequest 列出优惠券请求。
type ListCouponsRequest struct {
	// DiscountModel 折扣模型。可选。
	DiscountModel string `json:"discount_model,omitempty"`
	// Active 是否活跃。可选。
	Active bool `json:"active,omitempty"`
	// DurationType 持续时间类型。可选。
	DurationType string `json:"duration_type,omitempty"`
	// FromExpiresAt 过期时间起始。可选。
	FromExpiresAt string `json:"from_expires_at,omitempty"`
	// ToExpiresAt 过期时间截止。可选。
	ToExpiresAt string `json:"to_expires_at,omitempty"`
	// FromCreatedAt 创建时间起始。可选。
	FromCreatedAt string `json:"from_created_at,omitempty"`
	// ToCreatedAt 创建时间截止。可选。
	ToCreatedAt string `json:"to_created_at,omitempty"`
	// Page 分页游标。可选。
	Page string `json:"page,omitempty"`
	// PageSize 每页数量。可选。
	PageSize int32 `json:"page_size,omitempty"`
}

// CreateCoupon creates a new coupon.
// CreateCoupon 创建优惠券。
// 官方文档: https://www.airwallex.com/docs/api/billing/coupons/create.md
func (s *Service) CreateCoupon(ctx context.Context, req *CreateCouponRequest, opts ...sdk.RequestOption) (*Coupon, error) {
	var resp Coupon
	err := s.doer.Do(ctx, "POST", "/api/v1/coupons/create", req, &resp, opts...)
	return &resp, err
}

// GetCoupon retrieves a coupon by ID.
// GetCoupon 根据 ID 获取优惠券。
// 官方文档: https://www.airwallex.com/docs/api/billing/coupons/retrieve.md
func (s *Service) GetCoupon(ctx context.Context, id string, opts ...sdk.RequestOption) (*Coupon, error) {
	var resp Coupon
	err := s.doer.Do(ctx, "GET", "/api/v1/coupons/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateCoupon updates an existing coupon.
// UpdateCoupon 更新优惠券。
// 官方文档: https://www.airwallex.com/docs/api/billing/coupons/update.md
func (s *Service) UpdateCoupon(ctx context.Context, id string, req *UpdateCouponRequest, opts ...sdk.RequestOption) (*Coupon, error) {
	var resp Coupon
	err := s.doer.Do(ctx, "POST", "/api/v1/coupons/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ListCoupons lists coupons with optional filters.
// ListCoupons 列出优惠券。
// 官方文档: https://www.airwallex.com/docs/api/billing/coupons/list.md
func (s *Service) ListCoupons(ctx context.Context, req *ListCouponsRequest, opts ...sdk.RequestOption) (*ListResult[Coupon], error) {
	var resp ListResult[Coupon]
	err := s.doer.Do(ctx, "GET", "/api/v1/coupons", req, &resp, opts...)
	return &resp, err
}
