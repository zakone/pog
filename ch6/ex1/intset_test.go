package intset

import "testing"

func TestIntsetLen(t *testing.T) {
    var tests = []struct {
        para1    int
        para2    int
        para3    int
        expected int
    }{
        {1, 100, 200, 3},
        {1, 10, 20, 3},
        {1, 10, 100, 3},
    }
    for _, test := range tests {
        var x IntSet
        x.Add(test.para1)
        x.Add(test.para2)
        x.Add(test.para3)
        if got := x.Len(); got != test.expected {
            t.Errorf("%s Len is %d, expected %d", x.String(), got, test.expected)
        }
    }
}

func TestIntsetRemove(t *testing.T) {
    var tests = []struct {
        para1    int
        para2    int
        para3    int
        expected string
    }{
        {1, 100, 200, "{100 200}"},
        {1, 10, 20, "{10 20}"},
        {1, 10, 100, "{10 100}"},
    }
    for _, test := range tests {
        var x IntSet
        x.Add(test.para1)
        x.Add(test.para2)
        x.Add(test.para3)
        x.Remove(test.para1)
        if got := x.String(); got != test.expected {
            t.Errorf("result %s, expected %s", got, test.expected)
        }
    }
}

func TestIntsetClear(t *testing.T) {
    var tests = []struct {
        para1    int
        para2    int
        para3    int
        expected string
    }{
        {1, 100, 200, "{}"},
        {1, 10, 20, "{}"},
        {1, 10, 100, "{}"},
    }
    for _, test := range tests {
        var x IntSet
        x.Add(test.para1)
        x.Add(test.para2)
        x.Add(test.para3)
        x.Clear()
        if got := x.String(); got != test.expected {
            t.Errorf("result %s, expected %s", got, test.expected)
        }
    }
}

func TestIntsetCopy(t *testing.T) {
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
        x.Add(test.para1)
        x.Add(test.para2)
        x.Add(test.para3)
        if got := x.Copy().String(); got != test.expected {
            t.Errorf("result %s, expected %s", got, test.expected)
        }
    }
}
