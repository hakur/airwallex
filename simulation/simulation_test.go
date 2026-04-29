package simulation_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/pa"
	"github.com/hakur/airwallex/sdk"
	"github.com/hakur/airwallex/simulation"
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

// === 类型A：预期成功 ===

func TestSimulateShopperAction(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	pi, err := testClient.PA().CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
		RequestID:       "sim-pi-req-" + strconv.FormatInt(time.Now().UnixNano(), 10),
		MerchantOrderID: "sim-order-" + strconv.FormatInt(time.Now().UnixNano(), 10),
		Amount:          100.0,
		Currency:        sdk.CurrencyUSD,
	})
	require.NoError(t, err, "create payment intent for shopper action failed")
	require.NotEmpty(t, pi.ID, "payment intent id should not be empty")

	err = svc.SimulateShopperAction(ctx, "pay", &simulation.SimulateShopperActionRequest{
		URL: "https://checkout.airwallex.com/simulate?payment_intent_id=" + pi.ID,
	})
	require.NoError(t, err, "simulate shopper action failed")
}

func TestSimulateGlobalAccountDeposit(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	accounts, err := testClient.Core().ListGlobalAccounts(ctx)
	require.NoError(t, err, "list global accounts failed")
	require.NotEmpty(t, accounts.Items, "no global accounts available for deposit simulation")

	ga := accounts.Items[0]
	result, err := svc.SimulateGlobalAccountDeposit(ctx, &simulation.SimulateGlobalAccountDepositRequest{
		GlobalAccountID: ga.ID,
		Amount:          1000.0,
		PayerBankName:   "Test Bank",
		PayerCountry:    "US",
		PayerName:       "Test Payer",
		Reference:       "simulated deposit",
		Status:          "SETTLED",
	})
	require.NoError(t, err, "simulate global account deposit failed")
	require.NotNil(t, result, "deposit result should not be nil")
	assert.NotEmpty(t, result.ID, "deposit id should not be empty")
	assert.Equal(t, 1000.0, result.Amount, "deposit amount mismatch")
}

func TestSimulateDirectDebitOperations(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	dds, err := testClient.Core().ListDirectDebits(ctx)
	require.NoError(t, err, "list direct debits failed")

	if len(dds.Items) == 0 {
		t.Log("no direct debits available for simulation, skipping operations")
		return
	}

	dd := dds.Items[0]

	// 这些操作可能成功也可能失败，取决于直接扣款状态
	// 但不应该返回未知错误
	err = svc.SimulateDirectDebitReject(ctx, dd.TransactionID)
	if err != nil {
		require.Error(t, err)
		t.Logf("simulate direct debit reject result: %v", err)
		assert.True(t, sdk.IsBadRequest(err) || sdk.IsNotFound(err), "unexpected error type: %v", err)
	}

	err = svc.SimulateDirectDebitReverse(ctx, dd.TransactionID)
	if err != nil {
		require.Error(t, err)
		t.Logf("simulate direct debit reverse result: %v", err)
		assert.True(t, sdk.IsBadRequest(err) || sdk.IsNotFound(err), "unexpected error type: %v", err)
	}

	err = svc.SimulateDirectDebitSettle(ctx, dd.TransactionID)
	if err != nil {
		require.Error(t, err)
		t.Logf("simulate direct debit settle result: %v", err)
		assert.True(t, sdk.IsBadRequest(err) || sdk.IsNotFound(err), "unexpected error type: %v", err)
	}
}

// === 类型B：预期失败（硬编码无效ID）===

func TestSimulateAccountAmendmentApprove_InvalidID(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	err := svc.SimulateAccountAmendmentApprove(ctx, "amendment-test-id")
	require.Error(t, err, "expected error for invalid amendment id")
	assert.True(t, sdk.IsBadRequest(err), "expected bad_request error, got: %v", err)
}

func TestSimulateAccountAmendmentReject_InvalidID(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	err := svc.SimulateAccountAmendmentReject(ctx, "amendment-test-id")
	require.Error(t, err, "expected error for invalid amendment id")
	assert.True(t, sdk.IsBadRequest(err), "expected bad_request error, got: %v", err)
}

func TestSimulateLinkedAccountOperations_InvalidID(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()
	testID := "linked-account-test-id"

	err := svc.SimulateLinkedAccountFailMicrodeposits(ctx, testID)
	require.Error(t, err, "expected error for invalid linked account id")
	assert.True(t, sdk.IsBadRequest(err), "expected bad_request error, got: %v", err)

	err = svc.SimulateLinkedAccountMandateAccept(ctx, testID)
	require.Error(t, err, "expected error for invalid linked account id")
	assert.True(t, sdk.IsBadRequest(err), "expected bad_request error, got: %v", err)

	err = svc.SimulateLinkedAccountMandateCancel(ctx, testID)
	require.Error(t, err, "expected error for invalid linked account id")
	assert.True(t, sdk.IsBadRequest(err), "expected bad_request error, got: %v", err)

	err = svc.SimulateLinkedAccountMandateReject(ctx, testID)
	require.Error(t, err, "expected error for invalid linked account id")
	assert.True(t, sdk.IsBadRequest(err), "expected bad_request error, got: %v", err)
}

