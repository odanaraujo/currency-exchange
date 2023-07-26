package service

import (
	"github.com/odanaraujo/currency-exchange/config/exception"
	"github.com/odanaraujo/currency-exchange/config/logger"
	"github.com/odanaraujo/currency-exchange/src/model"
	"go.uber.org/zap"
)

func (ed *exchangeDomainService) SaveExchangeCurrency(domainInterface model.ExchangeDomainInterface) *exception.Exception {

	logger.Info("init save exchange currency model", zap.String("journey", "ExchangeCurrency"))
	if err := ed.validTypeString(domainInterface); !err {
		ex := exception.NewBadRequestError("Incorrect or empty field type")
		return ex
	}
	if err := ed.validTypeFloat(domainInterface); !err {
		ex := exception.NewBadRequestError("Incorrect or empty field type")
		return ex
	}
	return nil
}

func (ed *exchangeDomainService) validTypeFloat(domainInterface model.ExchangeDomainInterface) bool {
	return domainInterface.GetRate() == 0 || domainInterface.GetAmount() == 0
}

func (ed *exchangeDomainService) validTypeString(domainInterface model.ExchangeDomainInterface) bool {
	return domainInterface.GetFromCurrency() == "" || domainInterface.GetToCurrency() == "" ||
		len(domainInterface.GetFromCurrency()) < 2 || len(domainInterface.GetToCurrency()) < 2
}
