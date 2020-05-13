package errors

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

// MalformedTODOErr when todo is not properly formatted
func MalformedTODOErr(filename string, lines []string, linecnt int) error {
	msg := color.RedString("ERROR: Malformed todo.\n") +
		printSourceLocation(filename, lines, linecnt) +
		color.CyanString("\t> TODO should match pattern - \"// TODO [TASK_ID]: comment\"\n")

	return errors.New(msg)
}

// IssueClosedErr when referenced todo issue is closed
func IssueClosedErr(filename string, lines []string, linecnt int) error {
	msg := color.RedString("ERROR: Issue is closed.\n") +
		printSourceLocation(filename, lines, linecnt)

	return errors.New(msg)
}

// IssueNonExistentErr when referenced todo issue doesn't exist
func IssueNonExistentErr(filename string, lines []string, linecnt int) error {
	msg := color.RedString("ERROR: Issue doesn't exist.\n") +
		printSourceLocation(filename, lines, linecnt)

	return errors.New(msg)
}

func printSourceLocation(filename string, lines []string, linecnt int) string {
	res := ""
	for i, line := range lines {
		res += fmt.Sprintf("%s:%d: %s", filename, linecnt+i, line)
	}

	return res
}
