// Package tinysum provides a number of quazi-primitives to assist with 32 and 64-bit CRC sums
package tinysum

import (
	"fmt"
	"hash/crc32"
	"hash/crc64"
)

// Sum32 returns the 32bit sum of b
func Sum32(b []byte) (sum uint64) {
	h := crc32.NewIEEE()
	h.Write(b)
	sum = uint64(h.Sum32())
	return
}

// Sum64 returns the 64bit sum of b
func Sum64(b []byte) (sum uint64) {
	sum = crc64.Checksum(b, crc64.MakeTable(crc64.ECMA))
	return
}

// OffsetStringSum returns an offset CRC32 of the provided string
func OffsetStringSum(s string, offset uint32) (tiny string) {
	tiny = OffsetByteSum([]byte(s), offset)
	return
}

// OffsetStringSum64 returns an offset CRC64 of the provided string
func OffsetStringSum64(s string, offset uint32) (tiny string) {
	tiny = OffsetByteSum64([]byte(s), offset)
	return
}

// OffsetByteSum returns an offset CRC32 of the provided []byte
func OffsetByteSum(b []byte, offset uint32) (tiny string) {
	s64 := Sum32(b) + uint64(offset)
	tiny = fmt.Sprintf("%x", s64)
	return
}

// OffsetByteSum64 returns an offset CRC64 of the provided []byte
func OffsetByteSum64(b []byte, offset uint32) (tiny string) {
	s64 := Sum64(b) + uint64(offset)
	tiny = fmt.Sprintf("%x", s64)
	return
}

// StringSum returns a CRC32 of the provided string
func StringSum(s string) (tiny string) {
	s64 := Sum32([]byte(s))
	tiny = fmt.Sprintf("%x", s64)
	return
}

// StringSum64 returns a CRC64 of the provided string
func StringSum64(s string) (tiny string) {
	s64 := Sum64([]byte(s))
	tiny = fmt.Sprintf("%x", s64)
	return
}

// ByteSum returns a CRC32 of the provided []byte
func ByteSum(b []byte) (tiny string) {
	s64 := Sum32(b)
	tiny = fmt.Sprintf("%x", s64)
	return
}

// ByteSum64 returns a CRC64 of the provided []byte
func ByteSum64(b []byte) (tiny string) {
	s64 := Sum64(b)
	tiny = fmt.Sprintf("%x", s64)
	return
}
