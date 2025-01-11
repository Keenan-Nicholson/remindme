package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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

	// stop the bot gracefully
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Println("Bot is stopping!")
}
