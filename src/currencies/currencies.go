package currencies

import (
	"art/bots/akina/src/datalab"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

var forXAU float64

func CheckCurrencies(chatId int64) {
	currenciesArrPB, err := getPrivatData(datalab.GetDl().Currencies.PrivatBankApi)
	if err != nil {
		log.Println("The data of the PrivateBank was not received. ")
	}

	currenciesArrNbu, err := getNbuData(datalab.GetDl().Currencies.NbuApi)
	if err != nil {
		log.Println("The data of the NBU was not received. ")
	}

	if currenciesArrPB == nil && currenciesArrNbu == nil {
		return
	}

	superCurrenciesString := ""
	if currenciesArrPB != nil {
		for _, c := range *currenciesArrPB {
			c.makeAwr()
			superCurrenciesString += fmt.Sprintf("<b>1</b> %s  ➡️  <b>%s</b> %s\n", c.awr.Flag1, c.awr.Price, c.awr.Flag2)
		}
	}
	if currenciesArrNbu != nil {
		for _, c := range *currenciesArrNbu {
			if c.CurrencyName == "XAU" {
				c.makeAwr()
				superCurrenciesString += fmt.Sprintf("<b>1</b> %s  ➡️  <b>%s</b> %s\n", c.awr.Flag1, c.awr.Price, c.awr.Flag2)
				break
			}
		}
	}

	datalab.GetDl().Akina.SendMsg(chatId, superCurrenciesString)
}

func getPrivatData(link string) (*[]CurrencyPrivat, error) {
	body, err := getBody(link)
	if err != nil {
		return nil, err
	}

	var currenciesPrivatArr []CurrencyPrivat
	_ = json.Unmarshal(body, &currenciesPrivatArr)

	return &currenciesPrivatArr, nil
}

func getNbuData(link string) (*[]CurrencyNbu, error) {
	body, err := getBody(link)
	if err != nil {
		return nil, err
	}

	var currenciesArr []CurrencyNbu
	_ = json.Unmarshal(body, &currenciesArr)

	return &currenciesArr, nil
}

func getBody(link string) ([]byte, error) {
	httpResponse, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func preparePrice(data ...interface{}) string {
	v := 0.0
	if len(data) == 3 {
		var buy, sale float64

		buy, _ = strconv.ParseFloat(data[1].(string), 64)
		sale, _ = strconv.ParseFloat(data[2].(string), 64)

		v = (buy + sale) / 2
		if data[0].(string) == "USD" {
			forXAU = v
		}
	} else {
		v = math.Round((data[0].(float64) / 31.1) / forXAU)
	}
	return strings.TrimRight(strings.TrimRight(strconv.FormatFloat(v, 'f', 2, 64), "0"), ".")
}

func getFlag(currencyName string) string {
	return map[string]string{
		"UAH": "🇺🇦",
		"USD": "🇺🇸",
		"EUR": "🇪🇺",
		"RUR": "🇷🇺",
		"BTC": "  ₿  ",
		"XAU": "💰",
	}[currencyName]
}
