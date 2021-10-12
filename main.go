package main

import (
	"fmt"
	"net/http"

	"github.com/J-Obog/state-boundries/api"
	"github.com/gorilla/mux"
)

func main() {
	port := ":80"

	//create new router
	router := mux.NewRouter()
	
	//handle request
	router.HandleFunc("/api/borders/{state}", api.GetStateBorders).Methods("GET")
	
	//spin up server
	fmt.Printf("Server running on port %s 🚀🚀🚀", port)
	http.ListenAndServe(port, router)
}