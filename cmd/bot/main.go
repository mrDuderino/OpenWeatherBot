package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/mrDuderino/OpenWeatherBot/internal/app/commands"
	"github.com/mrDuderino/OpenWeatherBot/internal/service/weather"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 60,
	}
	weatherService := weather.NewService()
	commander := commands.NewCommander(bot, weatherService)

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)
		case "weather":
			commander.CurrentWeather(update.Message)
		default:
			commander.Default(update.Message)
		}
	}
}
