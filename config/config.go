package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// DefaultLocal contains the default filepath to the local todocheck config for the current repository
const DefaultLocal = ".todocheck.yaml"

var (
	windowsAbsolutePathPattern = regexp.MustCompile("^[A-Z]{1}:")
	gitRemoteOriginPattern     = regexp.MustCompile(`(?Um)url\s=\s\w+(://|@)(?P<origin>(?P<host>.+)?(:|/).+)(\.git)?$`)
)

// Local todocheck configuration struct definition
type Local struct {
	Origin               string       `yaml:"origin"`
	IssueTracker         IssueTracker `yaml:"issue_tracker"`
	IgnoredPaths         []string     `yaml:"ignored"`
	CustomTodos          []string     `yaml:"custom_todos"`
	Auth                 *Auth        `yaml:"auth"`
	MatchCaseInsensitive bool         `yaml:"match_case_insensitive"`
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

	cfg.CustomTodos = decodeEscapedReservedCharacters(cfg.CustomTodos)
	cfg.CustomTodos = addDefaultFormatIfMissing(cfg.CustomTodos)

	return cfg, nil
}

func fromFile(cfgPath string) (*Local, error) {
	bs, err := os.ReadFile(cfgPath)
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
	bs, err := os.ReadFile(basepath + "/.git/config")
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

// addDefaultFormatIfMissing trying find default TODO string and adding it if not exists
func addDefaultFormatIfMissing(todos []string) []string {
	var isExists bool
	for _, v := range todos {
		if v == "TODO" {
			isExists = true
			continue
		}
	}

	if !isExists {
		todos = append(todos, "TODO")
	}

	return todos
}

func decodeEscapedReservedCharacters(slice []string) []string {
	// Remove leading escaping "\" for reserved strings.
	// "@", "`" in YAML is reserved indicators
	//  https://yaml.org/spec/1.2/spec.html#id2772075
	for i, v := range slice {
		if strings.HasPrefix(v, "\\@") || strings.HasPrefix(v, "\\`") {
			slice[i] = v[1:]
		}
	}
	return slice
}
