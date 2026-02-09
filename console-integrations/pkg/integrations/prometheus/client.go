/*
Copyright 2025 The KubeStellar Authors.

Prometheus API client for KubeStellar Console integrations.
*/

package prometheus

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Client is a Prometheus API client.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Prometheus client.
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// QueryResult holds Prometheus query API response.
type QueryResult struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result    []struct {
			Metric map[string]string `json:"metric"`
			Value  []interface{}     `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

// Query executes a PromQL query.
func (c *Client) Query(ctx context.Context, query string, t time.Time) (*QueryResult, error) {
	u, err := url.Parse(c.baseURL + "/api/v1/query")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("query", query)
	q.Set("time", t.Format(time.RFC3339))
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("prometheus query failed: %s", resp.Status)
	}

	var result QueryResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RangeQuery executes a PromQL range query.
func (c *Client) RangeQuery(ctx context.Context, query string, start, end time.Time, step time.Duration) (interface{}, error) {
	u, err := url.Parse(c.baseURL + "/api/v1/query_range")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("query", query)
	q.Set("start", start.Format(time.RFC3339))
	q.Set("end", end.Format(time.RFC3339))
	q.Set("step", step.String())
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("prometheus range query failed: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
