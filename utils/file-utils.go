package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

type Project struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}

// CreateTempFile creates a temporary file to be opened in the editor
func CreateTempFile() *os.File {
	file, err := os.CreateTemp(os.TempDir(), "*")
	if err != nil {
		log.Fatalf("Unable to create new file: %v\n", err)
	}
	return file
}

// ReadFile returns the contents of the temp file as a string of bytes
func ReadFile(file *os.File) ([]byte, error) {
	bytes, err := os.ReadFile(file.Name())
	if err != nil {
		return []byte(""), errors.Wrap(err, fmt.Sprintf("Unable to read temp file: %s\n", file.Name()))
	}
	return bytes, nil
}

func ReadProjJson() ([]Project, error) {
	bytes, err := os.ReadFile("example_config.json")
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read proj.json")
	}
	var projectList []Project
	MyJson := []byte(bytes)
	err = json.Unmarshal(MyJson, &projectList)
	if err != nil {
		log.Fatalln(err)
	}
	return projectList, nil
}

func WriteProjJson(proj []byte, path string) error {
	newJson, err := json.Marshal(proj)
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%v/config.json", path), newJson, 0777)
	return err
}
