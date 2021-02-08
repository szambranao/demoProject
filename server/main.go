package main

import (
	"log"
	"os"
	"strconv"

	// Loading all .env file contents automatically

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

const bucketRegion = "us-east-1"

func main() {
	// Referenced from github.com/joho/godotenv/autoload
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Pull down env vars
	s3Bucket := os.Getenv("BUCKET")
	// secretKey := os.Getenv("SECRET_KEY")
	// bucketRegion := os.Getenv("REGION")
	portAsStr := os.Getenv("PORT")

	// Get a new Bucket object to interact with our S3 bucket through.
	bucket, err := NewBucket(s3Bucket, bucketRegion)
	if err != nil {
		log.Fatalf("Failed to create a new Bucket object: %v\n", err)
	}
	// fmt.Printf("%+v\n", bucket)

	// TODO: TEMPORARY! JUST FOR TESTING TO SEE IF WE'RE WIRED UP TO THE BUCKET IN THE CLOUD
	// contents, err := bucket.svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: &bucket.name})
	// if err != nil {
	// 	fmt.Errorf("Failed to list all objects in the S3 bucket: %v", err)
	// }

	// resp, err := bucket.svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: &bucket.name})
	// if err != nil {
	// 	fmt.Errorf("Unable to list items in bucket %q, %v", bucket, err)
	// }

	// for _, item := range resp.Contents {
	// 	fmt.Println("Name:         ", *item.Key)
	// 	fmt.Println("Last modified:", *item.LastModified)
	// 	fmt.Println("Size:         ", *item.Size)
	// 	fmt.Println("Storage class:", *item.StorageClass)
	// 	fmt.Println("")

	// 	fmt.Println("https://masdemoproject.s3.amazonaws.com/" + *item.Key)
	// }

	// const fileURL = "https://masdemoproject.s3.amazonaws.com/something.json"
	// response, err := http.Get(fileURL)
	// if err != nil {
	// 	fmt.Errorf("Failed to fetch JSON file from bucket: %v", err)
	// }
	// fmt.Println(response)
	contents, _ := bucket.CollectList()
	log.Printf("%+v\n", contents)

	// Initialize a Server object.
	port, err := strconv.Atoi(portAsStr)
	if err != nil {
		log.Fatalf("Failed to convert PORT env var to int: %v\n", err)
	}
	server := NewServer(port)
	server.Start()
}
