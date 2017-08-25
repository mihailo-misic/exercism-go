package pythagorean

const testVersion = 1

type Triplet [3]int

func Range(min, max int) (r []Triplet) {
	for min <= max {
		if min%3 == 0 {
			m := min / 3

			if 5*m > max {
				break
			}

			r = append(r, Triplet{3 * m, 4 * m, 5 * m})
			min += 3
		} else {
			min++
		}
	}

	return r
}

func Sum(p int) (r []Triplet) {

	return r
}
