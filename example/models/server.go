package models

import (
	"fmt"
	"time"

	"github.com/mickep76/kvstore"
	"github.com/pborman/uuid"
)

type Server struct {
	UUID     string    `json:"uuid"`
	Created  time.Time `json:"created"`
	Hostname string    `json:"hostname"`
	Bind     string    `json:"bind"`
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
	kvs, err := ds.Values(fmt.Sprintf("%s/%s", ds.prefix, "servers"))
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

func (ds *datastore) CreateServer(server *Server) error {
	return ds.Set(fmt.Sprintf("%s/servers/%s", ds.prefix, server.UUID), server, kvstore.WithLease(ds.lease))
}
