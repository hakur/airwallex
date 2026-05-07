// Package events provides typed webhook event structures for the online payments domain.
// Online Payments 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/online-payments.md
//
// 事件映射表:
//
//	payment_intent.created                          → PaymentIntentCreatedEvent                   (Data: PaymentIntentEventData)
//	payment_intent.requires_payment_method          → PaymentIntentRequiresPaymentMethodEvent     (Data: PaymentIntentEventData)
//	payment_intent.updated                          → PaymentIntentUpdatedEvent                   (Data: PaymentIntentEventData)
//	payment_intent.requires_capture                 → PaymentIntentRequiresCaptureEvent           (Data: PaymentIntentEventData)
//	payment_intent.requires_customer_action         → PaymentIntentRequiresCustomerActionEvent    (Data: PaymentIntentEventData)
//	payment_intent.pending                          → PaymentIntentPendingEvent                   (Data: PaymentIntentEventData)
//	payment_intent.pending_review                   → PaymentIntentPendingReviewEvent             (Data: PaymentIntentEventData)
//	payment_intent.succeeded                        → PaymentIntentSucceededEvent                 (Data: PaymentIntentEventData)
//	payment_intent.cancelled                        → PaymentIntentCancelledEvent                 (Data: PaymentIntentEventData)
//	payment_attempt.received                        → PaymentAttemptReceivedEvent                 (Data: PaymentAttemptEventData)
//	payment_attempt.authentication_failed           → PaymentAttemptAuthenticationFailedEvent     (Data: PaymentAttemptEventData)
//	payment_attempt.authentication_redirected       → PaymentAttemptAuthenticationRedirectedEvent (Data: PaymentAttemptEventData)
//	payment_attempt.pending_authorization           → PaymentAttemptPendingAuthorizationEvent     (Data: PaymentAttemptEventData)
//	payment_attempt.authorization_failed            → PaymentAttemptAuthorizationFailedEvent      (Data: PaymentAttemptEventData)
//	payment_attempt.authorized                      → PaymentAttemptAuthorizedEvent               (Data: PaymentAttemptEventData)
//	payment_attempt.capture_requested               → PaymentAttemptCaptureRequestedEvent         (Data: PaymentAttemptEventData)
//	payment_attempt.settled                         → PaymentAttemptSettledEvent                  (Data: PaymentAttemptEventData)
//	payment_attempt.paid                            → PaymentAttemptPaidEvent                     (Data: PaymentAttemptEventData)
//	payment_attempt.cancelled                       → PaymentAttemptCancelledEvent                (Data: PaymentAttemptEventData)
//	payment_attempt.expired                         → PaymentAttemptExpiredEvent                  (Data: PaymentAttemptEventData)
//	payment_attempt.risk_declined                   → PaymentAttemptRiskDeclinedEvent             (Data: PaymentAttemptEventData)
//	payment_attempt.failed_to_process               → PaymentAttemptFailedToProcessEvent          (Data: PaymentAttemptEventData)
//	payment_attempt.capture_failed                  → PaymentAttemptCaptureFailedEvent            (Data: PaymentAttemptEventData)
//	payment_consent.created                         → PaymentConsentCreatedEvent                  (Data: PaymentConsentEventData)
//	payment_consent.updated                         → PaymentConsentUpdatedEvent                  (Data: PaymentConsentEventData)
//	payment_consent.pending                         → PaymentConsentPendingEvent                  (Data: PaymentConsentEventData)
//	payment_consent.verified                        → PaymentConsentVerifiedEvent                 (Data: PaymentConsentEventData)
//	payment_consent.disabled                        → PaymentConsentDisabledEvent                 (Data: PaymentConsentEventData)
//	payment_consent.paused                          → PaymentConsentPausedEvent                   (Data: PaymentConsentEventData)
//	payment_consent.requires_payment_method         → PaymentConsentRequiresPaymentMethodEvent    (Data: PaymentConsentEventData)
//	payment_consent.requires_customer_action        → PaymentConsentRequiresCustomerActionEvent   (Data: PaymentConsentEventData)
//	payment_consent.verification_failed             → PaymentConsentVerificationFailedEvent       (Data: PaymentConsentEventData)
//	customer.created                                → CustomerCreatedEvent                        (Data: CustomerEventData)
//	customer.updated                                → CustomerUpdatedEvent                        (Data: CustomerEventData)
//	refund.received                                 → RefundReceivedEvent                         (Data: RefundEventData)
//	refund.accepted                                 → RefundAcceptedEvent                         (Data: RefundEventData)
//	refund.settled                                  → RefundSettledEvent                          (Data: RefundEventData)
//	refund.failed                                   → RefundFailedEvent                           (Data: RefundEventData)
//	payment_method.created                          → PaymentMethodCreatedEvent                   (Data: PaymentMethodEventData)
//	payment_method.updated                          → PaymentMethodUpdatedEvent                   (Data: PaymentMethodEventData)
//	payment_method.attached                         → PaymentMethodAttachedEvent                  (Data: PaymentMethodEventData)
//	payment_method.detached                         → PaymentMethodDetachedEvent                  (Data: PaymentMethodEventData)
//	payment_method.disabled                         → PaymentMethodDisabledEvent                  (Data: PaymentMethodEventData)
//	payment_dispute.requires_response               → PaymentDisputeRequiresResponseEvent         (Data: PaymentDisputeEventData)
//	payment_dispute.challenged                      → PaymentDisputeChallengedEvent               (Data: PaymentDisputeEventData)
//	payment_dispute.accepted                        → PaymentDisputeAcceptedEvent                 (Data: PaymentDisputeEventData)
//	payment_dispute.expired                         → PaymentDisputeExpiredEvent                  (Data: PaymentDisputeEventData)
//	payment_dispute.pending_closure                 → PaymentDisputePendingClosureEvent           (Data: PaymentDisputeEventData)
//	payment_dispute.pending_decision                → PaymentDisputePendingDecisionEvent          (Data: PaymentDisputeEventData)
//	payment_dispute.won                             → PaymentDisputeWonEvent                      (Data: PaymentDisputeEventData)
//	payment_dispute.lost                            → PaymentDisputeLostEvent                     (Data: PaymentDisputeEventData)
//	payment_dispute.reversed                        → PaymentDisputeReversedEvent                 (Data: PaymentDisputeEventData)
//	payment_link.created                            → PaymentLinkCreatedEvent                     (Data: PaymentLinkEventData)
//	payment_link.paid                               → PaymentLinkPaidEvent                        (Data: PaymentLinkEventData)
//	funds_split.created                             → FundsSplitCreatedEvent                      (Data: FundsSplitEventData)
//	funds_split.failed                              → FundsSplitFailedEvent                       (Data: FundsSplitEventData)
//	funds_split.released                            → FundsSplitReleasedEvent                     (Data: FundsSplitEventData)
//	funds_split.settled                             → FundsSplitSettledEvent                      (Data: FundsSplitEventData)
//	fraud.merchant_notified                         → FraudMerchantNotifiedEvent                  (Data: FraudEventData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- PaymentIntent Events ---

