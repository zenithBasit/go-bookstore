package main

import (
	"log"
	"net/http"

	"bookstore/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9000", r))
}
