package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/mickep76/kvstore"
	_ "github.com/mickep76/kvstore/etcdv3"

	"github.com/mickep76/kvstore/example/models"
)

func main() {
	usage := `client

Usage:
  client [--backend=<backend>] [--prefix=<prefix>] [--endpoints=<endpoints>] [--timeout=<seconds>] [--keepalive=<seconds>]
  client -h | --help
  client --version

Options:
  -h --help                             Show this screen.
  --version                             Show version.
  --backend=<backend>                   Key/value store backend. [default: etcdv3]
  --prefix=<prefix>                     Key/value store prefix. [default: /example]
  --endpoints=<endpoints>               Comma-delimited list of hosts in the key/value store cluster. [default: 127.0.0.1:2379]
  --timeout=<seconds>                   Connection timeout for key/value cluster in seconds. [default: 5]
  --keepalive=<seconds>                 Connection keepalive for key/value cluster in seconds. [default: 60]
`

	// Parse arguments.
	args, err := docopt.Parse(usage, nil, true, "client 0.0.1", false)
	if err != nil {
		log.Fatalf("parse args: %v", err)
	}

	// Get timeout.
	timeout, err := strconv.Atoi(args["--timeout"].(string))
	if err != nil {
		log.Fatalf("strconv: %v", err)
	}

	// Get keepalive.
	keepalive, err := strconv.Atoi(args["--keepalive"].(string))
	if err != nil {
		log.Fatalf("strconv: %v", err)
	}

	// Get prefix.
	prefix := args["--prefix"].(string)

	// Connect to etcd.
	log.Printf("connect to etcd")
	ds, err := models.NewDatastore("etcdv3", strings.Split(args["--endpoints"].(string), ","), prefix, keepalive, kvstore.WithTimeout(timeout), kvstore.WithEncoding("json"))
	if err != nil {
		log.Fatal(err)
	}

	// Create new client struct.
	log.Printf("create new client struct")
	hostname, _ := os.Hostname()
	c := models.NewClient(hostname)

	// Set client in etcd.
	log.Printf("create client in etcd")
	if err := ds.CreateClient(c); err != nil {
		log.Fatal(err)
	}

	// Create lease keepalive.
	log.Printf("create lease keepalive")
	ch, err := ds.Lease().KeepAlive()
	if err != nil {
		log.Fatal(err)
	}

	for {
		l := <-ch
		log.Print("send keepalive for lease")
		if l.Error != nil {
			log.Print(l.Error)
		}
	}
}
