package webhook

// EventName represents an Airwallex webhook event type name.
// EventName 表示 Airwallex webhook 事件类型名称。
type EventName = string

// --- Online Payments: PaymentIntent ---
const (
	// EventNamePaymentIntentCreated occurs when a PaymentIntent has been created.
	// EventNamePaymentIntentCreated 对应结构体: PaymentIntentCreatedEvent (Data: PaymentIntentEventData)
	EventNamePaymentIntentCreated EventName = "payment_intent.created"
	// EventNamePaymentIntentRequiresPaymentMethod occurs when a PaymentIntent requires a payment method.
	// EventNamePaymentIntentRequiresPaymentMethod 对应结构体: PaymentIntentRequiresPaymentMethodEvent (Data: PaymentIntentEventData)
	EventNamePaymentIntentRequiresPaymentMethod EventName = "payment_intent.requires_payment_method"
	// EventNamePaymentIntentUpdated occurs when a PaymentIntent has been updated.
	// EventNamePaymentIntentUpdated 对应结构体: PaymentIntentUpdatedEvent (Data: PaymentIntentEventData)
	EventNamePaymentIntentUpdated EventName = "payment_intent.updated"
	// EventNamePaymentIntentRequiresCapture occurs when merchant capture is required.
	// EventNamePaymentIntentRequiresCapture 对应结构体: PaymentIntentRequiresCaptureEvent (Data: PaymentIntentEventData)
	EventNamePaymentIntentRequiresCapture EventName = "payment_intent.requires_capture"
	// EventNamePaymentIntentRequiresCustomerAction occurs when additional customer action is required.
	// EventNamePaymentIntentRequiresCustomerAction 对应结构体: PaymentIntentRequiresCustomerActionEvent (Data: PaymentIntentEventData)
	EventNamePaymentIntentRequiresCustomerAction EventName = "payment_intent.requires_customer_action"
	// EventNamePaymentIntentPending occurs when the payment request has been accepted and is waiting for final result.
	// EventNamePaymentIntentPending 对应结构体: PaymentIntentPendingEvent (Data: PaymentIntentEventData)
	EventNamePaymentIntentPending EventName = "payment_intent.pending"
	// EventNamePaymentIntentPendingReview occurs when the payment request is authorized but undergoing risk review.
	// EventNamePaymentIntentPendingReview 对应结构体: PaymentIntentPendingReviewEvent (Data: PaymentIntentEventData)
	EventNamePaymentIntentPendingReview EventName = "payment_intent.pending_review"
	// EventNamePaymentIntentSucceeded occurs when a PaymentIntent has been fulfilled.
	// EventNamePaymentIntentSucceeded 对应结构体: PaymentIntentSucceededEvent (Data: PaymentIntentEventData)
	EventNamePaymentIntentSucceeded EventName = "payment_intent.succeeded"
	// EventNamePaymentIntentCancelled occurs when a PaymentIntent has been cancelled.
	// EventNamePaymentIntentCancelled 对应结构体: PaymentIntentCancelledEvent (Data: PaymentIntentEventData)
	EventNamePaymentIntentCancelled EventName = "payment_intent.cancelled"
)

// --- Online Payments: PaymentAttempt ---
const (
	// EventNamePaymentAttemptReceived occurs when a PaymentAttempt has been received.
	// EventNamePaymentAttemptReceived 对应结构体: PaymentAttemptReceivedEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptReceived EventName = "payment_attempt.received"
	// EventNamePaymentAttemptAuthenticationFailed occurs when authentication has failed.
	// EventNamePaymentAttemptAuthenticationFailed 对应结构体: PaymentAttemptAuthenticationFailedEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptAuthenticationFailed EventName = "payment_attempt.authentication_failed"
	// EventNamePaymentAttemptAuthenticationRedirected occurs when redirected for authentication.
	// EventNamePaymentAttemptAuthenticationRedirected 对应结构体: PaymentAttemptAuthenticationRedirectedEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptAuthenticationRedirected EventName = "payment_attempt.authentication_redirected"
	// EventNamePaymentAttemptPendingAuthorization occurs when the payment request is accepted and waiting for final result.
	// EventNamePaymentAttemptPendingAuthorization 对应结构体: PaymentAttemptPendingAuthorizationEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptPendingAuthorization EventName = "payment_attempt.pending_authorization"
	// EventNamePaymentAttemptAuthorizationFailed occurs when authorization has failed.
	// EventNamePaymentAttemptAuthorizationFailed 对应结构体: PaymentAttemptAuthorizationFailedEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptAuthorizationFailed EventName = "payment_attempt.authorization_failed"
	// EventNamePaymentAttemptAuthorized occurs when the PaymentAttempt has been authorized and is waiting to be captured.
	// EventNamePaymentAttemptAuthorized 对应结构体: PaymentAttemptAuthorizedEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptAuthorized EventName = "payment_attempt.authorized"
	// EventNamePaymentAttemptCaptureRequested occurs when the PaymentAttempt has been requested for capture.
	// EventNamePaymentAttemptCaptureRequested 对应结构体: PaymentAttemptCaptureRequestedEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptCaptureRequested EventName = "payment_attempt.capture_requested"
	// EventNamePaymentAttemptSettled occurs when funds have been received from the acquirer.
	// EventNamePaymentAttemptSettled 对应结构体: PaymentAttemptSettledEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptSettled EventName = "payment_attempt.settled"
	// EventNamePaymentAttemptPaid occurs when funds have been paid to the merchant's wallet.
	// EventNamePaymentAttemptPaid 对应结构体: PaymentAttemptPaidEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptPaid EventName = "payment_attempt.paid"
	// EventNamePaymentAttemptCancelled occurs when the PaymentAttempt has been cancelled.
	// EventNamePaymentAttemptCancelled 对应结构体: PaymentAttemptCancelledEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptCancelled EventName = "payment_attempt.cancelled"
	// EventNamePaymentAttemptExpired occurs when the PaymentAttempt has expired.
	// EventNamePaymentAttemptExpired 对应结构体: PaymentAttemptExpiredEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptExpired EventName = "payment_attempt.expired"
	// EventNamePaymentAttemptRiskDeclined occurs when the PaymentAttempt failed risk screening.
	// EventNamePaymentAttemptRiskDeclined 对应结构体: PaymentAttemptRiskDeclinedEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptRiskDeclined EventName = "payment_attempt.risk_declined"
	// EventNamePaymentAttemptFailedToProcess occurs when the PaymentAttempt failed to be processed.
	// EventNamePaymentAttemptFailedToProcess 对应结构体: PaymentAttemptFailedToProcessEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptFailedToProcess EventName = "payment_attempt.failed_to_process"
	// EventNamePaymentAttemptCaptureFailed occurs when capture has failed.
	// EventNamePaymentAttemptCaptureFailed 对应结构体: PaymentAttemptCaptureFailedEvent (Data: PaymentAttemptEventData)
	EventNamePaymentAttemptCaptureFailed EventName = "payment_attempt.capture_failed"
)

// --- Online Payments: PaymentConsent ---
const (
	// EventNamePaymentConsentCreated occurs when a PaymentConsent has been created.
	// EventNamePaymentConsentCreated 对应结构体: PaymentConsentCreatedEvent (Data: PaymentConsentEventData)
	EventNamePaymentConsentCreated EventName = "payment_consent.created"
	// EventNamePaymentConsentUpdated occurs when a PaymentConsent has been updated.
	// EventNamePaymentConsentUpdated 对应结构体: PaymentConsentUpdatedEvent (Data: PaymentConsentEventData)
	EventNamePaymentConsentUpdated EventName = "payment_consent.updated"
	// EventNamePaymentConsentPending occurs when the mandate has been submitted and is waiting for final result.
	// EventNamePaymentConsentPending 对应结构体: PaymentConsentPendingEvent (Data: PaymentConsentEventData)
	EventNamePaymentConsentPending EventName = "payment_consent.pending"
	// EventNamePaymentConsentVerified occurs when a PaymentConsent has been verified.
	// EventNamePaymentConsentVerified 对应结构体: PaymentConsentVerifiedEvent (Data: PaymentConsentEventData)
	EventNamePaymentConsentVerified EventName = "payment_consent.verified"
	// EventNamePaymentConsentDisabled occurs when a PaymentConsent has been disabled.
	// EventNamePaymentConsentDisabled 对应结构体: PaymentConsentDisabledEvent (Data: PaymentConsentEventData)
	EventNamePaymentConsentDisabled EventName = "payment_consent.disabled"
	// EventNamePaymentConsentPaused occurs when a PaymentConsent has been paused.
	// EventNamePaymentConsentPaused 对应结构体: PaymentConsentPausedEvent (Data: PaymentConsentEventData)
	EventNamePaymentConsentPaused EventName = "payment_consent.paused"
	// EventNamePaymentConsentRequiresPaymentMethod occurs when a payment method is required.
	// EventNamePaymentConsentRequiresPaymentMethod 对应结构体: PaymentConsentRequiresPaymentMethodEvent (Data: PaymentConsentEventData)
	EventNamePaymentConsentRequiresPaymentMethod EventName = "payment_consent.requires_payment_method"
	// EventNamePaymentConsentRequiresCustomerAction occurs when additional customer action is required.
	// EventNamePaymentConsentRequiresCustomerAction 对应结构体: PaymentConsentRequiresCustomerActionEvent (Data: PaymentConsentEventData)
	EventNamePaymentConsentRequiresCustomerAction EventName = "payment_consent.requires_customer_action"
	// EventNamePaymentConsentVerificationFailed occurs when verification has failed.
	// EventNamePaymentConsentVerificationFailed 对应结构体: PaymentConsentVerificationFailedEvent (Data: PaymentConsentEventData)
	EventNamePaymentConsentVerificationFailed EventName = "payment_consent.verification_failed"
)

