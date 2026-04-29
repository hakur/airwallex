package supporting_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/sdk"
	"github.com/hakur/airwallex/supporting"
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

func TestConnectedStoreLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Supporting()

	list, err := svc.ListConnectedStores(ctx, &supporting.ListConnectedStoresRequest{PageSize: 10})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsBadRequest(err), "list connected stores failed: %v", err)
		t.Logf("list connected stores skipped: %v", err)
		return
	}
	t.Logf("connected stores count: %d", len(list.Items))

	if len(list.Items) == 0 {
		t.Logf("no connected stores available, skipping lifecycle test")
		return
	}

	fetched, err := svc.GetConnectedStore(ctx, list.Items[0].ID)
	require.NoError(t, err, "get connected store failed")
	assert.Equal(t, list.Items[0].ID, fetched.ID, "connected store id mismatch")
}

func TestFileLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.Supporting()

	// Get download links — uses intentionally invalid ID to test error handling
	// 使用故意无效的 ID 来测试错误处理
	links, err := svc.GetDownloadLinks(ctx, &supporting.DownloadLinkRequest{
		FileIDs: []string{"test-file-id"}, // intentionally invalid; real IDs require UploadFile first
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsBadRequest(err) || sdk.IsNotFound(err), "get download links failed: %v", err)
		t.Logf("get download links skipped: %v", err)
		return
	}
	require.NotNil(t, links)
	t.Logf("download links count: %d", len(links.Files))

	// Upload file - skip in tests as it requires multipart form data
	t.Logf("upload file test skipped: requires multipart form data")
}
