package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/J-Obog/state-boundries/api"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	baseUri := fmt.Sprintf("%s:%s", host, port)
	
	//create and configure new router
	router := mux.NewRouter()
	router.Use(api.ReqLogger)
	router.Use(api.MimeTypeRes)
	
	//handle request
	router.HandleFunc("/api/borders/{state}", api.GetStateBorders).Methods("GET")
	router.HandleFunc("/api/borders/", api.GetAllStateBorders).Methods("GET")

	//spin up server
	fmt.Printf("Server running on port %s 🚀🚀🚀\n", port)
	http.ListenAndServe(baseUri, router)
}