SHELL = /bin/bash

all: gen build test

gen:
	go generate ./...
	go mod tidy

build:
	go build ./...

test:
	go test -race ./.

tools:
	go install github.com/go-bindata/go-bindata/go-bindata

.PHONY: gen build test tools
