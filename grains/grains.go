package grains

import "errors"

const testVersion = 1

func Square(i int) (uint64, error) {
	if i < 1 || i > 64 {
		return 0, errors.New("Invalid field requested.")
	}

	var r uint64 = 1

	for x := 1; x < i; x++ {
		r = r + r
	}

	return r, nil
}

func Total() uint64 {
	var r uint64 = 1
	var n uint64 = 1

	for i := 0; i < 64; i++ {
		r = r + r
		n += r
	}

	return n
}
