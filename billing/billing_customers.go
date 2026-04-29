package billing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// BillingCustomer represents a billing customer response.
// BillingCustomer 账单客户响应。
type BillingCustomer struct {
	// Address is the billing customer address. Conditional.
	// Address 地址。条件字段。
	Address *Address `json:"address,omitempty"`
	// CreatedAt is the billing customer creation time. Required.
	// CreatedAt 账单客户创建时间。必填。
	CreatedAt string `json:"created_at"`
	// DefaultBillingCurrency 账单客户的默认计费货币，3 字母 ISO-4217 格式。条件字段。
	DefaultBillingCurrency string `json:"default_billing_currency,omitempty"`
	// DefaultLegalEntityID 账单客户的默认法律实体 ID。条件字段。
	DefaultLegalEntityID string `json:"default_legal_entity_id,omitempty"`
	// Description 账单客户的附加描述。条件字段。
	Description string `json:"description,omitempty"`
	// Email is the billing customer email address. Conditional.
	// Email 账单客户的电子邮箱。条件字段。
	Email string `json:"email,omitempty"`
	// ID is the billing customer object ID. Required.
	// ID 账单客户对象的 ID。必填。
	ID string `json:"id"`
	// Metadata 附加到此对象的字符串键值对集合，用于存储额外信息。条件字段。
	Metadata map[string]string `json:"metadata,omitempty"`
	// Name 账单客户的名称。条件字段。
	Name string `json:"name,omitempty"`
	// Nickname 账单客户的内部昵称。条件字段。
	Nickname string `json:"nickname,omitempty"`
	// PhoneNumber 账单客户的电话号码。条件字段。
	PhoneNumber string `json:"phone_number,omitempty"`
	// TaxIdentificationNumber 账单客户的税务识别号。条件字段。
	TaxIdentificationNumber string `json:"tax_identification_number,omitempty"`
	// Type is the billing customer type, either BUSINESS or INDIVIDUAL. Conditional.
	// Type 账单客户的类型，BUSINESS 或 INDIVIDUAL 之一。条件字段。
	Type CustomerType `json:"type,omitempty"`
}

// CreateBillingCustomerRequest represents a request to create a billing customer.
// CreateBillingCustomerRequest 创建账单客户请求。
type CreateBillingCustomerRequest struct {
	// Address is the billing customer address. Conditional.
	// Address 地址。条件字段。
	Address *Address `json:"address,omitempty"`
	// DefaultBillingCurrency is the default billing currency (3-letter ISO-4217). Conditional.
	// DefaultBillingCurrency 账单客户的默认计费货币，3 字母 ISO-4217 格式。条件字段。
	DefaultBillingCurrency string `json:"default_billing_currency,omitempty"`
	// DefaultLegalEntityID is the default legal entity ID. Conditional.
	// DefaultLegalEntityID 账单客户的默认法律实体 ID。条件字段。
	DefaultLegalEntityID string `json:"default_legal_entity_id,omitempty"`
	// Description is an additional description for the billing customer. Conditional.
	// Description 账单客户的附加描述。条件字段。
	Description string `json:"description,omitempty"`
	// Email is the billing customer email address. Conditional.
	// Email 账单客户的电子邮箱。条件字段。
	Email string `json:"email,omitempty"`
	// Metadata is a set of key-value pairs for storing additional information. Conditional.
	// Metadata 附加到此对象的字符串键值对集合，用于存储额外信息。条件字段。
	Metadata map[string]string `json:"metadata,omitempty"`
	// Name is the billing customer name. Conditional.
	// Name 账单客户的名称。条件字段。
	Name string `json:"name,omitempty"`
	// Nickname is the internal nickname for the billing customer. Conditional.
	// Nickname 账单客户的内部昵称。条件字段。
	Nickname string `json:"nickname,omitempty"`
	// PhoneNumber 账单客户的电话号码。条件字段。
	PhoneNumber string `json:"phone_number,omitempty"`
	// RequestID is the merchant-specified unique request ID. Required.
	// RequestID 商户指定的唯一请求 ID。必填.
	RequestID string `json:"request_id"`
	// TaxIdentificationNumber is the billing customer tax identification number. Conditional.
	// TaxIdentificationNumber 账单客户的税务识别号。条件字段。
	TaxIdentificationNumber string `json:"tax_identification_number,omitempty"`
	// Type is the billing customer type, either BUSINESS or INDIVIDUAL. Conditional.
	// Type 账单客户的类型，BUSINESS 或 INDIVIDUAL 之一。条件字段。
	Type CustomerType `json:"type,omitempty"`
}

