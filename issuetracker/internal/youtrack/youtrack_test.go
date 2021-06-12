package youtrack

import (
	"fmt"
	"testing"
)

func Test_IssueTracker_IssueURLFor(t *testing.T) {
	var tests = []struct {
		input  string
		taskID string
		want   string
	}{
		{"https://youtrack.com/n/projects/1", "taskId", "https://youtrack.com/youtrack/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"youtrack.com/n/projects/1-2_3", "taskId", "https://youtrack.com/youtrack/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"yOUtrack.com/projects/1-2_3", "taskId", "https://youtrack.com/youtrack/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"httPS://youtrack.com/projects/1-2_3", "taskId", "https://youtrack.com/youtrack/api/issues/taskId?fields=customFields(value(isResolved))"},
		{"https://www.youtrack.com/n/projects/2459511", "taskId", "https://www.youtrack.com/youtrack/api/issues/taskId?fields=customFields(value(isResolved))"},
	}

	for _, tt := range tests {

		var it IssueTracker

		testname := fmt.Sprintf("%q", tt.input)
		t.Run(testname, func(t *testing.T) {
			it.Origin = tt.input
			res := it.IssueURLFor(tt.taskID)
			if res != tt.want {
				t.Errorf("got %s, want %s", res, tt.want)
			}
		})
	}
}
