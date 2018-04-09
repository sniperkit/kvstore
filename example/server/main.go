package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mickep76/kvstore"
	_ "github.com/mickep76/kvstore/etcdv3"

	"github.com/mickep76/kvstore/example/models"
)

var (
	kvc    kvstore.Conn
	prefix string
)

var clientHandler = kvstore.WatchHandler(func(kv kvstore.KeyValue) {
	log.Printf("client event: %s key: %s", kv.Event().Type, kv.Key())

	c := &models.Client{}
	kv.SetEncoding("json")
	if err := kv.Decode(c); err != nil {
		log.Print(err)
		return
	}

	log.Printf("client created: %s uuid: %s hostname: %s", c.Created, c.UUID, c.Hostname)
})

var clientAll = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	v, err := models.ClientAll(kvc, prefix)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	var b []byte
	b, _ = json.MarshalIndent(v, "", "  ")

	w.Write(b)
})

func main() {
	usage := `client

Usage:
  client [--backend=<backend>] [--prefix=<prefix>] [--endpoints=<endpoints>] [--timeout=<seconds>] [--bind=<address>]
  client -h | --help
  client --version

Options:
  -h --help                             Show this screen.
  --version                             Show version.
  --backend=<backend>                   Key/value store backend. [default: etcdv3]
  --prefix=<prefix>                     Key/value store prefix. [default: /example]
  --endpoints=<endpoints>               Comma-delimited list of hosts in the key/value store cluster. [default: 127.0.0.1:2379]
  --timeout=<seconds>                   Connection timeout for key/value cluster in seconds. [default: 5]
  --bind=<address>                      Bind to address and port. [default: 127.0.0.1:8080]
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

	// Get prefix.
	prefix = args["--prefix"].(string)

	// Connect to etcd.
	kvc, err = kvstore.Open("etcdv3", strings.Split(args["--endpoints"].(string), ","), kvstore.Timeout(timeout))
	if err != nil {
		log.Fatal(err)
	}

	// Create host watch.
	go func() {
		if err := kvc.Watch(fmt.Sprintf("%s/%s", prefix, "clients")).AddHandler(clientHandler).Start(); err != nil {
			log.Fatal(err)
		}
	}()

	// Create new router.
	router := mux.NewRouter()

	// Host handlers.
	router.Handle("/api/clients", clientAll).Methods("GET")

	// Start https listener.
	logr := handlers.LoggingHandler(os.Stdout, router)
	if err := http.ListenAndServe(args["--bind"].(string), logr); err != nil {
		log.Fatal("http listener:", err)
	}
}
