package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/fx"
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
	svc := client.FX()

	// 查询汇率
	slog.Info("=== 查询 USD/EUR 汇率 ===")
	rate, err := svc.GetRates(ctx, &fx.GetRatesRequest{
		BuyCurrency:  sdk.CurrencyUSD,
		SellCurrency: sdk.CurrencyEUR,
	})
	if err != nil {
		slog.Warn("查询汇率失败", "error", err)
	} else {
		slog.Info("获取汇率", "sell", rate.SellCurrency, "rate", rate.Rate, "buy", rate.BuyCurrency)
	}

	// 创建外汇兑换（用 EUR 买 USD）
	slog.Info("=== 创建外汇兑换 ===")
	conversion, err := svc.CreateConversion(ctx, &fx.CreateConversionRequest{
		RequestID:    "fx-req-" + time.Now().Format("20060102150405"),
		BuyCurrency:  sdk.CurrencyUSD,
		SellCurrency: sdk.CurrencyEUR,
		BuyAmount:    "1000",
	})
	if err != nil {
		slog.Warn("创建外汇兑换失败", "error", err)
		return
	}
	slog.Info("创建外汇兑换", "id", conversion.ConversionID)
	slog.Info("  买入", "amount", conversion.BuyAmount, "currency", conversion.BuyCurrency)
	slog.Info("  卖出", "amount", conversion.SellAmount, "currency", conversion.SellCurrency)
	slog.Info("  状态", "status", conversion.Status)

	// 查询兑换详情
	slog.Info("=== 查询外汇兑换详情 ===")
	fetched, err := svc.GetConversion(ctx, conversion.ConversionID)
	if err != nil {
		slog.Warn("查询外汇兑换详情失败", "error", err)
		return
	}
	slog.Info("兑换详情", "id", fetched.ConversionID, "status", fetched.Status)

	// 列出历史兑换
	slog.Info("=== 列出历史外汇兑换 ===")
	list, err := svc.ListConversions(ctx)
	if err != nil {
		slog.Warn("列出历史外汇兑换失败", "error", err)
		return
	}
	slog.Info("历史兑换记录", "count", len(list.Items))
	for _, c := range list.Items {
		slog.Info("  兑换记录",
			"id", c.ConversionID,
			"sell", c.SellAmount,
			"sell_currency", c.SellCurrency,
			"buy", c.BuyAmount,
			"buy_currency", c.BuyCurrency,
			"status", c.Status,
		)
	}
}
