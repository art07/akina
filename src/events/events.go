package events

import (
	"art/bots/akina/src/datalab"
	"time"
)

func CheckEvents(todayTime *time.Time) {
	eventsToday := make([]datalab.Event, 0, 8)
	for _, eventFromArr := range *datalab.GetDl().EventsArr {
		_, eventM, eventD := eventFromArr.Date.Date()
		_, todayM, todayD := todayTime.Date()
		if eventD == todayD && eventM == todayM {
			eventsToday = append(eventsToday, eventFromArr)
		}
	}

	if len(eventsToday) != 0 {
		sendEventsAsMsgsToChats(&eventsToday)
	}
}

func sendEventsAsMsgsToChats(eventsToday *[]datalab.Event) {
	for _, event := range *eventsToday {
		for _, chatId := range event.Chats {
			datalab.GetDl().Akina.SendMsg(chatId, event.Description)
		}
	}
}
