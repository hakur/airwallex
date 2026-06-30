package main

import (
	"context"
	"fmt"
	"log/slog"
	"os/exec"
	"runtime"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/pa"
	"github.com/hakur/airwallex/sdk"
)

func main() {
	client, err := airwallex.NewFromEnv("../../.env",
		sdk.WithBaseURL(sdk.SandboxURL),
	)
	if err != nil {
		slog.Error("创建客户端失败", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	svc := client.PA()

	// 支持多种支付方式: "card" | "wechatpay" | "alipay" | "airwallex_pay" | "kakaopay" | "visa"
	// 修改此变量切换支付方式。card 走 pre_auth 需手动捕获，wechatpay/alipay 走跳转链接
	paymentMethod := "card"

	var pmInput *pa.PaymentMethodInput
	var cardOpts *pa.CardPaymentMethodOptions
	var pmOpts *pa.PaymentMethodOptionsRequest
	switch paymentMethod {
	case "card":
		pmInput = &pa.PaymentMethodInput{
			Type: "card",
			Card: &pa.CardPaymentMethod{
				Number:      "4012000300001003",
				ExpiryMonth: "03",
				ExpiryYear:  "2030",
				CVC:         "123",
				Name:        "Test User",
			},
		}
		cardOpts = &pa.CardPaymentMethodOptions{
			AuthorizationType: pa.CardAuthPreAuth,
			AutoCapture:       new(false),
			ThreeDSAction:     pa.ThreeDSForce,
		}
		pmOpts = &pa.PaymentMethodOptionsRequest{Card: cardOpts}
		// card 支付支持 pre_auth 模式
		cardOpts = &pa.CardPaymentMethodOptions{
			AuthorizationType: pa.CardAuthPreAuth,
			AutoCapture:       new(false),
			ThreeDSAction:     pa.ThreeDSForce,
		}
	case "wechatpay":
		pmInput = &pa.PaymentMethodInput{
			Type:      "wechatpay",
			WechatPay: &pa.WechatPayPaymentMethod{Flow: "mobile_web"},
		}
	case "alipay":
		pmInput = &pa.PaymentMethodInput{
			Type:     "alipaycn",
			AlipayCN: &pa.AlipayPaymentMethod{Flow: "mobile_web", OSType: "ios"},
		}
	case "airwallex_pay":
		pmInput = &pa.PaymentMethodInput{
			Type:         "airwallex_pay",
			AirwallexPay: &pa.AirwallexPayPaymentMethod{PayerName: "Test User"},
		}
	case "kakaopay":
		pmInput = &pa.PaymentMethodInput{
			Type:     "kakaopay",
			KakaoPay: &pa.KakaoPayPaymentMethod{Flow: "mobile_web"},
		}
	case "visa":
		pmInput = &pa.PaymentMethodInput{
			Type: "visa",
			Visa: &pa.VisaPaymentMethod{},
		}
	default:
		slog.Error("不支持的支付方式", "method", paymentMethod)
		return
	}

	// 一步创建 PaymentIntent — 直接带支付方式，一步获取状态和跳转链接
	slog.Info("=== 1. 创建 PaymentIntent（一步法）===")
	intent, err := svc.CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
		RequestID:            "pi-req-" + time.Now().Format("20060102150405"),
		MerchantOrderID:      "order-" + time.Now().Format("20060102150405"),
		Amount:               100.0,
		Currency:             sdk.CurrencyUSD,
		OrderType:            "retail",
		ReturnURL:            "https://example.com/payment-callback",
		PaymentMethod:        pmInput,
		PaymentMethodOptions: pmOpts,
		Order: &pa.CreateOrderRequest{
			Cancellable:     true,
			CreatedAt:       "2026-06-29T12:00:00Z",
			PrepaymentModel: pa.PrepaymentModelFull,
		},
		Shipping: &pa.CreateShippingRequest{
			PhoneNumber:    "+1234567890",
			ShippingMethod: "express",
		},
		Billing: &pa.CreateBillingRequest{
			FirstName: "Demo",
			LastName:  "User",
			Address:   &pa.CreateAddressRequest{CountryCode: "US", City: "San Francisco"},
		},
	})
	if err != nil {
		slog.Error("创建 PaymentIntent 失败", "error", err)
		return
	}
	slog.Info("创建 PaymentIntent", "id", intent.ID, "status", intent.Status)

	// 获取 PaymentIntent
	slog.Info("=== 2. 获取 PaymentIntent ===")
	fetched, err := svc.GetPaymentIntent(ctx, intent.ID)
	if err != nil {
		if sdk.IsResourceNotFound(err) {
			slog.Warn("PaymentIntent 未找到", "id", intent.ID)
		} else {
			slog.Error("获取 PaymentIntent 失败", "error", err)
		}
		return
	}
	slog.Info("获取 PaymentIntent", "id", fetched.ID, "amount", fetched.Amount, "currency", fetched.Currency)

	// 根据状态处理后续流程
	slog.Info("=== 3. 处理状态 ===")
	switch intent.Status {
	case pa.PaymentIntentStatusSucceeded:
		slog.Info("支付已成功完成")

	case pa.PaymentIntentStatusRequiresCapture:
		// card 预授权成功，需要手动捕获
		captured, err := svc.CapturePaymentIntent(ctx, intent.ID, &pa.CapturePaymentIntentRequest{
			RequestID: "pi-capture-" + time.Now().Format("20060102150405"),
			Metadata:  map[string]any{"source": "demo"},
		})
		if err != nil {
			if sdk.IsInvalidStatusForOperation(err) {
				slog.Warn("PaymentIntent 状态不允许捕获", "id", intent.ID)
			} else {
				slog.Error("捕获 PaymentIntent 失败", "error", err)
			}
			return
		}
		slog.Info("捕获 PaymentIntent", "id", captured.ID, "status", captured.Status)

	case pa.PaymentIntentStatusRequiresCustomerAction:
		// wechatpay / alipay 需要用户扫码/跳转完成支付
		if intent.NextAction != nil && intent.NextAction.URL != "" {
			slog.Info("请完成支付，正在打开浏览器", "url", intent.NextAction.URL)
			if err := openBrowser(intent.NextAction.URL); err != nil {
				slog.Warn("自动打开浏览器失败，请手动访问", "url", intent.NextAction.URL, "error", err)
			}
		} else if intent.NextAction != nil && intent.NextAction.QRCode != "" {
			slog.Info("请扫码完成支付", "qrcode", intent.NextAction.QRCode)
		} else {
			slog.Info("支付需要用户操作", "next_action", intent.NextAction)
		}

	default:
		slog.Info("支付状态，请检查 PaymentIntent 详情", "status", intent.Status)
	}
}

// openBrowser 使用系统默认浏览器打开指定 URL。
func openBrowser(url string) error {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	case "linux":
		return exec.Command("xdg-open", url).Start()
	default:
		return fmt.Errorf("不支持的操作系统: %s", runtime.GOOS)
	}
}
