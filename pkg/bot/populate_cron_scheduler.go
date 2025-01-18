package bot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Keenan-Nicholson/remindme/pkg/database"
	"github.com/bwmarrin/discordgo"
)

func PopulateCronScheduleFromDatabase(s *discordgo.Session) error {

	currentTime := time.Now().UTC()

	// Query the database for all reminders

	rows, err := database.GetReminders()

	if err != nil {
		return fmt.Errorf("failed to query database: %w", err)
	}
	defer rows.Close()

	// Iterate over the rows and schedule each reminder
	for rows.Next() {
		channel_id := os.Getenv("DISCORD_CHANNEL_ID")

		channelID := channel_id

		var id int
		var created_at time.Time
		var username string
		var durationSeconds int
		var reminder string
		if err := rows.Scan(&id, &username, &durationSeconds, &reminder); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}

		if currentTime.Sub(created_at.Add(time.Duration(durationSeconds)*time.Second)) > 0 {
			// This reminder is in the past, delete it
			dbErr := database.DeleteReminder(id)
			if dbErr != nil {
				log.Println("Error deleting reminder from DB:", dbErr)
			} else {
				log.Println("Reminder deleted from DB!")
			}
		}

		if currentTime.Sub(created_at.Add(time.Duration(durationSeconds)*time.Second)) == 0 {
			// if somehow the reminder is set to the current time, send it immediately
			message := fmt.Sprintf("Hey <@%s>, this is your reminder: %s!", username, reminder)

			log.Printf("Executing reminder for user %s: %s", username, reminder)
			_, err := s.ChannelMessageSend(channelID, message)
			if err != nil {
				log.Println("Error sending message:", err)
			} else {
				log.Println("Message sent!")
			}
		}

		if currentTime.Sub(created_at.Add(time.Duration(durationSeconds)*time.Second)) < 0 {
			// This reminder is in the future, schedule it
			newDuration := created_at.Add(time.Duration(durationSeconds) * time.Second).Sub(currentTime)
			log.Printf("Scheduling reminder for user %s: %s in %s", username, reminder, newDuration)

			CreateOneTimeCronJob(s, newDuration, username, reminder, id)
		}

		time.Sleep(2000)
	}

	return nil
}
