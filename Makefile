TOP_DIR := $(shell pwd)

GONAME=$(shell basename "$(PWD)")
PID=/tmp/go-$(GONAME).pid
GO ?= go
GO_TEST_ARGS ?= "-count=1"

SRCFILES := $(shell GOPATH=$(GOPATH) $(GO) list ./bls ./bn256)

# Make the default target (first target) to all.
default: all

.PHONY: all
all: test

.PHONY: lint
lint: gofmt

.PHONY: gofmt
gofmt:
	$(GO) fmt $(SRCFILES)

.PHONY: vet
vet:
	$(GO) vet $(SRCFILES)

.PHONY: test
test: lint vet
	@GOPATH=$(GOPATH) $(GO) test -cover $(GO_TEST_ARGS) $(SRCFILES)

# Clean all generated files
.PHONY: clean
clean:
	rm -rf bin/
	rm -rf pkg/
