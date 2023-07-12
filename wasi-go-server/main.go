package main

import (
	"fmt"
	"os"

	"github.com/bytecodealliance/wasmtime-go/v9"
)

func makeImportMap(store *wasmtime.Store) map[string]*wasmtime.Func {
	var imports = map[string]*wasmtime.Func{
		"fd_write":          wasmtime.WrapFunc(store, func(int32, int32, int32, int32) int32 { return 0 }),
		"environ_get":       wasmtime.WrapFunc(store, func(int32, int32) int32 { return 0 }),
		"environ_sizes_get": wasmtime.WrapFunc(store, func(int32, int32) int32 { return 0 }),
		"proc_exit":         wasmtime.WrapFunc(store, func(int32) {}),
	}
	return imports
}

func main() {
	// Almost all operations in wasmtime require a contextual `store`
	// argument to share, so create that first
	store := wasmtime.NewStore(wasmtime.NewEngine())

	wasm, err := os.ReadFile("main-c99.wasm")
	if err != nil {
		panic(err)
	}

	// Once we have our binary `wasm` we can compile that into a `*Module`
	// which represents compiled JIT code.
	module, err := wasmtime.NewModule(store.Engine, wasm)
	if err != nil {
		panic(err)
	}

	// create a list of all modules we know about and want to support
	importModules := makeImportMap(store)

	modules := module.Imports()
	externFunks := make([]wasmtime.AsExtern, len(modules))
	for i := 0; i < len(modules); i++ {
		funcName := *modules[i].Name()
		// fmt.Printf("%v %v\n", modules[i].Module(), funcName)
		externFunks[i] = importModules[funcName]
	}

	// Next up we instantiate a module which is where we link in all our
	// imports. We've got one import so we pass that in here.
	// instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{item1, item2, item3, item4, item5})
	instance, err := wasmtime.NewInstance(store, module, externFunks)
	// -nostlib C can be created thusly...
	// instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		// panic(err)
	}

	// After we've instantiated we can lookup our `sum` function and call it.
	run := instance.GetFunc(store, "sum")
	if run == nil {
		panic("not a function")
	}

	// Here we actually call the sum function
	val, err := run.Call(store, 3, 4)
	if err != nil {
		panic(err)
	}
	fmt.Printf("sum from WASI module ---> 3 + 4 = %v\n", val)
}
