package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/odanaraujo/currency-exchange/src/controller"
	"github.com/odanaraujo/currency-exchange/src/controller/routes"
	"github.com/odanaraujo/currency-exchange/src/model/service"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	//init dependencies
	s := service.NewExchangeDomainService()
	exchangeController := controller.NewExchangeControllerInterface(s)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, exchangeController)

	if err := router.Run(":1987"); err != nil {
		log.Fatal(err)
	}
}
