# Airwallex Go SDK 完整使用指南

> 面向开发者与 AI Agent（LLM）的全面 SDK 使用参考。请先阅读本文档再编写涉及此 SDK 的代码。

## 快速开始

### 环境准备

在项目根目录创建 `.env` 文件（参考 `.env.example`）：

```env
AIRWALLEX_CLIENT_ID=your-client-id
AIRWALLEX_API_KEY=your-api-key
```

### 创建客户端

```go
import (
    "context"
    "github.com/hakur/airwallex"
    "github.com/hakur/airwallex/sdk"
)

// Sandbox 环境（推荐用于开发/测试）
client := airwallex.New("your-client-id", "your-api-key",
    sdk.WithBaseURL(sdk.SandboxURL),
    sdk.WithDebug(true))

// 从 .env 文件加载
client, err := airwallex.NewFromEnv(".env",
    sdk.WithBaseURL(sdk.SandboxURL))

// Production 环境
client := airwallex.New("your-client-id", "your-api-key",
    sdk.WithBaseURL(sdk.ServerURL))

ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
```

### 最小示例：查询余额

```go
balances, err := client.Core().GetCurrentBalances(ctx)
if err != nil {
    log.Fatal(err)
}
for _, b := range balances {
    log.Printf("%s: %s", b.Currency, b.AvailableAmount)
}
```

---

## 子包参考

SDK 包含 15 个业务域子包，通过 `client.XX()` 访问器调用：

| 子包 | 访问器 | 业务域 | Sandbox 状态 |
|------|--------|--------|:---:|
| `pa/` | `client.PA()` | Payment Acceptance（支付受理） | ✅ |
| `payouts/` | `client.Payouts()` | 付款/转账 | ✅ |
| `core/` | `client.Core()` | 核心资源（余额、存款、全球账户） | ✅ |
| `issuing/` | `client.Issuing()` | 发卡（持卡人、卡片、交易） | ✅ |
| `fx/` | `client.FX()` | 外汇（汇率、兑换、报价） | ✅ |
| `billing/` | `client.Billing()` | 计费（订阅、发票、价格、产品） | ✅ |
| `finance/` | `client.Finance()` | 财务（交易记录、结算） | ⚠️ 部分 401 |
| `scale/` | `client.Scale()` | 平台（账户、费用、流动性） | ✅ |
| `risk/` | `client.Risk()` | 风控（卖家、RFI、观察名单） | ⚠️ 部分 403 |
| `simulation/` | `client.Simulation()` | 沙箱模拟（触发测试事件） | ✅ |
| `webhook/` | — (独立函数) | Webhook 签名验证 | N/A |
| `spend/` | `client.Spend()` | 支出（账单、费用、采购订单） | ⚠️ 部分 401 |
| `supporting/` | `client.Supporting()` | 支持服务（文件、连接商店） | ⚠️ 部分不可用 |
| `capability/` | `client.Capability()` | 账户能力查询 | ✅ |
| `confirmation/` | `client.Confirmation()` | 确认函 | ⚠️ 400 |

### 各子包典型用法

