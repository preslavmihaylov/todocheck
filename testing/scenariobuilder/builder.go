// Package scenariobuilder allows you to construct
// test scenarios which are based on executing the real todocheck binary
// against projects to use as scenarios.
package scenariobuilder

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"regexp"

	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/testing/scenariobuilder/issuetracker"
)

type teardownFunc func()
type validateFunc func() error

// TodocheckScenario encapsulates the scenario the program is expected to execute.
// This let's you specify what are the program inputs & what is the expected outputs.
type TodocheckScenario struct {
	binaryLoc              string
	basepath               string
	cfgPath                string
	testCfgPath            string
	cfg                    *config.Local
	expectedAuthToken      string
	userOfflineToken       string
	deleteTokensCacheAfter bool
	expectedExitCode       int
	issueTracker           issuetracker.Type
	issues                 map[string]issuetracker.Status
	envVariables           map[string]string
	todoErrScenarios       []*TodoErrScenario
}

// NewScenario to execute against the todocheck program
func NewScenario() *TodocheckScenario {
	return &TodocheckScenario{
		binaryLoc:        "./todocheck",
		basepath:         ".",
		issues:           map[string]issuetracker.Status{},
		envVariables:     map[string]string{},
		expectedExitCode: 0,
	}
}

// WithBinary let's you specify the location of the todocheck binary to test with
func (s *TodocheckScenario) WithBinary(binaryLoc string) *TodocheckScenario {
	s.binaryLoc = binaryLoc
	return s
}

// WithBasepath let's you specify the --basepath flag passed to the program
func (s *TodocheckScenario) WithBasepath(basepath string) *TodocheckScenario {
	s.basepath = basepath
	return s
}

// WithConfig let's you specify the --config flag passed to the program.
// By default, it is also used for the test environment config.
// If you want to specify a different test environment config, see WithTestEnvConfig
func (s *TodocheckScenario) WithConfig(cfgPath string) *TodocheckScenario {
	s.cfgPath = cfgPath
	if s.testCfgPath == "" {
		s.testCfgPath = cfgPath
	}

	return s
}

// WithTestEnvConfig let's you specify a configuration used for the test environment, which can be different from the --config flag passed to todocheck
func (s *TodocheckScenario) WithTestEnvConfig(cfgPath string) *TodocheckScenario {
	s.testCfgPath = cfgPath
	return s
}

// WithIssueTracker let's you specify what issue tracker to execute the scenario with
func (s *TodocheckScenario) WithIssueTracker(issueTracker issuetracker.Type) *TodocheckScenario {
	s.issueTracker = issueTracker
	return s
}

// RequireAuthToken on each issue lookup
func (s *TodocheckScenario) RequireAuthToken(token string) *TodocheckScenario {
	s.expectedAuthToken = token
	return s
}

// WithEnvVariable sets a custom environment variable for the period of test execution
func (s *TodocheckScenario) WithEnvVariable(key, value string) *TodocheckScenario {
	s.envVariables[key] = value
	return s
}

// WithIssue sets up a mock issue in your issue tracker with the given status
func (s *TodocheckScenario) WithIssue(issueID string, status issuetracker.Status) *TodocheckScenario {
	s.issues[issueID] = status
	return s
}

// SetOfflineTokenWhenRequested by sending the specified token to the program's standard input
func (s *TodocheckScenario) SetOfflineTokenWhenRequested(token string) *TodocheckScenario {
	s.userOfflineToken = token
	return s
}

// DeleteTokensCacheAfter the test completes. The tokens cache is derived from the given config file
func (s *TodocheckScenario) DeleteTokensCacheAfter() *TodocheckScenario {
	s.deleteTokensCacheAfter = true
	return s
}

// ExpectTodoErr appends a new todo err scenario to expect from the program execution
func (s *TodocheckScenario) ExpectTodoErr(sc *TodoErrScenario) *TodocheckScenario {
	s.expectedExitCode = 2
	s.todoErrScenarios = append(s.todoErrScenarios, sc)
	return s
}

// ExpectExecutionError on program execution
func (s *TodocheckScenario) ExpectExecutionError() *TodocheckScenario {
	s.expectedExitCode = 1
	return s
}

