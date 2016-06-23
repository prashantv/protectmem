package protectmem

import (
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessSyscallProt(t *testing.T) {
	tests := []struct {
		access Access
		want   int
	}{
		{None, 0},
		{Read, syscall.PROT_READ},
		{Write, syscall.PROT_WRITE},
		{Read | Write, syscall.PROT_READ | syscall.PROT_WRITE},
		{8 /* unexpected */, 0},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.access.toSyscallProt(), "Mismatch for %v", tt.access)
	}
}

func TestAccessString(t *testing.T) {
	tests := []struct {
		access Access
		want   string
	}{
		{None, "none"},
		{Read, "read"},
		{Write, "write"},
		{Read | Write, "read+write"},
		{8 /* unexpected */, "Access(8)"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.access.String(), "Mismatch for %v", int(tt.access))
	}
}
