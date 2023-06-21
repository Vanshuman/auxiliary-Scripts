package main

import (
	"awsS3update/awsS3session"
	hello_world "awsS3update/awsS3session/hello-world"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"strings"
)

var (
	bucket  = "repute-junk"
	key     = "helloworld/"
	txtFile = "/Users/anshuman/GolandProjects/awsS3update/awsS3session/hello-world/hello.txt"
)
var sess *session.Session

func main() {
	if err := hello_world.ModifyTheFiles(txtFile); err != nil {
		fmt.Println("Error While Modifying the Json FILES: ", err.Error())
	}
	if err := manageAws(); err != nil {
		fmt.Println("Error While Uploading the files to s3 :: ", err.Error())
	}
}
func manageAws() error {
	//repute-junk/helloworld/'HW Onroll Employee KYC Form'/
	var err error
	sess, err = awsS3session.CreateSession()
	if err != nil {
		fmt.Println("Error While Creating a new session :: ", err.Error())
		return err
	}
	objPaths, err2 := awsS3session.ListObjects(bucket, key, sess)
	if err2 != nil {
		fmt.Println("Error while enlisting the objects:: ", err.Error())
		return err
	}
	if err3 := uploadAttributes(objPaths); err3 != nil {
		fmt.Println("Error while uploading the file :: ", err3.Error())
		return err
	}
	return nil
}
func uploadAttributes(objPaths []string) error {
	orderIds := strings.Join(hello_world.ReadFromTxtFile(txtFile), "")
	for _, v := range objPaths {
		folders := strings.Split(v, "/")
		if strings.Contains(orderIds, folders[2]) {
			dest := "/Users/anshuman/GolandProjects/awsS3update/awsS3session/hello-world/"
			//Uploading attributes.json
			destAttributes := fmt.Sprintf(dest+"%s/%s", folders[2], "attributes.json")
			if err := awsS3session.UploadFile(bucket, destAttributes, v, sess); err != nil {
				return err
			}
			//Uploading attributes-old.json
			destAttributesOld := fmt.Sprintf(dest+"%s/%s", folders[2], "attributes-old.json")
			if err := awsS3session.UploadFile(bucket, destAttributesOld, strings.TrimRight(v, "attributes.json")+"attributes-old.json", sess); err != nil {
				return err
			}
		}
	}
	return nil
}
