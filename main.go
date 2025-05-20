package main

import (
	"fmt"
	"github.com/joho/godotenv"

	"log"
	"net/http"

	"Clean-Architecture-Backend-Go/src/router"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//MONGO_URL := os.Getenv("MONGO_URL")

	r := router.SetupRouter()
	//r.Use(middleware.Logger)

	fmt.Println("Listening on port 3000")
	fmt.Println("http://localhost:3000/api/notes/")

	http.ListenAndServe(":3000", r)

}
