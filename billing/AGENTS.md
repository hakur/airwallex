# billing/

**Billing** — 订阅计费、发票、支付方式的完整生命周期管理。

## STRUCTURE

```
billing/
├── service.go              # Service struct, New(doer) constructor
├── product.go              # CRUD Product
├── prices.go               # CRUD Price
├── subscription.go         # CRUD Subscription + items 子资源
├── invoice.go              # CRUD Invoice + line items 子资源
├── checkouts.go            # CRUD Checkout
├── coupons.go              # CRUD Coupon
├── credit_notes.go         # CRUD CreditNote + line items 子资源
├── meters.go               # CRUD Meter + summaries
├── billing_customers.go    # CRUD BillingCustomer
├── billing_transactions.go # List/Get BillingTransaction
├── payment_sources.go      # CRUD PaymentSource (auto-charge)
├── usage_event.go          # Ingest / BatchIngest / Void UsageEvent
├── types.go                # Shared billing types
└── billing_test.go         # 测试入口
```

## WHERE TO LOOK

| Task | Location | Notes |
|------|----------|-------|
| 创建/管理产品 | `product.go` | 对标 Stripe Product，含 metadata 扩展字段 |
| 创建/管理价格 | `prices.go` | 支持 recurring / one_time 定价模型 |
| 订阅管理 | `subscription.go` | 含 `items` 子资源增删改 |
| 发票管理 | `invoice.go` | 含 `line_items` 子资源增删改 |
| 信用票据 | `credit_notes.go` | 结构与 invoice 类似，含 `line_items` |
| 计量计费 | `meters.go` | 含 meters/{id}/summaries 查询 |
| 计费客户 | `billing_customers.go` | 独立领域，与 payments 的 Customer 区分 |
| 支付来源 | `payment_sources.go` | 绑定 auto-charge 支付方式 |
| 用量事件 | `usage_event.go` | 批量写入，支持 void 撤销 |
| 共享类型 | `types.go` | BillingPeriod、BillingInterval 等枚举 |

## CONVENTIONS

- **Stripe-like API 设计**：products / prices / subscriptions / invoices 四层模型。
- **子资源嵌套**：subscriptions/{id}/items、invoices/{id}/line_items 直接写在父文件，不拆独立文件。
- **权限限制**：大多数 billing 端点在 Sandbox 中需特殊权限，测试可能返回 401。
- **批量写入**：UsageEvent 使用数组输入（`[]UsageEvent`），服务端保证幂等。
- **时间戳**：沿用 SDK 全局约定，API 结构体字段均为 `string` 类型。

## ANTI-PATTERNS

- **禁止创建独立的 billing client** — 统一使用根级 `Client.Billing()`。
- **禁止与 pa 包重复类型** — billing 和 payments 虽有共性（如 Customer），但字段集不同，各自维护。
