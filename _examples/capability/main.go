package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/capability"
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
	svc := client.Capability()

	// 获取资金限额
	slog.Info("获取资金限额")
	limits, err := svc.GetFundingLimits(ctx)
	if err != nil {
		if sdk.IsUnauthorized(err) {
			slog.Warn("无权访问资金限额", "error", err)
		} else {
			slog.Error("获取资金限额失败", "error", err)
		}
		return
	}
	slog.Info("资金限额数量", "count", len(limits.Items))
	for _, l := range limits.Items {
		slog.Info("限额详情",
			"currency", l.Currency,
			"limit", l.Limit,
			"status", l.Status,
			"type", l.Type,
		)
	}

	// 获取能力详情
	slog.Info("获取能力详情")
	cap, err := svc.GetAccountCapability(ctx, capability.CapabilityPaymentsVisa)
	if err != nil {
		if sdk.IsUnauthorized(err) {
			slog.Warn("无权访问能力详情", "error", err)
		} else {
			slog.Error("获取能力详情失败", "error", err)
		}
		return
	}
	slog.Info("能力详情", "id", cap.ID, "type", cap.EntityType, "status", cap.Status)
}
