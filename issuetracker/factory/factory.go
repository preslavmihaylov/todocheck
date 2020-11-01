package factory

import (
	"errors"

	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
	"github.com/preslavmihaylov/todocheck/issuetracker/github"
	"github.com/preslavmihaylov/todocheck/issuetracker/gitlab"
	"github.com/preslavmihaylov/todocheck/issuetracker/jira"
	"github.com/preslavmihaylov/todocheck/issuetracker/pivotaltracker"
	"github.com/preslavmihaylov/todocheck/issuetracker/redmine"
)

// NewIssueTrackerFrom is a static factory method for creating an issuetracker.IssueTracker instance based on the chosen issue tracker type
// in the configuration
func NewIssueTrackerFrom(issueTrackerType config.IssueTracker, origin string) (issuetracker.IssueTracker, error) {
	switch issueTrackerType {
	case config.IssueTrackerGithub:
		return &github.IssueTracker{Origin: origin}, nil
	case config.IssueTrackerJira:
		return &jira.IssueTracker{Origin: origin}, nil
	case config.IssueTrackerGitlab:
		return &gitlab.IssueTracker{Origin: origin}, nil
	case config.IssueTrackerRedmine:
		return &redmine.IssueTracker{Origin: origin}, nil
	case config.IssueTrackerPivotal:
		return &pivotaltracker.IssueTracker{Origin: origin}, nil
	}

	return nil, errors.New("unknown issue tracker " + string(issueTrackerType))
}
