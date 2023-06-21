package awsS3session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess *session.Session

func CreateSession() (*session.Session, error) {
	// Specify the AWS Region you want to connect to.
	awsRegion := "ap-south-1" // Replace with your desired AWS Region

	// Create a new session with AWS credentials.
	var err error
	sess, err = session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		return nil, err
	}

	return sess, nil
}
