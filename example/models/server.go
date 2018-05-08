package models

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

func (ds *datastore) AllServers() (Servers, error) {
	kvs, err := ds.Values("servers")
	if err != nil {
		return nil, err
	}

	servers := Servers{}
	return servers, kvs.Decode(&servers)
}

func (ds *datastore) QueryServers(q *qry.Query) (Servers, error) {
	kvs, err := ds.Values("servers")
	if err != nil {
		return nil, err
	}

	servers := Servers{}
	if err := kvs.Decode(&servers); err != nil {
		return nil, err
	}

	r, err := q.Tag("json").Eval(servers)
	if err != nil {
		return nil, err
	}

	return r.(Servers), nil
}

func (ds *datastore) FindServer(field string, value interface{}) (*Server, error) {
	all, err := ds.AllServers()
	if err != nil {
		return nil, err
	}

	r, err := qry.Eq(field, value).Eval(all)
	if err != nil {
		return nil, err
	}

	if len(r.(Servers)) > 0 {
		return r.(Servers)[0], nil
	}

	return nil, nil
}

func (ds *datastore) CreateServer(server *Server) error {
	return ds.Set(fmt.Sprintf("servers/%s", server.UUID), server, kvstore.WithLease(ds.lease))
}

func (ds *datastore) UpdateServer(server *Server) error {
	now := time.Now()
	server.Updated = &now
	return ds.Set(fmt.Sprintf("servers/%s", server.UUID), server, kvstore.WithLease(ds.lease))
}
