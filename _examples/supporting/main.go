package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/supporting"
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
	svc := client.Supporting()

	// 列出连接店铺（Connected Stores 为只读资源）
	slog.Info("列出连接店铺...")
	stores, err := svc.ListConnectedStores(ctx, &supporting.ListConnectedStoresRequest{})
	if err != nil {
		slog.Error("列出连接店铺失败", "error", err)
		os.Exit(0)
	}
	slog.Info("连接店铺数量", "count", len(stores.Items))
	for _, store := range stores.Items {
		slog.Info("连接店铺", "id", store.ID, "name", store.Name, "status", store.Status)
	}

	// 如果列表不为空，获取第一个店铺的详情
	if len(stores.Items) > 0 {
		slog.Info("获取店铺详情...")
		fetched, err := svc.GetConnectedStore(ctx, stores.Items[0].ID)
		if err != nil {
			slog.Error("获取店铺详情失败", "error", err)
			os.Exit(0)
		}
		slog.Info("店铺详情", "id", fetched.ID, "name", fetched.Name, "status", fetched.Status)
	}
}
