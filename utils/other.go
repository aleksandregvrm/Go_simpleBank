package util

import "fmt"

type OtherStruct struct {
	usefulness bool
	*MyDataMessaging
}

var MyStructFunction *MyDataMessaging

type CurrencyType string

var AvailableCurrencies = map[CurrencyType]string{
	"EUR": "Euro",
	"USD": "United States Dollar",
	"GBP": "British Pound",
}

const (
	currencyEUR CurrencyType = "EUR"
	CurrencyUSD CurrencyType = "something"
)

func (other OtherStruct) otherFunction() {
	if !other.usefulness {
		fmt.Printf("here comes the %s", MyStructFunction.name)
	}
}
