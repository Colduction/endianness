// Package for simple translation between numbers to strings, byte sequences, etc., and vice versa, written in Go
package endianness

import (
	"encoding/binary"
	"math"
	"unsafe"

	"github.com/colduction/nocopy"
)

var (
	BigEndian    bigEndian
	LittleEndian littleEndian
)

type (
	bigEndian    struct{}
	littleEndian struct{}
)

// Converts string to big-endian uint16 slice
func (bigEndian) StringToUint16Slice(s string) []uint16 {
	if len(s) <= 0 {
		return make([]uint16, 0)
	}
	bArr := []byte(s)
	return BigEndian.bytesToUint16(bArr)
}

// Converts string to big-endian uint32 slice
func (bigEndian) StringToUint32Slice(s string) []uint32 {
	if len(s) <= 0 {
		return make([]uint32, 0)
	}
	bArr := []byte(s)
	return BigEndian.bytesToUint32(bArr)
}

// Converts string to big-endian uint64 slice
func (bigEndian) StringToUint64Slice(s string) []uint64 {
	if len(s) <= 0 {
		return make([]uint64, 0)
	}
	bArr := []byte(s)
	return BigEndian.bytesToUint64(bArr)
}

// Converts string to big-endian float32 slice
func (bigEndian) StringToFloat32Slice(s string) []float32 {
	if len(s) <= 0 {
		return make([]float32, 0)
	}
	bArr := []byte(s)
	return BigEndian.bytesToFloat32(bArr)
}

// Converts string to big-endian float64 slice
func (bigEndian) StringToFloat64Slice(s string) []float64 {
	if len(s) <= 0 {
		return make([]float64, 0)
	}
	bArr := []byte(s)
	return BigEndian.bytesToFloat64(bArr)
}

// Converts byte slice to big-endian uint16 slice
func (bigEndian) bytesToUint16(b []byte) []uint16 {
	length := len(b)
	if length == 0 {
		return make([]uint16, 0)
	}
	arr := make([]uint16, 0, length)
	for i, tmp := 0, uint16(0); i < length; i = i + 2 {
		tmp = binary.BigEndian.Uint16(b[i : i+2])
		arr = append(arr, tmp)
	}
	if len(arr) != cap(arr) {
		arr = unsafe.Slice(unsafe.SliceData(arr), len(arr))
	}
	return arr
}

// Converts byte slice to big-endian uint32 slice
func (bigEndian) bytesToUint32(b []byte) []uint32 {
	length := len(b)
	if length == 0 {
		return make([]uint32, 0)
	}
	arr := make([]uint32, 0, length)
	for i, tmp := 0, uint32(0); i < length; i = i + 4 {
		tmp = binary.BigEndian.Uint32(b[i : i+4])
		arr = append(arr, tmp)
	}
	if len(arr) != cap(arr) {
		arr = unsafe.Slice(unsafe.SliceData(arr), len(arr))
	}
	return arr
}

// Converts byte slice to big-endian uint64 slice
func (bigEndian) bytesToUint64(b []byte) []uint64 {
	length := len(b)
	if length == 0 {
		return make([]uint64, 0)
	}
	arr := make([]uint64, 0, length)
	for i, tmp := 0, uint64(0); i < length; i = i + 8 {
		tmp = binary.BigEndian.Uint64(b[i : i+8])
		arr = append(arr, tmp)
	}
	if len(arr) != cap(arr) {
		arr = unsafe.Slice(unsafe.SliceData(arr), len(arr))
	}
	return arr
}

// Converts byte slice to big-endian float32 slice
func (bigEndian) bytesToFloat32(b []byte) []float32 {
	length := len(b)
	if length == 0 {
		return make([]float32, 0)
	}
	arr := make([]float32, 0, length)
	for i, tmp := 0, uint32(0); i < length; i = i + 4 {
		tmp = binary.BigEndian.Uint32(b[i : i+4])
		arr = append(arr, math.Float32frombits(tmp))
	}
	if len(arr) != cap(arr) {
		arr = unsafe.Slice(unsafe.SliceData(arr), len(arr))
	}
	return arr
}

// Converts byte slice to big-endian float64 slice
func (bigEndian) bytesToFloat64(b []byte) []float64 {
	length := len(b)
	if length == 0 {
		return make([]float64, 0)
	}
	arr := make([]float64, 0, length)
	for i, tmp := 0, uint64(0); i < length; i = i + 8 {
		tmp = binary.BigEndian.Uint64(b[i : i+8])
		arr = append(arr, math.Float64frombits(tmp))
	}
	if len(arr) != cap(arr) {
		arr = unsafe.Slice(unsafe.SliceData(arr), len(arr))
	}
	return arr
}

// Converts big-endian uint16 slice to string
func (bigEndian) Uint16SliceToString(a []uint16) string {
	length := len(a)
	if length == 0 {
		return ""
	}
	bArr := BigEndian.Uint16SliceToBytes(a)
	return nocopy.ByteSliceToString(bArr)
}

// Converts big-endian uint32 slice to string
func (bigEndian) Uint32SliceToString(a []uint32) string {
	length := len(a)
	if length == 0 {
		return ""
	}
	bArr := BigEndian.Uint32SliceToBytes(a)
	return nocopy.ByteSliceToString(bArr)
}

// Converts big-endian uint64 slice to string
func (bigEndian) Uint64SliceToString(a []uint64) string {
	length := len(a)
	if length == 0 {
		return ""
	}
	bArr := BigEndian.Uint64SliceToBytes(a)
	return nocopy.ByteSliceToString(bArr)
}

