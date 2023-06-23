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

func FilterFiles() []string {
	//Iterate over the files and Filter and write to text file.
	directory := "/Users/anshuman/GolandProjects/awsS3update/Hello-World/"
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	var orderIds []string
	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		filePath := filepath.Join(directory, file.Name())

		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Printf("Failed to read file %s: %v", filePath, err)
			continue
		}
		attributes := Attributes{}
		if err := json.Unmarshal(content, &attributes); err != nil {
			fmt.Println("Error in Parsing the files", err.Error())
		}
		if _, ok := attributes.Documents["VOTERID"]; ok {
			orderIds = append(orderIds, strings.TrimRight(file.Name(), ".json"))
		}
	}
	return orderIds
}
func WriteToTxtFile(objPaths []string) {
	filePath := "/Users/anshuman/GolandProjects/awsS3update/awsS3session/hello-world/hello.txt"

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
