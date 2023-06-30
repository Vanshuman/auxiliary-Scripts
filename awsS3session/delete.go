package awsS3session

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func DeleteS3obj(objKey []string) {

	// Create an S3 service client
	sess, _ := CreateSession()
	s3Svc := s3.New(sess)

	// Create the input parameter structure
	for _, o := range objKey {
		input := &s3.DeleteObjectInput{
			Bucket: aws.String("repute-junk"),
			Key:    aws.String(o),
		}

		// Delete the object
		_, err := s3Svc.DeleteObject(input)
		if err != nil {
			fmt.Println("Failed to delete object:", err)
			os.Exit(1)
		}
		fmt.Println("Object deleted successfully", o)
	}

	fmt.Println("Object deleted successfully!")
}
