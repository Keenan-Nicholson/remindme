package bot

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func RunBot() (*discordgo.Session, error) {
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
	appID := os.Getenv("DISCORD_APP_ID")

	// Register the command
	_, err = discord.ApplicationCommandCreate(appID, "", command)
	if err != nil {
		log.Fatalf("Error creating slash command: %v", err)
	}

	fmt.Println("Bot is running!")

	return discord, nil
}
