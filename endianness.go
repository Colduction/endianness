// Package endianness for simple translation between numbers bytes sequences and other types written in Go
package endianness

import (
	"encoding/binary"
	"unsafe"

	"github.com/colduction/nocopy"
)

var LittleEndian littleEndian

type littleEndian struct{}

// Converts string to little-endian uint16 slice
func (littleEndian) StringToUint16Slice(s string) []uint16 {
	if len(s) <= 0 {
		return make([]uint16, 0)
	}
	bArr := []byte(s)
	return LittleEndian.bytesToUint16(bArr)
}

// Converts string to little-endian uint32 slice
func (littleEndian) StringToUint32Slice(s string) []uint32 {
	if len(s) <= 0 {
		return make([]uint32, 0)
	}
	bArr := []byte(s)
	return LittleEndian.bytesToUint32(bArr)
}

// Converts string to little-endian uint64 slice
func (littleEndian) StringToUint64Slice(s string) []uint64 {
	if len(s) <= 0 {
		return make([]uint64, 0)
	}
	bArr := []byte(s)
	return LittleEndian.bytesToUint64(bArr)
}

// Converts byte slice to little-endian uint16 slice
func (littleEndian) bytesToUint16(b []byte) []uint16 {
	length := len(b)
	if length == 0 {
		return make([]uint16, 0)
	}
	arr := make([]uint16, 0, length)
	for i, tmp := 0, uint16(0); i < length; i = i + 2 {
		tmp = binary.LittleEndian.Uint16(b[i : i+2])
		arr = append(arr, tmp)
	}
	if len(arr) != cap(arr) {
		arr = unsafe.Slice(unsafe.SliceData(arr), len(arr))
	}
	return arr
}

// Converts byte slice to little-endian uint32 slice
func (littleEndian) bytesToUint32(b []byte) []uint32 {
	length := len(b)
	if length == 0 {
		return make([]uint32, 0)
	}
	arr := make([]uint32, 0, length)
	for i, tmp := 0, uint32(0); i < length; i = i + 4 {
		tmp = binary.LittleEndian.Uint32(b[i : i+4])
		arr = append(arr, tmp)
	}
	if len(arr) != cap(arr) {
		arr = unsafe.Slice(unsafe.SliceData(arr), len(arr))
	}
	return arr
}

// Converts byte slice to little-endian uint64 slice
func (littleEndian) bytesToUint64(b []byte) []uint64 {
	length := len(b)
	if length == 0 {
		return make([]uint64, 0)
	}
	arr := make([]uint64, 0, length)
	for i, tmp := 0, uint64(0); i < length; i = i + 8 {
		tmp = binary.LittleEndian.Uint64(b[i : i+8])
		arr = append(arr, tmp)
	}
	if len(arr) != cap(arr) {
		arr = unsafe.Slice(unsafe.SliceData(arr), len(arr))
	}
	return arr
}

// Converts little-endian uint16 slice to string
func (littleEndian) Uint16SliceToString(a []uint16) string {
	length := len(a)
	if length == 0 {
		return ""
	}
	bArr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		bArr = binary.LittleEndian.AppendUint16(bArr, a[i])
	}
	return nocopy.ByteSliceToString(bArr)
}

// Converts little-endian uint32 slice to string
func (littleEndian) Uint32SliceToString(a []uint32) string {
	length := len(a)
	if length == 0 {
		return ""
	}
	bArr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		bArr = binary.LittleEndian.AppendUint32(bArr, a[i])
	}
	return nocopy.ByteSliceToString(bArr)
}

// Converts little-endian uint64 slice to string
func (littleEndian) Uint64SliceToString(a []uint64) string {
	length := len(a)
	if length == 0 {
		return ""
	}
	bArr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		bArr = binary.LittleEndian.AppendUint64(bArr, a[i])
	}
	return nocopy.ByteSliceToString(bArr)
}

// Converts little-endian uint16 slice to byte slice
func (littleEndian) Uint16SliceToBytes(a []uint16) []byte {
	length := len(a)
	if length == 0 {
		return make([]byte, 0)
	}
	bArr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		bArr = binary.LittleEndian.AppendUint16(bArr, a[i])
	}
	return bArr
}

// Converts little-endian uint32 slice to byte slice
func (littleEndian) Uint32SliceToBytes(a []uint32) []byte {
	length := len(a)
	if length == 0 {
		return make([]byte, 0)
	}
	bArr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		bArr = binary.LittleEndian.AppendUint32(bArr, a[i])
	}
	return bArr
}

// Converts little-endian uint64 slice to byte slice
func (littleEndian) Uint64SliceToBytes(a []uint64) []byte {
	length := len(a)
	if length == 0 {
		return make([]byte, 0)
	}
	bArr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		bArr = binary.LittleEndian.AppendUint64(bArr, a[i])
	}
	return bArr
}
