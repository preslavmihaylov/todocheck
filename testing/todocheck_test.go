package testing

import (
	"testing"

	"github.com/preslavmihaylov/todocheck/testing/scenariobuilder"
)

func TestSingleLineMalformedTodos(t *testing.T) {
	err := scenariobuilder.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/simple_todos").
		WithConfig("./scenarios/simple_todos/.todocheck.yaml").
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithSourceFile("scenarios/simple_todos/main.go").WithLineNum(3).
				ExpectLine("// TODO: This is a malformed todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithSourceFile("scenarios/simple_todos/main.go").WithLineNum(4).
				ExpectLine("// TODO: This is a malformed todo 2")).Run()
	if err != nil {
		t.Errorf("todocheck scenario failed: %s", err)
	}
}
