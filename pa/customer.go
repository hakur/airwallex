package pa

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// Customer represents a payment customer.
// Customer 表示支付客户信息。
type Customer struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// MerchantCustomerID is the merchant customer unique identifier. Optional.
	// MerchantCustomerID 商户客户唯一标识符。可选。
	MerchantCustomerID string `json:"merchant_customer_id,omitempty"`
	// Email is the email address. Optional.
	// Email 邮箱地址。可选。
	Email string `json:"email,omitempty"`
	// FirstName is the first name. Optional.
	// FirstName 名字。可选。
	FirstName string `json:"first_name,omitempty"`
	// LastName is the last name. Optional.
	// LastName 姓氏。可选。
	LastName string `json:"last_name,omitempty"`
	// PhoneNumber is the phone number. Optional.
	// PhoneNumber 电话号码。可选。
	PhoneNumber string `json:"phone_number,omitempty"`
	// BusinessName is the business name. Optional.
	// BusinessName 企业名称。可选。
	BusinessName string `json:"business_name,omitempty"`
	// Address is the address information. Optional.
	// Address 地址信息。可选。
	Address map[string]any `json:"address,omitempty"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
	// CreatedAt is the creation time. Optional.
	// CreatedAt 创建时间。可选。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt is the update time. Optional.
	// UpdatedAt 更新时间。可选。
	UpdatedAt string `json:"updated_at,omitempty"`
}

// CreateCustomerRequest is the request to create a customer.
// CreateCustomerRequest 创建客户请求。
type CreateCustomerRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// MerchantCustomerID is the merchant customer unique identifier. Optional.
	// MerchantCustomerID 商户客户唯一标识符。可选。
	MerchantCustomerID string `json:"merchant_customer_id,omitempty"`
	// Email is the email address. Required.
	// Email 邮箱地址。必填。
	Email string `json:"email"`
	// FirstName is the first name. Optional.
	// FirstName 名字。可选。
	FirstName string `json:"first_name,omitempty"`
	// LastName is the last name. Optional.
	// LastName 姓氏。可选。
	LastName string `json:"last_name,omitempty"`
	// PhoneNumber is the phone number. Optional.
	// PhoneNumber 电话号码。可选。
	PhoneNumber string `json:"phone_number,omitempty"`
	// BusinessName is the business name. Optional.
	// BusinessName 企业名称。可选。
	BusinessName string `json:"business_name,omitempty"`
	// Address is the address information. Optional.
	// Address 地址信息。可选。
	Address map[string]any `json:"address,omitempty"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// UpdateCustomerRequest is the request to update a customer.
// UpdateCustomerRequest 更新客户请求。
type UpdateCustomerRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Email is the email address. Optional.
	// Email 邮箱地址。可选。
	Email string `json:"email,omitempty"`
	// FirstName is the first name. Optional.
	// FirstName 名字。可选。
	FirstName string `json:"first_name,omitempty"`
	// LastName is the last name. Optional.
	// LastName 姓氏。可选。
	LastName string `json:"last_name,omitempty"`
	// PhoneNumber is the phone number. Optional.
	// PhoneNumber 电话号码。可选。
	PhoneNumber string `json:"phone_number,omitempty"`
	// BusinessName is the business name. Optional.
	// BusinessName 企业名称。可选。
	BusinessName string `json:"business_name,omitempty"`
	// Address is the address information. Optional.
	// Address 地址信息。可选。
	Address map[string]any `json:"address,omitempty"`
	// Metadata is additional metadata. Optional.
	// Metadata 元数据。可选。
	Metadata map[string]any `json:"metadata,omitempty"`
}

// CreateCustomer creates a customer.
// CreateCustomer 创建客户。
// 官方文档: https://www.airwallex.com/docs/api/payments/customers/create.md
func (s *Service) CreateCustomer(ctx context.Context, req *CreateCustomerRequest, opts ...sdk.RequestOption) (*Customer, error) {
	var resp Customer
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/customers/create", req, &resp, opts...)
	return &resp, err
}

// GetCustomer retrieves a customer by ID.
// GetCustomer 根据 ID 获取客户。
// 官方文档: https://www.airwallex.com/docs/api/payments/customers/retrieve.md
func (s *Service) GetCustomer(ctx context.Context, id string, opts ...sdk.RequestOption) (*Customer, error) {
	var resp Customer
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/customers/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateCustomer updates a customer.
// UpdateCustomer 更新客户。
// 官方文档: https://www.airwallex.com/docs/api/payments/customers/update.md
func (s *Service) UpdateCustomer(ctx context.Context, id string, req *UpdateCustomerRequest, opts ...sdk.RequestOption) (*Customer, error) {
	var resp Customer
	err := s.doer.Do(ctx, "POST", "/api/v1/pa/customers/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ListCustomers lists customers.
// ListCustomers 列出客户。
// 官方文档: https://www.airwallex.com/docs/api/payments/customers/list.md
func (s *Service) ListCustomers(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[Customer], error) {
	var resp sdk.ListResult[Customer]
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/customers", nil, &resp, opts...)
	return &resp, err
}
