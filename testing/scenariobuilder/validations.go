package scenariobuilder

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/preslavmihaylov/todocheck/authmanager/authstore"
	"github.com/preslavmihaylov/todocheck/common"
	"github.com/preslavmihaylov/todocheck/config"
	"gopkg.in/yaml.v2"
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

func validateAuthTokensCache(tokensCache string, url, expectedToken string) validateFunc {
	return func() error {
		if expectedToken == "" {
			return nil
		} else if tokensCache == config.DefaultTokensCache() {
			return errors.New("tokens_cache is not set in the configuration. It must be set for auth token scenarios")
		}

		authTokensBs, err := ioutil.ReadFile(tokensCache)
		if err != nil {
			return fmt.Errorf("couldn't read auth tokens config file %s: %w", tokensCache, err)
		}

		var authTokensCfg *authstore.Config
		err = yaml.Unmarshal(authTokensBs, &authTokensCfg)
		if err != nil {
			return fmt.Errorf("failed to unmarshal auth tokens cfg %s: %w", tokensCache, err)
		}

		if authTokensCfg.Tokens[url] != expectedToken {
			return fmt.Errorf("Expected auth token not found in tokens cache %s\n\nExpected: %s Actual: %s",
				tokensCache, expectedToken, authTokensCfg.Tokens[url])
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
