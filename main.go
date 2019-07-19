package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/getlantern/systray"
	"github.com/robfig/cron"
)

type state struct {
	Cron *cron.Cron
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
	url := "https://coinmarketcap.com/currencies/bitcoin/"
	response, err := http.Get(url)
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
