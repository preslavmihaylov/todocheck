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
	metadata      map[string]string
}

// NewTodoErr returns a new todo err scenario
func NewTodoErr() *TodoErrScenario {
	return &TodoErrScenario{
		metadata: map[string]string{},
	}
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

// WithMetadata extends existing metadata with a multiple of key value pairs
// expected within the `metadata` field of the TODOs
func (s *TodoErrScenario) WithMetadata(metadata *map[string]string) *TodoErrScenario {
	for k, v := range *metadata {
		s.metadata[k] = v
	}
	return s
}

// WithMetadataEntry stores a single key-value pair that is expected in the `metadata` field
func (s *TodoErrScenario) WithMetadataEntry(key string, value string) *TodoErrScenario {
	s.metadata[key] = value
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

type TodoErrForJSON struct {
	Type     string            `json:"type"`
	Filename string            `json:"filename"`
	Line     int               `json:"line"`
	Message  string            `json:"message"`
	Metadata map[string]string `json:"metadata"`
}

func (s *TodoErrForJSON) isEqual(other *TodoErrForJSON) bool {
	basic := s.Type == other.Type &&
		s.Filename == other.Filename &&
		s.Line == other.Line &&
		s.Message == other.Message
	if !basic {
		return false
	}
	for k, v := range s.Metadata {
		if other.Metadata[k] != v {
			return false
		}
	}

	return true
}

func (s *TodoErrScenario) ToTodoErrForJSON() *TodoErrForJSON {
	res := &TodoErrForJSON{
		Type:     string(s.errType),
		Filename: s.sourceFile,
		Line:     s.sourceLineNum,
		Message:  "",
		Metadata: s.metadata,
	}

	if s.errType == errors.TODOErrTypeMalformed {
		res.Message = "TODO should match pattern - TODO {task_id}:"
	}

	return res
}
