package currencies

type CurrencyPrivat struct {
	CurrencyName string `json:"ccy"`
	BaseCcy      string `json:"base_ccy"`
	Buy          string `json:"buy"`
	Sale         string `json:"sale"`
	awr          Awr
}

func (c *CurrencyPrivat) makeAwr() {
	c.awr.Flag1 = getFlag(c.CurrencyName)
	c.awr.Price = preparePrice(c.CurrencyName, c.Buy, c.Sale)
	c.awr.Flag2 = getFlag(c.BaseCcy)
}

type CurrencyNbu struct {
	R030         int     `json:"r030"`
	Txt          string  `json:"txt"`
	Rate         float64 `json:"rate"`
	CurrencyName string  `json:"cc"`
	ExchangeDate string  `json:"exchangedate"`
	awr          Awr
}

func (c *CurrencyNbu) makeAwr() {
	c.awr.Flag1 = getFlag(c.CurrencyName)
	c.awr.Price = preparePrice(c.Rate)
	c.awr.Flag2 = getFlag("USD")
}

type Awr struct {
	Flag1 string
	Price string
	Flag2 string
}
