package scenariobuilder

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/preslavmihaylov/todocheck/config/authtokens"
	"gopkg.in/yaml.v2"
)

func validateAuthTokensCache(tokensCache string, url, expectedToken string) validateFunc {
	return func() error {
		if expectedToken == "" {
			return nil
		} else if tokensCache == authtokens.DefaultConfigFile() {
			return errors.New("tokens_cache is not set in the configuration. It must be set for auth token scenarios")
		}

		authTokensBs, err := ioutil.ReadFile(tokensCache)
		if err != nil {
			return fmt.Errorf("couldn't read auth tokens config file %s: %w", tokensCache, err)
		}

		var authTokensCfg *authtokens.Config
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
