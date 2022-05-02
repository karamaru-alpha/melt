//go:build tools
// +build tools

package tools

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/google/wire/cmd/wire"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
