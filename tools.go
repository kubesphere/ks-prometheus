//+build tools

// Package tools tracks dependencies for tools that used in the build process.
// See https://github.com/golang/go/wiki/Modules
package stack

import (
	_ "github.com/brancz/gojsontoyaml"
	_ "github.com/google/go-jsonnet/cmd/jsonnet"
	_ "github.com/jsonnet-bundler/jsonnet-bundler/cmd/jb"
)
