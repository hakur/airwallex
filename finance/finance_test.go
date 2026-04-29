package finance_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/finance"
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

func TestListFinancialTransactions(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Finance()

	transactions, err := svc.ListFinancialTransactions(ctx)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsBadRequest(err), "list financial transactions failed: %v", err)
		t.Logf("list financial transactions skipped: %v", err)
		return
	}
	assert.NotNil(t, transactions, "financial transactions list should not be nil")
	t.Logf("financial transactions count: %d", len(transactions.Items))

	if len(transactions.Items) > 0 {
		ft, err := svc.GetFinancialTransaction(ctx, transactions.Items[0].ID)
		require.NoError(t, err, "get financial transaction failed")
		assert.Equal(t, transactions.Items[0].ID, ft.ID, "financial transaction id mismatch")
		t.Logf("got financial transaction: %s", ft.ID)
	}
}

func TestFinancialReportsLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Finance()

	reports, err := svc.ListFinancialReports(ctx)
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsBadRequest(err), "list financial reports failed: %v", err)
		t.Logf("list financial reports skipped: %v", err)
		return
	}
	assert.NotNil(t, reports, "financial reports list should not be nil")
	t.Logf("financial reports count: %d", len(reports.Items))

	if len(reports.Items) > 0 {
		report, err := svc.GetFinancialReport(ctx, reports.Items[0].ID)
		require.NoError(t, err, "get financial report failed")
		assert.Equal(t, reports.Items[0].ID, report.ID, "financial report id mismatch")
		t.Logf("got financial report: %s status=%s", report.ID, report.Status)

		// Get report content
		content, err := svc.GetFinancialReportContent(ctx, reports.Items[0].ID)
		if err != nil {
			require.True(t, sdk.IsUnauthorized(err) || sdk.IsNotFound(err) || sdk.IsBadRequest(err), "get financial report content failed: %v", err)
			t.Logf("get financial report content skipped: %v", err)
		} else {
			require.NotNil(t, content)
			t.Logf("got financial report content: %d bytes", len(content))
		}
	}
}

func TestSettlementLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Finance()

	list, err := svc.ListSettlements(ctx, &finance.ListSettlementsRequest{
		Currency:      sdk.CurrencyUSD,
		FromSettledAt: time.Now().AddDate(0, -1, 0).Format("2006-01-02T15:04:05Z"),
		ToSettledAt:   time.Now().Format("2006-01-02T15:04:05Z"),
		Status:        finance.SettlementStatusSettled,
		PageSize:      10,
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsBadRequest(err) || sdk.IsValidationError(err), "list settlements failed: %v", err)
		t.Logf("list settlements skipped: %v", err)
		return
	}
	assert.NotNil(t, list, "settlements list should not be nil")
	t.Logf("settlements count: %d", len(list.Items))

	if len(list.Items) > 0 {
		settlement, err := svc.GetSettlement(ctx, list.Items[0].ID)
		require.NoError(t, err, "get settlement failed")
		assert.Equal(t, list.Items[0].ID, settlement.ID, "settlement id mismatch")
		t.Logf("got settlement: %s status=%s", settlement.ID, settlement.Status)
	}

	// 遍历结算列表，找到第一个能成功获取报告的 settlement
	var report *finance.SettlementReport
	for _, item := range list.Items {
		r, err := svc.GetSettlementReport(ctx, item.ID, &finance.GetSettlementReportRequest{})
		if err == nil {
			report = r
			t.Logf("got settlement report: %s settlement_id=%s", report.ID, item.ID)
			break
		}
		t.Logf("settlement %s report not available: %v", item.ID, err)
	}
	if report == nil {
		t.Logf("no settlement report available for %d settlements", len(list.Items))
	} else {
		require.NotNil(t, report)
		require.NotEmpty(t, report.ID, "report ID should not be empty")
	}
}
