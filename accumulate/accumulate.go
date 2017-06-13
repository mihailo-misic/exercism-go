package accumulate

const testVersion = 1

func Accumulate(vals []string, funk func(string) string) []string {
	var r []string

	for _, v := range vals {
		r = append(r, funk(v))
	}

	return r
}
