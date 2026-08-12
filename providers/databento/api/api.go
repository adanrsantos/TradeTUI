package api

import (
	"encoding/json"
	"fmt"
	"github.com/adanrsantos/TradeTUI/providers/databento/types"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func SaveCandles(candles []types.OHLCV) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Failed to get working directory: %w", err)
	}

	dataDir := filepath.Join(cwd, "data", "databento")

	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("Failed to create data directory: %w", err)
	}

	timestamp := time.Now().Format("2006-01-02T15-04-05")
	filename := timestamp + ".jsonl"

	filePath := filepath.Join(dataDir, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Failed ot create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	for _, candle := range candles {
		if err := encoder.Encode(candle); err != nil {
			return fmt.Errorf("Failed to encode candle: %w", err)
		}
	}

	return nil
}

func HistoryRequest(query types.Query, apiKey string) ([]types.OHLCV, error) {
	params := url.Values{}

	params.Set("dataset", query.Dataset.Value)
	params.Set("symbols", query.Symbol.Value)
	params.Set("schema", query.Schema.Value)
	params.Set("start", query.StartDate.Format(time.RFC3339))
	params.Set("end", query.EndDate.Format(time.RFC3339))

	if *query.Limit != 0 {
		params.Set("limit", strconv.Itoa(*query.Limit))
	}

	params.Set("encoding", "json")
	params.Set("stype_in", "continuous")

	reqURL := "https://hist.databento.com/v0/timeseries.get_range?" + params.Encode()

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	req.SetBasicAuth(apiKey, "")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to contact Databento: %w", err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var candles []types.OHLCV

	for {
		var candle types.OHLCV

		err := decoder.Decode(&candle)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("Failed to decode response body: %w", err)
		}

		candles = append(candles, candle)
	}

	return candles, nil
}

func HistoryEstimateCost(query types.Query, apiKey string) (float64, error) {
	params := url.Values{}

	params.Set("dataset", query.Dataset.Value)
	params.Set("symbols", query.Symbol.Value)
	params.Set("schema", query.Schema.Value)
	params.Set("start", query.StartDate.Format(time.RFC3339))
	params.Set("end", query.EndDate.Format(time.RFC3339))
	if *query.Limit != 0 {
		params.Set("limit", strconv.Itoa(*query.Limit))
	}
	params.Set("stype_in", "continuous")

	reqURL := "https://hist.databento.com/v0/metadata.get_cost?" + params.Encode()

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, fmt.Errorf("Failed to create request: %w", err)
	}

	req.SetBasicAuth(apiKey, "")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("Failed to contact Databento: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("Failed to read Databento response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(
			"Databento request failed (%s): %s",
			resp.Status,
			string(body),
		)
	}

	var cost float64
	if err := json.Unmarshal(body, &cost); err != nil {
		return 0, fmt.Errorf(
			"Failed to decode Databento response: %w; response: %s",
			err,
			string(body),
		)
	}

	return cost, nil
}
