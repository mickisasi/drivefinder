package sites

import (
	"context"
	"time"

	"github.com/mickisasi/drivefinder/database"
	"github.com/mickisasi/drivefinder/models"
	"github.com/mickisasi/drivefinder/sites/autotrader"
	"go.mongodb.org/mongo-driver/bson"
)

type Monitor struct {
	Keywords   string
	UserId     string
	PostalCode string
	MaxPrice   float64
}

type IExternalApi interface {
	StartChecking(string, string)
	GetSite() string
}

func LoadMonitors(db *database.Database) []*Monitor {
	monitorsCollection := db.OpenCollection("monitors")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := monitorsCollection.Find(ctx, bson.D{})
	if err != nil {
		return []*Monitor{}
	}

	var monitorsModels []models.Monitor
	var monitors []*Monitor

	if err = cursor.All(ctx, &monitorsModels); err != nil {
		return []*Monitor{}
	}

	for _, monitor := range monitorsModels {
		newMonitor := &Monitor{
			Keywords:   monitor.Keywords,
			UserId:     monitor.UserID,
			PostalCode: monitor.PostalCode,
			MaxPrice:   monitor.MaxPrice,
		}
		monitors = append(monitors, newMonitor)
	}

	return monitors
}

func StartMonitor(db *database.Database, monitor *Monitor) {
	at := autotrader.New(db)
	go at.StartChecking(monitor.Keywords, monitor.PostalCode, monitor.UserId, monitor.MaxPrice, db)
}
