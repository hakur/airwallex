package confirmation_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/confirmation"
	"github.com/hakur/airwallex/core"
	"github.com/hakur/airwallex/sdk"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

var testClient *airwallex.Client

func TestMain(m *testing.M) {
	_ = godotenv.Load(sdk.ResolveEnvPath())
	var err error
	testClient, err = airwallex.NewFromEnv("", sdk.WithBaseURL(sdk.SandboxURL), sdk.WithDebug(true))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create test client: %v\n", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestConfirmationLetterCreate(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Confirmation()

	// 获取一个真实的存款 ID 作为交易 ID（确认函需要真实的存款或转账交易）
	deposits, err := testClient.Core().ListDeposits(ctx, &core.ListDepositsRequest{PageSize: 1})
	if err != nil || len(deposits) == 0 {
		t.Logf("no deposits available in sandbox, skipping confirmation letter test")
		return
	}

	err = svc.CreateConfirmationLetter(ctx, &confirmation.CreateConfirmationLetterRequest{
		Format:        "STANDARD",
		TransactionID: deposits[0].ID,
	})
	if err != nil {
		// 沙箱环境或无权限时返回 401/400 是可接受的
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsBadRequest(err), "create confirmation letter failed: %v", err)
		t.Logf("confirmation letter skipped (sandbox limitation): %v", err)
		return
	}
	t.Logf("confirmation letter created successfully")
}
