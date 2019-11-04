package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file not found")
	}
}
func main() {
	r := NewRouter()
	port := os.Getenv("Port")
	if port == "" {
		port = "8080"
	}
	host := os.Getenv("Host")
	log.Printf("Server started!")
	log.Fatal(http.ListenAndServe(host+":"+port, r))
}