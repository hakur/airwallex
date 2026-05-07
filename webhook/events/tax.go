// Package events provides typed webhook event structures for the tax domain.
// Tax 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/tax.md
//
// 事件映射表:
//
//	tax.tax_form.tax_id_verification_submitted → TaxFormTaxIDVerificationSubmittedEvent (Data: TaxForm)
//	tax.tax_form.tax_id_verification_failed → TaxFormTaxIDVerificationFailedEvent (Data: TaxForm)
//	tax.tax_form.ready                      → TaxFormReadyEvent                (Data: TaxForm)
//	tax.tax_form.expired                    → TaxFormExpiredEvent              (Data: TaxForm)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Tax Form Events ---

// TaxFormTaxIDVerificationSubmittedEvent represents the tax.tax_form.tax_id_verification_submitted webhook event.
// TaxFormTaxIDVerificationSubmittedEvent 表示税表已提交并等待验证的事件。
type TaxFormTaxIDVerificationSubmittedEvent struct {
	Event
	Data TaxForm `json:"data"`
}

// TaxFormTaxIDVerificationFailedEvent represents the tax.tax_form.tax_id_verification_failed webhook event.
// TaxFormTaxIDVerificationFailedEvent 表示税表 tax_id 验证失败的事件。
type TaxFormTaxIDVerificationFailedEvent struct {
	Event
	Data TaxForm `json:"data"`
}

// TaxFormReadyEvent represents the tax.tax_form.ready webhook event.
// TaxFormReadyEvent 表示税表 tax_id 验证成功并已激活或准备提交的事件。
type TaxFormReadyEvent struct {
	Event
	Data TaxForm `json:"data"`
}

// TaxFormExpiredEvent represents the tax.tax_form.expired webhook event.
// TaxFormExpiredEvent 表示税表已过期的事件。
type TaxFormExpiredEvent struct {
	Event
	Data TaxForm `json:"data"`
}

// TaxForm represents a tax form object in webhook payload.
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/tax.md
//
// TaxForm 表示 webhook payload 中的税表对象。
type TaxForm struct {
	ID                         string `json:"id"`
	TaxpayerExternalIdentifier string `json:"taxpayer_external_identifier"`
	Type                       string `json:"type"`
}
