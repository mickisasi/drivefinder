package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mickisasi/drivefinder/controllers"
	"github.com/mickisasi/drivefinder/database"
)

func ListingRoutes(db *database.Database, routes *gin.Engine) {
	routes.GET("/listings", controllers.GetListings(db))
}
