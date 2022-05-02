include .env
export

## API起動
.PHONY: run
run:
	docker-compose up --build

## 全てのテストを走らせる
.PHONY: test
test:
	go test ./pkg/... ./cmd/...

## `go generate`する
.PHONY: go-generate
go-generate:
	go generate ./pkg/...

## git差分範囲で`go generate`する
.PHONY: go-generate-git-diff
go-generate-git-diff:
	sh ./scripts/go-generate-git-diff.sh

## proto定義ファイルから自動生成する
.PHONY: protoc
protoc:
	sh ./scripts/protoc.sh

## localで走らせるパッケージ群のinstall
.PHONY: local-install
local-install:
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/google/wire/cmd/wire@v0.5.0
	go install github.com/skeema/skeema@v1.7.1
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2

## DBの破壊的マイグレーション
.PHONY: force-migrate
force-migrate:
	sh ./scripts/force-migrate.sh

## docker内のsqlに入る
.PHONY: mysql
mysql:
	docker-compose exec mysql bash -c 'mysql -u$$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE'