// PaymentIntentCreatedEvent represents the payment_intent.created webhook event.
type PaymentIntentCreatedEvent struct {
	Event
	Data PaymentIntentEventData `json:"data"`
}

// PaymentIntentRequiresPaymentMethodEvent represents the payment_intent.requires_payment_method webhook event.
type PaymentIntentRequiresPaymentMethodEvent struct {
	Event
	Data PaymentIntentEventData `json:"data"`
}

// PaymentIntentUpdatedEvent represents the payment_intent.updated webhook event.
type PaymentIntentUpdatedEvent struct {
	Event
	Data PaymentIntentEventData `json:"data"`
}

// PaymentIntentRequiresCaptureEvent represents the payment_intent.requires_capture webhook event.
type PaymentIntentRequiresCaptureEvent struct {
	Event
	Data PaymentIntentEventData `json:"data"`
}

// PaymentIntentRequiresCustomerActionEvent represents the payment_intent.requires_customer_action webhook event.
type PaymentIntentRequiresCustomerActionEvent struct {
	Event
	Data PaymentIntentEventData `json:"data"`
}

// PaymentIntentPendingEvent represents the payment_intent.pending webhook event.
type PaymentIntentPendingEvent struct {
	Event
	Data PaymentIntentEventData `json:"data"`
}

// PaymentIntentPendingReviewEvent represents the payment_intent.pending_review webhook event.
type PaymentIntentPendingReviewEvent struct {
	Event
	Data PaymentIntentEventData `json:"data"`
}