// --- Online Payments: Customer ---
const (
	// EventNameCustomerCreated occurs when a customer has been created.
	// EventNameCustomerCreated 对应结构体: CustomerCreatedEvent (Data: CustomerEventData)
	EventNameCustomerCreated EventName = "customer.created"
	// EventNameCustomerUpdated occurs when a customer has been updated.
	// EventNameCustomerUpdated 对应结构体: CustomerUpdatedEvent (Data: CustomerEventData)
	EventNameCustomerUpdated EventName = "customer.updated"
)

// --- Online Payments: Refund ---
const (
	// EventNameRefundReceived occurs when a refund request has been received.
	// EventNameRefundReceived 对应结构体: RefundReceivedEvent (Data: RefundEventData)
	EventNameRefundReceived EventName = "refund.received"
	// EventNameRefundAccepted occurs when a refund request has been accepted.
	// EventNameRefundAccepted 对应结构体: RefundAcceptedEvent (Data: RefundEventData)
	EventNameRefundAccepted EventName = "refund.accepted"
	// EventNameRefundSettled occurs when a refund has been settled.
	// EventNameRefundSettled 对应结构体: RefundSettledEvent (Data: RefundEventData)
	EventNameRefundSettled EventName = "refund.settled"
	// EventNameRefundFailed occurs when a refund has failed.
	// EventNameRefundFailed 对应结构体: RefundFailedEvent (Data: RefundEventData)
	EventNameRefundFailed EventName = "refund.failed"
)

// --- Online Payments: PaymentMethod ---
const (
	// EventNamePaymentMethodCreated occurs when a payment method has been created.
	// EventNamePaymentMethodCreated 对应结构体: PaymentMethodCreatedEvent (Data: PaymentMethodEventData)
	EventNamePaymentMethodCreated EventName = "payment_method.created"
	// EventNamePaymentMethodUpdated occurs when a payment method has been updated.
	// EventNamePaymentMethodUpdated 对应结构体: PaymentMethodUpdatedEvent (Data: PaymentMethodEventData)
	EventNamePaymentMethodUpdated EventName = "payment_method.updated"
	// EventNamePaymentMethodAttached occurs when a payment method has been attached to a customer.
	// EventNamePaymentMethodAttached 对应结构体: PaymentMethodAttachedEvent (Data: PaymentMethodEventData)
	EventNamePaymentMethodAttached EventName = "payment_method.attached"
	// EventNamePaymentMethodDetached occurs when a payment method has been detached from a customer.
	// EventNamePaymentMethodDetached 对应结构体: PaymentMethodDetachedEvent (Data: PaymentMethodEventData)
	EventNamePaymentMethodDetached EventName = "payment_method.detached"
	// EventNamePaymentMethodDisabled occurs when a payment method has been disabled.
	// EventNamePaymentMethodDisabled 对应结构体: PaymentMethodDisabledEvent (Data: PaymentMethodEventData)
	EventNamePaymentMethodDisabled EventName = "payment_method.disabled"
)

// --- Online Payments: PaymentDispute ---
const (
	// EventNamePaymentDisputeRequiresResponse occurs when a dispute requires your response.
	// EventNamePaymentDisputeRequiresResponse 对应结构体: PaymentDisputeRequiresResponseEvent (Data: PaymentDisputeEventData)
	EventNamePaymentDisputeRequiresResponse EventName = "payment_dispute.requires_response"
	// EventNamePaymentDisputeChallenged occurs when you have challenged the dispute.
	// EventNamePaymentDisputeChallenged 对应结构体: PaymentDisputeChallengedEvent (Data: PaymentDisputeEventData)
	EventNamePaymentDisputeChallenged EventName = "payment_dispute.challenged"
	// EventNamePaymentDisputeAccepted occurs when you have accepted the dispute.
	// EventNamePaymentDisputeAccepted 对应结构体: PaymentDisputeAcceptedEvent (Data: PaymentDisputeEventData)
	EventNamePaymentDisputeAccepted EventName = "payment_dispute.accepted"
	// EventNamePaymentDisputeExpired occurs when the dispute expired without response.
	// EventNamePaymentDisputeExpired 对应结构体: PaymentDisputeExpiredEvent (Data: PaymentDisputeEventData)
	EventNamePaymentDisputeExpired EventName = "payment_dispute.expired"
	// EventNamePaymentDisputePendingClosure occurs when a pre-arbitration request will be auto-accepted.
	// EventNamePaymentDisputePendingClosure 对应结构体: PaymentDisputePendingClosureEvent (Data: PaymentDisputeEventData)
	EventNamePaymentDisputePendingClosure EventName = "payment_dispute.pending_closure"
	// EventNamePaymentDisputePendingDecision occurs when the dispute response was escalated to arbitration.
	// EventNamePaymentDisputePendingDecision 对应结构体: PaymentDisputePendingDecisionEvent (Data: PaymentDisputeEventData)
	EventNamePaymentDisputePendingDecision EventName = "payment_dispute.pending_decision"
	// EventNamePaymentDisputeWon occurs when the issuing bank accepted your response.
	// EventNamePaymentDisputeWon 对应结构体: PaymentDisputeWonEvent (Data: PaymentDisputeEventData)
	EventNamePaymentDisputeWon EventName = "payment_dispute.won"
	// EventNamePaymentDisputeLost occurs when the issuing bank did not accept your response.
	// EventNamePaymentDisputeLost 对应结构体: PaymentDisputeLostEvent (Data: PaymentDisputeEventData)
	EventNamePaymentDisputeLost EventName = "payment_dispute.lost"
	// EventNamePaymentDisputeReversed occurs when the dispute has been reversed.
	// EventNamePaymentDisputeReversed 对应结构体: PaymentDisputeReversedEvent (Data: PaymentDisputeEventData)
	EventNamePaymentDisputeReversed EventName = "payment_dispute.reversed"
)

// --- Online Payments: PaymentLink ---
const (
	// EventNamePaymentLinkCreated occurs when a PaymentLink has been created.
	// EventNamePaymentLinkCreated 对应结构体: PaymentLinkCreatedEvent (Data: PaymentLinkEventData)
	EventNamePaymentLinkCreated EventName = "payment_link.created"
	// EventNamePaymentLinkPaid occurs when a PaymentLink received a new payment.
	// EventNamePaymentLinkPaid 对应结构体: PaymentLinkPaidEvent (Data: PaymentLinkEventData)
	EventNamePaymentLinkPaid EventName = "payment_link.paid"
)

// --- Online Payments: FundsSplit ---
const (
	// EventNameFundsSplitCreated occurs when a request to split funds has been created.
	// EventNameFundsSplitCreated 对应结构体: FundsSplitCreatedEvent (Data: FundsSplitEventData)
	EventNameFundsSplitCreated EventName = "funds_split.created"
	// EventNameFundsSplitFailed occurs when a request to split funds has failed.
	// EventNameFundsSplitFailed 对应结构体: FundsSplitFailedEvent (Data: FundsSplitEventData)
	EventNameFundsSplitFailed EventName = "funds_split.failed"
	// EventNameFundsSplitReleased occurs when the release instruction has been received.
	// EventNameFundsSplitReleased 对应结构体: FundsSplitReleasedEvent (Data: FundsSplitEventData)
	EventNameFundsSplitReleased EventName = "funds_split.released"
	// EventNameFundsSplitSettled occurs when funds have been split and settled.
	// EventNameFundsSplitSettled 对应结构体: FundsSplitSettledEvent (Data: FundsSplitEventData)
	EventNameFundsSplitSettled EventName = "funds_split.settled"
)

// --- Online Payments: Fraud ---
const (
	// EventNameFraudMerchantNotified occurs when a payment is identified as fraudulent.
	// EventNameFraudMerchantNotified 对应结构体: FraudMerchantNotifiedEvent (Data: FraudEventData)
	EventNameFraudMerchantNotified EventName = "fraud.merchant_notified"
)

// --- Issuing: Card Status ---
const (
	// EventNameIssuingCardPending occurs when card creation is pending review.
	// EventNameIssuingCardPending 对应结构体: IssuingCardEvent (Data: IssuingCardEventData)
	EventNameIssuingCardPending EventName = "issuing.card.pending"
	// EventNameIssuingCardFailed occurs when card creation failed.
	// EventNameIssuingCardFailed 对应结构体: IssuingCardEvent (Data: IssuingCardEventData)
	EventNameIssuingCardFailed EventName = "issuing.card.failed"
	// EventNameIssuingCardActive occurs when the card is active.
	// EventNameIssuingCardActive 对应结构体: IssuingCardEvent (Data: IssuingCardEventData)
	EventNameIssuingCardActive EventName = "issuing.card.active"
	// EventNameIssuingCardInactive occurs when the card is frozen.
	// EventNameIssuingCardInactive 对应结构体: IssuingCardEvent (Data: IssuingCardEventData)
	EventNameIssuingCardInactive EventName = "issuing.card.inactive"
	// EventNameIssuingCardBlocked occurs when the card is suspended.
	// EventNameIssuingCardBlocked 对应结构体: IssuingCardEvent (Data: IssuingCardEventData)
	EventNameIssuingCardBlocked EventName = "issuing.card.blocked"
	// EventNameIssuingCardLost occurs when the card is marked as lost.
	// EventNameIssuingCardLost 对应结构体: IssuingCardEvent (Data: IssuingCardEventData)
	EventNameIssuingCardLost EventName = "issuing.card.lost"
	// EventNameIssuingCardStolen occurs when the card is reported as stolen.
	// EventNameIssuingCardStolen 对应结构体: IssuingCardEvent (Data: IssuingCardEventData)
	EventNameIssuingCardStolen EventName = "issuing.card.stolen"
	// EventNameIssuingCardClosed occurs when the card is cancelled.
	// EventNameIssuingCardClosed 对应结构体: IssuingCardEvent (Data: IssuingCardEventData)
	EventNameIssuingCardClosed EventName = "issuing.card.closed"
	// EventNameIssuingCardExpired occurs when the card has expired.
	// EventNameIssuingCardExpired 对应结构体: IssuingCardEvent (Data: IssuingCardEventData)
	EventNameIssuingCardExpired EventName = "issuing.card.expired"
)

