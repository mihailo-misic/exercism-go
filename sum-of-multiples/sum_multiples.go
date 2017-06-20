package summultiples

const testVersion = 2

// TODO Use a map instead of slice to store the used ints
func SumMultiples(lim int, ints ...int) int {
	r := 0

	for i := 0; i < lim; i++ {
		used := []int{}

		for _, v := range ints {
			if i%v == 0 && !intInSlice(i, used) {
				r += i
				used = append(used, i)
			}
		}
	}

	return r
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
