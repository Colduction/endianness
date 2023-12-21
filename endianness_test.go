package endianness_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/colduction/endianness"
)

const TEST_STRING string = "[5,7,2,0,4]:cognito:dhtxdV5uXHBpbG9ibW5KXGZjIjYdXG1namRtbSIoHVxzc3BqZmNyRV8bOCJQQlhHRFswKzE3NSssMjI1HSUgaWwdMyAxMCwnNzUqLCssMi8rGyoiYXNpIDotMikvMi0vKy4weQ"

var (
	TEST_BE_UINT16_SLICE []uint16 = endianness.BigEndian.StringToUint16Slice(TEST_STRING)
	TEST_BE_UINT32_SLICE []uint32 = endianness.BigEndian.StringToUint32Slice(TEST_STRING)
	TEST_BE_UINT64_SLICE []uint64 = endianness.BigEndian.StringToUint64Slice(TEST_STRING)

	TEST_LE_UINT16_SLICE []uint16 = endianness.LittleEndian.StringToUint16Slice(TEST_STRING)
	TEST_LE_UINT32_SLICE []uint32 = endianness.LittleEndian.StringToUint32Slice(TEST_STRING)
	TEST_LE_UINT64_SLICE []uint64 = endianness.LittleEndian.StringToUint64Slice(TEST_STRING)
)

func BenchPerCoreConfigs(b *testing.B, f func(b *testing.B)) {
	b.Helper()
	coreConfigs := []int{1, 2, 4, 8, 12}
	for _, n := range coreConfigs {
		name := fmt.Sprintf("%d cores", n)
		b.Run(name, func(b *testing.B) {
			runtime.GOMAXPROCS(n)
			f(b)
		})
	}
}

func BenchmarkBEStringToUint16Slice(b *testing.B) {
	BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				endianness.BigEndian.StringToUint16Slice(TEST_STRING)
			}
		})
	})
}

func BenchmarkLEStringToUint16Slice(b *testing.B) {
	BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				endianness.LittleEndian.StringToUint16Slice(TEST_STRING)
			}
		})
	})
}

func BenchmarkBEStringToUint32Slice(b *testing.B) {
	BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				endianness.BigEndian.StringToUint32Slice(TEST_STRING)
			}
		})
	})
}

func BenchmarkLEStringToUint32Slice(b *testing.B) {
	BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				endianness.LittleEndian.StringToUint32Slice(TEST_STRING)
			}
		})
	})
}

func BenchmarkBEStringToUint64Slice(b *testing.B) {
	BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				endianness.BigEndian.StringToUint64Slice(TEST_STRING)
			}
		})
	})
}

func BenchmarkLEStringToUint64Slice(b *testing.B) {
	BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				endianness.LittleEndian.StringToUint64Slice(TEST_STRING)
			}
		})
	})
}

func BenchmarkBEUint16SliceToString(b *testing.B) {
	BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				endianness.BigEndian.Uint16SliceToString(TEST_BE_UINT16_SLICE)
			}
		})
	})
}

func BenchmarkLEUint16SliceToString(b *testing.B) {
	BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				endianness.LittleEndian.Uint16SliceToString(TEST_LE_UINT16_SLICE)
			}
		})
	})
}

func BenchmarkBEUint32SliceToString(b *testing.B) {
	BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				endianness.BigEndian.Uint32SliceToString(TEST_BE_UINT32_SLICE)
			}
		})
	})
}

func BenchmarkLEUint32SliceToString(b *testing.B) {
	BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				endianness.LittleEndian.Uint32SliceToString(TEST_LE_UINT32_SLICE)
			}
		})
	})
}

func BenchmarkBEUint64SliceToString(b *testing.B) {
	BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				endianness.BigEndian.Uint64SliceToString(TEST_BE_UINT64_SLICE)
			}
		})
	})
}

func BenchmarkLEUint64SliceToString(b *testing.B) {
	BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				endianness.LittleEndian.Uint64SliceToString(TEST_LE_UINT64_SLICE)
			}
		})
	})
}
