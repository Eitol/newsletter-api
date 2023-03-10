//go:build tools
// +build tools

//nolint:typecheck // These imports make sense for tools builds
package tools

import (
	_ "github.com/go-task/task/v3/cmd/task"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/pressly/goose/v3/cmd/goose"
	_ "github.com/swaggo/swag/cmd/swag"
	_ "gotest.tools/gotestsum"
	_ "mvdan.cc/gofumpt"
)
