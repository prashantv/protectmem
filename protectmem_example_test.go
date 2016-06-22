package protectmem_test

import (
	"unsafe"

	"github.com/prashantv/protectmem"
)

func Example() {
	type S struct {
		F1 string
		F2 string
	}

	// Allocate S using protectmem
	sAlloc := protectmem.Allocate(unsafe.Sizeof(S{}))
	sPtr := (*S)(sAlloc.Ptr())

	// sPtr can be used as a normal Go struct
	sPtr.F1 = "f1"
	sPtr.F2 = "more data"

	// Now you can protect the memory from being modified
	sAlloc.Protect(protectmem.Read)

	// Reads are fine:
	_ = sPtr.F1

	// Writes will cause a fault
	sPtr.F1 = "this will fail"

	// You can also protect from both reads and writes
	sAlloc.Protect(protectmem.None)

	// Even reads will fail.
	_ = sPtr.F1
}
