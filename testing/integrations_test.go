package testing

import (
	"testing"

	"github.com/preslavmihaylov/todocheck/checker/errors"
	"github.com/preslavmihaylov/todocheck/testing/scenariobuilder"
)

func TestPublicGithubIntegration(t *testing.T) {
	err := baseGithubScenario().
		OnlyRunOnCI().
		WithAuthTokenFromEnv("TESTS_GITHUB_APITOKEN").
		WithConfig("./test_configs/integrations/github_public.yaml").
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPrivateGithubIntegration(t *testing.T) {
	err := baseGithubScenario().
		OnlyRunOnCI().
		WithAuthTokenFromEnv("TESTS_GITHUB_APITOKEN").
		WithConfig("./test_configs/integrations/github_private.yaml").
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPublicGitlabIntegration(t *testing.T) {
	err := baseGitlabScenario().
		OnlyRunOnCI().
		WithAuthTokenFromEnv("TESTS_GITLAB_APITOKEN").
		WithConfig("./test_configs/integrations/gitlab_public.yaml").
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPrivateGitlabIntegration(t *testing.T) {
	err := baseGitlabScenario().
		OnlyRunOnCI().
		WithAuthTokenFromEnv("TESTS_GITLAB_APITOKEN").
		WithConfig("./test_configs/integrations/gitlab_private.yaml").
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPivotalTrackerIntegration(t *testing.T) {
	err := scenariobuilder.NewScenario().
		OnlyRunOnCI().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/integrations/pivotaltracker").
		WithAuthTokenFromEnv("TESTS_PIVOTALTRACKER_APITOKEN").
		WithConfig("./test_configs/integrations/pivotaltracker.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/pivotaltracker/main.go", 5).
				ExpectLine("// TODO #175938853: A finished todo (closed)")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/pivotaltracker/main.go", 7).
				ExpectLine("// TODO #175938860: A delivered todo (closed)")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/pivotaltracker/main.go", 11).
				ExpectLine("// TODO #175938899: A rejected todo (closed)")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/pivotaltracker/main.go", 13).
				ExpectLine("// TODO #175938883: An accepted todo (closed)")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/integrations/pivotaltracker/main.go", 15).
				ExpectLine("// TODO #199938883: A non-existent issue")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPublicRedmineIntegration(t *testing.T) {
	err := scenariobuilder.NewScenario().
		OnlyRunOnCI().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/integrations/redmine").
		WithConfig("./test_configs/integrations/redmine_public.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine/main.go", 5).
				ExpectLine("// TODO 3: A closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine/main.go", 6).
				ExpectLine("// TODO 4: An issue with feedback")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine/main.go", 7).
				ExpectLine("// TODO 5: A resolved issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine/main.go", 8).
				ExpectLine("// TODO 6: A rejected issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine/main.go", 9).
				ExpectLine("// TODO 14: a non-existent issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine/main.go", 13).
				ExpectLine("// TODO #3: A closed issue")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func baseGithubScenario() *scenariobuilder.TodocheckScenario {
	return scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/integrations/github").
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
				ExpectLine("// TODO #3: A non-existent issue"))
}

func baseGitlabScenario() *scenariobuilder.TodocheckScenario {
	return baseGithubScenario()
}
