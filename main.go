package main

import (
	"fmt"
	"os"

	"github.com/preslavmihaylov/todocheck/checker"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/fetcher"
	"github.com/preslavmihaylov/todocheck/issuetracker"
	"github.com/preslavmihaylov/todocheck/matchers"
	"github.com/preslavmihaylov/todocheck/traverser/comments"
)

var authToken = "SECRET"

// TODO:
// * Extract auth token to ~/.config/todocheck/auth.yaml
// * Add a --closes option which indicates that an issue is to be closed as a result of a PR
// * Add github integration
// * specify basepath via a parameter
// * Add caching for task statuses
func main() {
	rcConfig, err := config.NewTodocheckRC(".todocheckrc")
	if err != nil {
		fmt.Printf("couldn't open .todocheckrc file: %s\n", err)
		os.Exit(1)
	}

	todoErrs := []error{}
	chk := checker.New(
		fetcher.NewFetcher(rcConfig.Origin, authToken, issuetracker.FromString(rcConfig.Type)))

	commentsTraverser := comments.New(func(comment, filepath string, lines []string, linecnt int) error {
		chk.SetMatcher(matchers.ForFile(filepath))
		if !chk.IsTODO(comment) {
			return nil
		}

		todoErr, err := chk.Check(comment, filepath, lines, linecnt)
		if err != nil {
			return fmt.Errorf("couldn't check todo line: %w", err)
		} else if todoErr != nil {
			todoErrs = append(todoErrs, todoErr)
		}

		return nil
	})

	err = commentsTraverser.TraversePath(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(todoErrs) > 0 {
		printTodoErrs(todoErrs)
		os.Exit(1)
	}
}

func printTodoErrs(errs []error) {
	for _, err := range errs {
		fmt.Println(err.Error())
	}
}
