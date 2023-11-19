package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mickisasi/drivefinder/controllers"
	"github.com/mickisasi/drivefinder/database"
)

func UserRoutes(db *database.Database, routes *gin.Engine) {
	routes.GET("/users", controllers.GetUsers(db))
	routes.POST("/users/", controllers.CreateUser(db))
}
