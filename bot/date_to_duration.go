package bot

import (
	"fmt"
	"time"
)

func ConvertDateToDuration(year int, month int, day int, hour int, minute int) time.Duration {
	// Get the current time
	now := time.Now()
	fmt.Println(now)
	now = now.UTC()
	fmt.Println(now)

	// Create a time object for the target date
	targetTime := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
	fmt.Println(targetTime)

	// Calculate the duration between the current time and the target time in 24 hour format
	duration := targetTime.Sub(now)
	fmt.Println(duration)

	return duration
}