// PaymentIntentSucceededEvent represents the payment_intent.succeeded webhook event.
type PaymentIntentSucceededEvent struct {
	Event
	Data PaymentIntentEventData `json:"data"`
}

// PaymentIntentCancelledEvent represents the payment_intent.cancelled webhook event.
type PaymentIntentCancelledEvent struct {
	Event
	Data PaymentIntentEventData `json:"data"`
}

// PaymentIntentEventData wraps the PaymentIntent object in webhook data.
type PaymentIntentEventData struct {
	Object PaymentIntent `json:"object"`
}

// PaymentIntent represents a payment intent (from webhook payload).
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/online-payments.md
type PaymentIntent struct {
	ID                   string               `json:"id"`
	RequestID            string               `json:"request_id,omitempty"`
	MerchantOrderID      string               `json:"merchant_order_id,omitempty"`
	Amount               float64              `json:"amount"`
	Currency             string               `json:"currency"`
	CapturedAmount       float64              `json:"captured_amount,omitempty"`
	PaymentMethodOptions PaymentMethodOptions `json:"payment_method_options,omitempty"`
	Status               string               `json:"status"`
	CreatedAt            string               `json:"created_at"`
	UpdatedAt            string               `json:"updated_at,omitempty"`
}

// PaymentMethodOptions represents payment method options.
type PaymentMethodOptions struct {
	Card CardOptions `json:"card,omitempty"`
}

// CardOptions represents card-specific options.
type CardOptions struct {
	RiskControl   RiskControl `json:"risk_control,omitempty"`
	ThreeDSAction string      `json:"three_ds_action,omitempty"`
}

// RiskControl represents risk control settings.
type RiskControl struct {
	ThreeDomainSecureAction string `json:"three_domain_secure_action,omitempty"`
	ThreeDSAction           string `json:"three_ds_action,omitempty"`
}

// --- PaymentAttempt Events ---

// PaymentAttemptReceivedEvent represents the payment_attempt.received webhook event.
type PaymentAttemptReceivedEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptAuthenticationFailedEvent represents the payment_attempt.authentication_failed webhook event.
type PaymentAttemptAuthenticationFailedEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptAuthenticationRedirectedEvent represents the payment_attempt.authentication_redirected webhook event.
type PaymentAttemptAuthenticationRedirectedEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptPendingAuthorizationEvent represents the payment_attempt.pending_authorization webhook event.
type PaymentAttemptPendingAuthorizationEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptAuthorizationFailedEvent represents the payment_attempt.authorization_failed webhook event.
type PaymentAttemptAuthorizationFailedEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptAuthorizedEvent represents the payment_attempt.authorized webhook event.
type PaymentAttemptAuthorizedEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptCaptureRequestedEvent represents the payment_attempt.capture_requested webhook event.
type PaymentAttemptCaptureRequestedEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptSettledEvent represents the payment_attempt.settled webhook event.
type PaymentAttemptSettledEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptPaidEvent represents the payment_attempt.paid webhook event.
type PaymentAttemptPaidEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptCancelledEvent represents the payment_attempt.cancelled webhook event.
type PaymentAttemptCancelledEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptExpiredEvent represents the payment_attempt.expired webhook event.
type PaymentAttemptExpiredEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptRiskDeclinedEvent represents the payment_attempt.risk_declined webhook event.
type PaymentAttemptRiskDeclinedEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptFailedToProcessEvent represents the payment_attempt.failed_to_process webhook event.
type PaymentAttemptFailedToProcessEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptCaptureFailedEvent represents the payment_attempt.capture_failed webhook event.
type PaymentAttemptCaptureFailedEvent struct {
	Event
	Data PaymentAttemptEventData `json:"data"`
}

// PaymentAttemptEventData wraps the PaymentAttempt object in webhook data.
type PaymentAttemptEventData struct {
	Object PaymentAttempt `json:"object"`
}

