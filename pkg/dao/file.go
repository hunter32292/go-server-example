package dao

import (
	"io/ioutil"
	"os"
)

// FileLoadInData - Load in the Data from a file as a stream of bytes
func FileLoadInData(path string) ([]byte, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
