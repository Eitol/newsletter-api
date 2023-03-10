# Newsletter API

API for managing subscriptions to newsletters.

# Table of Contents
0. [Challenge](#challenge)
1. [How to run](#how-to-run)
    1. [Requirements](#requirements)
    2. [Install Task](#install-task)
    2. [Install tools](#install-tools)
    4. [Run](#run)
2. [Generate documentation](#generate-documentation)
3. [Testing](#testing)
4. [Linting](#linting)

# Challenge
([Back to Table of Contents](#table-of-contents))

1. Implement the endpoints described in the [API documentation](./docs/swagger.yaml) (see [Generate documentation](#generate-documentation))
2. Using the same design patterns, implement a POST endpoint for creating new subscriptions

## Requirements
([Back to Table of Contents](#table-of-contents))

1. This challenge is meant to be completed under 90 minutes
2. The code should be legible and easy to understand
3. The code should be tested

# How to run

## Requirements
([Back to Table of Contents](#table-of-contents))

- [Go](https://go.dev/)

## Install Task
([Back to Table of Contents](#table-of-contents))

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

## Install tools
([Back to Table of Contents](#table-of-contents))

```bash
task install
```

## Run
([Back to Table of Contents](#table-of-contents))

```bash
task run
```

# Generate documentation

```bash
task docs
```

# Testing

```bash
task test
```

# Linting

```bash
task lint
```