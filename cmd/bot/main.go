package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/mrDuderino/OpenWeatherBot/internal/service/weather"
	"log"
	"os"
	"strconv"
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
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		case "weather":
			weatherCommand(bot, update.Message, weatherService)
		default:
			defaultBehavior(bot, update.Message)
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Hello!\nWeather Bot will help you to receive weather updates.\nEnter /weather and city name to start!",
	)
	bot.Send(msg)
}

func weatherCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, ws *weather.Service) {
	outputMsgText := "Wheather:\n\n"
	w := ws.Show()
	outputMsgText += "Temperature: " +
		strconv.FormatFloat(w.Temperature, 'f', 0, 64) + "ºC\n" +
		"Humidity: " + strconv.FormatFloat(w.Humidity, 'f', 0, 64) + "%"

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID
	bot.Send(msg)
}
