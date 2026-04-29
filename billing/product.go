package billing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// Product represents a product response.
// Product 产品响应。
type Product struct {
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Active 是否活跃。必填。
	Active bool `json:"active"`
	// CreatedAt 创建时间。必填。
	CreatedAt string `json:"created_at"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// Name 名称。必填。
	Name string `json:"name"`
	// Unit 单位。必填。
	Unit string `json:"unit"`
	// UpdatedAt 更新时间。必填。
	UpdatedAt string `json:"updated_at"`
}

// CreateProductRequest represents a request to create a product.
// CreateProductRequest 创建产品请求。
type CreateProductRequest struct {
	// RequestID is the unique request identifier. Required.
	// RequestID 请求唯一标识符。必填。
	RequestID string `json:"request_id"`
	// Active 是否活跃。可选。
	Active bool `json:"active,omitempty"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// Name 名称。必填。
	Name string `json:"name"`
	// Unit 单位。可选。
	Unit string `json:"unit,omitempty"`
}

// UpdateProductRequest represents a request to update a product.
// UpdateProductRequest 更新产品请求。
type UpdateProductRequest struct {
	// Active 是否活跃。可选。
	Active bool `json:"active,omitempty"`
	// Description 描述。可选。
	Description string `json:"description,omitempty"`
	// Metadata 元数据。可选。
	Metadata map[string]string `json:"metadata,omitempty"`
	// Name 名称。可选。
	Name string `json:"name,omitempty"`
	// Unit 单位。可选。
	Unit string `json:"unit,omitempty"`
}

// ListProductsRequest represents a request to list products.
// ListProductsRequest 获取产品列表请求。
type ListProductsRequest struct {
	// Page 分页游标。可选。
	Page string `json:"page,omitempty"`
	// PageSize 每页数量。可选。
	PageSize int32 `json:"page_size,omitempty"`
}

// CreateProduct creates a new product.
// CreateProduct 创建产品。
// 官方文档: https://www.airwallex.com/docs/api/billing/products/create.md
func (s *Service) CreateProduct(ctx context.Context, req *CreateProductRequest, opts ...sdk.RequestOption) (*Product, error) {
	var resp Product
	err := s.doer.Do(ctx, "POST", "/api/v1/products/create", req, &resp, opts...)
	return &resp, err
}

// GetProduct retrieves a product by ID.
// GetProduct 根据 ID 获取产品。
// 官方文档: https://www.airwallex.com/docs/api/billing/products/retrieve.md
func (s *Service) GetProduct(ctx context.Context, id string, opts ...sdk.RequestOption) (*Product, error) {
	var resp Product
	err := s.doer.Do(ctx, "GET", "/api/v1/products/"+id, nil, &resp, opts...)
	return &resp, err
}

// UpdateProduct updates an existing product.
// UpdateProduct 更新产品。
// 官方文档: https://www.airwallex.com/docs/api/billing/products/update.md
func (s *Service) UpdateProduct(ctx context.Context, id string, req *UpdateProductRequest, opts ...sdk.RequestOption) (*Product, error) {
	var resp Product
	err := s.doer.Do(ctx, "POST", "/api/v1/products/"+id+"/update", req, &resp, opts...)
	return &resp, err
}

// ListProducts lists products with optional filters.
// ListProducts 列出产品。
// 官方文档: https://www.airwallex.com/docs/api/billing/products/list.md
func (s *Service) ListProducts(ctx context.Context, req *ListProductsRequest, opts ...sdk.RequestOption) (*ListResult[Product], error) {
	var resp ListResult[Product]
	err := s.doer.Do(ctx, "GET", "/api/v1/products", req, &resp, opts...)
	return &resp, err
}
