package api

import (
	"encoding/json"
	"fmt"
	"github.com/adanrsantos/TradeTUI/providers/databento/types"
	"net/http"
	"net/url"
	// "strconv"
	"io"
)

func HistoryEstimateCost(query types.Query, apiKey string) (float64, error) {
	params := url.Values{}

	params.Set("dataset", query.Dataset.Value)
	params.Set("symbols", query.Symbol.Value)
	params.Set("schema", query.Schema.Value)
	params.Set(
		"start",
		query.StartDate.Format("2006-01-02T15:04:05"),
	)
	params.Set(
		"end",
		query.EndDate.Format("2006-01-02T15:04:05"),
	)

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
