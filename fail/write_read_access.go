package main

import (
	"fmt"
	"unsafe"

	"github.com/prashantv/protectmem"
)

type s struct {
	F1 string
	F2 string
	f3 int
}

func main() {
	alloc := protectmem.Allocate(unsafe.Sizeof(s{}))
	s := (*s)(alloc.Ptr())

	s.F1 = "f1"
	s.F2 = "f2"
	s.f3 = 1

	alloc.Protect(protectmem.Read)

	// Expect a crash now
	fmt.Println("crash", 28)
	s.F1 = "overwrite"
}
