package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/preslavmihaylov/todocheck/authmanager"
	"github.com/preslavmihaylov/todocheck/authmanager/authmiddleware"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/fetcher"
	"github.com/preslavmihaylov/todocheck/issuetracker"
	"github.com/preslavmihaylov/todocheck/traverser/todoerrs"
)

// TODO:
// * Add a --closes option which indicates that an issue is to be closed as a result of a PR
// * Add caching for task statuses
func main() {
	var basepath = flag.String("basepath", ".", "The path for the project to todocheck. Defaults to current directory")
	var cfgPath = flag.String("config", "", "The project configuration file to use. Will use the one from the basepath if not specified")
	flag.Parse()

	localCfg, err := config.NewLocal(*cfgPath, *basepath)
	if err != nil {
		log.Fatalf("couldn't open configuration file: %s\n", err)
	}

	err = authmanager.AcquireToken(localCfg)
	if err != nil {
		log.Fatalf("couldn't acquire token from config: %s\n", err)
	}

	baseURL, err := issuetracker.BaseURLFor(localCfg.IssueTracker, localCfg.Origin)
	if err != nil {
		log.Fatalf("couldn't get base url for origin %s & issue tracker %s: %s\n",
			localCfg.Origin, localCfg.IssueTracker, err)
	}

	f := fetcher.NewFetcher(baseURL, localCfg.IssueTracker, authmiddleware.For(localCfg))
	todoErrs := []error{}
	traverser := todoerrs.NewTraverser(f, localCfg.IgnoredPaths, func(todoErr error) error {
		todoErrs = append(todoErrs, todoErr)
		return nil
	})

	err = traverser.TraversePath(*basepath)
	if err != nil {
		log.Fatalf("couldn't traverse basepath: %s", err)
	}

	if len(todoErrs) > 0 {
		printTodoErrs(todoErrs)
		os.Exit(2)
	}
}

func printTodoErrs(errs []error) {
	for _, err := range errs {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
