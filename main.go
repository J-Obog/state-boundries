package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/J-Obog/state-boundries/api"
	"github.com/gorilla/mux"
)

func main() {
	//host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	//create and configure new router
	router := mux.NewRouter()
	router.Use(api.ReqLogger)
	router.Use(api.MimeTypeRes)
	
	//handle request
	router.HandleFunc("/api/borders/{state}", api.GetStateBorders).Methods("GET")
	router.HandleFunc("/api/borders/", api.GetAllStateBorders).Methods("GET")

	//spin up server
	fmt.Printf("Server running on port %s 🚀🚀🚀\n", port)
	http.ListenAndServe(":" + port, router)
}