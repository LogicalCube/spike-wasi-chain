# Need:
# 	- clang ~14
#	- wasmtime-cli ~0.38.1
CC=clang
WASM_RUNNER=wasmtime

LLVM_VERSION=14
WASM_SDK_URL=https://github.com/WebAssembly/wasi-sdk/releases/download

ENTRY=main

build:
#	-nodefaultlibs   : keeps the linker from trying to use -lgcc
#	-no-entry        : keeps the linker from needing a main
#	-export-dynamic  : make the EXPORTs work
#	-nostdlib        : if you want to do *everything* by hand
#	-Wl,--export-all : if you don't want to specify exports
	$(CC) \
		-v -nodefaultlibs \
		-Wl,--no-entry \
		-Wl,--export-dynamic \
		-O2 -s \
		-DNDEBUG \
		--target=wasm32-wasi \
		--sysroot=$(PWD)/wasi-sysroot -lc \
		-o $(ENTRY).wasm \
		$(PWD)/lib/wasi/libclang_rt.builtins-wasm32.a \
		$(ENTRY).c
	cp $(ENTRY).wasm ../wasi-go-server/$(ENTRY)-c99.wasm

run:
# 	Run the "main" function
	$(WASM_RUNNER) $(PWD)/$(ENTRY).wasm 
# 	Run the "sum" function
	$(WASM_RUNNER) $(PWD)/$(ENTRY).wasm --invoke sum 12 30

clean:
	rm $(ENTRY).wasm

# #########################################################################
# Run these to download the libraries needed to build wasm
preflight: download_sysroot download_builtins

download_sysroot:
# 	Makes the wasi-sysroot directory
#	This is needed to use <stdio.h> type stuff
	curl -LO $(WASM_SDK_URL)/wasi-sdk-$(LLVM_VERSION)/wasi-sysroot-$(LLVM_VERSION).0.tar.gz
	tar xzvf wasi-sysroot-$(LLVM_VERSION).0.tar.gz
	rm wasi-sysroot-$(LLVM_VERSION).0.tar.gz

download_builtins:
#	Makes the wasi-sysroot directory
	curl -LO $(WASM_SDK_URL)/wasi-sdk-$(LLVM_VERSION)/libclang_rt.builtins-wasm32-wasi-$(LLVM_VERSION).0.tar.gz
	tar xzvf libclang_rt.builtins-wasm32-wasi-$(LLVM_VERSION).0.tar.gz
	rm libclang_rt.builtins-wasm32-wasi-$(LLVM_VERSION).0.tar.gz
