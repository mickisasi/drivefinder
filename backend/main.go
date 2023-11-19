package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mickisasi/drivefinder/database"
	"github.com/mickisasi/drivefinder/routes"
	"github.com/mickisasi/drivefinder/sites"
)

func main() {
	fmt.Println("Starting Backend...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")
	DB_URL := os.Getenv("MONGO_DB_URL")

	db, err := database.New(DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.MonitorRoutes(db, router)
	routes.ListingRoutes(db, router)
	routes.UserRoutes(db, router)

	existingMonitors := sites.LoadMonitors(db)
	for _, monitor := range existingMonitors {
		go sites.StartMonitor(db, monitor)
	}

	err = router.Run(":" + PORT)
	if err != nil {
		log.Fatal(err)
	}
}
