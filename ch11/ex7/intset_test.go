package intset

import (
    "math/rand"
    "testing"
    "time"
)

var seed = time.Now().UTC().UnixNano()

func BenchmarkIntSetAdd(b *testing.B) {
    rng := rand.New(rand.NewSource(seed))
    n := rng.Intn(10000)
    for i := 0; i < b.N; i++ {
        var x IntSet
        x.Add(n)
    }
}

func BenchmarkIntSetUnionWith(b *testing.B) {
    rng := rand.New(rand.NewSource(seed))
    m := rng.Intn(10000)
    n := rng.Intn(10000)
    for i := 0; i < b.N; i++ {
        var x IntSet
        var y IntSet
        x.Add(m)
        y.Add(n)
        x.UnionWith(&y)
    }
}

func BenchmarkMapSetAdd(b *testing.B) {
    rng := rand.New(rand.NewSource(seed))
    n := rng.Intn(10000)
    for i := 0; i < b.N; i++ {
        var x IntSetMap
        x.Add(n)
    }
}

func BenchmarkMapSetUnionWith(b *testing.B) {
    rng := rand.New(rand.NewSource(seed))
    m := rng.Intn(10000)
    n := rng.Intn(10000)
    for i := 0; i < b.N; i++ {
        var x IntSetMap
        var y IntSetMap

        x.Add(m)
        y.Add(n)
        x.UnionWith(&y)
    }
}

// BenchmarkIntSetAdd-4          100000         15789 ns/op        7228 B/op          8 allocs/op
// BenchmarkIntSetUnionWith-4    100000         17076 ns/op        9743 B/op         16 allocs/op
// BenchmarkMapSetAdd-4          100000         14435 ns/op        5520 B/op          3 allocs/op
// BenchmarkMapSetUnionWith-4    100000         15038 ns/op        5712 B/op          7 allocs/op
// ok      _/Users/takano/golang/learning/ch11/ex7 6.896s
