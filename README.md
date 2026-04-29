# Airwallex Go SDK

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.26.1-blue)](https://golang.org)
[![GoDoc](https://godoc.org/github.com/hakur/airwallex?status.svg)](https://godoc.org/github.com/hakur/airwallex)

Airwallex（空中云汇）Go SDK 封装，由 opencode + Deepseek v4 pro + kimi 2.6 生成，请使用时 coding agent 粘贴本项目的 Readme.md 的URL以让大模型自行理解这个项目

> 📖 [api索引 (docs/api-endpoints.md)](docs/api-endpoints.md) — 已实现的API列表、单元测试状态
> 📖 [使用指南 (docs/full-usage.md)](docs/full-usage.md) — 子包参考、类型约定、测试规范、常见陷阱
> 📖 [示例程序 (_examples/)](_examples/) — 19 个独立可运行示例
> 📖 [English Documentation (README_EN.md)](README_EN.md)

## 安装

```bash
go get github.com/hakur/airwallex@latest
```

## 快速开始

### 方式一：显式传入密钥

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/hakur/airwallex"
    "github.com/hakur/airwallex/sdk"
)

func main() {
    // 创建客户端（Production 环境）
    // 正式服务器地址是 sdk.ServerURL
    // client := airwallex.New("your-client-id", "your-api-key", sdk.WithBaseURL(sdk.ServerURL))
    // 从环境变量中加载，参考 项目根目录的 .env.example 
    // client, err := airwallex.NewFromEnv("", sdk.WithBaseURL(sdk.ServerURL)) // 读取环境变量
    // client, err := airwallex.NewFromEnv(".env", sdk.WithBaseURL(sdk.ServerURL)) // 读取环境变量文件
    // 或使用 Sandbox 环境
    client := airwallex.New("your-client-id", "your-api-key", sdk.WithBaseURL(sdk.SandboxURL))
    ctx,cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()

    // 查询余额
    balances, err := client.Core().GetCurrentBalances(ctx)
    if err != nil {
        log.Fatal(err)
    }
    for _, b := range balances {
        log.Printf("%s: %.2f", b.Currency, b.AvailableAmount)
    }
}
```

## 平台多商户（未测试）

```go
// Client 级别：所有请求自动携带 x-on-behalf-of
client := airwallex.New(id, key, sdk.WithOnBehalfOf("acct_123"))

// 请求级别：覆盖默认值
client.Payouts().CreateTransfer(ctx, req, sdk.WithRequestOnBehalfOf("acct_456"))
```

## 子包

| 子包 | 业务域 | 状态 | 示例 |
|------|--------|------|------|
| `pa` | Payment Acceptance | ⚠️ 核心可用 | `client.PA().CreatePaymentIntent(ctx, req)` |
| `payouts` | Payouts | ⚠️ 核心可用 | `client.Payouts().CreateTransfer(ctx, req)` |
| `core` | Core Resources | ✅ 推荐 | `client.Core().GetCurrentBalances(ctx)` |
| `issuing` | Issuing | ⚠️ 需要激活 | `client.Issuing().CreateCard(ctx, req)` |
| `fx` | Transactional FX | ✅ 推荐 | `client.FX().CreateConversion(ctx, req)` |
| `billing` | Billing | ⚠️ 高级功能受限 | `client.Billing().CreateInvoice(ctx, req)` |
| `finance` | Finance | ✅ 推荐 | `client.Finance().ListFinancialTransactions(ctx)` |
| `scale` | Scale | ⚠️ 需要平台权限 | `client.Scale().CreateAccount(ctx, req)` |
| `risk` | Risk | ⚠️ 核心可用 | `client.Risk().CreateSeller(ctx, req)` |
| `simulation` | Simulation | ⚠️ 沙箱依赖 | `client.Simulation().SimulateShopperAction(ctx, piID)` |
| `webhook` | Webhook | ⚠️ 验证可用 | `webhook.VerifySignature(payload, tsHeader, sigHeader, secret)` |
| `spend` | Spend | 🚫 沙箱不可用 | `client.Spend().ListBills(ctx, req)` |
| `supporting` | Supporting | ✅ 推荐 | `client.Supporting().ListConnectedStores(ctx, req)` |
| `capability` | Account Capability | ⚠️ 核心可用 | `client.Capability().GetAccountCapability(ctx, "cap_payments")` |
| `confirmation` | Confirmation | ✅ 推荐 | `client.Confirmation().CreateConfirmationLetter(ctx, req)` |

## API 端点总览索引

> 用于自动化同步：每个 Go 源文件对应一个官方 API 端点总览页。
> 标注 `⚠️ 待补充` 的文件尚无已知 API 总览 URL，需人工补充后删除 `⚠️ 待补充` 前缀。

### 已有总览 URL

| Go 文件 | API 端点总览 |
|---------|-------------|
| `billing/billing_customers.go` | `https://www.airwallex.com/docs/api/billing/billing_customers.md` |
| `billing/billing_transactions.go` | `https://www.airwallex.com/docs/api/billing/billing_transactions.md` |
| `billing/checkouts.go` | `https://www.airwallex.com/docs/api/billing/billing_checkouts.md` |
| `billing/coupons.go` | `https://www.airwallex.com/docs/api/billing/coupons.md` |
| `billing/credit_notes.go` | `https://www.airwallex.com/docs/api/billing/credit_notes.md` |
| `billing/invoice.go` | `https://www.airwallex.com/docs/api/billing/invoices.md` |
| `billing/meters.go` | `https://www.airwallex.com/docs/api/billing/meters.md` |
| `billing/payment_sources.go` | `https://www.airwallex.com/docs/api/billing/payment_sources.md` |
| `billing/prices.go` | `https://www.airwallex.com/docs/api/billing/prices.md` |
| `billing/product.go` | `https://www.airwallex.com/docs/api/billing/products.md` |
| `billing/subscription.go` | `https://www.airwallex.com/docs/api/billing/subscriptions.md` |
| `billing/usage_event.go` | `https://www.airwallex.com/docs/api/billing/usage_events.md` |
| `capability/account_capability.go` | `https://www.airwallex.com/docs/api/account_capability/account_capability.md` |
| `confirmation/confirmation_letter.go` | `https://www.airwallex.com/docs/api/confirmation_letter/confirmation_letter.md` |
| `core/balance.go` | `https://www.airwallex.com/docs/api/core_resources/balances.md` |
| `core/deposit.go` | `https://www.airwallex.com/docs/api/core_resources/deposits.md` |
| `core/direct_debits.go` | `https://www.airwallex.com/docs/api/core_resources/direct_debits.md` |
| `core/global_account.go` | `https://www.airwallex.com/docs/api/core_resources/global_accounts.md` |
| `core/linked_account.go` | `https://www.airwallex.com/docs/api/core_resources/linked_accounts.md` |
| `finance/financial_reports.go` | `https://www.airwallex.com/docs/api/finance/financial_reports.md` |
| `finance/financial_transaction.go` | `https://www.airwallex.com/docs/api/finance/financial_transactions.md` |
| `finance/settlement.go` | `https://www.airwallex.com/docs/api/finance/settlements.md` |
| `fx/conversion.go` | `https://www.airwallex.com/docs/api/fx/conversions.md` |
| `fx/conversion_amendment.go` | `https://www.airwallex.com/docs/api/fx/conversion_amendments.md` |
| `fx/quote.go` | `https://www.airwallex.com/docs/api/fx/quotes.md` |
| `fx/rate.go` | `https://www.airwallex.com/docs/api/fx/rates.md` |
| `issuing/authorization.go` | `https://www.airwallex.com/docs/api/issuing/authorizations.md` |
| `issuing/card.go` | `https://www.airwallex.com/docs/api/issuing/cards.md` |
| `issuing/cardholder.go` | `https://www.airwallex.com/docs/api/issuing/cardholders.md` |
| `issuing/card_transaction_events.go` | `https://www.airwallex.com/docs/api/issuing/card_transaction_events.md` |
| `issuing/card_transaction_lifecycles.go` | `https://www.airwallex.com/docs/api/issuing/card_transaction_lifecycles.md` |
| `issuing/config.go` | `https://www.airwallex.com/docs/api/issuing/config.md` |
| `issuing/digital_wallet_tokens.go` | `https://www.airwallex.com/docs/api/issuing/digital_wallet_tokens.md` |
| `issuing/merchant_brands.go` | `https://www.airwallex.com/docs/api/issuing/merchant_brands.md` |
| `issuing/transaction.go` | `https://www.airwallex.com/docs/api/issuing/transactions.md` |
| `issuing/transaction_disputes.go` | `https://www.airwallex.com/docs/api/issuing/transaction_disputes.md` |
| `pa/config.go` | `https://www.airwallex.com/docs/api/payments/config.md` |
| `pa/conversion_quotes.go` | `https://www.airwallex.com/docs/api/payments/conversion_quotes.md` |
| `pa/customer.go` | `https://www.airwallex.com/docs/api/payments/customers.md` |
| `pa/customs_declarations.go` | `https://www.airwallex.com/docs/api/payments/customs_declarations.md` |
| `pa/funds_splits.go` | `https://www.airwallex.com/docs/api/payments/funds_splits.md` |
| `pa/funds_split_reversals.go` | `https://www.airwallex.com/docs/api/payments/funds_split_reversals.md` |
| `pa/pa_reference_data.go` | `https://www.airwallex.com/docs/api/payments/reference_data.md` |
| `pa/payment_attempts.go` | `https://www.airwallex.com/docs/api/payments/payment_attempts.md` |
| `pa/payment_consents.go` | `https://www.airwallex.com/docs/api/payments/payment_consents.md` |
| `pa/payment_disputes.go` | `https://www.airwallex.com/docs/api/payments/payment_disputes.md` |
| `pa/payment_intent.go` | `https://www.airwallex.com/docs/api/payments/payment_intents.md` |
| `pa/payment_link.go` | `https://www.airwallex.com/docs/api/payments/payment_links.md` |
| `pa/payment_method.go` | `https://www.airwallex.com/docs/api/payments/payment_methods.md` |
| `pa/refund.go` | `https://www.airwallex.com/docs/api/payments/refunds.md` |
| `pa/settlement_records.go` | `https://www.airwallex.com/docs/api/payments/settlement_records.md` |
| `pa/terminals.go` | `https://www.airwallex.com/docs/api/payments/terminals.md` |
| `payouts/batch_transfer.go` | `https://www.airwallex.com/docs/api/payouts/batch_transfers.md` |
| `payouts/beneficiary.go` | `https://www.airwallex.com/docs/api/payouts/beneficiaries.md` |
| `payouts/transfer.go` | `https://www.airwallex.com/docs/api/payouts/transfers.md` |
| `risk/fraud_feedback.go` | `https://www.airwallex.com/docs/api/risk/fraud_feedback_issuing.md` |
| `risk/rfi.go` | `https://www.airwallex.com/docs/api/risk/request_for_information_rfi.md` |
| `risk/seller.go` | `https://www.airwallex.com/docs/api/risk/sellers.md` |
| `risk/watchlist.go` | `https://www.airwallex.com/docs/api/risk/watchlist.md` |
| `scale/account.go` | `https://www.airwallex.com/docs/api/scale/accounts.md` |
| `scale/charges.go` | `https://www.airwallex.com/docs/api/scale/charges.md` |
| `scale/connected_account_transfers.go` | `https://www.airwallex.com/docs/api/scale/connected_account_transfers.md` |
| `scale/hosted_flows.go` | `https://www.airwallex.com/docs/api/scale/hosted_flows.md` |
| `scale/invitation_links.go` | `https://www.airwallex.com/docs/api/scale/invitation_links.md` |
| `scale/platform_liquidity_programs.go` | `https://www.airwallex.com/docs/api/scale/platform_liquidity_programs.md` |
| `scale/platform_reports.go` | `https://www.airwallex.com/docs/api/scale/platform_reports.md` |
| `scale/psp_settlement_deposits.go` | `https://www.airwallex.com/docs/api/scale/psp_settlement_deposits.md` |
| `scale/psp_settlement_intents.go` | `https://www.airwallex.com/docs/api/scale/psp_settlement_intents.md` |
| `scale/psp_settlement_splits.go` | `https://www.airwallex.com/docs/api/scale/psp_settlement_splits.md` |
| `simulation/accounts.go` | `https://www.airwallex.com/docs/api/simulation/accounts.md` |
| `simulation/billing.go` | `https://www.airwallex.com/docs/api/simulation/billing.md` |
| `simulation/deposits.go` | `https://www.airwallex.com/docs/api/simulation/deposits.md` |
| `simulation/issuing.go` | `https://www.airwallex.com/docs/api/simulation/issuing.md` |
| `simulation/linked_accounts.go` | `https://www.airwallex.com/docs/api/simulation/linked_accounts.md` |
| `simulation/payments.go` | `https://www.airwallex.com/docs/api/simulation/payments.md` |
| `simulation/payouts.go` | `https://www.airwallex.com/docs/api/simulation/payouts.md` |
| `simulation/rfis.go` | `https://www.airwallex.com/docs/api/simulation/request_for_information_rfi.md` |
| `simulation/terminals.go` | `https://www.airwallex.com/docs/api/simulation/terminal_simulation.md` |
| `simulation/transfers.go` | `https://www.airwallex.com/docs/api/simulation/transfers.md` |
| `spend/bill.go` | `https://www.airwallex.com/docs/api/spend/bills.md` |
| `spend/expense.go` | `https://www.airwallex.com/docs/api/spend/expenses.md` |
| `spend/purchase_order.go` | `https://www.airwallex.com/docs/api/spend/purchase_orders.md` |
| `spend/reimbursement.go` | `https://www.airwallex.com/docs/api/spend/reimbursements.md` |
| `spend/vendor.go` | `https://www.airwallex.com/docs/api/spend/vendors.md` |
| `supporting/connected_store.go` | `https://www.airwallex.com/docs/api/supporting_services/connected_stores.md` |
| `supporting/file.go` | `https://www.airwallex.com/docs/api/supporting_services/file_service/download_links_files.md` |
| `supporting/reference_data.go` | `https://www.airwallex.com/docs/api/supporting_services/reference_data.md` |
| `webhook/verify.go` | `https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events.md` |

### 待补充 API 总览 URL

| Go 文件 | API 端点总览 |
|---------|-------------|
| `webhook/endpoint.go` | ⚠️ 待补充（用户直接粘贴 markdown，URL 待确认） |