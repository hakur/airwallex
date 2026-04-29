package sdk

import (
	"encoding/json"
	"fmt"
)

// APIError 表示 Airwallex API 返回的业务错误。
type APIError struct {
	Code    string          `json:"code"`
	Source  string          `json:"source,omitempty"`
	Message string          `json:"message"`
	Details json.RawMessage `json:"details,omitempty"`
}

// Error 实现 error 接口。
func (e *APIError) Error() string {
	if e.Source != "" {
		return fmt.Sprintf("airwallex api error [%s] source=%s: %s", e.Code, e.Source, e.Message)
	}
	return fmt.Sprintf("airwallex api error [%s]: %s", e.Code, e.Message)
}

// ErrorCode 定义 Airwallex 业务错误码常量。
type ErrorCode = string

const (
	ErrorCodeAlreadyExists                  ErrorCode = "already_exists"
	ErrorCodeAmountAboveLimit               ErrorCode = "amount_above_limit"
	ErrorCodeAmountAboveTransferMethodLimit ErrorCode = "amount_above_transfer_method_limit"
	ErrorCodeAmountBelowLimit               ErrorCode = "amount_below_limit"
	ErrorCodeBadRequest                     ErrorCode = "bad_request"
	ErrorCodeCannotBeEdited                 ErrorCode = "can_not_be_edited"
	ErrorCodeConcurrentUpdate               ErrorCode = "concurrent_update"
	ErrorCodeConversionCreateFailed         ErrorCode = "conversion_create_failed"
	ErrorCodeCredentialsExpired             ErrorCode = "credentials_expired"
	ErrorCodeCredentialsInvalid             ErrorCode = "credentials_invalid"
	ErrorCodeFieldRequired                  ErrorCode = "field_required"
	ErrorCodeInvalidArgument                ErrorCode = "invalid_argument"
	ErrorCodeInvalidCurrencyPair            ErrorCode = "invalid_currency_pair"
	ErrorCodeInvalidTransferDate            ErrorCode = "invalid_transfer_date"
	ErrorCodeInvalidConversionDate          ErrorCode = "invalid_conversion_date"
	ErrorCodeNotFound                       ErrorCode = "not_found"
	ErrorCodeServiceUnavailable             ErrorCode = "service_unavailable"
	ErrorCodeTermAgreementRequired          ErrorCode = "term_agreement_is_required"
	ErrorCodeTooManyRequests                ErrorCode = "too_many_requests"
	ErrorCodeUnauthorized                   ErrorCode = "unauthorized"
	ErrorCodeUnsupportedCountryCode         ErrorCode = "unsupported_country_code"
	ErrorCodeUnsupportedCurrency            ErrorCode = "unsupported_currency"
	ErrorCodeUnsupportedTransferMethod      ErrorCode = "unsupported_transfer_method"
	ErrorCodeProviderDeclined               ErrorCode = "provider_declined"
	ErrorCodeValidationError                ErrorCode = "validation_error"
	ErrorCodeDuplicateRequest               ErrorCode = "duplicate_request"
	ErrorCodeInvalidStatusForOperation      ErrorCode = "invalid_status_for_operation"
	ErrorCodeForbidden                      ErrorCode = "forbidden"
	ErrorCodeResourceNotFound               ErrorCode = "resource_not_found"
	ErrorCodeMethodNotAllowed               ErrorCode = "method_not_allowed"
	ErrorCodeOperationFailed                ErrorCode = "operation_failed"
	ErrorCodeValidationFailed               ErrorCode = "validation_failed"
	ErrorCodeInternalError                  ErrorCode = "internal_error"
	ErrorCodeConfirmFundingUnsupported      ErrorCode = "confirm_funding_unsupported"
)

// IsBadRequest 判断错误是否为 bad_request（400）。
func IsBadRequest(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeBadRequest)
	}
	return false
}

// IsInvalidArgument 判断错误是否为请求参数无效（400）。
func IsInvalidArgument(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeInvalidArgument)
	}
	return false
}

// IsProviderDeclined 判断错误是否为支付提供方拒绝。
func IsProviderDeclined(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeProviderDeclined)
	}
	return false
}

// IsValidationError 判断错误是否为验证错误（如激活码无效）。
func IsValidationError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeValidationError)
	}
	return false
}

// IsDuplicateRequest 判断错误是否为重复请求。
func IsDuplicateRequest(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeDuplicateRequest)
	}
	return false
}

// IsInvalidStatusForOperation 判断错误是否为操作状态无效。
func IsInvalidStatusForOperation(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeInvalidStatusForOperation)
	}
	return false
}

// IsMethodNotAllowed 判断错误是否为方法不允许（405）。
func IsMethodNotAllowed(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeMethodNotAllowed)
	}
	return false
}

// IsOperationFailed 判断错误是否为操作失败（500）。
func IsOperationFailed(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeOperationFailed)
	}
	return false
}

// IsForbidden 判断错误是否为禁止访问（403）。
func IsForbidden(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeForbidden)
	}
	return false
}

// IsResourceNotFound 判断错误是否为资源不存在。
func IsResourceNotFound(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeResourceNotFound)
	}
	return false
}

// IsConfirmFundingUnsupported 判断错误是否为不支持确认资金来源。
func IsConfirmFundingUnsupported(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeConfirmFundingUnsupported)
	}
	return false
}

// IsValidationFailed 判断错误是否为模式验证失败。
func IsValidationFailed(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeValidationFailed)
	}
	return false
}

// IsCredentialsInvalid 判断错误是否为凭证无效。
func IsCredentialsInvalid(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeCredentialsInvalid)
	}
	return false
}

// IsInternalError 判断错误是否为内部错误（500）。
func IsInternalError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == string(ErrorCodeInternalError)
	}
	return false
}
