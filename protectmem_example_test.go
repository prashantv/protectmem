package protectmem_test

import (
	"fmt"
	"net"
	"unsafe"

	"github.com/prashantv/protectmem"
)

func Example_allocateAndProtect() {
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
	fmt.Println(sPtr.F1)
}

func Example_allocateSliceAndProtect() {
	var ints []int

	// Initializes ints to be a slice with 10 capacity, 0 length.
	alloc := protectmem.AllocateSlice(&ints, 10)

	for i := 0; i < cap(ints); i++ {
		ints = append(ints, i)
	}

	// Make the slice read only, can still print the slice fine.
	alloc.Protect(protectmem.Read)
	fmt.Println(ints)

	// ints cannot be written to, the following would crash
	ints[1] = 5

	// We can also protect the slice from any reads or writes
	alloc.Protect(protectmem.None)

	// Now even a read will fail
	fmt.Println(ints)
}

func ExampleAllocate() {
	ipAlloc := protectmem.Allocate(unsafe.Sizeof(net.IPAddr{}))
	ip := (*net.IPAddr)(ipAlloc.Ptr())

	_, _ = ipAlloc, ip
}

func ExampleAllocateSlice() {
	var strings []string
	alloc := protectmem.AllocateSlice(&strings, 100)

	_ = alloc

	fmt.Println("length", len(strings), "cap", cap(strings))
	// Output:
	// length 0 cap 100
}

func ExampleAllocation_Access() {
	alloc := protectmem.Allocate(unsafe.Sizeof(net.IPAddr{}))

	// Access returns the permissions for the alloc.
	// By default, everything starts off with read + write.
	fmt.Println(alloc.Access())

	// Output:
	// read+write
}

func ExampleAllocation_Protect() {
	ipAlloc := protectmem.Allocate(unsafe.Sizeof(net.IPAddr{}))
	ip := (*net.IPAddr)(ipAlloc.Ptr())

	// Protect can be used to change the access permissions for the alloc.
	ipAlloc.Protect(protectmem.Read)

	// ip cannot be written to, any writes will crash.
	ip.Zone = "this will fail"
}

func ExampleAllocation_Ptr() {
	ipAlloc := protectmem.Allocate(unsafe.Sizeof(net.IPAddr{}))

	// Ptr returns an unsafe.Pointer which can be casted into a pointer
	// of the type specified in Allocate.
	ip := (*net.IP)(ipAlloc.Ptr())

	_, _ = ipAlloc, ip
}
