package commands

import (
	"art/bots/akina/src/currencies"
	"art/bots/akina/src/datalab"
	"art/bots/akina/src/weather"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func ChooseAction(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "hi":
		datalab.GetDl().Akina.SendMsg(update.Message.Chat.ID, datalab.GetDl().Akina.MainPhrases["greeting"])
	case "w":
		weather.CheckWeather(update.Message.CommandArguments(), update.Message.Chat.ID)
	case "c":
		currencies.CheckCurrencies(update.Message.Chat.ID)
	case "about":
		datalab.GetDl().Akina.SendMsg(update.Message.Chat.ID, datalab.GetDl().Akina.About)
	case "chatinfo":
		log.Printf("%#v\n%#v\n%#v\n", update.Message, update.Message.From, update.Message.Chat)
		log.Printf("\nLastWatchedVideos:\n%s\n", (*datalab.GetDl().Youtube.Categories)[0].LastWatchedVideo)
		log.Printf("%s\n", (*datalab.GetDl().Youtube.Categories)[1].LastWatchedVideo)
		log.Printf("%s\n", (*datalab.GetDl().Youtube.Categories)[2].LastWatchedVideo)
	default:
		datalab.GetDl().Akina.SendMsg(update.Message.Chat.ID, datalab.GetDl().Akina.MainPhrases["unknownCommand"]+update.Message.Text)
	}
}
