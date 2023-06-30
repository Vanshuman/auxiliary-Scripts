package main

import (
	"awsS3update/awsS3session"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"strings"
)

var (
	bucket  = "repute-hrms-greytip"
	key     = "\"HW Onroll Employee KYC Form\""
	txtFile = "/Users/anshuman/GolandProjects/awsS3update/awsS3session/VID-Self-NotInAdd.txt"
)
var sess *session.Session

func main() {
	//if err := modifyAttributes(); err != nil {
	//	fmt.Println(err.Error())
	//}
	//if err := manageAws(); err != nil {
	//	fmt.Println(err.Error())
	//}
	awsS3session.FilterFiles()
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
	if err3 := uploadAttributes(objPaths, sess); err3 != nil {
		fmt.Println("Error while uploading the file :: ", err3.Error())
		return err
	}
	return nil
}
func uploadAttributes(objPaths []string, sess *session.Session) error {

	for _, objPath := range objPaths {
		folder := strings.Split(objPath, "/")
		if strings.Contains(strings.Join(awsS3session.ReadFromTxtFile(txtFile), ""), folder[1]) {
			//upload  attributes-old.json
			destOLD := "/Users/anshuman/GolandProjects/awsS3update/awsS3session/FailedVoterID/AddressWVoterID-NEW/" + folder[1] + "/attributes-old.json"
			pathOld := strings.TrimSuffix(objPath, "attributes.json") + "attributes-old.json"
			if err := awsS3session.UploadFile(bucket, destOLD, pathOld, sess); err != nil {
				return err
			}
			//upload attributes.json
			dest := "/Users/anshuman/GolandProjects/awsS3update/awsS3session/FailedVoterID/AddressWVoterID-NEW/" + folder[1] + "/attributes.json"
			if err := awsS3session.UploadFile(bucket, dest, objPath, sess); err != nil {
				return err
			}
		}
	}
	return nil
}
func renameWithoutDeleting() error {
	sess, err := awsS3session.CreateSession()
	if err != nil {
		fmt.Println("Error in Creating the session:: ", err.Error())
		return err
	}
	objPaths, err2 := awsS3session.ListObjects(bucket, key, sess)
	if err2 != nil {
		fmt.Println("Error while enlisting the objects:: ", err2.Error())
		return err2
	}
	fmt.Println("Size of objPaths ::= ", len(objPaths))
	for _, o := range objPaths {
		newKey := fmt.Sprintf(strings.TrimRight(o, "attributes.json")+"%s", "attributes-check.json")
		err := awsS3session.CopyObjToS3(&bucket, &newKey, &o, sess)
		if err != nil {
			fmt.Println("Error in Renaming :: ", err.Error())
			return err
		}
	}
	return nil
}
func listObj() (error, []string) {
	sess, err := awsS3session.CreateSession()
	if err != nil {
		fmt.Println("Error in Creating the session:: ", err.Error())
		return err, nil
	}
	objPaths, err2 := awsS3session.ListObjects(bucket, key, sess)
	if err2 != nil {
		fmt.Println("Error while enlisting the objects:: ", err2.Error())
		return err2, nil
	}
	return nil, objPaths
}
func downLoadObjects() error {
	_, objPaths := listObj()

	for _, o := range objPaths {
		floderStr := strings.Split(o, "/")
		dest := fmt.Sprintf("/Users/anshuman/GolandProjects/awsS3update/HelloWorld/%s/%s", floderStr[1], floderStr[2])
		if err := awsS3session.DownloadFile(bucket, o, dest, sess); err != nil {
			fmt.Println("Errror in downLoading", err.Error())
			return err
		}
	}
	return nil

}
func modifyAttributes() error {
	if err := awsS3session.ModifyTheFiles(txtFile); err != nil {
		fmt.Println("Error in Modifying the files ", err.Error())
		return err
	}
	return nil
}
