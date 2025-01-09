package main

import (
	"log"

	"github.com/Keenan-Nicholson/remindme/bot"
	"github.com/joho/godotenv"
)

func main() {

	bot.SetupLogger()

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	discord, err := bot.RunBot()
	if err != nil {
		log.Fatal(err)
	}
	defer discord.Close()

	discord.AddHandler(bot.TimerCommandHandler)
	discord.AddHandler(bot.DateCommandHandler)

	log.Println("Bot is stopping!")
}
