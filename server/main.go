package main

import (
	"log"
	"os"
	"strconv"

	// Loading all .env file contents automatically
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Referenced from github.com/joho/godotenv/autoload
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// s3Bucket := os.Getenv("S3_BUCKET")
	// secretKey := os.Getenv("SECRET_KEY")

	// Pull down env vars
	portAsStr := os.Getenv("PORT")

	// Initialize a Server object.
	port, err := strconv.Atoi(portAsStr)
	if err != nil {
		log.Fatalf("Failed to convert PORT env var to int: %v\n", err)
	}
	server := NewServer(port)
	server.Start()
}
