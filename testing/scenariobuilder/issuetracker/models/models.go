package models

// Status Jira struct
type Status struct {
	Name string `json:"name"`
}

// Fields Jira struct
type Fields struct {
	Status `json:"status"`
}

// JiraTask JSON model as returned by Rest API
type JiraTask struct {
	Fields `json:"fields"`
}
