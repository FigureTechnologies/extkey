GOLANGCI_VERSION = 1.47.2

SRC = $(shell find . -name *.go)
GO  = go

all: extkey

extkey: $(SRC)
	$(GO) build -o extkey ./cmd/extkey

.PHONY: test
test:
	$(GO) test -mod=readonly -race ./... 

.PHONY: docker
docker:
	docker build -f docker/Dockerfile -t provenanceio/extkey .

.PHONY: clean
clean:
	rm -f extkey

###########
# Linting #
###########
LINTER := $(shell command -v golangci-lint 2> /dev/null)
MISSPELL := $(shell command -v misspell 2> /dev/null)
GOIMPORTS := $(shell command -v goimports 2> /dev/null)

.PHONY: gofmt
gofmt:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*.pb.go" | xargs gofmt -s -w

.PHONY: check-goimports
check-goimports:
ifndef GOIMPORTS
	echo "Fetching goimports"
	go install golang.org/x/tools/cmd/goimports
endif

.PHONY: goimports
goimports: check-goimports
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*.pb.go" | xargs goimports -w -local github.com/provenance-io/cosmovisor

.PHONY: check-gomisspell
check-gomisspell:
ifndef MISSPELL
	echo "Fetching misspell"
	go install github.com/client9/misspell/cmd/misspell
endif

.PHONY: gomisspell
gomisspell: check-gomisspell
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*.pb.go" | xargs misspell -w

.PHONY: check-lint
check-lint:
ifndef LINTER
	echo "Fetching golangci-lint"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_VERSION)
endif

.PHONY: lint
lint: check-lint goimports gofmt gomisspell
	golangci-lint run
