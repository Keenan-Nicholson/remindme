package bot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron/v2"
)

func CreateOneTimeCronJob(discord *discordgo.Session, duration time.Duration, userID string, reminder string) {
	channel_id := os.Getenv("DISCORD_CHANNEL_ID")
	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}

	// Define the job and task asynchronously
	_, err = s.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(duration)),
		),
		gocron.NewTask(
			func() {
				// Send a message to a Discord channel
				channelID := channel_id // Replace with the desired channel ID
				message := fmt.Sprintf("Hey <@%s>, this is your reminder: %s!", userID, reminder)
				_, err := discord.ChannelMessageSend(channelID, message)
				if err != nil {
					log.Println("Error sending message:", err)
				} else {
					log.Println("Message sent!")
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
