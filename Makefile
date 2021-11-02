SHELL = /bin/bash

all: gen build

gen:
	go generate ./...
	go mod tidy

build:
	go build ./...

tools:
	go install github.com/go-bindata/go-bindata/go-bindata

.PHONY: gen build tools
