# _examples/

**20 standalone `package main` examples demonstrating SDK usage across all business domains.**

## STRUCTURE

| Directory | Domain | What it shows |
|---|---|---|
| `quickstart/` | core | Minimal: load from .env, query balances |
| `payment_intent/` | pa | PaymentIntent lifecycle (create, confirm, capture) |
| `payout/` | payouts | Create beneficiary + transfer |
| `platform/` | core | Multi-merchant `OnBehalfOf` |
| `webhook_server/` | webhook | HTTP server receiving and verifying webhooks |
| `fx_conversion/` | fx | Rates and conversions |
| `issuing/` | issuing | Cardholder + virtual card |
| `billing_invoice/` | billing | Create invoice |
| `global_account/` | core | Global account creation |
| `subscription_recurring/` | billing+pa | Recurring auto-charge |
| `capability/` | capability | Account capability queries |
| `confirmation/` | confirmation | Confirmation letters |
| `finance/` | finance | Financial transactions |
| `risk/` | risk | Sellers |
| `scale/` | scale | Platform accounts |
| `simulation/` | simulation | Sandbox simulations |
| `spend/` | spend | Bills and expenses |
| `supporting/` | supporting | Supporting services |
| `subscription_yearly_discount/` | billing | Yearly subscription + coupon discount |
| `payment_callback/` | webhook+pa | PaymentIntent lifecycle + webhook callback demo |

## WHERE TO LOOK

| Task | Location | Notes |
|---|---|---|
| Add new example | Create new subdir, single `main.go` | Follow `quickstart` pattern |
| Minimal example | `quickstart/main.go` | ~30 lines, best starting point |

## CONVENTIONS

- Each example is a self-contained `package main` with exactly one `main.go`
- All use `sdk.WithBaseURL(sdk.SandboxURL)` for safety
- Error handling: `log.Fatal(err)` — examples are not production code
- Context: uses `context.Background()`
- Pattern: create client → call API → print result

## ANTI-PATTERNS

- Do NOT add complex logic — examples should be copy-paste ready (under 100 lines)
- Do NOT require setup beyond `.env` with credentials
- Do NOT add tests to examples — tests live in domain packages
