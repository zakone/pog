package popcount

import (
	"testing"
)

func BenchmarkPopCountWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCountLoop(10)
	}
}

func BenchmarkPopCountWithoutLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCount(10)
	}
}

func BenchmarkPopCountWithoutTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCountOnlyLoop(10)
	}
}

func BenchmarkPCLastClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCountLastClear(10)
	}
}

/***
BenchmarkPopCountWithLoop-4    	100000000	        20.0 ns/op
BenchmarkPopCountWithoutLoop-4 	200000000	         8.70 ns/op
BenchmarkPopCountWithoutTable-4	10000000	       123 ns/op
BenchmarkPCLastClear-4         	300000000	         4.23 ns/op
ok  	_/Users/yangfei/golang/work/ch2/ex5	7.722s
***/