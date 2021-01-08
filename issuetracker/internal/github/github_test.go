package github

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
		{"https://github.com/preslavmihaylov1/todocheck", "1", "https://api.github.com/repos/preslavmihaylov1/todocheck/issues/1"},
		{"github.com/preslavmihaylov/todocheck/", "8", "https://api.github.com/repos/preslavmihaylov/todocheck/issues/8"},
		{"github.com/uSER-1989/todocheck/", "8", "https://api.github.com/repos/user-1989/todocheck/issues/8"},
		{"github.com/user2020/hyphen-1/", "8", "https://api.github.com/repos/user2020/hyphen-1/issues/8"},
		{"github.com/user2020/hyphen-1_underscore/", "8", "https://api.github.com/repos/user2020/hyphen-1_underscore/issues/8"},
		{"GITHUB.com/u1u/hyphen-1_underscore_x-2/", "8", "https://api.github.com/repos/u1u/hyphen-1_underscore_x-2/issues/8"},
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
