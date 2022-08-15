package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/go-stockapi/router"
)

func main() {
	r := router.Router()
	fmt.Print("Starting Server on port 8000.....")

	log.Fatal(http.ListenAndServe(":8000", r))
}