package awsS3session

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CopyObjToS3(bucket, key, source *string, sess *session.Session) error {
	svc := s3.New(sess)
	copyObjInput := s3.CopyObjectInput{
		Bucket:     bucket,
		Key:        key,
		CopySource: source}
	if _, err := svc.CopyObject(&copyObjInput); err != nil {
		fmt.Println("Error in Copying Obj :: ", err.Error())
		return err
	}
	return nil
}
