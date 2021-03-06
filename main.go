package main

import (
	"envelope/middleware"
	"envelope/models"
	"envelope/router"
	"fmt"
	"log"
	"net/http"
)

// TODO user input validation
// TODO CSRF for all POST requests
func main() {
	middleware.InitDb(false)
	fmt.Println("Db initialized")
	r := router.Router()
	models.StartTokenCleanup()
	log.Fatal(http.ListenAndServe(":8080", r))
}

