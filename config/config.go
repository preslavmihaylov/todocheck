package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

// DefaultLocal contains the default filepath to the local todocheck config for the current repository
const DefaultLocal = ".todocheck.yaml"

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

var windowsAbsolutePathPattern = regexp.MustCompile("^[A-Z]{1}:")

var originPatterns = map[IssueTracker]*regexp.Regexp{
	IssueTrackerJira:    regexp.MustCompile(``),
	IssueTrackerGithub:  regexp.MustCompile(`^(https://)?(www\.)?github\.com/[a-zA-Z0-9\-_]+/[a-zA-Z0-9\-_]+`),
	IssueTrackerGitlab:  regexp.MustCompile(`^(https://)?(www\.)?gitlab\.com/[a-zA-Z0-9\-_]+/[a-zA-Z0-9\-_]+`),
	IssueTrackerPivotal: regexp.MustCompile(`^(https://)?(www\.)?pivotaltracker/n/projects/[0-9]+`),
	IssueTrackerRedmine: regexp.MustCompile(``),
}

// Local todocheck configuration struct definition
type Local struct {
	Origin       string       `yaml:"origin"`
	IssueTracker IssueTracker `yaml:"issue_tracker"`
	IgnoredPaths []string     `yaml:"ignored"`
	Auth         *Auth        `yaml:"auth"`
}

// NewLocal configuration from a given file path
func NewLocal(cfgPath, basepath string) (*Local, error) {
	if cfgPath == "" {
		cfgPath = basepath + "/" + DefaultLocal
	}

	if !exists(cfgPath) {
		return nil, fmt.Errorf("file %s not found", cfgPath)
	}

	bs, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't open local configuration (%s): %w", cfgPath, err)
	}

	cfg := &Local{Auth: defaultAuthCfg()}
	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal local configuration (%s): %w", cfgPath, err)
	}

	cfg.Auth.TokensCache = prependBasepath(cfg.Auth.TokensCache, basepath)

	prependDoublestarGlob(cfg.IgnoredPaths, basepath)
	trimTrailingSlashesFromDirs(cfg.IgnoredPaths)
	removeCurrentDirReference(cfg.IgnoredPaths)

	return cfg, nil
}

// Validate validates the values of given configuration
func (l *Local) Validate() []error {
	var errors []error

	if err := l.validateIssueTracker(); err != nil {
		errors = append(errors, err)
	}

	if 0 == len(errors) {
		// l.IssueTracker is sure to be in the map after validating above
		pattern := originPatterns[l.IssueTracker]
		if !pattern.MatchString(l.Origin) {
			errors = append(errors, fmt.Errorf("origin is not valid for issue tracker: %s", l.IssueTracker))
		}
	}

	return errors
}

func (l *Local) validateIssueTracker() error {
	for _, issueTracker := range validIssueTrackers {
		if l.IssueTracker == issueTracker {
			return nil
		}
	}
	return fmt.Errorf("invalid issue tracker: %q is not supported. the valid issue trackers are: %v",
		l.IssueTracker, validIssueTrackers)
}

func exists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func trimTrailingSlashesFromDirs(dirs []string) {
	for i, dir := range dirs {
		dirs[i] = strings.TrimRight(dir, "/")
	}
}

func prependDoublestarGlob(dirs []string, basepath string) {
	for i := range dirs {
		dirs[i] = "**/" + dirs[i]
	}
}

func prependBasepath(path, basepath string) string {
	if !isRelativePath(path) {
		return path
	}

	if basepath[len(basepath)-1] != '/' {
		basepath = basepath + "/"
	}

	return basepath + path
}

func removeCurrentDirReference(dirs []string) {
	for i := range dirs {
		if dirs[i][:2] == "./" {
			dirs[i] = dirs[i][2:]
		}
	}
}

func isRelativePath(path string) bool {
	return path[0] != '/' && path[0] != '~' && !windowsAbsolutePathPattern.MatchString(path)
}
