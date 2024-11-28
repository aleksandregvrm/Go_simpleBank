package util

var SupportedCurrencies = [9]string{
	"USD",
	"EUR",
	"CAD",
	"AED",
	"GEL",
	"GBP",
	"BRL",
	"RUB",
	"NGN",
}

func IsSupportedCurrency(currency string) bool {
	for _, avalCurrency := range SupportedCurrencies {
		if avalCurrency == currency {
			return true
		}
	}
	return false
}
