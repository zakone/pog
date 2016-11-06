package main

import "fmt"
import "bytes"

type IntSet struct {
    words []uint64
}

func main() {
    var x IntSet
    x.Add(1)
    x.Add(3)
    x.Add(100)
    x.Add(300)
    fmt.Println(x.String())

    elems := x.Elems()
    fmt.Println(elems)

}

func (s *IntSet) Elems() []int {
    var elems []int
    for i, tword := range s.words {
        for j := 0; j < 64; j++ {
            if tword&(1<<uint(j)) != 0 {
                elems = append(elems, 64*i+j)
            }
        }
    }
    return elems
}
func (s *IntSet) IntersectWith(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] &= tword
        }
    }
}

func (s *IntSet) DifferenceWith(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] &^= tword
        }
    }
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] ^= tword
        } else {
            s.words = append(s.words, tword)
        }
    }
}

func (s *IntSet) Len() int {
    leng := 0
    for _, w := range s.words {
        leng += popCountLastClear(w)
    }
    return leng
}

func (s *IntSet) Remove(x int) {
    if s.Has(x) {
        word, bit := x/64, uint(x%64)
        s.words[word] ^= 1 << bit
    }
}

func (s *IntSet) Clear() {
    s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
    var y IntSet
    y.words = append(y.words, s.words...)
    return &y
}

func popCountLastClear(x uint64) int {
    var tmp byte
    for x != 0 {
        x = x & (x - 1)
        tmp += 1
    }
    return int(tmp)
}

func (s *IntSet) UnionWith(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] |= tword
        } else {
            s.words = append(s.words, tword)
        }
    }
}

func (s *IntSet) Has(x int) bool {
    word, bit := x/64, uint(x%64)
    return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
    word, bit := x/64, uint(x%64)
    for word >= len(s.words) {
        s.words = append(s.words, 0)
    }
    s.words[word] |= 1 << bit
}

func (s *IntSet) String() string {
    var buf bytes.Buffer
    buf.WriteByte('{')
    for i, word := range s.words {
        if word == 0 {
            continue
        }
        for j := 0; j < 64; j++ {
            if word&(1<<uint(j)) != 0 {
                if buf.Len() > len("{") {
                    buf.WriteByte(' ')
                }
                fmt.Fprintf(&buf, "%d", 64*i+j)
            }
        }
    }
    buf.WriteByte('}')
    return buf.String()
}
