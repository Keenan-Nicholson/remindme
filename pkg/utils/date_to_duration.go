package utils

import (
	"time"
)

func ConvertDateToDuration(year int, month int, day int, hour int, minute int, location *time.Location) time.Duration {
	// Get the current time
	now := time.Now()
	now = now.UTC()

	// Create a time object for the target date
	targetTime := time.Date(year, time.Month(month), day, hour, minute, 0, 0, location)

	// Calculate the duration between the current time and the target time in 24 hour format
	duration := targetTime.Sub(now)

	return duration
}
