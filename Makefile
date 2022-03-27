## go generate
.PHONY: go-generate
go-generate:
	go generate ./pkg/...

## go generate git diff
.PHONY: go-generate-git-diff
go-generate-git-diff:
	sh ./scripts/go-generate-git-diff.sh
