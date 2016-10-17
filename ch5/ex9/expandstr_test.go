package expandstr

import "testing"

func TestExpend(t *testing.T) {
    var tests = []struct {
        s        string
        part     string
        expected string
    }{
        {"$word $word $word", "$word", "$word_replace $word_replace $word_replace"},
    }
    for _, test := range tests {
        if got := Expend(test.s, test.part, F_replace); got != test.expected {
            t.Errorf("%s IsWrongComma", got)
        }
    }
}
