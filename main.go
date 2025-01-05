package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
)

func runBot() (*discordgo.Session, error) {
	secret := os.Getenv("DISCORD_BOT_TOKEN")

	discord, err := discordgo.New("Bot " + secret)
	if err != nil {
		log.Fatal(err)
	}

	discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bot is running!")

	return discord, nil
}

func handleCronJob(discord *discordgo.Session) {
	channel_id := os.Getenv("DISCORD_CHANNEL_ID")

	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}

	// Define the job and task
	_, err = s.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			func() {
				// Send a message to a Discord channel
				channelID := channel_id // Replace with the desired channel ID
				message := "Hello, World!"
				_, err := discord.ChannelMessageSend(channelID, message)
				if err != nil {
					fmt.Println("Error sending message:", err)
				} else {
					fmt.Println("Message sent!")
				}
			},
		),
	)
	if err != nil {
		// Handle error
		fmt.Println("Error creating job:", err)
		return
	}

	// each job has a unique id
	fmt.Println("Job ID:", s.Jobs()[0].ID())

	// start the scheduler
	s.Start()

	// block until you are ready to shut down
	select {
	case <-time.After(time.Minute): // Runs for 1 minute
	}

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		// handle error
		fmt.Println("Error shutting down scheduler:", err)
	}
}

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	discord, err := runBot()
	if err != nil {
		log.Fatal(err)
	}
	defer discord.Close()

	// Start the cron job to send messages to the server channel
	handleCronJob(discord)

	// Wait for a signal to stop the bot gracefully
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	fmt.Println("Bot is stopping!")
}
