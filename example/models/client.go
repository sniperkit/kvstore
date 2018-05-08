package models

import (
	"fmt"
	"time"

	"github.com/mickep76/kvstore"
	"github.com/mickep76/qry"
	"github.com/pborman/uuid"
)

type Client struct {
	UUID     string     `json:"uuid" kvstore:"unique"`
	Created  time.Time  `json:"created"`
	Updated  *time.Time `json:"updated,omitempty"`
	Hostname string     `json:"hostname" kvstore:"unique"`
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
	return clients, kvs.Decode(&clients)
}

func (ds *datastore) QueryClients(q *qry.Query) (Clients, error) {
	kvs, err := ds.Values("clients")
	if err != nil {
		return nil, err
	}

	clients := Clients{}
	if err := kvs.Decode(&clients); err != nil {
		return nil, err
	}

	r, err := q.Tag("json").Eval(clients)
	if err != nil {
		return nil, err
	}

	return r.(Clients), nil
}

func (ds *datastore) FindClient(field string, value interface{}) (*Client, error) {
	all, err := ds.AllClients()
	if err != nil {
		return nil, err
	}

	r, err := qry.Eq(field, value).Eval(all)
	if err != nil {
		return nil, err
	}

	if len(r.(Clients)) > 0 {
		return r.(Clients)[0], nil
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
