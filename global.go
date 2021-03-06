package als

import (
	"github.com/abh/geoip"
)

const (
	KEY_TYPE_STRING   = "string" // default type
	KEY_TYPE_IP       = "ip"
	KEY_TYPE_FLOAT    = "float"
	KEY_TYPE_INT      = "int"
	KEY_TYPE_MONEY    = "money"
	KEY_TYPE_RANGE    = "range"
	KEY_TYPE_BASEFILE = "base_file"

	KEY_NAME_CURRENCY = "currency"
	KEY_NAME_IP       = "ip"
)

var (
	geo *geoip.GeoIP

	currency_table = map[string]float64{
		"VND": 0.000047,
		"NZD": 0.84,
		"HUF": 0.0045,
		"GBP": 1.6,
		"COP": 0.00053,
		"MXN": 0.078,
		"PHP": 0.023,
		"AUD": 0.94,
		"PLN": 0.32,
		"EUR": 1.35,
		"THB": 0.032,
		"MYR": 0.32,
		"BRL": 0.45,
		"INR": 0.016,
		"CAD": 0.97,
		"SAR": 0.27,
		"VEF": 0.16,
		"ARS": 0.17,
		"CZK": 0.052,
		"DKK": 0.18,
		"USD": 1.0,
		"CLP": 0.002007,
		"ZAR": 0.100801,
		"PEN": 0.36049,
		"NOK": 0.166719,
		"TRY": 0.503221,
		"GTQ": 0.12550,
		"IDR": 0.00009,
		"SEK": 0.154521,
		"SGD": 0.802761,
		"AED": 0.272257,
		"RUB": 0.030961,
		"CHF": 1.09902,
		"TWD": 0.034031,
		"ILS": 0.282239,
		"HKD": 0.128961,
	}
)
