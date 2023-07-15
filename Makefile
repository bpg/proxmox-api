GOFMT_FILES?=$$(find . -name '*.go')

default: test

fmt:
	gofmt -s -w $(GOFMT_FILES)

test:
	go test -race ./...

lint:
	go run -modfile=tools/go.mod github.com/golangci/golangci-lint/cmd/golangci-lint run --fix

#release-build:
#	go run -modfile=tools/go.mod github.com/goreleaser/goreleaser build --clean --skip-validate

.PHONY: fmt targets test
