package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"github.com/Keenan-Nicholson/remindme/pkg/bot"
	"github.com/Keenan-Nicholson/remindme/pkg/database"
	"github.com/Keenan-Nicholson/remindme/pkg/utils"
)

func main() {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir("data", 0755)
	}

	utils.SetupLogger()

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

	database.InitDB()

	bot.PopulateCronScheduleFromDatabase(discord)

	// stop the bot gracefully
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Println("Bot is stopping!")
}
