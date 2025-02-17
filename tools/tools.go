//go:build tools
// +build tools

package main

import (
	_ "go.opentelemetry.io/collector/cmd/builder" // go.mod

	_ "go.opentelemetry.io/collector/cmd/mdatagen" // go.mod

	_ "golang.org/x/tools/cmd/goimports" // go.mod

	_ "honnef.co/go/tools/cmd/staticcheck" // go.mod
)