// --- Issuing: Cardholder Status ---
const (
	// EventNameIssuingCardholderIncomplete occurs when cardholder needs additional details.
	// EventNameIssuingCardholderIncomplete 对应结构体: IssuingCardholderEvent (Data: IssuingCardholderEventData)
	EventNameIssuingCardholderIncomplete EventName = "issuing.cardholder.incomplete"
	// EventNameIssuingCardholderPending occurs when cardholder needs to pass review.
	// EventNameIssuingCardholderPending 对应结构体: IssuingCardholderEvent (Data: IssuingCardholderEventData)
	EventNameIssuingCardholderPending EventName = "issuing.cardholder.pending"
	// EventNameIssuingCardholderReady occurs when cardholder has passed review.
	// EventNameIssuingCardholderReady 对应结构体: IssuingCardholderEvent (Data: IssuingCardholderEventData)
	EventNameIssuingCardholderReady EventName = "issuing.cardholder.ready"
	// EventNameIssuingCardholderDisabled occurs when cardholder has been disabled.
	// EventNameIssuingCardholderDisabled 对应结构体: IssuingCardholderEvent (Data: IssuingCardholderEventData)
	EventNameIssuingCardholderDisabled EventName = "issuing.cardholder.disabled"
	// EventNameIssuingCardholderDeleted occurs when cardholder has been deleted.
	// EventNameIssuingCardholderDeleted 对应结构体: IssuingCardholderEvent (Data: IssuingCardholderEventData)
	EventNameIssuingCardholderDeleted EventName = "issuing.cardholder.deleted"
)

// --- Issuing: Transaction ---
const (
	// EventNameIssuingTransactionSucceeded occurs when a transaction succeeded.
	// EventNameIssuingTransactionSucceeded 对应结构体: IssuingTransactionEvent (Data: IssuingTransactionEventData)
	EventNameIssuingTransactionSucceeded EventName = "issuing.transaction.succeeded"
	// EventNameIssuingTransactionFailed occurs when a transaction failed.
	// EventNameIssuingTransactionFailed 对应结构体: IssuingTransactionEvent (Data: IssuingTransactionEventData)
	EventNameIssuingTransactionFailed EventName = "issuing.transaction.failed"
)

// --- Issuing: Transaction Lifecycle ---
const (
	// EventNameIssuingCardTransactionLifecycleCreated occurs when a new card transaction lifecycle is created.
	// EventNameIssuingCardTransactionLifecycleCreated 对应结构体: IssuingCardTransactionLifecycleEvent (Data: IssuingCardTransactionLifecycleEventData)
	EventNameIssuingCardTransactionLifecycleCreated EventName = "issuing.card_transaction_lifecycle.created"
	// EventNameIssuingCardTransactionLifecycleModified occurs when a card transaction lifecycle is modified.
	// EventNameIssuingCardTransactionLifecycleModified 对应结构体: IssuingCardTransactionLifecycleEvent (Data: IssuingCardTransactionLifecycleEventData)
	EventNameIssuingCardTransactionLifecycleModified EventName = "issuing.card_transaction_lifecycle.modified"
)

// --- Issuing: Card Transaction Status ---
const (
	// EventNameIssuingCardTransactionAuthorized occurs when a card transaction is authorized.
	// EventNameIssuingCardTransactionAuthorized 对应结构体: IssuingCardTransactionStatusEvent (Data: IssuingCardTransactionStatusEventData)
	EventNameIssuingCardTransactionAuthorized EventName = "issuing.card_transaction.authorized"
	// EventNameIssuingCardTransactionVerified occurs when a card transaction is verified.
	// EventNameIssuingCardTransactionVerified 对应结构体: IssuingCardTransactionStatusEvent (Data: IssuingCardTransactionStatusEventData)
	EventNameIssuingCardTransactionVerified EventName = "issuing.card_transaction.verified"
	// EventNameIssuingCardTransactionCleared occurs when a card transaction is cleared.
	// EventNameIssuingCardTransactionCleared 对应结构体: IssuingCardTransactionStatusEvent (Data: IssuingCardTransactionStatusEventData)
	EventNameIssuingCardTransactionCleared EventName = "issuing.card_transaction.cleared"
	// EventNameIssuingCardTransactionReversed occurs when a card transaction is reversed.
	// EventNameIssuingCardTransactionReversed 对应结构体: IssuingCardTransactionStatusEvent (Data: IssuingCardTransactionStatusEventData)
	EventNameIssuingCardTransactionReversed EventName = "issuing.card_transaction.reversed"
	// EventNameIssuingCardTransactionExpired occurs when a card transaction has expired.
	// EventNameIssuingCardTransactionExpired 对应结构体: IssuingCardTransactionStatusEvent (Data: IssuingCardTransactionStatusEventData)
	EventNameIssuingCardTransactionExpired EventName = "issuing.card_transaction.expired"
	// EventNameIssuingCardTransactionDeclined occurs when a card transaction is declined.
	// EventNameIssuingCardTransactionDeclined 对应结构体: IssuingCardTransactionStatusEvent (Data: IssuingCardTransactionStatusEventData)
	EventNameIssuingCardTransactionDeclined EventName = "issuing.card_transaction.declined"
	// EventNameIssuingCardTransactionModified occurs when a card transaction was updated without status change.
	// EventNameIssuingCardTransactionModified 对应结构体: IssuingCardTransactionStatusEvent (Data: IssuingCardTransactionStatusEventData)
	EventNameIssuingCardTransactionModified EventName = "issuing.card_transaction.modified"
)

// --- Issuing: Card Transaction Events ---
const (
	// EventNameIssuingCardTransactionEventSuccess occurs when a transaction event was processed successfully.
	// EventNameIssuingCardTransactionEventSuccess 对应结构体: IssuingCardTransactionEvent (Data: IssuingCardTransactionEventData)
	EventNameIssuingCardTransactionEventSuccess EventName = "issuing.card_transaction_event.success"
	// EventNameIssuingCardTransactionEventFailed occurs when a transaction event was declined or failed.
	// EventNameIssuingCardTransactionEventFailed 对应结构体: IssuingCardTransactionEvent (Data: IssuingCardTransactionEventData)
	EventNameIssuingCardTransactionEventFailed EventName = "issuing.card_transaction_event.failed"
	// EventNameIssuingCardTransactionEventEnriched occurs when enriched data was added to a transaction event.
	// EventNameIssuingCardTransactionEventEnriched 对应结构体: IssuingCardTransactionEvent (Data: IssuingCardTransactionEventData)
	EventNameIssuingCardTransactionEventEnriched EventName = "issuing.card_transaction_event.enriched"
)

// --- Issuing: Disputes ---
const (
	// EventNameIssuingTransactionDisputeCreated occurs when a dispute draft is created.
	// EventNameIssuingTransactionDisputeCreated 对应结构体: IssuingTransactionDisputeEvent (Data: IssuingTransactionDisputeEventData)
	EventNameIssuingTransactionDisputeCreated EventName = "issuing.transaction_dispute.created"
	// EventNameIssuingTransactionDisputeExpired occurs when the dispute validity period is over.
	// EventNameIssuingTransactionDisputeExpired 对应结构体: IssuingTransactionDisputeEvent (Data: IssuingTransactionDisputeEventData)
	EventNameIssuingTransactionDisputeExpired EventName = "issuing.transaction_dispute.expired"
	// EventNameIssuingTransactionDisputeSubmitted occurs when the dispute is submitted to Airwallex.
	// EventNameIssuingTransactionDisputeSubmitted 对应结构体: IssuingTransactionDisputeEvent (Data: IssuingTransactionDisputeEventData)
	EventNameIssuingTransactionDisputeSubmitted EventName = "issuing.transaction_dispute.submitted"
	// EventNameIssuingTransactionDisputeRejected occurs when the dispute is rejected.
	// EventNameIssuingTransactionDisputeRejected 对应结构体: IssuingTransactionDisputeEvent (Data: IssuingTransactionDisputeEventData)
	EventNameIssuingTransactionDisputeRejected EventName = "issuing.transaction_dispute.rejected"
	// EventNameIssuingTransactionDisputeCanceled occurs when the dispute is canceled.
	// EventNameIssuingTransactionDisputeCanceled 对应结构体: IssuingTransactionDisputeEvent (Data: IssuingTransactionDisputeEventData)
	EventNameIssuingTransactionDisputeCanceled EventName = "issuing.transaction_dispute.canceled"
	// EventNameIssuingTransactionDisputeAccepted occurs when the dispute is submitted to the card scheme.
	// EventNameIssuingTransactionDisputeAccepted 对应结构体: IssuingTransactionDisputeEvent (Data: IssuingTransactionDisputeEventData)
	EventNameIssuingTransactionDisputeAccepted EventName = "issuing.transaction_dispute.accepted"
	// EventNameIssuingTransactionDisputeWon occurs when the dispute is won.
	// EventNameIssuingTransactionDisputeWon 对应结构体: IssuingTransactionDisputeEvent (Data: IssuingTransactionDisputeEventData)
	EventNameIssuingTransactionDisputeWon EventName = "issuing.transaction_dispute.won"
	// EventNameIssuingTransactionDisputeLost occurs when the dispute is lost.
	// EventNameIssuingTransactionDisputeLost 对应结构体: IssuingTransactionDisputeEvent (Data: IssuingTransactionDisputeEventData)
	EventNameIssuingTransactionDisputeLost EventName = "issuing.transaction_dispute.lost"
)

