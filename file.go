package file

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	Path string
	Name string
}

func (fl *File) IsExistFile() bool {
	_, err := fl.ReadFile()
	if err != nil {
		return false
	}

	return true
}

func (fl *File) fullPath() string {
	return filepath.Join(fl.Path, fl.Name)
}

func (fl *File) CreateFile() error {
	file := fl.fullPath()

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		dir := filepath.Dir(file)

		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}

		createdFile, err := os.Create(file)
		if err != nil {
			return err
		}

		defer func(createdFile *os.File) {
			err = createdFile.Close()
		}(createdFile)
	}

	return nil
}

func (fl *File) ReadFile() (string, error) {
	file := fl.fullPath()

	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		err = f.Close()
	}(f)

	fContent, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(fContent), nil
}

func (fl *File) WriteFile(rowData []byte) error {
	file := fl.fullPath()

	err := ioutil.WriteFile(file, rowData, 0)
	if err != nil {
		return err
	}

	return nil
}