```go
// Payment Acceptance — 支付意图
pi, _ := client.PA().CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
    RequestID: "req-" + time.Now().Format("20060102150405"),
    Amount:    "100.00",
    Currency:  sdk.CurrencyUSD,
})
confirmed, _ := client.PA().ConfirmPaymentIntent(ctx, pi.ID, &pa.ConfirmPaymentIntentRequest{RequestID: "conf-" + time.Now().Format("20060102150405")})

// Payouts — 转账
beneficiary, _ := client.Payouts().CreateBeneficiary(ctx, &payouts.CreateBeneficiaryRequest{...})
transfer, _ := client.Payouts().CreateTransfer(ctx, &payouts.CreateTransferRequest{
    RequestID:    "tfr-" + time.Now().Format("20060102150405"),
    BeneficiaryID: beneficiary.ID,
    SourceAmount:   "500.00",
    SourceCurrency: sdk.CurrencyUSD,
})

// FX — 汇率与兑换
rates, _ := client.FX().GetRates(ctx, &fx.GetRatesRequest{
    BuyCurrency:  sdk.CurrencyHKD,
    SellCurrency: sdk.CurrencyUSD,
    BuyAmount:    "1000.00",
})

// Issuing — 发卡
cardholder, _ := client.Issuing().CreateCardholder(ctx, &issuing.CreateCardholderRequest{...})
card, _ := client.Issuing().CreateCard(ctx, &issuing.CreateCardRequest{CardholderID: cardholder.ID})

// Billing — 订阅计费
product, _ := client.Billing().CreateProduct(ctx, &billing.CreateProductRequest{Name: "My Product"})
price, _ := client.Billing().CreatePrice(ctx, &billing.CreatePriceRequest{
    Currency:    sdk.CurrencyUSD,
    UnitAmount:  "29.99",
    ProductID:   product.ID,
})
subscription, _ := client.Billing().CreateSubscription(ctx, &billing.CreateSubscriptionRequest{
    BillingCustomerID: customer.ID,
    CollectionMethod:  billing.CollectionMethodChargeOnCheckout,
    Items: []billing.SubscriptionItemInput{{PriceID: price.ID, Quantity: 1}},
})

// Simulation — 沙箱模拟
client.Simulation().SimulateShopperAction(ctx, paymentIntentID, &simulation.SimulateShopperActionRequest{
    Action: simulation.ShopperActionConfirmPayment,
})
client.Simulation().SimulateGlobalAccountDeposit(ctx, &simulation.SimulateGlobalAccountDepositRequest{
    GlobalAccountID: gaID,
    Amount:          "100.00",
})

// Webhook — 签名验证（无需 client 实例）
webhook.VerifySignature(payload, timestampHeader, signatureHeader, webhookSecret)

// Scale — 平台多商户
client.Scale().CreateHostedFlow(ctx, &scale.CreateHostedFlowRequest{Type: scale.HostedFlowTypeKYC})
client.Scale().CreatePSPSettlementIntent(ctx, &scale.CreatePSPSettlementIntentRequest{...})
```

---

## 字段类型约定 ⚠️ 最重要

### 🔴 货币金额：必须用 `string`

API 返回的金额字段格式为字符串（如 `"100.00"`），使用 `float64` 会导致 **JSON 反序列化失败**或精度丢失。

```go
// ✅ 正确
type PaymentIntent struct {
    Amount string `json:"amount"`
}

// ❌ 错误 — 会导致 unmarshal error
type PaymentIntent struct {
    Amount float64 `json:"amount"`
}
```

**货币金额字段示例**：`Amount`, `Fee`, `Net`, `Balance`, `UnitAmount`, `FlatAmount`, `SourceAmount`, `TransferAmount`, `BillingAmount`, `TransactionAmount`, `TotalAmount`, `GrossAmount`, `FeeAmount`, 等等。

### 🟢 汇率：保持 `float64`

API 的汇率字段（`client_rate`, `awx_rate`, `mid_rate`）是真正的浮点数，不是货币金额。

```go
// ✅ 正确
type Conversion struct {
    ClientRate float64 `json:"client_rate,omitempty"`
    BuyAmount  string  `json:"buy_amount"`
}
```

### 🟢 百分比/非货币数值：保持 `float64`

`DefaultTaxPercent`, `PercentageOff`, `RollingPercentage`, `Quantity`, `UpperBound` 等不是货币金额。

### 🔵 时间戳：必须用 `string`

```go
// ✅ 正确
CreatedAt string `json:"created_at,omitempty"`

// ❌ 错误
CreatedAt time.Time `json:"created_at,omitempty"`
```

### 🔵 枚举：用 type alias

```go
// ✅ 正确 — type alias 保持 string 赋值兼容
type Status = string

// ❌ 错误 — 需要显式类型转换
type Status string
```

---

## 测试规范

### 强制规则

1. **只用 testify**：`require.NoError(t, err)` / `assert.Equal(t, expected, actual)`
2. **禁止 `t.Skip()`**：不允许跳过测试
3. **禁止裸 `if` 判断**：不要写 `if err != nil { t.Fatal(err) }`
4. **幂等 RequestID**：`"prefix-" + time.Now().Format("20060102150405")`
5. **真实 Sandbox 环境**：不 mock，直接调用 Sandbox API

```go
// ✅ 正确
require.NoError(t, err, "create failed")
assert.Equal(t, expected, actual, "should match")

// ❌ 错误
if err != nil {
    t.Fatal(err)
}
t.Skip("not implemented")
```

