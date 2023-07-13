package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/bytecodealliance/wasmtime-go/v9"
)

// makeImportMap creates a map of all the functions we want to import into the wasm modules
func makeImportMap(store *wasmtime.Store) map[string]*wasmtime.Func {
	var imports = map[string]*wasmtime.Func{
		"fd_write":          wasmtime.WrapFunc(store, func(int32, int32, int32, int32) int32 { return 0 }),
		"environ_get":       wasmtime.WrapFunc(store, func(int32, int32) int32 { return 0 }),
		"environ_sizes_get": wasmtime.WrapFunc(store, func(int32, int32) int32 { return 0 }),
		"proc_exit":         wasmtime.WrapFunc(store, func(int32) {}),
	}
	return imports
}

// walkPlugins walks the plugins directory and returns a list of all wasm files
func walkPlugins(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() && path.Ext(p) == ".wasm" {
			files = append(files, p)
		}
		return nil
	})
	return files, err
}

func main() {
	// Almost all operations in wasmtime require a contextual `store`
	// argument to share, so create that first
	store := wasmtime.NewStore(wasmtime.NewEngine())

	// get all the wasm files
	files, err := walkPlugins("plugins")
	if err != nil {
		panic(err)
	}

	// create a list of all module's functions we want to fake import
	importModules := makeImportMap(store)

	var plugins = make(map[string]*wasmtime.Instance, len(files))
	// try to load every wasm file as a module
	for i := 0; i < len(files); i++ {
		wasm, err := os.ReadFile(files[i])
		if err != nil {
			panic(err)
		}

		// Once we have our binary `wasm` we can compile that into a `*Module`
		module, err := wasmtime.NewModule(store.Engine, wasm)
		if err != nil {
			panic(err)
		}

		modules := module.Imports()
		// fake out the import functions the wasm module is looking for
		// Rust and C make some system level functions that must exist
		externFunks := make([]wasmtime.AsExtern, len(modules))
		for i := 0; i < len(modules); i++ {
			funcName := *modules[i].Name()
			externFunks[i] = importModules[funcName]
		}

		// Instantiate a module, and link in all our fake imports
		instance, err := wasmtime.NewInstance(store, module, externFunks)
		if err != nil {
			panic(err)
		}

		// If we make it this far, the instance is made
		plugins[files[i]] = instance
	}

	// Loop over all the loaded wasm files and try to call `sum`
	// which we are just assuming exists at this point.
	for i := range plugins {
		instance := plugins[i]

		// lookup our `sum` function and call it.
		run := instance.GetFunc(store, "sum")
		if run == nil {
			panic("not a function")
		}

		// Here we actually call the sum function...
		val, err := run.Call(store, 3, 4)
		if err != nil {
			panic(err)
		}

		fmt.Printf("sum from %s module ---> 3 + 4 = %v\n", i, val)
	}

}
