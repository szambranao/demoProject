package main

import (
	"fmt"

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

func (b *Bucket) CollectList() ([]string, error) {

	finalList := make([]string, 0)

	resp, err := b.svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: &b.name})
	if err != nil {
		return nil, fmt.Errorf("Failed to list all objects in the S3 bucket: %v", err)
	}

	for _, item := range resp.Contents {
		fileURL := bucketPrefix + b.name + bucketSuffix + *item.Key
		finalList = append(finalList, fileURL)
	}

	return finalList, nil
}
