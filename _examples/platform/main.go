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

	// 平台商户使用自己的密钥初始化
	client, err := airwallex.NewFromEnv("../../.env", sdk.WithBaseURL(sdk.SandboxURL), sdk.WithDebug(true))
	if err != nil {
		slog.Error("创建客户端失败", "error", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 查询平台商户自身余额
	balances, err := client.Core().GetCurrentBalances(ctx)
	if err != nil {
		slog.Error("查询余额失败", "error", err)
		os.Exit(1)
	}
	slog.Info("平台商户余额")
	for _, b := range balances {
		slog.Info("余额", "currency", b.Currency, "available", b.AvailableAmount)
	}

	// 代理子商户查询余额（假设子商户 account_id 为 acct_123）
	subAccountID := "acct_123"
	subBalances, err := client.Core().GetCurrentBalances(ctx, sdk.WithRequestOnBehalfOf(subAccountID))
	if err != nil {
		if sdk.IsUnauthorized(err) {
			slog.Warn("无权代理子商户（需要平台权限或子商户未授权）", "subAccountID", subAccountID)
		} else {
			slog.Error("查询子商户余额失败", "subAccountID", subAccountID, "error", err)
		}
	} else {
		slog.Info("子商户余额", "subAccountID", subAccountID)
		for _, b := range subBalances {
			slog.Info("子商户余额详情", "currency", b.Currency, "available", b.AvailableAmount)
		}
	}
}
