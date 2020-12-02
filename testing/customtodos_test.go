package testing

import (
	"testing"

	"github.com/preslavmihaylov/todocheck/checker/errors"
	"github.com/preslavmihaylov/todocheck/testing/scenariobuilder"
	"github.com/preslavmihaylov/todocheck/testing/scenariobuilder/issuetracker"
)

func TestScriptsCustomTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/custom_todos/scripts/").
		WithConfig("./test_configs/no_issue_tracker_and_custom_todos.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("123", issuetracker.StatusOpen).
		WithIssue("321", issuetracker.StatusClosed).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/scripts/script.sh", 1).
				ExpectLine("# This is a malformed TODO")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/scripts/script.sh", 3).
				ExpectLine("curl \"localhost:8080\" # This is a TODO comment at the end of the line")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/scripts/script.sh", 5).
				ExpectLine("# This is a malformed ToDo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/scripts/script.sh", 7).
				ExpectLine("curl \"localhost:8080\" # This is a ToDo comment at the end of the line")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/scripts/script.sh", 9).
				ExpectLine("# This is a malformed @fixme")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/scripts/script.sh", 11).
				ExpectLine("curl \"localhost:8080\" # This is a @fixme comment at the end of the line")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/scripts/script.sh", 13).
				ExpectLine("# This is a malformed @fix")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/scripts/script.sh", 15).
				ExpectLine("curl \"localhost:8080\" # This is a @fix comment at the end of the line")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/scripts/script.bash", 3).
				ExpectLine("# A malformed TODO comment")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/scripts/script.bash", 5).
				ExpectLine("# TODO 321: This is an invalid todo, marked against a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/scripts/script.bash", 6).
				ExpectLine("curl \"localhost:8080\" # TODO 567: This is an invalid todo, marked against a non-existent issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/scripts/script.bash", 8).
				ExpectLine("# A malformed ToDo comment")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/scripts/script.bash", 10).
				ExpectLine("# ToDo 321: This is an invalid todo, marked against a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/scripts/script.bash", 11).
				ExpectLine("curl \"localhost:8080\" # ToDo 567: This is an invalid todo, marked against a non-existent issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/scripts/script.bash", 13).
				ExpectLine("# A malformed @fixme comment")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/scripts/script.bash", 15).
				ExpectLine("# @fixme 321: This is an invalid todo, marked against a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/scripts/script.bash", 16).
				ExpectLine("curl \"localhost:8080\" # @fixme 567: This is an invalid todo, marked against a non-existent issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/scripts/script.bash", 18).
				ExpectLine("# A malformed @fix comment")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/scripts/script.bash", 20).
				ExpectLine("# @fix 321: This is an invalid todo, marked against a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/scripts/script.bash", 21).
				ExpectLine("curl \"localhost:8080\" # @fix 567: This is an invalid todo, marked against a non-existent issue")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestStandardCustomTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/custom_todos/standard/").
		WithConfig("./test_configs/no_issue_tracker_and_custom_todos.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("123", issuetracker.StatusOpen).
		WithIssue("321", issuetracker.StatusClosed).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/standard/main.go", 3).
				ExpectLine("// A malformed TODO comment")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/standard/main.go", 5).
				ExpectLine("// TODO 321: This is an invalid todo, marked against a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/standard/main.go", 6).
				ExpectLine("/*").
				ExpectLine(" * TODO 567: This is an invalid multiline todo, marked against a non-existent issue").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/standard/main.go", 10).
				ExpectLine("// A malformed ToDo comment")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/standard/main.go", 12).
				ExpectLine("// ToDo 321: This is an invalid todo, marked against a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/standard/main.go", 13).
				ExpectLine("/*").
				ExpectLine(" * ToDo 567: This is an invalid multiline todo, marked against a non-existent issue").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/standard/main.go", 17).
				ExpectLine("// A malformed @fixme comment")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/standard/main.go", 19).
				ExpectLine("// @fixme 321: This is an invalid todo, marked against a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/standard/main.go", 20).
				ExpectLine("/*").
				ExpectLine(" * @fixme 567: This is an invalid multiline todo, marked against a non-existent issue").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/standard/main.go", 24).
				ExpectLine("// A malformed @fix comment")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/standard/main.go", 26).
				ExpectLine("// @fix 321: This is an invalid todo, marked against a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/standard/main.go", 27).
				ExpectLine("/*").
				ExpectLine(" * @fix 567: This is an invalid multiline todo, marked against a non-existent issue").
				ExpectLine(" */")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPythonCustomTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/custom_todos/python").
		WithConfig("./test_configs/no_issue_tracker_and_custom_todos.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("1", issuetracker.StatusOpen).
		WithIssue("234", issuetracker.StatusClosed).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/python/main.py", 3).
				ExpectLine("# This is a single-line malformed TODO")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/python/main.py", 5).
				ExpectLine("\"\"\"").
				ExpectLine("And this is a multiline malformed TODO").
				ExpectLine("It should be parsed properly").
				ExpectLine("\"\"\"")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/python/main.py", 10).
				ExpectLine("'''").
				ExpectLine("This is the same multiline malformed TODO").
				ExpectLine("but with single-quotes").
				ExpectLine("'''")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/python/main.py", 15).
				ExpectLine("myvar = 5 # This is a malformed TODO at the end of a line")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/python/main.py", 19).
				ExpectLine("hello = \"hello\" # TODO 234: This is an invalid todo, with a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/python/main.py", 21).
				ExpectLine("\"\"\"").
				ExpectLine("TODO 234: This is an invalid todo, marked against a closed issue").
				ExpectLine("\"\"\"")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/python/main.py", 25).
				ExpectLine("'''").
				ExpectLine("TODO 234: This is an invalid todo,").
				ExpectLine("marked against a closed issue with single quotes").
				ExpectLine("'''")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/python/main.py", 31).
				ExpectLine("# This is a single-line malformed @fix")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/python/main.py", 33).
				ExpectLine("\"\"\"").
				ExpectLine("And this is a multiline malformed @fix").
				ExpectLine("It should be parsed properly").
				ExpectLine("\"\"\"")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/python/main.py", 38).
				ExpectLine("'''").
				ExpectLine("This is the same multiline malformed @fix").
				ExpectLine("but with single-quotes").
				ExpectLine("'''")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/python/main.py", 43).
				ExpectLine("myvar = 5 # This is a malformed @fix at the end of a line")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/python/main.py", 47).
				ExpectLine("hello = \"hello\" # @fix 234: This is an invalid todo, with a closed issue")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/python/main.py", 49).
				ExpectLine("\"\"\"").
				ExpectLine("@fix 234: This is an invalid todo, marked against a closed issue").
				ExpectLine("\"\"\"")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/python/main.py", 53).
				ExpectLine("'''").
				ExpectLine("@fix 234: This is an invalid todo,").
				ExpectLine("marked against a closed issue with single quotes").
				ExpectLine("'''")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGroovyCustomTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/custom_todos/groovy").
		WithConfig("./test_configs/no_issue_tracker_and_custom_todos.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("1", issuetracker.StatusOpen).
		WithIssue("2", issuetracker.StatusClosed).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 1).
				ExpectLine("//TODO: regular inline comment")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 11).
				ExpectLine("/*").
				ExpectLine("* TODO: Multi-line invalid todo").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 15).
				ExpectLine("/**").
				ExpectLine("* TODO: groovydoc invalid todo").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 19).
				ExpectLine("// TODO 2: The issue is closed")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 21).
				ExpectLine("// TODO 3: The issue is non-existent")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 52).
				ExpectLine("/*").
				ExpectLine("* TODO 2: Invalid todo as issue is closed").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 56).
				ExpectLine("/**").
				ExpectLine("* TODO 2: Invalid todo as issue is closed").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 61).
				ExpectLine("//@fix: regular inline comment")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 71).
				ExpectLine("/*").
				ExpectLine("* @fix: Multi-line invalid todo").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 75).
				ExpectLine("/**").
				ExpectLine("* @fix: groovydoc invalid todo").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 79).
				ExpectLine("// @fix 2: The issue is closed")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 81).
				ExpectLine("// @fix 3: The issue is non-existent")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 112).
				ExpectLine("/*").
				ExpectLine("* @fix 2: Invalid todo as issue is closed").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/groovy/main.groovy", 116).
				ExpectLine("/**").
				ExpectLine("* @fix 2: Invalid todo as issue is closed").
				ExpectLine("*/")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

