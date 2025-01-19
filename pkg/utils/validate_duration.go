package utils

import (
	"errors"
	"time"
)

// Define min and max durations
const (
	MinDuration = 1 * time.Second
	MaxDuration = time.Hour*24*365*4 + time.Hour*24*366 // ~5 years with leap year
)

// ValidateDuration returns a valid duration or an error if it's out of bounds
func ValidateDuration(duration time.Duration) (time.Duration, error) {
	if duration <= MinDuration {
		return 0, errors.New("duration is too short, minimum is 1 second")
	}
	if duration > MaxDuration {
		return 0, errors.New("duration exceeds maximum limit of ~5 years")
	}
	return duration, nil
}
