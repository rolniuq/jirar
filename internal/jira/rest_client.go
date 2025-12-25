// Package jira provides REST client implementation for Jira API.
package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"

	"jirar/internal/config"
)

// restClient implements the Client interface using REST API.
type restClient struct {
	client *resty.Client
	config *config.JiraConfig
	logger *logrus.Logger
}

// NewClient creates a new Jira REST client.
func NewClient(cfg *config.JiraConfig, logger *logrus.Logger) Client {
	client := resty.New().
		SetTimeout(30 * time.Second).
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return r.StatusCode() >= 500 || err != nil
		})

	return &restClient{
		client: client,
		config: cfg,
		logger: logger,
	}
}

// SearchIssues implements Client interface.
func (c *restClient) SearchIssues(ctx context.Context, jql string, opts ...SearchOption) (*SearchResult, error) {
	options := &SearchOptions{
		Limit:   50,
		StartAt: 0,
		Fields:  []string{"summary", "status", "priority", "assignee", "updated", "created", "project"},
	}

	for _, opt := range opts {
		opt(options)
	}

	url := fmt.Sprintf("%s/rest/api/3/search", c.config.Domain)

	resp, err := c.client.R().
		SetContext(ctx).
		SetBasicAuth(c.config.Email, c.config.Token).
		SetQueryParam("jql", jql).
		SetQueryParam("fields", "summary,status,priority,assignee,updated,created,project").
		SetQueryParam("maxResults", fmt.Sprintf("%d", options.Limit)).
		SetQueryParam("startAt", fmt.Sprintf("%d", options.StartAt)).
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		c.logger.WithError(err).Error("Failed to search issues")
		return nil, fmt.Errorf("search failed: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		c.logger.WithField("status", resp.StatusCode()).Error("Search request failed")
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode())
	}

	var result SearchResult
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		c.logger.WithError(err).Error("Failed to parse search response")
		return nil, fmt.Errorf("parse response failed: %w", err)
	}

	c.logger.WithFields(logrus.Fields{
		"total":    result.Total,
		"returned": len(result.Issues),
	}).Debug("Search completed")

	return &result, nil
}

// GetIssue implements Client interface.
func (c *restClient) GetIssue(ctx context.Context, key string) (*Issue, error) {
	url := fmt.Sprintf("%s/rest/api/3/issue/%s", c.config.Domain, key)

	resp, err := c.client.R().
		SetContext(ctx).
		SetBasicAuth(c.config.Email, c.config.Token).
		SetQueryParam("fields", "summary,status,priority,assignee,updated,created,project,description,reporter").
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		c.logger.WithError(err).Error("Failed to get issue")
		return nil, fmt.Errorf("get issue failed: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		c.logger.WithFields(logrus.Fields{
			"key":    key,
			"status": resp.StatusCode(),
		}).Error("Get issue request failed")
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode())
	}

	var issue Issue
	if err := json.Unmarshal(resp.Body(), &issue); err != nil {
		c.logger.WithError(err).Error("Failed to parse issue response")
		return nil, fmt.Errorf("parse response failed: %w", err)
	}

	c.logger.WithField("key", issue.Key).Debug("Issue retrieved successfully")
	return &issue, nil
}

// GetCurrentUser implements Client interface.
func (c *restClient) GetCurrentUser(ctx context.Context) (*User, error) {
	url := fmt.Sprintf("%s/rest/api/3/myself", c.config.Domain)

	resp, err := c.client.R().
		SetContext(ctx).
		SetBasicAuth(c.config.Email, c.config.Token).
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		c.logger.WithError(err).Error("Failed to get current user")
		return nil, fmt.Errorf("get user failed: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		c.logger.WithField("status", resp.StatusCode()).Error("Get user request failed")
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode())
	}

	var user User
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		c.logger.WithError(err).Error("Failed to parse user response")
		return nil, fmt.Errorf("parse response failed: %w", err)
	}

	c.logger.WithField("user", user.DisplayName).Debug("Current user retrieved")
	return &user, nil
}

// ValidateCredentials implements Client interface.
func (c *restClient) ValidateCredentials(ctx context.Context) error {
	_, err := c.GetCurrentUser(ctx)
	return err
}
