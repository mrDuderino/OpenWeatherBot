package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mrDuderino/OpenWeatherBot/internal/service/weather"
	"strconv"
	"strings"
	"time"
)

func (c *Commander) CurrentWeather(inputMessage *tgbotapi.Message) {
	argCity := inputMessage.CommandArguments()
	w := c.weatherService.Show(argCity)
	messageText := ""
	if w == nil {
		messageText += "City does not exist. Or you make a mistake in city name."
	} else {
		messageText += buildOutput(w, argCity)
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, messageText)
	c.bot.Send(msg)
}

func buildOutput(w *weather.WeatherList, cityName string) string {
	var outputMsgText strings.Builder
	reportTime := time.Unix(w.List[0].Timestamp, 0)
	outputMsgText.WriteString(cityName + " report for " + reportTime.String()[:10])
	outputMsgText.WriteString("\n\n* Condition: ")
	outputMsgText.WriteString(w.List[0].WeatherDecription[0].Description)
	outputMsgText.WriteString("\n* Temperature, ºC: ")
	outputMsgText.WriteString(strconv.FormatFloat(w.List[0].MainInfo.Temperature, 'f', 0, 64))
	outputMsgText.WriteString("\n* Atmospheric pressure on the sea level, hPa: ")
	outputMsgText.WriteString(strconv.Itoa(w.List[0].MainInfo.Pressure))
	outputMsgText.WriteString("\n* Atmospheric pressure on the ground level, hPa: ")
	outputMsgText.WriteString(strconv.Itoa(w.List[0].MainInfo.GroundLevel))
	outputMsgText.WriteString("\n* Humidity, %: ")
	outputMsgText.WriteString(strconv.FormatFloat(w.List[0].MainInfo.Humidity, 'f', 0, 64))

	return outputMsgText.String()
}
