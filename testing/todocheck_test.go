package testing

import (
	"testing"

	"github.com/preslavmihaylov/todocheck/testing/scenariobuilder"
	"github.com/preslavmihaylov/todocheck/testing/scenariobuilder/issuetracker"
)

func TestSingleLineMalformedTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/singleline_todos").
		WithConfig("./test_configs/no_issue_tracker.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithLocation("scenarios/singleline_todos/other.go", 3).
				ExpectLine("// TODO: This is a todo in a separate go file")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithLocation("scenarios/singleline_todos/main.go", 5).
				ExpectLine("// TODO: This is a malformed todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithLocation("scenarios/singleline_todos/main.go", 6).
				ExpectLine("// TODO: This is a malformed todo 2")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithLocation("scenarios/singleline_todos/main.go", 10).
				ExpectLine("func main() { // TODO: This is a todo comment at the end of a line")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithLocation("scenarios/singleline_todos/main.go", 15).
				ExpectLine("// TODO comment without colons")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithLocation("scenarios/singleline_todos/main.go", 17).
				ExpectLine("// This is a TODO comment at the middle of it")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestMultiLineMalformedTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/multiline_todos").
		WithConfig("./test_configs/no_issue_tracker.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithLocation("scenarios/multiline_todos/main.go", 3).
				ExpectLine("/*").
				ExpectLine(" * This is a multiline").
				ExpectLine(" * TODO comment.").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithLocation("scenarios/multiline_todos/main.go", 8).
				ExpectLine("func main() { /*").
				ExpectLine("	 * This is a multiline TODO comment").
				ExpectLine("	 * spanning several lines").
				ExpectLine("	 */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithLocation("scenarios/multiline_todos/main.go", 18).
				ExpectLine("/* This is a single-line multi-line TODO comment */")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestAnnotatedTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/annotated_todos").
		WithConfig("./test_configs/no_issue_tracker.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("J123", issuetracker.StatusClosed).
		WithIssue("J321", issuetracker.StatusOpen).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeIssueClosed).
				WithLocation("scenarios/annotated_todos/main.go", 3).
				ExpectLine("// TODO J123: This is a todo, annotated with a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeIssueNonExistent).
				WithLocation("scenarios/annotated_todos/main.go", 7).
				ExpectLine("// TODO J456: This is an invalid todo, annotated with a non-existent issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeIssueClosed).
				WithLocation("scenarios/annotated_todos/main.go", 14).
				ExpectLine("/*").
				ExpectLine(" * This is an invalid TODO J123:").
				ExpectLine(" * as the issue is closed").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeIssueNonExistent).
				WithLocation("scenarios/annotated_todos/main.go", 19).
				ExpectLine("/*").
				ExpectLine(" * TODO J456:").
				ExpectLine(" * This issue doesn't exist").
				ExpectLine(" */")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestAuthTokensCache(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/auth_tokens_cache").
		WithConfig("./test_configs/auth_tokens.yaml").
		WithIssueTracker(issuetracker.Jira).
		RequireAuthToken("123456").
		WithIssue("J123", issuetracker.StatusOpen).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestOfflineToken(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/offline_token").
		WithConfig("./test_configs/offline_token.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("J123", issuetracker.StatusOpen).
		RequireAuthToken("123456").
		SetOfflineTokenWhenRequested("123456").
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestIgnoredDirectories(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/ignored_dirs").
		WithConfig("./test_configs/ignored_dirs.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithLocation("scenarios/ignored_dirs/main.go", 3).
				ExpectLine("// This is a malformed TODO")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestIgnoredDirectoriesWithDotDot(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("../testing/scenarios/ignored_dirs").
		WithConfig("./test_configs/ignored_dirs.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithLocation("../testing/scenarios/ignored_dirs/main.go", 3).
				ExpectLine("// This is a malformed TODO")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}
