package tests

import (
	"fmt"
	fl "github.com/ssh-connection-manager/file"
	"math/rand"
	"os/user"
	"testing"
)

func randomString(count int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomStr := make([]rune, count)
	for i := range randomStr {
		randomStr[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(randomStr)
}

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	return usr.HomeDir + "qweeqwweqwqwe" + "/test/"
}

func TestCreateFile(t *testing.T) {
	homeDir := getHomeDir()
	randomStr := randomString(5)

	file := fl.File{Name: randomStr, Path: homeDir}
	err := file.CreateFile()
	if err != nil {
		t.Fatal("Error creating file")
	}

	if !file.IsExistFile() {
		t.Fatal("File dont created")
	}
}

func TestWriteToFile(t *testing.T) {
	randomStr := randomString(5) + ".txt"
	homeDir := getHomeDir()

	file := fl.File{Name: randomStr, Path: homeDir}
	err := file.CreateFile()
	if err != nil {
		t.Fatal("Error creating file")
	}

	randomText := randomString(100)
	err = file.WriteFile([]byte(randomText))
	if err != nil {
		t.Fatal("Error write to file")
	}

	fileText, err := file.ReadFile()
	if err != nil {
		t.Fatal("Error read file")
	}

	if fileText != randomText {
		t.Fatal("Error random text != text from file")
	}
}

func TestReadFile(t *testing.T) {
	files := [7]fl.File{
		{Name: randomString(5) + ".json", Path: getHomeDir()},
		{Name: randomString(5) + ".txt", Path: getHomeDir()},
		{Name: randomString(5) + ".PNG", Path: getHomeDir()},
		{Name: randomString(5) + ".PDF", Path: getHomeDir()},
		{Name: randomString(5) + ".PDF", Path: getHomeDir()},
		{Name: randomString(5) + ".DOC", Path: getHomeDir()},
		{Name: randomString(5), Path: getHomeDir()},
	}

	for _, file := range files {
		err := file.CreateFile()
		if err != nil {
			t.Fatal("Error creating file")
		}

		randomText := randomString(100)

		err = file.WriteFile([]byte(randomText))
		if err != nil {
			t.Fatalf("Error write to file %s", file)
		}

		fileText, err := file.ReadFile()
		if err != nil {
			t.Fatal("Error read file")
		}

		if fileText != randomText {
			t.Fatalf("Error random text != text from file - is file %s", file.Path+file.Name)
		}
	}
}
