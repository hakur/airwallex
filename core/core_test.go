package core_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/core"
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

func TestGetCurrentBalances(t *testing.T) {
	ctx := context.Background()

	balances, err := testClient.Core().GetCurrentBalances(ctx)
	require.NoError(t, err, "get current balances failed")
	t.Logf("got %d currency balances", len(balances))
	for _, b := range balances {
		if b.AvailableAmount > 0 {
			t.Logf("  %s: available=%.2f total=%.2f", b.Currency, b.AvailableAmount, b.TotalAmount)
		}
	}
}

func TestDirectDebitsListAndGet(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Core()

	list, err := svc.ListDirectDebits(ctx)
	require.NoError(t, err, "list direct debits failed")
	t.Logf("direct debits count: %d", len(list.Items))

	if len(list.Items) > 0 {
		dd, err := svc.GetDirectDebit(ctx, list.Items[0].TransactionID)
		require.NoError(t, err, "get direct debit failed")
		assert.Equal(t, list.Items[0].TransactionID, dd.TransactionID, "direct debit id mismatch")
		t.Logf("got direct debit: %s status=%s", dd.TransactionID, dd.Status)

		// 尝试取消直接扣款（根据状态可能成功或失败）
		err = svc.CancelDirectDebit(ctx, dd.TransactionID)
		if err != nil {
			assert.True(t, sdk.IsBadRequest(err) || sdk.IsNotFound(err) || sdk.IsValidationError(err),
				"unexpected error type: %v", err)
		}
	}
}

func TestDepositLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Core()

	// List deposits first
	list, err := svc.ListDeposits(ctx, &core.ListDepositsRequest{PageSize: 10})
	require.NoError(t, err, "list deposits failed")
	t.Logf("deposits count: %d", len(list))

	if len(list) > 0 {
		dep, err := svc.GetDeposit(ctx, list[0].ID)
		require.NoError(t, err, "get deposit failed")
		assert.Equal(t, list[0].ID, dep.ID, "deposit id mismatch")
		t.Logf("got deposit: %s status=%s", dep.ID, dep.Status)
	}
}

func TestBalanceHistory(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Core()

	resp, err := svc.GetBalanceHistory(ctx, &core.GetBalanceHistoryRequest{PageSize: 10})
	require.NoError(t, err, "get balance history failed")
	assert.NotNil(t, resp, "balance history response should not be nil")
	t.Logf("balance history items: %d", len(resp.Items))
}

func TestGlobalAccountLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Core()

	// List global accounts
	list, err := svc.ListGlobalAccounts(ctx)
	require.NoError(t, err, "list global accounts failed")
	assert.NotNil(t, list, "global accounts list should not be nil")
	t.Logf("global accounts count: %d", len(list.Items))

	if len(list.Items) > 0 {
		ga := list.Items[0]
		fetched, err := svc.GetGlobalAccount(ctx, ga.ID)
		require.NoError(t, err, "get global account failed")
		assert.Equal(t, ga.ID, fetched.ID, "global account id mismatch")
		t.Logf("got global account: %s currency=%s", fetched.ID, fetched.Currency)

		// Test transactions
		_, err = svc.GetGlobalAccountTransactions(ctx, ga.ID, &core.GetGlobalAccountTransactionsRequest{PageSize: 10})
		if err != nil {
			require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err), "get global account transactions failed: %v", err)
			t.Logf("get global account transactions skipped: %v", err)
		}
	}
}

func TestLinkedAccountLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Core()

	list, err := svc.ListLinkedAccounts(ctx)
	require.NoError(t, err, "list linked accounts failed")
	assert.NotNil(t, list, "linked accounts list should not be nil")
	t.Logf("linked accounts count: %d", len(list.Items))

	if len(list.Items) > 0 {
		la := list.Items[0]
		fetched, err := svc.GetLinkedAccount(ctx, la.ID)
		require.NoError(t, err, "get linked account failed")
		assert.Equal(t, la.ID, fetched.ID, "linked account id mismatch")
		t.Logf("got linked account: %s type=%s", fetched.ID, fetched.Type)
	}
}
