package queenattack

import (
	"github.com/pkg/errors"
)

const testVersion = 2

func CanQueenAttack(x string, y string) (att bool, err error) {
	// Check if empty.
	if x == "" || y == "" {
		return false, errors.New("Empty string")
	}
	// Check if same.
	if x == y {
		return false, errors.New("Same spot")
	}
	// Check if off-board.
	if !(checkOnBoard(x) && checkOnBoard(y)) {
		return false, errors.New("Off board")
	}
	// Check vertically & horizontally.
	if x[0] == y[0] || x[1] == y[1] {
		return true, nil
	}
	// Check diagonals.
	return checkDiagonals(x, y)
}

func checkOnBoard(pos string) bool {
	valid := []byte{
		49, 50, 51, 52, 53, 54, 55, 56, // 1-8
		97, 98, 99, 100, 101, 102, 103, 104, // a-h
	}
	a := false
	b := false

	for _, v := range valid {
		if pos[0] == v {
			a = true
		}
		if pos[1] == v {
			b = true
		}
	}

	return a && b
}

func checkDiagonals(x, y string) (att bool, err error) {
	yPos := int(y[0])*100 + int(y[1])

	for n := 1; n < 8; n++ {
		xPos := (int(x[0])+n)*100 + int(x[1]) + n
		if xPos == yPos {
			return true, nil
		}

		xPos = (int(x[0])-n)*100 + int(x[1]) - n
		if xPos == yPos {
			return true, nil
		}

		xPos = (int(x[0])+n)*100 + int(x[1]) - n
		if xPos == yPos {
			return true, nil
		}

		xPos = (int(x[0])-n)*100 + int(x[1]) + n
		if xPos == yPos {
			return true, nil
		}
	}

	return false, nil
}
