package helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/safe_msvc_user/insfractruture/utils"
)

func UploadFileToS3(bucket, filename string) (string, error) {
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(utils.AWS_REGION))
	if err != nil {
		log.Printf("Error loading SDK config, %v", err)
		return "", err
	}
	client := s3.NewFromConfig(config)

	file, err := os.Open(utils.UPLOADS + filename)
	if err != nil {
		log.Printf("Error opening file %q, %v", filename, err)
		return "", err
	}

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
		ACL:    types.ObjectCannedACLPublicRead, // Puedes ajustar los permisos según sea necesario, por ejemplo, PublicRead o Private.
	})
	publicURL := fmt.Sprintf(utils.AWS_URL_S3, bucket, filename)

	file.Close()
	if err != nil {
		log.Printf("Error uploading file %q to bucket %q, %v", filename, bucket, err)
		return "", err
	} else {
		err := os.Remove(filename)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("File %s deleted successfully", filename)
		}
	}
	log.Printf("File '%s' uploaded to bucket '%s'\n", filename, bucket)
	return publicURL, nil
}

func RemoveFileToS3(bucket, filename string) (string, error) {
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(utils.AWS_REGION))
	if err != nil {
		log.Printf("Error loading SDK config, %v", err)
		return "", err
	}
	client := s3.NewFromConfig(config)
	_, err = client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		log.Printf("Error deleting file %q from bucket %q, %v", filename, bucket, err)
		return "", err
	}
	log.Printf("File '%s' deleted successfully from bucket '%s'\n", filename, bucket)
	return "", nil
}