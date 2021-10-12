package main

import (
	"fmt"
	"net/http"

	"github.com/J-Obog/state-boundries/api"
	"github.com/gorilla/mux"
)

func main() {
	host := "localhost"
	port := "80"
	baseUri := fmt.Sprintf("%s:%s", host, port)

	//create new router
	router := mux.NewRouter()
	
	//handle request
	router.HandleFunc("/api/borders/{state}", api.GetStateBorders).Methods("GET")
	
	//spin up server
	fmt.Printf("Server running on port %s 🚀🚀🚀", port)
	http.ListenAndServe(baseUri, router)
}