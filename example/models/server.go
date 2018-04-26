package models

import (
	"fmt"
	"time"

	"github.com/mickep76/kvstore"
	"github.com/mickep76/kvstore/query"
	"github.com/pborman/uuid"
)

type Server struct {
	UUID     string     `json:"uuid"`
	Created  time.Time  `json:"created"`
	Updated  *time.Time `json:"updated,omitempty"`
	Hostname string     `json:"hostname"`
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
	for _, kv := range kvs {
		s := &Server{}
		if err := kv.Decode(s); err != nil {
			return nil, err
		}

		servers = append(servers, s)
	}

	return servers, nil
}

func (ds *datastore) FindServer(field string, value interface{}) (*Server, error) {
	servers, err := ds.AllServers()
	if err != nil {
		return nil, err
	}

	for _, s := range servers {
		v, err := query.FieldValue(s, field)
		if err != nil {
			return nil, err
		}
		if fmt.Sprint(value) == fmt.Sprint(v) {
			return s, nil
		}
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
