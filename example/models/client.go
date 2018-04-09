package models

import (
	"fmt"
	"time"

	"github.com/mickep76/kvstore"
	"github.com/pborman/uuid"
)

type Client struct {
	UUID     string    `json:"uuid"`
	Created  time.Time `json:"created"`
	Hostname string    `json:"hostname"`
}

type Clients []*Client

func NewClient(hostname string) *Client {
	return &Client{
		UUID:     uuid.New(),
		Created:  time.Now(),
		Hostname: hostname,
	}
}

func ClientAll(kvc kvstore.Conn, prefix string) (Clients, error) {
	kvs, err := kvc.Values(fmt.Sprintf("%s/%s", prefix, "clients"))
	if err != nil {
		return nil, err
	}

	clients := Clients{}
	for _, kv := range kvs {
		c := &Client{}
		kv.SetEncoding("json")
		if err := kv.Decode(c); err != nil {
			return nil, err
		}

		clients = append(clients, c)
	}

	return clients, nil
}
