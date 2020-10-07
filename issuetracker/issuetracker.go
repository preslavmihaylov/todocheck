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

// ErrUnsupportedHealthCheck is returned when the health check doesn't support the given issue tracker
var ErrUnsupportedHealthCheck = errors.New("unsupported issue tracker for health check")

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
	case config.IssueTrackerRedmine:
		return &models.RedmineTask{}
	default:
		return nil
	}
}

// TaskURLSuffixFor the given issue tracker. Returns the appropriate task ID to append to rest api request.
// In most use-cases, the taskID is returned as-is as it is simply appended to the issue tracker URL origin (e.g. issuetracker.com/issues/{taskID}.
// However, some issue trackers might expect the taskID to be appended in a special format
func TaskURLSuffixFor(taskID string, issueTracker config.IssueTracker) string {
	switch issueTracker {
	case config.IssueTrackerRedmine:
		return taskID + ".json"
	case config.IssueTrackerPivotal:
		if len(taskID) > 0 && taskID[0] == '#' {
			return taskID[1:]
		}

		return taskID
	default:
		return taskID
	}
}

// BaseURLFor returns the task-fetching base url given the issue tracker type and the site origin
func BaseURLFor(issueTracker config.IssueTracker, origin string) (string, error) {
	switch issueTracker {
	case config.IssueTrackerJira:
		return fmt.Sprintf("%s/rest/api/latest/issue/", origin), nil
	case config.IssueTrackerGithub:
		scheme, owner, repo := parseGithubDetails(origin)
		return fmt.Sprintf("%s//api.github.com/repos/%s/%s/issues/", scheme, owner, repo), nil
	case config.IssueTrackerGitlab:
		tokens := common.RemoveEmptyTokens(strings.Split(origin, "/"))
		if tokens[0] != "http:" && tokens[0] != "https:" {
			tokens = append([]string{"https:"}, tokens...)
		}

		scheme, host, owner, repo := tokens[0], tokens[1], tokens[2], tokens[3]
		urlEncodedProject := url.QueryEscape(fmt.Sprintf("%s/%s", owner, repo))
		return fmt.Sprintf("%s//%s/api/v4/projects/%s/issues/", scheme, host, urlEncodedProject), nil
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
	case config.IssueTrackerRedmine:
		return fmt.Sprintf("%s/issues/", origin), nil
	default:
		return "", errors.New("unknown issue tracker " + string(issueTracker))
	}
}

// HealthcheckURL returns the health check base url given the issue tracker type and the site origin
func HealthcheckURL(issueTracker config.IssueTracker, origin string) (string, error) {
	switch issueTracker {
	case config.IssueTrackerGithub:
		scheme, owner, repo := parseGithubDetails(origin)
		return fmt.Sprintf("%s//api.github.com/repos/%s/%s", scheme, owner, repo), nil
	default:
		return "", ErrUnsupportedHealthCheck
	}
}

func parseGithubDetails(origin string) (scheme, owner, repo string) {
	tokens := common.RemoveEmptyTokens(strings.Split(origin, "/"))
	if tokens[0] == "github.com" {
		tokens = append([]string{"https:"}, tokens...)
	}

	scheme, owner, repo = tokens[0], tokens[2], tokens[3]
	return
}
