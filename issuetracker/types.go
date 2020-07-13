package issuetracker

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/issuetracker/models"
)

// Type of issue tracker enum
type Type string

// Issue tracker types
const (
	Invalid      Type = ""
	Jira              = "JIRA"
	GithubPublic      = "GITHUB_PUBLIC"
	Gitlab            = "GITLAB"
)

// TaskFor gets the corresponding task model, based on the issue tracker type
func TaskFor(issueTracker Type) models.Task {
	switch issueTracker {
	case Jira:
		return &models.JiraTask{}
	case GithubPublic:
		return &models.GithubTask{}
	case Gitlab:
		return &models.GitlabTask{}
	default:
		return nil
	}
}

// BaseURLFor returns the task-fetching base url given the issue tracker type and the site origin
func BaseURLFor(issueTracker Type, origin string) (string, error) {
	switch issueTracker {
	case Jira:
		return origin + "/rest/api/latest/issue/", nil
	case GithubPublic:
		tokens := common.RemoveEmptyTokens(strings.Split(origin, "/"))
		if tokens[0] == "github.com" {
			tokens = append([]string{"https:"}, tokens...)
		}

		scheme, owner, repo := tokens[0], tokens[2], tokens[3]
		return fmt.Sprintf("%s//api.github.com/repos/%s/%s/issues/", scheme, owner, repo), nil
	case Gitlab:
		tokens := common.RemoveEmptyTokens(strings.Split(origin, "/"))
		if tokens[0] == "gitlab.com" {
			tokens = append([]string{"https:"}, tokens...)
		}

		scheme, owner, repo := tokens[0], tokens[2], tokens[3]
		urlEncodedProject := url.QueryEscape(fmt.Sprintf("%s/%s", owner, repo))
		return fmt.Sprintf("%s//gitlab.com/api/v4/projects/%s/issues/", scheme, urlEncodedProject), nil
	default:
		return "", errors.New("unknown issue tracker type")
	}
}

// FromString converts a string-encoded issue tracker type to the correct type
func FromString(str string) Type {
	switch str {
	case Jira:
		return Jira
	case GithubPublic:
		return GithubPublic
	case Gitlab:
		return Gitlab
	default:
		return Invalid
	}
}
