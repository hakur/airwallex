package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/issuing"
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
	svc := client.Issuing()

	// 列出现有持卡人
	slog.Info("列出持卡人")
	cardholders, err := svc.ListCardholders(ctx)
	if err != nil {
		slog.Error("列出持卡人失败", "error", err)
		return
	}
	slog.Info("持卡人数量", "count", len(cardholders.Items))
	for _, ch := range cardholders.Items {
		name := ""
		if ch.Individual != nil && ch.Individual.Name != nil {
			name = ch.Individual.Name.FirstName + " " + ch.Individual.Name.LastName
		}
		slog.Info("持卡人", "id", ch.ID, "name", name, "email", ch.Email, "status", ch.Status)
	}

	// 创建持卡人
	slog.Info("创建持卡人")
	now := time.Now().Format("20060102150405")
	cardholder, err := svc.CreateCardholder(ctx, &issuing.CreateCardholderRequest{
		RequestID: "ch-req-" + now,
		Type:      issuing.CardholderTypeIndividual,
		Email:     "test-" + now + "@example.com",
		Individual: &issuing.CardholderIndividual{
			Name: &issuing.CardholderName{
				FirstName: "Test",
				LastName:  "User",
			},
			DateOfBirth: "1990-01-01",
			Address: &issuing.CardholderAddress{
				City:     "Shanghai",
				Country:  "CN",
				Line1:    "123 Nanjing Road",
				Postcode: "200000",
			},
			ExpressConsentObtained: "yes",
		},
	})
	if err != nil {
		if sdk.IsUnauthorized(err) {
			slog.Warn("无权创建持卡人（需要 Issuing 权限）", "error", err)
		} else {
			slog.Error("创建持卡人失败", "error", err)
		}
		return
	}
	chName := ""
	if cardholder.Individual != nil && cardholder.Individual.Name != nil {
		chName = cardholder.Individual.Name.FirstName + " " + cardholder.Individual.Name.LastName
	}
	slog.Info("创建持卡人成功", "id", cardholder.ID, "name", chName)

	// 创建卡片
	slog.Info("创建卡片")
	card, err := svc.CreateCard(ctx, &issuing.CreateCardRequest{
		RequestID:    "card-req-" + time.Now().Format("20060102150405"),
		CardholderID: cardholder.ID,
		Program: issuing.CardProgram{
			Purpose: "CONSUMER",
			Type:    "DEBIT",
		},
		IsPersonalized: true,
		FormFactor:     "VIRTUAL",
		CreatedBy:      "Test User",
	})
	if err != nil {
		if sdk.IsInvalidArgument(err) || sdk.IsUnauthorized(err) {
			slog.Warn("创建卡片需要额外的发卡权限/激活，跳过", "error", err)
		} else {
			slog.Error("创建卡片失败", "error", err)
			return
		}
	} else {
		slog.Info("创建卡片成功", "id", card.ID, "status", card.Status)
	}

	// 列出所有卡片
	slog.Info("列出所有卡片")
	cards, err := svc.ListCards(ctx)
	if err != nil {
		slog.Error("列出卡片失败", "error", err)
		return
	}
	slog.Info("卡片数量", "count", len(cards.Items))
	for _, c := range cards.Items {
		slog.Info("卡片", "id", c.ID, "brand", c.Brand, "status", c.Status)
	}
}
