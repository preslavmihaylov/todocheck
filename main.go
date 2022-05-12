package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/preslavmihaylov/todocheck/authmanager"
	todocheckerrors "github.com/preslavmihaylov/todocheck/checker/errors"
	"github.com/preslavmihaylov/todocheck/config"
	"github.com/preslavmihaylov/todocheck/fetcher"
	"github.com/preslavmihaylov/todocheck/issuetracker/factory"
	"github.com/preslavmihaylov/todocheck/logger"
	"github.com/preslavmihaylov/todocheck/traverser/todoerrs"
	"github.com/preslavmihaylov/todocheck/validation"
)

// set dynamically on build time. See Makefile for more info
var version string

// TODO:
// * Add a --closes option which indicates that an issue is to be closed as a result of a PR
// * Add caching for task statuses
func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var basepath = fs.String("basepath", ".", "The path for the project to todocheck. Defaults to current directory")
	var cfgPath = fs.String("config", "", "The project configuration file to use. Will use the one from the basepath if not specified")
	var format = fs.String("format", "standard", "The output format to use. Available formats - standard, json")
	var verboseRequested = fs.Bool("verbose", false, "Make todocheck more talkative")
	var versionRequested = fs.Bool("version", false, "Show the current version of todocheck")
	fs.BoolVar(versionRequested, "v", *versionRequested, "Show the current version of todocheck (shorthand)")

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	if args := fs.Args(); len(args) > 0 {
		log.Fatalf("Unexpected arguments: %s\n", args)
	}

	logger.Setup(*verboseRequested)

	if *versionRequested {
		fmt.Println(version)
		os.Exit(0)
	}

	localCfg, err := config.NewLocal(*cfgPath, *basepath)
	if err != nil {
		log.Fatalf("couldn't open configuration file: %s\n", err)
	}

	tracker, err := factory.NewIssueTrackerFrom(localCfg.IssueTracker, localCfg.Auth, localCfg.Origin)
	if err != nil {
		log.Fatalf("couldn't create new issue tracker: %s\n", err)
	}

	err = authmanager.AcquireToken(localCfg, tracker)
	if err != nil {
		log.Fatalf("couldn't acquire token from config: %s\n", err)
	}

	if errors := validation.Validate(localCfg, tracker); len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}

		os.Exit(1)
	}

	f := fetcher.NewFetcher(tracker)

	todoErrs := []*todocheckerrors.TODO{}
	traverser := todoerrs.NewTraverser(f, localCfg.IgnoredPaths, localCfg.CustomTodos, localCfg.MatchCaseInsensitive, func(todoErr *todocheckerrors.TODO) error {
		todoErrs = append(todoErrs, todoErr)
		return nil
	})

	err = traverser.TraversePath(*basepath)
	if err != nil {
		log.Fatalf("couldn't traverse basepath: %s", err)
	}

	if len(todoErrs) > 0 {
		err = printTodoErrs(todoErrs, *format)
		if err != nil {
			panic(err)
		}
		os.Exit(2)
	}
}

func printTodoErrs(errs []*todocheckerrors.TODO, format string) error {
	if len(errs) == 0 {
		if format == "json" {
			fmt.Println("[]")
		}

		return nil
	}

	switch format {
	case "standard":
		for _, err := range errs {
			fmt.Fprintln(color.Error, err.Error())
		}
		return nil
	case "json":
		out := fmt.Sprintf("[%s", must(errs[0].ToJSON()))
		for _, err := range errs[1:] {
			out += fmt.Sprintf(",%s", must(err.ToJSON()))
		}

		out += "]"
		fmt.Println(out)
		return nil
	}

	return errors.New("unrecognized output format: " + format)
}

func must(bs []byte, err error) []byte {
	if err != nil {
		panic(err)
	}

	return bs
}
