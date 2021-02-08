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

	// Pull down env vars
	s3Bucket := os.Getenv("BUCKET")
	// secretKey := os.Getenv("SECRET_KEY")
	bucketRegion := os.Getenv("REGION")
	portAsStr := os.Getenv("PORT")

	// Get a new Bucket object to interact with our S3 bucket through.
	bucket, err := NewBucket(s3Bucket, bucketRegion)
	if err != nil {
		log.Fatalf("Failed to create a new Bucket object: %v\n", err)
	}

	// Initialize a Server object.
	port, err := strconv.Atoi(portAsStr)
	if err != nil {
		log.Fatalf("Failed to convert PORT env var to int: %v\n", err)
	}
	server := NewServer(port, bucket)
	server.Start()
}
