package checker

import (
	"fmt"

	checkererrors "github.com/preslavmihaylov/todocheck/checker/errors"
	"github.com/preslavmihaylov/todocheck/fetcher"
	"github.com/preslavmihaylov/todocheck/matchers"
	"github.com/preslavmihaylov/todocheck/taskstatus"
)

// Checker for todo lines
type Checker struct {
	statusFetcher *fetcher.Fetcher
	todoMatcher   matchers.TodoMatcher
}

// New checker
func New(statusFetcher *fetcher.Fetcher) *Checker {
	return &Checker{statusFetcher, nil}
}

// SetMatcher sets the todo matcher
func (c *Checker) SetMatcher(todoMatcher matchers.TodoMatcher) {
	c.todoMatcher = todoMatcher
}

// IsTODO line, without performing any validity checks
func (c *Checker) IsTODO(line string) bool {
	return c.todoMatcher.IsMatch(line)
}

// Check if todo line is valid
func (c *Checker) Check(comment, filename string, lines []string, linecnt int) (error, error) {
	if !c.todoMatcher.IsValid(comment) {
		return checkererrors.MalformedTODOErr(filename, lines, linecnt), nil
	}

	taskID, err := c.todoMatcher.ExtractIssueRef(comment)
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
		return checkererrors.IssueClosedErr(filename, lines, linecnt), nil
	case taskstatus.NonExistent:
		return checkererrors.IssueNonExistentErr(filename, lines, linecnt), nil
	}

	return nil, nil
}
