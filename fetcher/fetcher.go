package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/preslavmihaylov/todocheck/authmanager/authmiddleware"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
	"github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"
)

// Fetcher for task statuses by contacting task management web apps' rest api
type Fetcher struct {
	origin         string
	tracker        config.IssueTracker
	authMiddleware authmiddleware.Func
}

// NewFetcher instance
func NewFetcher(origin string, tracker config.IssueTracker, authMw authmiddleware.Func) *Fetcher {
	return &Fetcher{origin, tracker, authMw}
}

// Fetch a task's status based on task ID
func (f *Fetcher) Fetch(taskID string) (taskstatus.TaskStatus, error) {
	req, err := http.NewRequest("GET", f.origin+taskID, nil)
	if err != nil {
		return taskstatus.None, fmt.Errorf("failed creating new GET request: %w", err)
	}

	f.authMiddleware(req)

	hclient := &http.Client{}
	resp, err := hclient.Do(req)
	if err != nil {
		return taskstatus.None, fmt.Errorf("couldn't execute GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return taskstatus.None, fmt.Errorf("couldn't read response body: %w", err)
	} else if resp.StatusCode == http.StatusNotFound {
		return taskstatus.NonExistent, nil
	} else if resp.StatusCode != http.StatusOK {
		return taskstatus.None, fmt.Errorf("bad status code upon fetching task: %d - %s", resp.StatusCode, string(body))
	}

	task := issuetracker.TaskFor(f.tracker)
	err = json.Unmarshal(body, &task)
	if err != nil {
		return taskstatus.None, fmt.Errorf("couldn't unmarshal response task JSON: %w", err)
	}

	return task.GetStatus(), nil
}
