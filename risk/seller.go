package risk

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// SellerAddress represents a seller's address.
// SellerAddress 表示卖家地址。
type SellerAddress struct {
	// AddressLine1 is the first line of the address.
	// AddressLine1 地址第一行。
	AddressLine1 string `json:"address_line1,omitempty"`
	// AddressLine2 is the second line of the address.
	// AddressLine2 地址第二行。
	AddressLine2 string `json:"address_line2,omitempty"`
	// CountryCode is the 2-letter ISO 3166-2 country code.
	// CountryCode 国家代码（2位ISO 3166-2代码）。
	CountryCode string `json:"country_code,omitempty"`
	// Postcode is the postal code.
	// Postcode 邮政编码。
	Postcode string `json:"postcode,omitempty"`
	// State is the state or province.
	// State 州或省。
	State string `json:"state,omitempty"`
	// Suburb is the suburb or city.
	// Suburb 郊区或城市。
	Suburb string `json:"suburb,omitempty"`
}

// SellerWebsite represents a seller's website.
// SellerWebsite 表示卖家网站。
type SellerWebsite struct {
	// URL is the website URL.
	// URL 网站URL。
	URL string `json:"url,omitempty"`
}

// SellerDetails represents detailed information about a seller.
// SellerDetails represents detailed information about a seller.
// SellerDetails 表示卖家详细信息。
type SellerDetails struct {
	// Address is the seller's address.
	// Address 卖家地址。
	Address *SellerAddress `json:"address,omitempty"`
	// BusinessIdentificationNumber is the business identification number.
	// BusinessIdentificationNumber 企业识别号码。
	BusinessIdentificationNumber string `json:"business_identification_number,omitempty"`
	// Email is the contact email address.
	// Email 联系邮箱。
	Email string `json:"email,omitempty"`
	// IndustryCode is the industry code.
	// IndustryCode 行业代码。
	IndustryCode string `json:"industry_code,omitempty"`
	// LegalEntityName is the legal entity name.
	// LegalEntityName 法律实体名称。
	LegalEntityName string `json:"legal_entity_name,omitempty"`
	// MerchantCategoryCode is the merchant category code.
	// MerchantCategoryCode 商户类别代码。
	MerchantCategoryCode string `json:"merchant_category_code,omitempty"`
	// PhoneNumber is the phone number.
	// PhoneNumber 电话号码。
	PhoneNumber string `json:"phone_number,omitempty"`
	// RegistrationCountry is the registration country code.
	// RegistrationCountry 注册国家代码。
	RegistrationCountry string `json:"registration_country,omitempty"`
	// RegistrationDate is the registration date (YYYY-MM-DD format).
	// RegistrationDate 注册日期（YYYY-MM-DD格式）。
	RegistrationDate string `json:"registration_date,omitempty"`
	// TradingName is the trading name.
	// TradingName 经营名称。
	TradingName string `json:"trading_name,omitempty"`
	// Websites is the list of websites.
	// Websites 网站列表。
	Websites []SellerWebsite `json:"websites,omitempty"`
}

// Seller represents seller information.
// Seller 表示卖家信息。
type Seller struct {
	// ID is the unique seller identifier.
	// ID 卖家唯一标识符。
	ID string `json:"id,omitempty"`
	// Status is the seller status.
	// Status 卖家状态。
	Status string `json:"status,omitempty"`
	// CreatedAt is the creation timestamp (ISO8601 format).
	// CreatedAt 创建时间戳（ISO8601格式）。
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt is the last update timestamp (ISO8601 format).
	// UpdatedAt 最后更新时间戳（ISO8601格式）。
	UpdatedAt string `json:"updated_at,omitempty"`
	// Details is the detailed seller information.
	// Details 卖家详细信息。
	Details *SellerDetails `json:"details,omitempty"`
}

// CreateSellerRequest represents a request to create a seller.
// CreateSellerRequest 创建卖家请求。
type CreateSellerRequest struct {
	// RequestID is the unique request identifier.
	// RequestID 请求唯一标识符。
	RequestID string `json:"request_id,omitempty"`
	// Address is the seller's address.
	// Address 卖家地址。
	Address *SellerAddress `json:"address,omitempty"`
	// BusinessIdentificationNumber is the business identification number.
	// BusinessIdentificationNumber 企业识别号码。
	BusinessIdentificationNumber string `json:"business_identification_number,omitempty"`
	// Email is the contact email address.
	// Email 联系邮箱。
	Email string `json:"email,omitempty"`
	// IndustryCode is the industry code.
	// IndustryCode 行业代码。
	IndustryCode string `json:"industry_code,omitempty"`
	// LegalEntityName is the legal entity name.
	// LegalEntityName 法律实体名称。
	LegalEntityName string `json:"legal_entity_name,omitempty"`
	// MerchantCategoryCode is the merchant category code.
	// MerchantCategoryCode 商户类别代码。
	MerchantCategoryCode string `json:"merchant_category_code,omitempty"`
	// PhoneNumber is the phone number.
	// PhoneNumber 电话号码。
	PhoneNumber string `json:"phone_number,omitempty"`
	// RegistrationCountry is the registration country code.
	// RegistrationCountry 注册国家代码。
	RegistrationCountry string `json:"registration_country,omitempty"`
	// RegistrationDate is the registration date (YYYY-MM-DD format).
	// RegistrationDate 注册日期（YYYY-MM-DD格式）。
	RegistrationDate string `json:"registration_date,omitempty"`
	// TradingName is the trading name.
	// TradingName 经营名称。
	TradingName string `json:"trading_name,omitempty"`
	// Websites is the list of websites.
	// Websites 网站列表。
	Websites []SellerWebsite `json:"websites,omitempty"`
}

// CreateSeller creates a new seller.
// CreateSeller 创建卖家。
// 官方文档: https://www.airwallex.com/docs/api/risk/sellers/create.md
func (s *Service) CreateSeller(ctx context.Context, req *CreateSellerRequest, opts ...sdk.RequestOption) (*Seller, error) {
	var resp Seller
	err := s.doer.Do(ctx, "POST", "/api/v1/sellers/create", req, &resp, opts...)
	return &resp, err
}

// GetSeller retrieves a seller by ID.
// GetSeller 根据ID获取卖家。
// 官方文档: https://www.airwallex.com/docs/api/risk/sellers/retrieve.md
func (s *Service) GetSeller(ctx context.Context, id string, opts ...sdk.RequestOption) (*Seller, error) {
	var resp Seller
	err := s.doer.Do(ctx, "GET", "/api/v1/sellers/"+id, nil, &resp, opts...)
	return &resp, err
}

// ListSellers lists all sellers.
// ListSellers 列出卖家。
// 官方文档: https://www.airwallex.com/docs/api/risk/sellers/list.md
func (s *Service) ListSellers(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[Seller], error) {
	var resp sdk.ListResult[Seller]
	err := s.doer.Do(ctx, "GET", "/api/v1/sellers", nil, &resp, opts...)
	return &resp, err
}

// DeactivateSeller deactivates a seller by ID.
// DeactivateSeller 根据ID停用卖家。
// 官方文档: https://www.airwallex.com/docs/api/risk/sellers/deactivate.md
func (s *Service) DeactivateSeller(ctx context.Context, id string, opts ...sdk.RequestOption) (*Seller, error) {
	var resp Seller
	err := s.doer.Do(ctx, "POST", "/api/v1/sellers/"+id+"/deactivate", nil, &resp, opts...)
	return &resp, err
}
