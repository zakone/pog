package maxmin

import "testing"

func TestMaxOne(t *testing.T) {
    var tests = []struct {
        para1    int
        para2    int
        para3    int
        expected int
    }{
        {1, 2, 3, 3},
        {-1, -2, -3, -1},
    }
    for _, test := range tests {
        if got, _ := MaxOne(test.para1, test.para2, test.para3); got != test.expected {
            t.Errorf("Max Number wrong: %d", got)
        }
    }
    _, ok := MaxOne()
    if ok != false {
        t.Errorf("None data should return false")
    }
}

func TestMinOne(t *testing.T) {
    var tests = []struct {
        para1    int
        para2    int
        para3    int
        expected int
    }{
        {1, 2, 3, 1},
        {-1, -2, -3, -3},
    }
    for _, test := range tests {
        if got, _ := MinOne(test.para1, test.para2, test.para3); got != test.expected {
            t.Errorf("Max Number wrong: %d", got)
        }
    }
    _, ok := MinOne()
    if ok != false {
        t.Errorf("None data should return false")
    }
}
