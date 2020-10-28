package testing

import (
	"testing"

	"github.com/preslavmihaylov/todocheck/checker/errors"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/testing/scenariobuilder"
	"github.com/preslavmihaylov/todocheck/testing/scenariobuilder/issuetracker"
	"github.com/preslavmihaylov/todocheck/validation"
)

func TestSingleLineMalformedTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/singleline_todos").
		WithConfig("./test_configs/no_issue_tracker.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/singleline_todos/other.go", 3).
				ExpectLine("// TODO: This is a todo in a separate go file")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/singleline_todos/main.go", 5).
				ExpectLine("// TODO: This is a malformed todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/singleline_todos/main.go", 6).
				ExpectLine("// TODO: This is a malformed todo 2")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/singleline_todos/main.go", 10).
				ExpectLine("func main() { // TODO: This is a todo comment at the end of a line")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/singleline_todos/main.go", 15).
				ExpectLine("// TODO comment without colons")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/singleline_todos/main.go", 17).
				ExpectLine("// This is a TODO comment at the middle of it")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestFirstlineMalformedTodo(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/firstline_comment").
		WithConfig("./test_configs/no_issue_tracker.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/firstline_comment/main.cpp", 1).
				ExpectLine("// This is an invalid TODO on the very first line of the file")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/firstline_comment/other.cpp", 1).
				ExpectLine("// This is another first-line TODO comment in a second file")).
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
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/multiline_todos/main.go", 3).
				ExpectLine("/*").
				ExpectLine(" * This is a multiline").
				ExpectLine(" * TODO comment.").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/multiline_todos/main.go", 8).
				ExpectLine("func main() { /*").
				ExpectLine("	 * This is a multiline TODO comment").
				ExpectLine("	 * spanning several lines").
				ExpectLine("	 */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
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
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/annotated_todos/main.go", 3).
				ExpectLine("// TODO J123: This is a todo, annotated with a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/annotated_todos/main.go", 7).
				ExpectLine("// TODO J456: This is an invalid todo, annotated with a non-existent issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/annotated_todos/main.go", 14).
				ExpectLine("/*").
				ExpectLine(" * This is an invalid TODO J123:").
				ExpectLine(" * as the issue is closed").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
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

func TestScriptsTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/scripts").
		WithConfig("./test_configs/no_issue_tracker.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("123", issuetracker.StatusOpen).
		WithIssue("321", issuetracker.StatusClosed).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/scripts/script.sh", 1).
				ExpectLine("# This is a malformed TODO")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/scripts/script.sh", 5).
				ExpectLine("curl \"localhost:8080\" # This is a TODO comment at the end of the line")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/scripts/script.bash", 3).
				ExpectLine("# A malformed TODO comment")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/scripts/script.bash", 7).
				ExpectLine("# TODO 321: This is an invalid todo, marked against a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/scripts/script.bash", 9).
				ExpectLine("curl \"localhost:8080\" # TODO 567: This is an invalid todo, marked against a non-existent issue")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPythonTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/python").
		WithConfig("./test_configs/no_issue_tracker.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("1", issuetracker.StatusOpen).
		WithIssue("234", issuetracker.StatusClosed).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/python/main.py", 3).
				ExpectLine("# This is a single-line malformed TODO")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/python/main.py", 5).
				ExpectLine("\"\"\"").
				ExpectLine("And this is a multiline malformed TODO").
				ExpectLine("It should be parsed properly").
				ExpectLine("\"\"\"")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/python/main.py", 10).
				ExpectLine("'''").
				ExpectLine("This is the same multiline malformed TODO").
				ExpectLine("but with single-quotes").
				ExpectLine("'''")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/python/main.py", 15).
				ExpectLine("myvar = 5 # This is a malformed TODO at the end of a line")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/python/main.py", 19).
				ExpectLine("hello = \"hello\" # TODO 234: This is an invalid todo, with a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/python/main.py", 21).
				ExpectLine("\"\"\"").
				ExpectLine("TODO 234: This is an invalid todo, marked against a closed issue").
				ExpectLine("\"\"\"")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/python/main.py", 25).
				ExpectLine("'''").
				ExpectLine("TODO 234: This is an invalid todo,").
				ExpectLine("marked against a closed issue with single quotes").
				ExpectLine("'''")).
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
		DeleteTokensCacheAfter().
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestAuthTokenViaEnvVariable(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/offline_token").
		WithConfig("./test_configs/offline_token.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("J123", issuetracker.StatusOpen).
		RequireAuthToken("123456").
		WithEnvVariable("TODOCHECK_AUTH_TOKEN", "123456").
		DeleteTokensCacheAfter().
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
				WithType(errors.TODOErrTypeMalformed).
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
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("../testing/scenarios/ignored_dirs/main.go", 3).
				ExpectLine("// This is a malformed TODO")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestInvalidIssueTracker(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithConfig("./test_configs/invalid_issue_tracker.yaml").
		ExpectExecutionError().
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestValidGithubAccess(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/repo_health_checks").
		WithTestEnvConfig("./test_configs/valid_github_access.yaml").
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestInvalidGithubAccess(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithTestEnvConfig("./test_configs/invalid_github_access.yaml").
		ExpectExecutionError().
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestInvalidOrigins(t *testing.T) {
	invalidConfigPaths := []string{
		"./test_configs/invalid_github_https.yaml",
		"./test_configs/invalid_github_origin.yaml",
		"./test_configs/invalid_github_www.yaml",
		"./test_configs/invalid_gitlab_origin.yaml",
		"./test_configs/invalid_gitlab_port.yaml",
		"./test_configs/invalid_issue_tracker.yaml",
		"./test_configs/invalid_jira_origin.yaml",
		"./test_configs/invalid_jira_port.yaml",
		"./test_configs/invalid_offline_url.yaml",
		"./test_configs/invalid_pivotal_origin.yaml",
		"./test_configs/invalid_redmine_origin.yaml",
		"./test_configs/invalid_redmine_port.yaml",
	}

	for _, path := range invalidConfigPaths {
		cfg, err := config.NewLocal(path, ".")
		if err != nil {
			t.Errorf("%s", err)
			continue
		}
		errors := validation.Validate(cfg)
		if 0 == len(errors) {
			t.Errorf("%s should be invalid", path)
		}
	}
}

func TestValidOrigins(t *testing.T) {
	validConfigPaths := []string{
		"./test_configs/valid_github_https.yaml",
		"./test_configs/valid_github_origin.yaml",
		"./test_configs/valid_github_www.yaml",
		"./test_configs/valid_gitlab_origin.yaml",
		"./test_configs/valid_gitlab_port.yaml",
		"./test_configs/valid_gitlab_subdomain.yaml",
		"./test_configs/valid_jira_origin.yaml",
		"./test_configs/valid_jira_port.yaml",
		"./test_configs/valid_jira_subdomain.yaml",
		"./test_configs/valid_pivotal_origin.yaml",
		"./test_configs/valid_redmine_origin.yaml",
		"./test_configs/valid_redmine_port.yaml",
		"./test_configs/valid_redmine_subdomain.yaml",
	}

	for _, path := range validConfigPaths {
		cfg, err := config.NewLocal(path, ".")
		if err != nil {
			t.Errorf("%s", err)
			continue
		}
		errors := validation.Validate(cfg)
		if len(errors) > 0 {
			t.Errorf("%s should be valid but has errors: %v", path, errors)
		}
	}
}

func TestInvalidOfflineURL(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithConfig("./test_configs/invalid_offline_url.yaml").
		ExpectExecutionError().
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestNonExistentOfflineURL(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithConfig("./test_configs/non_existent_offline_url.yaml").
		ExpectExecutionError().
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestTraversingNonExistentDirectory(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("../testing/scenarios/non_existent_dir").
		WithConfig("./test_configs/no_issue_tracker.yaml").
		ExpectExecutionError().
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestDefaultAuthInConfig(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/no_auth_section").
		WithConfig("./test_configs/no_auth_section.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/no_auth_section/main.go", 3).
				ExpectLine("// TODO - malformed todo")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

// Test that the configuration path can be derived implicitly from the basepath
func TestConfigDerivedFromBasepath(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/basepath_config").
		WithTestEnvConfig("./scenarios/basepath_config/.todocheck.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/basepath_config/main.go", 3).
				ExpectLine("// TODO - malformed todo")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestConfigAutoDetectWithSSHGitConfig(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/auto_detect_config").
		WithTestEnvConfig("./scenarios/auto_detect_config/expected_config.yaml").
		WithGitConfig("git@github.com:preslavmihaylov/todocheck.git").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/auto_detect_config/main.go", 3).
				ExpectLine("// TODO - malformed todo")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestConfigAutoDetectWithHTTPSGitConfig(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/auto_detect_config").
		WithTestEnvConfig("./scenarios/auto_detect_config/expected_config.yaml").
		WithGitConfig("https://github.com/preslavmihaylov/todocheck").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/auto_detect_config/main.go", 3).
				ExpectLine("// TODO - malformed todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/auto_detect_config/main.go", 5).
				ExpectLine("// TODO #1: closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/auto_detect_config/main.go", 7).
				ExpectLine("// TODO 2: closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/auto_detect_config/main.go", 9).
				ExpectLine("// TODO #9999999: non-existent issue")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGroovyTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/groovy").
		WithConfig("./test_configs/no_issue_tracker.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("1", issuetracker.StatusOpen).
		WithIssue("2", issuetracker.StatusClosed).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/groovy/main.groovy", 1).
				ExpectLine("//TODO: regular inline comment")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/groovy/main.groovy", 11).
				ExpectLine("/*").
				ExpectLine("* TODO: Multi-line invalid todo").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/groovy/main.groovy", 15).
				ExpectLine("/**").
				ExpectLine("* TODO: groovydoc invalid todo").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/groovy/main.groovy", 19).
				ExpectLine("// TODO 2: The issue is closed")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/groovy/main.groovy", 21).
				ExpectLine("// TODO 3: The issue is non-existent")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/groovy/main.groovy", 52).
				ExpectLine("/*").
				ExpectLine("* TODO 2: Invalid todo as issue is closed").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/groovy/main.groovy", 56).
				ExpectLine("/**").
				ExpectLine("* TODO 2: Invalid todo as issue is closed").
				ExpectLine("*/")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

// Also tests Rust TODOs as rust uses the same comment syntax
func TestSwiftTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/swift").
		WithConfig("./test_configs/no_issue_tracker.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("1", issuetracker.StatusOpen).
		WithIssue("2", issuetracker.StatusClosed).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/swift/main.swift", 4).
				ExpectLine("print(\"Hello, World!\") // TODO: An invalid todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/swift/main.swift", 6).
				ExpectLine("/*").
				ExpectLine(" * TODO: An invalid multi-line todo").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/swift/main.swift", 10).
				ExpectLine("/// TODO: An invalid docstring todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/swift/main.swift", 12).
				ExpectLine("/*").
				ExpectLine(" /* TODO: An invalid nested todo */").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/swift/main.swift", 22).
				ExpectLine("/*").
				ExpectLine(" /*").
				ExpectLine("    /* TODO 2: invalid todo as issue is closed */").
				ExpectLine(" */").
				ExpectLine("*/")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPHPTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/php").
		WithConfig("./test_configs/no_issue_tracker.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("1", issuetracker.StatusOpen).
		WithIssue("2", issuetracker.StatusClosed).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/php/main.php", 2).
				ExpectLine("// TODO: malformed todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/php/main.php", 4).
				ExpectLine("// TODO 2: The issue is closed")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/php/main.php", 5).
				ExpectLine("// TODO 3: The issue is non-existent")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/php/main.php", 7).
				ExpectLine("# TODO: malformed todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/php/main.php", 9).
				ExpectLine("# TODO 2: The issue is closed")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/php/main.php", 10).
				ExpectLine("# TODO 3: The issue is non-existent")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/php/main.php", 19).
				ExpectLine("/*").
				ExpectLine(" * TODO: Multi-line invalid todo").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/php/main.php", 27).
				ExpectLine("/*").
				ExpectLine(" * TODO 2: issue is closed").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/php/main.php", 31).
				ExpectLine("/*").
				ExpectLine(" * TODO 3: issue is non-existent").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/php/main.php", 35).
				ExpectLine("/**").
				ExpectLine(" * TODO: docstring invalid todo").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/php/main.php", 43).
				ExpectLine("/**").
				ExpectLine(" * TODO 2: issue is closed").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/php/main.php", 47).
				ExpectLine("/**").
				ExpectLine(" * TODO 3: issue is non-existent").
				ExpectLine(" */")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

// This test prints the --version flag of the binary and expects that no other output,
// other than the version flag is printed.
//
// The target basepath includes files with todo errors, but no todo errors are expected to be printed as
// the version flagg overrides the normal execution of the program and exits the program after printing it.
func TestPrintingVersionFlagStopsProgram(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/php").
		WithConfig("./test_configs/no_issue_tracker.yaml").
		WithVersionFlag().
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}
