package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/preslavmihaylov/todocheck/checker"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/fetcher"
	"github.com/preslavmihaylov/todocheck/issuetracker"
	"github.com/preslavmihaylov/todocheck/matchers"
	"github.com/preslavmihaylov/todocheck/traverser/comments"
)

// TODO:
// * Add a --closes option which indicates that an issue is to be closed as a result of a PR
// * Add github integration
// * Add caching for task statuses
func main() {
	var basepath = flag.String("basepath", ".", "The path for the project to todocheck. Defaults to current directory")
	var cfgPath = flag.String("config", config.DefaultLocal, "The project configuration file to use")
	flag.Parse()

	localCfg, err := config.NewLocal(*cfgPath, *basepath)
	if err != nil {
		log.Fatalf("couldn't open configuration file: %s\n", err)
	}

	err = localCfg.Auth.AcquireToken()
	if err != nil {
		log.Fatalf("couldn't acquire token from config: %s\n", err)
	}

	baseURL, err := issuetracker.BaseURLFor(issuetracker.FromString(localCfg.IssueTrackerType), localCfg.Origin)
	if err != nil {
		log.Fatalf("couldn't get base url from origin %s & issue tracker %s: %s\n",
			localCfg.Origin, localCfg.IssueTrackerType, err)
	}

	todoErrs := []error{}
	chk := checker.New(
		fetcher.NewFetcher(baseURL, localCfg.Auth.Token, issuetracker.FromString(localCfg.IssueTrackerType)))

	commentsTraverser := comments.New(localCfg.IgnoredPaths, matchers.SupportedFileExtensions(),
		func(comment, filepath string, lines []string, linecnt int) error {
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

	err = commentsTraverser.TraversePath(*basepath)
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
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
