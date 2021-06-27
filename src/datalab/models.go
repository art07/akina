package datalab

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

type dataLab struct {
	Akina      *akina
	Owm        *owm
	Cities     string
	EventsArr  *[]Event
	Heroku     *heroku
	Currencies *currencies
	Youtube    *youtube
	Loc        *time.Location
}

type akina struct {
	Bot         *tgbotapi.BotAPI
	Name        string
	Meaning     string
	About       string
	Token       string
	MainPhrases map[string]string
}

func (a *akina) SendMsg(id int64, botMsg string) {
	msgConf := tgbotapi.NewMessage(id, botMsg)
	msgConf.ParseMode = "html"
	_, _ = a.Bot.Send(msgConf)
}

func (a *akina) SendMsgAsAnswer(id int64, msgId int, botMsg string) {
	msgConf := tgbotapi.NewMessage(id, botMsg)
	msgConf.ReplyToMessageID = msgId
	msgConf.ParseMode = "html"
	_, _ = a.Bot.Send(msgConf)
}

type Event struct {
	Date        time.Time
	Description string
	Chats       []int64
}

type youtube struct {
	Name               string
	Token              string
	MainPartOfHttpUrl  string
	MainPartOfHttpUrl2 string
	MainPartOfYbUrl    string
	Categories         []Category
}

type owm struct {
	Name           string
	MainPartOfLink string
	Token          string
}

type heroku struct {
	Name          string
	HerokuAppLink string
}

type currencies struct {
	PrivatBankApi string
	NbuApi        string
}

type Category struct {
	Name             string
	BestChannelsArr  []string
	LastWatchedVideo string
}
