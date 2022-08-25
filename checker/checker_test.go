package checker

import (
	"errors"
	"reflect"
	"testing"

	checkerrors "github.com/preslavmihaylov/todocheck/checker/errors"
	"github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"
)

func TestCheck(t *testing.T) {
	fetcher := mockFetcher{}
	checker := New(&fetcher)
	matcher := mockMatcher{}

	testLines := []string{}
	testLineCnt := 0

	testData := []struct {
		comment, filename string
		todoErr           *checkerrors.TODO
		err               error
	}{
		{"NotMatch", "", nil, nil},
		{"NotValid", "test.go", checkerrors.MalformedTODOErr("test.go", testLines, testLineCnt), nil},
		{"FailedFetch", "", nil, errors.New("")},
		{"ClosedIssue", "test.go", checkerrors.IssueClosedErr("test.go", testLines, testLineCnt, "ClosedIssue"), nil},
		{"NonExistentIssue", "test.go", checkerrors.IssueNonExistentErr("test.go", testLines, testLineCnt, "NonExistentIssue"), nil},
		{"Valid", "", nil, nil},
	}
	for _, tt := range testData {
		t.Run(tt.comment, func(t *testing.T) {

			todoErr, err := checker.Check(matcher, tt.comment, tt.filename, testLines, testLineCnt)
			if !reflect.DeepEqual(todoErr, tt.todoErr) {
				t.Errorf("Expected toddErr to be %v, got %v", tt.todoErr, todoErr)
			}
			if (err == nil) != (tt.err == nil) { // Don't care about the error string
				t.Errorf("Expected err to be %v, got %v", tt.err, err)
			}
		})
	}

	t.Run("NilMatcher", func(t *testing.T) {
		todoErr, err := checker.Check(nil, "", "", testLines, testLineCnt)
		if todoErr != nil {
			t.Errorf("Expected toddErr to be nil, got %v", todoErr)
		}
		if err == nil { // Don't care about the error string
			t.Errorf("Expected err to be not nil")
		}
	})
	t.Run("InvalidExtract", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Logf("Recovered in %v", r)
			}
		}()
		_, err := checker.Check(matcher, "InvalidExtract", "", testLines, testLineCnt)
		if err != nil {
			t.Errorf("Expected err to be nil, got %v", err)
		}
		t.Errorf("Expected code to panic")
	})
}

type mockMatcher struct {
}

func (m mockMatcher) IsMatch(expr string) bool {
	return expr != "NotMatch"
}
func (m mockMatcher) IsValid(expr string) bool {
	return expr != "NotValid"
}
func (m mockMatcher) ExtractIssueRef(expr string) (string, error) {
	if expr == "InvalidExtract" {
		return expr, errors.New("Invalid todo")
	}
	return expr, nil
}

type mockFetcher struct {
}

func (f *mockFetcher) Fetch(taskID string) (taskstatus.TaskStatus, error) {
	if taskID == "FailedFetch" {
		return 0, errors.New("FailedFetch")
	}
	if taskID == "ClosedIssue" {
		return 2, nil
	}
	if taskID == "NonExistentIssue" {
		return 3, nil
	}
	return 0, nil
}
