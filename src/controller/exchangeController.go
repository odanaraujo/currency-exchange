package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/currency-exchange/config/exception"
	"github.com/odanaraujo/currency-exchange/src/model"
	"github.com/odanaraujo/currency-exchange/src/model/request"
	"net/http"
	"strconv"
)

var (
	ExchangeDomainInterface model.ExchangeDomainInterface
)

func CurrencyConverter(c *gin.Context) {

	amountParams := c.Params.ByName("amount")
	from := c.Params.ByName("from")
	to := c.Params.ByName("to")
	rateParams := c.Params.ByName("rate")

	amount, _ := strconv.ParseFloat(amountParams, 64)
	rate, _ := strconv.ParseFloat(rateParams, 64)

	exchangeRequest := request.ExchangeRequest{
		FromCurrency: from,
		ToCurrency:   to,
		Amount:       amount,
		Rate:         rate,
	}

	validateParameterFields(c, exchangeRequest)

	fmt.Println(exchangeRequest)
}

func validateParameterFields(c *gin.Context, exchangeRequest request.ExchangeRequest) {

	if exchangeRequest.ValidTypeFloat() || exchangeRequest.ValidTypeString() {
		excepResponse := exception.NewBadRequestError("Parameter fields are incorrect")
		c.JSON(http.StatusBadRequest, excepResponse)
		return
	}
}
