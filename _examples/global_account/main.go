package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/core"
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
	svc := client.Core()

	// 查询账户余额
	slog.Info("=== 查询账户余额 ===")
	balances, err := svc.GetCurrentBalances(ctx)
	if err != nil {
		slog.Error("查询余额失败", "error", err)
		return
	}
	slog.Info("持有多种货币余额", "count", len(balances))
	for _, b := range balances {
		if b.AvailableAmount > 0 {
			slog.Info("余额", "currency", b.Currency, "available", b.AvailableAmount, "total", b.TotalAmount)
		}
	}

	// 列出全球账户
	slog.Info("=== 列出全球账户 ===")
	accounts, err := svc.ListGlobalAccounts(ctx)
	if err != nil {
		slog.Warn("列出全球账户失败", "error", err)
		return
	}
	slog.Info("全球账户", "count", len(accounts.Items))
	for _, ga := range accounts.Items {
		slog.Info("账户", "id", ga.ID, "currency", ga.Currency, "country", ga.CountryCode, "status", ga.Status)
		slog.Info("  详情", "account_name", ga.AccountName, "bank", ga.BankName)
	}

	// 创建全球账户（例如：美元账户）
	slog.Info("=== 创建全球账户 ===")
	created, err := svc.CreateGlobalAccount(ctx, &core.CreateGlobalAccountRequest{
		RequestID:   "ga-req-" + time.Now().Format("20060102150405"),
		Currency:    sdk.CurrencyUSD,
		CountryCode: sdk.CountryCodeUS,
	})
	if err != nil {
		if sdk.IsBadRequest(err) {
			slog.Warn("创建全球账户请求参数不合法（可能需要额外的 required_features 参数）", "error", err)
		} else {
			slog.Warn("创建全球账户失败", "error", err)
		}
		return
	}
	slog.Info("创建全球账户", "id", created.ID)
	slog.Info("  币种", "currency", created.Currency)
	slog.Info("  国家", "country", created.CountryCode)
	slog.Info("  账户名", "name", created.AccountName)
	slog.Info("  账号", "number", created.AccountNumber)
	slog.Info("  银行", "bank", created.BankName)
	slog.Info("  SWIFT", "swift", created.BankSwiftCode)

	// 获取银行详情
	slog.Info("=== 获取账户银行详情 ===")
	bankDetails, err := svc.GetGlobalAccountBankDetails(ctx, created.ID)
	if err != nil {
		if sdk.IsResourceNotFound(err) {
			slog.Warn("全球账户银行详情不存在", "id", created.ID)
		} else {
			slog.Warn("获取银行详情失败", "error", err)
		}
	} else {
		slog.Info("银行详情", "details", bankDetails)
	}
}
