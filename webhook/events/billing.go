// Package events provides typed webhook event structures for the billing domain.
// Billing 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/billing.md
//
// 事件映射表:
//
//	subscription.created                   → SubscriptionCreatedEvent           (Data: billing.Subscription)
//	subscription.in_trial                  → SubscriptionInTrialEvent           (Data: billing.Subscription)
//	subscription.active                    → SubscriptionActiveEvent            (Data: billing.Subscription)
//	subscription.unpaid                    → SubscriptionUnpaidEvent            (Data: billing.Subscription)
//	subscription.modified                  → SubscriptionModifiedEvent          (Data: billing.Subscription)
//	subscription.cancelled                 → SubscriptionCancelledEvent         (Data: billing.Subscription)
//	invoice.created                        → InvoiceCreatedEvent                (Data: billing.Invoice)
//	invoice.finalized                      → InvoiceFinalizedEvent              (Data: billing.Invoice)
//	invoice.voided                         → InvoiceVoidedEvent                 (Data: billing.Invoice)
//	invoice.updated                        → InvoiceUpdatedEvent                (Data: billing.Invoice)
//	invoice.payment.paid                   → InvoicePaymentPaidEvent            (Data: billing.Invoice)
//	billing_transaction.created            → BillingTransactionCreatedEvent     (Data: billing.BillingTransaction)
//	billing_transaction.succeeded          → BillingTransactionSucceededEvent   (Data: billing.BillingTransaction)
//	billing_transaction.cancelled          → BillingTransactionCancelledEvent   (Data: billing.BillingTransaction)
//	billing_checkout.created               → BillingCheckoutCreatedEvent        (Data: billing.Checkout)
//	billing_checkout.cancelled             → BillingCheckoutCancelledEvent      (Data: billing.Checkout)
//	billing_checkout.completed             → BillingCheckoutCompletedEvent      (Data: billing.Checkout)
//	usage_event.aggregation_failed         → UsageEventAggregationFailedEvent   (Data: UsageEventAggregationFailedData)
//	credit_note.created                    → CreditNoteCreatedEvent             (Data: billing.CreditNote)
//	credit_note.finalized                  → CreditNoteFinalizedEvent           (Data: billing.CreditNote)
//	credit_note.voided                     → CreditNoteVoidedEvent              (Data: billing.CreditNote)
//	subscription.updated (legacy)          → SubscriptionUpdatedEvent           (Data: BillingSubscriptionEventData)
//	invoice.sent (legacy)                  → InvoiceSentEvent                   (Data: BillingInvoiceEventData)
//	invoice.paid (legacy)                  → InvoicePaidEvent                   (Data: BillingInvoiceEventData)
//	invoice.payment_failed (legacy)        → InvoicePaymentFailedEvent          (Data: BillingInvoiceEventData)
//	invoice.payment_attempt_failed (legacy) → InvoicePaymentAttemptFailedEvent  (Data: BillingInvoiceEventData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

import (
	"github.com/hakur/airwallex/billing"
	"github.com/hakur/airwallex/sdk"
)

// --- Subscription Events (API version 2025-06-16 and later) ---

// SubscriptionCreatedEvent represents the subscription.created webhook event.
type SubscriptionCreatedEvent struct {
	Event
	Data billing.Subscription `json:"data"`
}

// SubscriptionInTrialEvent represents the subscription.in_trial webhook event.
type SubscriptionInTrialEvent struct {
	Event
	Data billing.Subscription `json:"data"`
}

// SubscriptionActiveEvent represents the subscription.active webhook event.
type SubscriptionActiveEvent struct {
	Event
	Data billing.Subscription `json:"data"`
}

// SubscriptionUnpaidEvent represents the subscription.unpaid webhook event.
type SubscriptionUnpaidEvent struct {
	Event
	Data billing.Subscription `json:"data"`
}

// SubscriptionModifiedEvent represents the subscription.modified webhook event.
type SubscriptionModifiedEvent struct {
	Event
	Data billing.Subscription `json:"data"`
}

// SubscriptionCancelledEvent represents the subscription.cancelled webhook event.
type SubscriptionCancelledEvent struct {
	Event
	Data billing.Subscription `json:"data"`
}

// --- Invoice Events (API version 2025-06-16 and later) ---

// InvoiceCreatedEvent represents the invoice.created webhook event.
type InvoiceCreatedEvent struct {
	Event
	Data billing.Invoice `json:"data"`
}

// InvoiceFinalizedEvent represents the invoice.finalized webhook event.
type InvoiceFinalizedEvent struct {
	Event
	Data billing.Invoice `json:"data"`
}

// InvoiceVoidedEvent represents the invoice.voided webhook event.
type InvoiceVoidedEvent struct {
	Event
	Data billing.Invoice `json:"data"`
}

// InvoiceUpdatedEvent represents the invoice.updated webhook event.
type InvoiceUpdatedEvent struct {
	Event
	Data billing.Invoice `json:"data"`
}

// InvoicePaymentPaidEvent represents the invoice.payment.paid webhook event.
type InvoicePaymentPaidEvent struct {
	Event
	Data billing.Invoice `json:"data"`
}

