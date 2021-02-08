package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Setup prefix and suffix for AWS S3 bucket
const (
	bucketPrefix = "https://"
	bucketSuffix = ".s3.amazonaws.com/"
)

// AWS S3 bucket struct definition
type Bucket struct {
	name   string
	region string
	svc    *s3.S3 //svc is a type reference to an AWS client
}

// New AWS S3 bucket builder function
func NewBucket(name, region string) (*Bucket, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to establish a new AWS session: %v", err)
	}

	svc := s3.New(sess)
	return &Bucket{name, region, svc}, nil
}

type totalJSON struct {
	Title string
	Tasks []TaskType
}

// Sets up Task list struct
type TaskType struct {
	Title string
}

// Collects list of tasks from S3 Bucket
func (b *Bucket) CollectList() (totalJSON, error) {

	finalList := make([]string, 0)

	resp, err := b.svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: &b.name})
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range resp.Contents {
		fileURL := bucketPrefix + b.name + bucketSuffix + *item.Key
		finalList = append(finalList, fileURL)
	}

	listResp, err := http.Get(finalList[0])
	if err != nil {
		log.Fatal(err)
	}

	var listOfTasks totalJSON
	err = json.NewDecoder(listResp.Body).Decode(&listOfTasks)
	if err != nil {
		log.Fatal(err)
	}

	return listOfTasks, nil
}
