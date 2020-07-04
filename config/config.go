package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// DefaultLocal contains the default filepath to the local todocheck config for the current repository
const DefaultLocal = ".todocheck.yaml"

// ErrNotFound when the specified file is not found
var ErrNotFound = errors.New("file not found")

// Local todocheck configuration struct definition
type Local struct {
	Origin           string   `yaml:"origin"`
	IssueTrackerType string   `yaml:"issue_tracker"`
	IgnoredPaths     []string `yaml:"ignored"`
	Auth             *Auth    `yaml:"auth"`
}

// NewLocal configuration from a given file path
func NewLocal(cfgPath, basepath string) (*Local, error) {
	if !exists(cfgPath) {
		return nil, ErrNotFound
	}

	bs, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't open local configuration (%s): %w", cfgPath, err)
	}

	var cfg *Local
	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal local configuration (%s): %w", cfgPath, err)
	}

	trimTrailingSlashesFromDirs(cfg.IgnoredPaths)
	prependBasepath(cfg.IgnoredPaths, basepath)

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

func prependBasepath(dirs []string, basepath string) {
	if basepath[len(basepath)-1] != '/' {
		basepath = basepath + "/"
	}

	for i := range dirs {
		dirs[i] = basepath + dirs[i]
	}
}
