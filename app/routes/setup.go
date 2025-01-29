package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vantutran2k1-commits-collector/producer/app/controllers"
	"github.com/vantutran2k1-commits-collector/producer/app/services"
	"github.com/vantutran2k1-commits-collector/producer/config"
)

var s *Services
var c *Controllers
var router *gin.Engine

type Services struct {
	CommitService services.CommitService
}

type Controllers struct {
	CommitController controllers.CommitController
}

func setupServices() {
	s = &Services{
		CommitService: services.NewCommitService(config.KafkaProducerClient),
	}
}

func setupControllers(services *Services) {
	c = &Controllers{
		CommitController: controllers.NewCommitController(services.CommitService),
	}
}

func setupRouter() {
	if config.AppEnv.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router = gin.Default()

	//router.Use(cors.New(cors.Config{
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	//	AllowCredentials: true,
	//}))
}

func setupRoutes() {
	setupServices()
	setupControllers(s)
	setupRouter()
}
