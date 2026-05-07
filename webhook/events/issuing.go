// Package events provides typed webhook event structures for the issuing domain.
// Issuing 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/issuing.md
//
// 事件映射表:
//
//	issuing.card.pending                          → IssuingCardPendingEvent                      (Data: IssuingCardEventData)
//	issuing.card.failed                           → IssuingCardFailedEvent                       (Data: IssuingCardEventData)
//	issuing.card.active                           → IssuingCardActiveEvent                       (Data: IssuingCardEventData)
//	issuing.card.inactive                         → IssuingCardInactiveEvent                     (Data: IssuingCardEventData)
//	issuing.card.blocked                          → IssuingCardBlockedEvent                      (Data: IssuingCardEventData)
//	issuing.card.lost                             → IssuingCardLostEvent                         (Data: IssuingCardEventData)
//	issuing.card.stolen                           → IssuingCardStolenEvent                       (Data: IssuingCardEventData)
//	issuing.card.closed                           → IssuingCardClosedEvent                       (Data: IssuingCardEventData)
//	issuing.card.expired                          → IssuingCardExpiredEvent                      (Data: IssuingCardEventData)
//	issuing.cardholder.incomplete                 → IssuingCardholderIncompleteEvent             (Data: IssuingCardholderEventData)
//	issuing.cardholder.pending                    → IssuingCardholderPendingEvent                (Data: IssuingCardholderEventData)
//	issuing.cardholder.ready                      → IssuingCardholderReadyEvent                  (Data: IssuingCardholderEventData)
//	issuing.cardholder.disabled                   → IssuingCardholderDisabledEvent               (Data: IssuingCardholderEventData)
//	issuing.cardholder.deleted                    → IssuingCardholderDeletedEvent                (Data: IssuingCardholderEventData)
//	issuing.transaction.succeeded                 → IssuingTransactionSucceededEvent             (Data: IssuingTransactionEventData)
//	issuing.transaction.failed                    → IssuingTransactionFailedEvent                (Data: IssuingTransactionEventData)
//	issuing.card_transaction_lifecycle.created    → IssuingCardTransactionLifecycleCreatedEvent  (Data: IssuingCardTransactionLifecycleEventData)
//	issuing.card_transaction_lifecycle.modified   → IssuingCardTransactionLifecycleModifiedEvent (Data: IssuingCardTransactionLifecycleEventData)
//	issuing.card_transaction.authorized           → IssuingCardTransactionAuthorizedEvent        (Data: IssuingCardTransactionStatusEventData)
//	issuing.card_transaction.verified             → IssuingCardTransactionVerifiedEvent          (Data: IssuingCardTransactionStatusEventData)
//	issuing.card_transaction.cleared              → IssuingCardTransactionClearedEvent           (Data: IssuingCardTransactionStatusEventData)
//	issuing.card_transaction.reversed             → IssuingCardTransactionReversedEvent          (Data: IssuingCardTransactionStatusEventData)
//	issuing.card_transaction.expired              → IssuingCardTransactionExpiredEvent           (Data: IssuingCardTransactionStatusEventData)
//	issuing.card_transaction.declined             → IssuingCardTransactionDeclinedEvent          (Data: IssuingCardTransactionStatusEventData)
//	issuing.card_transaction.modified             → IssuingCardTransactionModifiedEvent          (Data: IssuingCardTransactionStatusEventData)
//	issuing.card_transaction_event.success        → IssuingCardTransactionEventSuccessEvent      (Data: IssuingCardTransactionEventData)
//	issuing.card_transaction_event.failed         → IssuingCardTransactionEventFailedEvent       (Data: IssuingCardTransactionEventData)
//	issuing.card_transaction_event.enriched       → IssuingCardTransactionEventEnrichedEvent     (Data: IssuingCardTransactionEventData)
//	issuing.transaction_dispute.created           → IssuingTransactionDisputeCreatedEvent        (Data: IssuingTransactionDisputeEventData)
//	issuing.transaction_dispute.expired           → IssuingTransactionDisputeExpiredEvent        (Data: IssuingTransactionDisputeEventData)
//	issuing.transaction_dispute.submitted         → IssuingTransactionDisputeSubmittedEvent      (Data: IssuingTransactionDisputeEventData)
//	issuing.transaction_dispute.rejected          → IssuingTransactionDisputeRejectedEvent       (Data: IssuingTransactionDisputeEventData)
//	issuing.transaction_dispute.canceled          → IssuingTransactionDisputeCanceledEvent       (Data: IssuingTransactionDisputeEventData)
//	issuing.transaction_dispute.accepted          → IssuingTransactionDisputeAcceptedEvent       (Data: IssuingTransactionDisputeEventData)
//	issuing.transaction_dispute.won               → IssuingTransactionDisputeWonEvent            (Data: IssuingTransactionDisputeEventData)
//	issuing.transaction_dispute.lost              → IssuingTransactionDisputeLostEvent           (Data: IssuingTransactionDisputeEventData)
//	issuing.reissue.succeeded                     → IssuingReissueSucceededEvent                 (Data: IssuingReissueEventData)
//	issuing.card.low_remaining_transaction_limit  → IssuingCardLowRemainingTransactionLimitEvent (Data: IssuingCardAlertEventData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Card Status Events ---

