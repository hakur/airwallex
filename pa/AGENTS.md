# Payment Acceptance (`pa`)

## OVERVIEW
Payment Acceptance, creating, managing, and confirming payments.

## STRUCTURE

| File | Purpose |
|------|---------|
| `service.go` | `Service` struct, `New(doer)` constructor |
| `payment_intent.go` | Create/Confirm/Capture/Get/Update/Cancel/List PaymentIntent |
| `customer.go` | Create/Get/Update/List Customer |
| `payment_method.go` | List/Attach/Detach PaymentMethod + typed input structs |
| `payment_attempts.go` | Get/List PaymentAttempt |
| `payment_disputes.go` | Get/List PaymentDispute |
| `refund.go` | Create/Get/List Refund |
| `payment_link.go` | Create/Get/List PaymentLink |
| `conversion_quotes.go` | Create/Get ConversionQuote |
| `config.go` | Payment method types, banks, currencies, reserve plan, Apple Pay domains |
| `customs_declarations.go` | Create/Get/Update/Redeclare/List CustomsDeclaration |
| `funds_splits.go` | Create/Get/List FundsSplit |
| `funds_split_reversals.go` | Create/Get/List FundsSplitReversal |
| `settlement_records.go` | Get/List SettlementRecord |
| `terminals.go` | POS terminals: CRUD + Activate/Deactivate/Terminate/ResetPassword/CancelOp/ProcessPayment |
| `pa_reference_data.go` | GetReferenceData, ListReferenceDataCountries/Currencies |
| `payment_consents.go` | Create/Get/Cancel/List PaymentConsent |
| `pa_test.go` | Sandbox tests: lifecycle tests for major resources |

## WHERE TO LOOK

| Task | Location | Notes |
|------|----------|-------|
| Add payment method | `payment_method.go` | `PaymentMethodInput` (request) vs `PaymentMethod` (response). 60+ unimplemented method types listed at bottom of file |
| Add terminal operation | `terminals.go` | 10 methods including POS lifecycle + activation/deactivation |
| Add webhook handler | outside `pa/` | `webhook/verify.go` or root `client.go` |
| Fix Sandbox 404s | `customs_declarations.go`, `funds_splits.go`, `funds_split_reversals.go` | Create works; Get/List return 404 in Sandbox |

## CONVENTIONS

- Input structs (Create/Update/Confirm/Capture) vs response structs — separate types
- `PaymentMethodInput` uses pointer + `omitempty` for optional payment method fields
- Query params built manually with `url.Values` when `req != nil` (see `terminals.go`, `config.go`)
- Nested complex objects: use `map[string]any` as fallback, refine later with real API responses
- All API paths prefix: `/api/v1/pa/...`

## ANTI-PATTERNS

- Do NOT add `applepay` or `googlepay` method structs — requires frontend SDK token, out of scope
- Do NOT create generic CRUD helper — each resource has unique schema and query param needs
- Do NOT use `time.Time` in API structs — use `string` (see root `AGENTS.md`)
