package stringjoin

import "testing"

func TestJoin(t *testing.T) {
    var tests = []struct {
        para1    string
        para2    string
        para3    string
        para4    string
        expected string
    }{
        {",", "a", "b", "c", "a,b,c"},
    }
    for _, test := range tests {
        if got := Join(test.para1, test.para2, test.para3, test.para4); got != test.expected {
            t.Errorf("Join error: %d", got)
        }
    }
    if Join(",") != "" {
        t.Errorf("should return nil string")
    }
}
