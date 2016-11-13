package intset

import "testing"

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var x IntSet
        x.Add(1)
    }
}

func BenchmarkUnionWith(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var x IntSet
        var y IntSet
        x.Add(1)
        y.Add(2)
        x.UnionWith(&y)
    }
}
