package issuetracker

import (
	"encoding/json"

	"github.com/preslavmihaylov/todocheck/testing/scenariobuilder/issuetracker/models"
)

// Type is an enum specifying the target mock issue tracker to use
type Type string

// Issue tracker types available for test scenarios
const (
	Jira Type = "Jira"
)

// Status is an enum specifying the expected issue status while building your test scenario
type Status string

// Possible issue statuses to specify for your test scenarios
const (
	StatusClosed Status = "Closed"
	StatusOpen   Status = "Open"
)

var trackerToIssuePath = map[Type]string{
	Jira: "/rest/api/latest/issue/",
}

// IssueURLFrom builds the appropriate expected issue url, given the issue tracker type & issue id
func IssueURLFrom(t Type, issue string) string {
	path, ok := trackerToIssuePath[t]
	if !ok {
		panic("unknown issue tracker received: " + string(t))
	}

	return path + issue
}

// BuildResponseFor given issue tracker type, issue ID and issue status
func BuildResponseFor(t Type, issue string, status Status) []byte {
	switch t {
	case Jira:
		return must(json.Marshal(&models.JiraTask{
			Fields: models.Fields{
				Status: models.Status{
					Name: string(status),
				},
			},
		}))
	default:
		panic("unknown issue tracker received: " + string(t))
	}
}

func must(res []byte, err error) []byte {
	if err != nil {
		panic("couldn't marshal response: " + err.Error())
	}

	return res
}