// --- Issuing: Reissue ---
const (
	// EventNameIssuingReissueSucceeded occurs when a card reissue succeeded.
	// EventNameIssuingReissueSucceeded 对应结构体: IssuingReissueEvent (Data: IssuingReissueEventData)
	EventNameIssuingReissueSucceeded EventName = "issuing.reissue.succeeded"
)

// --- Issuing: Card Alerts ---
const (
	// EventNameIssuingCardLowRemainingTransactionLimit occurs when the card's remaining limit falls below threshold.
	// EventNameIssuingCardLowRemainingTransactionLimit 对应结构体: IssuingCardAlertEvent (Data: IssuingCardAlertEventData)
	EventNameIssuingCardLowRemainingTransactionLimit EventName = "issuing.card.low_remaining_transaction_limit"
)

// --- Deposits ---
const (
	// EventNameDepositPending occurs when a deposit is pending.
	// EventNameDepositPending 对应结构体: DepositPendingEvent (Data: Deposit)
	EventNameDepositPending EventName = "deposit.pending"
	// EventNameDepositSettled occurs when a deposit is successfully settled to Wallet.
	// EventNameDepositSettled 对应结构体: DepositSettledEvent (Data: Deposit)
	EventNameDepositSettled EventName = "deposit.settled"
	// EventNameDepositRejected occurs when a deposit is rejected after being reviewed by Airwallex or the clearing system.
	// EventNameDepositRejected 对应结构体: DepositRejectedEvent (Data: Deposit)
	EventNameDepositRejected EventName = "deposit.rejected"
	// EventNameDepositReversed occurs when a deposit is recalled by the payer's bank after the funds were settled to your Wallet.
	// EventNameDepositReversed 对应结构体: DepositReversedEvent (Data: Deposit)
	EventNameDepositReversed EventName = "deposit.reversed"
)

// --- Account ---
const (
	// EventNameAccountActive occurs when an account is active.
	// EventNameAccountActive 对应结构体: AccountActiveEvent (Data: AccountEventData)
	EventNameAccountActive EventName = "account.active"
	// EventNameAccountConnected occurs when an account is connected to another account.
	// EventNameAccountConnected 对应结构体: AccountConnectedEvent (Data: AccountConnectedEventData)
	EventNameAccountConnected EventName = "account.connected"
	// EventNameAccountSuspended occurs when account verification is unsuccessful.
	// EventNameAccountSuspended 对应结构体: AccountSuspendedEvent (Data: AccountEventData)
	EventNameAccountSuspended EventName = "account.suspended"
	// EventNameAccountActionRequired occurs when an account requires further action to proceed with verification.
	// EventNameAccountActionRequired 对应结构体: AccountActionRequiredEvent (Data: AccountActionRequiredEventData)
	EventNameAccountActionRequired EventName = "account.action_required"
	// EventNameAccountSubmitted occurs when an account is submitted.
	// EventNameAccountSubmitted 对应结构体: AccountSubmittedEvent (Data: AccountSubmittedEventData)
	EventNameAccountSubmitted EventName = "account.submitted"
)

// --- Account Capability ---
const (
	// EventNameAccountCapabilityEnabled occurs when a capability has been enabled successfully.
	// EventNameAccountCapabilityEnabled 对应结构体: AccountCapabilityEnabledEvent (Data: AccountCapabilityEventData)
	EventNameAccountCapabilityEnabled EventName = "account_capability.enabled"
	// EventNameAccountCapabilityDisabled occurs when a capability has been disabled or the enablement request has been rejected.
	// EventNameAccountCapabilityDisabled 对应结构体: AccountCapabilityDisabledEvent (Data: AccountCapabilityEventData)
	EventNameAccountCapabilityDisabled EventName = "account_capability.disabled"
	// EventNameAccountCapabilityPending occurs when a capability is pending.
	// EventNameAccountCapabilityPending 对应结构体: AccountCapabilityPendingEvent (Data: AccountCapabilityEventData)
	EventNameAccountCapabilityPending EventName = "account_capability.pending"
)

// --- Linked Accounts ---
const (
	// EventNameLinkedAccountRequiresAction occurs when a Linked Account requires further action to proceed.
	// EventNameLinkedAccountRequiresAction 对应结构体: LinkedAccountRequiresActionEvent (Data: LinkedAccountEventData)
	EventNameLinkedAccountRequiresAction EventName = "linked_account.requires_action"
	// EventNameLinkedAccountProcessing occurs when a Linked Account is in the process of being set up.
	// EventNameLinkedAccountProcessing 对应结构体: LinkedAccountProcessingEvent (Data: LinkedAccountEventData)
	EventNameLinkedAccountProcessing EventName = "linked_account.processing"
	// EventNameLinkedAccountSucceeded occurs when a Linked Account verification is successful and is active.
	// EventNameLinkedAccountSucceeded 对应结构体: LinkedAccountSucceededEvent (Data: LinkedAccountEventData)
	EventNameLinkedAccountSucceeded EventName = "linked_account.succeeded"
	// EventNameLinkedAccountFailed occurs when a Linked Account failed verification or is inactive.
	// EventNameLinkedAccountFailed 对应结构体: LinkedAccountFailedEvent (Data: LinkedAccountEventData)
	EventNameLinkedAccountFailed EventName = "linked_account.failed"
	// EventNameLinkedAccountSuspended occurs when a Linked Account is suspended.
	// EventNameLinkedAccountSuspended 对应结构体: LinkedAccountSuspendedEvent (Data: LinkedAccountEventData)
	EventNameLinkedAccountSuspended EventName = "linked_account.suspended"
)

// --- Platform: PlatformReport ---
const (
	// EventNamePlatformReportCompleted occurs when creation of Platform report completed.
	// EventNamePlatformReportCompleted 对应结构体: PlatformReportCompletedEvent (Data: PlatformReport)
	EventNamePlatformReportCompleted EventName = "platform_report.completed"
	// EventNamePlatformReportFailed occurs when creation of Platform report failed.
	// EventNamePlatformReportFailed 对应结构体: PlatformReportFailedEvent (Data: PlatformReport)
	EventNamePlatformReportFailed EventName = "platform_report.failed"
)

// --- Billing: Subscription ---
const (
	// EventNameSubscriptionCreated occurs when a subscription has been created.
	// EventNameSubscriptionCreated 对应结构体: SubscriptionCreatedEvent (Data: billing.Subscription)
	EventNameSubscriptionCreated EventName = "subscription.created"
	// EventNameSubscriptionInTrial occurs when the subscription transits to IN_TRIAL status.
	// EventNameSubscriptionInTrial 对应结构体: SubscriptionInTrialEvent (Data: billing.Subscription)
	EventNameSubscriptionInTrial EventName = "subscription.in_trial"
	// EventNameSubscriptionActive occurs when the subscription transits to ACTIVE status.
	// EventNameSubscriptionActive 对应结构体: SubscriptionActiveEvent (Data: billing.Subscription)
	EventNameSubscriptionActive EventName = "subscription.active"
	// EventNameSubscriptionUnpaid occurs when the subscription transits to UNPAID status.
	// EventNameSubscriptionUnpaid 对应结构体: SubscriptionUnpaidEvent (Data: billing.Subscription)
	EventNameSubscriptionUnpaid EventName = "subscription.unpaid"
	// EventNameSubscriptionModified occurs when the subscription is updated.
	// EventNameSubscriptionModified 对应结构体: SubscriptionModifiedEvent (Data: billing.Subscription)
	EventNameSubscriptionModified EventName = "subscription.modified"
	// EventNameSubscriptionCancelled occurs when the subscription has been cancelled.
	// EventNameSubscriptionCancelled 对应结构体: SubscriptionCancelledEvent (Data: billing.Subscription)
	EventNameSubscriptionCancelled EventName = "subscription.cancelled"
	// EventNameSubscriptionUpdated occurs when the subscription has been changed (legacy, API version 2025-04-25 and before).
	// EventNameSubscriptionUpdated 对应结构体: SubscriptionUpdatedEvent (Data: BillingSubscriptionEventData)
	EventNameSubscriptionUpdated EventName = "subscription.updated"
)

