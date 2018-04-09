package models

import (
	"time"

	"github.com/pborman/uuid"
)

type Client struct {
	UUID     string    `json:"uuid"`
	Created  time.Time `json:"created"`
	Hostname string    `json:"hostname"`
}

func NewClient(hostname string) *Client {
	return &Client{
		UUID:     uuid.New(),
		Created:  time.Now(),
		Hostname: hostname,
	}
}
