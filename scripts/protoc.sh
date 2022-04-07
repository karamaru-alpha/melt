#!/bin/sh

# proto定義Dir
definition_dir=./proto/api

# serverコード自動生成Dir
server_dir=./pkg/domain

# 自動生成
protoc \
	--go_out=paths=source_relative:${server_dir} \
	--go-grpc_out=require_unimplemented_servers=false,paths=source_relative:${server_dir} \
	${definition_dir}/*.proto


# ライブラリ整備
go mod tidy
