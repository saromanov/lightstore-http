package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/saromanov/lightstore/store"
)

var port = flag.Int("port", 8080, "port")

var ls *store.Lightstore

func set(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Set new key")
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get value by the key")
}

func main() {
	ls = store.Open(nil)
	http.HandleFunc("/set", set)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
