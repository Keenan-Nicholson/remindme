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

	// Define the slash command
	command := &discordgo.ApplicationCommand{
		Name:        "setreminder",
		Description: "Create a reminder.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "duration",
				Description: "<duration> days, hours, minutes, or seconds.",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
			},
			{
				Name:        "unit",
				Description: "days, hours, minutes, or seconds.",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "user",
				Description: "The user to remind.",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    true,
			},
			{
				Name:        "reminder",
				Description: "The reminder message.",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	}

	// Replace with your actual application ID
	appID := os.Getenv("DISCORD_BOT_APP_ID")

	// Register the command
	_, err = discord.ApplicationCommandCreate(appID, "", command)
	if err != nil {
		log.Fatalf("Error creating slash command: %v", err)
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

	// Define the job and task asynchronously
	_, err = s.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(duration)),
		),
		gocron.NewTask(
			func() {
				// Send a message to a Discord channel
				channelID := channel_id // Replace with the desired channel ID
				message := fmt.Sprintf("Hey <@%s>, this is your reminder to %s!", userID, reminder)
				fmt.Println("Sending message:", message)
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

func commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "setreminder" {

		duration := i.ApplicationCommandData().Options[0].IntValue()
		unit := i.ApplicationCommandData().Options[1].StringValue()
		userID := i.ApplicationCommandData().Options[2].UserValue(s).ID
		reminder := i.ApplicationCommandData().Options[3].StringValue()

		fmt.Println(unit, duration, userID, reminder)

		var timeDuration time.Duration
		switch unit {
		case "days":
			timeDuration = time.Duration(duration) * 24 * time.Hour
		case "hours":
			timeDuration = time.Duration(duration) * time.Hour
		case "minutes":
			timeDuration = time.Duration(duration) * time.Minute
		case "seconds":
			timeDuration = time.Duration(duration) * time.Second
		default:
			return
		}

		// Handle the cron job
		go handleCronJob(s, timeDuration, userID, reminder)

		// Respond to the interaction
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Reminder set!",
			},
		})
		if err != nil {
			fmt.Println("Error sending interaction response:", err)
		}
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

	discord.AddHandler(commandHandler)

	// Wait for a signal to stop the bot gracefully
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	fmt.Println("Bot is stopping!")
}
