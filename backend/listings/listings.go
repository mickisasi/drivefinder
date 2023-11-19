package listings

import (
	"context"
	"fmt"
	"time"

	"github.com/mickisasi/drivefinder/database"
	"github.com/mickisasi/drivefinder/models"
)

type Listing struct {
	Name        string
	Make        string
	Model       string
	Description string
	Picture     string
	Price       float64
	Url         string
	Site        string
	CreatedAt   time.Time
}

func (listing *Listing) InsertNewListing(db *database.Database) {
	listingModel := models.Listing{
		Name:        listing.Name,
		Make:        listing.Make,
		Model:       listing.Model,
		Description: listing.Description,
		Picture:     listing.Picture,
		Url:         listing.Url,
		Price:       listing.Price,
		Site:        listing.Site,
		CreatedAt:   listing.CreatedAt,
	}

	// TODO: IT INSERTS DUPS SOMETIMES
	_, err := db.OpenCollection("listings").InsertOne(context.TODO(), listingModel)
	if err != nil {
		fmt.Println(err)
		return
	}
}