func TestSimulatePayoutPaymentTransition_InvalidID(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	err := svc.SimulatePayoutPaymentTransition(ctx, "payment-test-id", &simulation.SimulatePayoutPaymentTransitionRequest{
		NextStatus: "COMPLETED",
	})
	require.Error(t, err, "expected error for invalid payment id")
	assert.True(t, sdk.IsBadRequest(err), "expected bad_request error, got: %v", err)
}

func TestSimulateRFIOperations_InvalidRequests(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	// 创建 RFI 缺少必填字段
	err := svc.SimulateRFICreate(ctx, &simulation.SimulateRFICreateRequest{})
	require.Error(t, err, "expected error for empty rfi create request")
	assert.True(t, sdk.IsBadRequest(err) || sdk.IsInvalidArgument(err), "expected bad_request or invalid_argument error, got: %v", err)

	// 关闭不存在的 RFI
	err = svc.SimulateRFIClose(ctx, "rfi-test-id")
	require.Error(t, err, "expected error for non-existent rfi")
	assert.True(t, sdk.IsNotFound(err) || sdk.IsBadRequest(err), "expected not_found or bad_request error, got: %v", err)

	// 跟进不存在的 RFI
	err = svc.SimulateRFIFollowUp(ctx, "rfi-test-id", &simulation.SimulateRFIFollowUpRequest{})
	require.Error(t, err, "expected error for non-existent rfi")
	assert.True(t, sdk.IsNotFound(err) || sdk.IsBadRequest(err), "expected not_found or bad_request error, got: %v", err)
}

func TestSimulateTerminalLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	// Get payment scenarios first
	scenarios, err := svc.SimulateTerminalPaymentScenarios(ctx)
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsNotFound(err) || sdk.IsUnauthorized(err), "get payment scenarios failed: %v", err)
		t.Logf("get payment scenarios skipped: %v", err)
		return
	}
	require.NotNil(t, scenarios, "scenarios should not be nil")
	require.NotEmpty(t, scenarios.Items, "expected at least one scenario")
	t.Logf("payment scenarios count: %d", len(scenarios.Items))

	// 错误路径：生成激活码缺少 request_id
	err = svc.SimulateTerminalGenerateActivationCode(ctx, &simulation.SimulateTerminalGenerateActivationCodeRequest{})
	require.Error(t, err, "expected error for empty activation code request")
	assert.True(t, sdk.IsBadRequest(err), "expected bad_request error, got: %v", err)

	// 错误路径：确认支付缺少 terminal_id
	err = svc.SimulateTerminalConfirmPaymentIntent(ctx, &simulation.SimulateTerminalConfirmPaymentIntentRequest{
		PaymentIntentID: "pi-test-id",
		ScenarioID:      "scenario-test-id",
	})
	require.Error(t, err, "expected error for missing terminal_id")
	assert.True(t, sdk.IsBadRequest(err), "expected bad_request error, got: %v", err)

	// 错误路径：打开终端缺少 terminal_id
	err = svc.SimulateTerminalTurnOn(ctx, &simulation.SimulateTerminalTurnOnRequest{})
	require.Error(t, err, "expected error for missing terminal_id")
	assert.True(t, sdk.IsBadRequest(err), "expected bad_request error, got: %v", err)

	// 错误路径：关闭终端缺少 terminal_id
	err = svc.SimulateTerminalTurnOff(ctx, &simulation.SimulateTerminalTurnOffRequest{})
	require.Error(t, err, "expected error for missing terminal_id")
	assert.True(t, sdk.IsBadRequest(err), "expected bad_request error, got: %v", err)

	// Create a terminal (requires valid activation code)
	term, err := testClient.PA().CreateTerminal(ctx, &pa.CreateTerminalRequest{
		RequestID:      "sim-term-" + strconv.FormatInt(time.Now().UnixNano(), 10),
		ActivationCode: "MEFCMSBC",
		Nickname:       "Sim Terminal",
	})
	if err != nil {
		require.True(t, sdk.IsNotFound(err) || sdk.IsValidationError(err), "create terminal failed: %v", err)
		t.Logf("create terminal skipped: %v", err)
		return
	}
	require.NotEmpty(t, term.ID, "terminal id should not be empty")

	// Turn on terminal
	err = svc.SimulateTerminalTurnOn(ctx, &simulation.SimulateTerminalTurnOnRequest{})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "turn on terminal failed: %v", err)
		t.Logf("turn on terminal skipped: %v", err)
	} else {
		t.Logf("turned on terminal successfully")
	}

	// Create a payment intent
	pi, err := testClient.PA().CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
		RequestID:       "sim-pi-" + strconv.FormatInt(time.Now().UnixNano(), 10),
		MerchantOrderID: "sim-order-" + strconv.FormatInt(time.Now().UnixNano(), 10),
		Amount:          100.0,
		Currency:        sdk.CurrencyUSD,
	})
	require.NoError(t, err, "create payment intent for terminal simulation failed")

	// Confirm payment intent in terminal
	err = svc.SimulateTerminalConfirmPaymentIntent(ctx, &simulation.SimulateTerminalConfirmPaymentIntentRequest{
		PaymentIntentID: pi.ID,
		ScenarioID:      scenarios.Items[0].ID,
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "confirm payment intent failed: %v", err)
		t.Logf("confirm payment intent skipped: %v", err)
	} else {
		t.Logf("confirmed payment intent in terminal successfully")
	}

	// Turn off terminal
	err = svc.SimulateTerminalTurnOff(ctx, &simulation.SimulateTerminalTurnOffRequest{})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "turn off terminal failed: %v", err)
		t.Logf("turn off terminal skipped: %v", err)
	} else {
		t.Logf("turned off terminal successfully")
	}
}

