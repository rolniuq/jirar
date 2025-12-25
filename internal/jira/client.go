// Package jira provides client interfaces and types for interacting with Jira API.
package jira

import (
	"context"
)

// Client defines the interface for Jira API operations.
type Client interface {
	// SearchIssues searches for issues using JQL
	SearchIssues(ctx context.Context, jql string, opts ...SearchOption) (*SearchResult, error)

	// GetIssue retrieves a single issue by key
	GetIssue(ctx context.Context, key string) (*Issue, error)

	// GetCurrentUser retrieves information about the authenticated user
	GetCurrentUser(ctx context.Context) (*User, error)

	// ValidateCredentials tests if the current credentials are valid
	ValidateCredentials(ctx context.Context) error
}

// SearchOptions configures how search results are returned.
type SearchOptions struct {
	Limit   int
	StartAt int
	Fields  []string
	Expand  []string
}

// SearchOption applies configuration to search options.
type SearchOption func(*SearchOptions)

// WithLimit sets the maximum number of results to return.
func WithLimit(limit int) SearchOption {
	return func(opts *SearchOptions) {
		opts.Limit = limit
	}
}

// WithStartAt sets the starting index for results.
func WithStartAt(startAt int) SearchOption {
	return func(opts *SearchOptions) {
		opts.StartAt = startAt
	}
}

// WithFields specifies which fields to include in the response.
func WithFields(fields ...string) SearchOption {
	return func(opts *SearchOptions) {
		opts.Fields = append(opts.Fields, fields...)
	}
}

// WithExpand specifies which fields to expand.
func WithExpand(expand ...string) SearchOption {
	return func(opts *SearchOptions) {
		opts.Expand = append(opts.Expand, expand...)
	}
}