// Converts big-endian float64 slice to string
func (bigEndian) Float32SliceToString(a []float32) string {
	length := len(a)
	if length == 0 {
		return ""
	}
	bArr := BigEndian.Float32SliceToBytes(a)
	return nocopy.ByteSliceToString(bArr)
}

// Converts big-endian float64 slice to string
func (bigEndian) Float64SliceToString(a []float64) string {
	length := len(a)
	if length == 0 {
		return ""
	}
	bArr := BigEndian.Float64SliceToBytes(a)
	return nocopy.ByteSliceToString(bArr)
}

// Converts big-endian uint16 slice to byte slice
func (bigEndian) Uint16SliceToBytes(a []uint16) []byte {
	length := len(a)
	if length == 0 {
		return make([]byte, 0)
	}
	bArr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		bArr = binary.BigEndian.AppendUint16(bArr, a[i])
	}
	return bArr
}

// Converts big-endian uint32 slice to byte slice
func (bigEndian) Uint32SliceToBytes(a []uint32) []byte {
	length := len(a)
	if length == 0 {
		return make([]byte, 0)
	}
	bArr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		bArr = binary.BigEndian.AppendUint32(bArr, a[i])
	}
	return bArr
}

// Converts big-endian uint64 slice to byte slice
func (bigEndian) Uint64SliceToBytes(a []uint64) []byte {
	length := len(a)
	if length == 0 {
		return make([]byte, 0)
	}
	bArr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		bArr = binary.BigEndian.AppendUint64(bArr, a[i])
	}
	return bArr
}

// Converts big-endian float32 slice to byte slice
func (bigEndian) Float32SliceToBytes(a []float32) []byte {
	length := len(a)
	if length == 0 {
		return make([]byte, 0)
	}
	bArr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		bArr = binary.BigEndian.AppendUint32(bArr, math.Float32bits(a[i]))
	}
	return bArr
}

// Converts big-endian float64 slice to byte slice
func (bigEndian) Float64SliceToBytes(a []float64) []byte {
	length := len(a)
	if length == 0 {
		return make([]byte, 0)
	}
	bArr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		bArr = binary.BigEndian.AppendUint64(bArr, math.Float64bits(a[i]))
	}
	return bArr
}

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

// Converts string to little-endian float32 slice
func (littleEndian) StringToFloat32Slice(s string) []float32 {
	if len(s) <= 0 {
		return make([]float32, 0)
	}
	bArr := []byte(s)
	return LittleEndian.bytesToFloat32(bArr)
}

// Converts string to little-endian float64 slice
func (littleEndian) StringToFloat64Slice(s string) []float64 {
	if len(s) <= 0 {
		return make([]float64, 0)
	}
	bArr := []byte(s)
	return LittleEndian.bytesToFloat64(bArr)
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

// Converts byte slice to little-endian float32 slice
func (littleEndian) bytesToFloat32(b []byte) []float32 {
	length := len(b)
	if length == 0 {
		return make([]float32, 0)
	}
	arr := make([]float32, 0, length)
	for i, tmp := 0, uint32(0); i < length; i = i + 4 {
		tmp = binary.LittleEndian.Uint32(b[i : i+4])
		arr = append(arr, math.Float32frombits(tmp))
	}
	if len(arr) != cap(arr) {
		arr = unsafe.Slice(unsafe.SliceData(arr), len(arr))
	}
	return arr
}

// Converts byte slice to little-endian float64 slice
func (littleEndian) bytesToFloat64(b []byte) []float64 {
	length := len(b)
	if length == 0 {
		return make([]float64, 0)
	}
	arr := make([]float64, 0, length)
	for i, tmp := 0, uint64(0); i < length; i = i + 8 {
		tmp = binary.LittleEndian.Uint64(b[i : i+8])
		arr = append(arr, math.Float64frombits(tmp))
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
	bArr := LittleEndian.Uint16SliceToBytes(a)
	return nocopy.ByteSliceToString(bArr)
}

// Converts little-endian uint32 slice to string
func (littleEndian) Uint32SliceToString(a []uint32) string {
	length := len(a)
	if length == 0 {
		return ""
	}
	bArr := LittleEndian.Uint32SliceToBytes(a)
	return nocopy.ByteSliceToString(bArr)
}

// Converts little-endian uint64 slice to string
func (littleEndian) Uint64SliceToString(a []uint64) string {
	length := len(a)
	if length == 0 {
		return ""
	}
	bArr := LittleEndian.Uint64SliceToBytes(a)
	return nocopy.ByteSliceToString(bArr)
}

// Converts little-endian float64 slice to string
func (littleEndian) Float32SliceToString(a []float32) string {
	length := len(a)
	if length == 0 {
		return ""
	}
	bArr := LittleEndian.Float32SliceToBytes(a)
	return nocopy.ByteSliceToString(bArr)
}

// Converts little-endian float64 slice to string
func (littleEndian) Float64SliceToString(a []float64) string {
	length := len(a)
	if length == 0 {
		return ""
	}
	bArr := LittleEndian.Float64SliceToBytes(a)
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

// Converts little-endian float32 slice to byte slice
func (littleEndian) Float32SliceToBytes(a []float32) []byte {
	length := len(a)
	if length == 0 {
		return make([]byte, 0)
	}
	bArr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		bArr = binary.LittleEndian.AppendUint32(bArr, math.Float32bits(a[i]))
	}
	return bArr
}

// Converts little-endian float64 slice to byte slice
func (littleEndian) Float64SliceToBytes(a []float64) []byte {
	length := len(a)
	if length == 0 {
		return make([]byte, 0)
	}
	bArr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		bArr = binary.LittleEndian.AppendUint64(bArr, math.Float64bits(a[i]))
	}
	return bArr
}
