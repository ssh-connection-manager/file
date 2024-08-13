package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateFile(filePath string) {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		dir := filepath.Dir(filePath)

		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}

		createdFile, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
		}

		defer createdFile.Close()
	}
}

func ReadFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(fContent), nil
}

func WriteFile(path string, rowData []byte) error {
	err := ioutil.WriteFile(path, rowData, 0)
	if err != nil {
		return errors.New("error when writing to a file")
	}

	return nil
}
