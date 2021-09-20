package errors

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
)

// TODOErrType is an enum representing the type of todo error
type TODOErrType string

// supported todo error types enum
const (
	TODOErrTypeMalformed        TODOErrType = "Malformed todo"
	TODOErrTypeIssueClosed      TODOErrType = "Issue is closed"
	TODOErrTypeNonExistentIssue TODOErrType = "Issue doesn't exist"
)

// TODO encapsulates the todo error information
type TODO struct {
	errType  TODOErrType
	filename string
	lines    []string
	linecnt  int
	metadata map[string]string
}

// ToJSON converts the todo error into json format
func (err *TODO) ToJSON() ([]byte, error) {
	res := &struct {
		Type     string            `json:"type"`
		Filename string            `json:"filename"`
		Line     int               `json:"line"`
		Message  string            `json:"message"`
		Metadata map[string]string `json:"metadata"`
	}{
		Type:     string(err.errType),
		Filename: err.filename,
		Line:     err.linecnt,
		Message:  "",
		Metadata: err.metadata,
	}

	if err.errType == TODOErrTypeMalformed {
		res.Message = "TODO should match pattern - TODO {task_id}:"
	}

	return json.Marshal(res)
}

func (err *TODO) Error() string {
	return err.String()
}

func (err *TODO) String() string {
	msg := color.RedString("ERROR: " + string(err.errType) + "\n")
	msg += printSourceLocation(err.filename, err.lines, err.linecnt)
	if err.errType == TODOErrTypeMalformed {
		msg += color.CyanString("\t> TODO should match pattern - TODO {task_id}:\n")
	}

	return msg
}

// MalformedTODOErr when todo is not properly formatted
func MalformedTODOErr(filename string, lines []string, linecnt int) *TODO {
	return &TODO{
		errType:  TODOErrTypeMalformed,
		filename: filename,
		lines:    lines,
		linecnt:  linecnt,
		metadata: make(map[string]string),
	}
}

// IssueClosedErr when referenced todo issue is closed
func IssueClosedErr(filename string, lines []string, linecnt int, issueID string) *TODO {
	return &TODO{
		errType:  TODOErrTypeIssueClosed,
		filename: filename,
		lines:    lines,
		linecnt:  linecnt,
		metadata: map[string]string{
			"issueID": issueID,
		},
	}
}

// IssueNonExistentErr when referenced todo issue doesn't exist
func IssueNonExistentErr(filename string, lines []string, linecnt int, issueID string) *TODO {
	return &TODO{
		errType:  TODOErrTypeNonExistentIssue,
		filename: filename,
		lines:    lines,
		linecnt:  linecnt,
		metadata: map[string]string{
			"issueID": issueID,
		},
	}
}

func printSourceLocation(filename string, lines []string, linecnt int) string {
	res := ""
	for i, line := range lines {
		res += fmt.Sprintf("%s:%d: %s", filename, linecnt+i, line)
	}

	return res
}
