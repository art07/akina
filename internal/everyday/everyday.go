package everyday

import (
	"art/bots/akina/internal/currencies"
	"art/bots/akina/internal/datalab"
	"art/bots/akina/internal/events"
	"art/bots/akina/internal/weather"
	"art/bots/akina/internal/youtube"
	"time"
)

func StartEveryDayJob(atH, atM int) {
	for {
		h, m, _ := time.Now().In(datalab.GetDl().Loc).Clock()
		if h == atH && m == atM {
			break
		}
	}

	for {
		// Запуск горутины из ПОТОКА X.
		go func() {
			todayTime := time.Now().In(datalab.GetDl().Loc)

			weather.CheckWeather(datalab.GetDl().Cities, datalab.ToOurGroup)
			events.CheckEvents(&todayTime)

			switch todayTime.Weekday() {
			case time.Monday:
				// do some job
			case time.Tuesday:
				// do some job
			case time.Wednesday:
				currencies.CheckCurrencies(datalab.ToOurGroup)
			case time.Thursday:
				// do some job
			case time.Friday:
				youtube.CheckTheBestVideo(&datalab.GetDl().Youtube.Categories[0])
			case time.Saturday:
				youtube.CheckTheBestVideo(&datalab.GetDl().Youtube.Categories[1])
			case time.Sunday:
				youtube.CheckTheBestVideo(&datalab.GetDl().Youtube.Categories[2])
			}
		}()

		// ПОТОК X переходит в ожидание.
		time.Sleep(time.Hour * 24)
	}
}
