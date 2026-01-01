package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AlphaVantageResponse struct {
	TimeSeries map[string]map[string]string `json:"Time Series (Daily)"`
}

type AlphaVantageFetcher struct {
	ApiKey string
}

func NewAlphaVantageFetcher(apikey string) *AlphaVantageFetcher {
	return &AlphaVantageFetcher{ApiKey: apikey}
}

func (f *AlphaVantageFetcher) FetchDaily(symbol string) (AlphaVantageResponse, error) {
	url := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&outputsize=compact&apikey=%s",
		symbol, f.ApiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return AlphaVantageResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	var data AlphaVantageResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return AlphaVantageResponse{}, err
	}

	return data, nil
}
