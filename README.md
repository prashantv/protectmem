# protectmem [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]

A Golang library to allocate memory that can have access protection changed after allocation

### Not intended for production ###

This library is not intended for producton code, it uses `unsafe` heavily and
relies on implementation details of how slices and interfaces are represented
in memory.


[doc-img]: https://godoc.org/github.com/prashantv/protectmem?status.svg
[doc]: https://godoc.org/github.com/prashantv/protectmem
[ci-img]: https://travis-ci.org/prashantv/protectmem.svg?branch=master
[ci]: https://travis-ci.org/prashantv/protectmem
[cov-img]: https://coveralls.io/repos/github/prashantv/protectmem/badge.svg?branch=master
[cov]: https://coveralls.io/github/prashantv/protectmem?branch=master
