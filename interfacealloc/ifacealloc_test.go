package interfacealloc

import "testing"

type MyByte byte
type MyStruct struct { b byte }

func BenchmarkIfaceAllocUInt64(b *testing.B) {
	var table [10]interface{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			table[(i*j) % 10] = uint64(i*j)
		}
	}
}

func BenchmarkIfaceAllocInt32(b *testing.B) {
	var table [10]interface{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			table[(i*j) % 10] = int32(i*j)
		}
	}
}

func BenchmarkIfaceAllocByte(b *testing.B) {
	var table [10]interface{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			table[(i*j) % 10] = byte(i*j)
		}
	}
}

func BenchmarkIfaceAllocInt16(b *testing.B) {
	var table [10]interface{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			table[(i*j) % 10] = int16(i*j)
		}
	}
}

func BenchmarkIfaceAllocRune(b *testing.B) {
	var table [10]interface{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			table[(i*j) % 10] = rune(i*j)
		}
	}
}

func BenchmarkIfaceAllocMyByte(b *testing.B) {
	var table [10]interface{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			table[(i*j) % 10] = MyByte(i*j)
		}
	}
}

func BenchmarkIfaceAllocMyStruct(b *testing.B) {
	var table [10]interface{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			table[(i*j) % 10] = MyStruct{b: byte(i*j)}
		}
	}
}
