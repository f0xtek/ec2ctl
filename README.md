# ec2ctl

`ec2ctl` is a command line tool, written in Go, to manage the power state of an EC2 instance, given its 'Name' tag.

This was born out of frustration from having to either log into the AWS Console, or use multiple AWS CLI commands to simply power on/off an EC2 instance.

## Prerequisites

The following must be configured before you are able to use this tool:

Linux/Mac:

- `~/.aws/config`
- `~/.aws/credentials`

Windows:

- `%HOME\.aws\config`
- `%HOME\.aws\credentials`

## Installation

### Building From Source

To build the tool from source, clone this repo, ensure that you have Go >=1.10 installed and get the following dependencies:

```bash
$ go get -u gitub.com/drmarconi/ec2ctl
$ cd %GOPATH/src/github.com/drmarconi/ec2ctl
$ go get -u github.com/aws/aws-sdk-go/aws && \
github.com/aws/aws-sdk-go/aws/aws/err && \
github.com/aws/aws-sdk-go/service/ec2 && \
github.com/spf13/cobra && \
github.com/spf13/viper
$ go install
```

The `ec2ctl` command will be installed in `$GOPATH/bin`.

## Usage

```bash
$ ec2ctl
  Manages the power state for an EC2 instance, given its 'Name' tag.
  For example:

          ec2ctl start --name MyInstance

  Usage:
    ec2ctl [command]

  Available Commands:
    help        Help about any command
    start       Start an EC2 instance given its 'Name' tag.
    stop        Stop an EC2 instance given its 'Name' tag.

  Flags:
    -h, --help             help for ec2ctl
    -n, --name string      Specify the value for the EC2 'Name' tag
    -p, --profile string   Specify the AWS CLI profile to use for authentication. (default "default")

  Use "ec2ctl [command] --help" for more information about a command.
```

### Starting an EC2 Instance

To start an EC2 instance, run the `ec2ctl start` command, and pass in the AWS profile and Name tag for the instance:

```bash
ec2ctl start --profile default-profile --name MyInstance
```

### Stopping an EC2 Instance

Stopping an EC2 instance is exactly the same as starting, but instead using the `ec2ctl stop` command:

```bash
ec2ctl stop --profile default-profile --name MyInstance
```
