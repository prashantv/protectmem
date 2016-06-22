package protectmem

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"
)

// Allocation represents an allocation created by this package that can be protected.
type Allocation struct {
	access Access
	data   []byte
}

// Ptr returns an unsafe.Pointer that should be casted into the type that
// the user wants to use.
func (a *Allocation) Ptr() unsafe.Pointer {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&a.data))
	return unsafe.Pointer(sh.Data)
}

// Access returns the current access level for the page.
func (a *Allocation) Access() Access {
	return a.access
}

// Protect protectes the allocation with the given access settings.
func (a *Allocation) Protect(access Access) {
	if err := syscall.Mprotect(a.data, access.toSyscallProt()); err != nil {
		panic(fmt.Errorf("failed to Protect memory using mprotect: %v", err))
	}
	a.access = access
}
