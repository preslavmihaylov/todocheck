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
	"os/exec"
	"regexp"

	"github.com/preslavmihaylov/todocheck/testing/scenariobuilder/issuetracker"
)

type teardownFunc func()

// TodocheckScenario encapsulates the scenario the program is expected to execute.
// This let's you specify what are the program inputs & what is the expected outputs.
type TodocheckScenario struct {
	binaryLoc        string
	basepath         string
	cfgPath          string
	authToken        string
	expectedExitCode int
	issueTracker     issuetracker.Type
	issues           map[string]issuetracker.Status
	todoErrScenarios []*TodoErrScenario
}

// NewScenario to execute against the todocheck program
func NewScenario() *TodocheckScenario {
	return &TodocheckScenario{
		binaryLoc:        "./todocheck",
		basepath:         ".",
		cfgPath:          ".todocheck.yaml",
		issues:           map[string]issuetracker.Status{},
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

// WithConfig let's you specify the --config flag passed to the program
func (s *TodocheckScenario) WithConfig(cfgPath string) *TodocheckScenario {
	s.cfgPath = cfgPath
	return s
}

// WithIssueTracker let's you specify what issue tracker to execute the scenario with
func (s *TodocheckScenario) WithIssueTracker(issueTracker issuetracker.Type) *TodocheckScenario {
	s.issueTracker = issueTracker
	return s
}

// RequireAuthToken on each issue lookup
func (s *TodocheckScenario) RequireAuthToken(token string) *TodocheckScenario {
	s.authToken = token
	return s
}

// WithIssue sets up a mock issue in your issue tracker with the given status
func (s *TodocheckScenario) WithIssue(issueID string, status issuetracker.Status) *TodocheckScenario {
	s.issues[issueID] = status
	return s
}

// ExpectTodoErr appends a new todo err scenario to expect from the program execution
func (s *TodocheckScenario) ExpectTodoErr(sc *TodoErrScenario) *TodocheckScenario {
	s.expectedExitCode = 1
	s.todoErrScenarios = append(s.todoErrScenarios, sc)
	return s
}

// Run sets up the environment & executes the configured scenario
func (s *TodocheckScenario) Run() error {
	teardown, err := s.setupTestEnvironment()
	if err != nil {
		return fmt.Errorf("couldn't setup test environment: %s", err)
	}
	defer teardown()

	cmd := exec.Command(s.binaryLoc, "--basepath", s.basepath, "--config", s.cfgPath)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); !ok || exitError.ExitCode() != s.expectedExitCode {
			return fmt.Errorf("program exited with an unexpected exit code: %s", err)
		}
	}

	fmt.Println("(standard output follows. Standard output is ignored & not validated...)")
	fmt.Println(stdout.String())

	return validateTodoErrs(stderr.String(), s.todoErrScenarios)
}

func (s *TodocheckScenario) setupTestEnvironment() (teardownFunc, error) {
	mockSrv := s.setupMockIssueTrackerServer()
	teardownIssueTrackerCfg, err := setupMockIssueTrackerCfg(s.cfgPath, mockSrv.URL)
	if err != nil {
		return nil, fmt.Errorf("couldn't setup mock issue tracker: %w", err)
	}

	return func() {
		teardownIssueTrackerCfg()
		mockSrv.Close()
	}, nil
}

func (s *TodocheckScenario) setupMockIssueTrackerServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if s.authToken != "" && r.Header.Get("Authorization") != "Bearer "+s.authToken {
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
