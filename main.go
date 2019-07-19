package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	getPrice()
}

func onExit() {}

func getPrice() {
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
