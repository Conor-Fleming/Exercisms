// Package gigasecond
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond takes a value of type time.Time, t and returns the time.Time plus on Gigasecond.
func AddGigasecond(t time.Time) time.Time {
	//gigasecond is onethousand millions seconds
	gigaSecond := time.Second * 1000000000

	//add our gigasecond to our given time
	result := t.Add(gigaSecond)

	return result
}
