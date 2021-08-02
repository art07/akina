package commands

import (
	"art/bots/akina/internal/currencies"
	"art/bots/akina/internal/datalab"
	"art/bots/akina/internal/db"
	"art/bots/akina/internal/weather"
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
		log.Printf("%#v\n%#v\n%#v\n\n", update.Message, update.Message.From, update.Message.Chat)
	case "dbinfo":
		db.InitDbJob(1)
	default:
		datalab.GetDl().Akina.SendMsg(update.Message.Chat.ID, datalab.GetDl().Akina.MainPhrases["unknownCommand"]+update.Message.Text)
	}
}
