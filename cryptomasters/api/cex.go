package api

import (
	"cryptomasters.com/datatypes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)

		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("error getting rate for %s", currency)
	}
	rate := datatypes.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}
