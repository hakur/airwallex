package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/billing"
	"github.com/hakur/airwallex/sdk"
)

// 包年订阅 + 首年优惠示例
// 演示：创建年度订阅计划，使用优惠券享受首年折扣

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	client, err := airwallex.NewFromEnv("../../.env",
		sdk.WithBaseURL(sdk.SandboxURL),
		sdk.WithDebug(true),
	)
	if err != nil {
		slog.Error("创建客户端失败", "error", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	svc := client.Billing()

	slog.Info("=== 包年订阅 + 首年优惠券 ===")

	// 1. 创建计费客户
	slog.Info("1. 创建计费客户...")
	customer, err := svc.CreateBillingCustomer(ctx, &billing.CreateBillingCustomerRequest{
		RequestID: "ycus-" + time.Now().Format("20060102150405"),
		Email:     "yearly-vip@example.com",
		Name:      "包年 VIP 客户",
		Type:      billing.CustomerTypeIndividual,
	})
	if err != nil {
		if sdk.IsUnauthorized(err) {
			slog.Warn("需要 Billing 权限，跳过", "error", err)
		} else {
			slog.Error("创建计费客户失败", "error", err)
			os.Exit(1)
		}
	} else {
		slog.Info("计费客户已创建", "id", customer.ID, "email", customer.Email)
	}

	// 2. 创建产品
	slog.Info("2. 创建年度会员产品...")
	product, err := svc.CreateProduct(ctx, &billing.CreateProductRequest{
		RequestID:   "yprod-" + time.Now().Format("20060102150405"),
		Name:        "年度 VIP 会员",
		Description: "包年订阅，享全年会员权益",
	})
	if err != nil {
		if sdk.IsUnauthorized(err) {
			slog.Warn("需要 Billing 权限，跳过", "error", err)
		} else {
			slog.Error("创建产品失败", "error", err)
			os.Exit(1)
		}
	} else {
		slog.Info("产品已创建", "id", product.ID, "name", product.Name)
	}

	// 3. 创建年度价格（$299/年）
	var price *billing.Price
	if product != nil && product.ID != "" {
		slog.Info("3. 创建年度价格（$299/年）...")
		price, err = svc.CreatePrice(ctx, &billing.CreatePriceRequest{
			RequestID:  "yprice-" + time.Now().Format("20060102150405"),
			ProductID:  product.ID,
			Currency:   sdk.CurrencyUSD,
			UnitAmount: 299.00,
			Recurring: &billing.Recurring{
				Period:     1,
				PeriodUnit: billing.PeriodUnitYear,
			},
		})
		if err != nil {
			if sdk.IsUnauthorized(err) {
				slog.Warn("需要 Billing 权限，跳过", "error", err)
			} else if sdk.IsValidationError(err) {
				slog.Warn("价格验证失败", "error", err)
			} else {
				slog.Error("创建价格失败", "error", err)
				os.Exit(1)
			}
		} else {
			slog.Info("年度价格已创建",
				"id", price.ID,
				"amount", price.UnitAmount,
				"currency", price.Currency,
				"period", "1YEAR",
			)
		}
	}

	// 4. 创建首年优惠券（20% off）
	var coupon *billing.Coupon
	slog.Info("4. 创建首年 20% 优惠券...")
	coupon, err = svc.CreateCoupon(ctx, &billing.CreateCouponRequest{
		RequestID:     "ycoupon-" + time.Now().Format("20060102150405"),
		Name:          "首年 20% 优惠",
		DiscountModel: billing.DiscountModelPercentage,
		PercentageOff: 20,
		DurationType:  billing.DiscountDurationTypeOnce,
	})
	if err != nil {
		if sdk.IsUnauthorized(err) {
			slog.Warn("需要 Billing 权限，跳过", "error", err)
		} else {
			slog.Error("创建优惠券失败", "error", err)
			os.Exit(1)
		}
	} else {
		slog.Info("优惠券已创建",
			"id", coupon.ID,
			"name", coupon.Name,
			"model", coupon.DiscountModel,
			"percent_off", coupon.PercentageOff,
		)
	}

	// 5. 创建订阅（应用优惠券）
	var subscription *billing.Subscription
	if price != nil && price.ID != "" && customer != nil && customer.ID != "" && coupon != nil && coupon.ID != "" {
		slog.Info("5. 创建订阅（应用首年优惠）...")
		subscription, err = svc.CreateSubscription(ctx, &billing.CreateSubscriptionRequest{
			RequestID:         "ysub-" + time.Now().Format("20060102150405"),
			BillingCustomerID: customer.ID,
			CollectionMethod:  billing.CollectionMethodChargeOnCheckout,
			Items: []billing.SubscriptionItemInput{
				{
					PriceID:  price.ID,
					Quantity: 1,
				},
			},
			Discounts: []billing.DiscountInput{
				{
					Coupon: &billing.CouponRef{ID: coupon.ID},
					Type:   billing.DiscountTypeCoupon,
				},
			},
		})
		if err != nil {
			if sdk.IsUnauthorized(err) {
				slog.Warn("需要 Billing 权限，跳过", "error", err)
			} else {
				slog.Error("创建订阅失败", "error", err)
				os.Exit(1)
			}
		} else {
			slog.Info("订阅已创建",
				"id", subscription.ID,
				"status", subscription.Status,
				"currency", subscription.Currency,
			)
		}
	}

	// 6. 验证订阅详情
	if subscription != nil && subscription.ID != "" {
		slog.Info("6. 查询订阅详情...")
		fetched, err := svc.GetSubscription(ctx, subscription.ID)
		if err != nil {
			slog.Error("查询订阅失败", "error", err)
		} else {
			slog.Info("订阅详情",
				"id", fetched.ID,
				"status", fetched.Status,
				"discounts_count", len(fetched.AppliedDiscounts),
			)
			for _, d := range fetched.AppliedDiscounts {
				slog.Info("已应用折扣",
					"model", d.DiscountModel,
					"percent_off", d.PercentageOff,
					"duration_type", d.DurationType,
				)
			}
		}
	}

	// 7. 列出活跃订阅
	slog.Info("7. 列出活跃订阅...")
	subList, err := svc.ListSubscriptions(ctx, &billing.ListSubscriptionsRequest{PageSize: 10})
	if err != nil {
		if sdk.IsUnauthorized(err) {
			slog.Warn("需要 Billing 权限，跳过", "error", err)
		} else {
			slog.Error("列出订阅失败", "error", err)
			os.Exit(1)
		}
	} else {
		slog.Info("当前订阅数", "count", len(subList.Items))
		for _, s := range subList.Items {
			slog.Info("订阅", "id", s.ID, "status", s.Status, "customer", s.BillingCustomerID)
		}
	}

	// 8. 取消订阅（清洁沙箱数据）
	if subscription != nil && subscription.ID != "" {
		slog.Info("8. 取消订阅（清理沙箱数据）...")
		time.Sleep(200 * time.Millisecond) // 避免 "already updated" 竞态
		cancelled, err := svc.CancelSubscription(ctx, subscription.ID, &billing.CancelSubscriptionRequest{
			RequestID:         "ysub-cancel-" + time.Now().Format("20060102150405"),
			ProrationBehavior: billing.ProrationBehaviorNone,
		})
		if err != nil {
			if sdk.IsUnauthorized(err) || sdk.IsValidationError(err) {
				slog.Warn("取消订阅跳过", "error", err)
			} else {
				slog.Error("取消订阅失败", "error", err)
				os.Exit(1)
			}
		} else {
			slog.Info("订阅已取消（沙箱数据已清理）", "id", cancelled.ID, "status", cancelled.Status)
		}
	}

	slog.Info("=== 流程完成 ===")
	slog.Info("说明：")
	slog.Info("- 年度订阅按年计费，首年使用优惠券享受 20% 折扣")
	slog.Info("- 优惠券 DurationType=ONCE 表示仅首年生效，次年恢复原价")
	slog.Info("- 步骤 8 取消订阅是为了清理沙箱测试数据，生产环境请勿取消")
}