// IssuingCardPendingEvent represents the issuing.card.pending webhook event.
type IssuingCardPendingEvent struct {
	Event
	Data IssuingCardEventData `json:"data"`
}

// IssuingCardFailedEvent represents the issuing.card.failed webhook event.
type IssuingCardFailedEvent struct {
	Event
	Data IssuingCardEventData `json:"data"`
}

// IssuingCardActiveEvent represents the issuing.card.active webhook event.
type IssuingCardActiveEvent struct {
	Event
	Data IssuingCardEventData `json:"data"`
}

// IssuingCardInactiveEvent represents the issuing.card.inactive webhook event.
type IssuingCardInactiveEvent struct {
	Event
	Data IssuingCardEventData `json:"data"`
}

// IssuingCardBlockedEvent represents the issuing.card.blocked webhook event.
type IssuingCardBlockedEvent struct {
	Event
	Data IssuingCardEventData `json:"data"`
}

// IssuingCardLostEvent represents the issuing.card.lost webhook event.
type IssuingCardLostEvent struct {
	Event
	Data IssuingCardEventData `json:"data"`
}

// IssuingCardStolenEvent represents the issuing.card.stolen webhook event.
type IssuingCardStolenEvent struct {
	Event
	Data IssuingCardEventData `json:"data"`
}

// IssuingCardClosedEvent represents the issuing.card.closed webhook event.
type IssuingCardClosedEvent struct {
	Event
	Data IssuingCardEventData `json:"data"`
}

// IssuingCardExpiredEvent represents the issuing.card.expired webhook event.
type IssuingCardExpiredEvent struct {
	Event
	Data IssuingCardEventData `json:"data"`
}

// IssuingCardEventData contains card status information (flat structure).
type IssuingCardEventData struct {
	CardID string `json:"card_id,omitempty"`
	Status string `json:"status,omitempty"`
}

// --- Cardholder Status Events ---

// IssuingCardholderIncompleteEvent represents the issuing.cardholder.incomplete webhook event.
type IssuingCardholderIncompleteEvent struct {
	Event
	Data IssuingCardholderEventData `json:"data"`
}

// IssuingCardholderPendingEvent represents the issuing.cardholder.pending webhook event.
type IssuingCardholderPendingEvent struct {
	Event
	Data IssuingCardholderEventData `json:"data"`
}

// IssuingCardholderReadyEvent represents the issuing.cardholder.ready webhook event.
type IssuingCardholderReadyEvent struct {
	Event
	Data IssuingCardholderEventData `json:"data"`
}

// IssuingCardholderDisabledEvent represents the issuing.cardholder.disabled webhook event.
type IssuingCardholderDisabledEvent struct {
	Event
	Data IssuingCardholderEventData `json:"data"`
}

// IssuingCardholderDeletedEvent represents the issuing.cardholder.deleted webhook event.
type IssuingCardholderDeletedEvent struct {
	Event
	Data IssuingCardholderEventData `json:"data"`
}

// IssuingCardholderEventData contains cardholder status information (flat structure).
type IssuingCardholderEventData struct {
	CardholderID string `json:"cardholder_id,omitempty"`
	Status       string `json:"status,omitempty"`
}

// --- Transaction Events ---

// IssuingTransactionSucceededEvent represents the issuing.transaction.succeeded webhook event.
type IssuingTransactionSucceededEvent struct {
	Event
	Data IssuingTransactionEventData `json:"data"`
}

// IssuingTransactionFailedEvent represents the issuing.transaction.failed webhook event.
type IssuingTransactionFailedEvent struct {
	Event
	Data IssuingTransactionEventData `json:"data"`
}

