package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/confirmation"
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
	svc := client.Confirmation()

	// 创建确认函
	slog.Info("=== 创建确认函 ===")
	err = svc.CreateConfirmationLetter(ctx, &confirmation.CreateConfirmationLetterRequest{
		Format:        "STANDARD",
		TransactionID: "txn-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		if sdk.IsForbidden(err) {
			slog.Warn("无权创建确认函", "error", err)
		} else if sdk.IsBadRequest(err) {
			slog.Warn("创建确认函请求参数不合法", "error", err)
		} else {
			slog.Warn("创建确认函失败", "error", err)
		}
		return
	}
	slog.Info("确认函创建请求已提交（无 JSON 响应体，直接返回 PDF 文件流）")
}
