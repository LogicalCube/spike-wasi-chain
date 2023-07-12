GOOS=js
GOARCH=wasm
ENTRY=main
BINPATH=$(ENTRY).wasm

all: build

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build \
		-ldflags="-s -w" \
		-trimpath -o $(BINPATH) \
		$(ENTRY).go
	cp $(ENTRY).wasm ../wasi-go-server/$(ENTRY)-go.wasm
