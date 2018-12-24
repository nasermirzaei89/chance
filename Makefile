export GO111MODULE=on
export PACKAGE=`go list`

default: build build-cmd lint test

build:
	go build

build-cmd:
	go build -o ./cmd/chance/chance ./cmd/chance

format:
	which goimports || go get -u -v golang.org/x/tools/cmd/goimports
	find $(PWD) -type f -name "*.go" | xargs -n 1 -I R goimports -w R
	find $(PWD) -type f -name "*.go" | xargs -n 1 -I R gofmt -s -w R

lint:
	go fmt $(PACKAGE)/... # fixme
	go vet $(PACKAGE)/... # fixme
	which gocyclo || go get -u github.com/fzipp/gocyclo
	find $(PWD) -type f -name "*.go" | xargs -n 1 -I {} gocyclo -over 15 {} # fixme
	which golint || go get -u golang.org/x/lint/golint
	go list $(PACKAGE)/... | xargs -n 1 -I {} golint {} # fixme
	which ineffassign || go get -u github.com/gordonklaus/ineffassign
	find $(PWD) -type f -name "*.go" | xargs -n 1 -I {} ineffassign {} # fixme
	which misspell || go get -u github.com/client9/misspell/cmd/misspell
	find $(PWD) -type f -name "*.go" | xargs -n 1 -I {} misspell {} # fixme

test:
	go test -v ./...