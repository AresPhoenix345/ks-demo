/*
Copyright 2025 The KubeStellar Authors.

GitHub API client for workflow and Actions integration.
*/

package github

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client is a GitHub API client for Actions.
type Client struct {
	token      string
	httpClient *http.Client
}

// NewClient creates a new GitHub client.
func NewClient(token string) *Client {
	return &Client{
		token: token,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// WorkflowRun represents a GitHub Actions workflow run.
type WorkflowRun struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	Conclusion string    `json:"conclusion"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	HeadBranch string    `json:"head_branch"`
	HTMLURL    string    `json:"html_url"`
}

// ListWorkflowRuns lists workflow runs for a repository.
func (c *Client) ListWorkflowRuns(ctx context.Context, owner, repo string, perPage int) ([]WorkflowRun, error) {
	if perPage <= 0 {
		perPage = 30
	}
	u := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/runs?per_page=%d", owner, repo, perPage)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api error: %s", resp.Status)
	}

	var result struct {
		WorkflowRuns []WorkflowRun `json:"workflow_runs"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.WorkflowRuns, nil
}

// DispatchWorkflow triggers a workflow_dispatch event.
func (c *Client) DispatchWorkflow(ctx context.Context, owner, repo, workflowID, ref string, inputs map[string]string) error {
	u := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows/%s/dispatches", owner, repo, workflowID)
	body := map[string]interface{}{
		"ref": ref,
	}
	if len(inputs) > 0 {
		body["inputs"] = inputs
	}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("github dispatch failed: %s", resp.Status)
	}
	return nil
}
