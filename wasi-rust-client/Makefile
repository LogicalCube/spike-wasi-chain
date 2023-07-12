# Need:
#	- wasmtime-cli ~10

WASM_RUNNER=wasmtime
ENTRY=main

build:
	rustup target add wasm32-wasi
	rustc $(ENTRY).rs \
		--target wasm32-wasi
	cp $(ENTRY).wasm ../wasi-go-server/$(ENTRY)-rust.wasm

run:
	$(WASM_RUNNER) main.wasm