// PaymentAttempt represents a payment attempt (from webhook payload).
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/online-payments.md
type PaymentAttempt struct {
	ID                 string               `json:"id"`
	PaymentIntentID    string               `json:"payment_intent_id"`
	PaymentConsentID   string               `json:"payment_consent_id,omitempty"`
	Currency           string               `json:"currency"`
	Amount             float64              `json:"amount"`
	PaymentMethod      WebhookPaymentMethod `json:"payment_method,omitempty"`
	AuthenticationData AuthenticationData   `json:"authentication_data,omitempty"`
	CapturedAmount     float64              `json:"captured_amount,omitempty"`
	RefundedAmount     float64              `json:"refunded_amount,omitempty"`
	SettleVia          string               `json:"settle_via,omitempty"`
	Status             string               `json:"status"`
	CreatedAt          string               `json:"created_at"`
	UpdatedAt          string               `json:"updated_at,omitempty"`
}

// WebhookPaymentMethod represents a payment method in webhook payload.
type WebhookPaymentMethod struct {
	ID         string   `json:"id,omitempty"`
	CustomerID string   `json:"customer_id,omitempty"`
	Type       string   `json:"type,omitempty"`
	Card       CardInfo `json:"card,omitempty"`
	Status     string   `json:"status,omitempty"`
	CreatedAt  string   `json:"created_at,omitempty"`
	UpdatedAt  string   `json:"updated_at,omitempty"`
}

// CardInfo represents card information.
type CardInfo struct {
	Bin               string `json:"bin,omitempty"`
	Brand             string `json:"brand,omitempty"`
	CardType          string `json:"card_type,omitempty"`
	ExpiryMonth       string `json:"expiry_month,omitempty"`
	ExpiryYear        string `json:"expiry_year,omitempty"`
	Fingerprint       string `json:"fingerprint,omitempty"`
	IsCommercial      bool   `json:"is_commercial,omitempty"`
	IssuerCountryCode string `json:"issuer_country_code,omitempty"`
	IssuerName        string `json:"issuer_name,omitempty"`
	Last4             string `json:"last4,omitempty"`
	Name              string `json:"name,omitempty"`
	NumberType        string `json:"number_type,omitempty"`
}

// AuthenticationData represents authentication information.
type AuthenticationData struct {
	AVSResult string         `json:"avs_result,omitempty"`
	CVCResult string         `json:"cvc_result,omitempty"`
	DSData    map[string]any `json:"ds_data,omitempty"` // 官方文档显示为 {}，类型未明确
	FraudData FraudData      `json:"fraud_data,omitempty"`
}

// FraudData represents fraud check information.
type FraudData struct {
	Score string `json:"score,omitempty"`
}

// --- PaymentConsent Events ---

// PaymentConsentCreatedEvent represents the payment_consent.created webhook event.
type PaymentConsentCreatedEvent struct {
	Event
	Data PaymentConsentEventData `json:"data"`
}

// PaymentConsentUpdatedEvent represents the payment_consent.updated webhook event.
type PaymentConsentUpdatedEvent struct {
	Event
	Data PaymentConsentEventData `json:"data"`
}

// PaymentConsentPendingEvent represents the payment_consent.pending webhook event.
type PaymentConsentPendingEvent struct {
	Event
	Data PaymentConsentEventData `json:"data"`
}

// PaymentConsentVerifiedEvent represents the payment_consent.verified webhook event.
type PaymentConsentVerifiedEvent struct {
	Event
	Data PaymentConsentEventData `json:"data"`
}

// PaymentConsentDisabledEvent represents the payment_consent.disabled webhook event.
type PaymentConsentDisabledEvent struct {
	Event
	Data PaymentConsentEventData `json:"data"`
}

// PaymentConsentPausedEvent represents the payment_consent.paused webhook event.
type PaymentConsentPausedEvent struct {
	Event
	Data PaymentConsentEventData `json:"data"`
}

// PaymentConsentRequiresPaymentMethodEvent represents the payment_consent.requires_payment_method webhook event.
type PaymentConsentRequiresPaymentMethodEvent struct {
	Event
	Data PaymentConsentEventData `json:"data"`
}

// PaymentConsentRequiresCustomerActionEvent represents the payment_consent.requires_customer_action webhook event.
type PaymentConsentRequiresCustomerActionEvent struct {
	Event
	Data PaymentConsentEventData `json:"data"`
}

// PaymentConsentVerificationFailedEvent represents the payment_consent.verification_failed webhook event.
type PaymentConsentVerificationFailedEvent struct {
	Event
	Data PaymentConsentEventData `json:"data"`
}

