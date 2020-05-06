package checker

import (
	"fmt"
	"regexp"

	checkererrors "github.com/preslavmihaylov/todocheck/checker/errors"
	"github.com/preslavmihaylov/todocheck/taskstatus"
)

// Checker for todo lines
type Checker struct {
	statusFetcher *taskstatus.Fetcher
}

// New checker
func New(statusFetcher *taskstatus.Fetcher) *Checker {
	return &Checker{statusFetcher}
}

// IsTODO line, without performing any validity checks
func (c *Checker) IsTODO(line string) bool {
	isTodoLine, err := regexp.MatchString("^\\s*//\\s*TODO[ :]", line)
	if err != nil {
		panic(err)
	}

	return isTodoLine
}

// Check if todo line is valid
func (c *Checker) Check(filename, line string, linecnt int) (error, error) {
	if isMalformed(line) {
		return checkererrors.MalformedTODOErr(filename, line, linecnt), nil
	}

	status, err := c.statusFetcher.Fetch(extractTaskID(line))
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

func isMalformed(todo string) bool {
	match, err := regexp.MatchString("\\s*// TODO [a-zA-Z0-9\\-]+:.*", todo)
	if err != nil {
		panic(err)
	}

	return !match
}

func extractTaskID(line string) string {
	pattern := regexp.MustCompile("^\\s*// TODO ([a-zA-Z0-9\\-]+):.*")

	return pattern.FindStringSubmatch(line)[1]
}
