.PHONY: all
all: tidy add-copyright format lint cover build

# ==============================================================================
# Usage

define USAGE_OPTIONS

Options:
  DEBUG            Whether to generate debug symbols. Default is 0.
  VERSION          The version information compiled into binaries.
                   The default is obtained from gsemver or git.
  V                Set to 1 enable verbose build. Default is 0.
endef
export USAGE_OPTIONS

# ==============================================================================

# build options
ROOT_PACKAGE=github.com/mariner-group/marinerd


# includes
include scripts/make-rules/copyright.mk
include scripts/make-rules/common.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/tools.mk

## build: Build source code for host platform.
.PHONY: build
build:
	@$(MAKE) go.build

## build.multiarch: Build source code for multiple platforms. See option PLATFORMS.
.PHONY: build.multiarch
build.multiarch:
	@$(MAKE) go.build.multiarch

.PHONY: tidy
tidy:
	@$(GO) mod tidy

## add-copyright: Ensures source code files have copyright license headers.
.PHONY: add-copyright
add-copyright:
	@$(MAKE) copyright.add

## clean: Remove all files that are created by building.
.PHONY: clean
clean:
	@echo "===========> Cleaning all build output"
	@-rm -vrf $(OUTPUT_DIR)

## lint: Check syntax and styling of go sources.
.PHONY: lint
lint:
	@$(MAKE) go.lint

## cover: Run unit test and get test coverage.
.PHONY: cover
cover:
	@$(MAKE) go.test.cover

## release: Release marinerd
.PHONY: release
release:
	@$(MAKE) release.run

## format: Gofmt (reformat) package sources (exclude vendor dir if existed).
.PHONY: format
format: tools.verify.golines tools.verify.goimports
	@echo "===========> Formating codes"
	@$(FIND) -type f -name '*.go' | $(XARGS) gofmt -s -w
	@$(FIND) -type f -name '*.go' | $(XARGS) goimports -w -local $(ROOT_PACKAGE)
	@$(FIND) -type f -name '*.go' | $(XARGS) golines -w --max-len=120 --reformat-tags --shorten-comments --ignore-generated .
	@$(GO) mod edit -fmt

## verify-copyright: Verify the boilerplate headers for all files.
.PHONY: verify-copyright
verify-copyright:
	@$(MAKE) copyright.verify

## test: Run unit test.
.PHONY: test
test:
	@$(MAKE) go.test

## tools: install dependent tools.
.PHONY: tools
tools:
	@$(MAKE) tools.install

## help: Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"