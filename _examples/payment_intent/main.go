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
	// 修改此变量切换支付方式
	paymentMethod := "card"

	// 创建 PaymentIntent
	intent, err := svc.CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
		RequestID:       "pi-req-" + time.Now().Format("20060102150405"),
		MerchantOrderID: "order-" + time.Now().Format("20060102150405"),
		Amount:          100.0,
		Currency:        sdk.CurrencyUSD,
	})
	if err != nil {
		slog.Error("创建 PaymentIntent 失败", "error", err)
		return
	}
	slog.Info("创建 PaymentIntent", "id", intent.ID, "status", intent.Status)

	// 获取 PaymentIntent
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

	// 根据支付方式构建 PaymentMethodInput
	var pmInput *pa.PaymentMethodInput
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

	// 确认 PaymentIntent
	confirmed, err := svc.ConfirmPaymentIntent(ctx, intent.ID, &pa.ConfirmPaymentIntentRequest{
		RequestID:     "pi-confirm-" + time.Now().Format("20060102150405"),
		PaymentMethod: pmInput,
	})
	if err != nil {
		if sdk.IsForbidden(err) {
			slog.Warn("确认 PaymentIntent 权限不足", "error", err)
		} else if sdk.IsResourceNotFound(err) {
			slog.Warn("PaymentIntent 不存在", "id", intent.ID)
		} else if sdk.IsInvalidStatusForOperation(err) {
			slog.Warn("PaymentIntent 状态不允许确认", "id", intent.ID, "status", intent.Status)
		} else {
			slog.Error("确认 PaymentIntent 失败", "error", err)
		}
		return
	}
	slog.Info("确认 PaymentIntent", "id", confirmed.ID, "status", confirmed.Status)

	// 根据状态处理后续流程
	switch confirmed.Status {
	case pa.PaymentIntentStatusSucceeded:
		slog.Info("支付已成功完成")

	case pa.PaymentIntentStatusRequiresCapture:
		// card 支付成功授权，需要手动捕获
		captured, err := svc.CapturePaymentIntent(ctx, intent.ID, &pa.CapturePaymentIntentRequest{
			RequestID: "pi-capture-" + time.Now().Format("20060102150405"),
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
		// wechatpay / alipay 需要用户扫码完成支付
		redirectURL, ok := confirmed.NextAction["url"].(string)
		if ok && redirectURL != "" {
			slog.Info("请完成支付，正在打开浏览器", "url", redirectURL)
			if err := openBrowser(redirectURL); err != nil {
				slog.Warn("自动打开浏览器失败，请手动访问", "url", redirectURL, "error", err)
			}
		} else {
			slog.Info("支付需要用户操作", "next_action", confirmed.NextAction)
		}

	default:
		slog.Info("支付状态，请检查 PaymentIntent 详情", "status", confirmed.Status)
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
