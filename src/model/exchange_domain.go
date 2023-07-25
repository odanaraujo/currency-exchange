package model

import "github.com/odanaraujo/currency-exchange/config/exception"

func NewExchangeDomain(fromCurrency, toCurrency string, amount, rate float64) ExchangeDomainInterface {
	return &ExchangeDomain{fromCurrency, toCurrency, amount, rate}
}

type ExchangeDomain struct {
	FromCurrency string
	ToCurrency   string
	Amount       float64
	Rate         float64
}

type ExchangeDomainInterface interface {
	SaveConvertedCurrencies(ExchangeDomain) *exception.Exception
}
