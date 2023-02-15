package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Hello!\nWeather Bot will help you to receive weather updates.\nEnter /weather and city name to start!",
	)
	c.bot.Send(msg)
}
