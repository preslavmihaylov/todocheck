package main

import (
	"fmt"
	"os"

	"github.com/preslavmihaylov/todocheck/checker"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/fetcher"
	"github.com/preslavmihaylov/todocheck/issuetracker"
	"github.com/preslavmihaylov/todocheck/matchers"
	"github.com/preslavmihaylov/todocheck/traverser"
)

var authToken = "SECRET"

// TODO:
// * Extract auth token to ~/.config/todocheck/auth.yaml
// * Handle multi-line comments
// * Handle comment on current line
// * specify basepath via a parameter
// * Add github integration
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
	err = traverser.TraversePath(".", func(filename, line string, linecnt int) error {
		chk.SetMatcher(matchers.ForFile(filename))
		if !chk.IsTODO(line) {
			return nil
		}

		todoErr, err := chk.Check(filename, line, linecnt)
		if err != nil {
			return fmt.Errorf("couldn't check todo line: %w", err)
		} else if todoErr != nil {
			todoErrs = append(todoErrs, todoErr)
		}

		return nil
	})
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