// PaymentConsentEventData wraps the PaymentConsent object in webhook data.
type PaymentConsentEventData struct {
	Object PaymentConsent `json:"object"`
}

// PaymentConsent represents a payment consent (from webhook payload).
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/online-payments.md
type PaymentConsent struct {
	ID                    string               `json:"id"`
	RequestID             string               `json:"request_id,omitempty"`
	CustomerID            string               `json:"customer_id"`
	MerchantTriggerReason string               `json:"merchant_trigger_reason,omitempty"`
	NextTriggeredBy       string               `json:"next_triggered_by,omitempty"`
	PaymentMethod         ConsentPaymentMethod `json:"payment_method,omitempty"`
	Status                string               `json:"status"`
	CreatedAt             string               `json:"created_at"`
	UpdatedAt             string               `json:"updated_at,omitempty"`
}

// ConsentPaymentMethod represents a payment method within a consent.
type ConsentPaymentMethod struct {
	ID   string   `json:"id,omitempty"`
	Card CardInfo `json:"card,omitempty"`
	Type string   `json:"type,omitempty"`
}

// --- Customer Events ---

// CustomerCreatedEvent represents the customer.created webhook event.
type CustomerCreatedEvent struct {
	Event
	Data CustomerEventData `json:"data"`
}

// CustomerUpdatedEvent represents the customer.updated webhook event.
type CustomerUpdatedEvent struct {
	Event
	Data CustomerEventData `json:"data"`
}

// CustomerEventData wraps the Customer object in webhook data.
type CustomerEventData struct {
	Object Customer `json:"object"`
}

// Customer represents a customer (from webhook payload).
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/online-payments.md
type Customer struct {
	ID                 string         `json:"id"`
	Email              string         `json:"email,omitempty"`
	FirstName          string         `json:"first_name,omitempty"`
	LastName           string         `json:"last_name,omitempty"`
	MerchantCustomerID string         `json:"merchant_customer_id,omitempty"`
	PhoneNumber        string         `json:"phone_number,omitempty"`
	RequestID          string         `json:"request_id,omitempty"`
	AdditionalInfo     AdditionalInfo `json:"additional_info,omitempty"`
	CreatedAt          string         `json:"created_at"`
	UpdatedAt          string         `json:"updated_at,omitempty"`
}

// AdditionalInfo represents additional customer information.
type AdditionalInfo struct {
	RegisteredViaSocialMedia bool `json:"registered_via_social_media,omitempty"`
}

// --- Refund Events ---

// RefundReceivedEvent represents the refund.received webhook event.
type RefundReceivedEvent struct {
	Event
	Data RefundEventData `json:"data"`
}

// RefundAcceptedEvent represents the refund.accepted webhook event.
type RefundAcceptedEvent struct {
	Event
	Data RefundEventData `json:"data"`
}

// RefundSettledEvent represents the refund.settled webhook event.
type RefundSettledEvent struct {
	Event
	Data RefundEventData `json:"data"`
}

// RefundFailedEvent represents the refund.failed webhook event.
type RefundFailedEvent struct {
	Event
	Data RefundEventData `json:"data"`
}

// RefundEventData wraps the Refund object in webhook data.
type RefundEventData struct {
	Object Refund `json:"object"`
}

// Refund represents a refund (from webhook payload).
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/online-payments.md
type Refund struct {
	RequestID        string  `json:"request_id,omitempty"`
	ID               string  `json:"id"`
	PaymentAttemptID string  `json:"payment_attempt_id,omitempty"`
	PaymentIntentID  string  `json:"payment_intent_id,omitempty"`
	Amount           float64 `json:"amount"`
	Currency         string  `json:"currency"`
	Reason           string  `json:"reason,omitempty"`
	Status           string  `json:"status"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at,omitempty"`
}

// --- PaymentMethod Events ---

// PaymentMethodCreatedEvent represents the payment_method.created webhook event.
type PaymentMethodCreatedEvent struct {
	Event
	Data PaymentMethodEventData `json:"data"`
}

// PaymentMethodUpdatedEvent represents the payment_method.updated webhook event.
type PaymentMethodUpdatedEvent struct {
	Event
	Data PaymentMethodEventData `json:"data"`
}

