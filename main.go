package main

import (
	"log"
	"net/http"
	"sport-venue-management/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.UserDetailRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8001", r))
}
