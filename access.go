package protectmem

import (
	"fmt"
	"syscall"
)

// Access returns the access level for an allocation.
type Access int

// List of different access bits that can be ORed together.
const (
	None Access = (1 << iota)
	Read
	Write
)

func (a Access) toSyscallProt() int {
	var prot int
	if a&Read != 0 {
		prot = prot | syscall.PROT_READ
	}
	if a&Write != 0 {
		prot = prot | syscall.PROT_WRITE
	}
	return prot
}

func (a Access) String() string {
	switch a {
	case None:
		return "none"
	case Read:
		return "read"
	case Write:
		return "write"
	case Read | Write:
		return "read+write"
	}

	return fmt.Sprintf("Access(%v)", int(a))
}
