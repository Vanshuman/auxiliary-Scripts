package awsS3session

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func UploadFile(bucketName, filePath, objPath string, sess *session.Session) error {

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create an S3 service client
	svc := s3.New(sess)

	// Set the bucket name and file key
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objPath),
		Body:   file,
	}

	// Upload the file to S3
	_, err = svc.PutObject(input)
	if err != nil {
		return err
	}

	fmt.Println("File uploaded successfully.")
	fmt.Println(filePath)
	return nil
}
