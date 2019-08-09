// Package timep wraps functions that return time.Time to return *time.Time.
//
// It's handy if you need it.
package timep

import (
	"time"
)

// Date returns the Time corresponding to
//	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
// in the appropriate zone for that time in the given location.
//
// See https://godoc.org/time#Date for more details.
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) *time.Time {
	t := time.Date(year, month, day, hour, min, sec, nsec, loc)
	return &t
}

// Now returns the current local time.
//
// See https://godoc.org/time#Now for more details.
func Now() *time.Time {
	t := time.Now()
	return &t
}

// Parse parses a formatted string and returns the time value it represents.
//
// See https://godoc.org/time#Parse for more details.
func Parse(layout, value string) (*time.Time, error) {
	t, err := time.Parse(layout, value)
	return &t, err
}

// ParseInLocation is like Parse but differs in two important ways.
//
// See https://godoc.org/time#ParseInLocation for more details.
func ParseInLocation(layout, value string, loc *time.Location) (*time.Time, error) {
	t, err := time.ParseInLocation(layout, value, loc)
	return &t, err
}

// Unix returns the local Time corresponding to the given Unix time,
//
// See https://godoc.org/time#Unix for more details.
func Unix(sec int64, nsec int64) *time.Time {
	t := time.Unix(sec, nsec)
	return &t
}
