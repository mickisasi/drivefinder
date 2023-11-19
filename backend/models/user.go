package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"    json:"id"`
	AccessToken    string             `bson:"access_token"     json:"access_token"    validate:"required"`
	DiscordWebhook string             `bson:"discord_webhoook" json:"discord_webhook" validate:"required"`
	PhoneNumber    string             `bson:"phone_number"     json:"phone_number"    validate:"required"`
	CreatedAt      time.Time          `bson:"created_at"       json:"created_at"      validate:"required"` // Time it was added to DB
}
