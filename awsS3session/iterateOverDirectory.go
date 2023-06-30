package awsS3session

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetDirectoryNames() []string {
	rootDir := "/Users/anshuman/GolandProjects/awsS3update/HelloWorld/"
	// Specify the root directory here
	var objPaths []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {

			str := strings.TrimPrefix(path, "/Users/anshuman/GolandProjects/awsS3update/awsS3session/HelloWorld/")
			if str != "" {
				objPaths = append(objPaths, str)
				fmt.Println("Directory:", str)
			}

		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
	return objPaths
}

func GetKeysForAttributeOld(dirPats []string) []string {
	var objPaths []string
	for _, o := range dirPats {
		str := fmt.Sprintf("repute-junk/helloworld/'HW Onroll Employee KYC Form'/%s/attributes-old.json", o)
		objPaths = append(objPaths, str)
	}
	return objPaths
}
