package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/spend"
	"github.com/hakur/airwallex/sdk"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	client, err := airwallex.NewFromEnv("../../.env", sdk.WithBaseURL(sdk.SandboxURL))
	if err != nil {
		slog.Error("创建客户端失败", "error", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	svc := client.Spend()

	// 列出账单（Bills 为只读资源，不支持创建/更新）
	slog.Info("列出账单...")
	bills, err := svc.ListBills(ctx, &spend.ListBillsRequest{})
	if err != nil {
		slog.Error("列出账单失败", "error", err)
		os.Exit(0)
	}
	slog.Info("账单数量", "count", len(bills.Items))
	for _, bill := range bills.Items {
		slog.Info("账单", "id", bill.ID, "amount", bill.BillingAmount, "currency", bill.BillingCurrency, "status", bill.Status)
	}

	// 获取账单详情
	if len(bills.Items) > 0 {
		slog.Info("获取账单详情...")
		fetched, err := svc.GetBill(ctx, bills.Items[0].ID)
		if err != nil {
			slog.Error("获取账单详情失败", "error", err)
			os.Exit(0)
		}
		slog.Info("账单详情", "id", fetched.ID, "amount", fetched.BillingAmount, "currency", fetched.BillingCurrency, "status", fetched.Status)
	}
}
