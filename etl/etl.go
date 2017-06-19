package etl

import "strings"

const testVersion = 1

func Transform(m map[int][]string) map[string]int {
	// Calculating the required length of resulting map.
	t := 0
	for _, v := range m {
		t += len(v)
	}
	r := make(map[string]int, t)

	// Filling in the result.
	for k, v := range m {
		for _, v := range v {
			v := strings.ToLower(v)
			r[v] = k
		}
	}

	return r
}
