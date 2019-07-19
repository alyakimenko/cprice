package main

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/getlantern/systray"
	"github.com/getsentry/sentry-go"
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
	s := &state{
		SelectedCurrency: btc,
		CurrencyNames: map[string]string{
			btc: "bitcoin",
			eth: "ethereum",
			xrp: "ripple",
			ltc: "litecoin",
		},
		MenuItems: map[string]*systray.MenuItem{},
	}
	systray.Run(s.onReady, s.onExit)
}

func (s *state) onReady() {
	sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	})

	s.updatePrice()

	s.Cron = cron.New()
	s.Cron.AddFunc("@every 30s", s.updatePrice)
	s.Cron.Start()

	for currency := range s.CurrencyNames {
		s.MenuItems[currency] = systray.AddMenuItem(currency, "")
	}

	for {
		select {
		case <-s.MenuItems[btc].ClickedCh:
			s.SelectedCurrency = btc
			s.updatePrice()
		case <-s.MenuItems[eth].ClickedCh:
			s.SelectedCurrency = eth
			s.updatePrice()
		case <-s.MenuItems[xrp].ClickedCh:
			s.SelectedCurrency = xrp
			s.updatePrice()
		case <-s.MenuItems[ltc].ClickedCh:
			s.SelectedCurrency = ltc
			s.updatePrice()
		}
	}
}

func (s *state) onExit() {
	s.Cron.Stop()
}

func (s *state) updatePrice() {
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://coinmarketcap.com/currencies/" + s.CurrencyNames[s.SelectedCurrency]
	response, err := client.Get(url)
	if err != nil {
		sentry.CaptureException(err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		sentry.CaptureException(errors.New("Status code is not OK"))
		return
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		sentry.CaptureException(err)
		return
	}

	price := document.Find(".details-panel-item--price__value").Text()
	systray.SetTitle(s.SelectedCurrency + " $" + price)
}
