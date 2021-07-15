package azureboards

import (
	"fmt"
	"testing"
)

func Test_IssueTracker_IssueURLFor(t *testing.T) {

	var tests = []struct {
		testName string
		input    string
		taskID   string
		want     string
	}{
		{"URL without special symbols", "https://dev.azure.com/todoerruser/todocheck", "1", "https://dev.azure.com/todoerruser/todocheck/_apis/wit/workitems/1?api-version=6.0"},
		{"URL with special symbols", "https://dev.azure.com/123foo998!/thebigproject123.", "2", "https://dev.azure.com/123foo998!/thebigproject123./_apis/wit/workitems/2?api-version=6.0"},
		{"URL without https prefix", "dev.azure.com/todoerruser/todocheck", "225", "https://dev.azure.com/todoerruser/todocheck/_apis/wit/workitems/225?api-version=6.0"},
		{"URL with special symbol in username", "dev.azure.com/foo.bar/quixproject", "123", "https://dev.azure.com/foo.bar/quixproject/_apis/wit/workitems/123?api-version=6.0"},
	}
	for _, tt := range tests {
		var it IssueTracker

		testname := fmt.Sprintf("%q", tt.testName)
		t.Run(testname, func(t *testing.T) {
			it.Origin = tt.input
			res := it.IssueURLFor(tt.taskID)
			if res != tt.want {
				t.Errorf("got %s, want %s", res, tt.want)
			}
		})
	}

}
