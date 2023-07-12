# WASI Stuff

Playing around with server side WASM. Examples include:

- C program that compiles into a wasi modules
- Go application that loads the wasi file

## Build the C wasm

```
cd wasi-c99-client
make preflight
make build
```

If you don't get any errors, that should create _main.wasm_ in the 
_wasi-c99-client_ and copy that to the _wasi-go-server_ directory.

The _wasm_ file contains the function _sum_ which the go server 
will call

## Build the Go Client

```
cd wasi-go-server
make install
make run
```

If you don't get any errors, that will run the Go code, load the
_main.wasm_ file, instantiate and call the _sum_ function.

Depending on how you compile the C wasm file, you might need to modify
the go server code. Use `make decompile` to have a look at what the
wasm file requires for callbacks.

## Notes

- wat2wasm: translate from WebAssembly text format to the WebAssembly binary format
- wasm2wat: the inverse of wat2wasm, translate from the binary format back to the text format (also known as a .wat)
- wasm-objdump: print information about a wasm binary. It is similar to objdump.
- wasm-strip: remove sections of a WebAssembly binary file
- wasm-validate: validate a file in the WebAssembly binary format

