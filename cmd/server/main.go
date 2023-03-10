package main

import (
	"github.com/Eitol/newsletter-api/pkg/newsletter/delivery/httpserver"
	"log"
	"os"
	"strconv"
)

// @contact.name                Grupo MContigo
// @title                       Newsletter API
// @version                     1.0
// @description                 Newsletter API
func main() {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Error converting port to int: %v", err)
	}
	httpserver.RunHttpServer(port)
}
