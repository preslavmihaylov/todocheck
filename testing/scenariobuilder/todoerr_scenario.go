package scenariobuilder

import (
	"fmt"

	"github.com/preslavmihaylov/todocheck/checker/errors"
)

// TodoErrScenario encapsulates a test scenario for an expected todo err the program should return.
type TodoErrScenario struct {
	errType       errors.TODOErrType
	sourceFile    string
	sourceLineNum int
	contents      []string
}

// NewTodoErr returns a new todo err scenario
func NewTodoErr() *TodoErrScenario {
	return &TodoErrScenario{}
}

// WithType specifies the expected todo err type for the given scenario
func (s *TodoErrScenario) WithType(t errors.TODOErrType) *TodoErrScenario {
	s.errType = t
	return s
}

// WithSourceFile specifies the expected source file for the given todo err scenario
func (s *TodoErrScenario) WithSourceFile(sf string) *TodoErrScenario {
	s.sourceFile = sf
	return s
}

// WithLineNum specifies the expected starting line number for the given todo err scenario
func (s *TodoErrScenario) WithLineNum(n int) *TodoErrScenario {
	s.sourceLineNum = n
	return s
}

// WithLocation specifies the expected source file & line number for the given todo err scenario
func (s *TodoErrScenario) WithLocation(source string, linenum int) *TodoErrScenario {
	return s.WithSourceFile(source).WithLineNum(linenum)
}

// ExpectLine appends an expected error line to the todo err scenario. more than one line can be specified.
func (s *TodoErrScenario) ExpectLine(line string) *TodoErrScenario {
	s.contents = append(s.contents, line)
	return s
}

func (s *TodoErrScenario) String() string {
	str := fmt.Sprintf("ERROR: %s", s.errType)
	for i := 0; i < len(s.contents); i++ {
		str += fmt.Sprintf("\n%s:%d: %s", s.sourceFile, i+s.sourceLineNum, s.contents[i])
	}

	if s.errType == errors.TODOErrTypeMalformed {
		str += "\n\t> TODO should match pattern - TODO {task_id}:"
	}

	return str
}
