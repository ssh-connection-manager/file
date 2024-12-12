package tests

import (
	fl "github.com/ssh-connection-manager/file"
	"testing"
)

func TestCreateFile(t *testing.T) {
	testDir := GetDirForTests()

	randomStr := RandomString(5)

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
	testDir := GetDirForTests()

	randomStr := RandomString(5) + ".txt"

	file := fl.File{Name: randomStr, Path: testDir}

	err := file.CreateFile()
	if err != nil {
		t.Fatal("Error creating file")
	}

	if !file.IsExistFile() {
		t.Fatal("File dont created")
	}

	randomText := RandomString(100)
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
		{Name: RandomString(5) + ".json", Path: GetDirForTests()},
		{Name: RandomString(5) + ".txt", Path: GetDirForTests()},
		{Name: RandomString(5) + ".PNG", Path: GetDirForTests()},
		{Name: RandomString(5) + ".PDF", Path: GetDirForTests()},
		{Name: RandomString(5) + ".PDF", Path: GetDirForTests()},
		{Name: RandomString(5) + ".DOC", Path: GetDirForTests()},
		{Name: RandomString(5), Path: GetDirForTests()},
	}

	for _, file := range files {
		err := file.CreateFile()
		if err != nil {
			t.Fatal("Error creating file")
		}

		if !file.IsExistFile() {
			t.Fatal("File dont created")
		}

		randomText := RandomString(100)

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
	testDir := GetDirForTests()

	randomStr := RandomString(5)
	randomStr2 := RandomString(5)

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
