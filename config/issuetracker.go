package config

import (
	"regexp"
)

// IssueTracker enum
type IssueTracker string

// Issue tracker types
const (
	IssueTrackerInvalid  IssueTracker = ""
	IssueTrackerJira                  = "JIRA"
	IssueTrackerGithub                = "GITHUB"
	IssueTrackerGitlab                = "GITLAB"
	IssueTrackerPivotal               = "PIVOTAL_TRACKER"
	IssueTrackerRedmine               = "REDMINE"
	IssueTrackerYoutrack              = "YOUTRACK"
)

var ValidIssueTrackerAuthTypes = map[IssueTracker][]AuthType{
	IssueTrackerGithub:   {AuthTypeNone, AuthTypeAPIToken},
	IssueTrackerGitlab:   {AuthTypeNone, AuthTypeAPIToken},
	IssueTrackerPivotal:  {AuthTypeNone, AuthTypeAPIToken},
	IssueTrackerRedmine:  {AuthTypeNone, AuthTypeAPIToken},
	IssueTrackerJira:     {AuthTypeNone, AuthTypeOffline},
	IssueTrackerYoutrack: {AuthTypeNone, AuthTypeAPIToken},
}

var validIssueTrackers = []IssueTracker{
	IssueTrackerJira,
	IssueTrackerGithub,
	IssueTrackerGitlab,
	IssueTrackerPivotal,
	IssueTrackerRedmine,
	IssueTrackerYoutrack,
}

var originPatterns = map[IssueTracker]*regexp.Regexp{
	IssueTrackerJira:     regexp.MustCompile(`^(https?://)?[a-zA-Z0-9\-]+(\.[a-zA-Z0-9]+)+(:[0-9]+)?$`),
	IssueTrackerGithub:   regexp.MustCompile(`^(https?://)?(www\.)?github\.com/[\w-]+/[\w-]+`),
	IssueTrackerGitlab:   regexp.MustCompile(`^(https?://)?[a-zA-Z0-9\-]+(\.[a-zA-Z0-9]+)+(:[0-9]+)?/[\w-]+/[\w-]+$`),
	IssueTrackerPivotal:  regexp.MustCompile(`^(https?://)?(www\.)?pivotaltracker\.com/n/projects/[0-9]+`),
	IssueTrackerRedmine:  regexp.MustCompile(`^(https?://)?[a-zA-Z0-9\-]+(\.[a-zA-Z0-9]+)+(:[0-9]+)?$`),
	IssueTrackerYoutrack: regexp.MustCompile(`^(https?://)?(www\.)?[0-9A-z-]{2,63}\.myjetbrains\.com/?.*`),
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

// IsValidAuthType checks if the given auth type is among the valid auth types for the given issue tracker
func (it IssueTracker) IsValidAuthType(authType AuthType) bool {
	for _, validType := range ValidIssueTrackerAuthTypes[it] {
		if authType == validType {
			return true
		}
	}
	return false
}
