package simulation

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// PaymentDisputeStatus represents the payment dispute status.
// PaymentDisputeStatus 支付争议状态。
type PaymentDisputeStatus = string

const (
	PaymentDisputeStatusOpen    PaymentDisputeStatus = "OPEN"
	PaymentDisputeStatusClosed  PaymentDisputeStatus = "CLOSED"
	PaymentDisputeStatusExpired PaymentDisputeStatus = "EXPIRED"
	PaymentDisputeStatusWon     PaymentDisputeStatus = "WON"
	PaymentDisputeStatusLost    PaymentDisputeStatus = "LOST"
)

// PaymentDisputeStage represents the payment dispute stage.
// PaymentDisputeStage 支付争议阶段。
type PaymentDisputeStage = string

const (
	PaymentDisputeStagePreDispute     PaymentDisputeStage = "PRE_DISPUTE"
	PaymentDisputeStageDispute        PaymentDisputeStage = "DISPUTE"
	PaymentDisputeStagePreArbitration PaymentDisputeStage = "PRE_ARBITRATION"
	PaymentDisputeStageArbitration    PaymentDisputeStage = "ARBITRATION"
)

// PaymentDisputeMode represents the payment dispute mode.
// PaymentDisputeMode 支付争议模式。
type PaymentDisputeMode = string

const (
	PaymentDisputeModeCBPP PaymentDisputeMode = "CBPP"
	PaymentDisputeModeRPB  PaymentDisputeMode = "RPB"
)

// PaymentDispute represents a payment dispute response.
// PaymentDispute 表示支付争议响应。
type PaymentDispute struct {
	// ID 唯一标识符。可选。
	ID string `json:"id,omitempty"`
	// PaymentIntentID 支付意图ID。可选。
	PaymentIntentID string `json:"payment_intent_id,omitempty"`
	// PaymentAttemptID 支付尝试ID。可选。
	PaymentAttemptID string `json:"payment_attempt_id,omitempty"`
	// MerchantOrderID 商户订单ID。可选。
	MerchantOrderID string `json:"merchant_order_id,omitempty"`
	// CustomerID 客户ID。可选。
	CustomerID string `json:"customer_id,omitempty"`
	// CustomerName 客户名称。可选。
	CustomerName string `json:"customer_name,omitempty"`
	// Currency 币种。可选。
	Currency string `json:"currency,omitempty"`
	// Amount 金额。可选。
	Amount float64 `json:"amount,omitempty"`
	// Status 状态。可选。
	Status PaymentDisputeStatus `json:"status,omitempty"`
	// Stage 阶段。可选。
	Stage PaymentDisputeStage `json:"stage,omitempty"`
	// Mode 模式。可选。
	Mode PaymentDisputeMode `json:"mode,omitempty"`
	// TransactionType 交易类型。可选。
	TransactionType string `json:"transaction_type,omitempty"`
	// PaymentMethodType 支付方式类型。可选。
	PaymentMethodType string `json:"payment_method_type,omitempty"`
	// CardBrand 卡品牌。可选。
	CardBrand string `json:"card_brand,omitempty"`
	// AcquirerReferenceNumber 收单机构参考号。可选。
	AcquirerReferenceNumber string `json:"acquirer_reference_number,omitempty"`
	// DueAt 截止日期。可选。
	DueAt string `json:"due_at,omitempty"`
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
	// IssuerComment 发卡行备注。可选。
	IssuerComment string `json:"issuer_comment,omitempty"`
	// IssuerDocuments 发卡行文档列表。可选。
	IssuerDocuments []string `json:"issuer_documents,omitempty"`
	// Reason 争议原因。可选。
	Reason *DisputeReason `json:"reason,omitempty"`
	// AcceptDetails 接受详情。可选。
	AcceptDetails []AcceptDetail `json:"accept_details,omitempty"`
	// ChallengeDetails 挑战详情。可选。
	ChallengeDetails []ChallengeDetail `json:"challenge_details,omitempty"`
	// Refunds 退款列表。可选。
	Refunds []RefundSummary `json:"refunds,omitempty"`
}

// DisputeReason represents a dispute reason.
// DisputeReason 表示争议原因。
type DisputeReason struct {
	Description  string `json:"description,omitempty"`
	OriginalCode string `json:"original_code,omitempty"`
	Type         string `json:"type,omitempty"`
}

// AcceptDetail represents an accepted dispute detail.
// AcceptDetail 表示接受争议详情。
type AcceptDetail struct {
	AcceptedAt  string        `json:"accepted_at,omitempty"`
	AcceptedBy  string        `json:"accepted_by,omitempty"`
	Description string        `json:"description,omitempty"`
	Reason      string        `json:"reason,omitempty"`
	Stage       string        `json:"stage,omitempty"`
	Refund      *AcceptRefund `json:"refund,omitempty"`
}

