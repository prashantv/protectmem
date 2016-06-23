// +build 386 amd64

package protectmem

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// On x86 systems, the MMU does not support write-only.
func TestReadAfterWriteOnly(t *testing.T) {
	alloc := Allocate(unsafe.Sizeof(S{}))
	s := (*S)(alloc.Ptr())

	s.F1 = "f1"
	s.F2 = 2
	s.f4 = "f4"

	alloc.Protect(Write)
	assert.Equal(t, "f1", s.F1, "F1 mismatch")
}
