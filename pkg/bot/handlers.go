package bot

import (
	"log"
	"time"

	"github.com/Keenan-Nicholson/remindme/pkg/database"
	"github.com/Keenan-Nicholson/remindme/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

func TimerCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "settimer" {

		duration := i.ApplicationCommandData().Options[0].IntValue()
		unit := i.ApplicationCommandData().Options[1].StringValue()
		userID := i.ApplicationCommandData().Options[2].UserValue(s).ID
		reminder := i.ApplicationCommandData().Options[3].StringValue()
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
}

func DateCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "setdate" {

		year := int(i.ApplicationCommandData().Options[0].IntValue())
		month := int(i.ApplicationCommandData().Options[1].IntValue())
		day := int(i.ApplicationCommandData().Options[2].IntValue())
		hour := int(i.ApplicationCommandData().Options[3].IntValue())
		minute := int(i.ApplicationCommandData().Options[4].IntValue())
		userID := i.ApplicationCommandData().Options[5].UserValue(s).ID
		reminder := i.ApplicationCommandData().Options[6].StringValue()
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
}
