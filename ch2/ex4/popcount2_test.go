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

/***
BenchmarkPopCountWithLoop-4    	100000000	        20.0 ns/op
BenchmarkPopCountWithoutLoop-4 	200000000	         9.24 ns/op
BenchmarkPopCountWithoutTable-4	10000000	       124 ns/op
ok  	_/Users/yangfei/golang/work/ch2/ex4	6.183s
***/