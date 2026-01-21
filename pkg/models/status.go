package models

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"os"
// )

// Top level fields from status json
type Status struct {
	Client  *ClientStatus  `json:"client"`
	Cluster *ClusterStatus `json:"cluster"`
}

// const RelativeJsonFileLocation = "test/data"
//
// func GetStatusFromFile(fileName string) (*FullStatus, error) {
// 	var status FullStatus
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get working directory")
// 	}
// 	testFilePath := fmt.Sprintf("%s/../%s/%s", wd, RelativeJsonFileLocation, fileName)
// 	f, err := os.Open(testFilePath)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open test file %s", testFilePath)
// 	}
// 	defer f.Close()
//
// 	jsonBytes, err := io.ReadAll(f)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read test file %s", testFilePath)
// 	}
// 	err = json.Unmarshal(jsonBytes, &status)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal test file %s", testFilePath)
// 	}
// 	return &status, nil
// }
