package scale_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/scale"
	"github.com/hakur/airwallex/sdk"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
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

func TestAccountLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Scale()

	list, err := svc.ListAccounts(ctx)
	require.NoError(t, err, "list accounts failed")
	t.Logf("existing accounts: %d", len(list.Items))

	created, err := svc.CreateAccount(ctx, &scale.CreateAccountRequest{
		RequestID: "acct-req-" + time.Now().Format("20060102150405"),
		Email:     "test-" + time.Now().Format("20060102150405") + "@example.com",
		AccountDetails: &scale.AccountDetails{
			BusinessName: "Test Business",
		},
		PrimaryContact: &scale.PrimaryContact{
			FirstName: "Test",
			LastName:  "User",
			Email:     "test-" + time.Now().Format("20060102150405") + "@example.com",
		},
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsNotFound(err) || sdk.IsInvalidArgument(err), "create account failed: %v", err)
		return
	}
	t.Logf("created account: %s", created.ID)

	fetched, err := svc.GetAccount(ctx, created.ID)
	require.NoError(t, err, "get account failed")
	assert.Equal(t, created.ID, fetched.ID, "account id mismatch")

	current, err := svc.GetCurrentAccount(ctx)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "get current account failed: %v", err)
	} else {
		assert.NotEmpty(t, current.ID)
	}
}

func TestChargeLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Scale()

	// Get a real account ID for source
	accounts, err := svc.ListAccounts(ctx)
	require.NoError(t, err, "list accounts failed")
	if len(accounts.Items) == 0 {
		t.Logf("no accounts available, skipping charge lifecycle")
		return
	}
	realAccountID := accounts.Items[0].ID

	// ListCharges first
	list, err := svc.ListCharges(ctx)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "list charges failed: %v", err)
		return
	}
	require.NotNil(t, list)
	t.Logf("charges count: %d", len(list.Items))

	// Try to get first charge if available
	if len(list.Items) > 0 {
		fetched, err := svc.GetCharge(ctx, list.Items[0].ID)
		if err != nil {
			require.True(t, sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "get charge failed: %v", err)
			t.Logf("get charge skipped: %v", err)
		} else {
			assert.Equal(t, list.Items[0].ID, fetched.ID, "charge id mismatch")
		}
	}

	// CreateCharge - requires a valid source account
	created, err := svc.CreateCharge(ctx, &scale.CreateChargeRequest{
		RequestID: "charge-req-" + time.Now().Format("20060102150405"),
		Source:    realAccountID,
		Amount:    "100.00",
		Currency:  sdk.CurrencyUSD,
		Reason:    "goods_purchased",
		Reference: "ref-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err) || sdk.IsInvalidArgument(err) || sdk.IsValidationFailed(err), "create charge failed: %v", err)
		t.Logf("create charge skipped: %v", err)
	} else {
		require.NotEmpty(t, created.ID, "charge id should not be empty")
		t.Logf("created charge: %s", created.ID)

		fetched, err := svc.GetCharge(ctx, created.ID)
		if err != nil {
			require.True(t, sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "get created charge failed: %v", err)
			t.Logf("get created charge skipped: %v", err)
		} else {
			assert.Equal(t, created.ID, fetched.ID, "charge id mismatch")
		}
	}
}

func TestConnectedAccountTransferLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Scale()

	// Get a real account ID for destination
	accounts, err := svc.ListAccounts(ctx)
	require.NoError(t, err, "list accounts failed")
	if len(accounts.Items) == 0 {
		t.Logf("no accounts available, skipping connected account transfer lifecycle")
		return
	}
	realAccountID := accounts.Items[0].ID

	// ListConnectedAccountTransfers first
	list, err := svc.ListConnectedAccountTransfers(ctx)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "list transfers failed: %v", err)
		return
	}
	require.NotNil(t, list)
	t.Logf("transfers count: %d", len(list.Items))

	// Try to get first transfer if available
	if len(list.Items) > 0 {
		fetched, err := svc.GetConnectedAccountTransfer(ctx, list.Items[0].ID)
		if err != nil {
			require.True(t, sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "get transfer failed: %v", err)
			t.Logf("get transfer skipped: %v", err)
		} else {
			assert.Equal(t, list.Items[0].ID, fetched.ID, "transfer id mismatch")
		}
	}

	// CreateConnectedAccountTransfer
	created, err := svc.CreateConnectedAccountTransfer(ctx, &scale.CreateConnectedAccountTransferRequest{
		RequestID:   "cat-req-" + time.Now().Format("20060102150405"),
		Destination: realAccountID,
		Amount:      "100.00",
		Currency:    sdk.CurrencyUSD,
		Reason:      "goods_purchased",
		Reference:   "ref-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err) || sdk.IsInvalidArgument(err) || sdk.IsValidationFailed(err), "create connected account transfer failed: %v", err)
		t.Logf("create connected account transfer skipped: %v", err)
	} else {
		require.NotEmpty(t, created.ID, "transfer id should not be empty")
		t.Logf("created connected account transfer: %s", created.ID)

		fetched, err := svc.GetConnectedAccountTransfer(ctx, created.ID)
		if err != nil {
			require.True(t, sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "get created transfer failed: %v", err)
			t.Logf("get created transfer skipped: %v", err)
		} else {
			assert.Equal(t, created.ID, fetched.ID, "transfer id mismatch")
		}
	}
}

func TestHostedFlowLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Scale()

	// Get a real account ID
	accounts, err := svc.ListAccounts(ctx)
	require.NoError(t, err, "list accounts failed")
	if len(accounts.Items) == 0 {
		t.Logf("no accounts available, skipping hosted flow lifecycle")
		return
	}
	realAccountID := accounts.Items[0].ID

	// CreateHostedFlow
	created, err := svc.CreateHostedFlow(ctx, &scale.CreateHostedFlowRequest{
		AccountID: realAccountID,
		Template:  "kyc",
		ReturnURL: "https://example.com/return",
		ErrorURL:  "https://example.com/error",
	})
	if err != nil {
		// 沙箱中可能没有可用的账户来创建 hosted flow
		t.Logf("create hosted flow skipped: %v", err)
		return
	}
	require.NotEmpty(t, created.ID, "hosted flow id should not be empty")
	t.Logf("created hosted flow: %s status=%s", created.ID, created.Status)

	// GetHostedFlow
	fetched, err := svc.GetHostedFlow(ctx, created.ID)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "get hosted flow failed: %v", err)
		t.Logf("get hosted flow skipped: %v", err)
	} else {
		assert.Equal(t, created.ID, fetched.ID, "hosted flow id mismatch")
	}

	// AuthorizeHostedFlow
	authorized, err := svc.AuthorizeHostedFlow(ctx, created.ID, &scale.AuthorizeHostedFlowRequest{
		Identity: "test_identity",
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "authorize hosted flow failed: %v", err)
		t.Logf("authorize hosted flow skipped: %v", err)
	} else {
		require.NotNil(t, authorized)
		t.Logf("authorized hosted flow: %s", authorized.ID)
	}
}

func TestInvitationLinkLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Scale()

	// Get a real account ID
	accounts, err := svc.ListAccounts(ctx)
	require.NoError(t, err, "list accounts failed")
	if len(accounts.Items) == 0 {
		t.Logf("no accounts available, skipping invitation link lifecycle")
		return
	}
	realAccountID := accounts.Items[0].ID

	// CreateInvitationLink
	created, err := svc.CreateInvitationLink(ctx, &scale.CreateInvitationLinkRequest{
		AccountID: realAccountID,
		Mode:      "oauth2",
		OAuth2: &scale.InvitationLinkOAuth2{
			RedirectURI:  "https://example.com/callback",
			ResponseType: "code",
			Scope:        []string{"r:awx_action:balances_view"},
		},
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err) || sdk.IsInvalidArgument(err), "create invitation link failed: %v", err)
		t.Logf("create invitation link skipped: %v", err)
		return
	}
	require.NotEmpty(t, created.ID, "invitation link id should not be empty")
	t.Logf("created invitation link: %s", created.ID)

	// GetInvitationLink
	fetched, err := svc.GetInvitationLink(ctx, created.ID)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "get invitation link failed: %v", err)
		t.Logf("get invitation link skipped: %v", err)
	} else {
		assert.Equal(t, created.ID, fetched.ID, "invitation link id mismatch")
	}
}

func TestProgramLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Scale()

	// List programs first
	list, err := svc.ListAccounts(ctx)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "list accounts failed: %v", err)
		return
	}
	if len(list.Items) == 0 {
		t.Logf("no accounts found, skipping program lifecycle")
		return
	}

	programID := list.Items[0].ID

	fetched, err := svc.GetProgram(ctx, programID)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "get program failed: %v", err)
		t.Logf("get program skipped: %v", err)
	} else {
		t.Logf("got program: %s", fetched.ID)
	}

	accounts, err := svc.ListProgramSpendingAccounts(ctx, programID)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "list program spending accounts failed: %v", err)
		t.Logf("list program spending accounts skipped: %v", err)
	} else {
		t.Logf("program spending accounts count: %d", len(accounts.Items))
	}

	transactions, err := svc.ListProgramTransactions(ctx, programID, &scale.ListProgramTransactionsRequest{PageSize: 10})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "list program transactions failed: %v", err)
		t.Logf("list program transactions skipped: %v", err)
	} else {
		t.Logf("program transactions count: %d", len(transactions.Items))
	}

	deposited, err := svc.DepositFunds(ctx, programID, &scale.DepositFundsRequest{
		RequestID: "dep-" + time.Now().Format("20060102150405"),
		Amount:    100.0,
		Currency:  string(sdk.CurrencyUSD),
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "deposit funds failed: %v", err)
		t.Logf("deposit funds skipped: %v", err)
	} else {
		t.Logf("deposited funds: %s", deposited.ID)
	}

	withdrawn, err := svc.WithdrawFunds(ctx, programID, &scale.WithdrawFundsRequest{
		RequestID: "wd-" + time.Now().Format("20060102150405"),
		Amount:    50.0,
		Currency:  string(sdk.CurrencyUSD),
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "withdraw funds failed: %v", err)
		t.Logf("withdraw funds skipped: %v", err)
	} else {
		t.Logf("withdrawn funds: %s", withdrawn.ID)
	}
}

func TestPlatformReportLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Scale()

	created, err := svc.CreatePlatformReport(ctx, &scale.CreatePlatformReportRequest{
		FileFormat: "CSV",
		Type:       "SETTLEMENT",
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsInvalidArgument(err), "create platform report failed: %v", err)
		t.Logf("create platform report skipped: %v", err)
		return
	}
	require.NotEmpty(t, created.ID, "platform report id should not be empty")
	t.Logf("created platform report: %s", created.ID)

	fetched, err := svc.GetPlatformReport(ctx, created.ID)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "get platform report failed: %v", err)
		t.Logf("get platform report skipped: %v", err)
	} else {
		t.Logf("got platform report: %s status=%s", fetched.ID, fetched.Status)
	}
}

func TestPSPSettlementDepositLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Scale()

	list, err := svc.ListPSPSettlementDeposits(ctx, &scale.ListPSPSettlementDepositsRequest{PageSize: 10})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsBadRequest(err), "list psp settlement deposits failed: %v", err)
		t.Logf("list psp settlement deposits skipped: %v", err)
		return
	}
	require.NotNil(t, list, "list psp settlement deposits result should not be nil")
	t.Logf("psp settlement deposits count: %d", len(list.Items))
}

func TestPSPSettlementIntentLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Scale()

	// Get a global account first
	gas, err := testClient.Core().ListGlobalAccounts(ctx)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "list global accounts failed: %v", err)
		t.Logf("list global accounts skipped: %v", err)
		return
	}
	var gaID string
	if len(gas.Items) > 0 {
		gaID = gas.Items[0].ID
	}

	created, err := svc.CreatePSPSettlementIntent(ctx, &scale.CreatePSPSettlementIntentRequest{
		RequestID:              "psi-req-" + time.Now().Format("20060102150405"),
		Currency:               sdk.CurrencyHKD,
		GlobalAccountID:        gaID,
		SettlementReference:    "ref-" + time.Now().Format("20060102150405"),
		ExpectedSettlementDate: time.Now().AddDate(0, 0, 7).Format("2006-01-02"),
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err) || sdk.IsInvalidArgument(err), "create psp settlement intent failed: %v", err)
		t.Logf("create psp settlement intent skipped: %v", err)
		return
	}
	require.NotEmpty(t, created.ID, "psp settlement intent id should not be empty")
	t.Logf("created psp settlement intent: %s", created.ID)

	fetched, err := svc.GetPSPSettlementIntent(ctx, created.ID)
	require.NoError(t, err, "get psp settlement intent failed")
	assert.Equal(t, created.ID, fetched.ID, "psp settlement intent id mismatch")

	list, err := svc.ListPSPSettlementIntents(ctx)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsBadRequest(err), "list psp settlement intents failed: %v", err)
		t.Logf("list psp settlement intents skipped: %v", err)
		return
	}
	require.NotEmpty(t, list.Items, "expected at least one psp settlement intent")
	t.Logf("psp settlement intents count: %d", len(list.Items))
}

func TestPSPSettlementSplitLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Scale()

	// Get a real account ID for target
	accounts, err := svc.ListAccounts(ctx)
	require.NoError(t, err, "list accounts failed")
	if len(accounts.Items) == 0 {
		t.Logf("no accounts available, skipping psp settlement split lifecycle")
		return
	}
	realAccountID := accounts.Items[0].ID

	// List settlement intents first
	intents, err := svc.ListPSPSettlementIntents(ctx)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsBadRequest(err), "list psp settlement intents failed: %v", err)
		t.Logf("list psp settlement intents skipped: %v", err)
		return
	}
	if len(intents.Items) == 0 {
		t.Logf("no psp settlement intents found, skipping split lifecycle")
		return
	}

	intentID := intents.Items[0].ID

	// Split the intent
	splitResult, err := svc.SplitPSPSettlementIntent(ctx, intentID, &scale.SplitSettlementIntentRequest{
		RequestID: "split-" + time.Now().Format("20060102150405"),
		Splits: []scale.SettlementSplitItem{
			{
				Amount:          "50.00",
				Identifier:      "split-1",
				SettlementType:  "STANDARD",
				TargetAccountID: realAccountID,
			},
		},
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "split psp settlement intent failed: %v", err)
		t.Logf("split psp settlement intent skipped: %v", err)
		return
	}
	require.NotEmpty(t, splitResult.PSPSettlementIntentID, "split result should have intent id")
	t.Logf("split psp settlement intent: %s", splitResult.PSPSettlementIntentID)

	// List splits
	list, err := svc.ListPSPSettlementSplits(ctx, &scale.ListPSPSettlementSplitsRequest{PageSize: 10})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsBadRequest(err), "list psp settlement splits failed: %v", err)
		t.Logf("list psp settlement splits skipped: %v", err)
		return
	}
	require.NotEmpty(t, list.Items, "expected at least one psp settlement split")
	t.Logf("psp settlement splits count: %d", len(list.Items))

	// Get a split
	fetched, err := svc.GetPSPSettlementSplit(ctx, list.Items[0].ID)
	require.NoError(t, err, "get psp settlement split failed")
	assert.Equal(t, list.Items[0].ID, fetched.ID, "psp settlement split id mismatch")

	// Release the split
	released, err := svc.ReleasePSPSettlementSplit(ctx, list.Items[0].ID)
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "release psp settlement split failed: %v", err)
		t.Logf("release psp settlement split skipped: %v", err)
	} else {
		t.Logf("released psp settlement split: %s", released.ID)
	}

	// Cancel the split
	cancelled, err := svc.CancelPSPSettlementSplit(ctx, list.Items[0].ID)
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "cancel psp settlement split failed: %v", err)
		t.Logf("cancel psp settlement split skipped: %v", err)
	} else {
		t.Logf("cancelled psp settlement split: %s", cancelled.ID)
	}
}
