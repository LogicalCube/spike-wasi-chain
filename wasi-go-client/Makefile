###
# Note go < 1.21 wont compile to WASI, tinygo will though
# but tinygo requires 1.15 - 1.18 so choose your edge.
#
# error: requires go version 1.15 through 1.18, got go1.20
###

GO=tinygo
GOOS=js
GOARCH=wasm
ENTRY=main
BINPATH=$(ENTRY).wasm

all: build

build:
	@echo Read the Makefile
#	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build \
#		-ldflags="-s -w" \
#		-trimpath -o $(BINPATH) \
#		$(ENTRY).go
#
#	$(GO) build \
#		-o $(ENTRY).wasm \
#		-target=wasi \
#		$(ENTRY).go
#
#	cp $(ENTRY).wasm ../wasi-go-server/$(ENTRY)-go.wasm
