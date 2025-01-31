package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/AbhishekCS3459/goGraph/api"
)

func main() {
	r := mux.NewRouter()

	// pass r to handler funciton
	r.HandleFunc("/", api.HealthCheck).Methods("GET")
	r.HandleFunc("/find-paths", api.FindPathsHandler).Methods("POST")
	// start the server
	fmt.Println("Server is started in port 8080")
	http.ListenAndServe(":8080", r)

}