// Also tests Rust TODOs as rust uses the same comment syntax
func TestSwiftCustomTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/custom_todos/swift").
		WithConfig("./test_configs/no_issue_tracker_and_custom_todos.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("1", issuetracker.StatusOpen).
		WithIssue("2", issuetracker.StatusClosed).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/swift/main.swift", 4).
				ExpectLine("print(\"Hello, World!\") // TODO: An invalid todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/swift/main.swift", 6).
				ExpectLine("/*").
				ExpectLine(" * TODO: An invalid multi-line todo").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/swift/main.swift", 10).
				ExpectLine("/// TODO: An invalid docstring todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/swift/main.swift", 12).
				ExpectLine("/*").
				ExpectLine(" /* TODO: An invalid nested todo */").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/swift/main.swift", 22).
				ExpectLine("/*").
				ExpectLine(" /*").
				ExpectLine("    /* TODO 2: invalid todo as issue is closed */").
				ExpectLine(" */").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/swift/main.swift", 29).
				ExpectLine("print(\"Hello, World!\") // @fix: An invalid todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/swift/main.swift", 31).
				ExpectLine("/*").
				ExpectLine(" * @fix: An invalid multi-line todo").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/swift/main.swift", 35).
				ExpectLine("/// @fix: An invalid docstring todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/swift/main.swift", 37).
				ExpectLine("/*").
				ExpectLine(" /* @fix: An invalid nested todo */").
				ExpectLine("*/")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/swift/main.swift", 47).
				ExpectLine("/*").
				ExpectLine(" /*").
				ExpectLine("    /* @fix 2: invalid todo as issue is closed */").
				ExpectLine(" */").
				ExpectLine("*/")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPHPCustomTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/custom_todos/php").
		WithConfig("./test_configs/no_issue_tracker_and_custom_todos.yaml").
		WithIssueTracker(issuetracker.Jira).
		WithIssue("1", issuetracker.StatusOpen).
		WithIssue("2", issuetracker.StatusClosed).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/php/main.php", 2).
				ExpectLine("// TODO: malformed todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/php/main.php", 4).
				ExpectLine("// TODO 2: The issue is closed")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/php/main.php", 5).
				ExpectLine("// TODO 3: The issue is non-existent")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/php/main.php", 7).
				ExpectLine("# TODO: malformed todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/php/main.php", 9).
				ExpectLine("# TODO 2: The issue is closed")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/php/main.php", 10).
				ExpectLine("# TODO 3: The issue is non-existent")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/php/main.php", 19).
				ExpectLine("/*").
				ExpectLine(" * TODO: Multi-line invalid todo").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/php/main.php", 27).
				ExpectLine("/*").
				ExpectLine(" * TODO 2: issue is closed").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/php/main.php", 31).
				ExpectLine("/*").
				ExpectLine(" * TODO 3: issue is non-existent").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/php/main.php", 35).
				ExpectLine("/**").
				ExpectLine(" * TODO: docstring invalid todo").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/php/main.php", 43).
				ExpectLine("/**").
				ExpectLine(" * TODO 2: issue is closed").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/php/main.php", 47).
				ExpectLine("/**").
				ExpectLine(" * TODO 3: issue is non-existent").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/php/main.php", 52).
				ExpectLine("// @fix: malformed todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/php/main.php", 54).
				ExpectLine("// @fix 2: The issue is closed")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/php/main.php", 55).
				ExpectLine("// @fix 3: The issue is non-existent")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/php/main.php", 57).
				ExpectLine("# @fix: malformed todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/php/main.php", 59).
				ExpectLine("# @fix 2: The issue is closed")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/php/main.php", 60).
				ExpectLine("# @fix 3: The issue is non-existent")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/php/main.php", 69).
				ExpectLine("/*").
				ExpectLine(" * @fix: Multi-line invalid todo").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/php/main.php", 77).
				ExpectLine("/*").
				ExpectLine(" * @fix 2: issue is closed").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/php/main.php", 81).
				ExpectLine("/*").
				ExpectLine(" * @fix 3: issue is non-existent").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeMalformed).
				WithLocation("scenarios/custom_todos/php/main.php", 85).
				ExpectLine("/**").
				ExpectLine(" * @fix: docstring invalid todo").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeIssueClosed).
				WithLocation("scenarios/custom_todos/php/main.php", 93).
				ExpectLine("/**").
				ExpectLine(" * @fix 2: issue is closed").
				ExpectLine(" */")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(errors.TODOErrTypeNonExistentIssue).
				WithLocation("scenarios/custom_todos/php/main.php", 97).
				ExpectLine("/**").
				ExpectLine(" * @fix 3: issue is non-existent").
				ExpectLine(" */")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}
