package checker

import (
	"fmt"

	checkererrors "github.com/preslavmihaylov/todocheck/checker/errors"
	"github.com/preslavmihaylov/todocheck/matchers"
	"github.com/preslavmihaylov/todocheck/taskstatus"
)

// Checker for todo lines
type Checker struct {
	statusFetcher *taskstatus.Fetcher
	todoMatcher   matchers.TodoMatcher
}

// New checker
func New(statusFetcher *taskstatus.Fetcher, todoMatcher matchers.TodoMatcher) *Checker {
	return &Checker{statusFetcher, todoMatcher}
}

// IsTODO line, without performing any validity checks
func (c *Checker) IsTODO(line string) bool {
	return c.todoMatcher.IsMatch(line)
}

// Check if todo line is valid
func (c *Checker) Check(filename, line string, linecnt int) (error, error) {
	if !c.todoMatcher.IsValid(line) {
		return checkererrors.MalformedTODOErr(filename, line, linecnt), nil
	}

	taskID, err := c.todoMatcher.ExtractIssueRef(line)
	if err != nil {
		// should never happen after validating todo line
		panic(err)
	}

	status, err := c.statusFetcher.Fetch(taskID)
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch task status: %w", err)
	}

	switch status {
	case taskstatus.Closed:
		return checkererrors.IssueClosedErr(filename, line, linecnt), nil
	case taskstatus.NonExistent:
		return checkererrors.IssueNonExistentErr(filename, line, linecnt), nil
	}

	return nil, nil
}
