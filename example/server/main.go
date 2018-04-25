package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mickep76/encdec"
	_ "github.com/mickep76/encdec/json"
	"github.com/mickep76/kvstore"
	_ "github.com/mickep76/kvstore/etcdv3"

	"github.com/mickep76/kvstore/example/models"
)

type Handler struct {
	ds models.Datastore
}

var clientHandler = kvstore.WatchHandler(func(kv kvstore.KeyValue) {
	log.Printf("client event: %s key: %s", kv.Event().Type, kv.Key())

	c := &models.Client{}
	if err := kv.Decode(c); err != nil {
		log.Print(err)
		return
	}

	log.Printf("client created: %s uuid: %s hostname: %s", c.Created, c.UUID, c.Hostname)
})

func (h *Handler) allClients(w http.ResponseWriter, r *http.Request) {
	v, err := h.ds.AllClients()
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	b, _ := encdec.ToBytes("json", v, encdec.WithIndent("  "))
	w.Write(b)
}

func (h *Handler) allServers(w http.ResponseWriter, r *http.Request) {
	v, err := h.ds.AllServers()
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	b, _ := encdec.ToBytes("json", v, encdec.WithIndent("  "))
	w.Write(b)
}

func main() {
	usage := `client

Usage:
  client [--backend=<backend>] [--prefix=<prefix>] [--endpoints=<endpoints>] [--timeout=<seconds>] [--keepalive=<seconds>] [--bind=<address>]
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

	// Get keepalive.
	keepalive, err := strconv.Atoi(args["--keepalive"].(string))
	if err != nil {
		log.Fatalf("strconv: %v", err)
	}

	// Get prefix.
	prefix := args["--prefix"].(string)

	// Connect to etcd.
	log.Printf("connect to etcd")
	ds, err := models.NewDatastore("etcdv3", strings.Split(args["--endpoints"].(string), ","), keepalive, kvstore.WithTimeout(timeout), kvstore.WithEncoding("json"), kvstore.WithPrefix(prefix))
	if err != nil {
		log.Fatal(err)
	}

	// Create new server struct.
	log.Printf("create new server struct")
	hostname, _ := os.Hostname()
	s := models.NewServer(hostname, args["--bind"].(string))

	// Set client in etcd.
	log.Printf("create server in etcd")
	if err := ds.CreateServer(s); err != nil {
		log.Fatal(err)
	}

	// Create lease keepalive.
	log.Printf("create lease keepalive")
	ch, err := ds.Lease().KeepAlive()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			l := <-ch
			log.Print("send keepalive for lease")
			if l.Error != nil {
				log.Print(l.Error)
			}
		}
	}()

	// Create client watch.
	log.Printf("create client watch")
	go func() {
		if err := ds.Watch(fmt.Sprintf("%s/%s", prefix, "clients")).AddHandler(clientHandler).Start(); err != nil {
			log.Fatal(err)
		}
	}()

	// Create new router.
	log.Printf("create http router")
	router := mux.NewRouter()
	h := &Handler{ds: ds}

	// Client handlers.
	log.Printf("add route /api/clients")
	router.HandleFunc("/api/clients", h.allClients).Methods("GET")

	// Server handlers.
	log.Printf("add route /api/servers")
	router.HandleFunc("/api/servers", h.allServers).Methods("GET")

	// Start https listener.
	log.Printf("start http listener")
	logr := handlers.LoggingHandler(os.Stdout, router)
	if err := http.ListenAndServe(args["--bind"].(string), logr); err != nil {
		log.Fatal("http listener:", err)
	}
}
