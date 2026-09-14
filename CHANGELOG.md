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
