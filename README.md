# tinysum
Make a CRC32 or CRC64 checksum (with or without an offset) of a string or []byte

[![GoDoc](https://godoc.org/github.com/cognusion/tinysum?status.svg)](https://godoc.org/github.com/cognusion/tinysum)

```bash
go get github.com/cognusion/tinysum
```

## Why Offsets?
There may be times you don't want your checksum to be too predictable. If you
use (and store) an offset, it's still trivial for you to compare and validate,
but harder for someone else to validate it, even if they know- or think they 
know- the match string.