// --- Billing: Invoice ---
const (
	// EventNameInvoiceCreated occurs when the invoice has been created.
	// EventNameInvoiceCreated 对应结构体: InvoiceCreatedEvent (Data: billing.Invoice)
	EventNameInvoiceCreated EventName = "invoice.created"
	// EventNameInvoiceFinalized occurs when the invoice has been finalized.
	// EventNameInvoiceFinalized 对应结构体: InvoiceFinalizedEvent (Data: billing.Invoice)
	EventNameInvoiceFinalized EventName = "invoice.finalized"
	// EventNameInvoiceVoided occurs when the invoice has been voided.
	// EventNameInvoiceVoided 对应结构体: InvoiceVoidedEvent (Data: billing.Invoice)
	EventNameInvoiceVoided EventName = "invoice.voided"
	// EventNameInvoiceUpdated occurs when the invoice has been updated.
	// EventNameInvoiceUpdated 对应结构体: InvoiceUpdatedEvent (Data: billing.Invoice)
	EventNameInvoiceUpdated EventName = "invoice.updated"
	// EventNameInvoicePaymentPaid occurs when the invoice has been paid.
	// EventNameInvoicePaymentPaid 对应结构体: InvoicePaymentPaidEvent (Data: billing.Invoice)
	EventNameInvoicePaymentPaid EventName = "invoice.payment.paid"
	// EventNameInvoiceSent occurs when the invoice is ready to be paid (legacy, API version 2025-04-25 and before).
	// EventNameInvoiceSent 对应结构体: InvoiceSentEvent (Data: BillingInvoiceEventData)
	EventNameInvoiceSent EventName = "invoice.sent"
	// EventNameInvoicePaid occurs when the invoice has been paid successfully (legacy, API version 2025-04-25 and before).
	// EventNameInvoicePaid 对应结构体: InvoicePaidEvent (Data: BillingInvoiceEventData)
	EventNameInvoicePaid EventName = "invoice.paid"
	// EventNameInvoicePaymentFailed occurs when the invoice is marked as PAYMENT_FAILED (legacy, API version 2025-04-25 and before).
	// EventNameInvoicePaymentFailed 对应结构体: InvoicePaymentFailedEvent (Data: BillingInvoiceEventData)
	EventNameInvoicePaymentFailed EventName = "invoice.payment_failed"
	// EventNameInvoicePaymentAttemptFailed occurs when a payment attempt on the invoice failed (legacy, API version 2025-04-25 and before).
	// EventNameInvoicePaymentAttemptFailed 对应结构体: InvoicePaymentAttemptFailedEvent (Data: BillingInvoiceEventData)
	EventNameInvoicePaymentAttemptFailed EventName = "invoice.payment_attempt_failed"
)

// --- Billing: Billing Transaction ---
const (
	// EventNameBillingTransactionCreated occurs when the Billing Transaction has been created.
	// EventNameBillingTransactionCreated 对应结构体: BillingTransactionCreatedEvent (Data: billing.BillingTransaction)
	EventNameBillingTransactionCreated EventName = "billing_transaction.created"
	// EventNameBillingTransactionSucceeded occurs when the Billing Transaction has succeeded.
	// EventNameBillingTransactionSucceeded 对应结构体: BillingTransactionSucceededEvent (Data: billing.BillingTransaction)
	EventNameBillingTransactionSucceeded EventName = "billing_transaction.succeeded"
	// EventNameBillingTransactionCancelled occurs when the Billing Transaction has been cancelled.
	// EventNameBillingTransactionCancelled 对应结构体: BillingTransactionCancelledEvent (Data: billing.BillingTransaction)
	EventNameBillingTransactionCancelled EventName = "billing_transaction.cancelled"
)

// --- Billing: Billing Checkout ---
const (
	// EventNameBillingCheckoutCreated occurs when the Billing Checkout has been created.
	// EventNameBillingCheckoutCreated 对应结构体: BillingCheckoutCreatedEvent (Data: billing.Checkout)
	EventNameBillingCheckoutCreated EventName = "billing_checkout.created"
	// EventNameBillingCheckoutCancelled occurs when the Billing Checkout has been cancelled.
	// EventNameBillingCheckoutCancelled 对应结构体: BillingCheckoutCancelledEvent (Data: billing.Checkout)
	EventNameBillingCheckoutCancelled EventName = "billing_checkout.cancelled"
	// EventNameBillingCheckoutCompleted occurs when the Billing Checkout has completed.
	// EventNameBillingCheckoutCompleted 对应结构体: BillingCheckoutCompletedEvent (Data: billing.Checkout)
	EventNameBillingCheckoutCompleted EventName = "billing_checkout.completed"
)

// --- Billing: Usage Event ---
const (
	// EventNameUsageEventAggregationFailed occurs when failed Billing Usage Events have been aggregated.
	// EventNameUsageEventAggregationFailed 对应结构体: UsageEventAggregationFailedEvent (Data: UsageEventAggregationFailedData)
	EventNameUsageEventAggregationFailed EventName = "usage_event.aggregation_failed"
)

// --- Billing: Credit Note ---
const (
	// EventNameCreditNoteCreated occurs when the credit note has been created.
	// EventNameCreditNoteCreated 对应结构体: CreditNoteCreatedEvent (Data: billing.CreditNote)
	EventNameCreditNoteCreated EventName = "credit_note.created"
	// EventNameCreditNoteFinalized occurs when the credit note has been finalized.
	// EventNameCreditNoteFinalized 对应结构体: CreditNoteFinalizedEvent (Data: billing.CreditNote)
	EventNameCreditNoteFinalized EventName = "credit_note.finalized"
	// EventNameCreditNoteVoided occurs when the credit note has been voided.
	// EventNameCreditNoteVoided 对应结构体: CreditNoteVoidedEvent (Data: billing.CreditNote)
	EventNameCreditNoteVoided EventName = "credit_note.voided"
)

// --- Global Accounts ---
const (
	// EventNameGlobalAccountActive occurs when a Global Account is active and can be used.
	// EventNameGlobalAccountActive 对应结构体: GlobalAccountActiveEvent (Data: GlobalAccountEventData)
	EventNameGlobalAccountActive EventName = "global_account.active"
	// EventNameGlobalAccountClosed occurs when a Global Account has been closed.
	// EventNameGlobalAccountClosed 对应结构体: GlobalAccountClosedEvent (Data: GlobalAccountEventData)
	EventNameGlobalAccountClosed EventName = "global_account.closed"
	// EventNameGlobalAccountFailed occurs when a Global Account creation has failed.
	// EventNameGlobalAccountFailed 对应结构体: GlobalAccountFailedEvent (Data: GlobalAccountEventData)
	EventNameGlobalAccountFailed EventName = "global_account.failed"
	// EventNameGADepositNew occurs when a deposit is received via Global Account (legacy, API version 2025-02-14 and before).
	// EventNameGADepositNew 对应结构体: GADepositNewEvent (Data: GADepositNewEventData)
	EventNameGADepositNew EventName = "ga.new"
)

// --- Balance ---
const (
	// EventNameBalanceVATopUp occurs when wallet balance increases due to deposit received via Virtual Account.
	// EventNameBalanceVATopUp 对应结构体: BalanceVATopUpEvent (Data: BalanceVATopUpData)
	EventNameBalanceVATopUp EventName = "balance.va.top_up"
	// EventNameBalanceGATopUp occurs when wallet balance increases due to deposit received via Global Account.
	// EventNameBalanceGATopUp 对应结构体: BalanceGATopUpEvent (Data: BalanceGATopUpData)
	EventNameBalanceGATopUp EventName = "balance.ga.top_up"
	// EventNameBalanceAdjustment occurs when balance changes due to adjustment.
	// EventNameBalanceAdjustment 对应结构体: BalanceAdjustmentEvent (Data: BalanceAdjustmentData)
	EventNameBalanceAdjustment EventName = "balance.adjustment"
)

// --- Charges ---
const (
	// EventNameChargeNew occurs when a charge transaction is created successfully.
	// EventNameChargeNew 对应结构体: ChargeNewEvent (Data: ChargeEventData)
	EventNameChargeNew EventName = "charge.new"
	// EventNameChargePending occurs when a charge transaction is in progress.
	// EventNameChargePending 对应结构体: ChargePendingEvent (Data: ChargeEventData)
	EventNameChargePending EventName = "charge.pending"
	// EventNameChargeSettled occurs when funds have settled in the target wallet.
	// EventNameChargeSettled 对应结构体: ChargeSettledEvent (Data: ChargeEventData)
	EventNameChargeSettled EventName = "charge.settled"
	// EventNameChargeSuspended occurs when a charge transaction is suspended.
	// EventNameChargeSuspended 对应结构体: ChargeSuspendedEvent (Data: ChargeEventData)
	EventNameChargeSuspended EventName = "charge.suspended"
	// EventNameChargeFailed occurs when a charge transaction has failed.
	// EventNameChargeFailed 对应结构体: ChargeFailedEvent (Data: ChargeEventData)
	EventNameChargeFailed EventName = "charge.failed"
)

// --- Direct Debit Payouts ---
const (
	// EventNameDirectDebitCreated occurs when a direct debit payout instruction is received.
	// EventNameDirectDebitCreated 对应结构体: DirectDebitCreatedEvent (Data: DirectDebit)
	EventNameDirectDebitCreated EventName = "direct_debit.created"
	// EventNameDirectDebitInReview occurs when a direct debit payout is on hold for review.
	// EventNameDirectDebitInReview 对应结构体: DirectDebitInReviewEvent (Data: DirectDebit)
	EventNameDirectDebitInReview EventName = "direct_debit.in_review"
	// EventNameDirectDebitPending occurs when a direct debit payout is pending for settlement.
	// EventNameDirectDebitPending 对应结构体: DirectDebitPendingEvent (Data: DirectDebit)
	EventNameDirectDebitPending EventName = "direct_debit.pending"
	// EventNameDirectDebitRejected occurs when a direct debit payout was rejected after being reviewed.
	// EventNameDirectDebitRejected 对应结构体: DirectDebitRejectedEvent (Data: DirectDebit)
	EventNameDirectDebitRejected EventName = "direct_debit.rejected"
	// EventNameDirectDebitSettled occurs when a direct debit payout has been debited from your Wallet balance.
	// EventNameDirectDebitSettled 对应结构体: DirectDebitSettledEvent (Data: DirectDebit)
	EventNameDirectDebitSettled EventName = "direct_debit.settled"
	// EventNameDirectDebitReturned occurs when the request to dispute the payout was successful. The funds will be credited back to your Wallet balance.
	// EventNameDirectDebitReturned 对应结构体: DirectDebitReturnedEvent (Data: DirectDebit)
	EventNameDirectDebitReturned EventName = "direct_debit.returned"
)

