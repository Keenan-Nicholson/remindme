package bot

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/Keenan-Nicholson/remindme/pkg/database"
	"github.com/Keenan-Nicholson/remindme/pkg/utils"
)

type timerArgs struct {
	Duration int            `description:"Duration (int)"`
	Unit     string         `description:"Unit of time (days / hours / minutes / seconds)"`
	User     discordgo.User `description:"User"`
	Reminder string         `description:"Reminder message"`
}

func TimerCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate, args timerArgs) {
	duration := args.Duration
	unit := args.Unit
	userID := args.User.ID
	reminder := args.Reminder
	channelID := i.ChannelID
	guildID := i.GuildID

	log.Printf("unit: %s, duration: %d, userID: %s, reminder: %s, channel: %s, guild: %s\n", unit, duration, userID, reminder, channelID, guildID)

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
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Invalid unit of time. Please use days, hours, minutes, or seconds.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	validatedDuration, inputErr := utils.ValidateDuration(timeDuration)
	if inputErr != nil {
		log.Println("Error validating duration:", inputErr)
		return
	}

	// Insert reminder into the database and handle errors
	id, err := database.InsertReminder(userID, validatedDuration, reminder, channelID, guildID)
	if err != nil {
		log.Println("Error inserting reminder:", err)
	} else {
		log.Printf("Reminder created with ID: %d", id)
	}

	// Handle the cron job
	go CreateOneTimeCronJob(s, validatedDuration, userID, reminder, id, channelID)

	// Respond to the interaction
	responseErr := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Reminder set!",
			Flags:   1 << 6, // This sets the response to be ephemeral
		},
	})

	if responseErr != nil {
		log.Println("Error sending interaction response:", responseErr)
	}
}

type dateArgs struct {
	Year     int            `description:"Year"`
	Month    int            `description:"Month"`
	Day      int            `description:"Day"`
	Hour     int            `description:"Hour (24h)"`
	Minute   int            `description:"Minute"`
	User     discordgo.User `description:"User"`
	Reminder string         `description:"Reminder message"`
}

func DateCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate, args dateArgs) {
	year := args.Year
	month := args.Month
	day := args.Day
	hour := args.Hour
	minute := args.Minute
	userID := args.User.ID
	reminder := args.Reminder
	channelID := i.ChannelID
	guildID := i.GuildID

	log.Printf("year: %d, month: %d, day: %d, hour: %d, minute: %d, userID: %s, reminder: %s, channel: %s, guild: %s\n", year, month, day, hour, minute, userID, reminder, channelID, guildID)

	// Handle the cron job
	timeDuration := utils.ConvertDateToDuration(year, month, day, hour, minute)

	validatedDuration, inputErr := utils.ValidateDuration(timeDuration)
	if inputErr != nil {
		log.Println("Error validating duration:", inputErr)
		return
	}
	// Insert reminder into the database and handle errors
	id, err := database.InsertReminder(userID, validatedDuration, reminder, channelID, guildID)
	if err != nil {
		log.Println("Error inserting reminder:", err)
	} else {
		log.Printf("Reminder created with ID: %d", id)
	}

	go CreateOneTimeCronJob(s, validatedDuration, userID, reminder, id, channelID)

	// Respond to the interaction
	responseErr := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Reminder set!",
			Flags:   1 << 6, // This sets the response to be ephemeral
		},
	})

	if responseErr != nil {
		log.Println("Error sending interaction response:", err)
	}

}
