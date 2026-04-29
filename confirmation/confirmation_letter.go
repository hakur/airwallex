package confirmation

import (
	"context"

	"github.com/hakur/airwallex/sdk"
)

// CreateConfirmationLetterRequest represents a request to create a confirmation letter.
// CreateConfirmationLetterRequest 创建确认函请求。
type CreateConfirmationLetterRequest struct {
	// Format is the confirmation letter format. Required.
	// STANDARD — includes fee information.
	// NO_FEE_DISPLAY — hides fee information.
	// Format 确认函格式。必填。
	Format string `json:"format"`
	// TransactionID is the deposit or transfer transaction ID. Required.
	// TransactionID 存款或转账的 ID。必填。
	TransactionID string `json:"transaction_id"`
}

// CreateConfirmationLetter creates a confirmation letter (returns a PDF file stream).
// CreateConfirmationLetter 创建确认函（返回 PDF 文件流）。
// 官方文档: https://www.airwallex.com/docs/api/confirmation_letter/confirmation_letter/create.md
func (s *Service) CreateConfirmationLetter(ctx context.Context, req *CreateConfirmationLetterRequest, opts ...sdk.RequestOption) error {
	return s.doer.Do(ctx, "POST", "/api/v1/confirmation_letters/create", req, nil, opts...)
}

// --- Deprecated ---
// GetConfirmationLetter / UpdateConfirmationLetter / ListConfirmationLetters 已移除。
// 官方 API 仅提供 POST create 端点（返回 PDF 文件流，无 JSON 响应体）。
// ConfirmationLetter 响应结构体也已移除。
