package stringutil

import (
	"testing"
)

// https://godoc.org/golang.org/x/tools/cmd/benchcmp

const (
	TEST_STRING          string = "!oG ,olleH"
	TEST_STRING_REVERSED string = "Hello, Go!"
)

func BenchmarkReverseRange(b *testing.B) {
	b.SetBytes(int64(len(TEST_STRING)))

	for i := 0; i < b.N; i++ {
		if r := ReverseRange(TEST_STRING); r != TEST_STRING_REVERSED {
			b.Errorf("result(%v) != TEST_STRING_REVERSED(%v)\n", r, TEST_STRING_REVERSED)
		}
	}

	//enable allocs report for a single test
	//equal to -benchmem
	//b.ReportAllocs()
}

// ~15% faster
func BenchmarkReverseConvert(b *testing.B) {
	b.SetBytes(int64(len(TEST_STRING)))

	for i := 0; i < b.N; i++ {
		if r := ReverseConvert(TEST_STRING); r != TEST_STRING_REVERSED {
			b.Errorf("result(%v) != TEST_STRING_REVERSED(%v)\n", r, TEST_STRING_REVERSED)
		}
	}
}

// trash, this solution is mostly for education purposes
func BenchmarkReverseDeferred(b *testing.B) {
	b.SetBytes(int64(len(TEST_STRING)))

	for i := 0; i < b.N; i++ {
		if r := ReverseDeferred(TEST_STRING); r != TEST_STRING_REVERSED {
			b.Errorf("result(%v) != TEST_STRING_REVERSED(%v)\n", r, TEST_STRING_REVERSED)
		}
	}
}

func BenchmarkParallelReverseRange(b *testing.B) {
	b.SetBytes(int64(len(TEST_STRING)))

	// -cpu 4
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ReverseRange(TEST_STRING)
		}
	})
}

func BenchmarkParallelReverseConvert(b *testing.B) {
	b.SetBytes(int64(len(TEST_STRING)))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ReverseConvert(TEST_STRING)
		}
	})
}

func BenchmarkParallelReverseDeferred(b *testing.B) {
	b.SetBytes(int64(len(TEST_STRING)))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ReverseDeferred(TEST_STRING)
		}
	})
}

// $ go test -bench Reverse -benchmem
// goos: linux
// goarch: amd64
// pkg: github.com/itnelo/stringutil
// BenchmarkReverseRange-8             	10000000	         127 ns/op	  78.54 MB/s	      64 B/op	       2 allocs/op
// BenchmarkReverseConvert-8           	20000000	         110 ns/op	  90.19 MB/s	      16 B/op	       1 allocs/op
// BenchmarkParallelReverseRange-8     	50000000	        34.8 ns/op	 287.47 MB/s	      64 B/op	       2 allocs/op
// BenchmarkParallelReverseConvert-8   	50000000	        30.1 ns/op	 331.78 MB/s	      16 B/op	       1 allocs/op
// PASS
// ok  	github.com/itnelo/stringutil	7.050s

// $ bc <<<"scale=5; 127 / 110 * 100 - 100"
// 15.45400

// $ bc <<<"scale=5; 331.78 / 287.47 * 100 - 100"
// 15.41300
