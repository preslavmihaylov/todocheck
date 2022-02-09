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
		WithBasepath("./scenarios/integrations/redmine_public").
		WithConfig("./test_configs/integrations/redmine_public.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine_public/main.go", 5).
				ExpectLine("// TODO 3: A closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine_public/main.go", 6).
				ExpectLine("// TODO 4: An issue with feedback")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine_public/main.go", 7).
				ExpectLine("// TODO 5: A resolved issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine_public/main.go", 8).
				ExpectLine("// TODO 6: A rejected issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/integrations/redmine_public/main.go", 9).
				ExpectLine("// TODO 14: a non-existent issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine_public/main.go", 13).
				ExpectLine("// TODO #3: A closed issue")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPrivateRedmineIntegration(t *testing.T) {
	err := scenariobuilder.NewScenario().
		OnlyRunOnCI().
		WithAuthTokenFromEnv("TESTS_REDMINE_PRIVATE_APITOKEN").
		WithBinary("../todocheck").
		WithBasepath("./scenarios/integrations/redmine_private").
		WithConfig("./test_configs/integrations/redmine_private.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine_private/main.go", 5).
				ExpectLine("// TODO 9: A closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine_private/main.go", 6).
				ExpectLine("// TODO 10: An issue with feedback")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine_private/main.go", 7).
				ExpectLine("// TODO 11: A resolved issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine_private/main.go", 8).
				ExpectLine("// TODO 12: A rejected issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/integrations/redmine_private/main.go", 9).
				ExpectLine("// TODO 14: a non-existent issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/redmine_private/main.go", 13).
				ExpectLine("// TODO #9: A closed issue")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPublicAzureIntegration(t *testing.T) {
	err := scenariobuilder.NewScenario().
		OnlyRunOnCI().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/integrations/azureboards_public").
		WithConfig("./test_configs/integrations/azureboards_public.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/integrations/azureboards_public/main.go", 3).
				ExpectLine("// TODO MALFORMED Issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/azureboards_public/main.go", 6).
				ExpectLine("// TODO 3: An issue in CLOSED column")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/integrations/azureboards_public/main.go", 7).
				ExpectLine("// TODO 12345: A non-existent issue")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPrivateAzureIntegration(t *testing.T) {
	err := scenariobuilder.NewScenario().
		OnlyRunOnCI().
		WithAuthTokenFromEnv("TESTS_AZUREBOARDS_PRIVATE_APITOKEN").
		WithBinary("../todocheck").
		WithBasepath("./scenarios/integrations/azureboards_private").
		WithConfig("./test_configs/integrations/azureboards_private.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/integrations/azureboards_private/main.go", 3).
				ExpectLine("// TODO: 1 A malformed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/azureboards_private/main.go", 5).
				ExpectLine("// TODO 9: An issue in CLOSED column")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/integrations/azureboards_private/main.go", 7).
				ExpectLine("// TODO 999: A non-existent issue")).
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

func TestPublicYoutrackIntegration(t *testing.T) {
	err := scenariobuilder.NewScenario().
		OnlyRunOnCI().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/integrations/youtrack_public_incloud").
		WithAuthTokenFromEnv("TESTS_YOUTRACK_PUBLIC_INCLOUD_APITOKEN").
		WithConfig("./test_configs/integrations/youtrack_public_incloud.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/integrations/youtrack_public_incloud/main.go", 4).
				ExpectLine("// TODO DEMO-1: A closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/integrations/youtrack_public_incloud/main.go", 5).
				ExpectLine("// TODO DEMO-1234: A non-existent issue")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}
