package pascal

const testVersion = 1

func Triangle(i int) [][]int {
	var r = make([][]int, i)

	for x := 0; x < i; x++ {
		row := []int{}

		for y := 0; y <= x; y++ {
			if y == 0 || x == y {
				row = append(row, 1)
			} else {
				row = append(row, r[x-1][y]+r[x-1][y-1])
			}
		}

		r[x] = row
	}

	return r
}
