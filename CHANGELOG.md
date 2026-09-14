# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]


## [0.0.7] - 2022-11-08

### Added

- Docker support via provided Dockerfile
- Confirmation prompt before pushing changes to SSM Parameter Store
- Support for encrypted strings in parameters
- Persistence of encryption keys to disk

### Changed

- Directory configuration now uses `<dir-name>.yaml` instead of `.yaml` files
- Binary name changed
- Encryption keys are now sorted

### Security

- Updated AWS SDK for AWS SSO support

## [0.0.6] - 2018-11-13

_No user-facing changes were detected in the commits for this release._


## [0.0.5] - 2018-11-13

### Changed

- Require confirmation before pushing changes


## [0.0.4] - 2018-09-14

### Changed

- Encryption keys are now sorted for consistent ordering


## [0.0.3] - 2018-08-30

### Changed

- Binary name updated


## [0.0.2] - 2018-08-28

### Added

- Support for encrypted strings
- Store encryption key to disk

### Removed

- convertFromSSMParameters function


## [0.0.1] - 2018-08-27

### Added

- Initial release of ssm-parameter-store CLI tool

