# Airwallex Go SDK

**Generated:** 2026-04-27 | **Commit:** 3e81e37 | **Branch:** master

## OVERVIEW

Airwallex Go SDK — 覆盖 98 个 REST API 端点的支付平台客户端库。单模块 `github.com/hakur/airwallex`，Go 1.26.1+。

## STRUCTURE

```
./
├── client.go           # 主入口: New(), NewFromEnv(), Client
├── auth.go             # Bearer Token 自动刷新（内部）
├── sdk/                # → sdk/AGENTS.md
├── pa/                 # → pa/AGENTS.md
├── billing/            # → billing/AGENTS.md
├── core/               # 核心资源: 余额、交易、全球账户
├── payouts/            # 付款: 收款人、转账
├── fx/                 # 外汇: 兑换、报价、汇率
├── issuing/            # 发卡: 持卡人、卡管理
├── finance/            # 财务: 交易记录、结算
├── scale/              # 平台: 账户管理
├── risk/               # 风控: 卖家管理
├── simulation/         # 沙箱模拟端点
├── spend/              # 支出: 账单、费用、采购
├── webhook/            # Webhook 签名验证与端点管理
├── supporting/         # 支持服务 (部分端点已弃用)
├── capability/         # 账户能力查询
├── confirmation/       # 确认函
├── _examples/          # → _examples/AGENTS.md
└── openspec/           # 规范驱动开发: specs/, changes/
```

## WHERE TO LOOK

| Task | Location | Notes |
|------|----------|-------|
| 创建客户端 | `client.go:New()` / `NewFromEnv()` | 需 client_id + api_key |
| HTTP 抽象 | `sdk/sdk.go` `Doer` interface | 所有服务依赖此接口 |
| API 错误处理 | `sdk/errors.go` | `APIError` + 错误码常量 + 断言函数 |
| 服务访问器 | `client.go:PA()` / `.Payouts()` / ... | 15 个领域服务方法 |
| 选项配置 | `sdk/options.go` | `WithBaseURL`, `WithDebug`, `WithOnBehalfOf` |
| 枚举常量 | `sdk/currency.go`, `sdk/country.go` | `Currency` (36), `CountryCode` (38) |
| 环境加载 | `sdk/sdk.go:LoadEnv()` | 从 .env 读取凭证 |
| Webhook 验证 | `webhook/verify.go:VerifySignature()` | 无需 Client 实例 |
| 示例程序 | `_examples/*/main.go` | 18 个独立可运行示例 |
| OpenSpec 规范 | `openspec/specs/sdk-*/spec.md` | 按业务域分离 |
| Lint 配置 | `.golangci.yml` | v2 格式，启用 6 个 linter |
| 环境模板 | `.env.example` | `AIRWALLEX_CLIENT_ID` + `AIRWALLEX_API_KEY` |

## CODE MAP

| Symbol | Type | Location | Refs | Role |
|--------|------|----------|------|------|
| `New` | Func | `client.go:47` | ~18 | SDK 入口构造函数 |
| `NewFromEnv` | Func | `client.go:34` | ~18 | 从 .env 构造客户端 |
| `Client` | Struct | `client.go:77` | 全局 | 主客户端，组合 Doer + Authenticator |
| `(*Client).Do` | Method | `client.go:85` | 所有 | 统一 HTTP 请求方法 |
| `Doer` | Interface | `sdk/sdk.go:40` | 全项目 | HTTP 执行器抽象 |
| `APIError` | Struct | `sdk/errors.go:9` | 全局 | API 业务错误类型 |
| `SandboxURL` | Const | `sdk/sdk.go:23` | 全局 | `https://api-demo.airwallex.com` |
| `Service` | Struct | `{domain}/service.go` | 15 | 各业务域服务入口 |
| `VerifySignature` | Func | `webhook/verify.go` | 独立 | Webhook 签名验证 |

## CONVENTIONS

