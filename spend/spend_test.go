package spend_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/sdk"
	"github.com/hakur/airwallex/spend"
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

func TestBillLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Spend()

	list, err := svc.ListBills(ctx, &spend.ListBillsRequest{})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "list bills failed: %v", err)
		return
	}
	t.Logf("existing bills: %d", len(list.Items))

	if len(list.Items) == 0 {
		t.Logf("no bills available, skipping lifecycle test")
		return
	}

	fetched, err := svc.GetBill(ctx, list.Items[0].ID)
	require.NoError(t, err, "get bill failed")
	assert.Equal(t, list.Items[0].ID, fetched.ID, "bill id mismatch")

	// Sync bill
	synced, err := svc.SyncBill(ctx, list.Items[0].ID, &spend.SyncBillRequest{
		SyncStatus: spend.BillSyncStatusSynced,
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "sync bill failed: %v", err)
		t.Logf("sync bill skipped: %v", err)
	} else {
		t.Logf("synced bill: %s status=%s", synced.ID, synced.SyncStatus)
	}
}

func TestExpenseLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Spend()

	list, err := svc.ListExpenses(ctx, &spend.ListExpensesRequest{})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "list expenses failed: %v", err)
		return
	}
	t.Logf("existing expenses: %d", len(list.Items))

	if len(list.Items) == 0 {
		t.Logf("no expenses available, skipping lifecycle test")
		return
	}

	fetched, err := svc.GetExpense(ctx, list.Items[0].ID)
	require.NoError(t, err, "get expense failed")
	assert.Equal(t, list.Items[0].ID, fetched.ID, "expense id mismatch")

	// Sync expense
	synced, err := svc.SyncExpense(ctx, list.Items[0].ID, &spend.SyncExpenseRequest{
		SyncStatus: spend.ExpenseSyncStatusSynced,
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "sync expense failed: %v", err)
		t.Logf("sync expense skipped: %v", err)
	} else {
		t.Logf("synced expense: %s status=%s", synced.ID, synced.SyncStatus)
	}
}

func TestPurchaseOrderLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Spend()

	// 先创建供应商
	vendor, err := svc.CreateVendor(ctx, &spend.CreateVendorRequest{
		RequestID: "vn-req-" + time.Now().Format("20060102150405"),
		Name:      "Test Vendor for PO",
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "create vendor for purchase order failed: %v", err)
		return
	}

	list, err := svc.ListPurchaseOrders(ctx, &spend.ListPurchaseOrdersRequest{})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "list purchase orders failed: %v", err)
		return
	}
	t.Logf("existing purchase orders: %d", len(list.Items))

	created, err := svc.CreatePurchaseOrder(ctx, &spend.CreatePurchaseOrderRequest{
		RequestID:           "po-req-" + time.Now().Format("20060102150405"),
		LegalEntityID:       "le_test",
		VendorID:            vendor.ID,
		PurchaseOrderNumber: "PO-TEST-001",
		SyncStatus:          spend.POSyncStatusNotSynced,
		BillingCurrency:     "USD",
		LineItems: []spend.CreatePOLineItem{
			{
				Quantity:  "1",
				UnitPrice: "200.0",
			},
		},
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "create purchase order failed: %v", err)
		return
	}
	t.Logf("created purchase order: %s", created.ID)

	fetched, err := svc.GetPurchaseOrder(ctx, created.ID)
	require.NoError(t, err, "get purchase order failed")
	assert.Equal(t, created.ID, fetched.ID, "purchase order id mismatch")

	// Sync purchase order
	synced, err := svc.SyncPurchaseOrder(ctx, created.ID, &spend.SyncPurchaseOrderRequest{
		SyncStatus: spend.POSyncStatusSynced,
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "sync purchase order failed: %v", err)
		t.Logf("sync purchase order skipped: %v", err)
	} else {
		t.Logf("synced purchase order: %s status=%s", synced.ID, synced.SyncStatus)
	}
}

func TestReimbursementLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Spend()

	list, err := svc.ListReimbursementReports(ctx, &spend.ListReimbursementReportsRequest{})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "list reimbursement reports failed: %v", err)
		return
	}
	t.Logf("existing reimbursement reports: %d", len(list.Items))

	if len(list.Items) == 0 {
		t.Logf("no reimbursement reports available, skipping lifecycle test")
		return
	}

	fetched, err := svc.GetReimbursementReport(ctx, list.Items[0].ID)
	require.NoError(t, err, "get reimbursement report failed")
	assert.Equal(t, list.Items[0].ID, fetched.ID, "reimbursement report id mismatch")

	// Mark as paid
	marked, err := svc.MarkReimbursementReportAsPaid(ctx, list.Items[0].ID)
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "mark reimbursement report as paid failed: %v", err)
		t.Logf("mark reimbursement report as paid skipped: %v", err)
	} else {
		t.Logf("marked reimbursement report as paid: %s status=%s", marked.ID, marked.Status)
	}
}

func TestVendorLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Spend()

	list, err := svc.ListVendors(ctx, &spend.ListVendorsRequest{})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "list vendors failed: %v", err)
		return
	}
	t.Logf("existing vendors: %d", len(list.Items))

	created, err := svc.CreateVendor(ctx, &spend.CreateVendorRequest{
		RequestID:      "vn-req-" + time.Now().Format("20060102150405"),
		ExternalID:     "ext-" + time.Now().Format("20060102150405"),
		Name:           "test vendor",
		LegalEntityIDs: []string{"le_test"},
		Status:         spend.VendorStatusDraft,
		SyncStatus:     spend.VendorSyncStatusNotSynced,
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "create vendor failed: %v", err)
		return
	}
	t.Logf("created vendor: %s", created.ID)

	fetched, err := svc.GetVendor(ctx, created.ID)
	require.NoError(t, err, "get vendor failed")
	assert.Equal(t, created.ID, fetched.ID, "vendor id mismatch")

	// Sync vendor
	synced, err := svc.SyncVendor(ctx, created.ID, &spend.SyncVendorRequest{
		SyncStatus: spend.VendorSyncStatusSynced,
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsNotFound(err), "sync vendor failed: %v", err)
		t.Logf("sync vendor skipped: %v", err)
	} else {
		t.Logf("synced vendor: %s status=%s", synced.ID, synced.SyncStatus)
	}
}
