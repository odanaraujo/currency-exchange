package service

import (
	"github.com/odanaraujo/currency-exchange/config/exception"
	"github.com/odanaraujo/currency-exchange/src/model"
)

func NewExchangeDomainService() ExchangeDomainService {
	return &exchangeDomainService{}
}

type exchangeDomainService struct {
}

type ExchangeDomainService interface {
	SaveExchangeCurrency(domainInterface model.ExchangeDomainInterface) *exception.Exception
}
