package main

import (
	"github.com/anish-gyawali/assessment/server/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Run()
}
