package main

import (
	"github.com/alyakimenko/cprice/pkg/state"
	"github.com/getlantern/systray"
)

func main() {
	s := &state.State{
		SelectedCurrency: state.BTC,
		CurrencyNames: map[string]string{
			state.BTC: "bitcoin",
			state.ETH: "ethereum",
			state.XRP: "ripple",
			state.LTC: "litecoin",
		},
		MenuItems: map[string]*systray.MenuItem{},
	}
	systray.Run(s.OnReady, s.OnExit)
}
