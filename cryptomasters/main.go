package main

import "cryptomasters.com/api"

func main() {
	rate, err := api.GetRate("BTC")
	print(rate.Price, err)
}
