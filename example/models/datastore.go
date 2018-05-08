package models

import (
	"github.com/mickep76/kvstore"
	_ "github.com/mickep76/kvstore/etcdv3"
	"github.com/mickep76/qry"
)

type Datastore interface {
	AllClients() (Clients, error)
	QueryClients(qry.Query) (Clients, error)
	CreateClient(client *Client) error
	AllServers() (Servers, error)
	QueryServers(qry.Query) (Servers, error)
	CreateServer(server *Server) error
}

type datastore struct {
	lease kvstore.Lease
	kvstore.Conn
}

func NewDatastore(driver string, endpoints []string, keepalive int, options ...func(kvstore.Driver) error) (*datastore, error) {
	c, err := kvstore.Open(driver, endpoints, options...)
	if err != nil {
		return nil, err
	}

	l, err := c.Lease(keepalive)
	if err != nil {
		return nil, err
	}

	return &datastore{
		lease: l,
		Conn:  c,
	}, nil
}

func (ds *datastore) Lease() kvstore.Lease {
	return ds.lease
}
