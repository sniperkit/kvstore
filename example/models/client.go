package models

import (
	"fmt"
	"time"

	"github.com/mickep76/kvstore"
	"github.com/mickep76/kvstore/query"
	"github.com/pborman/uuid"
)

type Client struct {
	UUID     string     `json:"uuid"`
	Created  time.Time  `json:"created"`
	Updated  *time.Time `json:"updated,omitempty"`
	Hostname string     `json:"hostname"`
}

type Clients []*Client

func NewClient(hostname string) *Client {
	return &Client{
		UUID:     uuid.New(),
		Created:  time.Now(),
		Hostname: hostname,
	}
}

func (ds *datastore) AllClients() (Clients, error) {
	kvs, err := ds.Values("clients")
	if err != nil {
		return nil, err
	}

	clients := Clients{}
	for _, kv := range kvs {
		c := &Client{}
		if err := kv.Decode(c); err != nil {
			return nil, err
		}

		clients = append(clients, c)
	}

	return clients, nil
}

func (ds *datastore) FindClient(field string, value interface{}) (*Client, error) {
	clients, err := ds.AllClients()
	if err != nil {
		return nil, err
	}

	for _, c := range clients {
		v, err := query.FieldValue(c, field)
		if err != nil {
			return nil, err
		}
		if fmt.Sprint(value) == fmt.Sprint(v) {
			return c, nil
		}
	}

	return nil, nil
}

func (ds *datastore) CreateClient(client *Client) error {
	return ds.Set(fmt.Sprintf("clients/%s", client.UUID), client, kvstore.WithLease(ds.lease))
}

func (ds *datastore) UpdateClient(client *Client) error {
	now := time.Now()
	client.Updated = &now
	return ds.Set(fmt.Sprintf("clients/%s", client.UUID), client, kvstore.WithLease(ds.lease))
}
