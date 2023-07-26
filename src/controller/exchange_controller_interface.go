package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/currency-exchange/src/model/service"
)

func NewExchangeControllerInterface(exchangeService service.ExchangeDomainService) ExchangeControllerInterface {
	return &exchageControllerInterface{service: exchangeService}
}

type ExchangeControllerInterface interface {
	CurrencyConverter(c *gin.Context)
}

type exchageControllerInterface struct {
	service service.ExchangeDomainService
}
