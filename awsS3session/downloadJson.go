package awsS3session

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"path/filepath"
)

func DownloadFile(bucketName, objPath, destinationPath string, sess *session.Session) error {

	// Create an S3 service client
	svc := s3.New(sess)

	// Set the bucket name and file key
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objPath),
	}

	// Download the file from S3
	result, err := svc.GetObject(input)
	if err != nil {
		return err
	}
	defer result.Body.Close()

	// Create the destination directory if it doesn't exist
	err = os.MkdirAll(filepath.Dir(destinationPath), os.ModePerm)
	if err != nil {
		return err
	}

	// Create the destination file
	file, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the file contents to the destination file
	_, err = file.ReadFrom(result.Body)
	if err != nil {
		return err
	}

	fmt.Println("File downloaded successfully.")
	return nil
}
