package gitlab

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
		{"gitlab.com/PRESLAVmihaylov/project", "1", "https://gitlab.com/api/v4/projects/preslavmihaylov%2Fproject/issues/1"},
		{"gitlab.com/user-5/project8-9_1", "1111", "https://gitlab.com/api/v4/projects/user-5%2Fproject8-9_1/issues/1111"},
		{"https://gitLAB.com/user-1/project", "2020", "https://gitlab.com/api/v4/projects/user-1%2Fproject/issues/2020"},
		{"https://gitlab.com/u/project9-1_2-3-4", "2020", "https://gitlab.com/api/v4/projects/u%2Fproject9-1_2-3-4/issues/2020"},
		{"gitlab.myorg.com/PRESLAVmihaylov/project", "1", "https://gitlab.myorg.com/api/v4/projects/preslavmihaylov%2Fproject/issues/1"},
		{"gitlab.myorg.com/user-5/project8-9_1", "1111", "https://gitlab.myorg.com/api/v4/projects/user-5%2Fproject8-9_1/issues/1111"},
		{"https://gitLAB.myorg.com/user-1/project", "2020", "https://gitlab.myorg.com/api/v4/projects/user-1%2Fproject/issues/2020"},
		{"https://gitlab.myORG.com/u/project9-1_2-3-4", "2020", "https://gitlab.myorg.com/api/v4/projects/u%2Fproject9-1_2-3-4/issues/2020"},
		{"myorg.com/PRESLAVmihaylov/project", "20201225", "https://myorg.com/api/v4/projects/preslavmihaylov%2Fproject/issues/20201225"},
		{"myorg.co.uk/PreslavMihaylov/project", "20201226", "https://myorg.co.uk/api/v4/projects/preslavmihaylov%2Fproject/issues/20201226"},
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
