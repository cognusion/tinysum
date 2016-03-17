// Make a CRC32 checksum (with or without an offset) of a string or []byte
package tinysum

import (
	"fmt"
	"hash/crc32"
	"hash/crc64"
)

// Returns the 32bit sum of b
func Sum32(b []byte) (sum uint64) {
	h := crc32.NewIEEE()
	h.Write(b)
	sum = uint64(h.Sum32())
	return
}

// Returns the 64bit sum of b
func Sum64(b []byte) (sum uint64) {
	sum = crc64.Checksum(b, crc64.MakeTable(crc64.ECMA))
	return
}

// Returns an offset CRC32 of the provided string
func OffsetStringSum(s string, offset uint32) (tiny string) {
	tiny = OffsetByteSum([]byte(s), offset)
	return
}

// Returns an offset CRC64 of the provided string
func OffsetStringSum64(s string, offset uint32) (tiny string) {
	tiny = OffsetByteSum64([]byte(s), offset)
	return
}

// Returns an offset CRC32 of the provided []byte
func OffsetByteSum(b []byte, offset uint32) (tiny string) {
	s64 := Sum32(b) + uint64(offset)
	tiny = fmt.Sprintf("%x", s64)
	return
}

// Returns an offset CRC64 of the provided []byte
func OffsetByteSum64(b []byte, offset uint32) (tiny string) {
	s64 := Sum64(b) + uint64(offset)
	tiny = fmt.Sprintf("%x", s64)
	return
}

// Returns a CRC32 of the provided string
func StringSum(s string) (tiny string) {
	s64 := Sum32([]byte(s))
	tiny = fmt.Sprintf("%x", s64)
	return
}

// Returns a CRC64 of the provided string
func StringSum64(s string) (tiny string) {
	s64 := Sum64([]byte(s))
	tiny = fmt.Sprintf("%x", s64)
	return
}

// Returns a CRC32 of the provided []byte
func ByteSum(b []byte) (tiny string) {
	s64 := Sum32(b)
	tiny = fmt.Sprintf("%x", s64)
	return
}

// Returns a CRC64 of the provided []byte
func ByteSum64(b []byte) (tiny string) {
	s64 := Sum64(b)
	tiny = fmt.Sprintf("%x", s64)
	return
}
