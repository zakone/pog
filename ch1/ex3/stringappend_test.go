//test for Golang string join speed
package main

import (
	"testing"
)

func BenchmarkAppendOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = appendOperator()
	}
}

func BenchmarkStirngsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stringsJoin()
	}
}

func BenchmarkAppendHardCoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = appendHardCoding()
	}
}

func BenchmarkByteArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = byteArray()
	}
}

func BenchmarkByteBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = byteBuffer()
	}
}

/***
BenchmarkAppendOperator-4  	 5000000	       301 ns/op
BenchmarkStirngsJoin-4     	10000000	       181 ns/op
BenchmarkAppendHardCoding-4	20000000	       102 ns/op
BenchmarkByteArray-4       	 5000000	       262 ns/op
BenchmarkByteBuffer-4      	 1000000	      1009 ns/op
ok  	_/Users/yangfei/golang/work/ch1/ex3	8.623s
***/