package als

import (
	"fmt"
)

func MoneyInUsdCents(currency string, amount int) int {
	rates, ok := currency_table[currency]
	if !ok {
		fmt.Printf("invalid currency: %s\n", currency)
		return 0
	}

	cents := rates * float64(amount)
	return int(cents)
}
