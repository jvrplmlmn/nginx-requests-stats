APPLICATION ?= $$(basename $(CURDIR))

# --vendor: Enable vendoring support (skips 'vendor' directories and sets GO15VENDOREXPERIMENT=1).
# --tests: Include test files for linters that support this option
# --errors: Only show errors.
GOMETALINTER_REQUIRED_FLAGS := --vendor --tests --errors

# gotype is broken, see https://github.com/alecthomas/gometalinter/issues/91
# --deadline: Cancel linters if they have not completed within this duration.
# --line-length: Report lines longer than N (using lll).
GOMETALINTER_COMMON_FLAGS := --concurrency 2 --deadline 60s --line-length 120 --enable lll --disable gotype

.PHONY: lint
lint:
	$(GOPATH)/bin/gometalinter \
		$(GOMETALINTER_COMMON_FLAGS) \
		$(GOMETALINTER_REQUIRED_FLAGS) \
		.

.PHONY: check
check:
	$(GOPATH)/bin/gometalinter \
		--enable goimports \
		--disable errcheck \
		--disable golint \
		--fast \
		$(GOMETALINTER_COMMON_FLAGS) \
		$(GOMETALINTER_REQUIRED_FLAGS) \
		.

.PHONY: getdeps
getdeps:
	$(GOPATH)/bin/glide install

.PHONY: test
test: lint getdeps
	go test -cover $(shell glide novendor)

.PHONY: build
build: getdeps
	mkdir -p $(CURDIR)/pkg
	$(GOPATH)/bin/gox \
	  -osarch="linux/amd64 darwin/amd64" \
		-output "pkg/$(APPLICATION)_{{.OS}}_{{.Arch}}"

PACKAGES := \
	golang.org/x/tools/cmd/goimports \
	github.com/tools/godep \
	github.com/alecthomas/gometalinter \
	github.com/Masterminds/glide \
	github.com/mitchellh/gox

.PHONY: install-tools
install-tools:
	$(foreach pkg,$(PACKAGES),go get -u $(pkg);)
	$(GOPATH)/bin/gometalinter --install --update
