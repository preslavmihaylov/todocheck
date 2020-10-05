# Change Log

## 0.3.0 (2020-10-04)

### Features
 * Support for additional programming languages - Scala ([#32](https://github.com/preslavmihaylov/todocheck/issues/32))
 * Support for running `todocheck` on Windows ([#35](https://github.com/preslavmihaylov/todocheck/issues/35))
 * Support for using `todocheck` with a private gitlab server ([#40](https://github.com/preslavmihaylov/todocheck/issues/40))
 * Support for specifying issues with `#` for integrating with pivotal tracker ([#39](https://github.com/preslavmihaylov/todocheck/issues/39))
 * Improved validation of `.todocheck.yaml` configuration ([#25](https://github.com/preslavmihaylov/todocheck/issues/25), [#36](https://github.com/preslavmihaylov/todocheck/issues/36), [#44](https://github.com/preslavmihaylov/todocheck/issues/44))
 * `todocheck` now automatically detects issue tracker from git config if `.todocheck.yaml` is not explicitly provided for public github & gitlab repos ([#50](https://github.com/preslavmihaylov/todocheck/issues/50), [#58](https://github.com/preslavmihaylov/todocheck/issues/58))
 * Added `-v/--version` flag for showing currently installed `todocheck` version ([#67](https://github.com/preslavmihaylov/todocheck/issues/67))

### Bug fixes
N/A

### Breaking Changes
N/A

### Internal Improvements
 * Configured automatic CI pipeline for executing build & test on PRs ([#51](https://github.com/preslavmihaylov/todocheck/issues/51))
 * Wrote the project's [Contributing guide](./CONTRIBUTING.md) ([#24](https://github.com/preslavmihaylov/todocheck/issues/24))

## 0.2.0 (2020-08-01)

### Features
 * Add support for new programming languages - R ([#1](https://github.com/preslavmihaylov/todocheck/issues/1)), PHP ([#9](https://github.com/preslavmihaylov/todocheck/issues/9)), Rust ([#12](https://github.com/preslavmihaylov/todocheck/issues/12)), Swift ([#13](https://github.com/preslavmihaylov/todocheck/issues/13)), Groovy ([#14](https://github.com/preslavmihaylov/todocheck/issues/14))
 * Add support for new issue trackers - [Pivotal Tracker](https://www.pivotaltracker.com) ([#7](https://github.com/preslavmihaylov/todocheck/issues/7)), [Redmine](https://redmine.org/) ([#11](https://github.com/preslavmihaylov/todocheck/issues/11))
 * Support setting Authentication Token via Environment Variable ([#2](https://github.com/preslavmihaylov/todocheck/issues/2))
 * Support machine-friendly output - JSON ([#3](https://github.com/preslavmihaylov/todocheck/issues/3))
 * Add project logo

### Bug fixes
N/A

### Breaking Changes
N/A