// IssuingTransactionEventData contains transaction information (flat structure).
// Based on official docs: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/issuing.md
type IssuingTransactionEventData struct {
	TransactionID          string   `json:"transaction_id"`
	TransactionType        string   `json:"transaction_type"`
	TransactionAmount      float64  `json:"transaction_amount"`
	TransactionCurrency    string   `json:"transaction_currency"`
	BillingAmount          float64  `json:"billing_amount,omitempty"`
	BillingCurrency        string   `json:"billing_currency,omitempty"`
	Status                 string   `json:"status"`
	CardID                 string   `json:"card_id"`
	CardNickname           string   `json:"card_nickname,omitempty"`
	MaskedCardNumber       string   `json:"masked_card_number,omitempty"`
	AuthCode               string   `json:"auth_code,omitempty"`
	RetrievalRef           string   `json:"retrieval_ref,omitempty"`
	NetworkTransactionID   string   `json:"network_transaction_id,omitempty"`
	AcquiringInstitutionID string   `json:"acquiring_institution_id,omitempty"`
	DigitalWalletTokenID   string   `json:"digital_wallet_token_id,omitempty"`
	FailureReason          string   `json:"failure_reason,omitempty"`
	MatchedAuthorizations  []string `json:"matched_authorizations,omitempty"`
	PostedDate             string   `json:"posted_date"`
	TransactionDate        string   `json:"transaction_date"`
	// ClientData is client-specific data. Official docs show null without type definition.
	// ClientData 客户端特定数据。官方文档显示为 null，未定义类型。
	ClientData map[string]any  `json:"client_data,omitempty"`
	Merchant   IssuingMerchant `json:"merchant,omitempty"`
}

// IssuingMerchant represents merchant information in issuing events.
type IssuingMerchant struct {
	CategoryCode string `json:"category_code,omitempty"`
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
	Identifier   string `json:"identifier,omitempty"`
	Name         string `json:"name,omitempty"`
}

// --- Transaction Lifecycle Events ---

// IssuingCardTransactionLifecycleCreatedEvent represents the issuing.card_transaction_lifecycle.created webhook event.
type IssuingCardTransactionLifecycleCreatedEvent struct {
	Event
	Data IssuingCardTransactionLifecycleEventData `json:"data"`
}

// IssuingCardTransactionLifecycleModifiedEvent represents the issuing.card_transaction_lifecycle.modified webhook event.
type IssuingCardTransactionLifecycleModifiedEvent struct {
	Event
	Data IssuingCardTransactionLifecycleEventData `json:"data"`
}

// IssuingCardTransactionLifecycleEventData contains transaction lifecycle information.
type IssuingCardTransactionLifecycleEventData struct {
	LifecycleID   string `json:"lifecycle_id,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	Status        string `json:"status,omitempty"`
}

// --- Card Transaction Status Events ---

// IssuingCardTransactionAuthorizedEvent represents the issuing.card_transaction.authorized webhook event.
type IssuingCardTransactionAuthorizedEvent struct {
	Event
	Data IssuingCardTransactionStatusEventData `json:"data"`
}

// IssuingCardTransactionVerifiedEvent represents the issuing.card_transaction.verified webhook event.
type IssuingCardTransactionVerifiedEvent struct {
	Event
	Data IssuingCardTransactionStatusEventData `json:"data"`
}

// IssuingCardTransactionClearedEvent represents the issuing.card_transaction.cleared webhook event.
type IssuingCardTransactionClearedEvent struct {
	Event
	Data IssuingCardTransactionStatusEventData `json:"data"`
}

// IssuingCardTransactionReversedEvent represents the issuing.card_transaction.reversed webhook event.
type IssuingCardTransactionReversedEvent struct {
	Event
	Data IssuingCardTransactionStatusEventData `json:"data"`
}

// IssuingCardTransactionExpiredEvent represents the issuing.card_transaction.expired webhook event.
type IssuingCardTransactionExpiredEvent struct {
	Event
	Data IssuingCardTransactionStatusEventData `json:"data"`
}

// IssuingCardTransactionDeclinedEvent represents the issuing.card_transaction.declined webhook event.
type IssuingCardTransactionDeclinedEvent struct {
	Event
	Data IssuingCardTransactionStatusEventData `json:"data"`
}

// IssuingCardTransactionModifiedEvent represents the issuing.card_transaction.modified webhook event.
type IssuingCardTransactionModifiedEvent struct {
	Event
	Data IssuingCardTransactionStatusEventData `json:"data"`
}

// IssuingCardTransactionStatusEventData contains card transaction status information.
type IssuingCardTransactionStatusEventData struct {
	TransactionID string `json:"transaction_id,omitempty"`
	Status        string `json:"status,omitempty"`
}

// --- Card Transaction Events ---

// IssuingCardTransactionEventSuccessEvent represents the issuing.card_transaction_event.success webhook event.
type IssuingCardTransactionEventSuccessEvent struct {
	Event
	Data IssuingCardTransactionEventData `json:"data"`
}

// IssuingCardTransactionEventFailedEvent represents the issuing.card_transaction_event.failed webhook event.
type IssuingCardTransactionEventFailedEvent struct {
	Event
	Data IssuingCardTransactionEventData `json:"data"`
}

// IssuingCardTransactionEventEnrichedEvent represents the issuing.card_transaction_event.enriched webhook event.
type IssuingCardTransactionEventEnrichedEvent struct {
	Event
	Data IssuingCardTransactionEventData `json:"data"`
}

// IssuingCardTransactionEventData contains card transaction event information.
type IssuingCardTransactionEventData struct {
	EventID       string `json:"event_id,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	Status        string `json:"status,omitempty"`
}

