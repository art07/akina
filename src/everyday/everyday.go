package everyday

import (
	"art/bots/akina/src/currencies"
	"art/bots/akina/src/datalab"
	"art/bots/akina/src/events"
	"art/bots/akina/src/weather"
	"art/bots/akina/src/youtube"
	"time"
)

func StartEveryDayJob(atTime int) {
	for {
		h, m, _ := time.Now().In(datalab.GetDl().Loc).Clock()
		if h == atTime && m == 0 {
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
				youtube.CheckTheBestVideo(&(*datalab.GetDl().Youtube.Categories)[0])
			case time.Saturday:
				youtube.CheckTheBestVideo(&(*datalab.GetDl().Youtube.Categories)[1])
			case time.Sunday:
				youtube.CheckTheBestVideo(&(*datalab.GetDl().Youtube.Categories)[2])
			}
		}()

		// ПОТОК X переходит в ожидание.
		time.Sleep(time.Hour * 24)
	}
}
