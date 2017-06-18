package pascal

const testVersion = 1

func Triangle(i int) [][]int {
	var r = make([][]int, i)

	for x := 0; x < i; x++ {
		r[x] = []int{}

		for y := 0; y <= x; y++ {
			if y == 0 || x == y {
				r[x] = append(r[x], 1)
			} else {
				r[x] = append(r[x], r[x-1][y]+r[x-1][y-1])
			}
		}
	}

	return r
}
