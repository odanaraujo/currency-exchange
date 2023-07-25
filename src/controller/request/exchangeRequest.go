package request

type ExchangeRequest struct {
	FromCurrency string  `json:"from"`
	ToCurrency   string  `json:"to"`
	Amount       float64 `json:"amount"`
	Rate         float64 `json:"rate"`
}

func (er *ExchangeRequest) ValidTypeFloat() bool {
	return er.Rate == 0 || er.Amount == 0
}

func (er *ExchangeRequest) ValidTypeString() bool {
	return er.FromCurrency == "" || er.ToCurrency == "" || len(er.FromCurrency) < 2 || len(er.ToCurrency) < 2
}
