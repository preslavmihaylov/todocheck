package authtokens

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

var (
	// DefaultConfigPermissions the config file is created with
	DefaultConfigPermissions = os.FileMode(0700)
)

// DefaultConfigFile where auth tokens are stored by default
func DefaultConfigFile() string {
	dir, err := homedir.Dir()
	if err != nil {
		panic("couldn't read user home directory: " + err.Error())
	}

	return dir + "/.todocheck/authtokens.yaml"
}

// Config for storing user tokens for all todocheck origins
type Config struct {
	Tokens map[string]string `yaml:"tokens"`
}

// FromFile extracts the tokens configuration from the given file
func FromFile(filename string) (*Config, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("couldn't open %s: %w", filename, err)
	}

	var cfg *Config
	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s: %w", filename, err)
	}

	return cfg, nil
}

// Save the tokens configuration into the given file with default permissions
func (cfg *Config) Save(filename string) error {
	return cfg.SaveWithPerms(filename, DefaultConfigPermissions)
}

// SaveWithPerms the tokens configuration into the given file with given permissions
func (cfg *Config) SaveWithPerms(filename string, perms os.FileMode) error {
	bs, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal tokens config: %w", err)
	}

	err = ioutil.WriteFile(filename, bs, perms)
	if err != nil {
		return fmt.Errorf("failed to save marshaled tokens config: %w", err)
	}

	return nil
}

// CreateIfNotExists creates a new tokens config file if one is not present already
func CreateIfNotExists(filename string, perms os.FileMode) (*Config, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return FromFile(filename)
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("couldn't stat file %s: %w", filename, err)
	}

	dir := filepath.Dir(filename)
	err = os.MkdirAll(dir, perms)
	if err != nil {
		return nil, fmt.Errorf("couldn't mkdir %s: %w", dir, err)
	}

	emptyCfg := emptyConfig()
	err = emptyCfg.SaveWithPerms(filename, perms)
	if err != nil {
		return nil, fmt.Errorf("couldn't create tokens configuration file %s: %w", filename, err)
	}

	return FromFile(filename)
}

func emptyConfig() *Config {
	return &Config{
		Tokens: map[string]string{},
	}
}
