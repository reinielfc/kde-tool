BIN := kde-tool
VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := -s -w -X main.version=$(VERSION)
TARGETS := linux-amd64 linux-arm64

.PHONY: build

build:
	@go build -o $(BIN) -ldflags "$(LDFLAGS)" .

clean:
	@rm -f $(BIN)
	@rm -rf dist/

release: $(TARGETS:%=dist/$(BIN)-$(VERSION)-%)

dist/$(BIN)-$(VERSION)-linux-%: | dist
	GOOS=linux GOARCH=$* go build -o $@ -ldflags "$(LDFLAGS)" .
