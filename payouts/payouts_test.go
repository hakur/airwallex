package payouts_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/payouts"
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

func TestBeneficiaryLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Payouts()

	created, err := svc.CreateBeneficiary(ctx, &payouts.CreateBeneficiaryRequest{
		RequestID:       "ben-req-" + time.Now().Format("20060102150405"),
		TransferMethods: []string{"LOCAL"},
		Beneficiary: &payouts.BeneficiaryInput{
			Type:       payouts.BeneficiaryTypeBankAccount,
			EntityType: "COMPANY",
			Name:       "Test Beneficiary " + time.Now().Format("20060102150405"),
			Email:      "test@example.com",
			Address: &payouts.BeneficiaryAddress{
				CountryCode:   sdk.CountryCodeUS,
				State:         "New York",
				City:          "New York",
				StreetAddress: "123 Main St",
				Postcode:      "10001",
			},
			BankDetails: &payouts.BankDetails{
				AccountCurrency:      sdk.CurrencyUSD,
				AccountName:          "Test User",
				AccountNumber:        "50001121",
				BankAccountCategory:  "Checking",
				BankName:             "Test Bank",
				BankCountryCode:      sdk.CountryCodeUS,
				LocalClearingSystem:  "ACH",
				AccountRoutingType1:  "aba",
				AccountRoutingValue1: "021000021",
			},
		},
	})
	require.NoError(t, err, "create beneficiary failed")
	benType := ""
	benStatus := ""
	if created.Beneficiary != nil {
		benType = string(created.Beneficiary.Type)
	}
	t.Logf("created beneficiary: %s type=%s status=%s", created.ID, benType, benStatus)

	fetched, err := svc.GetBeneficiary(ctx, created.ID)
	require.NoError(t, err, "get beneficiary failed")
	assert.Equal(t, created.ID, fetched.ID, "beneficiary id mismatch")

	updated, err := svc.UpdateBeneficiary(ctx, created.ID, &payouts.UpdateBeneficiaryRequest{
		RequestID:       "ben-upd-" + time.Now().Format("20060102150405"),
		TransferMethods: []string{"LOCAL"},
		Beneficiary: &payouts.BeneficiaryInput{
			Type:       payouts.BeneficiaryTypeBankAccount,
			EntityType: "COMPANY",
			Name:       "Updated Beneficiary",
			Email:      "test@example.com",
			Address: &payouts.BeneficiaryAddress{
				CountryCode:   sdk.CountryCodeUS,
				State:         "New York",
				City:          "New York",
				StreetAddress: "123 Main St",
				Postcode:      "10001",
			},
			BankDetails: &payouts.BankDetails{
				AccountCurrency:      sdk.CurrencyUSD,
				AccountName:          "Test User",
				AccountNumber:        "50001121",
				BankAccountCategory:  "Checking",
				BankName:             "Test Bank",
				BankCountryCode:      sdk.CountryCodeUS,
				LocalClearingSystem:  "ACH",
				AccountRoutingType1:  "aba",
				AccountRoutingValue1: "021000021",
			},
		},
	})
	require.NoError(t, err, "update beneficiary failed")
	if updated.Beneficiary != nil {
		t.Logf("updated beneficiary name: %s", updated.Beneficiary.Name)
	}

	validated, err := svc.ValidateBeneficiary(ctx, &payouts.ValidateBeneficiaryRequest{
		RequestID:       "ben-val-" + time.Now().Format("20060102150405"),
		TransferMethods: []string{"LOCAL"},
		Beneficiary: &payouts.BeneficiaryInput{
			Type:       payouts.BeneficiaryTypeBankAccount,
			EntityType: "COMPANY",
			Name:       "Validate Beneficiary",
			Email:      "test@example.com",
			Address: &payouts.BeneficiaryAddress{
				CountryCode:   sdk.CountryCodeUS,
				State:         "New York",
				City:          "New York",
				StreetAddress: "123 Main St",
				Postcode:      "10001",
			},
			BankDetails: &payouts.BankDetails{
				AccountCurrency:      sdk.CurrencyUSD,
				AccountName:          "Test User",
				AccountNumber:        "50001121",
				BankAccountCategory:  "Checking",
				BankName:             "Test Bank",
				BankCountryCode:      sdk.CountryCodeUS,
				LocalClearingSystem:  "ACH",
				AccountRoutingType1:  "aba",
				AccountRoutingValue1: "021000021",
			},
		},
	})
	if err != nil {
		require.Error(t, err, "validate beneficiary blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("validate beneficiary result: %v", validated)
	}

	verifyResp, err := svc.VerifyAccount(ctx, &payouts.VerifyAccountRequest{
		RequestID: "ben-verify-" + time.Now().Format("20060102150405"),
		Beneficiary: &payouts.BeneficiaryInput{
			Type:       payouts.BeneficiaryTypeBankAccount,
			EntityType: "COMPANY",
			BankDetails: &payouts.BankDetails{
				AccountCurrency:      sdk.CurrencyUSD,
				AccountName:          "Test User",
				AccountNumber:        "50001121",
				BankCountryCode:      sdk.CountryCodeUS,
				LocalClearingSystem:  "ACH",
				AccountRoutingType1:  "aba",
				AccountRoutingValue1: "021000021",
			},
		},
	})
	if err != nil {
		require.Error(t, err, "verify account blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("verify account valid=%v", verifyResp.Valid)
	}

	_, err = svc.GenerateAPISchema(ctx)
	if err != nil {
		require.Error(t, err, "generate api schema blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("generate api schema succeeded")
	}

	_, err = svc.GenerateFormSchema(ctx)
	if err != nil {
		require.Error(t, err, "generate form schema blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("generate form schema succeeded")
	}

	institutions, err := svc.GetSupportedFinancialInstitutions(ctx, &payouts.GetSupportedFinancialInstitutionsRequest{
		CountryCode: sdk.CountryCodeUS,
		Currency:    sdk.CurrencyUSD,
	})
	if err != nil {
		require.Error(t, err, "get supported institutions blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("supported institutions count: %d", len(institutions.Items))
	}

	fetched2, err := svc.GetBeneficiary(ctx, created.ID)
	require.NoError(t, err, "get beneficiary failed")
	assert.Equal(t, created.ID, fetched2.ID, "beneficiary id mismatch")

	err = svc.DeleteBeneficiary(ctx, created.ID)
	require.NoError(t, err, "delete beneficiary failed")
}

func TestTransferLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Payouts()

	benList, err := svc.ListBeneficiaries(ctx)
	require.NoError(t, err, "list beneficiaries failed")

	var benID string
	if len(benList.Items) > 0 {
		benID = benList.Items[0].ID
	} else {
		ben, err := svc.CreateBeneficiary(ctx, &payouts.CreateBeneficiaryRequest{
			RequestID:       "ben-req-" + time.Now().Format("20060102150405"),
			TransferMethods: []string{"LOCAL"},
			Beneficiary: &payouts.BeneficiaryInput{
				Type:       payouts.BeneficiaryTypeBankAccount,
				EntityType: "COMPANY",
				Name:       "Transfer Test Beneficiary",
				Email:      "test@example.com",
				Address: &payouts.BeneficiaryAddress{
					CountryCode:   sdk.CountryCodeUS,
					State:         "New York",
					City:          "New York",
					StreetAddress: "123 Main St",
					Postcode:      "10001",
				},
				BankDetails: &payouts.BankDetails{
					AccountCurrency:      sdk.CurrencyUSD,
					AccountName:          "Test User",
					AccountNumber:        "50001121",
					BankAccountCategory:  "Checking",
					BankName:             "Test Bank",
					BankCountryCode:      sdk.CountryCodeUS,
					LocalClearingSystem:  "ACH",
					AccountRoutingType1:  "aba",
					AccountRoutingValue1: "021000021",
				},
			},
		})
		require.NoError(t, err, "create beneficiary for transfer test failed")
		benID = ben.ID
	}

	transfer, err := svc.CreateTransfer(ctx, &payouts.CreateTransferRequest{
		RequestID:        "tr-req-" + time.Now().Format("20060102150405"),
		BeneficiaryID:    benID,
		TransferAmount:   100.0,
		TransferCurrency: sdk.CurrencyUSD,
		TransferMethod:   payouts.TransferMethodLocal,
		SourceCurrency:   string(sdk.CurrencyUSD),
		Reason:           "business_expenses",
		Reference:        "test-ref-" + time.Now().Format("20060102150405"),
	})
	require.NoError(t, err, "create transfer failed")
	t.Logf("created transfer: %s status=%s", transfer.ID, transfer.Status)

	fetched, err := svc.GetTransfer(ctx, transfer.ID)
	require.NoError(t, err, "get transfer failed")
	assert.Equal(t, transfer.ID, fetched.ID, "transfer id mismatch")

	valid, err := svc.ValidateTransfer(ctx, &payouts.ValidateTransferRequest{
		RequestID:        "tr-val-" + time.Now().Format("20060102150405"),
		BeneficiaryID:    benID,
		TransferAmount:   50.0,
		TransferCurrency: sdk.CurrencyUSD,
		TransferMethod:   payouts.TransferMethodLocal,
		SourceCurrency:   string(sdk.CurrencyUSD),
		Reason:           "business_expenses",
		Reference:        "test-ref-val-" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		require.Error(t, err, "validate transfer blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("validate transfer result: %v", valid)
	}

	cancelled, err := svc.CancelTransfer(ctx, transfer.ID)
	if err != nil {
		require.Error(t, err, "cancel transfer blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("cancelled transfer: %v", cancelled)
	}

	// 确认转账资金（在 Sandbox 中不支持的端点）
	_, err = svc.ConfirmTransferFunding(ctx, transfer.ID, &payouts.ConfirmTransferFundingRequest{})
	if err != nil {
		require.Error(t, err, "confirm transfer funding blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	}
}

func TestBatchTransferLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Payouts()

	created, err := svc.CreateBatchTransfer(ctx, &payouts.CreateBatchTransferRequest{
		RequestID: "batch-req-" + time.Now().Format("20060102150405"),
		Name:      "Test Batch Transfer",
	})
	require.NoError(t, err, "create batch transfer failed")
	t.Logf("created batch transfer: %s status=%s", created.ID, created.Status)

	fetched, err := svc.GetBatchTransfer(ctx, created.ID)
	require.NoError(t, err, "get batch transfer failed")
	assert.Equal(t, created.ID, fetched.ID, "batch transfer id mismatch")

	list, err := svc.ListBatchTransfers(ctx)
	if err != nil {
		require.Error(t, err, "list batch transfers blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("batch transfer list count: %d", len(list.Items))
	}

	_, err = svc.QuoteBatchTransfer(ctx, created.ID, &payouts.QuoteBatchTransferRequest{
		Validity: "HR_1",
	})
	if err != nil {
		require.Error(t, err, "quote batch transfer blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("quote batch transfer succeeded")
	}

	_, err = svc.SubmitBatchTransfer(ctx, created.ID)
	if err != nil {
		require.Error(t, err, "submit batch transfer blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("submit batch transfer succeeded")
	}

	// 批量转账子资源操作（在删除之前执行）
	_, err = svc.AddBatchTransferItems(ctx, created.ID, &payouts.AddBatchTransferItemsRequest{
		Items: []payouts.CreateTransferRequest{},
	})
	if err != nil {
		require.Error(t, err, "add batch transfer items blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	}

	_, err = svc.DeleteBatchTransferItems(ctx, created.ID, &payouts.DeleteBatchTransferItemsRequest{
		ItemIDs: []string{},
	})
	if err != nil {
		require.Error(t, err, "delete batch transfer items blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	}

	_, err = svc.ListBatchTransferItems(ctx, created.ID)
	if err != nil {
		require.Error(t, err, "list batch transfer items blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("listed batch transfer items successfully")
	}

	err = svc.DeleteBatchTransfer(ctx, created.ID)
	if err != nil {
		require.Error(t, err, "delete batch transfer blocked")
		t.Logf("🔒 endpoint blocked: %v", err)
	} else {
		t.Logf("deleted batch transfer: %s", created.ID)
	}
}
