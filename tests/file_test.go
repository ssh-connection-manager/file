package tests

import (
	fl "github.com/ssh-connection-manager/file"
	"testing"
)

func TestCreateFile(t *testing.T) {
	file := fl.File{Name: "test1", Path: "/home/deniskorbakov"}
	err := file.CreateFile()
	if err != nil {
		t.Fatal("Error creating file")
	}

}

func TestWriteToFile(t *testing.T) {}

func TestReadFile(t *testing.T) {}
