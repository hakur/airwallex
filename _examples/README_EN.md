# Airwallex Go SDK Examples

This directory contains 19 standalone, runnable example programs demonstrating all SDK business domains.

## Prerequisites

Create a `.env` file in the project root:

```bash
cp ../.env.example ../.env
# Edit ../.env with real credentials
```

## Running an Example

```bash
cd <example-directory>
go run main.go
```

## Example List

| Example | Module | Description |
|---|---|---|
| `quickstart` | core | Minimal: load credentials from .env, query balances |
| `payment_intent` | pa | PaymentIntent lifecycle: create → confirm → capture |
| `payout` | payouts | Create beneficiary + transfer |
| `platform` | core | Multi-merchant `OnBehalfOf` |
| `webhook_server` | webhook | HTTP server receiving and verifying webhooks |
| `fx_conversion` | fx | Rates and conversions |
| `issuing` | issuing | Cardholder + virtual card |
| `billing_invoice` | billing | Billing customer + invoice management |
| `global_account` | core | Global account creation and bank details |
| `subscription_recurring` | billing + pa | Monthly/yearly recurring auto-charge |
| `subscription_yearly_discount` | billing | Yearly subscription + first-year coupon discount |
| `payment_callback` | webhook + pa | Payment success callback: create → confirm → capture → webhook verification |
| `capability` | capability | Account capability queries and enablement |
| `confirmation` | confirmation | Confirmation letter generation |
| `finance` | finance | Financial transaction queries |
| `risk` | risk | Seller management, risk operations |
| `scale` | scale | Platform multi-account management |
| `simulation` | simulation | Sandbox simulation events |
| `spend` | spend | Bill, expense, purchase order management |
| `supporting` | supporting | File upload/download, connected store management |

## Choosing an Example by Business Domain

| Domain | Recommended Examples |
|---|---|
| Payment Acceptance | `payment_intent` |
| Payouts / Transfers | `payout` |
| FX / Currency Exchange | `fx_conversion` |
| Issuing | `issuing` |
| Billing & Subscriptions | `billing_invoice`, `subscription_recurring`, `subscription_yearly_discount` |
| Global Accounts | `global_account` |
| Platform / Multi-Merchant | `platform`, `scale` |
| Webhook Verification | `webhook_server` |
| Risk Management | `risk` |
| Sandbox Simulation | `simulation` |
| Spend Management | `spend` |
| Supporting Services | `supporting` |

## Conventions

- Each example is a self-contained `package main` with one `main.go` file
- Uses `sdk.WithBaseURL(sdk.SandboxURL)` for safety
- Error handling via `log/slog` — examples are not production code
- Pattern: create client → call API → print result
- Idempotent request IDs: `time.Now().Format("20060102150405")`
