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
