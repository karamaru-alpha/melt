#!/bin/sh

# proto定義Dir
definition_dir=./proto/api

# serverコード自動生成Dir
server_dir=./pkg/domain/proto

# 自動生成
protoc \
  --proto_path=./proto \
  --proto_path="${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.7" \
	--go_out=paths=source_relative:${server_dir} \
	--go-grpc_out=require_unimplemented_servers=false,paths=source_relative:${server_dir} \
	--validate_out=lang=go,paths=source_relative:${server_dir} \
	${definition_dir}/*.proto


# ライブラリ整備
go mod tidy
