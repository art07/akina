/*
https://api.openweathermap.org/data/2.5/weather?q=mariupol&units=metric&appid=2ab2e5d6db591c7aeb7a8a64c5135dcc
*/

/*
{"coord":{"lon":37.5,"lat":47.0667},"weather":[{"id":801,"main":"Clouds","description":"few clouds",
"icon":"02d"}],"base":"stations","main":{"temp":25.61,"feels_like":25.82,"temp_min":25.61,"temp_max":25.61,
"pressure":1013,"humidity":61,"sea_level":1013,"grnd_level":1007},"visibility":10000,"wind":{"speed":7.69,
"deg":87,"gust":11.12},"clouds":{"all":20},"dt":1624193641,"sys":{"country":"UA","sunrise":1624152865,
"sunset":1624210121},"timezone":10800,"id":701822,"name":"Mariupol","cod":200}
*/

package weather

import (
	"art/bots/akina/internal/datalab"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

func CheckWeather(citiesAsString string, chatId int64) {
	if citiesAsString == "" {
		datalab.GetDl().Akina.SendMsg(chatId, "Input example: /w kyiv dnipro mariupol new_york")
		return
	}

	citiesForAnswer := make([]*weather, 0, 8)
	wrongCityNames := make([]string, 0, 8)

	for _, cityStr := range strings.Split(citiesAsString, " ") {
		// 1) Проверка валидности города.
		if !isCityValid(cityStr) {
			wrongCityNames = append(wrongCityNames, cityStr)
			continue
		}

		// 2) Создать http запрос для города.
		httpRequest, err := makeHttpRequest(strings.ToLower(strings.ReplaceAll(cityStr, "_", " ")))
		if err != nil {
			wrongCityNames = append(wrongCityNames, cityStr)
			continue
		}

		// 3) Создать объект погоды на основе http response.
		weatherInCity, err := getWeatherDataObj(httpRequest)
		if err != nil {
			wrongCityNames = append(wrongCityNames, cityStr)
			continue
		}

		// 4) Добавляю в список объектов городов.
		citiesForAnswer = append(citiesForAnswer, weatherInCity)
	}

	if len(citiesForAnswer) == 0 {
		if len(wrongCityNames) != 0 {
			datalab.GetDl().Akina.SendMsg(chatId, makeStringWrongCities(&wrongCityNames))
		}
		return
	}

	sort.SliceStable(citiesForAnswer, func(i, j int) bool {
		return citiesForAnswer[i].Main.Temp > citiesForAnswer[j].Main.Temp
	})

	superWeatherString := ""
	for _, w := range citiesForAnswer {
		superWeatherString += w.makeStringWeather()
	}

	if len(wrongCityNames) > 0 {
		superWeatherString += makeStringWrongCities(&wrongCityNames)
	}

	datalab.GetDl().Akina.SendMsg(chatId, superWeatherString)
}

func isCityValid(city string) bool {
	// https://golang.org/pkg/regexp/syntax/
	b, _ := regexp.MatchString(`^[[:alpha:]_]{4,}$`, city)
	return b
}

func makeHttpRequest(city string) (*http.Request, error) {
	httpRequest, err := http.NewRequest("GET", datalab.GetDl().Owm.MainPartOfLink, nil)
	if err != nil {
		return nil, err
	}

	query := httpRequest.URL.Query()
	query.Add("q", city)
	query.Add("units", "metric")
	query.Add("appid", datalab.GetDl().Owm.Token)
	httpRequest.URL.RawQuery = query.Encode() // ("bar=baz&foo=quux")

	return httpRequest, nil
}

func getWeatherDataObj(request *http.Request) (*weather, error) {
	var w weather

	httpResponse, err := http.Get(request.URL.String())
	if err != nil || httpResponse.StatusCode != 200 {
		return nil, errors.New("err != nil || httpResponse.StatusCode != 200")
	}
	//goland:noinspection GoUnhandledErrorResult
	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(body, &w)
	w.setEmoji()
	return &w, nil
}

func makeStringWrongCities(wrongCities *[]string) string {
	return fmt.Sprintf("<b>Wrong data:</b>\n%s", strings.Join(*wrongCities, "\n"))
}
