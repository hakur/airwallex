package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/finance"
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
	svc := client.Finance()

	// 查询财务交易
	slog.Info("查询财务交易...")
	transactions, err := svc.ListFinancialTransactions(ctx)
	if err != nil {
		slog.Error("查询财务交易失败", "error", err)
		os.Exit(0)
	}
	slog.Info("财务交易数", "count", len(transactions.Items))
	for _, tx := range transactions.Items {
		slog.Info("财务交易", "id", tx.ID, "amount", tx.Amount, "currency", tx.Currency, "type", tx.TransactionType, "status", tx.Status)
	}

	// 查询结算记录
	slog.Info("查询结算记录...")
	settlements, err := svc.ListSettlements(ctx, &finance.ListSettlementsRequest{})
	if err != nil {
		slog.Error("查询结算记录失败", "error", err)
		os.Exit(0)
	}
	slog.Info("结算记录数", "count", len(settlements.Items))
	for _, s := range settlements.Items {
		slog.Info("结算记录", "id", s.ID, "amount", s.Amount, "currency", s.Currency, "status", s.Status)
	}
}
