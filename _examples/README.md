# Airwallex Go SDK 示例

本目录包含 19 个可独立运行的示例程序，演示 SDK 的全部业务域功能。

## 前置条件

在项目根目录创建 `.env` 文件：

```bash
cp ../.env.example ../.env
# 编辑 ../.env 填入真实密钥
```

## 运行示例

```bash
cd <example-directory>
go run main.go
```

## 示例列表

| 示例 | 模块 | 说明 |
|---|---|---|
| `quickstart` | core | 从 .env 加载密钥，查询账户余额 |
| `payment_intent` | pa | PaymentIntent 生命周期：创建→确认→捕获 |
| `payout` | payouts | 创建收款人，发起转账 |
| `platform` | core | 使用 `OnBehalfOf` 代理子商户操作 |
| `webhook_server` | webhook | 启动 HTTP 服务接收并验证 webhook |
| `fx_conversion` | fx | 查询汇率，创建外汇兑换 |
| `issuing` | issuing | 创建持卡人，发行虚拟卡 |
| `billing_invoice` | billing | 创建账单客户，管理发票 |
| `global_account` | core | 创建全球账户，查询银行详情 |
| `subscription_recurring` | billing + pa | 包月/包年订阅，自动扣款签约与取消 |
| `subscription_yearly_discount` | billing | 包年订阅 + 首年优惠券折扣 |
| `payment_callback` | webhook + pa | 支付成功回调通知：创建→确认→捕获→webhook 签名验证 |
| `capability` | capability | 账户能力查询与开通 |
| `confirmation` | confirmation | 确认函生成 |
| `finance` | finance | 财务交易记录查询 |
| `risk` | risk | 卖家管理、风控操作 |
| `scale` | scale | 平台多商户账户管理 |
| `simulation` | simulation | 沙箱模拟测试事件 |
| `spend` | spend | 账单、费用、采购订单管理 |
| `supporting` | supporting | 文件上传下载、连接店铺管理 |

## 按业务场景选择示例

| 业务场景 | 推荐示例 |
|---|---|
| 收款（Payment Acceptance） | `payment_intent` |
| 付款/转账（Payouts） | `payout` |
| 外汇兑换（FX） | `fx_conversion` |
| 发卡（Issuing） | `issuing` |
| 计费订阅（Billing） | `billing_invoice`、`subscription_recurring`、`subscription_yearly_discount` |
| 全球账户（Global Account） | `global_account` |
| 平台多商户（Platform） | `platform`、`scale` |
| Webhook 验证 | `webhook_server` |
| 风控（Risk） | `risk` |
| 沙箱模拟（Simulation） | `simulation` |
| 支出管理（Spend） | `spend` |
| 支持服务（Supporting） | `supporting` |

## 约定

- 每个示例是独立的 `package main`，仅一个 `main.go` 文件
- 使用 `sdk.WithBaseURL(sdk.SandboxURL)` 确保安全性
- 错误处理使用 `log/slog`，非生产代码
- 模式：创建客户端 → 调用 API → 输出结果
- 幂等请求 ID：`time.Now().Format("20060102150405")`
