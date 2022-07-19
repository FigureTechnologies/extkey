SRC = $(shell find . -name *.go)
GO  = go

all: extkey

extkey: $(SRC)
	$(GO) build -o extkey ./cmd/extkey

clean:
	rm -f extkey
