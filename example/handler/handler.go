package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mickep76/kvstore/example/model"
)

type Handler struct {
	ds *model.Datastore
}

func NewHandler(ds *model.Datastore) *Handler {
	return &Handler{
		ds: ds,
	}
}

func writeError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

func write(w http.ResponseWriter, v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func writeDelete(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
