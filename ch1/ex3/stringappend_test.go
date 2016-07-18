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