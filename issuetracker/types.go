package issuetracker

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker/models"
)

// TaskFor gets the corresponding task model, based on the issue tracker type
func TaskFor(issueTracker config.IssueTracker) models.Task {
	switch issueTracker {
	case config.IssueTrackerJira:
		return &models.JiraTask{}
	case config.IssueTrackerGithub:
		return &models.GithubTask{}
	case config.IssueTrackerGitlab:
		return &models.GitlabTask{}
	case config.IssueTrackerPivotal:
		return &models.PivotalTrackerTask{}
	default:
		return nil
	}
}

// BaseURLFor returns the task-fetching base url given the issue tracker type and the site origin
func BaseURLFor(issueTracker config.IssueTracker, origin string) (string, error) {
	switch issueTracker {
	case config.IssueTrackerJira:
		return origin + "/rest/api/latest/issue/", nil
	case config.IssueTrackerGithub:
		tokens := common.RemoveEmptyTokens(strings.Split(origin, "/"))
		if tokens[0] == "github.com" {
			tokens = append([]string{"https:"}, tokens...)
		}

		scheme, owner, repo := tokens[0], tokens[2], tokens[3]
		return fmt.Sprintf("%s//api.github.com/repos/%s/%s/issues/", scheme, owner, repo), nil
	case config.IssueTrackerGitlab:
		tokens := common.RemoveEmptyTokens(strings.Split(origin, "/"))
		if tokens[0] == "gitlab.com" {
			tokens = append([]string{"https:"}, tokens...)
		}

		scheme, owner, repo := tokens[0], tokens[2], tokens[3]
		urlEncodedProject := url.QueryEscape(fmt.Sprintf("%s/%s", owner, repo))
		return fmt.Sprintf("%s//gitlab.com/api/v4/projects/%s/issues/", scheme, urlEncodedProject), nil
	case config.IssueTrackerPivotal:
		tokens := common.RemoveEmptyTokens(strings.Split(origin, "/"))
		if tokens[0] == "pivotaltracker.com" {
			tokens = append([]string{"https:"}, tokens...)
		}

		scheme, project := tokens[0], tokens[3]
		if tokens[2] == "n" {
			project = tokens[4]
		}

		return fmt.Sprintf("%s//www.pivotaltracker.com/services/v5/projects/%s/stories/", scheme, project), nil
	default:
		return "", errors.New("unknown issue tracker " + string(issueTracker))
	}
}
