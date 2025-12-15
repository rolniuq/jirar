package jira

type Status struct {
	Name string `json:"name"`
}

type Issue struct {
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields"`
}

type IssueFields struct {
	Summary string `json:"summary"`
	Status  Status `json:"status"`
	Updated string `json:"updated"`
}

type Jira struct{}

func (j *Jira) GetIssues() []Issue {
	return []Issue{}
}
