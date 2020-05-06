package models

// JiraTask JSON model as returned by Rest API
type JiraTask struct {
	Fields struct {
		Status struct {
			Name string `json:"name"`
		} `json:"status"`
	} `json:"fields"`
}
