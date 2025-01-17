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

	// Define the commands
	setTimerCommand := &discordgo.ApplicationCommand{
		Name:        "settimer",
		Description: "Create a timer-based reminder.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "duration",
				Description: "Duration (int)",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
			},
			{
				Name:        "unit",
				Description: "Unit of time",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "days",
						Value: "days",
					},
					{
						Name:  "hours",
						Value: "hours",
					},
					{
						Name:  "minutes",
						Value: "minutes",
					},
					{
						Name:  "seconds",
						Value: "seconds",
					},
				},
			},

			{
				Name:        "user",
				Description: "User",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    true,
			},

			{
				Name:        "reminder",
				Description: "Reminder Message",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	}

	setDateCommand := &discordgo.ApplicationCommand{
		Name:        "setdate",
		Description: "Create a reminder for a specific date and time (UTC).",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "year",
				Description: "Year",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
			},
			{
				Name:        "month",
				Description: "Month",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
			},
			{
				Name:        "day",
				Description: "Day",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
			},
			{
				Name:        "hour",
				Description: "Hour (24h)",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
			},
			{
				Name:        "minute",
				Description: "Minute",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
			},
			{
				Name:        "user",
				Description: "User",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    true,
			},
			{
				Name:        "reminder",
				Description: "Reminder Message",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	}

	appID := os.Getenv("DISCORD_APP_ID")
	guildID := os.Getenv("DISCORD_GUILD_ID")

	_, err = discord.ApplicationCommandCreate(appID, guildID, setTimerCommand)
	if err != nil {
		log.Fatalf("Error creating 'settimer' command: %v", err)
	}
	_, err = discord.ApplicationCommandCreate(appID, guildID, setDateCommand)
	if err != nil {
		log.Fatalf("Error creating 'setdate' command: %v", err)
	}

	fmt.Println("Bot is running!")
	return discord, nil
}
