package factory

import (
	"errors"

	"github.com/preslavmihaylov/todocheck/issuetracker/internal/azureboards"

	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/issuetracker"
	"github.com/preslavmihaylov/todocheck/issuetracker/internal/github"
	"github.com/preslavmihaylov/todocheck/issuetracker/internal/gitlab"
	"github.com/preslavmihaylov/todocheck/issuetracker/internal/jira"
	"github.com/preslavmihaylov/todocheck/issuetracker/internal/pivotaltracker"
	"github.com/preslavmihaylov/todocheck/issuetracker/internal/redmine"
	"github.com/preslavmihaylov/todocheck/issuetracker/internal/youtrack"
)

// NewIssueTrackerFrom is a static factory method for creating an issuetracker.IssueTracker instance based on the chosen issue tracker type
// in the configuration
func NewIssueTrackerFrom(issueTrackerType config.IssueTracker, authCfg *config.Auth, origin string) (issuetracker.IssueTracker, error) {
	switch issueTrackerType {
	case config.IssueTrackerGithub:
		return github.New(origin, authCfg)
	case config.IssueTrackerJira:
		return jira.New(origin, authCfg)
	case config.IssueTrackerGitlab:
		return gitlab.New(origin, authCfg)
	case config.IssueTrackerRedmine:
		return redmine.New(origin, authCfg)
	case config.IssueTrackerPivotal:
		return pivotaltracker.New(origin, authCfg)
	case config.IssueTrackerYoutrack:
		return youtrack.New(origin, authCfg)
	case config.IssueTrackerAzure:
		return azureboards.NewIssueTrackerAzure(origin, authCfg)
	}

	return nil, errors.New("unknown issue tracker " + string(issueTrackerType))
}
