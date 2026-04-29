package issuing

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// MerchantBrand represents merchant brand information.
// MerchantBrand 表示商户品牌信息。
type MerchantBrand struct {
	// ID is the unique identifier. Required.
	// ID 唯一标识符。必填。
	ID string `json:"id"`
	// Name is the name. Required.
	// Name 名称。必填。
	Name string `json:"name"`
	// LogoURL is the brand logo URL. Optional.
	// LogoURL 品牌 Logo URL。可选。
	LogoURL string `json:"logo_url,omitempty"`
}

// ListMerchantBrands lists merchant brands.
// 官方文档: https://www.airwallex.com/docs/api/issuing/merchant_brands/list.md
// ListMerchantBrands 列出商户品牌。
func (s *Service) ListMerchantBrands(ctx context.Context, opts ...sdk.RequestOption) (*sdk.ListResult[MerchantBrand], error) {
	var resp sdk.ListResult[MerchantBrand]
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/merchant_brands", nil, &resp, opts...)
	return &resp, err
}

// GetMerchantBrand retrieves a merchant brand by ID.
// 官方文档: https://www.airwallex.com/docs/api/issuing/merchant_brands/retrieve.md
// GetMerchantBrand 根据 ID 获取商户品牌。
func (s *Service) GetMerchantBrand(ctx context.Context, id string, opts ...sdk.RequestOption) (*MerchantBrand, error) {
	var resp MerchantBrand
	err := s.doer.Do(ctx, "GET", "/api/v1/issuing/merchant_brands/"+id, nil, &resp, opts...)
	return &resp, err
}
