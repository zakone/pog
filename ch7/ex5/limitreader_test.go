package limitreader

import "testing"
import "strings"
import "fmt"

func TestWriteCounter(t *testing.T) {
    var tests = []struct {
        input string
        limit int64
        want  string
    }{
        {"hello world", 5, "hello"},
        {"", 0, ""},
    }
    for _, test := range tests {
        buf := make([]byte, len(test.input))
        reader := LimitReader(strings.NewReader(test.input), test.limit)
        n, _ := reader.Read(buf)
        got := fmt.Sprintf("%s", buf[:n])
        if got != test.want {
            t.Errorf("%s, should be %s", got, test.want)
        }
    }
}