// --- Transfers ---
const (
	// EventNameTransferInApproval occurs when the transfer is submitted and waiting for approval.
	// EventNameTransferInApproval 对应结构体: TransferInApprovalEvent (Data: TransferEventData)
	EventNameTransferInApproval EventName = "payout.transfer.in_approval"
	// EventNameTransferApprovalRecalled occurs when transfer is recalled by a user through the web app.
	// EventNameTransferApprovalRecalled 对应结构体: TransferApprovalRecalledEvent (Data: TransferEventData)
	EventNameTransferApprovalRecalled EventName = "payout.transfer.approval_recalled"
	// EventNameTransferApprovalRejected occurs when transfer is rejected by an approver through the web app.
	// EventNameTransferApprovalRejected 对应结构体: TransferApprovalRejectedEvent (Data: TransferEventData)
	EventNameTransferApprovalRejected EventName = "payout.transfer.approval_rejected"
	// EventNameTransferApprovalBlocked occurs when transfer is blocked as the next approver cannot be found.
	// EventNameTransferApprovalBlocked 对应结构体: TransferApprovalBlockedEvent (Data: TransferEventData)
	EventNameTransferApprovalBlocked EventName = "payout.transfer.approval_blocked"
	// EventNameTransferScheduled occurs when the transfer is scheduled to be processed on the transfer date.
	// EventNameTransferScheduled 对应结构体: TransferScheduledEvent (Data: TransferEventData)
	EventNameTransferScheduled EventName = "payout.transfer.scheduled"
	// EventNameTransferOverdue occurs when the transfer has not been funded by the scheduled date.
	// EventNameTransferOverdue 对应结构体: TransferOverdueEvent (Data: TransferEventData)
	EventNameTransferOverdue EventName = "payout.transfer.overdue"
	// EventNameTransferProcessing occurs when the transfer is funded and being processed.
	// EventNameTransferProcessing 对应结构体: TransferProcessingEvent (Data: TransferEventData)
	EventNameTransferProcessing EventName = "payout.transfer.processing"
	// EventNameTransferSent occurs when the transfer has been sent from Airwallex.
	// EventNameTransferSent 对应结构体: TransferSentEvent (Data: TransferEventData)
	EventNameTransferSent EventName = "payout.transfer.sent"
	// EventNameTransferPaid occurs when the transfer has been processed successfully by our banking partner.
	// EventNameTransferPaid 对应结构体: TransferPaidEvent (Data: TransferEventData)
	EventNameTransferPaid EventName = "payout.transfer.paid"
	// EventNameTransferFailed occurs when the transfer has failed our banking partner's processing or was rejected by the recipient bank.
	// EventNameTransferFailed 对应结构体: TransferFailedEvent (Data: TransferEventData)
	EventNameTransferFailed EventName = "payout.transfer.failed"
	// EventNameTransferCancellationRequested occurs when transfer cancellation has been requested before SENT.
	// EventNameTransferCancellationRequested 对应结构体: TransferCancellationRequestedEvent (Data: TransferEventData)
	EventNameTransferCancellationRequested EventName = "payout.transfer.cancellation_requested"
	// EventNameTransferCancelled occurs when the transfer has been successfully cancelled.
	// EventNameTransferCancelled 对应结构体: TransferCancelledEvent (Data: TransferEventData)
	EventNameTransferCancelled EventName = "payout.transfer.cancelled"
	// EventNameTransferFundingRequiresFundingConfirmation occurs when the funding of this transfer requires your confirmation.
	// EventNameTransferFundingRequiresFundingConfirmation 对应结构体: TransferFundingRequiresFundingConfirmationEvent (Data: TransferEventData)
	EventNameTransferFundingRequiresFundingConfirmation EventName = "payout.transfer.funding.requires_funding_confirmation"
	// EventNameTransferFundingScheduled occurs when the funding is scheduled to be processed on the transfer date.
	// EventNameTransferFundingScheduled 对应结构体: TransferFundingScheduledEvent (Data: TransferEventData)
	EventNameTransferFundingScheduled EventName = "payout.transfer.funding.scheduled"
	// EventNameTransferFundingProcessing occurs when the funding for this transfer is being processed.
	// EventNameTransferFundingProcessing 对应结构体: TransferFundingProcessingEvent (Data: TransferEventData)
	EventNameTransferFundingProcessing EventName = "payout.transfer.funding.processing"
	// EventNameTransferFundingFunded occurs when this transfer is successfully funded.
	// EventNameTransferFundingFunded 对应结构体: TransferFundingFundedEvent (Data: TransferEventData)
	EventNameTransferFundingFunded EventName = "payout.transfer.funding.funded"
	// EventNameTransferFundingFailed occurs when this transfer failed to be funded.
	// EventNameTransferFundingFailed 对应结构体: TransferFundingFailedEvent (Data: TransferEventData)
	EventNameTransferFundingFailed EventName = "payout.transfer.funding.failed"
	// EventNameTransferFundingCancelled occurs when the funding for this transfer has been cancelled.
	// EventNameTransferFundingCancelled 对应结构体: TransferFundingCancelledEvent (Data: TransferEventData)
	EventNameTransferFundingCancelled EventName = "payout.transfer.funding.cancelled"
	// EventNameTransferFundingReversed occurs when the funding status has been updated to REVERSED.
	// EventNameTransferFundingReversed 对应结构体: TransferFundingReversedEvent (Data: TransferEventData)
	EventNameTransferFundingReversed EventName = "payout.transfer.funding.reversed"
)

// --- Spend: Card Expense ---
const (
	// EventNameSpendExpenseDraft occurs when a card expense has been created.
	// EventNameSpendExpenseDraft 对应结构体: SpendExpenseDraftEvent (Data: SpendExpenseEventData)
	EventNameSpendExpenseDraft EventName = "spend.expense.draft"
	// EventNameSpendExpenseAwaitingApproval occurs when a card expense has been submitted for approval, or has been approved but requires additional approvals.
	// EventNameSpendExpenseAwaitingApproval 对应结构体: SpendExpenseAwaitingApprovalEvent (Data: SpendExpenseEventData)
	EventNameSpendExpenseAwaitingApproval EventName = "spend.expense.awaiting_approval"
	// EventNameSpendExpenseUpdated occurs when a card expense or its underlying card transaction has been updated.
	// EventNameSpendExpenseUpdated 对应结构体: SpendExpenseUpdatedEvent (Data: SpendExpenseEventData)
	EventNameSpendExpenseUpdated EventName = "spend.expense.updated"
	// EventNameSpendExpenseRejected occurs when a card expense has been rejected.
	// EventNameSpendExpenseRejected 对应结构体: SpendExpenseRejectedEvent (Data: SpendExpenseEventData)
	EventNameSpendExpenseRejected EventName = "spend.expense.rejected"
	// EventNameSpendExpenseApproved occurs when a card expense has been fully approved.
	// EventNameSpendExpenseApproved 对应结构体: SpendExpenseApprovedEvent (Data: SpendExpenseEventData)
	EventNameSpendExpenseApproved EventName = "spend.expense.approved"
	// EventNameSpendExpenseArchived occurs when a card expense has been archived.
	// EventNameSpendExpenseArchived 对应结构体: SpendExpenseArchivedEvent (Data: SpendExpenseEventData)
	EventNameSpendExpenseArchived EventName = "spend.expense.archived"
	// EventNameSpendExpenseDeleted occurs when a card expense has been deleted.
	// EventNameSpendExpenseDeleted 对应结构体: SpendExpenseDeletedEvent (Data: SpendExpenseEventData)
	EventNameSpendExpenseDeleted EventName = "spend.expense.deleted"
)

