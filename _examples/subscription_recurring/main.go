package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/billing"
	"github.com/hakur/airwallex/pa"
	"github.com/hakur/airwallex/sdk"
)

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
	billingSvc := client.Billing()
	paSvc := client.PA()

	slog.Info("=== 包月/包年自动扣款完整流程 ===")

	// 1. 获取或创建产品（如：会员服务）
	slog.Info("1. 获取或创建产品...")
	var product *billing.Product
	prodList, err := billingSvc.ListProducts(ctx, &billing.ListProductsRequest{})
	if err != nil {
		if sdk.IsUnauthorized(err) {
			slog.Warn("跳过：需要 Billing 权限", "error", err)
		} else {
			slog.Error("列出产品失败", "error", err)
			os.Exit(1)
		}
	} else if len(prodList.Items) > 0 {
		product = &prodList.Items[0]
		slog.Info("使用已有产品", "id", product.ID, "name", product.Name)
	} else {
		product, err = billingSvc.CreateProduct(ctx, &billing.CreateProductRequest{
			RequestID:   "prod-req-" + time.Now().Format("20060102150405"),
			Name:        "高级会员服务",
			Description: "包含所有高级功能的月度订阅",
		})
		if err != nil {
			if sdk.IsUnauthorized(err) {
				slog.Warn("跳过：需要 Billing 权限", "error", err)
			} else if sdk.IsDuplicateRequest(err) {
				prodList, err = billingSvc.ListProducts(ctx, &billing.ListProductsRequest{})
				if err != nil {
					slog.Error("列出产品失败", "error", err)
					os.Exit(1)
				}
				if len(prodList.Items) > 0 {
					product = &prodList.Items[0]
					slog.Info("使用已有产品（重复请求后）", "id", product.ID, "name", product.Name)
				}
			} else {
				slog.Error("创建产品失败", "error", err)
				os.Exit(1)
			}
		} else {
			slog.Info("产品已创建", "id", product.ID, "name", product.Name)
		}
	}

	// 2. 创建价格计划（月费 $29.99）
	var price *billing.Price
	if product != nil && product.ID != "" {
		slog.Info("2. 创建价格计划（月费 $29.99）...")
		price, err = billingSvc.CreatePrice(ctx, &billing.CreatePriceRequest{
			RequestID:  "price-req-" + time.Now().Format("20060102150405"),
			ProductID:  product.ID,
			Currency:   sdk.CurrencyUSD,
			UnitAmount: 29.99,
			Recurring: &billing.Recurring{
				Period:     1,
				PeriodUnit: billing.PeriodUnitMonth,
			},
		})
		if err != nil {
			if sdk.IsUnauthorized(err) {
				slog.Warn("跳过：需要 Billing 权限", "error", err)
			} else {
				slog.Error("创建价格计划失败", "error", err)
				os.Exit(1)
			}
		} else {
			slog.Info("价格计划已创建", "id", price.ID, "amount", price.UnitAmount, "currency", price.Currency)
		}
	}

	// 3. 创建支付客户
	slog.Info("3. 创建支付客户...")
	customer, err := paSvc.CreateCustomer(ctx, &pa.CreateCustomerRequest{
		RequestID:          "cus-req-" + time.Now().Format("20060102150405"),
		MerchantCustomerID: "merchant-cus-" + time.Now().Format("20060102150405"),
		Email:              "user@example.com",
		FirstName:          "张",
		LastName:           "三",
	})
	if err != nil {
		slog.Error("创建客户失败", "error", err)
		os.Exit(1)
	}
	slog.Info("客户已创建", "id", customer.ID, "email", customer.Email)

	// 4. 创建支付授权（自动扣款授权）
	// 注意：真实场景中，PaymentMethodID 应由前端 SDK 收集用户卡信息后提供
	// 这里使用占位符演示流程，实际运行时需要替换为真实的 PaymentMethod ID
	slog.Info("4. 创建支付授权（授权自动扣款）...")
	slog.Info("注意：请先通过前端 SDK 收集支付方式，或使用 PaymentIntent 保存支付方式")
	slog.Info("然后替换下方 pm_xxx 为真实的 PaymentMethod ID")

	var consent *pa.PaymentConsent
	// consent, err = paSvc.CreatePaymentConsent(ctx, &pa.CreatePaymentConsentRequest{
	// 	CustomerID:      customer.ID,
	// 	PaymentMethodID: "pm_xxx", // 替换为真实的 PaymentMethod ID
	// })
	// if err != nil {
	// 	slog.Error("创建支付授权失败", "error", err)
	// 	os.Exit(1)
	// }
	// slog.Info("支付授权已创建", "id", consent.ID, "status", consent.Status)
	slog.Info("[演示模式] 跳过实际创建，请替换 pm_xxx 后取消注释")

	// 5. 创建订阅（签约包月）
	slog.Info("5. 创建订阅（签约包月服务）...")
	var subscription *billing.Subscription
	if price != nil && price.ID != "" && customer != nil && customer.ID != "" {
		subscription, err = billingSvc.CreateSubscription(ctx, &billing.CreateSubscriptionRequest{
			RequestID:         "sub-req-" + time.Now().Format("20060102150405"),
			BillingCustomerID: customer.ID,
			CollectionMethod:  billing.CollectionMethodChargeOnCheckout,
			Items: []billing.SubscriptionItemInput{
				{PriceID: price.ID, Quantity: 1},
			},
		})
		if err != nil {
			if sdk.IsUnauthorized(err) {
				slog.Warn("跳过：需要 Billing 权限", "error", err)
			} else {
				slog.Error("创建订阅失败", "error", err)
				os.Exit(1)
			}
		} else {
			slog.Info("订阅已创建", "id", subscription.ID, "status", subscription.Status)
		}
	} else {
		slog.Warn("跳过：缺少价格计划或客户")
	}

	// 6. 查询订阅状态
	if subscription != nil && subscription.ID != "" {
		slog.Info("6. 查询订阅状态...")
		fetched, err := billingSvc.GetSubscription(ctx, subscription.ID)
		if err != nil {
			slog.Error("查询订阅失败", "error", err)
			os.Exit(1)
		}
		slog.Info("订阅状态", "status", fetched.Status, "customer", fetched.BillingCustomerID)
	}

	// 7. 列出所有订阅
	slog.Info("7. 列出所有订阅...")
	subList, err := billingSvc.ListSubscriptions(ctx, &billing.ListSubscriptionsRequest{})
	if err != nil {
		if sdk.IsUnauthorized(err) {
			slog.Warn("跳过：需要 Billing 权限", "error", err)
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

	// 8. 取消订阅（停止包月扣款）
	if subscription != nil && subscription.ID != "" {
		slog.Info("8. 取消订阅...")
		cancelled, err := billingSvc.CancelSubscription(ctx, subscription.ID, &billing.CancelSubscriptionRequest{
			RequestID:         "sub-cancel-" + time.Now().Format("20060102150405"),
			ProrationBehavior: billing.ProrationBehaviorNone,
		})
		if err != nil {
			if sdk.IsUnauthorized(err) {
				slog.Warn("跳过：需要 Billing 权限", "error", err)
			} else {
				slog.Error("取消订阅失败", "error", err)
				os.Exit(1)
			}
		} else {
			slog.Info("订阅已取消", "id", cancelled.ID, "status", cancelled.Status)
		}
	}

	// 9. 取消支付授权（解除自动扣款绑定）
	if consent != nil && consent.ID != "" {
		slog.Info("9. 取消支付授权...")
		cancelledConsent, err := paSvc.CancelPaymentConsent(ctx, consent.ID, &pa.CancelPaymentConsentRequest{
			RequestID: "pc-cancel-" + time.Now().Format("20060102150405"),
		})
		if err != nil {
			slog.Error("取消支付授权失败", "error", err)
			os.Exit(1)
		}
		slog.Info("支付授权已取消", "id", cancelledConsent.ID, "status", cancelledConsent.Status)
	}

	slog.Info("=== 流程完成 ===")
	slog.Info("说明：")
	slog.Info("- 真实场景中，步骤 4 的 PaymentMethodID 需要由前端 SDK 提供")
	slog.Info("- 客户在前端输入卡信息后，前端调用 Airwallex SDK 生成 PaymentMethod")
	slog.Info("- 后端收到 PaymentMethod ID 后，创建 PaymentConsent 授权自动扣款")
	slog.Info("- 创建 Subscription 后，Airwallex 会按照 Price 配置的周期自动扣款")
	slog.Info("- 取消 Subscription 停止后续扣款，但历史已扣款项不受影响")
	slog.Info("- 取消 PaymentConsent 解除支付方式绑定，彻底关闭自动扣款能力")
}
