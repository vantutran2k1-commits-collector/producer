package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vantutran2k1-commits-collector/producer/app/constants"
	"github.com/vantutran2k1-commits-collector/producer/app/controllers"
	"github.com/vantutran2k1-commits-collector/producer/app/repositories"
	"github.com/vantutran2k1-commits-collector/producer/app/services"
	"github.com/vantutran2k1-commits-collector/producer/config"
)

var r *Repositories
var s *Services
var c *Controllers
var router *gin.Engine

type Repositories struct {
	CollectionJobRepository repositories.JobRepository
}

type Services struct {
	CommitService services.CommitService
}

type Controllers struct {
	CommitController controllers.CommitController
}

func setupRepositories() {
	r = &Repositories{
		CollectionJobRepository: repositories.NewJobRepository(config.Db),
	}
}

func setupServices(repositories *Repositories) {
	s = &Services{
		CommitService: services.NewCommitService(
			repositories.CollectionJobRepository,
			config.KafkaProducerClient,
		),
	}
}

func setupControllers(services *Services) {
	c = &Controllers{
		CommitController: controllers.NewCommitController(services.CommitService),
	}
}

func setupRouter() {
	if config.AppEnv.GinMode == constants.GinReleaseMode {
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
	setupRepositories()
	setupServices(r)
	setupControllers(s)
	setupRouter()
}
