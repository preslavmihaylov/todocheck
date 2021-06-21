package youtrack

import (
	"fmt"
	"testing"
)

func Test_IssueTracker_IssueURLFor(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		taskID   string
		expected string
	}{
		// YouTrack InCloud tests
		{"Test YouTrack InCloud base case", "youtrack.myjetbrains.com", "taskId", "https://youtrack.myjetbrains.com/youtrack/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"Test YouTrack InCloud with trailing slash", "youtrack.myjetbrains.com/", "taskId", "https://youtrack.myjetbrains.com/youtrack/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"Test YouTrack InCloud with trailing slash and sequence after instance name", "https://youtrack.myjetbrains.com/n/projects/1", "taskId", "https://youtrack.myjetbrains.com/youtrack/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"Test YouTrack InCloud with capital letters, trailing slash and sequence without http(s):// and www.", "yOUtrack.myjetbrains.com/thats/trailing/sequence", "taskId", "https://youtrack.myjetbrains.com/youtrack/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"Test YouTrack InCloud without http(s):// and www.", "youtrack.myjetbrains.com/thats/trailing/sequence", "taskId", "https://youtrack.myjetbrains.com/youtrack/api/issues/taskId?fields=customFields(value(isResolved))"},

		// Youtrack Standalone tests
		{"Test YouTrack Standalone base case", "youtrack-standalone.com", "taskId", "https://youtrack-standalone.com/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"Test YouTrack Standalone with http", "http://youtrack-standalone.com", "taskId", "http://youtrack-standalone.com/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"Test YouTrack Standalone with default port number", "youtrack.standalone.com:8080", "taskId", "https://youtrack.standalone.com:8080/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"Test YouTrack Standalone with non-default port number", "youtrack.standalone.com:12345", "taskId", "https://youtrack.standalone.com:12345/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"Test YouTrack Standalone with port number and trailing slash", "https://youtrack.com:8080/", "taskId", "https://youtrack.com:8080/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"Test YouTrack Standalone with www., port number and trailing slash and sequence", "https://www.youtrack.com:8080/trailing/seq", "taskId", "https://www.youtrack.com:8080/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"Test YouTrack Standalone with localhost", "localhost:8080", "taskId", "https://localhost:8080/api/issues/taskId?fields=customFields(value(isResolved))"},
	}

	for _, test := range tests {
		var it IssueTracker
		testname := fmt.Sprintf("%q", test.input)
		t.Run(testname, func(t *testing.T) {
			it.Origin = test.input
			res := it.IssueURLFor(test.taskID)
			if res != test.expected {
				t.Errorf("\nTest: %s\n\nest input: %s\nExpected: %s", test.name, res, test.expected)
			}
		})
	}
}
