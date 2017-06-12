package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	// Round up hours if negative
	if hour < 0 {
		hour = 24 + (hour % 24)
	}

	hour += minute / 60 // Get extra hours from overflowing minutes

	// Remove an extra hour if negative minutes are not exactly divisible by 60
	if minute < 0 && minute%60 != 0 {
		hour -= 1
	}
	// Round up hours if negative
	if hour < 0 {
		hour = 24 + (hour % 24)
	}

	hour = hour % 24     // Mold hours into 24 format with modulus
	minute = minute % 60 // Mold minutes into 60 format with modulus

	// Round up minutes if negative
	if minute < 0 {
		minute = 60 + minute
	}

	return Clock{hour, minute}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
	return New(clock.hour, clock.minute+minutes)
}
