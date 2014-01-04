package als

func MoneyInUsdCents(currency string, amount int) int {
	rates, ok := currency_table[currency]
	if !ok {
		panic("invalid currency: " + currency)
	}

	cents := rates * float64(amount)
	return int(cents)
}
