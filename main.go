package main

import (
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/getlantern/systray"
	"github.com/robfig/cron"
)

const (
	btc = "BTC"
	eth = "ETH"
	xrp = "XRP"
	ltc = "LTC"
)

type state struct {
	Cron             *cron.Cron
	SelectedCurrency string
	CurrencyNames    map[string]string
	MenuItems        map[string]*systray.MenuItem
}

func main() {
	s := &state{}
	systray.Run(s.onReady, s.onExit)
}

func (s *state) onReady() {
	s.updatePrice()

	s.Cron = cron.New()
	s.Cron.AddFunc("@every 10s", s.updatePrice)
	s.Cron.Start()
}

func (s *state) onExit() {
	s.Cron.Stop()
}

func (s *state) updatePrice() {
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://coinmarketcap.com/currencies/bitcoin/"
	response, err := client.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return
	}

	price := document.Find(".details-panel-item--price__value").Text()
	systray.SetTitle("BTC $" + price)
}
