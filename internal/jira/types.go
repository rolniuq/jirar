// Package jira provides types for Jira API responses.
package jira

import "time"

// Issue represents a Jira issue/ticket.
type Issue struct {
	Key    string `json:"key"`
	ID     string `json:"id"`
	Self   string `json:"self"`
	Fields Fields `json:"fields"`
}

// Fields contains all issue fields.
type Fields struct {
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	Priority    Priority  `json:"priority"`
	Assignee    User      `json:"assignee"`
	Reporter    User      `json:"reporter"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	DueDate     time.Time `json:"duedate"`
	Project     Project   `json:"project"`
	IssueType   IssueType `json:"issuetype"`
}

// Status represents issue status.
type Status struct {
	Name           string         `json:"name"`
	ID             string         `json:"id"`
	StatusCategory StatusCategory `json:"statusCategory"`
}

// StatusCategory represents the category of a status.
type StatusCategory struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	ColorName string `json:"colorName"`
}

// Priority represents issue priority.
type Priority struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	IconURL string `json:"iconUrl"`
}

// User represents a Jira user.
type User struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Email       string `json:"emailAddress"`
	Active      bool   `json:"active"`
}

// Project represents a Jira project.
type Project struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	ID   string `json:"id"`
}

// IssueType represents the type of issue.
type IssueType struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	IconURL string `json:"iconUrl"`
}

// SearchResult contains the results of a JQL search.
type SearchResult struct {
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	Issues     []Issue `json:"issues"`
}

// CurrentUser represents the authenticated user.
type CurrentUser struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Email       string `json:"emailAddress"`
	Active      bool   `json:"active"`
	TimeZone    string `json:"timeZone"`
}
