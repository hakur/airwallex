package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/billing"
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
	svc := client.Billing()

	// 先获取已存在的账单客户，避免硬编码 ID
	slog.Info("获取已有账单客户...")
	customers, err := svc.ListBillingCustomers(ctx, &billing.ListBillingCustomersRequest{})
	if err != nil {
		if sdk.IsUnauthorized(err) {
			slog.Warn("无权限列出账单客户", "error", err)
		} else {
			slog.Error("列出账单客户失败", "error", err)
			os.Exit(1)
		}
	}

	var customerID string
	if customers != nil && len(customers.Items) > 0 {
		customerID = customers.Items[0].ID
		slog.Info("使用已有账单客户", "id", customerID, "name", customers.Items[0].Name)
	} else {
		// 没有现有客户，创建新客户
		slog.Info("未找到已有账单客户，创建新客户...")
		newCust, err := svc.CreateBillingCustomer(ctx, &billing.CreateBillingCustomerRequest{
			RequestID: "bcus-req-" + time.Now().Format("20060102150405"),
			Name:      "示例客户 " + time.Now().Format("20060102"),
			Email:     "customer@example.com",
		})
		if err != nil {
			if sdk.IsUnauthorized(err) {
				slog.Warn("无权创建账单客户", "error", err)
				os.Exit(0)
			}
			slog.Error("创建账单客户失败", "error", err)
			os.Exit(1)
		}
		customerID = newCust.ID
		slog.Info("已创建账单客户", "id", customerID)
	}

	// 创建账单
	slog.Info("创建账单...")
	invoice, err := svc.CreateInvoice(ctx, &billing.CreateInvoiceRequest{
		RequestID:         "inv-req-" + time.Now().Format("20060102150405"),
		BillingCustomerID: customerID,
		Currency:          sdk.CurrencyUSD,
	})
	if err != nil {
		if sdk.IsUnauthorized(err) {
			slog.Warn("无权创建账单（需要 Billing 权限）", "error", err)
		} else {
			slog.Error("创建账单失败", "error", err)
		}
		os.Exit(0)
	}
	slog.Info("创建账单成功", "id", invoice.ID, "amount", invoice.TotalAmount, "currency", invoice.Currency, "status", invoice.Status)

	// 获取账单详情
	slog.Info("获取账单详情...")
	fetched, err := svc.GetInvoice(ctx, invoice.ID)
	if err != nil {
		slog.Error("获取账单失败", "error", err)
		os.Exit(0)
	}
	slog.Info("账单详情", "id", fetched.ID, "amount", fetched.TotalAmount, "currency", fetched.Currency)

	// 更新账单备注
	slog.Info("更新账单备注...")
	updated, err := svc.UpdateInvoice(ctx, invoice.ID, &billing.UpdateInvoiceRequest{
		Memo: "Updated via example",
	})
	if err != nil {
		slog.Warn("更新账单失败", "error", err)
	} else {
		slog.Info("更新后备注", "memo", updated.Memo)
	}

	// 列出所有账单
	slog.Info("列出所有账单...")
	invoices, err := svc.ListInvoices(ctx, &billing.ListInvoicesRequest{})
	if err != nil {
		slog.Error("列出账单失败", "error", err)
		os.Exit(0)
	}
	slog.Info("账单列表", "count", len(invoices.Items))
	for _, inv := range invoices.Items {
		slog.Info("账单", "id", inv.ID, "amount", inv.TotalAmount, "currency", inv.Currency, "status", inv.Status)
	}
}