// UpdateBillingCustomerRequest represents a request to update a billing customer.
// UpdateBillingCustomerRequest 更新账单客户请求。
type UpdateBillingCustomerRequest struct {
	// Address is the billing customer address. Conditional.
	// Address 地址。条件字段。
	Address *Address `json:"address,omitempty"`
	// DefaultBillingCurrency is the default billing currency (3-letter ISO-4217). Conditional.
	// DefaultBillingCurrency 账单客户的默认计费货币，3 字母 ISO-4217 格式。条件字段。
	DefaultBillingCurrency string `json:"default_billing_currency,omitempty"`
	// DefaultLegalEntityID is the default legal entity ID. Conditional.
	// DefaultLegalEntityID 账单客户的默认法律实体 ID。条件字段。
	DefaultLegalEntityID string `json:"default_legal_entity_id,omitempty"`
	// Description is an additional description for the billing customer. Conditional.
	// Description 账单客户的附加描述。条件字段。
	Description string `json:"description,omitempty"`
	// Email is the billing customer email address. Conditional.
	// Email 账单客户的电子邮箱。条件字段。
	Email string `json:"email,omitempty"`
	// Metadata is a set of key-value pairs for storing additional information. Conditional.
	// Metadata 附加到此对象的字符串键值对集合，用于存储额外信息。条件字段。
	Metadata map[string]string `json:"metadata,omitempty"`
	// Name is the billing customer name. Conditional.
	// Name 账单客户的名称。条件字段。
	Name string `json:"name,omitempty"`
	// Nickname is the internal nickname for the billing customer. Conditional.
	// Nickname 账单客户的内部昵称。条件字段。
	Nickname string `json:"nickname,omitempty"`
	// PhoneNumber is the billing customer phone number. Conditional.
	// PhoneNumber 账单客户的电话号码。条件字段。
	PhoneNumber string `json:"phone_number,omitempty"`
	// TaxIdentificationNumber is the billing customer tax identification number. Conditional.
	// TaxIdentificationNumber 账单客户的税务识别号。条件字段。
	TaxIdentificationNumber string `json:"tax_identification_number,omitempty"`
	// Type is the billing customer type, either BUSINESS or INDIVIDUAL. Conditional.
	// Type 账单客户的类型，BUSINESS 或 INDIVIDUAL 之一。条件字段。
	Type CustomerType `json:"type,omitempty"`
}

// ListBillingCustomersRequest represents a request to list billing customers.
// ListBillingCustomersRequest 获取账单客户列表请求。
type ListBillingCustomersRequest struct {
	// FromCreatedAt is the start of the creation time range (ISO8601 inclusive). Optional.
	// FromCreatedAt created_at 的起始时间，ISO8601 格式（包含）。可选。
	FromCreatedAt string `json:"from_created_at,omitempty"`
	// ToCreatedAt is the end of the creation time range (ISO8601 exclusive). Optional.
	// ToCreatedAt created_at 的结束时间，ISO8601 格式（不包含）。可选。
	ToCreatedAt string `json:"to_created_at,omitempty"`
	// Email is the billing customer email address. Optional.
	// Email 账单客户的电子邮箱。可选。
	Email string `json:"email,omitempty"`
	// Page is the pagination cursor. Optional.
	// Page 分页游标。可选。
	Page string `json:"page,omitempty"`
	// PageSize is the number of items per page, defaults to 20. Optional.
	// PageSize 每页账单客户对象数量，默认为 20。可选。
	PageSize int32 `json:"page_size,omitempty"`
}

// CreateBillingCustomer creates a new billing customer.
// CreateBillingCustomer 创建账单客户。
// 官方文档: https://www.airwallex.com/docs/api/billing/billing_customers/create.md
func (s *Service) CreateBillingCustomer(ctx context.Context, req *CreateBillingCustomerRequest, opts ...sdk.RequestOption) (*BillingCustomer, error) {
	var resp BillingCustomer
	err := s.doer.Do(ctx, "POST", "/api/v1/billing_customers/create", req, &resp, opts...)
	return &resp, err
}

// GetBillingCustomer retrieves a billing customer by ID.
// GetBillingCustomer 根据 ID 检索账单客户。
// 官方文档: https://www.airwallex.com/docs/api/billing/billing_customers/retrieve.md
func (s *Service) GetBillingCustomer(ctx context.Context, id string, opts ...sdk.RequestOption) (*BillingCustomer, error) {
	var resp BillingCustomer
	err := s.doer.Do(ctx, "GET", "/api/v1/billing_customers/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateBillingCustomer updates a billing customer. Only fields provided in the request are updated.
// UpdateBillingCustomer 更新账单客户。仅更新请求中提供的字段，省略的字段保持不变。
// 官方文档: https://www.airwallex.com/docs/api/billing/billing_customers/update.md
func (s *Service) UpdateBillingCustomer(ctx context.Context, id string, req *UpdateBillingCustomerRequest, opts ...sdk.RequestOption) (*BillingCustomer, error) {
	var resp BillingCustomer
	err := s.doer.Do(ctx, "POST", "/api/v1/billing_customers/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ListBillingCustomers lists billing customers with optional filters.
// ListBillingCustomers 列出账单客户。
// 官方文档: https://www.airwallex.com/docs/api/billing/billing_customers/list.md
func (s *Service) ListBillingCustomers(ctx context.Context, req *ListBillingCustomersRequest, opts ...sdk.RequestOption) (*ListResult[BillingCustomer], error) {
	var resp ListResult[BillingCustomer]
	err := s.doer.Do(ctx, "GET", "/api/v1/billing_customers", req, &resp, opts...)
	return &resp, err
}
