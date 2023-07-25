package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/currency-exchange/config/exception"
	"github.com/odanaraujo/currency-exchange/config/logger"
	"github.com/odanaraujo/currency-exchange/src/controller/response"
	"github.com/odanaraujo/currency-exchange/src/model"
	"go.uber.org/zap"
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

	domain := model.NewExchangeDomain(from, to, amount, rate)

	validateParameterFields(c, *domain)

	if err := domain.SaveConvertedCurrencies(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("successful conversion",
		zap.String("journey", "CurrencyConverterController"))

	res := response.ExchangeResponse{domain.Rate, "$"}

	c.JSON(http.StatusOK, res)
}

func validateParameterFields(c *gin.Context, exchange model.ExchangeDomain) {

	if exchange.ValidTypeFloat() || exchange.ValidTypeString() {
		excepResponse := exception.NewBadRequestError("Parameter fields are incorrect")
		c.JSON(http.StatusBadRequest, excepResponse)
		return
	}
}
