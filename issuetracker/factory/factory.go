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
		return github.New(authCfg, origin)
	case config.IssueTrackerJira:
		return jira.New(authCfg, origin)
	case config.IssueTrackerGitlab:
		return gitlab.New(authCfg, origin)
	case config.IssueTrackerRedmine:
		return redmine.New(authCfg, origin)
	case config.IssueTrackerPivotal:
		return pivotaltracker.New(authCfg, origin)
	}

	return nil, errors.New("unknown issue tracker " + string(issueTrackerType))
}
