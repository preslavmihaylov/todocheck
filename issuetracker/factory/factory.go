package factory

import (
	"errors"

	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
	"github.com/preslavmihaylov/todocheck/issuetracker/internal/github"
	"github.com/preslavmihaylov/todocheck/issuetracker/internal/gitlab"
	"github.com/preslavmihaylov/todocheck/issuetracker/internal/jira"
	"github.com/preslavmihaylov/todocheck/issuetracker/internal/pivotaltracker"
	"github.com/preslavmihaylov/todocheck/issuetracker/internal/redmine"
)

// NewIssueTrackerFrom is a static factory method for creating an issuetracker.IssueTracker instance based on the chosen issue tracker type
// in the configuration
func NewIssueTrackerFrom(issueTrackerType config.IssueTracker, authCfg *config.Auth, origin string) (issuetracker.IssueTracker, error) {
	switch issueTrackerType {
	case config.IssueTrackerGithub:
		return &github.IssueTracker{Origin: origin, AuthCfg: authCfg}, nil
	case config.IssueTrackerJira:
		return &jira.IssueTracker{Origin: origin, AuthCfg: authCfg}, nil
	case config.IssueTrackerGitlab:
		return &gitlab.IssueTracker{Origin: origin, AuthCfg: authCfg}, nil
	case config.IssueTrackerRedmine:
		return &redmine.IssueTracker{Origin: origin, AuthCfg: authCfg}, nil
	case config.IssueTrackerPivotal:
		return &pivotaltracker.IssueTracker{Origin: origin, AuthCfg: authCfg}, nil
	}

	return nil, errors.New("unknown issue tracker " + string(issueTrackerType))
}
