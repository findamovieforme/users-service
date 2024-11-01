package main

import (
	"log"
	"net/http"

	"github.com/movierecuh/users-service/helpers"
	"github.com/movierecuh/users-service/routers"
)

func main() {
	helpers.InitEnv()
	router := routers.InitRouter()
	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
