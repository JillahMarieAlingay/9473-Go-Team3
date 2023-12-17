PROJECT_NAME := "finals-golang-copy"
PKG := "https://gitlab.com/Julliard_Macatuggal/finals-golang-copy/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

coverage: ## Generate global code coverage report
	./tools/coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	./tools/coverage.sh html;

dep: ## Get the dependencies
	@go get -v -d ./...

build: dep ## Build the binary file
	@go build -i -v $(PKG)

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)