// --- Spend: Reimbursement Report ---
const (
	// EventNameSpendReimbursementReportDraft occurs when a reimbursement report has been created.
	// EventNameSpendReimbursementReportDraft 对应结构体: SpendReimbursementReportDraftEvent (Data: SpendReimbursementReportEventData)
	EventNameSpendReimbursementReportDraft EventName = "spend.reimbursement_report.draft"
	// EventNameSpendReimbursementReportAwaitingApproval occurs when a reimbursement report has been submitted for approval.
	// EventNameSpendReimbursementReportAwaitingApproval 对应结构体: SpendReimbursementReportAwaitingApprovalEvent (Data: SpendReimbursementReportEventData)
	EventNameSpendReimbursementReportAwaitingApproval EventName = "spend.reimbursement_report.awaiting_approval"
	// EventNameSpendReimbursementReportAwaitingPayment occurs when a reimbursement report has been approved and is awaiting payment.
	// EventNameSpendReimbursementReportAwaitingPayment 对应结构体: SpendReimbursementReportAwaitingPaymentEvent (Data: SpendReimbursementReportEventData)
	EventNameSpendReimbursementReportAwaitingPayment EventName = "spend.reimbursement_report.awaiting_payment"
	// EventNameSpendReimbursementReportRejected occurs when a reimbursement report has been rejected.
	// EventNameSpendReimbursementReportRejected 对应结构体: SpendReimbursementReportRejectedEvent (Data: SpendReimbursementReportEventData)
	EventNameSpendReimbursementReportRejected EventName = "spend.reimbursement_report.rejected"
	// EventNameSpendReimbursementReportPaymentInProgress occurs when payment for a reimbursement report is being processed.
	// EventNameSpendReimbursementReportPaymentInProgress 对应结构体: SpendReimbursementReportPaymentInProgressEvent (Data: SpendReimbursementReportEventData)
	EventNameSpendReimbursementReportPaymentInProgress EventName = "spend.reimbursement_report.payment_in_progress"
	// EventNameSpendReimbursementReportPaid occurs when a reimbursement report has been paid through Airwallex.
	// EventNameSpendReimbursementReportPaid 对应结构体: SpendReimbursementReportPaidEvent (Data: SpendReimbursementReportEventData)
	EventNameSpendReimbursementReportPaid EventName = "spend.reimbursement_report.paid"
	// EventNameSpendReimbursementReportMarkAsPaid occurs when a reimbursement report has been marked as paid externally, outside Airwallex.
	// EventNameSpendReimbursementReportMarkAsPaid 对应结构体: SpendReimbursementReportMarkAsPaidEvent (Data: SpendReimbursementReportEventData)
	EventNameSpendReimbursementReportMarkAsPaid EventName = "spend.reimbursement_report.mark_as_paid"
	// EventNameSpendReimbursementReportDeleted occurs when a reimbursement report has been deleted.
	// EventNameSpendReimbursementReportDeleted 对应结构体: SpendReimbursementReportDeletedEvent (Data: SpendReimbursementReportEventData)
	EventNameSpendReimbursementReportDeleted EventName = "spend.reimbursement_report.deleted"
	// EventNameSpendReimbursementReportUpdated occurs when a reimbursement report has been updated.
	// EventNameSpendReimbursementReportUpdated 对应结构体: SpendReimbursementReportUpdatedEvent (Data: SpendReimbursementReportEventData)
	EventNameSpendReimbursementReportUpdated EventName = "spend.reimbursement_report.updated"
)

// --- Batch Transfers ---
const (
	// EventNameBatchTransfersDrafting occurs when the batch transfer is being drafted.
	// EventNameBatchTransfersDrafting 对应结构体: BatchTransfersDraftingEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersDrafting EventName = "payout.batch_transfers.drafting"
	// EventNameBatchTransfersInApproval occurs when the batch transfer is submitted for approval.
	// EventNameBatchTransfersInApproval 对应结构体: BatchTransfersInApprovalEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersInApproval EventName = "payout.batch_transfers.in_approval"
	// EventNameBatchTransfersApprovalRecalled occurs when the batch transfer is recalled by a user in the web app.
	// EventNameBatchTransfersApprovalRecalled 对应结构体: BatchTransfersApprovalRecalledEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersApprovalRecalled EventName = "payout.batch_transfers.approval_recalled"
	// EventNameBatchTransfersApprovalRejected occurs when the batch transfer is rejected by an approver in the web app.
	// EventNameBatchTransfersApprovalRejected 对应结构体: BatchTransfersApprovalRejectedEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersApprovalRejected EventName = "payout.batch_transfers.approval_rejected"
	// EventNameBatchTransfersApprovalBlocked occurs when batch transfer is blocked as next approver cannot be found.
	// EventNameBatchTransfersApprovalBlocked 对应结构体: BatchTransfersApprovalBlockedEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersApprovalBlocked EventName = "payout.batch_transfers.approval_blocked"
	// EventNameBatchTransfersScheduled occurs when the batch transfer is scheduled and will be processed to book transfers once it is funded.
	// EventNameBatchTransfersScheduled 对应结构体: BatchTransfersScheduledEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersScheduled EventName = "payout.batch_transfers.scheduled"
	// EventNameBatchTransfersOverdue occurs when the batch transfer has not been funded beyond the scheduled date.
	// EventNameBatchTransfersOverdue 对应结构体: BatchTransfersOverdueEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersOverdue EventName = "payout.batch_transfers.overdue"
	// EventNameBatchTransfersBooking occurs when the batch transfer is being processed to book transfers.
	// EventNameBatchTransfersBooking 对应结构体: BatchTransfersBookingEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersBooking EventName = "payout.batch_transfers.booking"
	// EventNameBatchTransfersPartiallyBooked occurs when some items in the batch failed to be booked.
	// EventNameBatchTransfersPartiallyBooked 对应结构体: BatchTransfersPartiallyBookedEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersPartiallyBooked EventName = "payout.batch_transfers.partially_booked"
	// EventNameBatchTransfersBooked occurs when all valid items in the batch transfer has been booked.
	// EventNameBatchTransfersBooked 对应结构体: BatchTransfersBookedEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersBooked EventName = "payout.batch_transfers.booked"
	// EventNameBatchTransfersFailed occurs when the batch transfer failed to be booked.
	// EventNameBatchTransfersFailed 对应结构体: BatchTransfersFailedEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersFailed EventName = "payout.batch_transfers.failed"
	// EventNameBatchTransfersCancellationRequested occurs when batch transfer cancellation has been requested before BOOKING.
	// EventNameBatchTransfersCancellationRequested 对应结构体: BatchTransfersCancellationRequestedEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersCancellationRequested EventName = "payout.batch_transfers.cancellation_requested"
	// EventNameBatchTransfersCancelled occurs when the batch transfer has been successfully cancelled.
	// EventNameBatchTransfersCancelled 对应结构体: BatchTransfersCancelledEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersCancelled EventName = "payout.batch_transfers.cancelled"
	// EventNameBatchTransfersFundingScheduled occurs when the funding is scheduled to be processed on the transfer date.
	// EventNameBatchTransfersFundingScheduled 对应结构体: BatchTransfersFundingScheduledEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersFundingScheduled EventName = "payout.batch_transfers.funding.scheduled"
	// EventNameBatchTransfersFundingProcessing occurs when the funding for this batch transfer is being processed.
	// EventNameBatchTransfersFundingProcessing 对应结构体: BatchTransfersFundingProcessingEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersFundingProcessing EventName = "payout.batch_transfers.funding.processing"
	// EventNameBatchTransfersFundingFunded occurs when this batch transfer is successfully funded.
	// EventNameBatchTransfersFundingFunded 对应结构体: BatchTransfersFundingFundedEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersFundingFunded EventName = "payout.batch_transfers.funding.funded"
	// EventNameBatchTransfersFundingCancelled occurs when the funding for this batch transfer is cancelled.
	// EventNameBatchTransfersFundingCancelled 对应结构体: BatchTransfersFundingCancelledEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersFundingCancelled EventName = "payout.batch_transfers.funding.cancelled"
	// EventNameBatchTransfersFundingFailed occurs when deposit via direct debit from a Linked Account is rejected.
	// EventNameBatchTransfersFundingFailed 对应结构体: BatchTransfersFundingFailedEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersFundingFailed EventName = "payout.batch_transfers.funding.failed"
	// EventNameBatchTransfersFundingReversed occurs when the funding status has been updated to REVERSED.
	// EventNameBatchTransfersFundingReversed 对应结构体: BatchTransfersFundingReversedEvent (Data: BatchTransferEventData)
	EventNameBatchTransfersFundingReversed EventName = "payout.batch_transfers.funding.reversed"
)

// --- Conversions ---
const (
	// EventNameConversionScheduled occurs when the conversion has been booked and scheduled for settlement.
	// EventNameConversionScheduled 对应结构体: ConversionScheduledEvent (Data: ConversionEventData)
	EventNameConversionScheduled EventName = "conversion.scheduled"
	// EventNameConversionOverdue occurs when the settlement attempt has failed due to insufficient funds past settlement cut-off time.
	// EventNameConversionOverdue 对应结构体: ConversionOverdueEvent (Data: ConversionEventData)
	EventNameConversionOverdue EventName = "conversion.overdue"
	// EventNameConversionSettled occurs when funds have been settled in the wallet.
	// EventNameConversionSettled 对应结构体: ConversionSettledEvent (Data: ConversionEventData)
	EventNameConversionSettled EventName = "conversion.settled"
	// EventNameConversionCancelled occurs when the conversion was cancelled and funds have returned to the client's wallet.
	// EventNameConversionCancelled 对应结构体: ConversionCancelledEvent (Data: ConversionEventData)
	EventNameConversionCancelled EventName = "conversion.cancelled"
)

// --- RFI ---
const (
	// EventNameRFIActionRequired occurs when RFI is pending answers.
	// EventNameRFIActionRequired 对应结构体: RFIActionRequiredEvent (Data: RFIActionRequiredEventData)
	EventNameRFIActionRequired EventName = "rfi.action_required"
	// EventNameRFIAnswered occurs when RFI is answered.
	// EventNameRFIAnswered 对应结构体: RFIAnsweredEvent (Data: RFIAnsweredEventData)
	EventNameRFIAnswered EventName = "rfi.answered"
	// EventNameRFIClosed occurs when RFI is finished and closed.
	// EventNameRFIClosed 对应结构体: RFIClosedEvent (Data: RFIClosedEventData)
	EventNameRFIClosed EventName = "rfi.closed"
)

