package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/Keenan-Nicholson/remindme/pkg/database"
)

func PopulateCronScheduleFromDatabase(s *discordgo.Session) error {
	currentTime := time.Now().UTC()

	// Query the database for all reminders
	rows, err := database.GetReminders()
	if err != nil {
		return fmt.Errorf("failed to query database: %w", err)
	}
	defer rows.Close()

	// Store all reminders in memory
	var reminders []struct {
		ID              int
		CreatedAt       time.Time
		Username        string
		DurationSeconds int
		Reminder        string
		ChannelID       string
		GuildID         string
	}

	for rows.Next() {
		var reminder struct {
			ID              int
			CreatedAt       time.Time
			Username        string
			DurationSeconds int
			Reminder        string
			ChannelID       string
			GuildID         string
		}
		if err := rows.Scan(&reminder.ID, &reminder.CreatedAt, &reminder.Username, &reminder.DurationSeconds, &reminder.Reminder, &reminder.ChannelID, &reminder.GuildID); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		reminders = append(reminders, reminder)
	}

	for _, r := range reminders {
		reminderTime := r.CreatedAt.Add(time.Duration(r.DurationSeconds) * time.Second)

		if currentTime.Sub(reminderTime) > 0 {
			// This reminder is in the past, delete it
			if dbErr := database.DeleteReminder(r.ID); dbErr != nil {
				log.Println("Error deleting reminder from DB:", dbErr)
			} else {
				log.Println("Reminder deleted from DB!")
			}
		} else if currentTime.Equal(reminderTime) {
			// Reminder is set to the current time, send it immediately
			message := fmt.Sprintf("Hey <@%s>, this is your reminder: %s!", r.Username, r.Reminder)
			log.Printf("Executing reminder for user %s: %s in channel %s", r.Username, r.Reminder, r.ChannelID)
			_, err := s.ChannelMessageSend(r.ChannelID, message)
			if err != nil {
				log.Println("Error sending message:", err)
			} else {
				log.Println("Message sent!")
			}
		} else {
			// Reminder is in the future, schedule it
			durationUntilReminder := reminderTime.Sub(currentTime)
			log.Printf("Scheduling reminder for user %s: %s in %s in channel %s in guild %s", r.Username, r.Reminder, durationUntilReminder, r.ChannelID, r.GuildID)
			CreateOneTimeCronJob(s, durationUntilReminder, r.Username, r.Reminder, r.ID, r.ChannelID)
		}
	}

	return nil
}
