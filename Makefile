include .env
export

.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: go-generate
go-generate:
	go generate ./pkg/...

## git差分範囲で`go generate`する
.PHONY: go-generate-git-diff
go-generate-git-diff:
	sh ./scripts/go-generate-git-diff.sh
