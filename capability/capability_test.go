package capability_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/capability"
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

func TestGetAccountCapability(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Capability()

	cap, err := svc.GetAccountCapability(ctx, capability.CapabilityPaymentsVisa)
	if err != nil {
		require.True(t, sdk.IsNotFound(err) || sdk.IsUnauthorized(err), "get account capability failed: %v", err)
		t.Logf("get account capability skipped: %v", err)
		return
	}
	assert.NotEmpty(t, cap.ID, "capability id should not be empty")
	t.Logf("got capability: %s status=%s", cap.ID, cap.Status)
}

func TestGetFundingLimits(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Capability()

	resp, err := svc.GetFundingLimits(ctx)
	require.NoError(t, err, "get funding limits failed")
	assert.NotNil(t, resp, "funding limits response should not be nil")
	t.Logf("funding limits count: %d", len(resp.Items))
}

func TestEnableAccountCapability(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Capability()

	enrollSME := true
	_, err := svc.EnableAccountCapability(ctx, capability.CapabilityPaymentsVisa, &capability.EnableAccountCapabilityRequest{
		RequestID:        "cap-req-test",
		ID:               capability.CapabilityPaymentsVisa,
		EnrollSMEProgram: &enrollSME,
		EntityType:       capability.EntityTypeBusiness,
	})
	if err != nil {
		require.True(t, sdk.IsBadRequest(err) || sdk.IsUnauthorized(err) || sdk.IsValidationError(err), "enable account capability failed: %v", err)
		t.Logf("enable account capability skipped: %v", err)
		return
	}
	t.Logf("enabled capability: %s", capability.CapabilityPaymentsVisa)
}
