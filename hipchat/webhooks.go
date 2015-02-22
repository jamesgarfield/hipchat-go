package hipchat

import (
	"encoding/json"
)

//Represents a room_message webhook event message
type RoomMessage struct {
	Event string           `json:"event"`
	Item  RoomMessageEvent `json:"item`
}

type RoomMessageEvent struct {
	Message WebhookMessage `json:"message"`
	Room    WebhookRoom    `json:"room"`
}

type WebhookMessage struct {
	Date    string `json:"date"`
	File    *File  `json:"file"`
	Id      string `json:"id"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type File struct {
	Name     string `json:"name"`
	Size     int    `json:"size"`
	ThumbURL string `json:"thumb_url"`
	URL      string `json:"url"`
}

type WebhookRoom struct {
	Id    string `json:"id"`
	Links struct {
		Members       string `json:"members"`
		Participtants string `json:"participants"`
		Self          string `json:"self"`
		Webhooks      string `json:"webhooks"`
	}
	Name string `json:"name"`
}

func (rm *RoomMessage) UnmarshallJSON(data []byte) (err error) {
	err = json.Unmarshal(data, rm)
	return
}
