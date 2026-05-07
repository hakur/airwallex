package webhook

// webhook post 数据结构和事件类型定义
//
// 更新工作流：
// 1. 当 Airwallex 官方文档更新时，使用以下 URL 查看最新事件类型和 payload 示例
// 2. 使用大模型阅读文档后生成新的/更新的事件结构体
// 3. 将生成的事件类型添加到 webhook/types.go 常量中
// 4. 将生成的类型化事件添加到 webhook/events/{domain}.go 中
// 5. 运行 go vet ./... && go test ./webhook/... 验证
//
// 数据结构说明：
// - Event: 通用事件外壳（verify.go + event.go）
// - EventName: 事件类型常量（types.go）
// - Typed Events: 类型化事件结构体（events/*.go）
// - Data 模式：
//   • 多数领域（online-payments 等）: {"object": {...}}
//   • issuing 等领域: 平铺对象
//
// 官方文档 URL 列表
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/account.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/account-capability.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/balance.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/billing.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/charges.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/connected-account-transfers.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/conversions.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/deposits.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/direct-debit-payouts.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/global-accounts.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/issuing.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/linked-accounts.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/online-payments.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/platform.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/platform-liquidity-program.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/psp-agnostic.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/rfi.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/spend.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/tax.md
// https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/transfers.md
