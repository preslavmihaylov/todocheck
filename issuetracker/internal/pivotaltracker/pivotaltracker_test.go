package pivotaltracker

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
		{"https://pivotaltracker.com/n/projects/1", "2020", "https://www.pivotaltracker.com/services/v5/projects/1/stories/2020"},
		{"pivotaltracker.com/n/projects/1-2_3", "2020", "https://www.pivotaltracker.com/services/v5/projects/1-2_3/stories/2020"},
		{"pIVOTALtracker.com/projects/1-2_3", "2020", "https://www.pivotaltracker.com/services/v5/projects/1-2_3/stories/2020"},
		{"httPS://pivotaltracker.com/projects/1-2_3", "2020", "https://www.pivotaltracker.com/services/v5/projects/1-2_3/stories/2020"},
		{"https://www.pivotaltracker.com/n/projects/2459511", "21", "https://www.pivotaltracker.com/services/v5/projects/2459511/stories/21"},
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
