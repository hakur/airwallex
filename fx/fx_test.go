package fx_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hakur/airwallex"
	"github.com/hakur/airwallex/fx"
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

func TestRateLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.FX()

	rate, err := svc.GetRates(ctx, &fx.GetRatesRequest{
		BuyCurrency:  sdk.CurrencyUSD,
		SellCurrency: sdk.CurrencyEUR,
		BuyAmount:    10000,
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "get rates failed: %v", err)
		return
	}
	require.NotNil(t, rate)
	assert.NotEmpty(t, rate.CurrencyPair)
	assert.Greater(t, rate.Rate, 0.0)
	t.Logf("rate: %s = %f", rate.CurrencyPair, rate.Rate)
}

func TestConversionLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.FX()

	created, err := svc.CreateConversion(ctx, &fx.CreateConversionRequest{
		RequestID:    "fx-req-" + time.Now().Format("20060102150405"),
		BuyCurrency:  sdk.CurrencyUSD,
		SellCurrency: sdk.CurrencyEUR,
		BuyAmount:    "1000",
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "create conversion failed: %v", err)
		return
	}
	t.Logf("created conversion: %s status=%s", created.ConversionID, created.Status)

	fetched, err := svc.GetConversion(ctx, created.ConversionID)
	require.NoError(t, err, "get conversion failed")
	assert.Equal(t, created.ConversionID, fetched.ConversionID, "conversion id mismatch")

	list, err := svc.ListConversions(ctx)
	require.NoError(t, err, "list conversions failed")
	require.NotEmpty(t, list.Items, "expected at least one conversion")
}

func TestQuoteLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.FX()

	created, err := svc.CreateQuote(ctx, &fx.CreateQuoteRequest{
		BuyCurrency:  sdk.CurrencyUSD,
		SellCurrency: sdk.CurrencyEUR,
		BuyAmount:    1000,
		Validity:     fx.QuoteValidityHr1,
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err), "create quote failed: %v", err)
		return
	}
	require.NotNil(t, created)
	t.Logf("created quote: %s rate=%f", created.ID, created.Rate)

	fetched, err := svc.GetQuote(ctx, created.ID)
	require.NoError(t, err, "get quote failed")
	assert.Equal(t, created.ID, fetched.ID, "quote id mismatch")
}

func TestConversionAmendmentLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := testClient.FX()

	// 先创建一个兑换以便有可用的 conversion_id
	conversion, err := svc.CreateConversion(ctx, &fx.CreateConversionRequest{
		RequestID:    "fx-amend-" + time.Now().Format("20060102150405"),
		BuyCurrency:  sdk.CurrencyUSD,
		SellCurrency: sdk.CurrencyEUR,
		BuyAmount:    "100",
	})
	if err != nil {
		require.True(t, sdk.IsUnauthorized(err) || sdk.IsBadRequest(err),
			"create conversion for amendment test failed: %v", err)
		return
	}

	amended, err := svc.CreateConversionAmendment(ctx, &fx.CreateConversionAmendmentRequest{
		RequestID:    "fx-amend-req-" + time.Now().Format("20060102150405"),
		ConversionID: conversion.ConversionID,
		Type:         fx.AmendmentTypeCancel,
	})
	if err != nil {
		// Sandbox 中修正操作可能因状态不允许而失败
		assert.True(t, sdk.IsBadRequest(err) || sdk.IsValidationError(err) || sdk.IsNotFound(err) || sdk.IsInvalidStatusForOperation(err),
			"unexpected error type: %v", err)
	} else {
		t.Logf("created conversion amendment: %s", amended.AmendmentID)
	}
}
