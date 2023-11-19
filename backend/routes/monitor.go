package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mickisasi/drivefinder/controllers"
	"github.com/mickisasi/drivefinder/database"
)

func MonitorRoutes(db *database.Database, routes *gin.Engine) {
	routes.GET("/monitors", controllers.GetMonitors(db))
	routes.POST("/monitors/", controllers.CreateMonitor(db))
}
