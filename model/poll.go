package model

import "time"

type Poll struct {
	Id             int       `json:"id"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedBy      int       `json:"created_by"`
	Options        []Option  `json:"options"`
	Title          string    `json:"title"`
	CreatedAt      time.Time `json:"created_at"`
	Expired        bool      `json:"expired"`
}

type Option struct {
	Image string `json:"image"`
	Name  string `json:"name"`
}

