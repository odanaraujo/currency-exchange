package response

type ExchangeResponse struct {
	CurrencyConverter float64 `json:"currency_converter"`
	CurrencySimbol    string  `json:"currency_simbol"`
}
