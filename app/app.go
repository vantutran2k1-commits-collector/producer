package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vantutran2k1-commits-collector/producer/app/routes"
	"github.com/vantutran2k1-commits-collector/producer/config"
	"log"
)

type CollectorApp struct {
	Router *gin.Engine
}

func InitApp() *CollectorApp {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	config.InitAppEnv()
	config.InitDb()
	config.InitKafkaProducerClient()

	router := routes.RegisterRoutes()
	return &CollectorApp{
		Router: router,
	}
}
