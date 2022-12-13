/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package daw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/grayson40/daw/types"
)

// Creates and returns a new commit
func newCommit(files []types.File, message string) types.Commit {
	c := types.Commit{
		Files:   files,
		Message: message,
	}
	return c
}

// Executes the commiting process
func ExecuteCommit(message string) {
	// Throw error if not an initialized repo
	if !IsInitialized() {
		fmt.Println("fatal: not a daw repository (or any of the parent directories): .daw")
		return
	}

	// Read staged staged files
	stagedFiles := GetStaged()

	if len(stagedFiles) == 0 {
		fmt.Println("no changes added to commit (use \"daw add\")")
		return
	}

	// Read commit stack
	commits := GetCommits()

	// Create new commit and write to commit stack
	commit := newCommit(stagedFiles, message)
	commits = append([]types.Commit{commit}, commits...)
	writeErr := writeCommit(commits)
	if writeErr != nil {
		panic(writeErr)
	}

	// Clear staged files
	if err := os.Truncate("./.daw/staged.json", 0); err != nil {
		panic(err)
	}
}

// Reads contents of json staged files and returns array of staged files
func GetStaged() []types.File {
	var files []types.File

	jsonFile, err := os.Open("./.daw/staged.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &files)

	return files
}

// Reads contents of json commit file and returns array of commits
func GetCommits() []types.Commit {
	var commits []types.Commit

	jsonFile, err := os.Open("./.daw/commits.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &commits)

	return commits
}

// Writes commit array to json file, returns err
func writeCommit(commits []types.Commit) error {
	file, err2 := json.MarshalIndent(commits, "", "\t")
	if err2 != nil {
		panic(err2)
	}

	err := ioutil.WriteFile("./.daw/commits.json", file, 0644)

	return err
}
