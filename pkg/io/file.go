/*
Copyright © 2024 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"os"

	constants "github.com/grayson40/daw/constants"
)

// Creates an empty file with inputted file name
func CreateEmptyFile(fileName string) {
	emptyBytes := []byte("")
	err := os.WriteFile(fileName, emptyBytes, 0644)
	if err != nil {
		panic(err)
	}
}

// Returns true if the working directory is an initialized repository
func IsInitialized() bool {
	if _, err := os.Stat(constants.DawDir); os.IsNotExist(err) {
		return false
	}
	return true
}

// Returns true if the inputted file exists
func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

// Clears file contents
func ClearFile(fileName string) {
	emptyBytes := []byte("")
	err := os.WriteFile(fileName, emptyBytes, 0644)
	if err != nil {
		panic(err)
	}
}

// Open file and return pointer
func OpenFile(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

// Open file and return contents
func ReadFile(fileName string) []byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

// Write to file
func WriteToFile(fileName string, data []byte) {
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		panic(err)
	}
}

// Make directory
func MakeDir(dirName string) {
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		panic(err)
	}
}
