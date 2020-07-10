package scenariobuilder

import (
	"fmt"
	"strings"

	"github.com/preslavmihaylov/todocheck/common"
)

func validateTodoErrs(programOutput string, scenarios []*TodoErrScenario) validateFunc {
	return func() error {
		chunks := common.RemoveEmptyTokens(strings.Split(programOutput, "\n\n"))
		if len(chunks) != len(scenarios) {
			return fmt.Errorf("Invalid amount of todo errors detected.\nExpected: %d, Actual: %d\n\n"+
				"(program output follows...)\n%s",
				len(scenarios), len(chunks), programOutput)
		}

		for i := range chunks {
			j := indexOfMatchingTodoScenario(scenarios, chunks[i])
			if j == -1 {
				return fmt.Errorf(
					"No matching todo detected in any of the scenarios"+
						"\n\nActual:\n%s\n\nRemaining scenarios:\n%s",
					chunks[i], printScenarios(scenarios))
			}

			scenarios = removeScenario(scenarios, j)
		}

		return nil
	}
}

func indexOfMatchingTodoScenario(scenarios []*TodoErrScenario, target string) int {
	for i := range scenarios {
		if scenarios[i].String() == target {
			return i
		}
	}

	return -1
}

func printScenarios(ss []*TodoErrScenario) string {
	res := ""
	for i, s := range ss {
		res += fmt.Sprintf("(scenario #%d)\n%s\n\n", i+1, s.String())
	}

	return res
}

func removeScenario(scenarios []*TodoErrScenario, i int) []*TodoErrScenario {
	return append(scenarios[:i], scenarios[i+1:]...)
}
