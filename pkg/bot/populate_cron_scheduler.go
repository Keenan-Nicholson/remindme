package bot

import (
	"fmt"
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

	fmt.Println(rows)

	// Iterate over the rows and schedule each reminder
	for rows.Next() {
		fmt.Println("making it here")
		// channel_id := os.Getenv("DISCORD_CHANNEL_ID")

		// channelID := channel_id

		var id int
		var created_at time.Time
		var username string
		var durationSeconds int
		var reminder string

		// it is not making it past this point, it is having trouble scanning the rows?
		scanErr := rows.Scan(&id, &created_at, &username, &durationSeconds, &reminder)
		if scanErr != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}

		fmt.Println("not making it here")

		fmt.Println("current time:", currentTime)
		fmt.Println("id:", id, "created_at:", created_at, "username:", username, "durationSeconds:", durationSeconds, "reminder:", reminder)
		// if currentTime.Sub(created_at.Add(time.Duration(durationSeconds)*time.Second)) > 0 {
		// 	// This reminder is in the past, delete it
		// 	dbErr := database.DeleteReminder(id)
		// 	if dbErr != nil {
		// 		log.Println("Error deleting reminder from DB:", dbErr)
		// 	} else {
		// 		log.Println("Reminder deleted from DB!")
		// 	}
		// }

		// if currentTime.Sub(created_at.Add(time.Duration(durationSeconds)*time.Second)) == 0 {
		// 	// if somehow the reminder is set to the current time, send it immediately
		// 	message := fmt.Sprintf("Hey <@%s>, this is your reminder: %s!", username, reminder)

		// 	log.Printf("Executing reminder for user %s: %s", username, reminder)
		// 	_, err := s.ChannelMessageSend(channelID, message)
		// 	if err != nil {
		// 		log.Println("Error sending message:", err)
		// 	} else {
		// 		log.Println("Message sent!")
		// 	}
		// }

		// if currentTime.Sub(created_at.Add(time.Duration(durationSeconds)*time.Second)) < 0 {
		// 	// This reminder is in the future, schedule it
		// 	newDuration := created_at.Add(time.Duration(durationSeconds) * time.Second).Sub(currentTime)
		// 	log.Printf("Scheduling reminder for user %s: %s in %s", username, reminder, newDuration)

		// 	CreateOneTimeCronJob(s, newDuration, username, reminder, id)
		// }
	}

	return nil
}
