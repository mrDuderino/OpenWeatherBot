package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (c *Commander) CurrentWeather(inputMessage *tgbotapi.Message) {
	outputMsgText := "Wheather:\n\n"
	w := c.weatherService.Show()
	outputMsgText += "Temperature: " +
		strconv.FormatFloat(w.Temperature, 'f', 0, 64) + "ºC\n" +
		"Humidity: " + strconv.FormatFloat(w.Humidity, 'f', 0, 64) + "%"

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	c.bot.Send(msg)
}
