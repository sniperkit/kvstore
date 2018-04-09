package main

import (
	"log"
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
  --prefix=<prefix>                     Key/value store prefix. [default: /dock2box]
  --endpoints=<endpoints>               Comma-delimited list of hosts in the key/value store cluster. [default: 127.0.0.1:2379]
  --timeout=<seconds>                   Connection timeout for key/value cluster in seconds. [default: 5]
  --keepalive=<seconds>                 Connection keepalive for key/value cluster in seconds. [default: 60]
  --bind=<address>                      Bind to address. [default: 0.0.0.0:8080]
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
	kvc, err := kvstore.Open("etcdv3", strings.Split(args["--endpoints"].(string), ","), kvstore.Timeout(timeout))
	if err != nil {
		log.Fatal(err)
	}

	// Create host watch.
	go func() {
		if err := kvc.Watch(fmt.Sprintf("%s/%s", prefix, "hosts")).AddHandler(hostHandler).Start(); err != nil {
			log.Fatal(err)
		}
	}()

	// Create new router.
	router := mux.NewRouter()

	// Host handlers.
	router.Handle("/api/hosts", AllHosts).Methods("GET")

	// Start https listener.
	logr := handlers.LoggingHandler(os.Stdout, router)
	if err := http.ListenAndServeTLS(args["--bind"].(string), args["--cert"].(string), args["--key"].(string), logr); err != nil {
		log.Fatal("https listener:", err)
	}
}
