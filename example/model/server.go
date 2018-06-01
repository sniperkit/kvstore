package model

import (
	"fmt"
	"time"

	"github.com/mickep76/kvstore"
	"github.com/mickep76/qry"
	"github.com/pborman/uuid"
)

type Server struct {
	UUID     string     `json:"uuid" kvstore:"unique"`
	Created  time.Time  `json:"created"`
	Updated  *time.Time `json:"updated,omitempty"`
	Hostname string     `json:"hostname" kvstore:"unique"`
	Bind     string     `json:"bind"`
}

type Servers []*Server

func NewServer(hostname string, bind string) *Server {
	return &Server{
		UUID:     uuid.New(),
		Created:  time.Now(),
		Hostname: hostname,
		Bind:     bind,
	}
}

func (ds *Datastore) AllServers() (Servers, error) {
	kvs, err := ds.Values("servers")
	if err != nil {
		return nil, err
	}

	servers := Servers{}
	if err := kvs.Decode(&servers); err != nil {
		return nil, err
	}

	return servers, nil
}

func (ds *Datastore) QueryServers(q *qry.Query) (Servers, error) {
	servers, err := ds.AllServers()
	if err != nil {
		return nil, err
	}

	filtered, err := q.Query(servers)
	if err != nil {
		return nil, err
	}

	return filtered.(Servers), nil
}

func (ds *Datastore) OneServer(uuid string) (*Server, error) {
	kvs, err := ds.Values(fmt.Sprintf("servers/%s", uuid))
	if err != nil {
		return nil, err
	}

	servers := Servers{}
	if err := kvs.Decode(&servers); err != nil {
		return nil, err
	}

	if len(servers) > 0 {
		return servers[0], nil
	}

	return nil, nil
}

func (ds *Datastore) CreateServer(server *Server) error {
	return ds.Set(fmt.Sprintf("servers/%s", server.UUID), server, kvstore.WithLease(ds.lease))
}

func (ds *Datastore) UpdateServer(server *Server) error {
	now := time.Now()
	server.Updated = &now
	return ds.Set(fmt.Sprintf("servers/%s", server.UUID), server, kvstore.WithLease(ds.lease))
}
