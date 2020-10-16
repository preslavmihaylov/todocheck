package config

import "regexp"

// IssueTracker enum
type IssueTracker string

// Issue tracker types
const (
	IssueTrackerInvalid IssueTracker = ""
	IssueTrackerJira                 = "JIRA"
	IssueTrackerGithub               = "GITHUB"
	IssueTrackerGitlab               = "GITLAB"
	IssueTrackerPivotal              = "PIVOTAL_TRACKER"
	IssueTrackerRedmine              = "REDMINE"
)

var validIssueTrackers = []IssueTracker{
	IssueTrackerJira,
	IssueTrackerGithub,
	IssueTrackerGitlab,
	IssueTrackerPivotal,
	IssueTrackerRedmine,
}

var originPatterns = map[IssueTracker]*regexp.Regexp{
	IssueTrackerJira:    regexp.MustCompile(`^(https?://)?[a-zA-Z0-9\-]+(\.[a-zA-Z0-9]+)+(:[0-9]+)?$`),
	IssueTrackerGithub:  regexp.MustCompile(`^(https?://)?(www\.)?github\.com/[\w-]+/[\w-]+`),
	IssueTrackerGitlab:  regexp.MustCompile(`^(https?://)?[a-zA-Z0-9\-]+(\.[a-zA-Z0-9]+)+(:[0-9]+)?/[\w-]+/[\w-]+$`),
	IssueTrackerPivotal: regexp.MustCompile(`^(https?://)?(www\.)?pivotaltracker\.com/n/projects/[0-9]+`),
	IssueTrackerRedmine: regexp.MustCompile(`^(https?://)?[a-zA-Z0-9\-]+(\.[a-zA-Z0-9]+)+(:[0-9]+)?$`),
}

// IsValid checks if the given issue tracker is among the valid enum values
func (it IssueTracker) IsValid() bool {
	for _, other := range validIssueTrackers {
		if it == other {
			return true
		}
	}

	return false
}

// IsValidOrigin checks if the given origin is among the valid patterns for the given issue tracker
func (it IssueTracker) IsValidOrigin(origin string) bool {
	pattern, ok := originPatterns[it]
	if !ok || !pattern.MatchString(origin) {
		return false
	}

	return true
}
