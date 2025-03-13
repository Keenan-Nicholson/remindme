package bot

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"pkg.nit.so/switchboard"
)

func RunBot() (*discordgo.Session, error) {
	secret := os.Getenv("DISCORD_BOT_TOKEN")
	guildId := os.Getenv("DISCORD_GUILD_ID")

	commandHandler := switchboard.Switchboard{}
	_ = commandHandler.AddCommand(&switchboard.Command{
		Name:        "settimer",
		Description: "Create a timer-based reminder.",
		Handler:     TimerCommandHandler,
		GuildID:     guildId,
	})
	_ = commandHandler.AddCommand(&switchboard.Command{
		Name:        "setdate",
		Description: "Create a date-based reminder.",
		Handler:     DateCommandHandler,
		GuildID:     guildId,
	})

	discord, err := discordgo.New("Bot " + secret)
	if err != nil {
		return nil, fmt.Errorf("error creating Discord session: %w", err)
	}

	discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	discord.AddHandler(commandHandler.HandleInteractionCreate)

	err = discord.Open()
	if err != nil {
		return nil, fmt.Errorf("error opening connection to Discord: %w", err)
	}

	appID := os.Getenv("DISCORD_APP_ID")

	err = commandHandler.SyncCommands(discord, appID)
	if err != nil {
		return nil, fmt.Errorf("error syncing commands: %w", err)
	}

	fmt.Println("Bot is running!")
	return discord, nil
}
