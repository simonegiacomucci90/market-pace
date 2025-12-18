package fetcher

import (
	"os"
	"testing"
)

func TestFetchDaily(t *testing.T) {
	apiKey := os.Getenv("ALPHAVANTAGE_API_KEY")
	if apiKey == "" {
		t.Skip("ALPHAVANTAGE_API_KEY not set, skipping test")
	}

	fetcher := NewAlphaVantageFetcher(apiKey)
	data, err := fetcher.FetchDaily("SPY")
	if err != nil {
		t.Fatalf("FetchDaily failed: %v", err)
	}

	if len(data.TimeSeries) == 0 {
		t.Fatalf("Expected some daily data, got none")
	}
}
