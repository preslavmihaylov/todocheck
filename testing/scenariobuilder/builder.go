// Package scenariobuilder allows you to construct
// test scenarios which are based on executing the real todocheck binary
// against projects to use as scenarios.
package scenariobuilder

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// TodocheckScenario encapsulates the scenario the program is expected to execute.
// This let's you specify what are the program inputs & what is the expected outputs.
type TodocheckScenario struct {
	binaryLoc        string
	basepath         string
	cfgPath          string
	expectedExitCode int
	todoErrScenarios []*TodoErrScenario
}

// NewScenario to execute against the todocheck program
func NewScenario() *TodocheckScenario {
	return &TodocheckScenario{
		binaryLoc:        "./todocheck",
		basepath:         ".",
		cfgPath:          ".todocheck.yaml",
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

// ExpectTodoErr appends a new todo err scenario to expect from the program execution
func (s *TodocheckScenario) ExpectTodoErr(sc *TodoErrScenario) *TodocheckScenario {
	s.expectedExitCode = 1
	s.todoErrScenarios = append(s.todoErrScenarios, sc)
	return s
}

// Run sets up the environment & executes the configured scenario
func (s *TodocheckScenario) Run() error {
	cmd := exec.Command(s.binaryLoc, "--basepath", s.basepath, "--config", s.cfgPath)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); !ok || exitError.ExitCode() != s.expectedExitCode {
			return fmt.Errorf("program exited with an unexpected exit code: %s", err)
		}
	}

	fmt.Println("(standard output follows. Standard output is ignored & not validated...)")
	fmt.Println(stdout.String())

	return validateTodoErrs(stderr.String(), s.todoErrScenarios)
}

func validateTodoErrs(programOutput string, scenarios []*TodoErrScenario) error {
	chunks := removeEmptyTokens(strings.Split(programOutput, "\n\n"))
	if len(chunks) != len(scenarios) {
		return fmt.Errorf("Invalid amount of todo errors detected.\nExpected: %d, Actual: %d\n\n"+
			"(program output follows...)\n%s",
			len(scenarios), len(chunks), programOutput)
	}

	for i := range chunks {
		if chunks[i] != scenarios[i].String() {
			out := fmt.Sprintf("Invalid todo error detected\n\nExpected:\n%s\n\nActual:\n%s",
				scenarios[i].String(), chunks[i])

			return errors.New(out)
		}
	}

	return nil
}

func removeEmptyTokens(ss []string) []string {
	res := []string{}
	for _, s := range ss {
		if s != "" {
			res = append(res, s)
		}
	}

	return res
}
