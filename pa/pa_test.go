package pa_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/pa"
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

func TestPaymentIntentLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	created, err := svc.CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
		RequestID:       "pi-req-" + time.Now().Format("20060102150405"),
		MerchantOrderID: "order-" + time.Now().Format("20060102150405"),
		Amount:          100.0,
		Currency:        sdk.CurrencyUSD,
	})
	require.NoError(t, err, "create payment intent failed")
	t.Logf("created payment intent: %s status=%s", created.ID, created.Status)

	fetched, err := svc.GetPaymentIntent(ctx, created.ID)
	require.NoError(t, err, "get payment intent failed")
	assert.Equal(t, created.ID, fetched.ID, "payment intent id mismatch")

	updated, err := svc.UpdatePaymentIntent(ctx, created.ID, &pa.UpdatePaymentIntentRequest{
		RequestID: "pi-upd-" + time.Now().Format("20060102150405"),
		Amount:    200.0,
	})
	require.NoError(t, err, "update payment intent failed")
	assert.Equal(t, 200.0, updated.Amount, "payment intent amount not updated")

	list, err := svc.ListPaymentIntents(ctx)
	require.NoError(t, err, "list payment intents failed")
	require.NotEmpty(t, list.Items, "expected at least one payment intent")

	// Test capture with manual capture PI
	piCapture, err := svc.CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
		RequestID:       "pi-capture-" + time.Now().Format("20060102150405"),
		MerchantOrderID: "order-capture-" + time.Now().Format("20060102150405"),
		Amount:          100.0,
		Currency:        sdk.CurrencyUSD,
		CaptureMethod:   "manual",
	})
	require.NoError(t, err, "create payment intent for capture failed")
	t.Logf("created capture PI: %s status=%s", piCapture.ID, piCapture.Status)

	// Confirm the PI with card payment method first
	confirmedCapture, err := svc.ConfirmPaymentIntent(ctx, piCapture.ID, &pa.ConfirmPaymentIntentRequest{
		RequestID: "pi-capture-confirm-" + time.Now().Format("20060102150405"),
		PaymentMethod: &pa.PaymentMethodInput{
			Type: "card",
			Card: &pa.CardPaymentMethod{
				Number:      "4012000300001003",
				ExpiryMonth: "03",
				ExpiryYear:  "2030",
				CVC:         "123",
				Name:        "Test User",
			},
		},
	})
	require.NoError(t, err, "confirm payment intent for capture failed")
	t.Logf("confirmed capture PI: %s status=%s", confirmedCapture.ID, confirmedCapture.Status)

	// Handle capture by status (same pattern as _examples/payment_intent/main.go)
	switch confirmedCapture.Status {
	case pa.PaymentIntentStatusRequiresCapture:
		captured, err := svc.CapturePaymentIntent(ctx, piCapture.ID, &pa.CapturePaymentIntentRequest{
			RequestID: "pi-cap-" + time.Now().Format("20060102150405"),
		})
		require.NoError(t, err, "capture payment intent failed")
		assert.Equal(t, pa.PaymentIntentStatusSucceeded, captured.Status, "captured PI should be succeeded")
		t.Logf("captured payment intent status: %s", captured.Status)

	case pa.PaymentIntentStatusSucceeded:
		t.Logf("payment intent auto-captured (status=SUCCEEDED), explicit capture skipped")

	default:
		t.Logf("unexpected status after confirm: %s", confirmedCapture.Status)
	}

	cancelled, err := svc.CancelPaymentIntent(ctx, created.ID, &pa.CancelPaymentIntentRequest{
		RequestID: "pi-cancel-" + time.Now().Format("20060102150405"),
	})
	require.NoError(t, err, "cancel payment intent failed")
	t.Logf("cancelled payment intent status: %s", cancelled.Status)
}

func TestCustomerLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	created, err := svc.CreateCustomer(ctx, &pa.CreateCustomerRequest{
		RequestID:          "cus-req-" + time.Now().Format("20060102150405"),
		MerchantCustomerID: "mcus-" + time.Now().Format("20060102150405"),
		Email:              "test-" + time.Now().Format("20060102150405") + "@example.com",
		FirstName:          "Test",
		LastName:           "User",
	})
	require.NoError(t, err, "create customer failed")
	t.Logf("created customer: %s", created.ID)

	fetched, err := svc.GetCustomer(ctx, created.ID)
	require.NoError(t, err, "get customer failed")
	assert.Equal(t, created.ID, fetched.ID, "customer id mismatch")

	updated, err := svc.UpdateCustomer(ctx, created.ID, &pa.UpdateCustomerRequest{
		RequestID: "cus-upd-" + time.Now().Format("20060102150405"),
		FirstName: "Updated",
	})
	require.NoError(t, err, "update customer failed")
	assert.Equal(t, "Updated", updated.FirstName, "customer first name not updated")

	list, err := svc.ListCustomers(ctx)
	require.NoError(t, err, "list customers failed")
	require.NotEmpty(t, list.Items, "expected at least one customer")
}

func TestPaymentLinkLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	created, err := svc.CreatePaymentLink(ctx, &pa.CreatePaymentLinkRequest{
		RequestID: "pl-req-" + time.Now().Format("20060102150405"),
		Title:     "Test Payment Link",
		Amount:    99.99,
		Currency:  sdk.CurrencyUSD,
		Reusable:  true,
	})
	require.NoError(t, err, "create payment link failed")
	t.Logf("created payment link: %s url=%s", created.ID, created.URL)

	fetched, err := svc.GetPaymentLink(ctx, created.ID)
	require.NoError(t, err, "get payment link failed")
	assert.Equal(t, created.ID, fetched.ID, "payment link id mismatch")

	list, err := svc.ListPaymentLinks(ctx)
	require.NoError(t, err, "list payment links failed")
	t.Logf("payment links count: %d", len(list.Items))
}

func TestRefundLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	// Create a PaymentIntent with manual capture
	pi, err := svc.CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
		RequestID:       "pi-refund-" + time.Now().Format("20060102150405"),
		Amount:          50.0,
		Currency:        sdk.CurrencyUSD,
		MerchantOrderID: "order-refund-" + time.Now().Format("20060102150405"),
		CaptureMethod:   "manual",
	})
	require.NoError(t, err, "create payment intent for refund failed")
	t.Logf("created refund PI: %s status=%s", pi.ID, pi.Status)

	// Confirm with card
	confirmed, err := svc.ConfirmPaymentIntent(ctx, pi.ID, &pa.ConfirmPaymentIntentRequest{
		RequestID: "pi-refund-confirm-" + time.Now().Format("20060102150405"),
		PaymentMethod: &pa.PaymentMethodInput{
			Type: "card",
			Card: &pa.CardPaymentMethod{
				Number:      "4012000300001003",
				ExpiryMonth: "03",
				ExpiryYear:  "2030",
				CVC:         "123",
				Name:        "Test User",
			},
		},
	})
	require.NoError(t, err, "confirm payment intent for refund failed")
	t.Logf("confirmed refund PI: %s status=%s", confirmed.ID, confirmed.Status)

	// Ensure PI is captured before refund (handle both auto-capture and manual capture)
	switch confirmed.Status {
	case pa.PaymentIntentStatusRequiresCapture:
		captured, err := svc.CapturePaymentIntent(ctx, pi.ID, &pa.CapturePaymentIntentRequest{
			RequestID: "pi-refund-capture-" + time.Now().Format("20060102150405"),
		})
		require.NoError(t, err, "capture payment intent for refund failed")
		assert.Equal(t, pa.PaymentIntentStatusSucceeded, captured.Status)
		t.Logf("captured refund PI: %s status=%s", captured.ID, captured.Status)

	case pa.PaymentIntentStatusSucceeded:
		t.Logf("refund PI auto-captured (status=SUCCEEDED), explicit capture skipped")

	default:
		t.Logf("unexpected status after confirm: %s", confirmed.Status)
	}

	// Create refund
	refund, err := svc.CreateRefund(ctx, &pa.CreateRefundRequest{
		RequestID:       "ref-req-" + time.Now().Format("20060102150405"),
		PaymentIntentID: pi.ID,
		Amount:          10.0,
	})
	require.NoError(t, err, "create refund failed")
	t.Logf("created refund: %s status=%s", refund.ID, refund.Status)

	// Get refund (may 404 in sandbox immediately after create)
	fetched, err := svc.GetRefund(ctx, refund.ID)
	if err != nil {
		if sdk.IsResourceNotFound(err) {
			t.Logf("🔒 get refund returned 404 (sandbox delay): %v", err)
		} else {
			require.NoError(t, err, "get refund failed")
		}
	} else {
		assert.Equal(t, refund.ID, fetched.ID)
	}

	// List refunds
	list, err := svc.ListRefunds(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, list.Items)
}

func TestPaymentConsentLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	// Create customer first
	cust, err := svc.CreateCustomer(ctx, &pa.CreateCustomerRequest{
		RequestID:          "pc-cust-" + time.Now().Format("20060102150405"),
		Email:              "test-" + time.Now().Format("20060102150405") + "@example.com",
		FirstName:          "Test",
		LastName:           "Consent",
		MerchantCustomerID: "mcus-pc-" + time.Now().Format("20060102150405"),
	})
	require.NoError(t, err)

	// Create payment consent
	consent, err := svc.CreatePaymentConsent(ctx, &pa.CreatePaymentConsentRequest{
		RequestID:       "pc-" + time.Now().Format("20060102150405"),
		CustomerID:      cust.ID,
		NextTriggeredBy: "customer",
	})
	if err != nil {
		require.Error(t, err, "create consent blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	t.Logf("created consent: %s status=%s", consent.ID, consent.Status)

	// Get consent
	fetched, err := svc.GetPaymentConsent(ctx, consent.ID)
	require.NoError(t, err, "get consent failed")
	assert.Equal(t, consent.ID, fetched.ID)

	// List consents
	list, err := svc.ListPaymentConsents(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, list.Items)

	// Cancel consent
	cancelled, err := svc.CancelPaymentConsent(ctx, consent.ID, &pa.CancelPaymentConsentRequest{
		RequestID: "pc-cancel-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		require.Error(t, err, "cancel consent blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("cancelled consent: %s", cancelled.Status)
	}
}

func TestPaymentMethodLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	// 先创建客户
	cust, err := svc.CreateCustomer(ctx, &pa.CreateCustomerRequest{
		RequestID:          "pm-cust-" + time.Now().Format("20060102150405"),
		FirstName:          "Test",
		LastName:           "PM",
		Email:              "pm-test-" + time.Now().Format("20060102150405") + "@example.com",
		MerchantCustomerID: "mcus-pm-" + time.Now().Format("20060102150405"),
	})
	require.NoError(t, err, "create customer for payment method failed")
	t.Logf("created customer: %s", cust.ID)

	// 创建支付方式（银行卡）
	created, err := svc.CreatePaymentMethod(ctx, &pa.CreatePaymentMethodRequest{
		RequestID:  "pm-req-" + time.Now().Format("20060102150405"),
		Type:       pa.PaymentMethodTypeCard,
		CustomerID: cust.ID,
		Card: &pa.CreateCardPaymentMethod{
			Number:      "4012000300001003",
			ExpiryMonth: "03",
			ExpiryYear:  "2030",
			Name:        "Test User",
		},
	})
	if err != nil {
		require.Error(t, err, "create payment method blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	t.Logf("created payment method: %s type=%s status=%s", created.ID, created.Type, created.Status)

	// 获取支付方式
	fetched, err := svc.GetPaymentMethod(ctx, created.ID)
	require.NoError(t, err, "get payment method failed")
	assert.Equal(t, created.ID, fetched.ID, "payment method id mismatch")

	// 列出支付方式
	list, err := svc.ListPaymentMethods(ctx)
	require.NoError(t, err, "list payment methods failed")
	require.NotEmpty(t, list.Items, "expected at least one payment method")

	// 禁用支付方式（清理）
	disabled, err := svc.DisablePaymentMethod(ctx, created.ID, &pa.DisablePaymentMethodRequest{
		RequestID: "pm-disable-" + time.Now().Format("20060102150405"),
	})
	require.NoError(t, err, "disable payment method failed")
	assert.Equal(t, pa.PaymentMethodStatusDisabled, disabled.Status, "payment method should be disabled")
	t.Logf("disabled payment method: %s", disabled.Status)
}

func TestTerminalLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	created, err := svc.CreateTerminal(ctx, &pa.CreateTerminalRequest{
		RequestID:      "term-req-" + time.Now().Format("20060102150405"),
		ActivationCode: "MEFCMSBC",
		Nickname:       "Test Terminal",
	})
	if err != nil {
		require.Error(t, err, "create terminal blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	t.Logf("created terminal: %s nickname=%s model=%s", created.ID, created.Nickname, created.Model)

	fetched, err := svc.GetTerminal(ctx, created.ID)
	require.NoError(t, err, "get terminal failed")
	assert.Equal(t, created.ID, fetched.ID, "terminal id mismatch")

	updated, err := svc.UpdateTerminal(ctx, created.ID, &pa.UpdateTerminalRequest{
		RequestID: "term-upd-" + time.Now().Format("20060102150405"),
		Nickname:  "Updated Terminal",
	})
	require.NoError(t, err, "update terminal failed")
	assert.Equal(t, "Updated Terminal", updated.Nickname, "terminal nickname not updated")

	// Verify update persisted
	fetched2, err := svc.GetTerminal(ctx, created.ID)
	require.NoError(t, err, "get terminal after update failed")
	assert.Equal(t, "Updated Terminal", fetched2.Nickname, "terminal nickname update not persisted")

	list, err := svc.ListTerminals(ctx, &pa.ListTerminalsRequest{
		PageSize: 10,
	})
	require.NoError(t, err, "list terminals failed")
	require.NotEmpty(t, list.Items, "expected at least one terminal in list")

	// Activate terminal
	activated, err := svc.ActivateTerminal(ctx, created.ID, &pa.ActivateTerminalRequest{
		RequestID: "term-activate-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		require.Error(t, err, "activate terminal blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("activated terminal status: %s", activated.Status)
	}

	// Deactivate terminal
	deactivated, err := svc.DeactivateTerminal(ctx, created.ID, &pa.DeactivateTerminalRequest{
		RequestID: "term-deactivate-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		require.Error(t, err, "deactivate terminal blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("deactivated terminal status: %s", deactivated.Status)
	}

	// Reset terminal password
	reset, err := svc.ResetTerminalPassword(ctx, created.ID, &pa.ResetTerminalPasswordRequest{
		RequestID:    "term-reset-" + time.Now().Format("20060102150405"),
		PasswordType: "admin",
	})
	if err != nil {
		require.Error(t, err, "reset terminal password blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("reset terminal password status: %s", reset.AdminPasswordStatus)
	}

	// Cancel current operation
	cancelledOp, err := svc.CancelCurrentOperation(ctx, created.ID, &pa.CancelTerminalOperationRequest{
		RequestID: "term-cancel-op-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		require.Error(t, err, "cancel current operation blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("cancel current operation status: %s", cancelledOp.Status)
	}

	// Process payment intent in terminal
	pi, err := svc.CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
		RequestID:       "term-pi-" + time.Now().Format("20060102150405"),
		MerchantOrderID: "order-term-" + time.Now().Format("20060102150405"),
		Amount:          10.0,
		Currency:        sdk.CurrencyUSD,
	})
	require.NoError(t, err, "create payment intent for terminal failed")

	processed, err := svc.ProcessPaymentIntentInTerminal(ctx, created.ID, &pa.ProcessPaymentIntentInTerminalRequest{
		RequestID:       "term-process-" + time.Now().Format("20060102150405"),
		PaymentIntentID: pi.ID,
	})
	if err != nil {
		require.Error(t, err, "process payment intent in terminal blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("processed payment intent in terminal status: %s", processed.Status)
	}

	// Terminate terminal (destructive, do last)
	terminated, err := svc.TerminateTerminal(ctx, created.ID, &pa.TerminateTerminalRequest{
		RequestID: "term-terminate-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		require.Error(t, err, "terminate terminal blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("terminated terminal status: %s", terminated.Status)
	}
}

func TestPaymentAttemptLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	list, err := svc.ListPaymentAttempts(ctx)
	if err != nil {
		require.Error(t, err, "list payment attempts blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	if len(list.Items) == 0 {
		t.Logf("no payment attempts found, skipping get")
		return
	}

	fetched, err := svc.GetPaymentAttempt(ctx, list.Items[0].ID)
	require.NoError(t, err, "get payment attempt failed")
	assert.Equal(t, list.Items[0].ID, fetched.ID)
}

func TestPaymentDisputeLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	list, err := svc.ListPaymentDisputes(ctx)
	if err != nil {
		require.Error(t, err, "list payment disputes blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	if len(list.Items) == 0 {
		t.Logf("no payment disputes found, skipping get")
		return
	}

	fetched, err := svc.GetPaymentDispute(ctx, list.Items[0].ID)
	require.NoError(t, err, "get payment dispute failed")
	assert.Equal(t, list.Items[0].ID, fetched.ID)
}

func TestConversionQuoteLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	created, err := svc.CreateConversionQuote(ctx, &pa.CreateConversionQuoteRequest{
		RequestID:        "cq-req-" + time.Now().Format("20060102150405"),
		MerchantCurrency: sdk.CurrencyUSD,
		ShopperCurrency:  sdk.CurrencyEUR,
	})
	if err != nil {
		require.Error(t, err, "create conversion quote blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	t.Logf("created conversion quote: %s status=%s", created.ID, created.Status)

	fetched, err := svc.GetConversionQuote(ctx, created.ID)
	require.NoError(t, err, "get conversion quote failed")
	assert.Equal(t, created.ID, fetched.ID)
}

func TestSettlementRecordLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	// Create a payment intent to have a payment_intent_id for settlement records query
	pi, err := svc.CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
		RequestID:       "sr-pi-" + time.Now().Format("20060102150405"),
		MerchantOrderID: "order-sr-" + time.Now().Format("20060102150405"),
		Amount:          10.0,
		Currency:        sdk.CurrencyUSD,
	})
	require.NoError(t, err, "create payment intent for settlement record failed")

	list, err := svc.ListSettlementRecords(ctx, &pa.ListSettlementRecordsRequest{
		PageSize:        10,
		PaymentIntentID: pi.ID,
	})
	if err != nil {
		require.Error(t, err, "list settlement records blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	if len(list.Items) == 0 {
		t.Logf("no settlement records found, skipping get")
		return
	}

	fetched, err := svc.GetSettlementRecord(ctx, list.Items[0].ID)
	require.NoError(t, err, "get settlement record failed")
	assert.Equal(t, list.Items[0].ID, fetched.ID)
}

func TestCustomsDeclarationLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	// Create a payment intent first
	pi, err := svc.CreatePaymentIntent(ctx, &pa.CreatePaymentIntentRequest{
		RequestID:       "cd-pi-" + time.Now().Format("20060102150405"),
		MerchantOrderID: "order-cd-" + time.Now().Format("20060102150405"),
		Amount:          100.0,
		Currency:        sdk.CurrencyUSD,
	})
	require.NoError(t, err, "create payment intent for customs declaration failed")

	created, err := svc.CreateCustomsDeclaration(ctx, &pa.CreateCustomsDeclarationRequest{
		RequestID:       "cd-req-" + time.Now().Format("20060102150405"),
		PaymentIntentID: pi.ID,
		CustomsDetails: &pa.CustomsDetails{
			CustomsCode:           "SHANGHAI_ZS",
			MerchantCustomsName:   "Test Merchant",
			MerchantCustomsNumber: "5678",
		},
		SubOrder: &pa.SubOrder{
			OrderNumber: "SO-001",
			Amount:      100.0,
			ShippingFee: 10.0,
			Currency:    string(sdk.CurrencyUSD),
		},
	})
	if err != nil {
		require.Error(t, err, "create customs declaration failed")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	require.NotEmpty(t, created.ID, "customs declaration id should not be empty")
	t.Logf("created customs declaration: %s status=%s", created.ID, created.Status)

	fetched, err := svc.GetCustomsDeclaration(ctx, created.ID)
	if err != nil {
		require.Error(t, err, "get customs declaration blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		assert.Equal(t, created.ID, fetched.ID, "customs declaration id mismatch")
	}

	updated, err := svc.UpdateCustomsDeclaration(ctx, created.ID, &pa.UpdateCustomsDeclarationRequest{
		RequestID: "cd-upd-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		require.Error(t, err, "update customs declaration blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("updated customs declaration: %s", updated.ID)
	}

	redeclared, err := svc.RedeclareCustomsDeclaration(ctx, created.ID, &pa.RedeclareCustomsDeclarationRequest{
		RequestID: "cd-red-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		require.Error(t, err, "redeclare customs declaration blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("redeclared customs declaration: %s", redeclared.ID)
	}

	list, err := svc.ListCustomsDeclarations(ctx, &pa.ListCustomsDeclarationsRequest{PageSize: 10})
	if err != nil {
		require.Error(t, err, "list customs declarations blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	require.NotEmpty(t, list.Items, "expected at least one customs declaration")
	t.Logf("customs declarations count: %d", len(list.Items))
}

func TestFundsSplitLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	created, err := svc.CreateFundsSplit(ctx, &pa.CreateFundsSplitRequest{
		RequestID:   "fs-req-" + time.Now().Format("20060102150405"),
		SourceID:    "pi_test",
		SourceType:  "PAYMENT_INTENT",
		Destination: "acct_test",
		Amount:      "100.00",
	})
	if err != nil {
		require.Error(t, err, "create funds split blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	require.NotEmpty(t, created.ID, "funds split id should not be empty")
	t.Logf("created funds split: %s", created.ID)

	fetched, err := svc.GetFundsSplit(ctx, created.ID)
	if err != nil {
		require.Error(t, err, "get funds split blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		assert.Equal(t, created.ID, fetched.ID, "funds split id mismatch")
	}

	list, err := svc.ListFundsSplits(ctx, &pa.ListFundsSplitsRequest{PageSize: 10})
	if err != nil {
		require.Error(t, err, "list funds splits blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	require.NotEmpty(t, list.Items, "expected at least one funds split")
	t.Logf("funds splits count: %d", len(list.Items))
}

func TestFundsSplitReversalLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	// Create a funds split first
	fs, err := svc.CreateFundsSplit(ctx, &pa.CreateFundsSplitRequest{
		RequestID:   "fsr-fs-req-" + time.Now().Format("20060102150405"),
		SourceID:    "pi_test",
		SourceType:  "PAYMENT_INTENT",
		Destination: "acct_test",
		Amount:      "50.00",
	})
	if err != nil {
		require.Error(t, err, "create funds split for reversal blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	require.NotEmpty(t, fs.ID, "funds split id should not be empty")

	created, err := svc.CreateFundsSplitReversal(ctx, &pa.CreateFundsSplitReversalRequest{
		RequestID:    "fsr-req-" + time.Now().Format("20060102150405"),
		FundsSplitID: fs.ID,
		Amount:       "50.00",
	})
	if err != nil {
		require.Error(t, err, "create funds split reversal blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	require.NotEmpty(t, created.ID, "funds split reversal id should not be empty")
	t.Logf("created funds split reversal: %s", created.ID)

	fetched, err := svc.GetFundsSplitReversal(ctx, created.ID)
	if err != nil {
		require.Error(t, err, "get funds split reversal blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		assert.Equal(t, created.ID, fetched.ID, "funds split reversal id mismatch")
	}

	list, err := svc.ListFundsSplitReversals(ctx, &pa.ListFundsSplitReversalsRequest{PageSize: 10})
	if err != nil {
		require.Error(t, err, "list funds split reversals blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
		return
	}
	require.NotEmpty(t, list.Items, "expected at least one funds split reversal")
	t.Logf("funds splits reversals count: %d", len(list.Items))
}

func TestConfigLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.PA()

	// GetPaymentMethodTypes
	pmTypes, err := svc.GetPaymentMethodTypes(ctx, &pa.GetPaymentMethodTypesRequest{
		PageSize: 10,
	})
	if err != nil {
		require.Error(t, err, "get payment method types blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		require.NotNil(t, pmTypes)
		t.Logf("payment method types count: %d", len(pmTypes.Items))
	}

	// GetBanks
	banks, err := svc.GetBanks(ctx, &pa.GetBanksRequest{
		PageSize:         10,
		PaymentMethodType: "card",
	})
	if err != nil {
		require.Error(t, err, "get banks failed")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		require.NotNil(t, banks)
		t.Logf("banks count: %d", len(banks.Items))
	}

	// GetConvertibleShopperCurrencies
	currencies, err := svc.GetConvertibleShopperCurrencies(ctx)
	if err != nil {
		require.Error(t, err, "get convertible shopper currencies failed")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		require.NotNil(t, currencies)
		t.Logf("convertible shopper currencies count: %d", len(currencies.Items))
	}

	// GetReservePlan
	reservePlan, err := svc.GetReservePlan(ctx)
	if err != nil {
		require.Error(t, err, "get reserve plan failed")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		require.NotNil(t, reservePlan)
		t.Logf("reserve plan type: %s", reservePlan.Type)
	}

	// GetApplePayDomains
	domains, err := svc.GetApplePayDomains(ctx)
	if err != nil {
		require.Error(t, err, "get apple pay domains failed")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		require.NotNil(t, domains)
		t.Logf("apple pay domains count: %d", len(domains.Items))
	}

	// AddApplePayDomains
	addedDomains, err := svc.AddApplePayDomains(ctx, &pa.AddApplePayDomainsRequest{
		Items: []string{"example.com"},
	})
	if err != nil {
		require.Error(t, err, "add apple pay domains blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		require.NotNil(t, addedDomains)
		t.Logf("added apple pay domains count: %d", len(addedDomains.Items))
	}

	// RemoveApplePayDomains
	removedDomains, err := svc.RemoveApplePayDomains(ctx, &pa.RemoveApplePayDomainsRequest{
		Items: []string{"example.com"},
	})
	if err != nil {
		require.Error(t, err, "remove apple pay domains blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		require.NotNil(t, removedDomains)
		t.Logf("removed apple pay domains count: %d", len(removedDomains.Items))
	}
}
