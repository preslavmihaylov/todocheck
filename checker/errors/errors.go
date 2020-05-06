package errors

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

// MalformedTODOErr when todo is not properly formatted
func MalformedTODOErr(filename, line string, linecnt int) error {
	msg := color.RedString("ERROR: Malformed todo.\n") +
		fmt.Sprintf("%s:%d: %s", filename, linecnt, line) +
		color.CyanString("\t> TODO should match pattern - \"// TODO [TASK_ID]: comment\"\n")

	return errors.New(msg)
}

// IssueClosedErr when referenced todo issue is closed
func IssueClosedErr(filename, line string, linecnt int) error {
	msg := color.RedString("ERROR: Issue is closed.\n") +
		fmt.Sprintf("%s:%d: %s", filename, linecnt, line)

	return errors.New(msg)
}

// IssueNonExistentErr when referenced todo issue doesn't exist
func IssueNonExistentErr(filename, line string, linecnt int) error {
	msg := color.RedString("ERROR: Issue doesn't exist.\n") +
		fmt.Sprintf("%s:%d: %s", filename, linecnt, line)

	return errors.New(msg)
}
