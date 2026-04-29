package spend

// AccountingFieldSelection represents an ERP accounting field selection.
// AccountingFieldSelection ERP 会计字段选择。
type AccountingFieldSelection struct {
	Type       string `json:"type"`
	Name       string `json:"name,omitempty"`
	ExternalID string `json:"external_id,omitempty"`
	Value      string `json:"value"`
	ValueLabel string `json:"value_label,omitempty"`
}

// Comment represents a comment.
// Comment 评论。
type Comment struct {
	Content   string `json:"content"`
	CreatedBy string `json:"created_by,omitempty"`
	CreatedAt string `json:"created_at"`
}

// Address represents an address.
// Address 地址。
type Address struct {
	StreetAddress string `json:"street_address,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omitempty"`
	Postcode      string `json:"postcode,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
}