// --- Billing Transaction Events (API version 2025-06-16 and later) ---

// BillingTransactionCreatedEvent represents the billing_transaction.created webhook event.
type BillingTransactionCreatedEvent struct {
	Event
	Data billing.BillingTransaction `json:"data"`
}

// BillingTransactionSucceededEvent represents the billing_transaction.succeeded webhook event.
type BillingTransactionSucceededEvent struct {
	Event
	Data billing.BillingTransaction `json:"data"`
}

// BillingTransactionCancelledEvent represents the billing_transaction.cancelled webhook event.
type BillingTransactionCancelledEvent struct {
	Event
	Data billing.BillingTransaction `json:"data"`
}

// --- Billing Checkout Events (API version 2025-06-16 and later) ---

// BillingCheckoutCreatedEvent represents the billing_checkout.created webhook event.
type BillingCheckoutCreatedEvent struct {
	Event
	Data billing.Checkout `json:"data"`
}

// BillingCheckoutCancelledEvent represents the billing_checkout.cancelled webhook event.
type BillingCheckoutCancelledEvent struct {
	Event
	Data billing.Checkout `json:"data"`
}

// BillingCheckoutCompletedEvent represents the billing_checkout.completed webhook event.
type BillingCheckoutCompletedEvent struct {
	Event
	Data billing.Checkout `json:"data"`
}

// --- Usage Event Aggregation Failed Event ---

// UsageEventAggregationFailedEvent represents the usage_event.aggregation_failed webhook event.
type UsageEventAggregationFailedEvent struct {
	Event
	Data UsageEventAggregationFailedData `json:"data"`
}

// UsageEventAggregationFailedData describes a summary of failure events during the aggregation period.
type UsageEventAggregationFailedData struct {
	// FromIngestedAt is the start timestamp of the time range when events were ingested into the system.
	FromIngestedAt string `json:"from_ingested_at"`
	// ToIngestedAt is the end timestamp of the time range when events were ingested into the system.
	ToIngestedAt string `json:"to_ingested_at"`
	// TotalFailedEventCount is the total number of events that failed to be processed during the specified time range.
	TotalFailedEventCount int `json:"total_failed_event_count"`
	// ErrorSamples is an array containing sample error cases from the failed events.
	ErrorSamples []UsageEventAggregationErrorSample `json:"error_samples"`
}

// UsageEventAggregationErrorSample represents a sample error case from the failed events.
type UsageEventAggregationErrorSample struct {
	// Error contains details about the error of the event ingestion.
	Error sdk.APIError `json:"error"`
	// MerchantEventID is a unique identifier for the event sample, provided by the user.
	MerchantEventID string `json:"merchant_event_id"`
}

// --- Credit Note Events (API version 2025-06-16 and later) ---

// CreditNoteCreatedEvent represents the credit_note.created webhook event.
type CreditNoteCreatedEvent struct {
	Event
	Data billing.CreditNote `json:"data"`
}

// CreditNoteFinalizedEvent represents the credit_note.finalized webhook event.
type CreditNoteFinalizedEvent struct {
	Event
	Data billing.CreditNote `json:"data"`
}

// CreditNoteVoidedEvent represents the credit_note.voided webhook event.
type CreditNoteVoidedEvent struct {
	Event
	Data billing.CreditNote `json:"data"`
}

// --- Legacy Events (API version 2025-04-25 and before) ---
//
// 以下事件适用于 API 版本 2025-04-25 及更早版本。
// 这些事件的 data 字段包含 object 包装，与新版本的扁平结构不同。

// BillingSubscriptionEventData wraps the subscription object for legacy webhook events.
type BillingSubscriptionEventData struct {
	Object billing.Subscription `json:"object"`
}

// BillingInvoiceEventData wraps the invoice object for legacy webhook events.
type BillingInvoiceEventData struct {
	Object billing.Invoice `json:"object"`
}

// SubscriptionUpdatedEvent represents the legacy subscription.updated webhook event.
// 适用于 API 版本 2025-04-25 及更早版本。
type SubscriptionUpdatedEvent struct {
	Event
	Data BillingSubscriptionEventData `json:"data"`
}

// InvoiceSentEvent represents the legacy invoice.sent webhook event.
// 适用于 API 版本 2025-04-25 及更早版本。
type InvoiceSentEvent struct {
	Event
	Data BillingInvoiceEventData `json:"data"`
}

// InvoicePaidEvent represents the legacy invoice.paid webhook event.
// 适用于 API 版本 2025-04-25 及更早版本。
type InvoicePaidEvent struct {
	Event
	Data BillingInvoiceEventData `json:"data"`
}

// InvoicePaymentFailedEvent represents the legacy invoice.payment_failed webhook event.
// 适用于 API 版本 2025-04-25 及更早版本。
type InvoicePaymentFailedEvent struct {
	Event
	Data BillingInvoiceEventData `json:"data"`
}

// InvoicePaymentAttemptFailedEvent represents the legacy invoice.payment_attempt_failed webhook event.
// 适用于 API 版本 2025-04-25 及更早版本。
type InvoicePaymentAttemptFailedEvent struct {
	Event
	Data BillingInvoiceEventData `json:"data"`
}
