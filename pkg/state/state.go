package state

import (
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/getlantern/systray"
	"github.com/robfig/cron"
)

const (
	BTC = "BTC"
	ETH = "ETH"
	XRP = "XRP"
	LTC = "LTC"
)

type State struct {
	Cron             *cron.Cron
	SelectedCurrency string
	CurrencyNames    map[string]string
	MenuItems        map[string]*systray.MenuItem
}

func (s *State) OnReady() {
	s.UpdatePrice()

	s.Cron = cron.New()
	s.Cron.AddFunc("@every 30s", s.UpdatePrice)
	s.Cron.Start()

	for currency := range s.CurrencyNames {
		s.MenuItems[currency] = systray.AddMenuItem(currency, "")
	}
	systray.AddSeparator()
	quit := systray.AddMenuItem("Quit", "")

	for {
		select {
		case <-s.MenuItems[BTC].ClickedCh:
			s.SelectedCurrency = BTC
			s.UpdatePrice()
		case <-s.MenuItems[ETH].ClickedCh:
			s.SelectedCurrency = ETH
			s.UpdatePrice()
		case <-s.MenuItems[XRP].ClickedCh:
			s.SelectedCurrency = XRP
			s.UpdatePrice()
		case <-s.MenuItems[LTC].ClickedCh:
			s.SelectedCurrency = LTC
			s.UpdatePrice()
		case <-quit.ClickedCh:
			systray.Quit()
		}
	}
}

func (s *State) OnExit() {
	s.Cron.Stop()
}

func (s *State) UpdatePrice() {
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://coinmarketcap.com/currencies/" + s.CurrencyNames[s.SelectedCurrency]
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
	systray.SetTitle(s.SelectedCurrency + " $" + price)
}
