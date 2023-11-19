package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mickisasi/drivefinder/database"
	"github.com/mickisasi/drivefinder/models"
	"github.com/mickisasi/drivefinder/sites"
	"go.mongodb.org/mongo-driver/bson"
)

func GetMonitors(db *database.Database) gin.HandlerFunc {
	monitorCollection := db.OpenCollection("monitors")

	return func(c *gin.Context) {
		cursor, err := monitorCollection.Find(c, bson.D{})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "error occured while fetching monitors"},
			)
			return
		}

		var monitors []models.Monitor
		if err = cursor.All(c, &monitors); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "error occured while decoding monitors"},
			)
			return
		}

		c.JSON(http.StatusOK, monitors)
	}
}

func CreateMonitor(db *database.Database) gin.HandlerFunc {
	monitorCollection := db.OpenCollection("monitors")

	return func(c *gin.Context) {
		var monitor models.Monitor

		if err := c.BindJSON(&monitor); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if monitor.Keywords == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "keywords cannot be empty"})
			return
		}

		if monitor.PostalCode == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "postalCode cannot be empty"})
			return
		}

		// TODO: ADD VERIFICATION WITH SUPABASE
		if monitor.UserID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "UserId cannot be empty"})
			return
		}

		monitor.CreatedAt = time.Now()

		_, err := monitorCollection.InsertOne(context.Background(), monitor)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// TODO: START MONITOR HERE

		m := sites.Monitor{
			Keywords:   monitor.Keywords,
			UserId:     monitor.UserID,
			PostalCode: monitor.PostalCode,
			MaxPrice:   monitor.MaxPrice,
		}
		go sites.StartMonitor(db, &m)
		c.IndentedJSON(http.StatusCreated, monitor)
	}
}
