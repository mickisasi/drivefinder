package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mickisasi/drivefinder/database"
	"github.com/mickisasi/drivefinder/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(db *database.Database) gin.HandlerFunc {
	monitorCollection := db.OpenCollection("users")

	return func(c *gin.Context) {
		cursor, err := monitorCollection.Find(c, bson.D{})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "error occured while fetching monitors"},
			)
			return
		}

		var monitors []models.User
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

func CreateUser(db *database.Database) gin.HandlerFunc {
	userCollection := db.OpenCollection("users")

	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user.AccessToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "access token cannot be empty"})
			return
		}

		// TODO: ADD VERIFICATION WITH SUPABASE

		user.CreatedAt = time.Now()

		_, err := userCollection.InsertOne(context.Background(), user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusCreated, user)
	}
}
