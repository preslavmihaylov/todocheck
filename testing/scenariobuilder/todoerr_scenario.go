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
	message       string
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

// WithJSONMetadata extends existing metadata with a multiple of key value pairs
// expected within the `metadata` field of the TODOs
func (s *TodoErrScenario) WithJSONMetadata(metadata map[string]string) *TodoErrScenario {
	for k, v := range metadata {
		s.metadata[k] = v
	}
	return s
}

// WithJSONMetadataEntry stores a single key-value pair that is expected in the `metadata` field
func (s *TodoErrScenario) WithJSONMetadataEntry(key string, value string) *TodoErrScenario {
	s.metadata[key] = value
	return s
}

// WithMessage stores the expected extracted message.
func (s *TodoErrScenario) WithMessage(message string) *TodoErrScenario {
	s.message = message
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

func metadatasMatch(this, other *TodoErrForJSON) bool {
	if len(this.Metadata) != len(other.Metadata) {
		return false
	}
	for k, v := range this.Metadata {
		if other.Metadata[k] != v {
			return false
		}
	}

	return true
}

func (s *TodoErrForJSON) equals(other *TodoErrForJSON) bool {
	return s.Type == other.Type &&
		s.Filename == other.Filename &&
		s.Line == other.Line &&
		s.Message == other.Message &&
		metadatasMatch(s, other)
}

func (s *TodoErrScenario) ToTodoErrForJSON() *TodoErrForJSON {
	res := &TodoErrForJSON{
		Type:     string(s.errType),
		Filename: s.sourceFile,
		Line:     s.sourceLineNum,
		Message:  s.message,
		Metadata: s.metadata,
	}

	if s.errType == errors.TODOErrTypeMalformed {
		res.Message = "TODO should match pattern - TODO {task_id}:"
	}

	return res
}
