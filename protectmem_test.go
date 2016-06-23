package protectmem

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type S struct {
	F1 string
	F2 int
	F3 int
	f4 string
}

func TestAllocate(t *testing.T) {
	sAlloc := Allocate(unsafe.Sizeof(S{}))
	s := (*S)(sAlloc.Ptr())
	s.F1 = "test"
	s.f4 = "private"

	expected := &S{F1: "test", f4: "private"}
	assert.Equal(t, expected, s, "S mismatch")

	// Once we protect with read only access, reads should work, but writes should fail
	sAlloc.Protect(Read)
	assert.Equal(t, expected, s, "S mismatch")
}

func TestAllocateSlice(t *testing.T) {
	var bs []byte
	alloc := AllocateSlice(&bs, 5)
	assert.Equal(t, 0, len(bs), "Wrong length")
	assert.Equal(t, 5, cap(bs), "Wrong capacity")

	for i := 0; i < cap(bs); i++ {
		bs = append(bs, 1)
	}

	alloc.Protect(Read)
	assert.Equal(t, []byte{1, 1, 1, 1, 1}, bs, "Unexpected slice values")
}

func TestAllocateInvalidSize(t *testing.T) {
	// Underflow to the maximum value for a uintptr.
	maxUintPtr := uintptr(0)
	maxUintPtr--

	assert.Panics(t, func() {
		Allocate(maxUintPtr)
	})
}

func TestAllocateSliceNotPtr(t *testing.T) {
	assert.Panics(t, func() {
		var bs []byte
		AllocateSlice(bs, 5)
	})
}

func TestAllocateSliceNotSlice(t *testing.T) {
	assert.Panics(t, func() {
		var s S
		AllocateSlice(&s, 5)
	})
}

func TestAccess(t *testing.T) {
	var bs []byte
	alloc := AllocateSlice(&bs, 5)
	assert.Equal(t, Read|Write, alloc.Access(), "Invalid initial Access")
	alloc.Protect(Read)
	assert.Equal(t, Read, alloc.Access(), "Invalid Access after Protect")
}

func TestProtectFail(t *testing.T) {
	// Cannot protect an invalid allocation
	var alloc Allocation
	assert.Panics(t, func() {
		alloc.Protect(10000)
	})
}
