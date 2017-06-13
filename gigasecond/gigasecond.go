package gigasecond

import (
	"math"
	"time"
)

const testVersion = 4

// Avg = 1.278s
func AddGigasecond(date time.Time) time.Time {
	return date.Add(time.Second * time.Duration(math.Pow(10, 9)))
}

/*
AvgTime = 1.943s
func AddGigasecond(date time.Time) time.Time {
    return date.Add(time.Second * time.Duration(math.Pow10(9)))
}

AvgTime = 2.819s
func AddGigasecond(date time.Time) time.Time {
    return date.Add(time.Second * 1e+09)
}

 AvgTime  = 2.822s
func AddGigasecond(date time.Time) time.Time {
    return date.Add(time.Second * 1000000000)
}

 AvgTime = 2.823s
func AddGigasecond(date time.Time) time.Time {
    return date.Add(time.Second * time.Duration(1000000000))
}

Note: These are the avg times from only 3 bench runs (not super precise)
*/