### 代码结构
- **每个业务域一个包**：直接在根目录 (`pa/`, `core/`, ...)，无 `pkg/` 或 `internal/` 隔离
- **每个包一个 `service.go`**：暴露 `Service` 结构体 + `New(doer sdk.Doer) *Service`
- **API 方法分散在多个文件**：每个资源一个文件（如 `payment_intent.go`, `customer.go`）
- **时间戳字段使用 `string` 类型**：保留 API 原始格式，避免时区问题

### 测试
- **禁止 `t.Skip`**（当前违规：`finance/`, `issuing/`, `simulation/` 测试文件）
- **使用 `assert`/`require`**，禁止 `if` 裸判断
- **真实 Sandbox 环境**：无 mock，直接调用 API
- **统一使用 `sdk.SandboxURL`** + `sdk.WithDebug(true)`
- **每包独立 `TestMain`**：初始化共享 `testClient`

### JSON 策略
- **生产**：宽松模式（Go 默认），API 新字段不 break SDK
- **测试**：应使用 `DisallowUnknownFields` 及时发现字段缺失（当前未实现）
- **嵌套不稳定对象**：先用 `map[string]any` 兜底，后续再细化为强类型
- **基于真实 API 响应修复**，不依赖可能过时的文档

## ANTI-PATTERNS (THIS PROJECT)

### FATAL — 禁止
- `sync.Mutex` → 必须用 `deadlock.Mutex` / `deadlock.RWMutex`
- `t.Skip` → 禁止跳过测试
- 日志库二次包装 → 必须用全局实例
- `x-request-id` 注入到 header → 应在请求体中
- 造轮子 → 优先用第三方库

### ERROR — 应避免
- 函数超过 200 行（不含空行/注释）
- `map[string]interface{}` → 用 `map[string]any`
- API 结构体用 `time.Time` → 用 `string`（当前违规：`webhook/verify.go:99`）
- 纯 mock 测试 → 必须基于真实环境
- 过度工程 → 只写需求范围内的代码

### WARN — 注意
- JSON tag 命名：API 字段用 snake_case，Go 内部用驼峰
- 嵌套对象先用 `map[string]any`，收集真实 JSON 后再强类型化
- applepay / googlepay 暂不实现（需前端 SDK 提供 token）
- 修改相邻代码 → 只改你引入变更的部分

## UNIQUE STYLES

- **Domain-as-package 架构**：15 个业务域包在根级，非标准 Go 布局但语义清晰
- **`_examples/` 前缀约定**：示例程序位于下划线前缀目录
- **`sdk.Doer` 接口注入**：所有服务通过构造函数接收 HTTP 执行器
- **Bearer Token 自动管理**：首次换取 + 过期前 5 分钟预刷新 + 401 兜底重试
- **OpenSpec 规范驱动**：每个领域有独立 `spec.md`，变更流程为 Explore→Propose→Apply→Archive
- **`.opencode/skills/`**：5 个项目自定义技能（`api-doc-alignment-check` 等）

## COMMANDS

```bash
# L1: 静态分析
go vet ./...
golangci-lint run

# L2: 代码重复 + AST 检查
# 由 golang-code-style 脚本自动执行

# L3: 编译 + 测试
go build ./...
go test ./...

# 格式化
go fmt ./...
```

## NOTES

- **单模块**：无 monorepo，无 go.work，无 internal/
- **无 CI**：无 `.github/workflows/`，无 Makefile
- **测试依赖凭证**：必须设置 `AIRWALLEX_CLIENT_ID` + `AIRWALLEX_API_KEY` 环境变量
- **SDK 环境变量键名**：`EnvClientIDKey = "AIRWALLEX_CLIENT_ID"`, `EnvAPIKey = "AIRWALLEX_API_KEY"`
- **当前活跃变更**：`openspec/changes/full-api-doc-alignment-check/`
- **已知违规**：3 处 `t.Skip`、1 处 `time.Time` 在 API 结构体中、1 处 `map[string]interface{}` 待改为 `map[string]any`
