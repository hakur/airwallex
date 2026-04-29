package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/risk"
	"github.com/hakur/airwallex/sdk"
)

// displayName 从卖家详情中获取展示名称。
func displayName(s *risk.Seller) string {
	if s != nil && s.Details != nil {
		return s.Details.TradingName
	}
	return "N/A"
}

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	client, err := airwallex.NewFromEnv("../../.env", sdk.WithBaseURL(sdk.SandboxURL))
	if err != nil {
		slog.Error("创建客户端失败", "error", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	svc := client.Risk()

	// 创建卖家
	slog.Info("创建卖家...")
	seller, err := svc.CreateSeller(ctx, &risk.CreateSellerRequest{
		RequestID:  "seller-req-" + time.Now().Format("20060102150405"),
		TradingName: "Test Seller " + time.Now().Format("20060102"),
	})
	if err != nil {
		if sdk.IsUnauthorized(err) {
			slog.Warn("无权创建卖家", "error", err)
		} else {
			slog.Error("创建卖家失败", "error", err)
		}
		os.Exit(0)
	}
	slog.Info("创建卖家成功", "id", seller.ID, "name", displayName(seller), "status", seller.Status)

	// 获取卖家详情
	slog.Info("获取卖家详情...")
	fetched, err := svc.GetSeller(ctx, seller.ID)
	if err != nil {
		slog.Error("获取卖家详情失败", "error", err)
		os.Exit(0)
	}
	slog.Info("卖家详情", "id", fetched.ID, "name", displayName(fetched), "status", fetched.Status)

	// 列出所有卖家
	slog.Info("列出所有卖家...")
	sellers, err := svc.ListSellers(ctx)
	if err != nil {
		slog.Error("列出卖家失败", "error", err)
		os.Exit(0)
	}
	slog.Info("卖家数量", "count", len(sellers.Items))
	for _, s := range sellers.Items {
		slog.Info("卖家", "id", s.ID, "name", displayName(&s), "status", s.Status)
	}
}
