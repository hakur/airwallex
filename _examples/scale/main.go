package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/scale"
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
	svc := client.Scale()

	// 列出连接账户
	slog.Info("列出连接账户")
	accounts, err := svc.ListAccounts(ctx)
	if err != nil {
		slog.Error("列出连接账户失败", "error", err)
		return
	}
	slog.Info("连接账户数量", "count", len(accounts.Items))
	for _, acct := range accounts.Items {
		slog.Info("连接账户", "id", acct.ID, "email", acct.Email, "status", acct.Status)
	}

	// 创建连接账户
	slog.Info("创建连接账户")
	now := time.Now().Format("20060102150405")
	// 注意：创建企业账户需要平台商户权限。
	// 如果当前账号没有平台权限，下面的调用会返回错误，
	// 示例仍会优雅地以 Warn 级别记录而不是崩溃退出。
	created, err := svc.CreateAccount(ctx, &scale.CreateAccountRequest{
		RequestID: "acct-req-" + now,
		Email:     "sub-account-" + now + "@example.com",
		PrimaryContact: &scale.PrimaryContact{
			FirstName: "Test",
			LastName:  "User",
			Email:     "test-" + now + "@example.com",
		},
		AccountDetails: &scale.AccountDetails{
			BusinessName: "Test Business",
		},
	})
	if err != nil {
		if sdk.IsUnauthorized(err) || sdk.IsInvalidArgument(err) {
			slog.Warn("创建账户失败（需要平台商户权限，当前账号可能不是平台商户）", "error", err)
		} else {
			slog.Error("创建账户失败", "error", err)
		}
		return
	}
	slog.Info("创建账户成功", "id", created.ID, "email", created.Email, "status", created.Status)

	// 获取账户详情
	slog.Info("获取账户详情")
	fetched, err := svc.GetAccount(ctx, created.ID)
	if err != nil {
		slog.Error("获取账户详情失败", "error", err)
		return
	}
	slog.Info("账户详情", "id", fetched.ID, "email", fetched.Email, "status", fetched.Status)
}
