package protectmem

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"
)

var _defaultAccess = Read | Write

// Allocate allocates space with the specified size that can be protected in future.
// The size is typically specified using unsafe.Sizeof
func Allocate(size uintptr) *Allocation {
	data := allocate(size, _defaultAccess)
	return &Allocation{_defaultAccess, data}
}

// AllocateSlice allocates the specified capacity slice and modifies the given ptr.
func AllocateSlice(slicePtr interface{}, capacity int) *Allocation {
	slicePtrType := reflect.TypeOf(slicePtr)
	if slicePtrType.Kind() != reflect.Ptr {
		panic("AllocateSlice: slicePtr must be a pointer to a slice")
	}
	sliceType := slicePtrType.Elem()
	if sliceType.Kind() != reflect.Slice {
		panic("AllocateSlice: slicePtr must be a pointer to a slice")
	}

	size := reflect.ArrayOf(capacity, sliceType.Elem()).Size()
	alloc := Allocate(size)

	// The slicePtr is actually two words: typeTable, sliceHeaderPtr
	slice := (*reflect.SliceHeader)((*[2]unsafe.Pointer)(unsafe.Pointer(&slicePtr))[1])

	// Set the underlying fields in the slice header.
	slice.Len = 0
	slice.Cap = capacity
	slice.Data = uintptr(alloc.Ptr())

	return alloc
}

func allocate(size uintptr, access Access) []byte {
	data, err := syscall.Mmap(-1 /* fd */, 0, int(size), access.toSyscallProt(), syscall.MAP_PRIVATE|syscall.MAP_ANON)
	if err != nil {
		panic(fmt.Errorf("failed to allocate memory using mmap: %v", err))
	}

	return data
}
