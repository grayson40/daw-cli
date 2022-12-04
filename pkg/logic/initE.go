package daw

import "os"

func ExecuteInit() {
	// Create .daw directory
	err := os.Mkdir(".daw", 0755)
	if err != nil {
		panic(err)
	}

	// Create json files to store staged, tracked, commits
	createEmptyFile(".daw/staged.json")
	createEmptyFile(".daw/tracked.json")
	createEmptyFile(".daw/commits.json")
}

// Creates an empty file with inputted file name
func createEmptyFile(fileName string) {
	emptyBytes := []byte("")
	err := os.WriteFile(fileName, emptyBytes, 0644)
	if err != nil {
		panic(err)
	}
}

// Returns true if the working directory is an initialized repository
func IsInitialized() bool {
	if _, err := os.Stat("./.daw"); os.IsNotExist(err) {
		return false
	}
	return true
}
