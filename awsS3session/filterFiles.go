package awsS3session

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func FilterFiles() {
	var orderIds []string
	for _, d := range getAllDirectories() {
		files, err := ioutil.ReadDir(d)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if file.IsDir() {
				continue // Skip directories
			}

			filePath := filepath.Join(d, file.Name())

			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Printf("Failed to read file %s: %v", filePath, err)
				break
			}
			attributes := Attributes{}
			if err := json.Unmarshal(content, &attributes); err != nil {
				fmt.Println("Error in Parsing the files", err.Error())
			}
			if v, ok := attributes.Documents["ADDRESS"]; ok {
				content, err := json.Marshal(v)
				if err != nil {
					fmt.Println("Error in changing into Bytes", err.Error())
				}

				if strings.Contains(string(content), "VOTERID") {
					//if v, ok := attributes.Documents["VOTERID"]; ok {
					//	content, _ := json.Marshal(v)
					//	if strings.Contains(string(content), "self_upload") {
					//		//fmt.Println("Error in changing into Bytes", err.Error())
					fmt.Println(strings.TrimPrefix(d, "/Users/anshuman/GolandProjects/awsS3update/HelloWorld/"))
					orderIds = append(orderIds, strings.TrimPrefix(d, "/Users/anshuman/GolandProjects/awsS3update/HelloWorld/"))
				}
				//	fmt.Println(strings.TrimLeft(d, "/Users/anshuman/GolandProjects/awsS3update/HelloWorld/"))
				//	orderIds = append(orderIds, strings.TrimLeft(d, "/Users/anshuman/GolandProjects/awsS3update/HelloWorld/"))
				//
			}
		}
	}

	WriteToTxtFile(orderIds)
}
func WriteToTxtFile(objPaths []string) {
	filePath := "/Users/anshuman/GolandProjects/awsS3update/awsS3session/AddWVID.txt"
	// Open the file in write mode with read and write permissions
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, o := range objPaths {
		_, err = file.WriteString(o + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	_, _ = file.WriteString("Total Count  " + strconv.Itoa(len(objPaths)))
}
func getAllDirectories() []string {
	//Iterate over the files and Filter and write to text file.
	directory := "/Users/anshuman/GolandProjects/awsS3update/HelloWorld/"
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	var directories []string
	for _, file := range files {
		if file.IsDir() {
			filePath := filepath.Join(directory, file.Name())
			directories = append(directories, filePath)
		}

	}
	return directories
}
