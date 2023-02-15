package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mrDuderino/OpenWeatherBot/internal/service/weather"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	weatherService *weather.Service
}

func NewCommander(bot *tgbotapi.BotAPI, weatherService *weather.Service) *Commander {
	return &Commander{
		bot:            bot,
		weatherService: weatherService,
	}
}
