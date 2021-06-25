package everyday

import (
	"art/bots/akina/src/datalab"
	"time"
)

func StartEveryDayJob(atTime int) {
	//for ; ; {
	//	h, m, _ := time.Now().In(datalab.GetDl().Loc).Clock()
	//	if h == atTime && m == 0 {
	//		break
	//	}
	//}

	for {
		// Запуск горутины из ПОТОКА X.
		go func() {
			todayTime := time.Now().In(datalab.GetDl().Loc)

			//events.CheckEvents(&todayTime)
			//weather.CheckWeather(datalab.GetDl().Cities, datalab.ToMe)
			//currencies.CheckCurrencies(datalab.ToMe)
			//youtube.CheckTheBestVideo(&(*datalab.GetDl().Youtube.Categories)[0])
			//youtube.CheckTheBestVideo(&(*datalab.GetDl().Youtube.Categories)[1])
			//youtube.CheckTheBestVideo(&(*datalab.GetDl().Youtube.Categories)[2])

			switch todayTime.Weekday() {
			case time.Monday:
				// do some job
			case time.Tuesday:
				// do some job
			case time.Wednesday:
				// do some job
			case time.Thursday:
				// do some job
			case time.Friday:
				// do some job
			case time.Saturday:
				// do some job
			case time.Sunday:
				// do some job
			}
		}()

		// ПОТОК X переходит в ожидание.
		time.Sleep(time.Hour * 24)
	}
}
