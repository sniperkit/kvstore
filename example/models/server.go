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

func ServerAll(kvc kvstore.Conn, prefix string) (Servers, error) {
	kvs, err := kvc.Values(fmt.Sprintf("%s/%s", prefix, "servers"))
	if err != nil {
		return nil, err
	}

	servers := Servers{}
	for _, kv := range kvs {
		s := &Server{}
		kv.SetEncoding("json")
		if err := kv.Decode(s); err != nil {
			return nil, err
		}

		servers = append(servers, s)
	}

	return servers, nil
}
