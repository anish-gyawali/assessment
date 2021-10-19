package main

import (
	"github.com/anish-gyawali/assessment/server/config"
	"github.com/anish-gyawali/assessment/server/controller"
	"github.com/anish-gyawali/assessment/server/repository"
	"github.com/anish-gyawali/assessment/server/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	dataRepository repository.DataRepository = repository.NewDataRepository(db)
	dataService    service.DataService       = service.NewDataService(dataRepository)
	dataController controller.DataController = controller.NewDataController(dataService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	dataRoutes := r.Group("api/datas")
	{
		dataRoutes.GET("/", dataController.All)
		dataRoutes.POST("/", dataController.Insert)
		dataRoutes.GET("/:id", dataController.FindByID)
		dataRoutes.PUT("/:id", dataController.Update)
		dataRoutes.DELETE("/:id", dataController.Delete)
	}
	r.Run()
}