// AcceptRefund represents an accepted dispute refund.
// AcceptRefund 表示接受争议退款信息。
type AcceptRefund struct {
	Amount float64 `json:"amount,omitempty"`
	Reason string  `json:"reason,omitempty"`
}

// ChallengeDetail represents a dispute challenge detail.
// ChallengeDetail 表示争议挑战详情。
type ChallengeDetail struct {
	ChallengedAt        string               `json:"challenged_at,omitempty"`
	ChallengedBy        string               `json:"challenged_by,omitempty"`
	Stage               string               `json:"stage,omitempty"`
	Reason              string               `json:"reason,omitempty"`
	RefundRefusalReason string               `json:"refund_refusal_reason,omitempty"`
	ProductDescription  string               `json:"product_description,omitempty"`
	ProductType         string               `json:"product_type,omitempty"`
	CustomerInfo        *CustomerInfo        `json:"customer_info,omitempty"`
	DeliveryInfo        *DeliveryInfo        `json:"delivery_info,omitempty"`
	OrderInfo           *OrderInfo           `json:"order_info,omitempty"`
	SellerInfo          *SellerInfo          `json:"seller_info,omitempty"`
	SupportingDocuments *SupportingDocuments `json:"supporting_documents,omitempty"`
}

// CustomerInfo represents customer information.
// CustomerInfo 表示客户信息。
type CustomerInfo struct {
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	PhoneNumber    string `json:"phone_number,omitempty"`
	BillingAddress string `json:"billing_address,omitempty"`
	IP             string `json:"ip,omitempty"`
	DeviceID       string `json:"device_id,omitempty"`
}

// DeliveryInfo represents delivery information.
// DeliveryInfo 表示配送信息。
type DeliveryInfo struct {
	Address         string  `json:"address,omitempty"`
	Name            string  `json:"name,omitempty"`
	PhoneNumber     string  `json:"phone_number,omitempty"`
	TrackingNumber  string  `json:"tracking_number,omitempty"`
	ShippingCompany string  `json:"shipping_company,omitempty"`
	ShippingMethod  string  `json:"shipping_method,omitempty"`
	Status          string  `json:"status,omitempty"`
	ShippedAt       string  `json:"shipped_at,omitempty"`
	DeliveredAt     string  `json:"delivered_at,omitempty"`
	FeeAmount       float64 `json:"fee_amount,omitempty"`
	FeeCurrency     string  `json:"fee_currency,omitempty"`
}

// OrderInfo represents order information.
// OrderInfo 表示订单信息。
type OrderInfo struct {
	ID            string    `json:"id,omitempty"`
	CreatedAt     string    `json:"created_at,omitempty"`
	InvoiceNumber string    `json:"invoice_number,omitempty"`
	TotalAmount   float64   `json:"total_amount,omitempty"`
	TotalCurrency string    `json:"total_currency,omitempty"`
	Products      []Product `json:"products,omitempty"`
}

// Product represents product information.
// Product 表示商品信息。
type Product struct {
	Name      string  `json:"name,omitempty"`
	Quantity  float64 `json:"quantity,omitempty"`
	UnitPrice float64 `json:"unit_price,omitempty"`
	Currency  string  `json:"currency,omitempty"`
}

// SellerInfo represents seller information.
// SellerInfo 表示卖家信息。
type SellerInfo struct {
	Name                 string `json:"name,omitempty"`
	StoreName            string `json:"store_name,omitempty"`
	StorePhysicalAddress string `json:"store_physical_address,omitempty"`
	StoreURL             string `json:"store_url,omitempty"`
}

// SupportingDocuments represents supporting documents.
// SupportingDocuments 表示支持文档。
type SupportingDocuments struct {
	Documents                       []Document `json:"documents,omitempty"`
	GeneratedFiles                  []string   `json:"generated_files,omitempty"`
	CustomerCommunicationDocuments  []string   `json:"customer_communication_documents,omitempty"`
	CustomerSignatureDocuments      []string   `json:"customer_signature_documents,omitempty"`
	DuplicateChargeDefenseDocuments []string   `json:"duplicate_charge_defense_documents,omitempty"`
	OtherDocuments                  []string   `json:"other_documents,omitempty"`
	ProofOfDeliveryDocuments        []string   `json:"proof_of_delivery_documents,omitempty"`
	ReceiptDocuments                []string   `json:"receipt_documents,omitempty"`
	RefundPolicyDocuments           []string   `json:"refund_policy_documents,omitempty"`
}

