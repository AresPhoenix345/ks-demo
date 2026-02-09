/*
Copyright 2025 The KubeStellar Authors.

Grafana API client for KubeStellar Console integrations.
*/

package grafana

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Client is a Grafana API client.
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new Grafana client.
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// DashboardSummary holds minimal dashboard metadata.
type DashboardSummary struct {
	UID  string `json:"uid"`
	Title string `json:"title"`
}

// ListDashboards returns dashboard summaries.
func (c *Client) ListDashboards(ctx context.Context) ([]DashboardSummary, error) {
	u, err := url.Parse(c.baseURL + "/api/search")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("type", "dash-db")
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("grafana list dashboards failed: %s", resp.Status)
	}

	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	dashboards := make([]DashboardSummary, 0, len(result))
	for _, d := range result {
		uid, _ := d["uid"].(string)
		title, _ := d["title"].(string)
		dashboards = append(dashboards, DashboardSummary{UID: uid, Title: title})
	}
	return dashboards, nil
}

// GetEmbedURL returns a signed embed URL for a dashboard (simplified; production would use Grafana auth).
func (c *Client) GetEmbedURL(uid string, from, to string) string {
	u, _ := url.Parse(c.baseURL + "/d/" + uid)
	q := u.Query()
	if from != "" {
		q.Set("from", from)
	}
	if to != "" {
		q.Set("to", to)
	}
	q.Set("kiosk", "tv")
	u.RawQuery = q.Encode()
	return u.String()
}
