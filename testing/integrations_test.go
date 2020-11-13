package testing

import (
	"testing"

	"github.com/preslavmihaylov/todocheck/checker/errors"
	"github.com/preslavmihaylov/todocheck/testing/scenariobuilder"
)

func TestPublicGithubIntegration(t *testing.T) {
	err := scenariobuilder.NewScenario().
		OnlyRunOnCI().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/integrations/github").
		WithConfig("./test_configs/integrations/github.yaml").
		WithAuthTokenFromEnv("TESTS_GITHUB_APITOKEN").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/github/main.go", 4).
				ExpectLine("// TODO 2: A closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/integrations/github/main.go", 5).
				ExpectLine("// TODO 3: A non-existent issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/github/main.go", 8).
				ExpectLine("// TODO #2: A closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/integrations/github/main.go", 9).
				ExpectLine("// TODO #3: A non-existent issue")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}
