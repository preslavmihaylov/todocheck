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

// TodocheckRC configuration
type TodocheckRC struct {
	Origin string `yaml:"origin"`
	Type   string `yaml:"type"`
}

// NewTodocheckRC configuration from a given file path
func NewTodocheckRC(filepath string) (*TodocheckRC, error) {
	if !exists(filepath) {
		return nil, ErrNotFound
	}

	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("couldn't open .todocheckrc file (%s): %w", filepath, err)
	}

	var cfg *TodocheckRC
	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal .todocheckrc (%s): %w", filepath, err)
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
