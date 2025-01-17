package bot

import (
	"time"
)

func ConvertDateToDuration(year int, month int, day int, hour int, minute int) time.Duration {
	// Get the current time
	now := time.Now()
	now = now.UTC()

	// Create a time object for the target date
	targetTime := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)

	// Calculate the duration between the current time and the target time in 24 hour format
	duration := targetTime.Sub(now)

	return duration
}
