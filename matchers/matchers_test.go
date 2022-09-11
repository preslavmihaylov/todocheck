package matchers

import (
	"reflect"
	"testing"

	"github.com/preslavmihaylov/todocheck/matchers/state"
)

func TestTodoMatcherForFile(t *testing.T) {
	testTodo := []string{"// TODO"}

	for extension, factory := range supportedMatchers {
		t.Run(extension, func(t *testing.T) {
			matcher := TodoMatcherForFile("test"+extension, testTodo)
			want := factory.newTodoMatcher(testTodo)
			if matcher != want {
				t.Errorf("got %v want %v", matcher, want)
			}
		})
	}

	t.Run("Unsupported extension", func(t *testing.T) {
		matcher := TodoMatcherForFile("test.md", testTodo)
		if matcher != nil {
			t.Errorf("Expected nil matcher")
		}
	})
}

func TestCommentMatcherForFile(t *testing.T) {
	var testCommentCallback state.CommentCallback

	for extension, factory := range supportedMatchers {
		t.Run(extension, func(t *testing.T) {
			matcher := CommentMatcherForFile("test"+extension, testCommentCallback)
			want := factory.newCommentsMatcher(testCommentCallback)
			if !reflect.DeepEqual(matcher, want) {
				t.Errorf("got %v want %v", matcher, want)
			}
		})
	}

	t.Run("Unsupported extension", func(t *testing.T) {
		matcher := CommentMatcherForFile("test.md", testCommentCallback)
		if matcher != nil {
			t.Errorf("Expected nil matcher")
		}
	})
}

func TestSupportedFileExtensions(t *testing.T) {
	extensions := SupportedFileExtensions()

	for _, ext := range extensions {
		if _, ok := supportedMatchers[ext]; !ok {
			t.Errorf("Extension %v is not in supported extensions", ext)
		}
	}

	if len(extensions) != len(supportedMatchers) {
		t.Errorf("Some extensions from supported extensions are missing")
	}
}
