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
				WithSourceFile("scenarios/simple_todos/other.go").WithLineNum(3).
				ExpectLine("// TODO: This is a todo in a separate go file")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithSourceFile("scenarios/simple_todos/main.go").WithLineNum(5).
				ExpectLine("// TODO: This is a malformed todo")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithSourceFile("scenarios/simple_todos/main.go").WithLineNum(6).
				ExpectLine("// TODO: This is a malformed todo 2")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithSourceFile("scenarios/simple_todos/main.go").WithLineNum(10).
				ExpectLine("func main() { // TODO: This is a todo comment at the end of a line")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithSourceFile("scenarios/simple_todos/main.go").WithLineNum(15).
				ExpectLine("// TODO comment without colons")).
		ExpectTodoErr(
			scenariobuilder.NewTodoErr().
				WithType(scenariobuilder.TodoErrTypeMalformed).
				WithSourceFile("scenarios/simple_todos/main.go").WithLineNum(17).
				ExpectLine("// This is a TODO comment at the middle of it")).
		Run()
	if err != nil {
		t.Errorf("%s", err)
	}
}
