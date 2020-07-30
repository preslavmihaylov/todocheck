package config

import (
	"fmt"
	"io/ioutil"
	"os"
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
	return path[0] != '/' && path[0] != '~'
}
