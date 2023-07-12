# sudo apt install wabt

BIN=wasi-server

install:
	go mod tidy

run:
	go run main.go

build:
	go build -o $(BIN)

clean:
	rm $(BIN)

# ####################################
# Decompile WASM to WAT for debugging
WASMS=$(wildcard *.wasm)
WATS=$(WASMS:.wasm=.wat)
%.wat : %.wasm
	wasm2wat $< -o $@
decompile: $(WATS)
