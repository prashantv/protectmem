# protectmem
A Golang library to allocate memory that can have access protection changed after allocation

### Not intended for production ###

This library is not intended for producton code, it uses `unsafe` heavily and
relies on implementation details of how slices and interfaces are represented
in memory.
