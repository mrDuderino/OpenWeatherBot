package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mrDuderino/OpenWeatherBot/internal/service/weather"
	"strconv"
	"strings"
)

func (c *Commander) CurrentWeather(inputMessage *tgbotapi.Message) {
	argCity := inputMessage.CommandArguments()
	w := c.weatherService.Show(argCity)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, buildOutput(w, argCity))
	c.bot.Send(msg)
}

func buildOutput(w *weather.WeatherList, cityName string) string {
	var outputMsgText strings.Builder
	outputMsgText.WriteString("Wheather in ")
	outputMsgText.WriteString(cityName)
	outputMsgText.WriteString("\n\n")
	outputMsgText.WriteString("Temperature: ")
	outputMsgText.WriteString(strconv.FormatFloat(w.List[0].MainInfo.Temperature, 'f', 0, 64))
	outputMsgText.WriteString("ºC\n")
	outputMsgText.WriteString("Pressure: ")
	outputMsgText.WriteString(strconv.Itoa(w.List[0].MainInfo.Pressure))
	outputMsgText.WriteString("\n")
	outputMsgText.WriteString("Sea level: ")
	outputMsgText.WriteString(strconv.Itoa(w.List[0].MainInfo.SeaLevel))
	outputMsgText.WriteString("\n")
	outputMsgText.WriteString("Humidity: ")
	outputMsgText.WriteString(strconv.FormatFloat(w.List[0].MainInfo.Humidity, 'f', 0, 64))
	outputMsgText.WriteString("%\n")

	return outputMsgText.String()
}
