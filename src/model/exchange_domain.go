package model

import "github.com/odanaraujo/currency-exchange/config/exception"

func NewExchangeDomain(fromCurrency, toCurrency string, amount, rate float64) *ExchangeDomain {
	return &ExchangeDomain{fromCurrency, toCurrency, amount, rate}
}

type ExchangeDomain struct {
	FromCurrency string
	ToCurrency   string
	Amount       float64
	Rate         float64
}

type ExchangeDomainInterface interface {
	SaveConvertedCurrencies(ExchangeDomainInterface) *exception.Exception
}

func (ed *ExchangeDomain) ValidTypeFloat() bool {
	return ed.Rate == 0 || ed.Amount == 0
}

func (ed *ExchangeDomain) ValidTypeString() bool {
	return ed.FromCurrency == "" || ed.ToCurrency == "" || len(ed.FromCurrency) < 2 || len(ed.ToCurrency) < 2
}
