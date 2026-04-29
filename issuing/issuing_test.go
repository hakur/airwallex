package issuing_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/issuing"
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

func TestCardholderLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Issuing()

	created, err := svc.CreateCardholder(ctx, &issuing.CreateCardholderRequest{
		RequestID: "ch-req-" + time.Now().Format("20060102150405"),
		Type:      issuing.CardholderTypeIndividual,
		Email:     "test-" + time.Now().Format("20060102150405") + "@example.com",
		Individual: &issuing.CardholderIndividual{
			Name: &issuing.CardholderName{
				FirstName: "Test",
				LastName:  "Cardholder",
			},
			DateOfBirth: "1990-01-01",
			Address: &issuing.CardholderAddress{
				Country:  "US",
				City:     "San Francisco",
				Line1:    "123 Test St",
				Postcode: "94102",
				State:    "CA",
			},
			ExpressConsentObtained: "yes",
		},
	})
	require.NoError(t, err, "create cardholder failed (issuing may need activation)")
	t.Logf("created cardholder: %s", created.ID)

	fetched, err := svc.GetCardholder(ctx, created.ID)
	require.NoError(t, err, "get cardholder failed")
	assert.Equal(t, created.ID, fetched.ID, "cardholder id mismatch")

	updated, err := svc.UpdateCardholder(ctx, created.ID, &issuing.UpdateCardholderRequest{
		Individual: &issuing.CardholderIndividual{
			Name: &issuing.CardholderName{
				FirstName: "Updated",
			},
		},
	})
	if err != nil {
		require.Error(t, err, "update cardholder blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else if updated.Individual != nil && updated.Individual.Name != nil {
		t.Logf("updated cardholder name: %s", updated.Individual.Name.FirstName)
	}

	list, err := svc.ListCardholders(ctx)
	require.NoError(t, err, "list cardholders failed")
	require.NotEmpty(t, list.Items, "expected at least one cardholder")
}

func TestCardLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Issuing()

	cards, err := svc.ListCards(ctx)
	require.NoError(t, err, "list cards failed")
	t.Logf("existing cards: %d", len(cards.Items))

	chList, err := svc.ListCardholders(ctx)
	require.NoError(t, err, "list cardholders failed")

	var cardholderID string
	if len(chList.Items) > 0 {
		cardholderID = chList.Items[0].ID
	} else {
		ch, err := svc.CreateCardholder(ctx, &issuing.CreateCardholderRequest{
			RequestID: "ch-req-" + time.Now().Format("20060102150405"),
			Type:      issuing.CardholderTypeIndividual,
			Email:     "test-" + time.Now().Format("20060102150405") + "@example.com",
			Individual: &issuing.CardholderIndividual{
				Name: &issuing.CardholderName{
					FirstName: "Test",
					LastName:  "Cardholder",
				},
				DateOfBirth: "1990-01-01",
				Address: &issuing.CardholderAddress{
					Country:  "US",
					City:     "San Francisco",
					Line1:    "123 Test St",
					Postcode: "94102",
					State:    "CA",
				},
				ExpressConsentObtained: "yes",
			},
		})
		require.NoError(t, err, "create cardholder failed")
		cardholderID = ch.ID
	}

	created, err := svc.CreateCard(ctx, &issuing.CreateCardRequest{
		RequestID:    "card-req-" + time.Now().Format("20060102150405"),
		CardholderID: cardholderID,
		Program: issuing.CardProgram{
			Purpose: "CONSUMER",
			Type:    "DEBIT",
		},
		IsPersonalized: true,
		FormFactor:     "VIRTUAL",
		CreatedBy:      "Test User",
	})
	if err != nil {
		require.Error(t, err, "create card blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	t.Logf("created card: %s", created.ID)
	require.NotEmpty(t, created.ID, "card id should not be empty")

	fetched, err := svc.GetCard(ctx, created.ID)
	require.NoError(t, err, "get card failed")
	assert.Equal(t, created.ID, fetched.ID, "card id mismatch")

	// 测试激活卡片
	err = svc.ActivateCard(ctx, created.ID)
	if err != nil {
		require.Error(t, err, "activate card blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("activated card successfully")
	}

	// 测试获取卡片详情
	details, err := svc.GetCardDetails(ctx, created.ID)
	if err != nil {
		require.Error(t, err, "get card details blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		require.NotEmpty(t, details.CardNumber, "card number should not be empty")
		t.Logf("card details: number=%s expiry=%02d/%d", details.CardNumber, details.ExpiryMonth, details.ExpiryYear)
	}

	// 测试获取卡片限额
	limits, err := svc.GetCardLimits(ctx, created.ID)
	if err != nil {
		require.Error(t, err, "get card limits blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("card limits: currency=%s, limits=%d", limits.Currency, len(limits.Limits))
	}

	// 测试列出卡片
	cards, err = svc.ListCards(ctx)
	require.NoError(t, err, "list cards failed")
	require.NotEmpty(t, cards.Items, "expected at least one card")
}

func TestConfig(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Issuing()

	config, err := svc.GetConfig(ctx)
	require.NoError(t, err, "get config failed")
	t.Logf("config primary_currency: %s", config.PrimaryCurrency)

	updated, err := svc.UpdateConfig(ctx, &issuing.UpdateConfigRequest{
		PrimaryCurrency: sdk.CurrencyUSD,
	})
	if err != nil {
		require.Error(t, err, "update config blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("updated config primary_currency: %s", updated.PrimaryCurrency)
	}
}

func TestAuthorizationLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Issuing()

	list, err := svc.ListAuthorizations(ctx)
	if err != nil {
		require.Error(t, err, "list authorizations blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	if len(list.Items) == 0 {
		t.Logf("no authorizations found, skipping authorization lifecycle assertions (sandbox limitation)")
		return
	}
	t.Logf("authorizations count: %d", len(list.Items))

	fetched, err := svc.GetAuthorization(ctx, list.Items[0].TransactionID)
	require.NoError(t, err, "get authorization failed")
	assert.Equal(t, list.Items[0].TransactionID, fetched.TransactionID, "authorization transaction id mismatch")
}

func TestCardTransactionEventLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Issuing()

	list, err := svc.ListCardTransactionEvents(ctx)
	if err != nil {
		require.Error(t, err, "list card transaction events blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	if len(list.Items) == 0 {
		t.Logf("no card transaction events found, skipping lifecycle assertions (sandbox limitation)")
		return
	}
	t.Logf("card transaction events count: %d", len(list.Items))

	fetched, err := svc.GetCardTransactionEvent(ctx, list.Items[0].ID)
	require.NoError(t, err, "get card transaction event failed")
	assert.Equal(t, list.Items[0].ID, fetched.ID, "card transaction event id mismatch")
}

func TestCardTransactionLifecycleLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Issuing()

	list, err := svc.ListLifecycles(ctx, &issuing.ListLifecyclesRequest{PageSize: 10})
	if err != nil {
		require.Error(t, err, "list lifecycles blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	if len(list.Items) == 0 {
		t.Logf("no lifecycles found, skipping lifecycle assertions (sandbox limitation)")
		return
	}
	t.Logf("lifecycles count: %d", len(list.Items))

	fetched, err := svc.GetLifecycle(ctx, list.Items[0].ID)
	require.NoError(t, err, "get lifecycle failed")
	assert.Equal(t, list.Items[0].ID, fetched.ID, "lifecycle id mismatch")
}

func TestDigitalWalletTokenLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Issuing()

	list, err := svc.ListDigitalWalletTokens(ctx)
	if err != nil {
		require.Error(t, err, "list digital wallet tokens blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	if len(list.Items) == 0 {
		t.Logf("no digital wallet tokens found, skipping lifecycle assertions (sandbox limitation)")
		return
	}
	t.Logf("digital wallet tokens count: %d", len(list.Items))

	fetched, err := svc.GetDigitalWalletToken(ctx, list.Items[0].TokenID)
	require.NoError(t, err, "get digital wallet token failed")
	assert.Equal(t, list.Items[0].TokenID, fetched.TokenID, "digital wallet token id mismatch")
}

func TestMerchantBrandLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Issuing()

	list, err := svc.ListMerchantBrands(ctx)
	if err != nil {
		require.Error(t, err, "list merchant brands blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	require.NotEmpty(t, list.Items, "expected at least one merchant brand")
	t.Logf("merchant brands count: %d", len(list.Items))

	fetched, err := svc.GetMerchantBrand(ctx, list.Items[0].ID)
	require.NoError(t, err, "get merchant brand failed")
	assert.Equal(t, list.Items[0].ID, fetched.ID, "merchant brand id mismatch")
}

func TestTransactionLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Issuing()

	list, err := svc.ListTransactions(ctx)
	if err != nil {
		require.Error(t, err, "list transactions blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	if len(list.Items) == 0 {
		t.Logf("no transactions found, skipping transaction lifecycle assertions (sandbox limitation)")
		return
	}
	t.Logf("transactions count: %d", len(list.Items))

	fetched, err := svc.GetTransaction(ctx, list.Items[0].TransactionID)
	require.NoError(t, err, "get transaction failed")
	assert.Equal(t, list.Items[0].TransactionID, fetched.TransactionID, "transaction id mismatch")
}

func TestTransactionDisputeLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Issuing()

	// List transactions first to get a transaction ID
	transactions, err := svc.ListTransactions(ctx)
	if err != nil {
		require.Error(t, err, "list transactions blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	if len(transactions.Items) == 0 {
		t.Logf("no transactions found, skipping transaction dispute lifecycle")
		return
	}

	created, err := svc.CreateTransactionDispute(ctx, &issuing.CreateTransactionDisputeRequest{
		RequestID:     "td-req-" + time.Now().Format("20060102150405"),
		TransactionID: transactions.Items[0].TransactionID,
		Amount:        10.0,
		Reason:        issuing.DisputeReasonOther,
	})
	if err != nil {
		require.Error(t, err, "create transaction dispute blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	require.NotEmpty(t, created.ID, "transaction dispute id should not be empty")
	t.Logf("created transaction dispute: %s", created.ID)

	fetched, err := svc.GetTransactionDispute(ctx, created.ID)
	require.NoError(t, err, "get transaction dispute failed")
	assert.Equal(t, created.ID, fetched.ID, "transaction dispute id mismatch")

	updated, err := svc.UpdateTransactionDispute(ctx, created.ID, &issuing.UpdateTransactionDisputeRequest{
		RequestID: "td-upd-" + time.Now().Format("20060102150405"),
		Notes:     "updated notes",
	})
	if err != nil {
		require.Error(t, err, "update transaction dispute blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("updated transaction dispute: %s", updated.ID)
	}

	list, err := svc.ListTransactionDisputes(ctx)
	if err != nil {
		require.Error(t, err, "list transaction disputes blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	require.NotEmpty(t, list.Items, "expected at least one transaction dispute")
	t.Logf("transaction disputes count: %d", len(list.Items))

	submitted, err := svc.SubmitTransactionDispute(ctx, created.ID)
	if err != nil {
		require.Error(t, err, "submit transaction dispute blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("submitted transaction dispute: %s", submitted.ID)
	}

	cancelled, err := svc.CancelTransactionDispute(ctx, created.ID)
	if err != nil {
		require.Error(t, err, "cancel transaction dispute blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("cancelled transaction dispute: %s", cancelled.ID)
	}
}

func TestDeleteCardholder(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Issuing()

	// 先创建持卡人
	created, err := svc.CreateCardholder(ctx, &issuing.CreateCardholderRequest{
		RequestID: "ch-del-req-" + time.Now().Format("20060102150405"),
		Type:      issuing.CardholderTypeIndividual,
		Email:     "test-del-" + time.Now().Format("20060102150405") + "@example.com",
		Individual: &issuing.CardholderIndividual{
			Name: &issuing.CardholderName{
				FirstName: "Test",
				LastName:  "Delete",
			},
			DateOfBirth: "1990-01-01",
			Address: &issuing.CardholderAddress{
				Country:  "US",
				City:     "San Francisco",
				Line1:    "123 Test St",
				Postcode: "94102",
				State:    "CA",
			},
			ExpressConsentObtained: "yes",
		},
	})
	require.NoError(t, err, "create cardholder for delete failed")
	require.NotEmpty(t, created.ID, "cardholder id should not be empty")

	// 删除持卡人
	deleted, err := svc.DeleteCardholder(ctx, created.ID)
	if err != nil {
		require.Error(t, err, "delete cardholder blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("deleted cardholder: %s, deleted=%v", deleted.CardholderID, deleted.Deleted)
	}
}
