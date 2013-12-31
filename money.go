package als

import (
	"errors"
)

func MoneyInUsd(currency string, amount float64) (usd float64, err error) {
	var ok bool
	usd, ok = currency_table[currency]
	if !ok {
		err = errors.New("invalid currency")
	}

	usd *= amount
	return
}
