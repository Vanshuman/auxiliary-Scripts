package awsS3session

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"strconv"
	"strings"
)

func ListObjects(bucketName, folderPath string, sess *session.Session) ([]string, error) {

	svc := s3.New(sess)
	// Set the bucket name
	i := 0
	var attributeFiles []string
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}
	for {

		result, err := svc.ListObjectsV2(input)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}

		input.ContinuationToken = result.NextContinuationToken
		fmt.Println("Result size ::", len(result.Contents))
		for _, item := range result.Contents {
			if strings.Contains(*item.Key, "HW Onroll Employee KYC Form") && strings.Contains(*item.Key, "attributes.json") {
				attributeFiles = append(attributeFiles, *item.Key)
			}
		}
		if *result.IsTruncated {
			fmt.Println("Result is truncated :: ", strconv.Itoa(i))
		}
		if !*result.IsTruncated {
			fmt.Printf("Stopped at iteration %d", i+1)
			break
		}
		i++
	}

	// List the objects in the bucket

	// Print the object names

	fmt.Println(len(attributeFiles))
	return attributeFiles, nil
}
