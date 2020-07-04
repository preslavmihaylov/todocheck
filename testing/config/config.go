// Package config allows you to construct
// test scenarios which are based on executing the real todocheck binary
// against projects to use as scenarios.
package config

import (
	"fmt"
	"os/exec"
)

// TodocheckScenario encapsulates the scenario the program is expected to execute.
// This let's you specify what are the program inputs & what is the expected outputs.
type TodocheckScenario struct {
	binaryLoc string
	basepath  string
	cfgPath   string
}

// NewScenario to execute against the todocheck program
func NewScenario() *TodocheckScenario {
	return &TodocheckScenario{
		binaryLoc: "./todocheck",
		basepath:  ".",
		cfgPath:   ".todocheck.yaml",
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

// Run sets up the environment & executes the configured scenario
func (s *TodocheckScenario) Run() error {
	cmd := exec.Command(s.binaryLoc, "--basepath", s.basepath, "--config", s.cfgPath)
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Couldn't execute todocheck binary. Received error: %w", err)
	}

	fmt.Println(string(out))
	return nil
}
