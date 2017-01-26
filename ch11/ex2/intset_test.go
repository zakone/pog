package intset

import "testing"
import "bytes"
import "fmt"

type IntSetMap struct {
    words map[uint64]bool
}

func (s *IntSetMap) Has(x int) bool {

    return x > 0 && s.words[uint64(x)]
}

func (s *IntSetMap) Add(x int) {
    if x < 0 {
        return
    }
    if s.words == nil {
        s.words = make(map[uint64]bool)
    }
    s.words[uint64(x)] = true
}

func (s *IntSetMap) UnionWith(t *IntSetMap) {
    for k, v := range t.words {
        s.words[k] = v
    }
}

func (s *IntSetMap) String() string {
    var buf bytes.Buffer
    buf.WriteByte('{')
    for k, v := range s.words {
        if v != false {
            if buf.Len() > len("{") {
                buf.WriteByte(' ')
            }
            fmt.Fprintf(&buf, "%d", k)
        }

    }
    buf.WriteByte('}')
    return buf.String()
}

func TestHas(t *testing.T) {
    var tests = []struct {
        para     int
        expected bool
    }{
        {1, true},
        {10, true},
        {300, true},
    }
    for _, test := range tests {
        var x IntSetMap
        x.words = make(map[uint64]bool)
        x.words[uint64(test.para)] = true
        if got := x.Has(test.para); got != test.expected {
            t.Errorf("IntSetMap has %d, result false, expected true", test.para)
        }
    }
}

func TestIntsetAdd(t *testing.T) {
    var tests = []struct {
        para     int
        expected bool
    }{
        {100, true},
        {-1, false},
        {10, true},
    }
    for _, test := range tests {
        var x IntSetMap
        x.Add(test.para)
        if got := x.Has(test.para); got != test.expected {
            t.Errorf("IntSetMap added %d, expected has it, but result %s", test.para, x.String())
        }
    }
}

func TestIntsetUnionWith(t *testing.T) {
    var tests = []struct {
        para1    int
        para2    int
        para3    int
        expected bool
    }{
        {1, 100, 200, true},
        {1, 10, 20, true},
        {1, 10, 100, true},
    }
    for _, test := range tests {
        var x IntSetMap
        var y IntSetMap
        x.Add(test.para1)
        x.Add(test.para2)
        y.Add(test.para3)
        x.UnionWith(&y)
        if got := x.Has(test.para3); got != test.expected {
            t.Errorf("union result %s, expected has %s result no", x.String(), test.para3)
        }
    }
}

// type IntSetMap struct {
//     words map[int][64]int
// }

// func (s *IntSetMap) Has(x int) bool {
//     word, bit := x/64, uint(x%64)
//     return word < len(s.words) && s.words[word][bit] == 1

// }

// func (s *IntSetMap) Add(x int) {
//     word, bit := x/64, uint(x%64)
//     s.words[word] = [64]int{0}
//     s.words[word][bit] = 1

// }

// func (s *IntSetMap) UnionWith(t *IntSetMap) {
//     for i, tword := range t.words {
//         if i < len(s.words) {
//             unionArrays(s.words[i], tword)
//         } else {
//             s.words[i] = tword
//         }
//     }
// }

// func (s *IntSetMap) String() string {
//     var buf bytes.Buffer
//     buf.WriteByte('{')
//     for i, tword := range s.words {
//         if len(tword) == 0 {
//             continue
//         }
//         for j := 0; j < 64; j++ {
//             if tword[j] != 0 {
//                 if buf.Len() > len("{") {
//                     buf.WriteByte(' ')
//                 }
//                 fmt.Fprintf(&buf, "%d", 64*i+j)
//             }
//         }
//     }
//     buf.WriteByte('}')
//     return buf.String()
// }

// func unionArrays(x, y [64]int) {
//     for i, v := range x {
//         x[i] = v | y[i]
//     }
// }
