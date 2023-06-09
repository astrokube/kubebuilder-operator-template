//go:build tools
// +build tools

package tools

import (
	_ "github.com/AlekSi/gocov-xml"
	_ "github.com/axw/gocov/gocov"

	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "golang.org/x/tools/cmd/goimports"

	_ "github.com/goreleaser/goreleaser"
)
