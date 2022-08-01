package fetcher

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/preslavmihaylov/todocheck/issuetracker"
	"github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"
)

var testErr = errors.New("")

func TestFetch(t *testing.T) {
	fetcher := NewFetcher(mockIssueTracker{})
	testJson, err := json.Marshal(mockTask{})
	if err != nil {
		t.Fatalf("Test json is bad")
	}

	testData := []struct {
		Task   string
		Client mockClient
		Status int
		Err    error
	}{
		{
			Task:   "GoodFetch",
			Client: mockClient{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(testJson)), Err: nil},
			Status: 1,
			Err:    nil,
		},
		{
			Task:   "BadURL",
			Client: mockClient{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(testJson)), Err: nil},
			Status: 0,
			Err:    testErr,
		},
		{
			Task:   "MiddlewareFailure",
			Client: mockClient{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(testJson)), Err: nil},
			Status: 0,
			Err:    testErr,
		},
		{
			Task:   "FailedSendingRequest",
			Client: mockClient{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(testJson)), Err: testErr},
			Status: 0,
			Err:    testErr,
		},
		{
			Task:   "BadReader",
			Client: mockClient{StatusCode: 200, Body: errReader(0), Err: nil},
			Status: 0,
			Err:    testErr,
		},
		{
			Task:   "ResponseStatusNotFound",
			Client: mockClient{StatusCode: 404, Body: io.NopCloser(bytes.NewReader(testJson)), Err: nil},
			Status: 3,
			Err:    nil,
		},
		{
			Task:   "ResponseBadStatus",
			Client: mockClient{StatusCode: 405, Body: io.NopCloser(bytes.NewReader(testJson)), Err: nil},
			Status: 0,
			Err:    testErr,
		},
		{
			Task:   "BadJson",
			Client: mockClient{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte{})), Err: nil},
			Status: 0,
			Err:    testErr,
		},
	}
	for _, tt := range testData {
		t.Run(tt.Task, func(t *testing.T) {
			fetcher.sendRequest = tt.Client.sendRequest
			taskStatus, err := fetcher.Fetch(tt.Task)
			if taskStatus != taskstatus.TaskStatus(tt.Status) {
				t.Errorf("Task status is %v, expected %v", taskStatus, taskstatus.TaskStatus(tt.Status))
			}
			if (err == nil) != (tt.Err == nil) { // Doesn't care about error message or type
				t.Errorf("Fetch error is %v, expected %v", err, tt.Err)
			}
		})
	}

}

// Mocking Task
type mockTask struct {
	Status string
}

func (t mockTask) GetStatus() (taskstatus.TaskStatus, error) {
	return taskstatus.Open, nil
}

// Mocking IssueTracker
type mockIssueTracker struct {
}

func (it mockIssueTracker) TaskModel() issuetracker.Task {
	return &mockTask{}
}

func (it mockIssueTracker) IssueURLFor(taskID string) string {
	if taskID == "BadURL" {
		return string(byte(' ') - 1) // This causes http.NewRequest to fail
	}
	return taskID
}

func (it mockIssueTracker) Exists() bool { // Never called
	return false
}

func (it mockIssueTracker) InstrumentMiddleware(r *http.Request) error {
	if r.URL.Path == "MiddlewareFailure" { // The taskID is set as URL path, so we're using that as our fail trigger
		return testErr
	}
	return nil
}

func (it mockIssueTracker) TokenAcquisitionInstructions() string { // Never called
	return ""
}

// Mocking sendRequest
type mockClient struct {
	StatusCode int
	Body       io.ReadCloser
	Err        error
}

func (c mockClient) sendRequest(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: c.StatusCode,
		Body:       c.Body,
	}, c.Err
}

type errReader int

func (r errReader) Read(body []byte) (int, error) {
	return 0, testErr
}

func (r errReader) Close() error {
	return nil
}
