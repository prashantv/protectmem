package protectmem

import "syscall"

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
