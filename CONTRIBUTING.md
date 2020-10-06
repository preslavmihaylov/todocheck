# Contributing

If you discover issues, have ideas for improvements or new features,
please report them to the [issue tracker](https://github.com/preslavmihaylov/todocheck/issues) of the repository or
submit a pull request. Please, try to follow these guidelines when you do so.

## Issue reporting

* Check that the issue has not already been reported.
* Check that the issue has not already been fixed in the latest code
  (a.k.a. `master`).
* Be clear, concise and precise in your description of the problem.
* Open an issue with a descriptive title and a summary in grammatically correct,
  complete sentences.
* Include any relevant code to the issue summary.

## Pull requests

* Fork the project.
* Write [good commit messages](https://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html).
* Use the [standard Go coding conventions and idioms](https://golang.org/doc/effective_go.html)
* If your change has a corresponding open GitHub issue, add the following in the description `closes #issue-number`.
* If the feature you are developing has external user impact (i.e. not a refactoring or internal improvement), then make sure to add appropriate tests for it.
  * The project has a set of integration tests, which run the actual program binary & test the output. Refer to `testing/todocheck_test.go` for examples.
  * If needed, extend the `testing/scenariobuilder` to incorporate your testing needs
* Make sure the test suite is passing locally. It will otherwise fail on CI nevertheless.
* Open a [pull request](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/about-pull-requests) that relates to *only* one subject with a clear title
  and description in grammatically correct, complete sentences.

## Local Development
The project uses standard Go tooling for build & run.

Testing is done via a custom framework, located in `testing/scenariobuilder`, which allows you to write integration tests, which run the `todocheck` binary & verify the program's output.

Build the project:
```bash
make build
```

Run the tests:
```bash
make test
```

Run `todocheck` locally:
```
go run main.go
```

Alternatively, copy the build output binary `./todocheck` to the relevant location and run from there.

Bundle binaries for next `todocheck` release:
```bash
./release.sh  <target-dir> <next-version>
```