// Document represents document information.
// Document 表示文档信息。
type Document struct {
	Description string   `json:"description,omitempty"`
	FileIDs     []string `json:"file_ids,omitempty"`
	Type        string   `json:"type,omitempty"`
}

// RefundSummary represents a refund summary.
// RefundSummary 表示退款摘要。
type RefundSummary struct {
	ID                      string `json:"id,omitempty"`
	AcquirerReferenceNumber string `json:"acquirer_reference_number,omitempty"`
}

// SimulatePaymentDisputeCreateRequest represents a request to create a simulated payment dispute.
// SimulatePaymentDisputeCreateRequest 创建模拟支付争议请求。
type SimulatePaymentDisputeCreateRequest struct {
	PaymentIntentID string   `json:"payment_intent_id"`
	ReasonCode      string   `json:"reason_code"`
	Stage           string   `json:"stage"`
	DueAt           string   `json:"due_at"`
	Amount          float64  `json:"amount,omitempty"`
	Comment         string   `json:"comment,omitempty"`
	Documents       []string `json:"documents,omitempty"`
}

// SimulatePaymentDisputeEscalateRequest represents a request to escalate a simulated payment dispute.
// SimulatePaymentDisputeEscalateRequest 升级模拟支付争议请求。
type SimulatePaymentDisputeEscalateRequest struct {
	Amount    float64  `json:"amount,omitempty"`
	Comment   string   `json:"comment,omitempty"`
	Documents []string `json:"documents,omitempty"`
	DueAt     string   `json:"due_at"`
}

// SimulatePaymentDisputeResolveRequest represents a request to resolve a simulated payment dispute.
// SimulatePaymentDisputeResolveRequest 解决模拟支付争议请求。
type SimulatePaymentDisputeResolveRequest struct {
	InFavorOf string  `json:"in_favor_of"`
	Amount    float64 `json:"amount,omitempty"`
	Comment   string  `json:"comment,omitempty"`
}

// SimulateShopperActionRequest represents a request to simulate a shopper action.
// SimulateShopperActionRequest 模拟购物者操作请求。
type SimulateShopperActionRequest struct {
	URL string `json:"url"`
}

// SimulatePaymentDisputeCreate creates a simulated payment dispute.
// SimulatePaymentDisputeCreate 创建模拟支付争议。
// 官方文档: https://www.airwallex.com/docs/api/simulation/payments/create_payment_disputes.md
func (s *Service) SimulatePaymentDisputeCreate(ctx context.Context, req *SimulatePaymentDisputeCreateRequest, opts ...sdk.RequestOption) (*PaymentDispute, error) {
	var resp PaymentDispute
	err := s.doer.Do(ctx, "POST", "/api/v1/simulation/pa/payment_disputes/create", req, &resp, opts...)
	return &resp, err
}

// SimulatePaymentDisputeEscalate escalates a simulated payment dispute.
// SimulatePaymentDisputeEscalate 升级模拟支付争议。
// 官方文档: https://www.airwallex.com/docs/api/simulation/payments/escalate_payment_disputes.md
func (s *Service) SimulatePaymentDisputeEscalate(ctx context.Context, id string, req *SimulatePaymentDisputeEscalateRequest, opts ...sdk.RequestOption) (*PaymentDispute, error) {
	var resp PaymentDispute
	err := s.doer.Do(ctx, "POST", "/api/v1/simulation/pa/payment_disputes/"+id+"/escalate", req, &resp, opts...)
	return &resp, err
}

// SimulatePaymentDisputeResolve resolves a simulated payment dispute.
// SimulatePaymentDisputeResolve 解决模拟支付争议。
// 官方文档: https://www.airwallex.com/docs/api/simulation/payments/resolve_payment_disputes.md
func (s *Service) SimulatePaymentDisputeResolve(ctx context.Context, id string, req *SimulatePaymentDisputeResolveRequest, opts ...sdk.RequestOption) (*PaymentDispute, error) {
	var resp PaymentDispute
	err := s.doer.Do(ctx, "POST", "/api/v1/simulation/pa/payment_disputes/"+id+"/resolve", req, &resp, opts...)
	return &resp, err
}

// SimulateShopperAction simulates a shopper action.
// SimulateShopperAction 模拟购物者操作。
// 官方文档: https://www.airwallex.com/docs/api/simulation/payments/shopper_actions.md
func (s *Service) SimulateShopperAction(ctx context.Context, action string, req *SimulateShopperActionRequest, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/simulation/pa/shopper_actions/"+action, req, nil, opts...)
}
