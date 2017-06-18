package series

const testVersion = 2

func All(i int, s string) []string {
	l := len(s)
	t := l - i + 1
	r := make([]string, t)

	for x := 0; x < t; x++ {
		r[x] = s[x : x+i]
	}

	return r
}

func UnsafeFirst(i int, s string) string {
	return s[:i]
}

func First(i int, s string) (res string, ok bool) {
	if len(s) >= i {
		return s[:i], true
	}

	return "", false
}
