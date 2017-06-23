package summultiples

const testVersion = 2

func SumMultiples(lim int, ints ...int) int {
	var r int

	for i := 0; i < lim; i++ {
		for _, v := range ints {
			if i%v == 0 {
				r += i
				break
			}
		}
	}

	return r
}
