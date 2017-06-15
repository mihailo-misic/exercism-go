package diffsquares

const testVersion = 1

func SquareOfSums(i int) int {
	r := 0

	for x := 1; x <= i; x++ {
		r += x
	}

	return r * r
}

func SumOfSquares(i int) int {
	r := 0

	for x := 1; x <= i; x++ {
		r += x * x
	}

	return r
}

func Difference(i int) int {
	return SquareOfSums(i) - SumOfSquares(i)
}
