package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Saikrishna-8/go-postgres/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting the server on the port 8089....")

	log.Fatal(http.ListenAndServe(":8089", r))
}