### 测试结构

每个包有独立的 `TestMain` 初始化共享 `testClient`：

```go
func TestMain(m *testing.M) {
    _ = godotenv.Load(sdk.ResolveEnvPath())
    testClient, _ = airwallex.NewFromEnv("", sdk.WithBaseURL(sdk.SandboxURL), sdk.WithDebug(true))
    os.Exit(m.Run())
}
```

测试函数使用 lifecycle 模式：创建 → 查询 → 更新 → 列表 → 删除/取消。

### 优雅跳过

Sandbox 中某些端点返回 401/403/404。测试应优雅处理：

```go
if err != nil {
    require.True(t, sdk.IsUnauthorized(err) || sdk.IsForbidden(err) || sdk.IsNotFound(err),
        "unexpected error: %v", err)
    t.Logf("test skipped due to sandbox limitation: %v", err)
    return
}
```

---

## 常见陷阱

### 1. 硬编码 Fake ID 🚫

**不要**在测试中硬编码 ID。始终从 List API 获取真实 ID：

```go
// ❌ 错误
svc.GetDownloadLinks(ctx, &supporting.DownloadLinkRequest{FileIDs: []string{"test-file-id"}})

// ✅ 正确
list, _ := svc.ListConnectedStores(ctx, &supporting.ListConnectedStoresRequest{PageSize: 10})
svc.GetConnectedStore(ctx, list.Items[0].ID)
```

唯一的例外是**故意测试错误处理**时使用假 ID——但必须加注释说明。

### 2. Subscription 竞态条件

更新订阅后立即取消可能触发 `"The subscription is already updated."` 错误。在操作之间增加短暂等待：

```go
svc.UpdateSubscription(ctx, id, req)
time.Sleep(200 * time.Millisecond)  // 避免竞态
svc.CancelSubscription(ctx, id, cancelReq)
```

或优雅处理：`require.True(t, sdk.IsUnauthorized(err) || sdk.IsValidationError(err))`

### 3. Sandbox 限制

以下端点类型在 Sandbox 中通常不可用：
- `finance/financial_reports` → 401
- `finance/settlements` → 401
- `risk/seller` → 403
- `spend/` 大部分端点 → 401
- `pa/customs_declarations` → 404
- `pa/funds_splits` → 404

### 4. API 版本

SDK 硬编码 API 版本 `2026-02-27`（在 `sdk/sdk.go`）。Changelog 中此日期之后的变更是未来的，不适用于当前 SDK。

### 5. 循环依赖

子包之间**禁止互相引用**。所有子包只依赖根 `github.com/hakur/airwallex` 和 `sdk/` 包。

### 6. 提交前检查

```bash
go build ./...       # 编译通过
go vet ./...         # 无警告
golangci-lint run    # lint 通过
```

---

## 架构约定

- **域驱动包结构**：15 个业务域包在根级，不使用 `pkg/` 或 `internal/`
- **`sdk.Doer` 接口注入**：所有 Service 通过构造函数接收 HTTP 执行器
- **Bearer Token 自动管理**：首次获取 + 过期前 5 分钟预刷新 + 401 兜底重试
- **OpenSpec 规范驱动**：变更流程为 Explore → Propose → Apply → Archive
- **示例程序**：`_examples/` 目录下 18 个独立可运行的 main.go

---

## 端点状态跟踪

完整的端点覆盖状态见 [`docs/api-endpoints.md`](api-endpoints.md)。

每个端点记录：HTTP 方法 + 路径 → Go 函数 → Go 文件 → 测试函数 → 状态标识。

状态标识：✅ PASS | 🔒 权限不足 | ⏭️ SKIP | 🚫 无法测试

---

## 常见问题

**Q: 为什么编译通过但运行时 JSON 解析失败？**
A: 检查 struct 中货币类型字段是否用了 `float64`。API 返回 `"amount": "100.00"`（字符串），但 Go 期望 float64 → unmarshal error。

**Q: 如何获取真实的资源 ID？**
A: 先调用 List API（如 `ListCustomers`）获取真实数据，从返回结果中提取 ID。不要硬编码。

**Q: 测试报 "validation_error" 怎么办？**
A: 检查请求参数是否符合 API 规范——查看对应函数注释中的 `// 官方文档:` URL。
