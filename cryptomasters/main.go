package main

import (
	"cryptomasters.com/api"
	"fmt"
	"sync"
)

func main() {
	currencies := []string{"BTC", "ETH", "LTC", "BCH"}
	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go (func(c string) {
			getCurrencyData(c)
			wg.Done()
		})(currency)
	}
	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("1 %s = %.2f %s\n", currency, rate.Price, "USD")
	} else {
		println(err.Error())
	}

}
