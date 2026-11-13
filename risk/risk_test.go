package risk_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/risk"
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

func TestFraudFeedbackLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Risk()

	// List authorizations first to get an authorization ID
	auths, err := testClient.Issuing().ListAuthorizations(ctx)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "list authorizations failed: %v", err)
		t.Logf("list authorizations skipped: %v", err)
		return
	}
	if len(auths.Items) == 0 {
		t.Logf("no authorizations found, skipping fraud feedback lifecycle")
		return
	}

	created, err := svc.CreateFraudFeedback(ctx, &risk.CreateFraudFeedbackRequest{
		AuthorizationID:    auths.Items[0].TransactionID,
		CardholderDecision: risk.CardholderDecisionAuthorized,
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsBadRequest(err) || sdk.IsNotFound(err), "create fraud feedback failed: %v", err)
		t.Logf("create fraud feedback skipped: %v", err)
		return
	}
	require.NotNil(t, created)
	t.Logf("created fraud feedback: authorization_id=%s", created.AuthorizationID)
}

func TestWatchlistLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Risk()

	created, err := svc.CreateWatchlistEntry(ctx, &risk.CreateWatchlistEntryRequest{
		RequestID: "wl-req-" + time.Now().Format("20060102150405"),
		Action:    risk.WatchlistActionBlock,
		Type:      risk.WatchlistItemTypeCustomerEmail,
		Value:     "test-" + time.Now().Format("20060102150405") + "@example.com",
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "create watchlist failed: %v", err)
		return
	}
	require.NotNil(t, created)
	t.Logf("created watchlist entry: %s action=%s", created.ID, created.Action)

	fetched, err := svc.GetWatchlistEntry(ctx, created.ID)
	require.NoError(t, err, "get watchlist entry failed")
	assert.Equal(t, created.ID, fetched.ID, "watchlist entry id mismatch")

	updated, err := svc.UpdateWatchlistEntry(ctx, created.ID, &risk.UpdateWatchlistEntryRequest{
		RequestID: "wl-upd-" + time.Now().Format("20060102150405"),
		Status:    risk.WatchlistItemStatusInactive,
	})
	require.NoError(t, err, "update watchlist entry failed")
	assert.Equal(t, risk.WatchlistItemStatusInactive, updated.Status)

	list, err := svc.ListWatchlistEntries(ctx)
	require.NoError(t, err, "list watchlist entries failed")
	require.NotEmpty(t, list.Items, "expected at least one watchlist entry")
}

func TestSellerLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Risk()

	created, err := svc.CreateSeller(ctx, &risk.CreateSellerRequest{
		RequestID:           "seller-req-" + time.Now().Format("20060102150405"),
		LegalEntityName:     "Test Seller " + time.Now().Format("20060102150405"),
		TradingName:         "Test Trading Name",
		Email:               "seller-test@example.com",
		RegistrationCountry: "AU",
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsForbidden(err), "create seller failed: %v", err)
		return
	}
	require.NotNil(t, created)
	t.Logf("created seller: %s status=%s", created.ID, created.Status)

	fetched, err := svc.GetSeller(ctx, created.ID)
	require.NoError(t, err, "get seller failed")
	assert.Equal(t, created.ID, fetched.ID, "seller id mismatch")

	deactivated, err := svc.DeactivateSeller(ctx, created.ID)
	require.NoError(t, err, "deactivate seller failed")
	assert.Equal(t, created.ID, deactivated.ID)
}

func TestRFILifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Risk()

	list, err := svc.ListRFIs(ctx)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "list rfis failed: %v", err)
		t.Logf("list rfis skipped: %v", err)
		return
	}
	require.NotNil(t, list)
	t.Logf("rfi count: %d", len(list.Items))
	if len(list.Items) == 0 {
		t.Logf("no rfis found, skipping get and respond")
		return
	}

	id := list.Items[0].ID
	fetched, err := svc.GetRFI(ctx, id)
	require.NoError(t, err, "get rfi failed")
	assert.Equal(t, id, fetched.ID, "rfi id mismatch")

	responded, err := svc.RespondRFI(ctx, id, &risk.RespondRFIRequest{
		Questions: []risk.RFIQuestion{},
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "respond rfi failed: %v", err)
		t.Logf("respond rfi skipped: %v", err)
		return
	}
	t.Logf("responded rfi: %s", responded.ID)
}
