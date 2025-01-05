package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
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

func handleCronJob(discord *discordgo.Session, duration time.Duration, userID string, reminder string) {
	channel_id := os.Getenv("DISCORD_CHANNEL_ID")

	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}

	// Define the job and task
	_, err = s.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(duration)),
		),
		gocron.NewTask(
			func() {
				// Send a message to a Discord channel
				channelID := channel_id // Replace with the desired channel ID
				message := fmt.Sprintf("Hey <@%s>, this is your reminder to%s!", userID, reminder)
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
	// fmt.Println("Job ID:", s.Jobs()[0].ID())

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

func parseDuration(message string) (time.Duration, error, string) {

	message_contents := strings.Split(message, "remindme!")
	message = message_contents[1]

	command := strings.Split(message, ":")[0]
	reminder := strings.Split(message, ":")[1]

	parts := strings.Split(command, " ")

	if len(parts) < 2 {
		return 0, fmt.Errorf("invalid message format"), ""
	}

	num, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid number format"), ""
	}

	unit := parts[1]
	var duration time.Duration
	switch unit {
	case "days":
		duration = time.Duration(num) * 24 * time.Hour
	case "hours":
		duration = time.Duration(num) * time.Hour
	case "minutes":
		duration = time.Duration(num) * time.Minute
	case "seconds":
		duration = time.Duration(num) * time.Second
	default:
		return 0, fmt.Errorf("invalid time unit"), ""
	}

	return duration, nil, reminder
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "remindme!") {
		duration, err, reminder := parseDuration(m.Content)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %v", err))
			return
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Reminder set for %v", duration))

		handleCronJob(s, duration, m.Author.ID, reminder)

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

	discord.AddHandler(messageCreate)

	// Wait for a signal to stop the bot gracefully
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	fmt.Println("Bot is stopping!")
}
