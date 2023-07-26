package model

type ExchangeDomainInterface interface {
	GetFromCurrency() string
	GetToCurrency() string
	GetAmount() float64
	GetRate() float64
}

func NewExchangeDomain(fromCurrency, toCurrency string, amount, rate float64) *exchangeDomain {
	return &exchangeDomain{fromCurrency, toCurrency, amount, rate}
}

type exchangeDomain struct {
	fromCurrency string
	toCurrency   string
	amount       float64
	rate         float64
}

func (ed *exchangeDomain) GetFromCurrency() string {
	return ed.fromCurrency
}
func (ed *exchangeDomain) GetToCurrency() string {
	return ed.toCurrency
}
func (ed *exchangeDomain) GetAmount() float64 {
	return ed.amount
}
func (ed *exchangeDomain) GetRate() float64 {
	return ed.rate
}
