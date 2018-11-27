package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.Int("port", 8080, "port")

func set(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Set new key")
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get value by the key")
}

func main() {
	http.HandleFunc("/set", set)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
