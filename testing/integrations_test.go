package testing

import (
	"testing"

	"github.com/preslavmihaylov/todocheck/checker/errors"
	"github.com/preslavmihaylov/todocheck/testing/scenariobuilder"
)

func TestPublicGithubIntegration(t *testing.T) {
	err := baseGithubScenario().
		OnlyRunOnCI().
		WithConfig("./test_configs/integrations/github_public.yaml").
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPrivateGithubIntegration(t *testing.T) {
	err := baseGithubScenario().
		OnlyRunOnCI().
		WithConfig("./test_configs/integrations/github_private.yaml").
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func baseGithubScenario() *scenariobuilder.TodocheckScenario {
	return scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/integrations/github").
		WithAuthTokenFromEnv("TESTS_GITHUB_APITOKEN").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/github/main.go", 3).
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
				ExpectLine("// TODO #3: A non-existent issue"))
}
