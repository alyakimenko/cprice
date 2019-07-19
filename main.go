package main

import (
	"net/http"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("test")
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
}
