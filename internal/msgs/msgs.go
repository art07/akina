package msgs

import (
	"art/bots/akina/internal/datalab"
	"art/bots/akina/internal/youtube"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

func ChooseAction(u tgbotapi.Update) {
	switch {
	case strings.Contains(u.Message.Text, "thanks"):
		datalab.GetDl().Akina.SendMsg(u.Message.Chat.ID, datalab.GetDl().Akina.MainPhrases["yaw"])
	case strings.Contains(u.Message.Text, "tyt"):
		youtube.CheckTheBestVideo(&datalab.GetDl().Youtube.Categories[0])
		youtube.CheckTheBestVideo(&datalab.GetDl().Youtube.Categories[1])
		youtube.CheckTheBestVideo(&datalab.GetDl().Youtube.Categories[2])
	default:
		datalab.GetDl().Akina.SendMsgAsAnswer(u.Message.Chat.ID, u.Message.MessageID, fmt.Sprintf("Hi, %s!", u.Message.From.FirstName))
	}
}
