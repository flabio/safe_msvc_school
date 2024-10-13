package helpers

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsService struct {
	S3Clinet *s3.Client
}

func (awsSvc AwsService) UploadFile(bucketName, bucketKey, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to open file %q, %v", filename, err)
	} else {
		defer file.Close()
		_, err = awsSvc.S3Clinet.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String("t-bucket-golang"),
			Key:    aws.String(bucketKey),
			Body:   file,
		})
		if err != nil {
			log.Fatalf("Unable to upload %q to %q, %v", filename, bucketName, err)
		} else {
			log.Printf("Successfully uploaded %q to %q", filename, bucketName)
		}
	}
	return err
}
