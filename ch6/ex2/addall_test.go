package intset

import "testing"

func TestIntsetAddall(t *testing.T) {
    var tests = []struct {
        para1    int
        para2    int
        para3    int
        expected string
    }{
        {1, 100, 200, "{1 100 200}"},
        {1, 10, 20, "{1 10 20}"},
        {1, 10, 100, "{1 10 100}"},
    }
    for _, test := range tests {
        var x IntSet
        x.AddAll(test.para1, test.para2, test.para3)
        if got := x.String(); got != test.expected {
            t.Errorf("result %s, expected %s", got, test.expected)
        }
    }
    var x IntSet
    x.AddAll()
    if got := x.String(); got != "{}" {
        t.Errorf("result %s, expected %s", got, "{}")
    }
}