// PaymentMethodAttachedEvent represents the payment_method.attached webhook event.
type PaymentMethodAttachedEvent struct {
	Event
	Data PaymentMethodEventData `json:"data"`
}

// PaymentMethodDetachedEvent represents the payment_method.detached webhook event.
type PaymentMethodDetachedEvent struct {
	Event
	Data PaymentMethodEventData `json:"data"`
}

// PaymentMethodDisabledEvent represents the payment_method.disabled webhook event.
type PaymentMethodDisabledEvent struct {
	Event
	Data PaymentMethodEventData `json:"data"`
}

// PaymentMethodEventData wraps the PaymentMethod object in webhook data.
type PaymentMethodEventData struct {
	Object PaymentMethod `json:"object"`
}

// PaymentMethod represents a payment method (from webhook payload).
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/online-payments.md
type PaymentMethod struct {
	ID         string   `json:"id,omitempty"`
	RequestID  string   `json:"request_id,omitempty"`
	CustomerID string   `json:"customer_id,omitempty"`
	Type       string   `json:"type,omitempty"`
	Card       CardInfo `json:"card,omitempty"`
	Status     string   `json:"status,omitempty"`
	CreatedAt  string   `json:"created_at,omitempty"`
	UpdatedAt  string   `json:"updated_at,omitempty"`
}

// --- PaymentDispute Events ---

// PaymentDisputeRequiresResponseEvent represents the payment_dispute.requires_response webhook event.
type PaymentDisputeRequiresResponseEvent struct {
	Event
	Data PaymentDisputeEventData `json:"data"`
}

// PaymentDisputeChallengedEvent represents the payment_dispute.challenged webhook event.
type PaymentDisputeChallengedEvent struct {
	Event
	Data PaymentDisputeEventData `json:"data"`
}

// PaymentDisputeAcceptedEvent represents the payment_dispute.accepted webhook event.
type PaymentDisputeAcceptedEvent struct {
	Event
	Data PaymentDisputeEventData `json:"data"`
}

// PaymentDisputeExpiredEvent represents the payment_dispute.expired webhook event.
type PaymentDisputeExpiredEvent struct {
	Event
	Data PaymentDisputeEventData `json:"data"`
}

// PaymentDisputePendingClosureEvent represents the payment_dispute.pending_closure webhook event.
type PaymentDisputePendingClosureEvent struct {
	Event
	Data PaymentDisputeEventData `json:"data"`
}

// PaymentDisputePendingDecisionEvent represents the payment_dispute.pending_decision webhook event.
type PaymentDisputePendingDecisionEvent struct {
	Event
	Data PaymentDisputeEventData `json:"data"`
}

// PaymentDisputeWonEvent represents the payment_dispute.won webhook event.
type PaymentDisputeWonEvent struct {
	Event
	Data PaymentDisputeEventData `json:"data"`
}

// PaymentDisputeLostEvent represents the payment_dispute.lost webhook event.
type PaymentDisputeLostEvent struct {
	Event
	Data PaymentDisputeEventData `json:"data"`
}

// PaymentDisputeReversedEvent represents the payment_dispute.reversed webhook event.
type PaymentDisputeReversedEvent struct {
	Event
	Data PaymentDisputeEventData `json:"data"`
}

// PaymentDisputeEventData wraps the PaymentDispute object in webhook data.
type PaymentDisputeEventData struct {
	Object PaymentDispute `json:"object"`
}

// PaymentDispute represents a payment dispute (from webhook payload).
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/online-payments.md
type PaymentDispute struct {
	DisputeID                 string  `json:"dispute_id"`
	PaymentIntentID           string  `json:"payment_intent_id,omitempty"`
	PaymentAttemptID          string  `json:"payment_attempt_id,omitempty"`
	MerchantAccount           string  `json:"merchant_account,omitempty"`
	CardScheme                string  `json:"card_scheme,omitempty"`
	DisputeReasonType         string  `json:"dispute_reason_type,omitempty"`
	DisputeOriginalReasonCode string  `json:"dispute_original_reason_code,omitempty"`
	AcceptReason              *string `json:"accept_reason,omitempty"` // 官方示例为 null
	Status                    string  `json:"status"`
	Stage                     string  `json:"stage,omitempty"`
	DisputeAmount             float64 `json:"dispute_amount"`
	DisputeCurrency           string  `json:"dispute_currency"`
	UpdatedBy                 string  `json:"updated_by,omitempty"`
	UpdatedAt                 string  `json:"updated_at"`
	CreatedAt                 string  `json:"created_at"`
}

