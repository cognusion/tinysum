// Simple helper library for making CRC32 checksums
package tinysum

import (
	"fmt"
	"hash/crc32"
)

// Returns an offset CRC32 of the provided string
func OffsetStringSum(s string, offset uint32) (tiny string) {
	h := crc32.NewIEEE()
	h.Write([]byte(s))
	s64 := uint64(h.Sum32() + offset)
	tiny = fmt.Sprintf("%x", s64)
	return
}

// Returns an offset CRC32 of the provided []byte
func OffsetByteSum(b []byte, offset uint32) (tiny string) {
	h := crc32.NewIEEE()
	h.Write(b)
	s64 := uint64(h.Sum32() + offset)
	tiny = fmt.Sprintf("%x", s64)
	return
}

// Returns a CRC32 of the provided string
func StringSum(s string) (tiny string) {
	h := crc32.NewIEEE()
	h.Write([]byte(s))
	s64 := uint64(h.Sum32())
	tiny = fmt.Sprintf("%x", s64)
	return
}

// Returns a CRC32 of the provided []byte
func ByteSum(b []byte) (tiny string) {
	h := crc32.NewIEEE()
	h.Write(b)
	s64 := uint64(h.Sum32())
	tiny = fmt.Sprintf("%x", s64)
	return
}
