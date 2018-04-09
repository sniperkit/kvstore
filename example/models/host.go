package models

import (
	"time"

	"github.com/pborman/uuid"
)

type Host struct {
	UUID     string    `json:"uuid"`
	Created  time.Time `json:"created"`
	Hostname string    `json:"hostname"`
}

func NewHost(hostname string) *Host {
	return &Host{
		UUID:     uuid.New(),
		Created:  time.Now(),
		Hostname: hostname,
	}
}
