package file

import "path/filepath"

func GetFullPath(filePath string, fileName string) string {
	return filepath.Join(filePath, fileName)
}

func IsExistFile(fullPathToFile string) bool {
	_, err := ReadFile(fullPathToFile)
	if err != nil {
		return false
	}

	return true
}
