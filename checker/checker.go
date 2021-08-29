package checker

import (
	"errors"
	"fmt"

	checkererrors "github.com/preslavmihaylov/todocheck/checker/errors"
	"github.com/preslavmihaylov/todocheck/fetcher"
	"github.com/preslavmihaylov/todocheck/issuetracker/taskstatus"
	"github.com/preslavmihaylov/todocheck/matchers"
)

// Checker for todo lines
type Checker struct {
	statusFetcher *fetcher.Fetcher
}

// New checker
func New(statusFetcher *fetcher.Fetcher) *Checker {
	return &Checker{statusFetcher}
}

// Check if todo line is valid
func (c *Checker) Check(
	matcher matchers.TodoMatcher, comment, filename string, lines []string, linecnt int,
) (*checkererrors.TODO, error) {
	if matcher == nil {
		return nil, errors.New("matcher is nil")
	}

	if !matcher.IsMatch(comment) {
		return nil, nil
	}

	if !matcher.IsValid(comment) {
		return checkererrors.MalformedTODOErr(filename, lines, linecnt), nil
	}

	taskID, err := matcher.ExtractIssueRef(comment)
	if err != nil {
		// should never happen after validating todo line
		panic("couldn't extract issue reference from a valid todo: " + err.Error())
	}

	status, err := c.statusFetcher.Fetch(taskID)
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch task status: %w", err)
	}

	// Extract message by slicing off the necessary tokens ("TODO {taskId}: ")
	colonPos := func() int {
		for i := range comment {
			if comment[i] == ':' {
				return i
			}
		}
		// Cannot happen as TODO format guarantees presence of colon
		return -1
	}()
	message := comment[colonPos+2:]

	switch status {
	case taskstatus.Closed:
		return checkererrors.IssueClosedErr(filename, lines, linecnt, taskID, message), nil
	case taskstatus.NonExistent:
		return checkererrors.IssueNonExistentErr(filename, lines, linecnt, taskID, message), nil
	}

	return nil, nil
}
