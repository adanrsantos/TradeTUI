package api

import (
	"bufio"
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
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	dataDir := filepath.Join(cwd, "data", "databento")

	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}

	data, err := json.MarshalIndent(candles, "", "\t")
	if err != nil {
		return fmt.Errorf("failed to encode candles: %w", err)
	}

	filePath := filepath.Join(dataDir, "candles.json")

	fmt.Println("Saving candles:", len(candles))
	fmt.Println("Saving to:", filePath)

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to save candles to %s: %w", filePath, err)
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
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(apiKey, "")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to contact Databento: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read Databento error response: %w", err)
		}

		return nil, fmt.Errorf(
			"Databento request failed (%s): %s",
			resp.Status,
			string(body),
		)
	}

	var candles []types.OHLCV

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		var candle types.OHLCV

		if err := json.Unmarshal(scanner.Bytes(), &candle); err != nil {
			return nil, fmt.Errorf(
				"failed to decode Databento candle: %w\nresponse: %s",
				err,
				scanner.Text(),
			)
		}

		candles = append(candles, candle)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf(
			"failed to read Databento response: %w",
			err,
		)
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
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(apiKey, "")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to contact Databento: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read Databento response: %w", err)
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
			"failed to decode Databento response: %w; response: %s",
			err,
			string(body),
		)
	}

	return cost, nil
}
