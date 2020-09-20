package config

import (
	"fmt"
	"io/ioutil"
	"net/url"
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

var (
	windowsAbsolutePathPattern = regexp.MustCompile("^[A-Z]{1}:")
	gitRemoteOriginPattern     = regexp.MustCompile(`(?Um)url\s=\s\w+(://|@)(?P<origin>(?P<host>.+)?(:|/).+)(\.git)?$`)
)

var originPatterns = map[IssueTracker]*regexp.Regexp{
	IssueTrackerJira:    regexp.MustCompile(`^(https?://)?[a-zA-Z0-9\-]+(\.[a-zA-Z0-9]+)+(:[0-9]+)?$`),
	IssueTrackerGithub:  regexp.MustCompile(`^(https?://)?(www\.)?github\.com/[\w-]+/[\w-]+`),
	IssueTrackerGitlab:  regexp.MustCompile(`^(https?://)?[a-zA-Z0-9\-]+(\.[a-zA-Z0-9]+)+(:[0-9]+)?/[\w-]+/[\w-]+$`),
	IssueTrackerPivotal: regexp.MustCompile(`^(https?://)?(www\.)?pivotaltracker\.com/n/projects/[0-9]+`),
	IssueTrackerRedmine: regexp.MustCompile(`^(https?://)?[a-zA-Z0-9\-]+(\.[a-zA-Z0-9]+)+(:[0-9]+)?$`),
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

	var (
		cfg *Local
		err error
	)

	if exists(cfgPath) {
		cfg, err = fromFile(cfgPath)
		if err != nil {
			return nil, err
		}
	} else {
		cfg, err = autoDetect(basepath)
		if err != nil {
			return nil, fmt.Errorf("file %s not found: unable to automatically detect issue tracker: %w", cfgPath, err)
		}
	}

	cfg.Auth.TokensCache = prependBasepath(cfg.Auth.TokensCache, basepath)

	prependDoublestarGlob(cfg.IgnoredPaths, basepath)
	trimTrailingSlashesFromDirs(cfg.IgnoredPaths)
	removeCurrentDirReference(cfg.IgnoredPaths)

	return cfg, nil
}

func fromFile(cfgPath string) (*Local, error) {
	bs, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't open local configuration (%s): %w", cfgPath, err)
	}

	cfg := &Local{Auth: defaultAuthCfg()}
	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal local configuration (%s): %w", cfgPath, err)
	}

	return cfg, nil
}

func autoDetect(basepath string) (*Local, error) {
	bs, err := ioutil.ReadFile(basepath + "/.git/config")
	if err != nil {
		return nil, err
	}

	result := map[string]string{}
	match := gitRemoteOriginPattern.FindStringSubmatch(string(bs))

	for i, group := range gitRemoteOriginPattern.SubexpNames() {
		result[group] = match[i]
	}

	var issueTracker IssueTracker

	switch result["host"] {
	case "github.com":
		issueTracker = IssueTrackerGithub
	case "gitlab.com":
		issueTracker = IssueTrackerGitlab
	default:
		return nil, fmt.Errorf("unable to auto-detect issue tracker")
	}

	// Since origin urls can be found in both formats of HTTP based URLs and SSH URIs,
	// it's necessary to replace colon with slash to convert it to a valid HTTP URL.
	// Example: git@github:username/repo.git, https://github.com/username/repo.git
	origin := strings.Replace(result["origin"], ":", "/", 1)

	fmt.Printf("Detected %q as issue tracker since no config file was found.\n", origin)

	return &Local{
		Auth:         defaultAuthCfg(),
		IssueTracker: issueTracker,
		Origin:       origin,
	}, nil
}

// Validate validates the values of given configuration
func (l *Local) Validate() []error {
	var errors []error

	if err := l.validateIssueTracker(); err != nil {
		errors = append(errors, err)
	}

	if err := l.validateAuthOfflineURL(); err != nil {
		errors = append(errors, err)
	}

	if l.IssueTracker != "" {
		pattern, ok := originPatterns[l.IssueTracker]
		if !ok || !pattern.MatchString(l.Origin) {
			errors = append(errors, fmt.Errorf("%s is not a valid origin for issue tracker %s", l.Origin, l.IssueTracker))
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

func (l *Local) validateAuthOfflineURL() error {
	if _, err := url.ParseRequestURI(l.Auth.OfflineURL); l.Auth.Type == AuthTypeOffline && err != nil {
		return fmt.Errorf("invalid offline URL: %q", l.Auth.OfflineURL)
	}

	return nil
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
