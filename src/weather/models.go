package weather

import (
	"fmt"
	"math"
)

type weather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
	Emoji    string
}

func (w *weather) setEmoji() {
	emj := "🚀"
	switch {
	case 233 > w.Weather[0].ID && w.Weather[0].ID > 199:
		emj = "🌩"
	case 322 > w.Weather[0].ID && w.Weather[0].ID > 299:
		emj = "🌧"
	case 532 > w.Weather[0].ID && w.Weather[0].ID > 499:
		emj = "🌦"
	case 623 > w.Weather[0].ID && w.Weather[0].ID > 599:
		emj = "❄"
	case 782 > w.Weather[0].ID && w.Weather[0].ID > 700:
		emj = "🌫"
	case w.Weather[0].ID == 800:
		emj = "☀"
	case 805 > w.Weather[0].ID && w.Weather[0].ID > 800:
		emj = "☁"
	}
	w.Emoji = emj
}

func (w *weather) makeStringWeather() string {
	return fmt.Sprintf("✅ <b>%s: %.0f°С</b> %s\nWind: %.0fm/s. Clouds: %d%%\n",
		w.Name, math.Round(w.Main.Temp), w.Emoji, math.Round(w.Wind.Speed), w.Clouds.All)
}
