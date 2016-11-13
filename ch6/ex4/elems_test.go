package intset

import "testing"

func TestElems(t *testing.T) {
    var tests = []struct {
        para1    int
        para2    int
        para3    int
        expected []int
    }{
        {1, 2, 10, []int{1, 2, 10}},
    }
    for _, test := range tests {
        var x IntSet
        x.Add(test.para1)
        x.Add(test.para2)
        x.Add(test.para3)
        got := x.Elems()
        for i, val := range got {
            if val != test.expected[i] {
                t.Errorf("result %d, expected %d", val, test.expected[i])
            }
        }
    }
}
