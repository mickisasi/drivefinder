package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mickisasi/drivefinder/database"
	"github.com/mickisasi/drivefinder/listings"
	"github.com/mickisasi/drivefinder/models"
	"github.com/mickisasi/drivefinder/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetListings(db *database.Database) gin.HandlerFunc {
	listsingsCollection := db.OpenCollection("listings")

	return func(c *gin.Context) {
		cursor, err := listsingsCollection.Find(c, bson.D{})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "error occured while fetching listings"},
			)
			return
		}

		var listings []models.Listing
		if err = cursor.All(c, &listings); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "error occured while decoding listings"},
			)
			return
		}

		c.JSON(http.StatusOK, listings)
	}
}

func CreateListing(db *database.Database) gin.HandlerFunc {
	listingCollection := db.OpenCollection("listings")

	return func(c *gin.Context) {
		var listingModel models.Listing

		if err := c.BindJSON(&listingModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		listingModel.CreatedAt = time.Now()

		_, err := listingCollection.InsertOne(context.Background(), listingModel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// TODO: START MONITOR HERE
		newListing := listings.Listing{
			Name:        listingModel.Name,
			Make:        listingModel.Make,
			Model:       listingModel.Model,
			Description: listingModel.Description,
			Url:         listingModel.Url,
			Site:        listingModel.Site,
			Picture:     listingModel.Picture,
			CreatedAt:   time.Now(),
			Price:       listingModel.Price,
		}

		NotifyUser("urmom2", &newListing, db)

		c.IndentedJSON(http.StatusCreated, listingModel)
	}
}

func NotifyUser(userId string, currentListing *listings.Listing, db *database.Database) {
	userCollection := db.OpenCollection("users")

	filter := bson.D{{"access_token", userId}}
	var result models.User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return
		}
		fmt.Println(err)
		return
	}

	// we have User
	if result.DiscordWebhook != "" {
		fmt.Println("Sending Webhook....")
		hook := utils.Webhook{
			URL:      result.DiscordWebhook,
			Username: "SUre",
		}

		hook.SendMessage("You made it here!")
	}

	if result.PhoneNumber != "" {
		fmt.Println("Starting phone call...")
	}
}
