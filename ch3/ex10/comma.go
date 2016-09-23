package comma

import "bytes"

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	var buf bytes.Buffer
	// for i := 0; i < len(s); i++ {
	// 	if len(s)%3 == 0 {
	// 		break
	// 	}
	// 	if (len(s)-i)%3 == 0 {
	// 		buf.WriteString(s[:i])
	// 		buf.WriteString(",")
	// 		s = s[i:]
	// 		break
	// 	}
	// }
	i := len(s) % 3
	buf.WriteString(s[:i] + ",")
	s = s[i:]
	for j := 0; j < len(s); j++ {
		buf.WriteByte(s[j])
		if (j+1)%3 == 0 && (j+1) != len(s) {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
