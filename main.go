package main

import (
	"art/bots/akina/src/commands"
	"art/bots/akina/src/datalab"
	"art/bots/akina/src/db"
	"art/bots/akina/src/everyday"
	"art/bots/akina/src/msgs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	dl := datalab.GetDl()

	// Создание объекта бота.
	akinaBot, err := tgbotapi.NewBotAPI(dl.Akina.Token)
	if err == nil {
		db.InitDbJob(0)
		dl.Akina.Bot = akinaBot
		dl.Akina.Bot.Debug = false
		log.Printf("\n%s started!\n\n\n", dl.Akina.Name)
		// Запустить ПОТОК X для ежедневной работы.
		go everyday.StartEveryDayJob(10, 0)
	} else {
		log.Panic(err)
	}

	// Создать объект UpdateConfig для получения updates.
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	// В этот канал сохраняются updates.
	updatesChannel, err := dl.Akina.Bot.GetUpdatesChan(updateConfig)

	// БЕСКОНЕЧНЫЙ ЦИКЛ СЧИТЫВАНИЯ ИЗ КАНАЛА-------------------------------------->
	for update := range updatesChannel {
		//log.Printf("%#v", update.Message)
		if update.Message == nil {
			continue
		} else {
			log.Printf("\n\nFrom > {%s} Msg > {%s}\n", update.Message.From.FirstName, update.Message.Text)
		}

		if update.Message.IsCommand() {
			/*Command*/
			commands.ChooseAction(update)
		} else {
			/*Message*/
			msgs.ChooseAction(update)
		}
	}
}
