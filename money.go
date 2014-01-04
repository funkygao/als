package als

// In dollar instead of cent
func MoneyInUsd(currency string, amount int) (usd int) {
	rates, ok := currency_table[currency]
	if !ok {
		panic("invalid currency: " + currency)
	}

	cents := rates * float64(amount)
	usd = int(cents) / 100

	return
}
