package hello_world

import (
	"awsS3update/awsS3session"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Read from hello.txt file
// Go to the json file
// Cast the  data into Attribute object
// Delete the voterId
// Write the modified data and original data into another file

func ModifyTheFiles(txtFile string) error {
	for _, r := range ReadFromTxtFile(txtFile) {
		fileName := strings.Trim(r, "\"")
		dest := fmt.Sprintf("/Users/anshuman/GolandProjects/awsS3update/Hello-World/%s.json", fileName)
		attr, err := awsS3session.InitAttributes(dest)
		if err != nil {
			fmt.Println("Error in converting the json file into Attribute obj :", err.Error())
			return err
		}
		if err1 := writeFiles(*attr, fileName, "attributes-old"); err1 != nil {
			fmt.Println("Problem while writing the attributes-old.json. ", err.Error())
		}
		awsS3session.ModifyAttributesJson(attr, "VOTERID")
		if err1 := writeFiles(*attr, fileName, "attributes"); err1 != nil {
			fmt.Println("Problem while writing the attributes.json. ", err.Error())
		}
	}
	return nil
}

func writeFiles(attr awsS3session.Attributes, orderId, name string) error {
	// Copying the data to attributes_old.json
	attrLoc := fmt.Sprintf("/Users/anshuman/GolandProjects/awsS3update/awsS3session/hello-world/%s/%s.json", orderId, name)
	err := os.MkdirAll(filepath.Dir(attrLoc), os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(attrLoc)
	if err != nil {
		return err
	}
	defer file.Close()

	attrByte, err := json.MarshalIndent(attr, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(attrLoc, attrByte, os.ModePerm)
	if err != nil {
		return err
	}

	fmt.Println("JSON files created successfully.")
	return nil
}

func ReadFromTxtFile(txtFile string) (files []string) {
	file, err := os.ReadFile(txtFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	files = strings.Split(string(file), "\n")
	return
}
