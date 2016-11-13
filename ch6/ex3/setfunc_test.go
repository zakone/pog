package intset

import "testing"

func TestIntersectWith(t *testing.T) {
    var tests = []struct {
        para1    int
        para2    int
        para3    int
        para4    int
        expected string
    }{
        {1, 10, 10, 100, "{10}"},
        {1, 10, 20, 100, "{}"},
        {1, 10, 1, 10, "{1 10}"},
    }
    for _, test := range tests {
        var x IntSet
        var y IntSet
        x.Add(test.para1)
        x.Add(test.para2)
        y.Add(test.para3)
        y.Add(test.para4)
        x.IntersectWith(&y)
        if got := x.String(); got != test.expected {
            t.Errorf("result %s, expected %s", got, test.expected)
        }
    }
}

func TestDifferenceWith(t *testing.T) {
    var tests = []struct {
        para1    int
        para2    int
        para3    int
        para4    int
        expected string
    }{
        {1, 10, 10, 100, "{1}"},
        {1, 10, 20, 100, "{1 10}"},
        {1, 10, 1, 10, "{}"},
    }
    for _, test := range tests {
        var x IntSet
        var y IntSet
        x.Add(test.para1)
        x.Add(test.para2)
        y.Add(test.para3)
        y.Add(test.para4)
        x.DifferenceWith(&y)
        if got := x.String(); got != test.expected {
            t.Errorf("result %s, expected %s", got, test.expected)
        }
    }
}

func TestSymmetricDifference(t *testing.T) {
    var tests = []struct {
        para1    int
        para2    int
        para3    int
        para4    int
        para5    int
        expected string
    }{
        {1, 10, 10, 200, 300, "{1 200 300}"},
        {1, 10, 1, 10, 100, "{100}"},
        {1, 10, 1, 1, 10, "{}"},
    }
    for _, test := range tests {
        var x IntSet
        var y IntSet
        x.Add(test.para1)
        x.Add(test.para2)
        y.Add(test.para3)
        y.Add(test.para4)
        y.Add(test.para5)
        x.SymmetricDifference(&y)
        if got := x.String(); got != test.expected {
            t.Errorf("result %s, expected %s", got, test.expected)
        }
    }
}
