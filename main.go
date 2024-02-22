// Run with go generate . && go run .
//
//go:generate cython demo.pyx
package main

// #cgo pkg-config: python-3.10-embed
// #include <Python.h>
// #include "demo.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	demoC := C.CString("demo")
	defer C.free(unsafe.Pointer(demoC))

	// we need to register the module ...

	C.PyImport_AppendInittab(demoC, (*[0]byte)(C.PyInit_demo))
	C.Py_Initialize()
	defer C.Py_Finalize()

	// and import the module
	moduleC := C.CString("demo")
	defer C.free(unsafe.Pointer(moduleC))
	C.PyImport_ImportModule(moduleC)

	// otherwise cython code that invokes other python code will crash
	fmt.Printf("hello() -> %d\n", C.hello())
}
