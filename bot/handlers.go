package bot

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func TimerCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "settimer" {

		duration := i.ApplicationCommandData().Options[0].IntValue()
		unit := i.ApplicationCommandData().Options[1].StringValue()
		userID := i.ApplicationCommandData().Options[2].UserValue(s).ID
		reminder := i.ApplicationCommandData().Options[3].StringValue()

		log.Printf("unit: %s, duration: %d, userID: %s, reminder: %s\n", unit, duration, userID, reminder)

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
		go CreateOneTimeCronJob(s, timeDuration, userID, reminder)

		// Respond to the interaction
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Reminder set!",
			},
		})
		if err != nil {
			log.Println("Error sending interaction response:", err)
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

		log.Printf("year: %d, month: %d, day: %d, hour: %d, minute: %d, userID: %s, reminder: %s\n", year, month, day, hour, minute, userID, reminder)

		// Handle the cron job
		duration := ConvertDateToDuration(year, month, day, hour, minute)

		go CreateOneTimeCronJob(s, duration, userID, reminder)

		// Respond to the interaction
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Reminder set!",
			},
		})
		if err != nil {
			log.Println("Error sending interaction response:", err)
		}
	}
}