// --- Tax: Tax Form ---
const (
	// EventNameTaxFormTaxIDVerificationSubmitted occurs when a tax form is submitted and pending verification.
	// EventNameTaxFormTaxIDVerificationSubmitted 对应结构体: TaxFormTaxIDVerificationSubmittedEvent (Data: TaxForm)
	EventNameTaxFormTaxIDVerificationSubmitted EventName = "tax.tax_form.tax_id_verification_submitted"
	// EventNameTaxFormTaxIDVerificationFailed occurs when a tax form's tax_id verification has failed.
	// EventNameTaxFormTaxIDVerificationFailed 对应结构体: TaxFormTaxIDVerificationFailedEvent (Data: TaxForm)
	EventNameTaxFormTaxIDVerificationFailed EventName = "tax.tax_form.tax_id_verification_failed"
	// EventNameTaxFormReady occurs when a tax form's tax_id verification is successful and is active or ready for submission.
	// EventNameTaxFormReady 对应结构体: TaxFormReadyEvent (Data: TaxForm)
	EventNameTaxFormReady EventName = "tax.tax_form.ready"
	// EventNameTaxFormExpired occurs when a tax form has expired.
	// EventNameTaxFormExpired 对应结构体: TaxFormExpiredEvent (Data: TaxForm)
	EventNameTaxFormExpired EventName = "tax.tax_form.expired"
)

// --- Platform Liquidity Program ---
const (
	// EventNamePlatformLiquidityProgramLowBalance occurs when available balance for Platform Liquidity Program has fallen below the threshold.
	// EventNamePlatformLiquidityProgramLowBalance 对应结构体: PlatformLiquidityProgramLowBalanceEvent (Data: PlatformLiquidityProgramLowBalanceData)
	EventNamePlatformLiquidityProgramLowBalance EventName = "platform_liquidity_program.low_balance"
)

// --- Connected Account Transfers ---
const (
	// EventNameConnectedAccountTransferNew occurs when a connected account transfer has been created.
	// EventNameConnectedAccountTransferNew 对应结构体: ConnectedAccountTransferNewEvent (Data: ConnectedAccountTransfer)
	EventNameConnectedAccountTransferNew EventName = "connected_account_transfer.new"
	// EventNameConnectedAccountTransferPending occurs when a connected account transfer is pending.
	// EventNameConnectedAccountTransferPending 对应结构体: ConnectedAccountTransferPendingEvent (Data: ConnectedAccountTransfer)
	EventNameConnectedAccountTransferPending EventName = "connected_account_transfer.pending"
	// EventNameConnectedAccountTransferSettled occurs when a connected account transfer has settled.
	// EventNameConnectedAccountTransferSettled 对应结构体: ConnectedAccountTransferSettledEvent (Data: ConnectedAccountTransfer)
	EventNameConnectedAccountTransferSettled EventName = "connected_account_transfer.settled"
	// EventNameConnectedAccountTransferSuspended occurs when a connected account transfer has been suspended.
	// EventNameConnectedAccountTransferSuspended 对应结构体: ConnectedAccountTransferSuspendedEvent (Data: ConnectedAccountTransfer)
	EventNameConnectedAccountTransferSuspended EventName = "connected_account_transfer.suspended"
	// EventNameConnectedAccountTransferFailed occurs when a connected account transfer has failed.
	// EventNameConnectedAccountTransferFailed 对应结构体: ConnectedAccountTransferFailedEvent (Data: ConnectedAccountTransfer)
	EventNameConnectedAccountTransferFailed EventName = "connected_account_transfer.failed"
)

// --- PSP Agnostic: PSP Settlement Intent ---
const (
	// EventNamePSPSettlementIntentNew occurs when a PSP settlement intent has been successfully created.
	// EventNamePSPSettlementIntentNew 对应结构体: PSPSettlementIntentNewEvent (Data: PSPSettlementIntent)
	EventNamePSPSettlementIntentNew EventName = "psp_settlement_intent.new"
	// EventNamePSPSettlementIntentCancelled occurs when a PSP settlement intent has been successfully canceled.
	// EventNamePSPSettlementIntentCancelled 对应结构体: PSPSettlementIntentCancelledEvent (Data: PSPSettlementIntent)
	EventNamePSPSettlementIntentCancelled EventName = "psp_settlement_intent.cancelled"
	// EventNamePSPSettlementIntentSubmitted occurs when a PSP settlement intent has been successfully submitted.
	// EventNamePSPSettlementIntentSubmitted 对应结构体: PSPSettlementIntentSubmittedEvent (Data: PSPSettlementIntent)
	EventNamePSPSettlementIntentSubmitted EventName = "psp_settlement_intent.submitted"
	// EventNamePSPSettlementIntentActionRequired occurs when a PSP settlement intent does not match the associated PSP settlement deposits.
	// EventNamePSPSettlementIntentActionRequired 对应结构体: PSPSettlementIntentActionRequiredEvent (Data: PSPSettlementIntent)
	EventNamePSPSettlementIntentActionRequired EventName = "psp_settlement_intent.action_required"
	// EventNamePSPSettlementIntentMatched occurs when a PSP settlement intent matches the associated PSP settlement deposits.
	// EventNamePSPSettlementIntentMatched 对应结构体: PSPSettlementIntentMatchedEvent (Data: PSPSettlementIntent)
	EventNamePSPSettlementIntentMatched EventName = "psp_settlement_intent.matched"
	// EventNamePSPSettlementIntentSettled occurs when all the PSP settlement splits of the PSP settlement intent have been settled.
	// EventNamePSPSettlementIntentSettled 对应结构体: PSPSettlementIntentSettledEvent (Data: PSPSettlementIntent)
	EventNamePSPSettlementIntentSettled EventName = "psp_settlement_intent.settled"
)

// --- PSP Agnostic: PSP Settlement Deposit ---
const (
	// EventNamePSPSettlementDepositNew occurs when funds have been deposited into a holding account.
	// EventNamePSPSettlementDepositNew 对应结构体: PSPSettlementDepositNewEvent (Data: PSPSettlementDeposit)
	EventNamePSPSettlementDepositNew EventName = "psp_settlement_deposit.new"
	// EventNamePSPSettlementDepositActionRequired occurs when a PSP settlement deposit does not match any PSP settlement intent.
	// EventNamePSPSettlementDepositActionRequired 对应结构体: PSPSettlementDepositActionRequiredEvent (Data: PSPSettlementDeposit)
	EventNamePSPSettlementDepositActionRequired EventName = "psp_settlement_deposit.action_required"
	// EventNamePSPSettlementDepositMatched occurs when PSP settlement deposits match the associated PSP settlement intent.
	// EventNamePSPSettlementDepositMatched 对应结构体: PSPSettlementDepositMatchedEvent (Data: PSPSettlementDeposit)
	EventNamePSPSettlementDepositMatched EventName = "psp_settlement_deposit.matched"
	// EventNamePSPSettlementDepositSettled occurs when all the PSP settlement splits associated with this PSP settlement deposit have been settled.
	// EventNamePSPSettlementDepositSettled 对应结构体: PSPSettlementDepositSettledEvent (Data: PSPSettlementDeposit)
	EventNamePSPSettlementDepositSettled EventName = "psp_settlement_deposit.settled"
)

// --- PSP Agnostic: PSP Settlement Split ---
const (
	// EventNamePSPSettlementSplitNew occurs when a PSP settlement split has been successfully created.
	// EventNamePSPSettlementSplitNew 对应结构体: PSPSettlementSplitNewEvent (Data: PSPSettlementSplit)
	EventNamePSPSettlementSplitNew EventName = "psp_settlement_split.new"
	// EventNamePSPSettlementSplitCreateFailed occurs when a PSP settlement split could not be created.
	// EventNamePSPSettlementSplitCreateFailed 对应结构体: PSPSettlementSplitCreateFailedEvent (Data: PSPSettlementSplit)
	EventNamePSPSettlementSplitCreateFailed EventName = "psp_settlement_split.create_failed"
	// EventNamePSPSettlementSplitCancelled occurs when a PSP settlement split has been successfully canceled.
	// EventNamePSPSettlementSplitCancelled 对应结构体: PSPSettlementSplitCancelledEvent (Data: PSPSettlementSplit)
	EventNamePSPSettlementSplitCancelled EventName = "psp_settlement_split.cancelled"
	// EventNamePSPSettlementSplitMatched occurs when a PSP settlement split is part of a PSP settlement intent that has been matched with its associated PSP deposits.
	// EventNamePSPSettlementSplitMatched 对应结构体: PSPSettlementSplitMatchedEvent (Data: PSPSettlementSplit)
	EventNamePSPSettlementSplitMatched EventName = "psp_settlement_split.matched"
	// EventNamePSPSettlementSplitPending occurs when a PSP settlement split has been successfully queued for releasing funds to the target account.
	// EventNamePSPSettlementSplitPending 对应结构体: PSPSettlementSplitPendingEvent (Data: PSPSettlementSplit)
	EventNamePSPSettlementSplitPending EventName = "psp_settlement_split.pending"
	// EventNamePSPSettlementSplitFailed occurs when the release of a PSP settlement split has failed.
	// EventNamePSPSettlementSplitFailed 对应结构体: PSPSettlementSplitFailedEvent (Data: PSPSettlementSplit)
	EventNamePSPSettlementSplitFailed EventName = "psp_settlement_split.failed"
	// EventNamePSPSettlementSplitSettled occurs when the release of a PSP settlement split has succeeded.
	// EventNamePSPSettlementSplitSettled 对应结构体: PSPSettlementSplitSettledEvent (Data: PSPSettlementSplit)
	EventNamePSPSettlementSplitSettled EventName = "psp_settlement_split.settled"
)
