# ssm-parameter-store
A CLI to pull, diff, and push SSM Parameter Store to and from disk


    Usage:
      ssm-parameter-store [flags]
      ssm-parameter-store [command]

    Available Commands:
      diff        diff SSM Parameters with those on disk
      help        Help about any command
      pull        pull SSM Parameters
      push        push SSM Parameters

    Flags:
      -d, --directory string   output directory (default "/home/justmiles/edo/parameter-store")
      -f, --format string      format type (default "yaml")
      -h, --help               help for ssm-parameter-store
      -p, --path strings       path (default [/])
          --version            version for ssm-parameter-store

    Use "ssm-parameter-store [command] --help" for more information about a command.

## Install
Download one of the [prebuilt binaries](https://github.com/justmiles/ssm-parameter-store/releases) add it to your system PATH.

## Examples

Pull from the Parameter Store to a local directory

    ssm-parameter-store pull

Diff changes on local disk with the remote Parameter Store

    ssm-parameter-store diff

Push the changes from local disk to remote Parameter Store

    ssm-paramter-store push

## Docker

When using Docker, mount the local parameter directory into the container and pass in AWS credentials, for example

    docker run -v <absolute path to a directory>:/my-parameters --rm -e AWS_DEFAULT_REGION -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY ssm-parameter-store diff -d my-parameters
