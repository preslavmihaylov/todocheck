package taskstatus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/preslavmihaylov/todocheck/models"
)

// TaskStatus of linked issue
type TaskStatus int

// Supported task statuses
const (
	None TaskStatus = iota
	Open
	Closed
	NonExistent
)

// Fetcher for task statuses by contacting task management web apps' rest api
type Fetcher struct {
	origin, authToken string
}

// NewFetcher instance
func NewFetcher(origin, authToken string) *Fetcher {
	return &Fetcher{origin, authToken}
}

// Fetch a task's status based on task ID
func (f *Fetcher) Fetch(taskID string) (TaskStatus, error) {
	req, err := http.NewRequest("GET", f.origin+taskID, nil)
	if err != nil {
		return None, fmt.Errorf("failed creating new GET request: %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+f.authToken)
	hclient := &http.Client{}
	resp, err := hclient.Do(req)
	if err != nil {
		return None, fmt.Errorf("couldn't execute GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return None, fmt.Errorf("couldn't read response body: %w", err)
	} else if resp.StatusCode == http.StatusNotFound {
		return NonExistent, nil
	} else if resp.StatusCode != http.StatusOK {
		return None, fmt.Errorf("bad status code upon fetching task: %w", err)
	}

	var task *models.JiraTask
	err = json.Unmarshal(body, &task)
	if err != nil {
		return None, fmt.Errorf("couldn't unmarshal response task JSON: %w", err)
	}

	switch task.Fields.Status.Name {
	case "Done":
		return Closed, nil
	default:
		return Open, nil
	}
}
