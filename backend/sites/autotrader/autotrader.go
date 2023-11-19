package autotrader

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/mickisasi/drivefinder/database"
	"github.com/mickisasi/drivefinder/listings"
	"github.com/mickisasi/drivefinder/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AutoTraderApi struct {
	keywords   string
	postalCode string
	database   *database.Database
}

func New(db *database.Database) *AutoTraderApi {
	return &AutoTraderApi{database: db}
}

func (api *AutoTraderApi) GetSite() string {
	return "autotrader"
}

func (api *AutoTraderApi) FetchRecentListings(
	keywords string,
	postalCode string,
) ([]*listings.Listing, error) {
	url := "https://mktapi.autotrader.ca/api/ad/search/v2"

	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(profiles.Chrome_110),
		tls_client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	params := req.URL.Query()
	params.Add("c2", "7,9,10,11")
	params.Add("ptradloc", fmt.Sprintf("50,%s", postalCode))
	params.Add("haspr", "True")
	params.Add("hcp", "True")
	params.Add("q", keywords)
	params.Add("srt", "2")
	params.Add("t", "20")
	params.Add("p", "0")

	req.URL.RawQuery = params.Encode()

	req.Header.Add("User-Agent", "samsung SM-G955N|Android 7.1.2|8.14.0.459")
	req.Header.Add("Authorization", "Basic MjE4MzA1MTg2Njo6")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Accept-Language", "en-ca")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("autotrader: invalid request response")
	}

	var data SuccessResponse
	json.NewDecoder(resp.Body).Decode(&data)

	var list []*listings.Listing

	for _, item := range data.Summaries {
		newListing := listings.Listing{
			Name:        fmt.Sprintf("%d %s %s %s", item.Year, item.Make, item.Model, item.Trim),
			Make:        item.Make,
			Model:       item.Model,
			Description: "",
			Url:         item.ContentURL,
			Site:        "autotrader",
			Picture:     item.MainPhotoURL,
			CreatedAt:   time.Now(),
			Price:       float64(item.Price),
		}
		list = append(list, &newListing)
	}

	if len(list) == 0 {
		return nil, errors.New("failed to fetch listings")
	}

	return list, nil
}

func (api *AutoTraderApi) StartChecking(
	keywords, postalCode, userId string,
	maxPrice float64,
	db *database.Database,
) {
	for {
		fmt.Printf("autotrader: checking for new listings with keywords: %s\n", keywords)

		currentListings, err := api.FetchRecentListings(keywords, postalCode)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, currentListing := range currentListings {
			currentListing.InsertNewListing(api.database)
		}

		for _, currentListing := range currentListings {
			if currentListing.Price <= maxPrice {
				NotifyUser(userId, currentListing, db)
			}
		}
		time.Sleep(time.Minute * 10)
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
	}

	if result.PhoneNumber != "" {
		fmt.Println("Starting phone call...")
	}
}
