package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/currency-exchange/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/exchange/:amount/:from/:to/:rate", controller.CurrencyConverter)
}
