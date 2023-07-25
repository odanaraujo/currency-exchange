package model

import (
	"github.com/odanaraujo/currency-exchange/config/exception"
	"github.com/odanaraujo/currency-exchange/config/logger"
	"go.uber.org/zap"
)

func (ed *ExchangeDomain) SaveConvertedCurrencies(ExchangeDomain) *exception.Exception {
	logger.Info("init saveConvertedCurrencies model", zap.String("journey", "saveConvertedCurrencies"))
	return nil
}
