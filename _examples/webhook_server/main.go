package main

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/hakur/airwallex/webhook"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	secret := os.Getenv("AIRWALLEX_WEBHOOK_SECRET")
	if secret == "" {
		slog.Error("请设置环境变量 AIRWALLEX_WEBHOOK_SECRET")
		os.Exit(1)
	}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		// 为每个请求设置超时上下文
		ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
		defer cancel()

		payload, err := io.ReadAll(r.Body)
		if err != nil {
			slog.WarnContext(ctx, "读取请求体失败", "error", err)
			http.Error(w, "读取请求体失败", http.StatusBadRequest)
			return
		}

		// Airwallex 使用两个独立 header: x-timestamp + x-signature
		timestampHeader := r.Header.Get("x-timestamp")
		signatureHeader := r.Header.Get("x-signature")
		if timestampHeader == "" || signatureHeader == "" {
			slog.WarnContext(ctx, "缺少签名 header")
			http.Error(w, "缺少签名 header", http.StatusBadRequest)
			return
		}

		// 验证签名（官方格式: HMAC-SHA256(secret, "{timestamp}{body}")）
		if err := webhook.VerifySignature(payload, timestampHeader, signatureHeader, secret); err != nil {
			slog.WarnContext(ctx, "签名验证失败", "error", err)
			http.Error(w, "签名验证失败", http.StatusUnauthorized)
			return
		}

		// 解析事件
		evt, err := webhook.ParseEvent(payload)
		if err != nil {
			slog.WarnContext(ctx, "解析事件失败", "error", err)
			http.Error(w, "解析事件失败", http.StatusBadRequest)
			return
		}

		slog.InfoContext(ctx, "收到 webhook", "event", evt.Event, "id", evt.ID)

		// 根据事件类型处理（使用 UnmarshalData 泛型解析具体负载）
		switch evt.Event {
		case "payment_intent.succeeded":
			type PaymentIntentData struct {
				ID       string  `json:"id"`
				Status   string  `json:"status"`
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			}
			data, err := webhook.UnmarshalData[PaymentIntentData](evt)
			if err == nil {
				slog.InfoContext(ctx, "PaymentIntent 成功", "id", data.ID, "amount", data.Amount, "currency", data.Currency)
			}
		default:
			slog.InfoContext(ctx, "未处理事件类型", "event", evt.Event)
		}

		w.WriteHeader(http.StatusOK)
	})

	port := ":8080"
	slog.Info("Webhook 服务启动", "url", "http://localhost"+port+"/webhook")
	if err := http.ListenAndServe(port, nil); err != nil {
		slog.Error("服务启动失败", "error", err)
		os.Exit(1)
	}
}
