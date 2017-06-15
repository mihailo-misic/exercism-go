package gigasecond

import (
	"time"
)

const testVersion = 4

//AvgTime = 1.33 ns/op
func AddGigasecond(date time.Time) time.Time {
	return date.Add(time.Second * 1e+09)
}

//AvgTime  = 1.33 ns/op
//func AddGigasecond(date time.Time) time.Time {
//    return date.Add(time.Second * 1000000000)
//}

//AvgTime = 1.33 ns/op
//func AddGigasecond(date time.Time) time.Time {
//    return date.Add(time.Second * time.Duration(1000000000))
//}

//AvgTime = 6.3 ns/op
//func AddGigasecond(date time.Time) time.Time {
//    return date.Add(time.Second * time.Duration(math.Pow10(9)))
//}

// Avg = 38.4 ns/op
//func AddGigasecond(date time.Time) time.Time {
//	return date.Add(time.Second * time.Duration(math.Pow(10, 9)))
//}
