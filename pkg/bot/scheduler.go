package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/Keenan-Nicholson/remindme/pkg/database"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron/v2"
)

func CreateOneTimeCronJob(discord *discordgo.Session, duration time.Duration, userID string, reminder string, uid int, channelID string) {
	channel_id := channelID
	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}

	rowID := uid

	// Define the job and task asynchronously
	_, err = s.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(duration)),
		),
		gocron.NewTask(
			func() {
				// Send a message to a Discord channel
				channelID := channel_id
				message := fmt.Sprintf("Hey <@%s>, this is your reminder: %s!", userID, reminder)
				_, err := discord.ChannelMessageSend(channelID, message)
				if err != nil {
					log.Println("Error sending message:", err)
				} else {
					log.Println("Message sent!")
				}

				dbErr := database.DeleteReminder(rowID)
				if dbErr != nil {
					log.Println("Error deleting reminder from DB:", dbErr)
				} else {
					log.Println("Reminder deleted from DB!")
				}
			},
		),
	)

	if err != nil {
		// Handle error
		log.Println("Error creating job:", err)
		return
	}

	// start the scheduler
	s.Start()

}