// Run sets up the environment & executes the configured scenario
func (s *TodocheckScenario) Run() error {
	var err error
	s.cfg, err = config.NewLocal(s.testCfgPath, s.basepath)
	if err != nil {
		return fmt.Errorf("couldn't initialize todocheck config: %w", err)
	}

	cmd := exec.Command(s.binaryLoc, "--basepath", s.basepath, "--config", s.cfgPath)
	teardown, err := s.setupTestEnvironment(cmd)
	if err != nil {
		return fmt.Errorf("couldn't setup test environment: %s", err)
	}
	defer teardown()

	var stdin, stdout, stderr bytes.Buffer
	if s.userOfflineToken != "" {
		stdin.Write([]byte(s.userOfflineToken + "\n"))
	}

	cmd.Stdin = &stdin
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	fmt.Println("(standard output follows. Standard output is ignored & not validated...)")
	fmt.Println(stdout.String())
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); !ok || exitError.ExitCode() != s.expectedExitCode {
			return fmt.Errorf("program exited with an unexpected exit code: %s", err)
		}
	}

	validateFuncs := []validateFunc{
		validateTodoErrs(stderr.String(), s.todoErrScenarios),
		validateAuthTokensCache(s.cfg.Auth.TokensCache, s.cfg.Auth.OfflineURL, s.expectedAuthToken),
	}

	if s.expectedExitCode == 1 {
		// skip todo err validation on execution error
		validateFuncs = validateFuncs[1:]
	}

	return validationPipeline(validateFuncs...)
}

func (s *TodocheckScenario) setupTestEnvironment(cmd *exec.Cmd) (teardownFunc, error) {
	s.setupEnvironmentVariables(cmd, s.envVariables)
	mockSrv := s.setupMockIssueTrackerServer()
	teardownIssueTrackerCfg, err := setupMockIssueTrackerCfg(s.testCfgPath, mockSrv.URL)
	if err != nil {
		return nil, fmt.Errorf("couldn't setup mock issue tracker: %w", err)
	}

	return func() {
		if s.deleteTokensCacheAfter {
			defer deleteTokensCache(s.cfg.Auth.TokensCache)
		}

		defer teardownIssueTrackerCfg()
		defer mockSrv.Close()
	}, nil
}

func (s *TodocheckScenario) setupEnvironmentVariables(cmd *exec.Cmd, envVariables map[string]string) {
	cmd.Env = os.Environ()
	for key, value := range envVariables {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}
}

func (s *TodocheckScenario) setupMockIssueTrackerServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if s.expectedAuthToken != "" && r.Header.Get("Authorization") != "Bearer "+s.expectedAuthToken {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		for issue := range s.issues {
			if r.URL.Path == issuetracker.IssueURLFrom(s.issueTracker, issue) {
				w.Write(issuetracker.BuildResponseFor(s.issueTracker, issue, s.issues[issue]))
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
	}))
}

func setupMockIssueTrackerCfg(cfgPath string, mockOrigin string) (teardownFunc, error) {
	patt := regexp.MustCompile("origin: \"?[a-zA-Z0-9._:/]+\"?")
	origBs, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't read config file %s: %w", cfgPath, err)
	}

	mockBs := patt.ReplaceAll(origBs, []byte(fmt.Sprintf("origin: %s", mockOrigin)))

	err = ioutil.WriteFile(cfgPath, mockBs, 0755)
	if err != nil {
		return nil, fmt.Errorf("couldn't writeback mock issue tracker origin in file %s: %w", cfgPath, err)
	}

	return func() {
		err := ioutil.WriteFile(cfgPath, origBs, 0755)
		if err != nil {
			panic("couldn't teardown mock issue tracker: " + err.Error())
		}
	}, nil
}

func deleteTokensCache(tokensCache string) {
	if tokensCache == config.DefaultTokensCache() {
		return
	}

	err := os.Remove(tokensCache)
	if err != nil {
		panic("couldn't teardown test environment: failed to delete tokens cache " + tokensCache)
	}
}

func validationPipeline(fs ...validateFunc) error {
	for _, f := range fs {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}
