package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/pa"
	"github.com/hakur/airwallex/simulation"
	"github.com/hakur/airwallex/sdk"
)

func main() {
	client, err := airwallex.NewFromEnv("../../.env", sdk.WithBaseURL(sdk.SandboxURL))
	if err != nil {
		slog.Error("创建客户端失败", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	svc := client.Simulation()

	// 先创建一个 PaymentIntent 用于模拟
	slog.Info("=== 创建 PaymentIntent 用于模拟 ===")
	intent, err := client.PA().CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
		RequestID:       "sim-pi-req-" + time.Now().Format("20060102150405"),
		MerchantOrderID: "sim-order-" + time.Now().Format("20060102150405"),
		Amount:          100.0,
		Currency:        sdk.CurrencyUSD,
	})
	if err != nil {
		slog.Warn("创建 PaymentIntent 失败", "error", err)
		return
	}
	slog.Info("创建 PaymentIntent", "id", intent.ID, "status", intent.Status)

	// 模拟 Shopper Action（如支付成功、3DS 验证等）
	slog.Info("=== 模拟 Shopper Action ===")
	err = svc.SimulateShopperAction(ctx, "pay", &simulation.SimulateShopperActionRequest{
		URL: "https://checkout.airwallex.com/simulate?payment_intent_id=" + intent.ID,
	})
	if err != nil {
		slog.Warn("模拟 Shopper Action 失败", "error", err)
	} else {
		slog.Info("Shopper Action 模拟成功")
	}

	// 验证 PaymentIntent 状态已变更
	slog.Info("=== 验证 PaymentIntent 状态 ===")
	updated, err := client.PA().GetPaymentIntent(ctx, intent.ID)
	if err != nil {
		slog.Warn("查询 PaymentIntent 失败", "error", err)
		return
	}
	slog.Info("PaymentIntent 当前状态", "status", updated.Status)

	// 模拟全球账户存款
	slog.Info("=== 模拟全球账户存款 ===")
	deposit, err := svc.SimulateGlobalAccountDeposit(ctx, &simulation.SimulateGlobalAccountDepositRequest{
		GlobalAccountID: "ga-" + time.Now().Format("20060102150405"),
		Amount:          1000.0,
		PayerBankName:   "Test Bank",
		PayerCountry:    "US",
		PayerName:       "Test Payer",
		Reference:       "sim-deposit-" + time.Now().Format("20060102150405"),
		StatementRef:    "stmt-" + time.Now().Format("20060102150405"),
		Status:          "COMPLETED",
	})
	if err != nil {
		slog.Warn("模拟存款失败", "error", err)
		return
	}
	slog.Info("存款模拟结果", "id", deposit.ID, "amount", deposit.Amount, "currency", deposit.Currency)
}
