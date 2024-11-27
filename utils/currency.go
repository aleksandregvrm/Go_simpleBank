package util

var SupportedCurrencies = [8]string{
	"USD",
	"EUR",
	"CAD",
	"AED",
	"GEL",
	"GBP",
	"BRL",
	"RUB",
}

func IsSupportedCurrency(currency string) bool {
	for _, avalCurrency := range SupportedCurrencies {
		if avalCurrency == currency {
			return true
		}
	}
	return false
}