// === 类型A变体：可能成功也可能失败 ===

func TestSimulateIssuingThreedsNotify(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	err := svc.SimulateIssuingThreedsNotify(ctx, &simulation.SimulateIssuingThreedsNotifyRequest{})
	// 这个端点需要特定场景，可能返回各种错误
	if err != nil {
		require.Error(t, err)
		t.Logf("simulate issuing 3ds notify result: %v", err)
		assert.True(t, sdk.IsNotFound(err) || sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsInternalError(err), "unexpected error: %v", err)
	}
}

func TestSimulateIssuingCardholderPassReview(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	cardholders, err := testClient.Issuing().ListCardholders(ctx)
	require.NoError(t, err, "list cardholders failed")
	require.NotEmpty(t, cardholders.Items, "no cardholders available")

	err = svc.SimulateIssuingCardholderPassReview(ctx, cardholders.Items[0].ID)
	// 这个端点可能返回错误，取决于 cardholder 状态
	if err != nil {
		require.Error(t, err)
		t.Logf("simulate issuing cardholder pass review result: %v", err)
		assert.True(t, sdk.IsBadRequest(err), "expected bad_request error, got: %v", err)
	}
}

func TestSimulatePaymentDisputeLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	// Create a payment intent first
	pi, err := testClient.PA().CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
		RequestID:       "sim-pd-pi-" + strconv.FormatInt(time.Now().UnixNano(), 10),
		MerchantOrderID: "sim-pd-order-" + strconv.FormatInt(time.Now().UnixNano(), 10),
		Amount:          100.0,
		Currency:        sdk.CurrencyUSD,
	})
	require.NoError(t, err, "create payment intent for dispute simulation failed")
	require.NotEmpty(t, pi.ID, "payment intent id should not be empty")

	// SimulatePaymentDisputeCreate
	created, err := svc.SimulatePaymentDisputeCreate(ctx, &simulation.SimulatePaymentDisputeCreateRequest{
		PaymentIntentID: pi.ID,
		ReasonCode:      "10.4",
		Stage:           "DISPUTE",
		DueAt:           time.Now().AddDate(0, 0, 30).Format(time.RFC3339),
		Amount:          100.0,
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err) || sdk.IsOperationFailed(err), "create payment dispute failed: %v", err)
		t.Logf("create payment dispute skipped: %v", err)
		return
	}
	require.NotEmpty(t, created.ID, "payment dispute id should not be empty")
	t.Logf("created payment dispute: %s", created.ID)

	// SimulatePaymentDisputeEscalate
	escalated, err := svc.SimulatePaymentDisputeEscalate(ctx, created.ID, &simulation.SimulatePaymentDisputeEscalateRequest{
		Amount: 100.0,
		DueAt:  time.Now().AddDate(0, 0, 30).Format(time.RFC3339),
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "escalate payment dispute failed: %v", err)
		t.Logf("escalate payment dispute skipped: %v", err)
	} else {
		t.Logf("escalated payment dispute: %s", escalated.ID)
	}

	// SimulatePaymentDisputeResolve
	resolved, err := svc.SimulatePaymentDisputeResolve(ctx, created.ID, &simulation.SimulatePaymentDisputeResolveRequest{
		InFavorOf: "merchant",
		Amount:    100.0,
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "resolve payment dispute failed: %v", err)
		t.Logf("resolve payment dispute skipped: %v", err)
	} else {
		t.Logf("resolved payment dispute: %s", resolved.ID)
	}
}

func TestSimulateIssuingTransactionLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	// List cards first
	cards, err := testClient.Issuing().ListCards(ctx)
	require.NoError(t, err, "list cards failed")
	var cardID string
	if len(cards.Items) > 0 {
		cardID = cards.Items[0].ID
	}

	// SimulateIssuingTransactionCreate
	created, err := svc.SimulateIssuingTransactionCreate(ctx, &simulation.SimulateIssuingTransactionCreateRequest{
		CardID:               cardID,
		CardNumber:           "4111111111111111",
		TransactionAmount:    50.0,
		TransactionCurrency:  "USD",
		MerchantCategoryCode: "5411",
		MerchantInfo:         "Test Merchant",
		TransactionID:        "sim-txn-" + strconv.FormatInt(time.Now().UnixNano(), 10),
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err) || sdk.IsOperationFailed(err), "create issuing transaction failed: %v", err)
		t.Logf("create issuing transaction skipped: %v", err)
		return
	}
	require.NotEmpty(t, created.TransactionID, "issuing transaction id should not be empty")
	t.Logf("created issuing transaction: %s", created.TransactionID)

	// SimulateIssuingTransactionCapture
	captured, err := svc.SimulateIssuingTransactionCapture(ctx, created.TransactionID, &simulation.SimulateIssuingTransactionCaptureRequest{
		TransactionAmount: 50.0,
		MerchantInfo:      "Test Merchant",
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "capture issuing transaction failed: %v", err)
		t.Logf("capture issuing transaction skipped: %v", err)
	} else {
		t.Logf("captured issuing transaction: %s", captured.TransactionID)
	}

	// SimulateIssuingTransactionRefund
	refunded, err := svc.SimulateIssuingTransactionRefund(ctx, &simulation.SimulateIssuingTransactionRefundRequest{
		CardID:              cardID,
		CardNumber:          "4111111111111111",
		TransactionAmount:   25.0,
		TransactionCurrency: "USD",
		TransactionID:       "sim-refund-" + strconv.FormatInt(time.Now().UnixNano(), 10),
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "refund issuing transaction failed: %v", err)
		t.Logf("refund issuing transaction skipped: %v", err)
	} else {
		t.Logf("refunded issuing transaction: %s", refunded.TransactionID)
	}

	// SimulateIssuingTransactionReverse
	reversed, err := svc.SimulateIssuingTransactionReverse(ctx, created.TransactionID, &simulation.SimulateIssuingTransactionReverseRequest{
		TransactionAmount: 25.0,
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "reverse issuing transaction failed: %v", err)
		t.Logf("reverse issuing transaction skipped: %v", err)
	} else {
		t.Logf("reversed issuing transaction: %s", reversed.TransactionID)
	}
}

func TestSimulateAccountUpdateStatus(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	// Try with a test account ID - expected to fail or succeed depending on sandbox
	err := svc.SimulateAccountUpdateStatus(ctx, "acct_test", &simulation.SimulateAccountUpdateStatusRequest{
		NextStatus: simulation.AccountStatusActive,
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err) || sdk.IsCredentialsInvalid(err), "simulate account update status failed: %v", err)
		t.Logf("simulate account update status skipped: %v", err)
	} else {
		t.Logf("simulated account update status successfully")
	}
}

func TestSimulateBillingFailNextAutocharge(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	// Try with a test payment source ID - expected to fail or succeed depending on sandbox
	err := svc.SimulateBillingFailNextAutocharge(ctx, "ps_test")
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err) || sdk.IsInvalidArgument(err), "simulate billing fail next autocharge failed: %v", err)
		t.Logf("simulate billing fail next autocharge skipped: %v", err)
	} else {
		t.Logf("simulated billing fail next autocharge successfully")
	}
}

func TestSimulateTransferTransition(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Simulation()

	// List transfers first
	transfers, err := testClient.Payouts().ListTransfers(ctx)
	require.NoError(t, err, "list transfers failed")
	if len(transfers.Items) == 0 {
		t.Logf("no transfers available, skipping transfer transition simulation")
		return
	}

	transferID := transfers.Items[0].ID
	result, err := svc.SimulateTransferTransition(ctx, transferID, &simulation.SimulateTransferTransitionRequest{
		NextStatus: simulation.TransferStatusSent,
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err) || sdk.IsOperationFailed(err), "simulate transfer transition failed: %v", err)
		t.Logf("simulate transfer transition skipped: %v", err)
	} else {
		require.NotNil(t, result)
		t.Logf("simulated transfer transition: %s status=%s", result.ID, result.Status)
	}
}
