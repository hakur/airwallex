# sdk — Shared Infrastructure

## OVERVIEW
Shared SDK infrastructure: types, errors, options, and constants consumed by all domain services.

## STRUCTURE

- `sdk.go` — Core constants (SandboxURL, ServerURL, APIVersion), Doer interface, env loading
- `options.go` — ClientOptions + With* option functions (WithBaseURL, WithDebug, WithOnBehalfOf)
- `request_option.go` — Per-request options (WithRequestOnBehalfOf)
- `errors.go` — APIError struct + ErrorCode constants + Is* predicates
- `currency.go` — Currency string type + 36 currency constants
- `country.go` — CountryCode string type + 38 country constants
- `types.go` — Shared types (Money, Address, ListOptions, ListResult[T], Pagination)

## WHERE TO LOOK

| Task | Location | Notes |
|------|----------|-------|
| Add SDK option | `options.go` | Follow `WithBaseURL` pattern |
| Add request option | `request_option.go` | Resty-based, sets header |
| Add error code | `errors.go` | Add `ErrorCode*` constant + `Is*` predicate |
| Add currency/country | `currency.go` / `country.go` | String type alias |
| Add shared type | `types.go` | Only cross-domain types |
| Generate request ID | `sdk.go:GenerateRequestID()` | UUIDv4 fallback to timestamp |

## CONVENTIONS

- All exported types — no internal/private SDK types
- Error codes: string constant + `Is*` predicate per code
- Options: functional options pattern `func(*ClientOptions)` in `options.go`
- Request options: `func(*resty.Request)` in `request_option.go`
- `ListResult[T]` generic for paginated API responses
- `Money` + `Address` reused across multiple domain packages

## ANTI-PATTERNS

- Do NOT add domain-specific types here — they go in their domain package
- Do NOT import domain packages (pa, billing, etc.) — this is bottom-level
- Do NOT add business logic — this package is pure types and utilities
