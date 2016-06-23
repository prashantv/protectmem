package protectmem

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestProtectSegmentationFaults(t *testing.T) {
	const failDir = "./testdata/fail"
	files, err := ioutil.ReadDir(failDir)
	require.NoError(t, err, "ReadDir failed")

	for _, f := range files {
		checkCrash(t, filepath.Join(failDir, f.Name()))
	}
}

func checkCrash(t *testing.T, path string) {
	const crashPrefix = "crash "

	cmd := exec.Command("go", "run", path)
	out, err := cmd.CombinedOutput()

	// We expect an ExitError since the command should fail.
	if !assert.IsType(t, &exec.ExitError{}, err, "Expected ExitError") {
		return
	}

	// Output should contain a "crash" line
	lines := strings.Split(string(out), "\n")

	var found string
	for _, line := range lines {
		if strings.HasPrefix(line, crashPrefix) {
			found = line
		}
	}
	if !assert.NotEmpty(t, found, "Failed to find crash line in output: %s", out) {
		return
	}

	// Make sure the crash line matches
	line := strings.TrimPrefix(found, crashPrefix)
	lookFor := filepath.Base(path) + ":" + line
	assert.Contains(t, string(out), lookFor, "Unexpected panic")
}
