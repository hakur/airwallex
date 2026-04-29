package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/payouts"
	"github.com/hakur/airwallex/sdk"
)

func main() {
	client, err := airwallex.NewFromEnv("../../.env",
		sdk.WithBaseURL(sdk.SandboxURL),
	)
	if err != nil {
		slog.Error("创建客户端失败", "error", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	svc := client.Payouts()

	// 优先复用已有收款人，避免重复创建
	var benID string
	benList, err := svc.ListBeneficiaries(ctx)
	if err != nil {
		slog.Error("列出收款人失败", "error", err)
		return
	}

	if len(benList.Items) > 0 {
		benID = benList.Items[0].ID
		slog.Info("复用已有收款人", "id", benID)
	} else {
		// 没有收款人时才创建
		ben, err := svc.CreateBeneficiary(ctx, &payouts.CreateBeneficiaryRequest{
			RequestID:       "ben-req-" + time.Now().Format("20060102150405"),
			TransferMethods: []string{"LOCAL"},
			Beneficiary: &payouts.BeneficiaryInput{
				Type:       payouts.BeneficiaryTypeBankAccount,
				EntityType: "COMPANY",
				Name:       "Test Beneficiary",
				Email:      "test@example.com",
				Address: &payouts.BeneficiaryAddress{
					CountryCode:   sdk.CountryCodeUS,
					State:         "New York",
					City:          "New York",
					StreetAddress: "123 Main St",
					Postcode:      "10001",
				},
				BankDetails: &payouts.BankDetails{
					AccountCurrency:      sdk.CurrencyUSD,
					AccountName:          "Test User",
					AccountNumber:        "50001121",
					BankAccountCategory:  "Checking",
					BankName:             "Test Bank",
					BankCountryCode:      sdk.CountryCodeUS,
					LocalClearingSystem:  "ACH",
					AccountRoutingType1:  "aba",
					AccountRoutingValue1: "021000021",
				},
			},
		})
		if err != nil {
			slog.Error("创建收款人失败", "error", err)
			return
		}
		benID = ben.ID
		slog.Info("创建收款人", "id", benID)
	}

	// 创建转账
	transfer, err := svc.CreateTransfer(ctx, &payouts.CreateTransferRequest{
		RequestID:        "tr-req-" + time.Now().Format("20060102150405"),
		BeneficiaryID:    benID,
		TransferAmount:   50.0,
		TransferCurrency: sdk.CurrencyUSD,
		TransferMethod:   payouts.TransferMethodLocal,
		SourceCurrency:   sdk.CurrencyUSD,
		Reason:           "business_expenses",
		Reference:        "test-ref-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		slog.Error("创建转账失败", "error", err)
		return
	}
	slog.Info("创建转账", "id", transfer.ID, "status", transfer.Status)

	// 获取转账详情
	fetched, err := svc.GetTransfer(ctx, transfer.ID)
	if err != nil {
		if sdk.IsResourceNotFound(err) {
			slog.Warn("转账不存在", "id", transfer.ID)
		} else {
			slog.Error("获取转账失败", "error", err)
		}
		return
	}
	slog.Info("获取转账", "id", fetched.ID, "amount", fetched.TransferAmount, "currency", fetched.TransferCurrency)
}
