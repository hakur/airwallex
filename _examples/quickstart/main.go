package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/sdk"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	// 从项目根目录的 .env 加载密钥
	client, err := airwallex.NewFromEnv("../../.env", sdk.WithBaseURL(sdk.SandboxURL))
	if err != nil {
		slog.Error("创建客户端失败", "error", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 查询余额
	balances, err := client.Core().GetCurrentBalances(ctx)
	if err != nil {
		slog.Error("查询余额失败", "error", err)
		os.Exit(1)
	}

	slog.Info("当前余额")
	for _, b := range balances {
		slog.Info("余额详情",
			"currency", b.Currency,
			"available", b.AvailableAmount,
			"total", b.TotalAmount,
		)
	}
}
