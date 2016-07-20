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


/***
BenchmarkPopCountWithLoop-4   	100000000	        20.5 ns/op
BenchmarkPopCountWithoutLoop-4	200000000	         7.86 ns/op
ok  	_/Users/yangfei/golang/work/ch2/ex3	4.434s
***/