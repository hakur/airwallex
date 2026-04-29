package pa

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// CardBrand represents a card brand type.
// CardBrand 表示卡品牌类型。
type CardBrand = string

const (
	// CardBrandVisa represents Visa cards.
	// CardBrandVisa Visa 卡。
	CardBrandVisa CardBrand = "visa"
	// CardBrandMastercard represents Mastercard cards.
	// CardBrandMastercard Mastercard 卡。
	CardBrandMastercard CardBrand = "mastercard"
	// CardBrandUnionPay represents UnionPay cards.
	// CardBrandUnionPay 银联卡。
	CardBrandUnionPay CardBrand = "union pay"
	// CardBrandAmericanExpress represents American Express cards.
	// CardBrandAmericanExpress 美国运通卡。
	CardBrandAmericanExpress CardBrand = "american express"
)

// CardType represents a card type within a brand.
// CardType 表示卡类型。
type CardType = string

const (
	// CardTypeCredit represents a credit card (Visa/Mastercard/UnionPay).
	// CardTypeCredit 信用卡。
	CardTypeCredit CardType = "credit"
	// CardTypeDebit represents a debit card (Visa/Mastercard/UnionPay).
	// CardTypeDebit 借记卡。
	CardTypeDebit CardType = "debit"
	// CardTypePrepaid represents a prepaid card (Visa/Mastercard/UnionPay).
	// CardTypePrepaid 预付卡。
	CardTypePrepaid CardType = "prepaid"
	// CardTypeCharge represents a charge card (Visa only).
	// CardTypeCharge 签账卡。
	CardTypeCharge CardType = "charge"
	// CardTypeDeferredDebit represents a deferred debit card (Visa only).
	// CardTypeDeferredDebit 延期借记卡。
	CardTypeDeferredDebit CardType = "deferred_debit"
	// CardTypeNotApplicable represents not applicable (Visa only).
	// CardTypeNotApplicable 不适用。
	CardTypeNotApplicable CardType = "not_applicable"
)

// IssuerAccountRangeForMerchantResponse represents BIN lookup result.
// IssuerAccountRangeForMerchantResponse 表示 BIN 查询结果。
type IssuerAccountRangeForMerchantResponse struct {
	// CardBrand is the card brand (visa, mastercard, union pay, american express). Required.
	// CardBrand 卡品牌。必填。
	CardBrand CardBrand `json:"card_brand"`
	// CardType is the card type within the brand. Required.
	// CardType 卡类型。必填。
	CardType CardType `json:"card_type"`
	// CommercialCard indicates whether the card is commercial. Required.
	// CommercialCard 是否为商务卡。必填。
	CommercialCard bool `json:"commercial_card"`
	// IssuerCountryCode is the issuer country code. Conditional.
	// IssuerCountryCode 发卡行国家代码。条件字段。
	IssuerCountryCode string `json:"issuer_country_code,omitempty"`
	// IssuerName is the issuer name. Conditional.
	// IssuerName 发卡行名称。条件字段。
	IssuerName string `json:"issuer_name,omitempty"`
	// ProductCode is the card product code. Conditional.
	// ProductCode 卡产品代码。条件字段。
	ProductCode string `json:"product_code,omitempty"`
	// ProductDescription is the description for product_code. Conditional.
	// ProductDescription 产品代码描述。条件字段。
	ProductDescription string `json:"product_description,omitempty"`
	// ProductSubtypeCode is the product subtype code. Conditional.
	// ProductSubtypeCode 产品子类型代码。条件字段。
	ProductSubtypeCode string `json:"product_subtype_code,omitempty"`
	// ProductSubtypeDescription is the description for product_subtype_code. Conditional.
	// ProductSubtypeDescription 产品子类型描述。条件字段。
	ProductSubtypeDescription string `json:"product_subtype_description,omitempty"`
}

// LookupBin retrieves BIN (Bank Identification Number) information using the PAN
// passed via the x-pan header parameter.
// LookupBin 使用 x-pan header 参数传入的 PAN 查询 BIN（银行识别号）信息。
//
// 官方文档: https://www.airwallex.com/docs/api/payments/reference_data/lookup.md
func (s *Service) LookupBin(ctx context.Context, pan string, opts ...sdk.RequestOption) ([]IssuerAccountRangeForMerchantResponse, error) {
	var resp []IssuerAccountRangeForMerchantResponse
	allOpts := []sdk.RequestOption{sdk.WithRequestHeader("x-pan", pan)}
	allOpts = append(allOpts, opts...)
	err := s.doer.Do(ctx, "GET", "/api/v1/pa/reference/bin/lookup", nil, &resp, allOpts...)
	return resp, err
}
