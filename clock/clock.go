package clock

import "fmt"

const testVersion = 4

type Clock struct {
    hour   int
    minute int
}

func New(hour, minute int) Clock {
    hour += minute / 60  // Get extra hours from overflowing minutes
    hour = hour % 24     // Get hours
    minute = minute % 60 // Get minutes

    return Clock{hour, minute}
}

func (clock Clock) String() string {
    return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
    return New(clock.hour, clock.minute + minutes)
}

//func main() {
//    fmt.Println(New(0,45).Add(40))
//}