package awsS3session

import (
	"encoding/json"
	"os"
)

type Attributes struct {
	Documents         map[string]interface{} `json:"documents"`
	AdditionalDetails *interface{}           `json:"additional_details"`
	MetaData          *interface{}           `json:"meta_data"`
	Version           *interface{}           `json:"version"`
}
type AttrMap struct {
	Attr    map[string]interface{} `json:"attributes"`
	DBcheck bool                   `json:"dbcheckdone"`
	source  string                 `json:"source"`
}

func InitAttributes(file string) (*Attributes, error) {
	file1, _ := os.Open(file)
	attr := new(Attributes)
	if err := json.NewDecoder(file1).Decode(attr); err != nil {
		return nil, err
	}

	defer file1.Close()
	return attr, nil
}
