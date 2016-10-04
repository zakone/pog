package commaDecimal

import "strings"

func comma(s string) string {
    n := len(s)
    if n <= 3 {
        return s
    }
    return comma(s[:n-3]) + "," + s[n-3:]
}

func commaDecimal(s string) string {
    if len(s) < 3 {
        return s
    }
    sign := ""
    if s[0] == '-' {
        sign = "-"
        s = s[1:]
    }
    if dot := strings.LastIndex(s, "."); dot >= 0 {
        sInt := s[:dot]
        sDec := s[dot+1:]
        return sign + comma(sInt) + "." + comma(sDec)
    } else {
        return sign + comma(s)
    }

}
