package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// ErrNotFound when the specified file is not found
var ErrNotFound = errors.New("file not found")

// Local todocheck configuration struct definition
type Local struct {
	Origin           string `yaml:"origin"`
	IssueTrackerType string `yaml:"issue_tracker"`
	Auth             *Auth  `yaml:"auth"`
}

// NewLocal configuration from a given file path
func NewLocal(filepath string) (*Local, error) {
	if !exists(filepath) {
		return nil, ErrNotFound
	}

	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("couldn't open %s file (%s): %w", DefaultLocal, filepath, err)
	}

	var cfg *Local
	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s (%s): %w", DefaultLocal, filepath, err)
	}

	return cfg, nil
}

func exists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
