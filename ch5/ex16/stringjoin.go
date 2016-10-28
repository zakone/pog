package stringjoin

func Join(sep string, vals ...string) string {
	res := ""
	if len(vals) == 0 {
		return res
	}
	for _, val := range vals {
		res = res + val + sep
	}
	return res[:len(res)-1]
}
