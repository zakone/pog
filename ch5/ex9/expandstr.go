package expandstr

import (
    "fmt"
    "strings"
)

func Expend(s, part string, f func(string) string) string {
    return strings.Replace(s, part, f(part), -1)
}

func F_replace(s string) string {
    return fmt.Sprintf("%s_replace", s)
}
