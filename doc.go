/*
Package protectmem provides functions to allocate memory that can be protected
after allocation. Protecting allows removing Write access or Read access after
the memory has been allocated.

This package is not intended to be used in production code, but is intended for
unit tests.
*/
package protectmem
