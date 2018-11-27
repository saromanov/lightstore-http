package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/saromanov/lightstore/store"
)

// KeyModel defines struct for setting key value pair
type KeyModel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var port = flag.Int("port", 8080, "port")

var ls *store.Lightstore

func set(w http.ResponseWriter, r *http.Request) {
	var payload KeyModel
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = ls.Write(func(txn *store.Txn) error {
		err := txn.Set([]byte(payload.Key), []byte(payload.Value))
		if err != nil {
			return err
		}
		return txn.Commit()
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get value by the key")
}

func main() {
	ls = store.Open(nil)
	http.HandleFunc("/set", set)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