// --- Dispute Events ---

// IssuingTransactionDisputeCreatedEvent represents the issuing.transaction_dispute.created webhook event.
type IssuingTransactionDisputeCreatedEvent struct {
	Event
	Data IssuingTransactionDisputeEventData `json:"data"`
}

// IssuingTransactionDisputeExpiredEvent represents the issuing.transaction_dispute.expired webhook event.
type IssuingTransactionDisputeExpiredEvent struct {
	Event
	Data IssuingTransactionDisputeEventData `json:"data"`
}

// IssuingTransactionDisputeSubmittedEvent represents the issuing.transaction_dispute.submitted webhook event.
type IssuingTransactionDisputeSubmittedEvent struct {
	Event
	Data IssuingTransactionDisputeEventData `json:"data"`
}

// IssuingTransactionDisputeRejectedEvent represents the issuing.transaction_dispute.rejected webhook event.
type IssuingTransactionDisputeRejectedEvent struct {
	Event
	Data IssuingTransactionDisputeEventData `json:"data"`
}

// IssuingTransactionDisputeCanceledEvent represents the issuing.transaction_dispute.canceled webhook event.
type IssuingTransactionDisputeCanceledEvent struct {
	Event
	Data IssuingTransactionDisputeEventData `json:"data"`
}

// IssuingTransactionDisputeAcceptedEvent represents the issuing.transaction_dispute.accepted webhook event.
type IssuingTransactionDisputeAcceptedEvent struct {
	Event
	Data IssuingTransactionDisputeEventData `json:"data"`
}

// IssuingTransactionDisputeWonEvent represents the issuing.transaction_dispute.won webhook event.
type IssuingTransactionDisputeWonEvent struct {
	Event
	Data IssuingTransactionDisputeEventData `json:"data"`
}

// IssuingTransactionDisputeLostEvent represents the issuing.transaction_dispute.lost webhook event.
type IssuingTransactionDisputeLostEvent struct {
	Event
	Data IssuingTransactionDisputeEventData `json:"data"`
}

// IssuingTransactionDisputeEventData contains dispute information.
type IssuingTransactionDisputeEventData struct {
	DisputeID     string `json:"dispute_id,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	Status        string `json:"status,omitempty"`
}

// --- Reissue Events ---

// IssuingReissueSucceededEvent represents the issuing.reissue.succeeded webhook event.
type IssuingReissueSucceededEvent struct {
	Event
	Data IssuingReissueEventData `json:"data"`
}

// IssuingReissueEventData contains card reissue information.
type IssuingReissueEventData struct {
	AccountID           string `json:"account_id,omitempty"`
	AccountBusinessName string `json:"account_business_name,omitempty"`
	CardID              string `json:"card_id,omitempty"`
	CardholderFirstName string `json:"cardholder_first_name,omitempty"`
	OldMaskedCardNumber string `json:"old_masked_card_number,omitempty"`
	NewMaskedCardNumber string `json:"new_masked_card_number,omitempty"`
	ReissueReason       string `json:"reissue_reason,omitempty"`
}

// --- Card Alert Events ---

// IssuingCardLowRemainingTransactionLimitEvent represents the issuing.card.low_remaining_transaction_limit webhook event.
type IssuingCardLowRemainingTransactionLimitEvent struct {
	Event
	Data IssuingCardAlertEventData `json:"data"`
}

// IssuingCardAlertEventData contains card alert information.
type IssuingCardAlertEventData struct {
	CardID            string  `json:"card_id,omitempty"`
	RemainingLimit    float64 `json:"remaining_limit,omitempty"`
	Threshold         float64 `json:"threshold,omitempty"`
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
}
