package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Webhook struct {
	URL       string
	Username  string  `json:"username"`
	AvatarUrl string  `json:"avatar_url"`
	Embeds    []Embed `json:"embeds"`
	Content   string  `json:"content"`
}

type Author struct {
	Name    string `json:"name"`
	IconUrl string `json:"icon_url"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Footer struct {
	Text    string `json:"text"`
	IconUrl string `json:"icon_url"`
}

type Thumbnail struct {
	Url string `json:"url"`
}

type Image struct {
	Url string `json:"url"`
}

type Embed struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Color       int       `json:"color"`
	Author      Author    `json:"author"`
	Footer      Footer    `json:"footer"`
	Fields      []Field   `json:"fields"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Image       Image     `json:"image"`
	Timestamp   string    `json:"timestamp"`
}

func (embed *Embed) AddField(name, value string, inline bool) {
	field := Field{
		Name:   name,
		Value:  value,
		Inline: inline,
	}
	embed.Fields = append(embed.Fields, field)
}

func (embed *Embed) AddThumbnail(imageUrl string) {
	thumbnail := Thumbnail{Url: imageUrl}
	embed.Thumbnail = thumbnail
}

func (embed *Embed) AddImage(imageUrl string) {
	image := Image{Url: imageUrl}
	embed.Image = image
}

func (embed *Embed) UseTimestamp() {
	currentTime := time.Now().UTC().Format(time.RFC3339)
	embed.Timestamp = currentTime
}

func (webhook *Webhook) SendMessage(message string) {
	data := &Webhook{
		Username:  webhook.Username,
		AvatarUrl: webhook.AvatarUrl,
		Embeds:    nil,
		Content:   message,
	}
	payload, _ := json.Marshal(data)

	_, err := http.Post(webhook.URL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return
	}
}

func (webhook *Webhook) SendEmbed(embed Embed) {
	embedList := make([]Embed, 1)
	embedList[0] = embed

	data := &Webhook{
		Username:  webhook.Username,
		AvatarUrl: webhook.AvatarUrl,
		Embeds:    embedList,
	}

	payload, _ := json.Marshal(data)

	_, err := http.Post(webhook.URL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return
	}
}