// --- PaymentLink Events ---

// PaymentLinkCreatedEvent represents the payment_link.created webhook event.
type PaymentLinkCreatedEvent struct {
	Event
	Data PaymentLinkEventData `json:"data"`
}

// PaymentLinkPaidEvent represents the payment_link.paid webhook event.
type PaymentLinkPaidEvent struct {
	Event
	Data PaymentLinkEventData `json:"data"`
}

// PaymentLinkEventData wraps the PaymentLink object in webhook data.
type PaymentLinkEventData struct {
	Object PaymentLink `json:"object"`
}

// PaymentLink represents a payment link (from webhook payload).
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/online-payments.md
type PaymentLink struct {
	Active                       bool                   `json:"active,omitempty"`
	Amount                       float64                `json:"amount,omitempty"`
	CollectableShopperInfo       CollectableShopperInfo `json:"collectable_shopper_info,omitempty"`
	CreatedAt                    string                 `json:"created_at,omitempty"`
	Currency                     string                 `json:"currency,omitempty"`
	CustomerID                   string                 `json:"customer_id,omitempty"`
	Description                  string                 `json:"description,omitempty"`
	ExpiresAt                    string                 `json:"expires_at,omitempty"`
	ID                           string                 `json:"id"`
	Metadata                     map[string]any         `json:"metadata,omitempty"` // 官方示例为 {"foo": "bar"}
	Reference                    string                 `json:"reference,omitempty"`
	Reusable                     bool                   `json:"reusable,omitempty"`
	Status                       string                 `json:"status,omitempty"`
	SuccessfulPaymentIntentCount int                    `json:"successful_payment_intent_count,omitempty"`
	SupportedCurrencies          []string               `json:"supported_currencies,omitempty"`
	Title                        string                 `json:"title,omitempty"`
	UpdatedAt                    string                 `json:"updated_at,omitempty"`
	URL                          string                 `json:"url,omitempty"`
}

// CollectableShopperInfo represents collectable shopper information.
type CollectableShopperInfo struct {
	Message         bool `json:"message,omitempty"`
	PhoneNumber     bool `json:"phone_number,omitempty"`
	Reference       bool `json:"reference,omitempty"`
	ShippingAddress bool `json:"shipping_address,omitempty"`
}

// --- FundsSplit Events ---

// FundsSplitCreatedEvent represents the funds_split.created webhook event.
type FundsSplitCreatedEvent struct {
	Event
	Data FundsSplitEventData `json:"data"`
}

// FundsSplitFailedEvent represents the funds_split.failed webhook event.
type FundsSplitFailedEvent struct {
	Event
	Data FundsSplitEventData `json:"data"`
}

// FundsSplitReleasedEvent represents the funds_split.released webhook event.
type FundsSplitReleasedEvent struct {
	Event
	Data FundsSplitEventData `json:"data"`
}

// FundsSplitSettledEvent represents the funds_split.settled webhook event.
type FundsSplitSettledEvent struct {
	Event
	Data FundsSplitEventData `json:"data"`
}

// FundsSplitEventData wraps the FundsSplit object in webhook data.
type FundsSplitEventData struct {
	Object FundsSplit `json:"object"`
}

// FundsSplit represents a funds split (from webhook payload).
// TODO: 官方文档未提供 payload 示例，字段基于 API 推断，需验证
type FundsSplit struct {
	ID       string  `json:"id,omitempty"`
	Status   string  `json:"status,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

// --- Fraud Events ---

// FraudMerchantNotifiedEvent represents the fraud.merchant_notified webhook event.
type FraudMerchantNotifiedEvent struct {
	Event
	Data FraudEventData `json:"data"`
}

// FraudEventData wraps the fraud notification object in webhook data.
type FraudEventData struct {
	Object FraudNotification `json:"object"`
}

// FraudNotification represents a fraud notification (from webhook payload).
// TODO: 官方文档未提供 payload 示例，需补充精确字段
type FraudNotification struct {
	PaymentIntentID string `json:"payment_intent_id,omitempty"`
	Reason          string `json:"reason,omitempty"`
}
