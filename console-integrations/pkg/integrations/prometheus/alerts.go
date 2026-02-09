/*
Copyright 2025 The KubeStellar Authors.

Prometheus AlertManager API client for KubeStellar Console integrations.
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

// Alert status from AlertManager.
type Alert struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	State       string            `json:"state"`
	Value       string            `json:"value"`
	ActiveAt    time.Time         `json:"activeAt"`
}

// AlertsResponse holds AlertManager API response.
type AlertsResponse struct {
	Status string  `json:"status"`
	Data   []Alert `json:"data"`
}

// GetAlerts fetches alerts from AlertManager (default: same host as Prometheus, /api/v2/alerts).
func (c *Client) GetAlerts(ctx context.Context, alertmanagerURL string) ([]Alert, error) {
	base := alertmanagerURL
	if base == "" {
		base = c.baseURL
	}
	u, err := url.Parse(base + "/api/v2/alerts")
	if err != nil {
		return nil, err
	}

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
		return nil, fmt.Errorf("alertmanager request failed: %s", resp.Status)
	}

	var result AlertsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data, nil
}
