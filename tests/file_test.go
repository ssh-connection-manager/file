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

func getDirForTests() string {
	usr, err := user.Current()

	if err != nil {
		fmt.Println(err)
	}

	return usr.HomeDir + "/test/"
}

func TestCreateFile(t *testing.T) {
	testDir := getDirForTests()

	randomStr := randomString(5)

	file := fl.File{Name: randomStr, Path: testDir}
	err := file.CreateFile()
	if err != nil {
		t.Fatal("Error creating file")
	}

	if !file.IsExistFile() {
		t.Fatal("File dont created")
	}
}

func TestWriteToFile(t *testing.T) {
	testDir := getDirForTests()

	randomStr := randomString(5) + ".txt"

	file := fl.File{Name: randomStr, Path: testDir}

	err := file.CreateFile()
	if err != nil {
		t.Fatal("Error creating file")
	}

	if !file.IsExistFile() {
		t.Fatal("File dont created")
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
		{Name: randomString(5) + ".json", Path: getDirForTests()},
		{Name: randomString(5) + ".txt", Path: getDirForTests()},
		{Name: randomString(5) + ".PNG", Path: getDirForTests()},
		{Name: randomString(5) + ".PDF", Path: getDirForTests()},
		{Name: randomString(5) + ".PDF", Path: getDirForTests()},
		{Name: randomString(5) + ".DOC", Path: getDirForTests()},
		{Name: randomString(5), Path: getDirForTests()},
	}

	for _, file := range files {
		err := file.CreateFile()
		if err != nil {
			t.Fatal("Error creating file")
		}

		if !file.IsExistFile() {
			t.Fatal("File dont created")
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

func TestIsExistFile(t *testing.T) {
	testDir := getDirForTests()

	randomStr := randomString(5)
	randomStr2 := randomString(5)

	file := fl.File{Name: randomStr, Path: testDir}
	fileWithDontExistName := fl.File{Name: randomStr2, Path: testDir}

	err := file.CreateFile()
	if err != nil {
		t.Fatal("Error creating file")
	}

	if !file.IsExistFile() {
		t.Fatal("Created file exists")
	}

	if fileWithDontExistName.IsExistFile() {
		t.Fatal("None create file is exist")
	}
}
